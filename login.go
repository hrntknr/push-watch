package main

import (
	"fmt"
	"net/url"
)

func login(username, password string) error {
	var login struct {
		Status  float64     `json:"status"`
		Request string      `json:"request"`
		ID      string      `json:"id"`
		Secret  string      `json:"secret"`
		Err     interface{} `json:"errors"`
	}
	if err := req(
		"https://api.pushover.net/1/users/login.json",
		"POST",
		url.Values{},
		url.Values{
			"email":    {username},
			"password": {password},
		},
		&login,
	); err != nil {
		return fmt.Errorf("login: %v", err)
	}
	if login.Status != 1 {
		return fmt.Errorf("login: %s", login.Err)
	}

	var register struct {
		Status  float64     `json:"status"`
		Request string      `json:"request"`
		ID      string      `json:"id"`
		Err     interface{} `json:"errors"`
	}
	if err := req(
		"https://api.pushover.net/1/devices.json",
		"POST",
		url.Values{},
		url.Values{
			"secret": {login.Secret},
			"name":   {LoginCmdDeviceName},
			"os":     {"O"},
		},
		&register,
	); err != nil {
		return fmt.Errorf("register: %v", err)
	}
	if register.Status != 1 {
		return fmt.Errorf("register: %s", register.Err)
	}

	fmt.Println("Success!")
	fmt.Println("Device ID:", register.ID)
	fmt.Println("Device Secret:", login.Secret)

	return nil
}
