package main

import (
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

func watch(deviceID string, secret string, cmd []string) error {
	var priority []int
	for _, p := range strings.Split(WatchCmdPriorityFilter, ",") {
		switch p {
		case "-2":
			priority = append(priority, -2)
		case "-1":
			priority = append(priority, -1)
		case "0":
			priority = append(priority, 0)
		case "1":
			priority = append(priority, 1)
		case "2":
			priority = append(priority, 2)
		}
	}

	if _, err := getMessage(deviceID, secret); err != nil {
		return err
	}

	ch := make(chan Message)

	go func() {
		for {
			message := <-ch
			find := false
			for _, p := range priority {
				if message.Priority == float64(p) {
					find = true
					break
				}
			}
			if find {
				if err := execCmd(cmd, message); err != nil {
					log.Print(err)
				}
			}
		}
	}()

	for {
		if err := connWs(deviceID, secret, ch); err != nil {
			return err
		}
		log.Printf("Reconnecting in 5 seconds...")
		time.Sleep(5 * time.Second)
	}
}

func connWs(deviceID string, secret string, ch chan Message) error {
	c, _, err := websocket.DefaultDialer.Dial("wss://client.pushover.net/push", nil)
	if err != nil {
		return err
	}
	defer c.Close()

	done := make(chan error)
	go func() {
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				done <- err
				return
			}
			switch string(message) {
			case "#":
				continue
			case "!":
				m, err := getMessage(deviceID, secret)
				if err != nil {
					log.Printf("getMessage: %v", err)
					done <- nil
					return
				}
				for _, msg := range m {
					ch <- msg
				}
			case "R":
				done <- nil
			case "E":
				done <- fmt.Errorf("permanent error")
			case "A":
				done <- fmt.Errorf("session closed")
			default:
				done <- fmt.Errorf("unknown message: %s", message)
			}
		}
	}()
	if err := c.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("login:%s:%s", deviceID, secret))); err != nil {
		return err
	}
	if err := <-done; err != nil {
		return err
	}
	return nil
}

func getMessage(deviceID string, secret string) ([]Message, error) {
	var messages struct {
		Status   float64     `json:"status"`
		Request  string      `json:"request"`
		Messages []Message   `json:"messages"`
		Err      interface{} `json:"errors"`
	}
	if err := req(
		"https://api.pushover.net/1/messages.json",
		"GET",
		url.Values{
			"secret": {secret},
			"device": {deviceID},
		},
		url.Values{},
		&messages,
	); err != nil {
		return nil, fmt.Errorf("messages: %v", err)
	}
	if messages.Status != 1 {
		return nil, fmt.Errorf("messages: %s", messages.Err)
	}

	if len(messages.Messages) != 0 {
		var updateHighest struct {
			Status  float64     `json:"status"`
			Request string      `json:"request"`
			Err     interface{} `json:"errors"`
		}
		if err := req(
			fmt.Sprintf("https://api.pushover.net/1/devices/%s/update_highest_message.json", deviceID),
			"POST",
			url.Values{},
			url.Values{
				"secret":  {secret},
				"message": {messages.Messages[len(messages.Messages)-1].ID},
			},
			&updateHighest,
		); err != nil {
			return nil, fmt.Errorf("updateHighest: %v", err)
		}
		if updateHighest.Status != 1 {
			return nil, fmt.Errorf("updateHighest: %s", updateHighest.Err)
		}
	}
	return messages.Messages, nil
}
