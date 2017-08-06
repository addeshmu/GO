package readinput

import (
	"bufio"
	"os"
	"strings"
)

// Getuserinput exported to get user input
func Getuserinput() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	return text
}
