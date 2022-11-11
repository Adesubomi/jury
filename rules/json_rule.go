package rules

import "github.com/adesubomi/jury"

type IsJson struct {
	jury.Rule
	V string
}

type JsonContains struct {
	jury.Rule
	V string
}

type JsonIsSubsetOf struct {
	jury.Rule
	V string
}

type ExactJson struct {
	jury.Rule
	V string
}
