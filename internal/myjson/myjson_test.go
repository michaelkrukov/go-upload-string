package myjson

import (
	"testing"
	"time"

	"github.com/michaelkrukov/go-upload-string/internal/pkg/payload"
)

func Test_SaveAndLoad(t *testing.T) {
	url := Save(payload.CreatePayload("example string"))

	if len(url) < 1 {
		t.Error("Returned url is incorrect")
		return
	}

	data := Load(url)

	if data.Content != "example string" {
		t.Error("Wrong content after loading data")
		return
	}

	receivedDatetime, _ := time.Parse(time.RFC3339, data.Created)

	d := time.Now().Sub(receivedDatetime)

	if d > 5*time.Second {
		t.Error("Returned bad date")
		return
	}
}

func Test_Load_Error(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Load didn't paniced on empty url")
		}
	}()

	Load("")
}

func Test_Save_Error(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Save didn't paniced on incorrect payload")
		}
	}()

	Save("a")
}
