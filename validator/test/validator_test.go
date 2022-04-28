package test

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type ValidatorTestSuite struct {
	suite.Suite
}

func (suite *ValidatorTestSuite) SetupTest() {

}

func (suite *ValidatorTestSuite) TestValidator() {

}

func TestValidateTestSuite(t *testing.T) {
	suite.Run(t, new(ValidatorTestSuite))
}
