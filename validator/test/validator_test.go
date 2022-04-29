package test

import (
	"github.com/miruken-go/go-playground/validator"
	"github.com/miruken-go/miruken"
	"github.com/stretchr/testify/suite"
	"testing"
)

//go:generate $GOPATH/bin/mirukentypes -tests

// User contains user information.
type User struct {
	Id             int
	FirstName      string    `validate:"required"`
	LastName       string    `validate:"required"`
	Age            uint8     `validate:"gte=0,lte=130"`
	Email          string    `validate:"required,email"`
	FavouriteColor string    `validate:"iscolor"`
	Addresses      []Address `validate:"required,dive,required"`
}

// Address houses a users address information.
type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
	Phone  string `validate:"required"`
}

// Command to create a User.
type CreateUser struct {
	User User
}

// UserHandler handles User commands.
type UserHandler struct {
	userId int
}

func (u *UserHandler) CreateUser(
	_ *miruken.Handles, create *CreateUser,
) User {
	user := create.User
	u.userId++
	user.Id = u.userId
	return user
}

type ValidatorTestSuite struct {
	suite.Suite
	handler miruken.Handler
}

func (suite *ValidatorTestSuite) SetupTest() {
	suite.handler = miruken.NewRegistration(
		miruken.WithHandlerTypes(HandlerTestTypes...),
		validator.WithGoPlaygroundValidation(),
	).Build()
}

func (suite *ValidatorTestSuite) TestValidator() {
	suite.Run("Valid Target", func() {
		var user User
		create := CreateUser{
			User{
				FirstName:      "Badger",
				LastName:       "Smith",
				Age:            52,
				Email:          "Badger.Smith@gmail.com",
				FavouriteColor: "#000",
				Addresses:      []Address{
					{
						Street: "Eavesdown Docks",
						City:   "Rockwall",
						Planet: "Persphone",
						Phone:  "none",
					},
				},
			},
		}
		if err := miruken.Invoke(suite.handler, &create, &user); err != nil {
			suite.Fail("unexpected error", err.Error())
		}
	})

	suite.Run("Invalid Target", func() {
		var user User
		create := CreateUser{
			User{
				Age: 200,
				FavouriteColor: "#000-",
				Addresses:[]Address{
					{
					},
				},
			},
		}
		if err := miruken.Invoke(suite.handler, &create, &user); err != nil {
		} else {
			suite.Fail("expected error")
		}
	})
}

func TestValidateTestSuite(t *testing.T) {
	suite.Run(t, new(ValidatorTestSuite))
}
