package solution

import (
	"github.com/kyokomi/emoji"
)

const (
	helloText = "Hello "
	worldText = ":world_map:"
)

func GetMessage() string {

	rendered := emoji.Sprint(helloText, worldText)
	rendered = rendered + "!"

	return rendered
}
