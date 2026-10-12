package wordy

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type Stack struct {
	items []int
}

// Push adds an item to the top of the stack
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Pop removes and returns the item from the top of the stack
func (s *Stack) Pop() (int, error) {
	if len(s.items) == 0 {
		return 0, errors.New("stack is empty")
	}
	// Get the top item
	item := s.items[len(s.items)-1]
	// Remove the top item
	s.items = s.items[:len(s.items)-1]
	return item, nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of items in the stack
func (s *Stack) Size() int {
	return len(s.items)
}

func Answer(question string) (int, bool) {
	numberRegex := `-?\d+`
	number := ""
	operation := ""
	stack := &Stack{}
	re := regexp.MustCompile(numberRegex)
	for _, word := range strings.Fields(question) {

		if !re.Match([]byte(word)) && matchOperation(word) == "" {
			continue
		}
		if re.Match([]byte(word)) {
			if word[len(word)-1] == '?' {
				number = word[:len(word)-1]
			} else {
				number = word
			}
		}

		switch operation {
		case "+":
			prev, _ := stack.Pop()
			integer, err := strconv.Atoi(number)
			if err != nil {
				return 0, false
			}
			stack.Push(prev + integer)
		case "-":
			prev, _ := stack.Pop()
			integer, err := strconv.Atoi(number)
			if err != nil {
				return 0, false
			}
			stack.Push(prev - integer)
		case "*":
			prev, _ := stack.Pop()
			integer, err := strconv.Atoi(number)
			if err != nil {
				return 0, false
			}
			stack.Push(prev * integer)
		case "/":
			prev, _ := stack.Pop()
			integer, err := strconv.Atoi(number)
			if err != nil {
				return 0, false
			}
			stack.Push(prev / integer)
		}
		operation = matchOperation(word)
		if stack.Size() == 0 {
			integer, err := strconv.Atoi(number)
			if err != nil {
				return 0, false
			}
			stack.Push(integer)
		}
	}

	if stack.IsEmpty() || stack.Size() > 1 {
		return 0, false
	}

	res, _ := stack.Pop()
	return res, true
}

func matchOperation(word string) string {
	switch word {
	case "plus":
		return "+"
	case "minus":
		return "-"
	case "multiplied":
		return "*"
	case "divided":
		return "/"
	}
	return ""
}
