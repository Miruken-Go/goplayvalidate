package validator

import (
	"errors"
	ut "github.com/go-playground/universal-translator"
	play "github.com/go-playground/validator/v10"
	"github.com/miruken-go/miruken"
	"strings"
)

type GoPlaygroundValidator struct {
	validate   *play.Validate
	translator ut.Translator
}

func (v *GoPlaygroundValidator) Constructor(
	validate *play.Validate,
	_ *struct{ miruken.Optional }, translator ut.Translator,
) {
	v.validate   = validate
	v.translator = translator
}

func (v *GoPlaygroundValidator) Validate(
	validates *miruken.Validates, target any,
) miruken.HandleResult {
	if !miruken.IsStruct(target) {
		return miruken.NotHandled
	}
	if err := v.validate.Struct(target); err != nil {
		switch e := err.(type) {
		case *play.InvalidValidationError:
			return miruken.NotHandled.WithError(err)
		case play.ValidationErrors:
			outcome := validates.Outcome()
			if v.translator == nil {
				v.buildValidationOutcome(outcome, e)
			} else {
				v.translateValidationOutcome(outcome, e)
			}
			return miruken.HandledAndStop
		}
	}
	return miruken.Handled
}

func  (v *GoPlaygroundValidator) buildValidationOutcome(
	outcome     *miruken.ValidationOutcome,
	fieldErrors play.ValidationErrors,
) {
	for _, err := range fieldErrors {
		var path string
		ns    := err.StructNamespace()
		parts := strings.SplitN(ns, ".", 2)
		if len(parts) > 1 { path = parts[1] }
		outcome.AddError(path, err)
	}
}

func  (v *GoPlaygroundValidator) translateValidationOutcome(
	outcome     *miruken.ValidationOutcome,
	fieldErrors play.ValidationErrors,
) {
	for field, msg := range fieldErrors.Translate(v.translator) {
		var path string
		parts := strings.SplitN(field, ".", 2)
		if len(parts) > 1 { path = parts[1] }
		outcome.AddError(path, errors.New(msg))
	}
}
