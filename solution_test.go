package solution

import (
	"testing"
)

const (
	worldKey  = ":world_map:"
	worldText = "Hello "
)

var testText = worldText + "\U0001f5fa\ufe0f!"

func TestMessage(t *testing.T) {
	rendered := GetMessage(worldText, worldKey)
	if testText != rendered {
		t.Error("GetMessage", rendered, testText)
	}
}
