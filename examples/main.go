package main

import (
	"encoding/json"
	"fmt"
	"github.com/adesubomi/jury"
	"github.com/adesubomi/jury/rules"
)

func main() {

	rules := jury.ValidationRules{
		"email": jury.Rules{
			rules.EmailRule{},
		},
	}

	var req string
	validation, passed := rules.Validate(req)
	if !passed {
		valErrors, _ := json.Marshal(validation.Errors)
		fmt.Println("    ✗ validation failed")
		fmt.Println("       [..] ", string(valErrors))
	} else {
		fmt.Println("    ✔ validation passed")
	}
}
