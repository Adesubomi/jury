package jury

type ValidationRules map[string]Rules

type ValidationResult struct {
	Errors map[string]string
}

func (valRules ValidationRules) Validate(request interface{}) (ValidationResult, bool) {
	result := new(ValidationResult)
	result.Errors = map[string]string{}
	passed := true

	for field, rules := range valRules {
		for _, rule := range rules {

			msg := rule.Validate(field)
			if msg != "" {
				passed = false
				result.Errors[field] = msg
				break
			}
		}
	}

	return *result, passed
}
