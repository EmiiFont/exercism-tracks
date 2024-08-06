package kindergarten

import (
  "errors"
  "strings"
  "sort"
)
// Define the Garden type here.
type Garden struct {
  plants map[string][]string
}
// The diagram argument starts each row with a '\n'.  This allows Go's
// raw string literals to present diagrams in source code nicely as two
// rows flush left, for example,
//
//     diagram := `
//     VVCCGG
//     VVCCGG`
// Map plant codes to plant names.
var plantMap = map[rune]string{
	'G': "grass",
	'C': "clover",
	'R': "radishes",
  'V': "violets",
}

func NewGarden(diagram string, children []string) (*Garden, error) {
  if diagram[0] != '\n' {
    return nil, errors.New("diagram error")
  }
  childrenSorted := []string{}
  childrenSorted = append(childrenSorted, children...)
  sort.Strings(childrenSorted)
  rows := strings.Split(diagram, "\n")
if rows[0] == "" {
     rows = rows[1:]
   }
	if len(rows) != 2 {
		return nil, errors.New("diagram should have exactly two rows")
	}
  rowLength := len(rows[0])
	if rowLength != len(rows[1]) {
		return nil, errors.New("rows should have the same length")
	}

	if rowLength%2 != 0 {
		return nil, errors.New("number of cups should be even")
	}

	g := &Garden{
		plants: make(map[string][]string),
	}

	for _, row := range rows {
		for _, char := range row {
			if _, ok := plantMap[char]; !ok {
				return nil, errors.New("invalid cup code found")
			}
		}
	}

	for i, student := range childrenSorted{
		start := i * 2
		end := start + 2
     
  		plants := []string{
			plantMap[rune(rows[0][start])],
			plantMap[rune(rows[0][end-1])],
			plantMap[rune(rows[1][start])],
			plantMap[rune(rows[1][end-1])],
		}

    if _, ok := g.plants[student]; ok {
      return nil, errors.New("duplicate student")
    }

		g.plants[student] = plants
	}

	return g, nil
  }

func (g *Garden) Plants(child string) ([]string, bool) {
 plants, ok := g.plants[child]
	return plants, ok
}
