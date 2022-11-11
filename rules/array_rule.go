package rules

import "github.com/adesubomi/jury"

type IsArray struct {
}

type ArrayContains struct {
	jury.Rule
	V []any
}

type ArrayIsSubsetOf struct {
	jury.Rule
	V []any
}

type ArrayIs struct {
	jury.Rule
	V      []any
	Strict bool
}
