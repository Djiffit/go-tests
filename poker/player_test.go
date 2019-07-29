package poker

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGETPlayers(t *testing.T) {

	store := &StubPlayerStore{
		map[string]int{
			"Pekka": 11,
			"Floyd": 10,
		},
		nil,
		nil,
	}

	server := NewPlayerServer(store)

	t.Run("returns Pekka's score", func(t *testing.T) {
		req := newGetScoreRequest("Pekka")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, req)

		got := response.Body.String()
		want := "11"

		assertResponseBody(t, got, want)
		assertStatusCode(t, response.Code, http.StatusOK)
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "10"

		assertResponseBody(t, got, want)
		assertStatusCode(t, response.Code, http.StatusOK)
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatusCode(t, response.Code, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {

	t.Run("it returns accepted on POST", func(t *testing.T) {
		server, _ := createStoreAndServer()
		request, _ := http.NewRequest(http.MethodPost, "/players/Pepper", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		assertStatusCode(t, response.Code, http.StatusAccepted)
	})

	t.Run("it increments wins on POST", func(t *testing.T) {
		server, store := createStoreAndServer()
		request, _ := http.NewRequest(http.MethodPost, "/players/Pekka", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assertStatusCode(t, response.Code, http.StatusAccepted)

		got := store.GetPlayerScore("Pekka")

		if got != 1 {
			t.Errorf("wins were not incremented properly, got %d, wanted %d", got, 1)
		}

		if !reflect.DeepEqual(store.wins, []string{"Pekka"}) {
			t.Errorf("wins were not recorded, got %v, wanted %v", store.wins, []string{"Pekka"})
		}
	})
}

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := NewStore()
	server := NewPlayerServer(store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWin(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWin(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWin(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(player))
	assertStatusCode(t, response.Code, http.StatusOK)

	assertResponseBody(t, response.Body.String(), "3")
}

func newPostWin(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, "/players/"+name, nil)
	return req
}

func createStoreAndServer() (*PlayerServer, *StubPlayerStore) {
	store := &StubPlayerStore{
		map[string]int{},
		nil,
		nil,
	}
	server := NewPlayerServer(store)
	return server, store
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}

func assertStatusCode(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("status was %d, expected %d", got, want)
	}
}

type StubPlayerStore struct {
	scores map[string]int
	wins   []string
	league League
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score, ok := s.scores[name]
	if ok {
		return score
	}
	return 0
}

func (s *StubPlayerStore) GetLeague() League {
	return s.league
}

func (s *StubPlayerStore) IncrementScore(name string) {
	s.scores[name]++
	s.wins = append(s.wins, name)
}
