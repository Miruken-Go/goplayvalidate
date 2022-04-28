package validator

import "github.com/miruken-go/miruken"
import playground "github.com/go-playground/validator/v10"

//go:generate $GOPATH/bin/mirukentypes -tests

type IntegrationValidator struct {
	validate *playground.Validate
}

func (v *IntegrationValidator) Constructor() {
	v.validate = playground.New()
}

func (v *IntegrationValidator) Validate(
	validates *miruken.Validates, target any,
) miruken.HandleResult {
	if miruken.IsStruct(target) {
		return miruken.NotHandled
	}
	if err := v.validate.Struct(target); err != nil {
		switch e := err.(type) {
		case *playground.InvalidValidationError:
			return miruken.NotHandled.WithError(err)
		case playground.ValidationErrors:
			addOutcomeErrors(validates.Outcome(), e)
		}
	}
	return miruken.Handled
}

func addOutcomeErrors(
	outcome  *miruken.ValidationOutcome,
	failures playground.ValidationErrors,
) {

}
