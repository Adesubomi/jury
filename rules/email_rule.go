package rules

import "github.com/adesubomi/jury"

type EmailRule struct {
	jury.Rule
	FieldName string
	Template  string
}

func (r EmailRule) Validate(defaultFieldName string) string {
	template := defaultFieldName + " is not valid"
	return template
}
