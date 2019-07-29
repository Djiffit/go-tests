package poker

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
)

type League []Player

func NewFileSystemPlayerStore(database *os.File) (*FileSystemPlayerStore, error) {
	err := initialisePlayerDBFile(database)

	if err != nil {
		return nil, fmt.Errorf("problem initialising player db file, %v", err)
	}

	league, err := NewLeague(database)

	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", database.Name(), err)
	}

	return &FileSystemPlayerStore{
		db:     json.NewEncoder(&tape{database}),
		league: league,
	}, nil
}

func initialisePlayerDBFile(file *os.File) error {
	file.Seek(0, 0)
	info, err := file.Stat()

	if err != nil {
		return fmt.Errorf("problem getting file info for %s, %v", file.Name(), err)
	}

	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, 0)
	}

	return nil
}

func (l League) Find(player string) *Player {
	for i, plr := range l {
		if plr.Name == player {
			return &l[i]
		}
	}
	return nil
}

func NewLeague(reader io.Reader) (league []Player, err error) {
	err = json.NewDecoder(reader).Decode(&league)
	if err != nil {
		err = fmt.Errorf("Problem parsing league %v", err)
	}
	return
}

type FileSystemPlayerStore struct {
	db     *json.Encoder
	league League
}

func (f *FileSystemPlayerStore) GetLeague() League {
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins > f.league[j].Wins
	})

	return f.league
}

func (f *FileSystemPlayerStore) GetPlayerScore(player string) (wins int) {
	plr := f.league.Find(player)

	if plr != nil {
		wins = plr.Wins
	}

	return
}

func (f *FileSystemPlayerStore) Serialize(players []Player) {
	f.db.Encode(players)
	f.league = players
}

func (l League) AddPlayer(name string, wins int) League {
	return append(l, Player{name, wins})
}

func (f *FileSystemPlayerStore) IncrementScore(player string) {
	league := f.league
	plr := league.Find(player)

	if plr != nil {
		plr.Wins++
	} else {
		league = league.AddPlayer(player, 1)
	}

	f.Serialize(league)
}
