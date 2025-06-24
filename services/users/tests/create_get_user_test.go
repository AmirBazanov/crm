package tests

import (
	usersv3 "crm/proto/gen/go/users/v3"
	"crm/services/users/tests/suite"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
)

func TestCreateGetUser_HappyTest(t *testing.T) {
	ctx, st := suite.New(t)

	user := usersv3.CreateUserRequest{
		Firstname: gofakeit.FirstName(),
		Lastname:  gofakeit.LastName(),
		Nickname:  gofakeit.Name(),
		Email:     gofakeit.Email(),
		Password:  gofakeit.Password(true, true, true, true, true, 32),
		Country:   usersv3.Country(rand.Intn(6)),
	}
	createUser, err := st.UserClient.CreateUser(ctx, &user)
	require.NoError(t, err)
	require.NotNil(t, createUser)
	getUser, err := st.UserClient.GetUser(ctx, &usersv3.GetUserRequest{Id: createUser.Id})
	require.NoError(t, err)
	require.NotNil(t, getUser)
}
