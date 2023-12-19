package validator

import (
	"regexp"
)

var EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&’*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\\.[a-zA-Z0-9-]+)*$")

type Validator struct {
	Errors map[string]string
}

func New() *Validator {
	return &Validator{Errors : make(map[string]string)}
}

func (v *Validator) valid() bool {
	return len(v.Errors) == 0
}

// AddError adds an error message to the map (so long as no entry already exists for
// the given key).
func (v *Validator) AddError(key, message string) {
	//when we get value from a map, it returns value and boolean ifExists. we can check that boolean. No need to traverse through the whole map
	if _, exists := v.Errors[key]; exists == false {
		v.Errors[key] = message
	}
}

// Check adds an error message to the map only if a validation check is not 'ok'.
func (v *Validator) Check(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}
}

// In returns true if a specific value is in a list of strings.
func In(value string, list ...string) bool {
	for _, v := range list {
		if v ==  value {
			return true
		}
	}
	return false
}

// Matches returns true if a string value matches a specific regexp pattern.
func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}

// Unique returns true if all string values in a slice are unique.
func Unique(values []string) bool {
	// MY WAY
	// for k, v := range values {
	// 	for compareKey, compareVal := range values {
	// 		if k != compareKey && v == compareVal {
	// 			return false
	// 		}
	// 	}
	// }
	// return true

	//BOOK WAY
	uniques := make(map[string]bool)
	for _, v := range values {
		uniques[v] = true
	}

	return len(uniques) == len(values)

}



	