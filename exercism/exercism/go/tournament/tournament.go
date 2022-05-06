package tournament

import (
	"encoding/csv"
	"fmt"
	"io"
	"sort"
)

type team struct {
	name    string
	matches int
	wins    int
	losses  int
	draws   int
	points  int
}

// Tally calculates the tournament table.
func Tally(r io.Reader, w io.Writer) error {
	// csv reader has the perfect solution to this.
	reader := csv.NewReader(r)
	reader.Comma = ';'
	reader.Comment = '#'
	reader.FieldsPerRecord = 3 // Home, visitor, outcome

	tally := make(map[string]team)

	// Iterate over the reader
	for {
		game, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		// Get Game results
		homePlayer := game[0]
		awayPlayer := game[1]
		gameResults := game[2]

		// Get Team statistics
		home := tally[homePlayer]
		visitor := tally[awayPlayer]
		home.name = homePlayer
		visitor.name = awayPlayer

		// Increase matches played
		home.matches++
		visitor.matches++

		// Give credit where credit is due
		switch gameResults {
		case "win":
			home.points += 3
			home.wins++
			visitor.losses++
		case "draw":
			home.points++
			home.draws++
			visitor.points++
			visitor.draws++
		case "loss":
			visitor.points += 3
			visitor.wins++
			home.losses++
		default:
			return fmt.Errorf("Unknown result: %s", gameResults)
		}

		tally[homePlayer] = home
		tally[awayPlayer] = visitor
	}

	// Create a teams (results) table/slice
	teams := make([]team, 0, len(tally))
	for _, v := range tally {
		teams = append(teams, v)
	}

	// Sort final resuls from points (not by name)
	sort.Slice(teams, func(i, j int) bool {
		if teams[i].points == teams[j].points {
			return teams[i].name < teams[j].name
		}
		return teams[i].points > teams[j].points
	})

	// Create the output
	shouldReturn, returnValue := outputString(w, teams)
	if shouldReturn {
		return returnValue
	}
	return nil
}

func outputString(w io.Writer, teams []team) (bool, error) {
	_, err := fmt.Fprintf(
		w, "%-30s | %2s | %2s | %2s | %2s | %2s\n",
		"Team", "MP", "W", "D", "L", "P")
	if err != nil {
		return true, err
	}
	for _, t := range teams {
		_, err := fmt.Fprintf(
			w, "%-30s | %2d | %2d | %2d | %2d | %2d\n",
			t.name, t.matches, t.wins, t.draws, t.losses, t.points,
		)
		if err != nil {
			return true, err
		}
	}
	return false, nil
}
