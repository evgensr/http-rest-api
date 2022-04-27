package teststore_test

import (
	"log"
	"testing"

	"github.com/evgensr/http-rest-api/internal/app/model"
	"github.com/evgensr/http-rest-api/internal/app/store"
	"github.com/evgensr/http-rest-api/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepositore_Create(t *testing.T) {

	s := teststore.New()
	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)

}

func TestUserRepository_FindByEmail(t *testing.T) {

	s := teststore.New()
	email := "test@gmail.com"
	r, err := s.User().FindByEmail(email)
	log.Println(r)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Email = email
	s.User().Create(u)
	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
	// log.Println(s)

}
