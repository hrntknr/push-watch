package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Message struct {
	ID       string  `json:"id_str"`
	UmID     string  `json:"umid_str"`
	Title    string  `json:"title"`
	Message  string  `json:"message"`
	App      string  `json:"app"`
	AID      string  `json:"aid_str"`
	Priority float64 `json:"priority"`
	Acked    float64 `json:"acked"`
}

func req(urlStr string, methods string, urlValues url.Values, values url.Values, res interface{}) error {
	u, err := url.Parse(urlStr)
	if err != nil {
		return err
	}
	u.RawQuery = urlValues.Encode()

	req, err := http.NewRequest(
		methods,
		u.String(),
		strings.NewReader(values.Encode()),
	)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, res); err != nil {
		return err
	}

	return nil
}
