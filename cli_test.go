package poker_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"

	poker "github.com/mizrahishahar/go-web-app"
)

func TestCLI(t *testing.T) {

	// var dummyStdIn = &bytes.Buffer{}
	var dummyStdOut = &bytes.Buffer{}
	// var dummyGame = poker.NewGame(dummyBlindAlerter, dummyPlayerStore)

	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader("5\nChris wins\n")
		game := &poker.GameSpy{}

		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		poker.AssertEqual(t, game.FinishedWith, "Chris")
	})

	t.Run("record cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("5\nCleo wins\n")
		game := &poker.GameSpy{}

		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		poker.AssertEqual(t, game.FinishedWith, "Cleo")
	})
	t.Run("it prompts the user to enter the number of players and starts the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("7\nruth wins")
		game := &poker.GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		gotPrompt := stdout.String()
		wantPrompt := poker.PlayerPrompt

		if gotPrompt != wantPrompt {
			t.Errorf("got %q, want %q", gotPrompt, wantPrompt)
		}

		if game.StartedWith != 7 {
			t.Errorf("wanted Start called with 7 but got %d", game.StartedWith)
		}
	})
	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("Pies\n")
		game := &poker.GameSpy{}
	
		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()
	
		if game.StartCalled {
			t.Errorf("game should not have started")
		}

		gotPrompt := stdout.String()

		wantPrompt := poker.PlayerPrompt + "you're so silly"

		if gotPrompt != wantPrompt {
			t.Errorf("got %q, want %q", gotPrompt, wantPrompt)
		}
	})
	t.Run("it prints an error when the finisher isn't by the format 'wins'", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("5\nLloyd is a killer")
		game := &poker.GameSpy{}
	
		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()
	
		if game.FinishCalled {
			t.Errorf("game should not have finished")
		}

		gotPrompt := stdout.String()

		wantPrompt := poker.PlayerPrompt + "Because of you, nobody wins"

		if gotPrompt != wantPrompt {
			t.Errorf("got %q, want %q", gotPrompt, wantPrompt)
		}
	})

}

type SpyBlindAlerter struct {
	alerts []ScheduledAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.alerts = append(s.alerts, ScheduledAlert{duration, amount})
}

type ScheduledAlert struct {
	At     time.Duration
	Amount int
}

func (s ScheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.Amount, s.At)
}

func assertScheduledAlert(t testing.TB, got, want ScheduledAlert) {
	t.Helper()

	if got.Amount != want.Amount {
		t.Errorf("got amount %d, want %d", got.Amount, want.Amount)
	}

	if got.At != want.At {
		t.Errorf("got scheduled time of %v, want %v", got.At, want.At)
	}
}
