package change

import (
	"errors"
	"fmt"
	"math"
)

func Change(coins []int, target int) ([]int, error) {
	if len(coins) == 0 {
		return []int{}, errors.New("no coins")
	}
	if target < 0 || target < coins[0] {
		return []int{}, errors.New("target is a negative")
	}
	change := []int{}

	for i := 0; i <= target; i++ {
		change = append(change, target+1)
	}

	change[0] = 0
	for i := 1; i <= target; i++ {
		for _, coin := range coins {
			if coin <= i {
				fmt.Printf("the coin: %d\n", coin)
				change[i] = int(math.Min(float64(change[i]), float64(change[i-coin]+1)))
			}
		}
	}
	if change[target] > target {
		return []int{}, errors.New("no change")
	}
	return change, nil
}
