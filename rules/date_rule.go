package rules

import "github.com/adesubomi/jury"

type DateFormat struct {
	jury.Rule
	V string
}

type DateIsBefore struct {
	V string
}

type DateIsAfter struct {
	V string
}

type DateIsBetween struct {
	Min string
	Max string
}

type Timezone struct {
}
