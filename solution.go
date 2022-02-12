package solution

import (
	"strings"

	"github.com/kyokomi/emoji"
)

const (
	helloText = "Hello "
	worldText = ":world_map:"
)

func GetMessage() string {

	rendered := emoji.Sprint(helloText, worldText)
	rendered = strings.TrimSuffix(rendered, " ") + "!"

	return rendered
}
