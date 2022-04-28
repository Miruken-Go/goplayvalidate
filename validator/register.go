package validator

import "github.com/miruken-go/miruken"

func WithGoPlaygroundValidation() miruken.Installer {
	return miruken.InstallerFunc(func(registration *miruken.Registration) {
		if registration.CanInstall(&_registrationTag) {
			registration.Install(miruken.WithValidation(false))
			registration.AddHandlerTypes(miruken.TypeOf[*IntegrationValidator]())
		}
	})
}

var _registrationTag any
