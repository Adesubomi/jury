package rules

import "github.com/adesubomi/jury"

type NumberBetween struct {
	Min, Max float64
}

type NumberIsEqualTo struct {
	jury.Rule
	V float64
}

type NumberIsLessThan struct {
	jury.Rule
	V float64
}

type NumberIsGreaterThan struct {
	jury.Rule
	V float64
}

type NumberIsPositive struct {
	jury.Rule
}

type NumberIsNegative struct {
	jury.Rule
}

type DigitsCount struct {
	jury.Rule
	V int
}

type DigitsCountBetween struct {
	jury.Rule
	Min, Max int
}
