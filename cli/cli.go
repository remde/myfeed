package cli

import (
	"bufio"
	"os"
)

func Start() {
	reader = bufio.NewReader(os.Stdin)
}
