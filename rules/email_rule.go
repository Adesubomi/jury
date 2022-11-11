package rules

import "github.com/adesubomi/jury"

type Email struct {
	jury.Rule
	FieldName string
	Template  string
}

func (r Email) Validate(defaultFieldName string) string {
	template := defaultFieldName + " is not valid"
	return template
}
