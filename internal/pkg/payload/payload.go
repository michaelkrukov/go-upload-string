package payload

import (
	"encoding/json"
	"io"
	"os"
	"strings"
	"time"
)

// Data is structure for storing contents of json object stored
// in myjson.
type Data struct {
	Content string `json:"content"`
	Created string `json:"created"`
}

// CreatePayload returns string representation of JSON object ready for
// uploading to myjson with passed string as content.
func CreatePayload(str string) string {
	jsonData := map[string]string{
		"content": str,
		"created": time.Now().UTC().Format(time.RFC3339),
	}

	jsonString, _ := json.Marshal(jsonData)

	return string(jsonString)
}

// FromStdin returns string representation of JSON object ready for
// uploading to myjson with content read from standart input.
func FromStdin() string {
	return FromReader(os.Stdin)
}

// FromReader returns string representation of JSON object ready for
// uploading to myjson with content read from specified io.Reader.
func FromReader(reader io.Reader) string {
	buffer := make([]byte, 8192)

	builder := strings.Builder{}

	for {
		n, err := reader.Read(buffer)

		if err != nil {
			if err != io.EOF {
				panic(err)
			}

			break
		}

		builder.WriteString(string(buffer[:n]))
	}

	return CreatePayload(builder.String())
}
