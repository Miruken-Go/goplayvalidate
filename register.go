package goplayvalidate

import (
	ut "github.com/go-playground/universal-translator"
	play "github.com/go-playground/validator/v10"
	"github.com/miruken-go/miruken"
)

// GoPlaygroundValidationInstaller enables validation support for
// https://github.com/go-playground/validator/
type GoPlaygroundValidationInstaller struct {
	validate   *play.Validate
	translator ut.Translator
}

func (v *GoPlaygroundValidationInstaller) Validator() *play.Validate {
	return v.validate
}

func (v *GoPlaygroundValidationInstaller) SetTranslator(translator ut.Translator) {
	v.translator = translator
}

func (v *GoPlaygroundValidationInstaller) Install(registration *miruken.Registration) {
	if registration.CanInstall(&_registrationTag) {
		registration.Install(miruken.WithValidation())
		registration.AddHandlerTypes(miruken.TypeOf[*validator]())
		registration.AddHandlers(miruken.NewProvider(v.validate))
		if trans := v.translator; trans != nil {
			registration.AddHandlers(miruken.NewProvider(trans))
		}
	}
}

func WithGoPlaygroundValidation(
	config ... func(installer *GoPlaygroundValidationInstaller),
) miruken.Installer {
	installer := &GoPlaygroundValidationInstaller{validate: play.New()}
	for _, configure := range config {
		if configure != nil {
			configure(installer)
		}
	}
	return installer
}

var _registrationTag byte
