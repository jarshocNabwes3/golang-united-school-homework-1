package solution

import (
	"strings"

	"github.com/kyokomi/emoji"
)

func GetMessage(template string, data string) string {
	rendered := emoji.Sprint(template, data)
	rendered = strings.TrimSuffix(rendered, " ") + "!"

	return rendered
}
