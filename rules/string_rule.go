package rules

import "github.com/adesubomi/jury"

type CharSet struct {
	Alpha       bool
	Num         bool
	SpecialChar bool
	Underscore  bool
	Dash        bool
	Chars       []rune
}

type StringContains struct {
	jury.Rule
	V string
}

type StringStartsWith struct {
	jury.Rule
	V string
}

type StringEndsWith struct {
	jury.Rule
	V string
}

type UUID struct {
	jury.Rule
}

type IPAddress struct {
	jury.Rule
}

type RegEx struct {
	jury.Rule
}

type ActiveUrl struct {
	jury.Rule
}
