package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"githib.com/igomonov88/game-service/app/service/api/handlers"
	"githib.com/igomonov88/game-service/business/contract"
	"githib.com/igomonov88/game-service/business/services/choice"
	"githib.com/igomonov88/game-service/business/services/game"
	"githib.com/igomonov88/game-service/internal/database/dbtest"
)

// Success and failure markers.
const (
	Success = "\u2713"
	Failed  = "\u2717"
)

type ChoicesTests struct {
	handler http.Handler
}

var expectedValues = map[int]string{
	1: "rock",
	2: "paper",
	3: "scissors",
	4: "lizard",
	5: "spock",
}

func Test_Choice(t *testing.T) {
	t.Parallel()
	test := dbtest.NewIntegration(t, c)
	t.Cleanup(test.Teardown)
	gameService := game.NewService(test.Log)
	choiceService := choice.NewService(test.Log, gameService)
	api := handlers.Handler(test.Log, choiceService, nil, nil)
	tests := ChoicesTests{handler: api}

	t.Run("getChoice", tests.getChoice)
	t.Run("getChoices", tests.getChoices)
}

func (ct *ChoicesTests) getChoice(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/choice", strings.NewReader(`{}`))
	w := httptest.NewRecorder()
	ct.handler.ServeHTTP(w, r)

	t.Logf("Given we request for a random choice")
	{
		if w.Code != http.StatusOK {
			t.Fatalf("\t%s\tTest %s:\tShould receive a status code of 200 for the response : %v", "getChoice", Failed, w.Code)
		}

		var response contract.ChoiceResponse

		if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
			t.Fatalf("\t%s\tTest %s:\tShould be able to unmarshal the response : %v", "getChoice", Failed, err)
		}

		if _, exist := expectedValues[response.ID]; !exist {
			t.Fatalf("\t%s\tTest %s:\tShould respond with valid id of specific choise: %v", "getChoice", Failed, response.ID)
		}

		if expectedValues[response.ID] != response.Name {
			t.Fatalf("\t%s\tTest %s:\tShould respond with valid name: %v for specific choise ID: %v but respond with name: %v", "getChoice", Failed, expectedValues[response.ID], response.ID, response.Name)
		}

		t.Logf("\t%s\tTest %s: \tShould get the expected result.", "getChoise", Success)
	}
}

func (ct *ChoicesTests) getChoices(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/choices", strings.NewReader(`{}`))
	w := httptest.NewRecorder()
	ct.handler.ServeHTTP(w, r)

	t.Logf("Given we request a list of choices")
	{
		if w.Code != http.StatusOK {
			t.Fatalf("\t%s\tTest %s:\tShould receive a status code of 200 for the choices : %v", "getChoices", Failed, w.Code)
		}

		var choices []contract.ChoiceResponse
		if err := json.NewDecoder(w.Body).Decode(&choices); err != nil {
			t.Fatalf("\t%s\tTest %s:\tShould be able to unmarshal the choices : %v", "getChoices", Failed, err)
		}

		if len(choices) != len(expectedValues) {
			t.Fatalf("\t%s\tTest %s:\tShould respond with valid number of choices: %v", "getChoices", Failed, len(choices))
		}

		for _, choice := range choices {
			if expectedValues[choice.ID] != choice.Name {
				t.Fatalf("\t%s\tTest %s:\tShould respond with valid name: %v for specific choise ID: %v but respond with name: %v", "getChoices", Failed, expectedValues[choice.ID], choice.ID, choice.Name)
			}

			if _, exist := expectedValues[choice.ID]; !exist {
				t.Fatalf("\t%s\tTest %s:\tShould respond with valid id of specific choise: %v", "getChoices", Failed, choice.ID)
			}
		}

		t.Logf("\t%s\tTest %s: \tShould get the expected result.", "getChoices", Success)
	}
}
