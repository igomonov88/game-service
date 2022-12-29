package choice_test

import (
	"context"
	"testing"

	"githib.com/igomonov88/game-service/business/services/choice"
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

func Test_GetChoice(t *testing.T) {
	t.Parallel()
	log := logger.Must(logger.New("choice-test"))
	ctx := middlewaries.WithRequestIDContext(context.Background(), "test")
	gameService := game.NewService(log)
	svc := choice.NewService(log, gameService, gameService)
	t.Run("getChoice", func(t *testing.T) {
		t.Log("Given we request for a random choice")
		{
			response, err := svc.GetChoice(ctx)
			if err != nil {
				t.Fatalf("\t%s\tTest %s:\tShould receive a valid choice : %v", "getChoice", Failed, err)
			}

			name, exist := expectedValues[response.ID]
			if !exist {
				t.Fatalf("\t%s\tTest %s:\tShould receive a valid choice ID", "getChoice", Failed)
			}

			if name != response.Name {
				t.Fatalf("\t%s\tTest %s:\tShould receive a valid choice name", "getChoice", Failed)
			}

			t.Logf("\t%s\tTest %s:\tShould receive a valid choice", "getChoice", Success)
		}
	})
	t.Run("getChoices", func(t *testing.T) {
		t.Log("Given we request for all available choices")
		{
			choices := svc.GetChoices(ctx)
			if len(choices) != len(expectedValues) {
				t.Fatalf("\t%s\tTest %s:\tShould receive a valid choices", "getChoices", Failed)
			}

			t.Logf("\t%s\tTest %s:\tShould receive a valid choices", "getChoices", Success)
		}
	})

}
