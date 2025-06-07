package hw08

import (
	"fmt"
	"strings"
)

type MultiError struct {
	errors []error
}

func (e *MultiError) Error() string {
	if len(e.errors) == 0 {
		return "no errors"
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%d errors occurred:\n", len(e.errors)))

	for _, err := range e.errors {
		sb.WriteString(fmt.Sprintf("\t* %s", err.Error()))
	}

	sb.WriteString("\n")
	return sb.String()
}

func Append(err error, errs ...error) *MultiError {
	multiError := &MultiError{}

	if err != nil {
		if me, ok := err.(*MultiError); ok {
			multiError.errors = append(multiError.errors, me.errors...)
		} else {
			multiError.errors = append(multiError.errors, err)
		}
	}

	multiError.errors = append(multiError.errors, errs...)

	return multiError
}
