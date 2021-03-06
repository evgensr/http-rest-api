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

func TestUserRepository_Find(t *testing.T) {
	s := teststore.New()
	u1 := model.TestUser(t)
	s.User().Create(u1)
	u2, err := s.User().Find(u1.ID)
	log.Println(u2)
	assert.NoError(t, err)
	assert.NotNil(t, u2)

}

func TestUserRepository_FindByEmail(t *testing.T) {

	s := teststore.New()
	u1 := model.TestUser(t)
	r, err := s.User().FindByEmail(u1.Email)
	log.Println(r)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.User().Create(u1)

	u2, err := s.User().FindByEmail(u1.Email)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
	// log.Println(s)

}
