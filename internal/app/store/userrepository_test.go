package store_test

import (
	"log"
	"testing"

	"github.com/evgensr/http-rest-api/internal/app/model"
	"github.com/evgensr/http-rest-api/internal/app/store"
	"github.com/stretchr/testify/assert"
)

func TestUserRepositore_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u, err := s.User().Create(&model.User{
		Email: "user@gmail.com",
	})
	if err != nil {
		log.Println(err)
	}

	assert.NoError(t, err)
	assert.NotNil(t, u)

}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	email := "user@gmail.com"

	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	u, err := s.User().Create(&model.User{
		Email: email,
	})
	assert.NoError(t, err)
	assert.NotNil(t, u)

}
