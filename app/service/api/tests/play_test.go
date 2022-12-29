package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"githib.com/igomonov88/game-service/app/service/api/handlers"
	"githib.com/igomonov88/game-service/business/contract"
	"githib.com/igomonov88/game-service/business/services/game"
	service "githib.com/igomonov88/game-service/business/services/play"
	"githib.com/igomonov88/game-service/business/storage"
	"githib.com/igomonov88/game-service/internal/database/dbtest"
)

type PlayTest struct {
	handler http.Handler
}

func Test_Play(t *testing.T) {
	t.Parallel()
	test := dbtest.NewIntegration(t)
	t.Cleanup(test.Teardown)
	gameServce := game.NewService(test.Log)
	storage := storage.NewService(test.DB)
	playService := service.NewService(test.Log, storage, gameServce, gameServce)
	api := handlers.Handler(test.Log, nil, playService, nil)
	tests := PlayTest{handler: api}

	t.Run("play", tests.playTest)
	t.Run("scoreboard", tests.scoreboardTest)
	t.Run("reset scoreboard", tests.resetScoreboardTest)
}

func (pt *PlayTest) playTest(t *testing.T) {
	r := httptest.NewRequest(http.MethodPost, "/play", strings.NewReader(`{"player": 1}`))
	w := httptest.NewRecorder()
	pt.handler.ServeHTTP(w, r)

	t.Logf("Given we make a play request")
	{
		if w.Code != http.StatusOK {
			t.Fatalf("\t%s\tTest %s:\tShould receive a status code of 200 for the response : %v", "play", Failed, w.Code)
		}

		var response contract.PlayResponse

		if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
			t.Fatalf("\t%s\tTest %s:\tShould be able to decode the response : %v", "play", Failed, err)
		}

		if _, exist := expectedValues[response.Player]; !exist {
			t.Fatalf("\t%s\tTest %s:\tShould receive a valid player choice : %v", "play", Failed, response.Player)
		}

		if _, exist := expectedValues[response.Computer]; !exist {
			t.Fatalf("\t%s\tTest %s:\tShould receive a valid computer choice : %v", "play", Failed, response.Computer)
		}

		if len(response.Results) == 0 {
			t.Fatalf("\t%s\tTest %s:\tShould receive a valid results : %v", "play", Failed, response.Results)
		}

		if response.Results != "win" && response.Results != "lose" && response.Results != "tie" {
			t.Fatalf("\t%s\tTest %s:\tShould receive a valid results : %v", "play", Failed, response.Results)
		}

		t.Logf("\t%s\tTest %s:\tShould receive a valid response", "play", Success)
	}
}

func (pt *PlayTest) scoreboardTest(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/scoreboard", strings.NewReader(``))
	w := httptest.NewRecorder()
	pt.handler.ServeHTTP(w, r)

	t.Logf("Given we make a scoreboard request")
	{
		if w.Code != http.StatusOK {
			t.Fatalf("\t%s\tTest %s:\tShould receive a status code of 200 for the response : %v", "play", Failed, w.Code)
		}

		var response []contract.PlayResponse

		if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
			t.Fatalf("\t%s\tTest %s:\tShould be able to decode the response : %v", "play", Failed, err)
		}

		if len(response) > 10 {
			t.Fatalf("\t%s\tTest %s:\tShould receive not more then 10 results in scoreboard request: %v", "play", Failed, response)
		}

		t.Logf("\t%s\tTest %s:\tShould receive a valid response", "play", Success)
	}
}

func (pt *PlayTest) resetScoreboardTest(t *testing.T) {
	r := httptest.NewRequest(http.MethodDelete, "/scoreboard", strings.NewReader(``))
	w := httptest.NewRecorder()
	pt.handler.ServeHTTP(w, r)

	t.Logf("Given we make a reset scoreboard request")
	{
		if w.Code != http.StatusNoContent {
			t.Fatalf("\t%s\tTest %s:\tShould receive a status code of 200 for the response : %v", "play", Failed, w.Code)
		}

		t.Logf("\t%s\tTest %s:\tShould receive a valid response", "play", Success)
	}
}
