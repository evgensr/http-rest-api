package model_test

import (
	"testing"

	"github.com/evgensr/http-rest-api/internal/app/model"
	"github.com/stretchr/testify/assert"
)

func TestUser_Validate(t *testing.T) {
	u := model.TestUser(t)
	assert.NoError(t, u.Validate())
}

func TestUser_BeforeCreate(t *testing.T) {

	testCases := []struct {
		name    string
		u       func() *model.User
		isValid bool
	}{
		{
			name: "Valid",
			u: func() *model.User {
				return model.TestUser(t)
			},
			isValid: true,
		},
		{
			name: "Empty email",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = ""
				return u
			},
			isValid: false,
		},
		{
			name: "Empty Invalid",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = "Invalid"
				return u
			},
			isValid: false,
		},
		{
			name: "Short password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = "short"
				return u
			},
			isValid: false,
		},
		{
			name: "Empty password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = ""
				return u
			},
			isValid: false,
		},
		{
			name: "With encrypt password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = ""
				u.EncryptedPassword = "encrypt"
				return u
			},
			isValid: true,
		},
	}
	for _, tc := range testCases {
		t.Run(t.Name(), func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().Validate())
			} else {
				assert.Error(t, tc.u().Validate())
			}
		})

	}

	// u := model.TestUser(t)
	// assert.NoError(t, u.BeforeCreate())
	// assert.NotEmpty(t, u.EncryptedPassword)
}
