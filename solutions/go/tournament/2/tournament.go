package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
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
		fmt.Println(line)
		if line == "" {
			continue
		}
		team1, team2 := parseLine(line)
		fmt.Printf("team1: %v, team2: %v\n", team1, team2)
		ref1, ok := results[team1.name]
		if !ok {
			results[team1.name] = &TeamResult{win: team1.win, loss: team1.loss, draw: team1.draw, points: team1.points, gamePlayed: team1.gamePlayed}
		} else {
			results[team1.name].win += team1.win
			results[team1.name].loss += team1.loss
			results[team1.name].draw += team1.draw
			results[team1.name].points += team1.points
			results[team1.name].gamePlayed += 1
		}
		ref2, ok := results[team2.name]
		if !ok {
			results[team2.name] = &TeamResult{win: team2.win, loss: team2.loss, draw: team2.draw, points: team2.points, gamePlayed: team2.gamePlayed}
		} else {
			results[team2.name].win += team2.win
			results[team2.name].loss += team2.loss
			results[team2.name].draw += team2.draw
			results[team2.name].points += team2.points
			results[team2.name].gamePlayed += 1
		}
		fmt.Printf("team1: %v, team2: %v\n", ref1, ref2)
	}

	sortedResults := sortTeamResultMapByPoints(results)
	fmt.Fprintf(writer, "%-30s | %2s | %2s | %2s | %2s | %2s\n", "Team", "MP", "W", "D", "L", "P")
	for i, v := range sortedResults {
		writeLines(writer, v.name, strconv.Itoa(v.gamePlayed), strconv.Itoa(v.win), strconv.Itoa(v.draw), strconv.Itoa(v.loss), strconv.Itoa(v.points))
		if i < len(sortedResults)-1 {
			writer.Write([]byte("\n"))
		}
	}
	return nil
}

func writeLines(writer io.Writer, args ...interface{}) {
	fmt.Fprintf(writer, "%-31s", args[0])
	fmt.Fprintf(writer, "%s", "|")
	fmt.Fprintf(writer, "%3s", args[1])
	fmt.Fprintf(writer, "%2s", "|")
	fmt.Fprintf(writer, "%3s", args[2])
	fmt.Fprintf(writer, "%2s", "|")
	fmt.Fprintf(writer, "%3s", args[3])
	fmt.Fprintf(writer, "%2s", "|")
	fmt.Fprintf(writer, "%3s", args[4])
	fmt.Fprintf(writer, "%2s", "|")
	fmt.Fprintf(writer, "%3s", args[5])
}

func parseLine(ie string) (team1 *TeamResult, team2 *TeamResult) {
	split := strings.Split(ie, ";")

	if len(split) != 3 {
		return &TeamResult{}, &TeamResult{}
	}

	fmt.Printf("split: %v\n", split)

	switch split[2] {
	case "win":
		return &TeamResult{name: split[0], win: 1, loss: 0, draw: 0, gamePlayed: 1, points: 3}, &TeamResult{name: split[1], win: 0, loss: 1, draw: 0, gamePlayed: 1, points: 0}
	case "loss":
		return &TeamResult{name: split[0], win: 0, loss: 1, draw: 0, gamePlayed: 1, points: 0}, &TeamResult{name: split[1], win: 1, loss: 0, draw: 0, gamePlayed: 1, points: 3}
	case "draw":
		return &TeamResult{name: split[0], win: 0, loss: 0, draw: 1, gamePlayed: 1, points: 1}, &TeamResult{name: split[1], win: 0, loss: 0, draw: 1, gamePlayed: 1, points: 1}
	default:
		return &TeamResult{}, &TeamResult{}
	}
}
