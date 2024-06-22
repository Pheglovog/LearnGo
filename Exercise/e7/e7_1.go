package main

import (
	"io"
	"os"
	"slices"
	"sort"
)

type Ranker interface {
	Ranking() []string
}

func RankPrinter(r Ranker, w io.Writer) {
	results := r.Ranking()
	for _, v := range results {
		io.WriteString(w, v)
		w.Write([]byte("\n"))
	}
}

type Team struct {
	teamName    string
	playerNames []string
}

type League struct {
	Teams []Team
	Wins  map[string]int
}

func (l *League) MatchResult(firstName string, firstScore int, secondName string, secondScore int) {
	if !slices.ContainsFunc(l.Teams, func(t Team) bool { return t.teamName == firstName }) {
		return
	}
	if !slices.ContainsFunc(l.Teams, func(t Team) bool { return t.teamName == secondName }) {
		return
	}
	if firstScore > secondScore {
		l.Wins[firstName]++
	} else if firstScore < secondScore {
		l.Wins[secondName]++
	}
}

func (l League) Ranking() []string {
	names := make([]string, 0, len(l.Teams))
	for _, v := range l.Teams {
		names = append(names, v.teamName)
	}
	sort.Slice(names, func(i, j int) bool {
		return l.Wins[names[i]] > l.Wins[names[j]]
	})
	// slices.SortFunc(names, func(a, b string) int {
	// 	if l.Wins[a] < l.Wins[b] {
	// 		return 1
	// 	} else if l.Wins[a] == l.Wins[b] {
	// 		return 0
	// 	} else {
	// 		return -1
	// 	}
	// })
	return names
}

func main() {
	l := League{
		Teams: []Team{
			{
				teamName:    "Italy",
				playerNames: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			{
				teamName:    "France",
				playerNames: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			{
				teamName:    "India",
				playerNames: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			{
				teamName:    "Nigeria",
				playerNames: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
		},
		Wins: map[string]int{},
	}
	l.MatchResult("Italy", 50, "France", 70)
	l.MatchResult("India", 85, "Nigeria", 80)
	l.MatchResult("Italy", 60, "India", 55)
	l.MatchResult("France", 100, "Nigeria", 110)
	l.MatchResult("Italy", 65, "Nigeria", 70)
	l.MatchResult("France", 95, "India", 80)
	RankPrinter(l, os.Stdout)
}
