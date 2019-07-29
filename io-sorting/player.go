package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func NewPlayerServer(store PlayerStore) *PlayerServer {
	s := PlayerServer{
		store,
		http.NewServeMux(),
	}

	router := http.NewServeMux()

	router.Handle("/league", http.HandlerFunc(s.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(s.playerHandler))

	s.Handler = router

	return &s
}

func (p *PlayerServer) getPlayer(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]
	score := p.store.GetPlayerScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, score)
}

func (p *PlayerServer) getLeagueTable() []Player {
	return p.store.GetLeague()
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	leagueTable := p.getLeagueTable()
	json.NewEncoder(w).Encode(leagueTable)
	w.WriteHeader(http.StatusOK)

}

func (p *PlayerServer) playerHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		p.postPlayer(w, r)
	case http.MethodGet:
		p.getPlayer(w, r)
	}
}

func (p *PlayerServer) postPlayer(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]
	p.store.IncrementScore(player)
	w.WriteHeader(http.StatusAccepted)
}

type PlayerStore interface {
	GetPlayerScore(name string) int
	IncrementScore(name string)
	GetLeague() League
}

type PlayerServer struct {
	store PlayerStore
	http.Handler
}

type Player struct {
	Name string
	Wins int
}
