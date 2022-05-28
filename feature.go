package goplayvalidate

import (
	ut "github.com/go-playground/universal-translator"
	play "github.com/go-playground/validator/v10"
	"github.com/miruken-go/miruken"
)

// GoPlaygroundValidationInstaller enables validation support
// for https://github.com/go-playground/validator/
type GoPlaygroundValidationInstaller struct {
	validate   *play.Validate
	translator ut.Translator
}

func (v *GoPlaygroundValidationInstaller) Validator() *play.Validate {
	return v.validate
}

func (v *GoPlaygroundValidationInstaller) UseTranslator(translator ut.Translator) {
	v.translator = translator
}

func (v *GoPlaygroundValidationInstaller) Install(setup *miruken.SetupBuilder) error {
	if setup.CanInstall(&_featureTag) {
		setup.Install(miruken.WithValidation())
		setup.RegisterHandlers(&validator{})
		setup.AddHandlers(miruken.NewProvider(v.validate))
		if trans := v.translator; trans != nil {
			setup.AddHandlers(miruken.NewProvider(trans))
		}
	}
	return nil
}

func WithGoPlaygroundValidation(
	config ... func(installer *GoPlaygroundValidationInstaller),
) miruken.Feature {
	installer := &GoPlaygroundValidationInstaller{validate: play.New()}
	for _, configure := range config {
		if configure != nil {
			configure(installer)
		}
	}
	return installer
}

var _featureTag byte
