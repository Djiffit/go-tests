package main

import (
	"fmt"
	"net/http"
)

type StubPlayerStore struct {
	scores map[string]int
	wins   []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score, ok := s.scores[name]
	if ok {
		return score
	}
	return 0
}

func (s *StubPlayerStore) IncrementScore(name string) {
	s.scores[name]++
	s.wins = append(s.wins, name)
}

type PlayerStore interface {
	GetPlayerScore(name string) int
	IncrementScore(name string)
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		p.postPlayer(w, r)
	case http.MethodGet:
		p.getPlayer(w, r)
	}
}

func (p *PlayerServer) getPlayer(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]
	score := p.store.GetPlayerScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, score)
}

func (p *PlayerServer) postPlayer(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]
	p.store.IncrementScore(player)
	w.WriteHeader(http.StatusAccepted)
}
