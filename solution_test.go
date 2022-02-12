package solution

import (
	"testing"
)

const testText = "Hello \U0001f5fa\ufe0f!"

func TestGetMessage(t *testing.T) {
	rendered := GetMessage()
	if testText != rendered {
		t.Error("GetMessage", rendered, testText)
	}
}
