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

// FromStdin created string representation of JSON object ready for
// uploading to myjson.
func FromStdin() string {
	buffer := make([]byte, 8192)

	builder := strings.Builder{}

	for {
		n, err := os.Stdin.Read(buffer)

		if err != nil {
			if err != io.EOF {
				panic(err)
			}

			break
		}

		builder.WriteString(string(buffer[:n]))
	}

	jsonData := map[string]string{
		"content": builder.String(),
		"created": time.Now().UTC().String(),
	}

	jsonString, _ := json.Marshal(jsonData)

	return string(jsonString)
}
