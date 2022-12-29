package game_test

import (
	"context"
	"testing"

	"githib.com/igomonov88/game-service/business/services/game"
	"githib.com/igomonov88/game-service/internal/logger"
	"githib.com/igomonov88/game-service/internal/middlewaries"
)

// Success and failure markers.
const (
	Success = "\u2713"
	Failed  = "\u2717"
)

var expectedValues = map[int]string{
	1: "rock",
	2: "paper",
	3: "scissors",
	4: "lizard",
	5: "spock",
}

func TestGame(t *testing.T) {
	t.Parallel()
	log := logger.Must(logger.New("game-test"))
	ctx := middlewaries.WithRequestIDContext(context.Background(), "test")
	svc := game.NewService(log)
	t.Run("getRandomChoiceID", func(t *testing.T) {
		t.Log("Given we want to test RandomChoiceID")
		if svc.RandomChoiceID(ctx) > 5 {
			t.Fatalf("\t%s\tTest %s:\tShould receive a valid choice ID", "getRandomChoiceID", Failed)
		}
		t.Logf("\t%s\tTest %s:\tShould receive a valid choice ID", "getRandomChoiceID", Success)
	})
	t.Run("getChoices", func(t *testing.T) {
		t.Log("Given we want to test AvailableChoices")
		{
			result := svc.AvailableChoices(ctx)
			if len(result) != len(expectedValues) {
				t.Fatalf("\t%s\tTest %s:\tShould receive a valid choices", "getChoices", Failed)
			}

			for id, name := range result {
				if name != expectedValues[id] {
					t.Fatalf("\t%s\tTest %s:\tShould receive a valid choices", "getChoices", Failed)
				}
			}
			t.Logf("\t%s\tTest %s:\tShould receive a valid choices", "getChoices", Success)
		}
	})
	t.Run("result", func(t *testing.T) {
		t.Log("Given we want to test Play")
		{
			tieResult := svc.Play(ctx, 1, 1)
			if tieResult != "tie" {
				t.Fatalf("\t%s\tTest %s:\tShould receive tie result if user and computer use the same values", "result", Failed)
			}

			winResult := svc.Play(ctx, 1, 3)
			if winResult != "win" {
				t.Fatalf("\t%s\tTest %s:\tShould receive win result if user plays rock and computer plays scissors", "result", Failed)
			}

			loseResult := svc.Play(ctx, 1, 2)
			if loseResult != "lose" {
				t.Fatalf("\t%s\tTest %s:\tShould receive lose result if user plays rock and computer plays paper", "result", Failed)
			}
			svc.Play(ctx, 2, 4)

			t.Logf("\t%s\tTest %s:\tShould receive a valid result", "result", Success)
		}
	})
}
