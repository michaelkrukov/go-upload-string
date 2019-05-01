package myjson

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/michaelkrukov/go-upload-string/internal/pkg/payload"
)

type createdResponse struct {
	URI string `json:"uri"`
}

func errorAndExit(reason string, code int) {
	fmt.Printf("Error: %s\n", reason)
	os.Exit(code)
}

// Save uploads passed string representation of JSON object to myjson.
func Save(payload string) string {
	resp, err := http.Post(
		"https://api.myjson.com/bins",
		"application/json",
		bytes.NewBuffer([]byte(payload)),
	)

	if err != nil {
		errorAndExit(err.Error(), 1)
	}

	defer resp.Body.Close()

	respContent, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		errorAndExit(err.Error(), 2)
	}

	jsonResponse := createdResponse{}

	json.Unmarshal(respContent, &jsonResponse)

	return jsonResponse.URI
}

// Load loads data uploaded to myjson and parses it to `payload.Data`.
func Load(uri string) payload.Data {
	resp, err := http.Get(uri)

	if err != nil {
		errorAndExit(err.Error(), 1)
	}

	defer resp.Body.Close()

	respContent, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		errorAndExit(err.Error(), 2)
	}

	payloadData := payload.Data{}

	json.Unmarshal(respContent, &payloadData)

	return payloadData
}
