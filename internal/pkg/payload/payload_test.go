package payload

import (
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
