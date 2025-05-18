package mocks

import (
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type DBMock struct {
	mock.Mock
	*gorm.DB
}
