package myjson

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/michaelkrukov/go-upload-string/internal/pkg/payload"
)

type createdResponse struct {
	URI string `json:"uri"`
}

// Save uploads passed string representation of JSON object to myjson.
func Save(payload string) string {
	resp, err := http.Post(
		"https://api.myjson.com/bins",
		"application/json",
		bytes.NewBuffer([]byte(payload)),
	)

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		err = fmt.Errorf("Status code wanted 200, but got %d", resp.StatusCode)
	}

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	respContent, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	jsonResponse := createdResponse{}

	json.Unmarshal(respContent, &jsonResponse)

	return jsonResponse.URI
}

// Load loads data uploaded to myjson and parses it to `payload.Data`.
func Load(uri string) payload.Data {
	resp, err := http.Get(uri)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	respContent, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	payloadData := payload.Data{}

	json.Unmarshal(respContent, &payloadData)

	return payloadData
}
