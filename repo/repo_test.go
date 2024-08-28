package repo

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	taskRepo TaskRepo
	userRepo UserRepo
}

func TestInit(t *testing.T) {

	//suite.Run(t, new(Suite))

}

func (s *Suite) SetupSuite() {

}
