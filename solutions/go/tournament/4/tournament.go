package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type TeamResult struct {
	name       string
	win        int
	loss       int
	draw       int
	gamePlayed int
	points     int
}

func sortTeamResultMapByPoints(tr map[string]*TeamResult) []*TeamResult {
	sorted := []*TeamResult{}

	for key, val := range tr {
		val.name = key
		sorted = append(sorted, val)
	}

	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].points == sorted[j].points {
			return sorted[i].name < sorted[j].name
		}
		return sorted[i].points > sorted[j].points
	})

	return sorted
}

func Tally(reader io.Reader, writer io.Writer) error {
	results := map[string]*TeamResult{}

	scanner := bufio.NewScanner(reader)

	// Read and print each line
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || line[0] == '#' {
			continue
		}
		team1, team2, err := parseLine(line)
		if err != nil {
			return err
		}
		_, ok := results[team1.name]
		if !ok {
			results[team1.name] = &TeamResult{win: team1.win, loss: team1.loss, draw: team1.draw, points: team1.points, gamePlayed: team1.gamePlayed}
		} else {
			results[team1.name].win += team1.win
			results[team1.name].loss += team1.loss
			results[team1.name].draw += team1.draw
			results[team1.name].points += team1.points
			results[team1.name].gamePlayed += 1
		}
		_, ok = results[team2.name]
		if !ok {
			results[team2.name] = &TeamResult{win: team2.win, loss: team2.loss, draw: team2.draw, points: team2.points, gamePlayed: team2.gamePlayed}
		} else {
			results[team2.name].win += team2.win
			results[team2.name].loss += team2.loss
			results[team2.name].draw += team2.draw
			results[team2.name].points += team2.points
			results[team2.name].gamePlayed += 1
		}
	}

	sortedResults := sortTeamResultMapByPoints(results)
	fmt.Fprintf(writer, "%-30s | %2s | %2s | %2s | %2s | %2s\n", "Team", "MP", "W", "D", "L", "P")
	for _, v := range sortedResults {
		fmt.Fprintf(writer, "%-30s | %2d | %2d | %2d | %2d | %2d\n", v.name, v.gamePlayed, v.win, v.draw, v.loss, v.points)
	}
	return nil
}

func parseLine(ie string) (team1 *TeamResult, team2 *TeamResult, err error) {
	split := strings.Split(ie, ";")

	if len(split) != 3 {
		return &TeamResult{}, &TeamResult{}, errors.New("invalid input")
	}

	switch split[2] {
	case "win":
		return &TeamResult{name: split[0], win: 1, loss: 0, draw: 0, gamePlayed: 1, points: 3}, &TeamResult{name: split[1], win: 0, loss: 1, draw: 0, gamePlayed: 1, points: 0}, nil
	case "loss":
		return &TeamResult{name: split[0], win: 0, loss: 1, draw: 0, gamePlayed: 1, points: 0}, &TeamResult{name: split[1], win: 1, loss: 0, draw: 0, gamePlayed: 1, points: 3}, nil
	case "draw":
		return &TeamResult{name: split[0], win: 0, loss: 0, draw: 1, gamePlayed: 1, points: 1}, &TeamResult{name: split[1], win: 0, loss: 0, draw: 1, gamePlayed: 1, points: 1}, nil
	default:
		return &TeamResult{}, &TeamResult{}, errors.New("invalid input")
	}
}
