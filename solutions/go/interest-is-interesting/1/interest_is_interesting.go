package interest

const (
	firstLayerInterestRate  = 3.213
	secondLayerInterestRate = 0.5
	thirdLayerInterestRate  = 1.621
	fourthLayerInterestRate = 2.475
)

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	switch {
	case balance < 0:
		return firstLayerInterestRate
	case balance >= 0 && balance < 1000:
		return secondLayerInterestRate
	case balance >= 1000 && balance < 5000:
		return thirdLayerInterestRate
	case balance >= 5000:
		return fourthLayerInterestRate
	}
	return firstLayerInterestRate
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	rate := InterestRate(balance)

	return (balance * float64(rate) / float64(100))
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return Interest(balance) + balance
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	sum := balance
	years := 0
	for sum < targetBalance {
		sum = AnnualBalanceUpdate(sum)
		years++
	}
	return years
}
