package payload

import (
	"errors"
	"strings"
	"testing"
)

func Test_CreatePayload(t *testing.T) {
	str := CreatePayload("123")

	if !strings.HasPrefix(str, `{`) || !strings.HasSuffix(str, `}`) {
		t.Error("CreatePayload: payload has incorrect object format")
	}

	if !strings.Contains(str, `"content":"123"`) {
		t.Error("CreatePayload: payload without correct data")
	}
}

func Test_FromReader(t *testing.T) {
	str := FromReader(strings.NewReader("123"))

	if !strings.HasPrefix(str, `{`) || !strings.HasSuffix(str, `}`) {
		t.Error("FromReader: payload has incorrect object format")
	}

	if !strings.Contains(str, `"content":"123"`) {
		t.Error("FromReader: payload without correct data")
	}
}

type FakeReader struct{}

func (fr FakeReader) Read(b []byte) (int, error) {
	return 0, errors.New("Something bad")
}

func Test_FromReader_WithErrro(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("FromReader didn't paniced on Read error")
		}
	}()

	FromReader(FakeReader{})
}
