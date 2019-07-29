package poker

import (
	"bufio"
	"io"
	"strings"
)

type CLI struct {
	store PlayerStore
	r     io.Reader
}

func NewCLI(store PlayerStore, reader io.Reader) *CLI {
	return &CLI{store, reader}
}

func (c *CLI) PlayPoker() {
	scan := bufio.NewScanner(c.r)
	if scan.Scan() {
		cmd := scan.Text()
		c.store.IncrementScore(extractWinner(cmd))
	}

}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}
