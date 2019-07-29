package poker

import (
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {

	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &StubPlayerStore{
			map[string]int{},
			[]string{},
			[]Player{},
		}

		cli := &CLI{playerStore, in}
		cli.PlayPoker()

		assertPlayerWin(t, playerStore, "Chris")
	})

	t.Run("record cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		playerStore := &StubPlayerStore{
			map[string]int{},
			[]string{},
			[]Player{},
		}

		cli := &CLI{playerStore, in}
		cli.PlayPoker()

		assertPlayerWin(t, playerStore, "Cleo")
	})

}

func assertPlayerWin(t *testing.T, playerStore *StubPlayerStore, want string) {
	t.Helper()

	if len(playerStore.wins) < 1 {
		t.Fatal("expected a win call")
	}

	got := playerStore.wins[0]

	if got != want {
		t.Errorf("expected %q to win got %q", want, got)
	}
}
