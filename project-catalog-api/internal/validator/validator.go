package validator

import (
	"slices"
	"strings"
	"unicode/utf8"
)

type Validator struct {
	Errors map[string]string
}

func New() *Validator {
	return &Validator{
		Errors: make(map[string]string),
	}
}

func (v *Validator) AddError(key, message string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}
}

func (v *Validator) Check(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}
}

func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

func NotBlank(value string) bool {
	return len(strings.TrimSpace(value)) != 0
}

func MinChar(value string, n int) bool {
	return utf8.RuneCountInString(strings.TrimSpace(value)) >= n
}

func MaxChar(value string, n int) bool {
	return utf8.RuneCountInString(strings.TrimSpace(value)) <= n
}

func PermittedValue[T comparable](value T, permittedValues ...T) bool {
	return slices.Contains(permittedValues, value)
}
