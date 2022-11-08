package rules

import "github.com/adesubomi/jury"

type LengthRule struct {
	jury.Rule
	Min, Max int
}

type ExactLengthRule struct {
	jury.Rule
	Length int
}
