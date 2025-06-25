package tests

import (
	"crm/go_libs/helpers/pointercheck"
	usersv3 "crm/proto/gen/go/users/v3"
	"crm/services/users/tests/suite"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
	"math/rand"
	"testing"
)

func TestCreateGetUser_Happy(t *testing.T) {
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

func TestGetUsers_Happy(t *testing.T) {
	ctx, st := suite.New(t)
	users, err := st.UserClient.GetUsers(ctx, &usersv3.GetUsersRequest{})
	require.NoError(t, err)
	require.NotNil(t, users)
}

func TestDeleteUser_Happy(t *testing.T) {
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
	deleteUser, err := st.UserClient.DeleteUser(ctx, &usersv3.DeleteUserRequest{Id: createUser.Id})
	require.NoError(t, err)
	require.NotNil(t, deleteUser)
}

func TestDeleteUsers_Fail(t *testing.T) {
	ctx, st := suite.New(t)
	deleteUser, err := st.UserClient.DeleteUser(ctx, &usersv3.DeleteUserRequest{Id: 0})
	require.Error(t, err)
	require.Nil(t, deleteUser)
	require.ErrorAs(t, err, &gorm.ErrRecordNotFound)
}

func TestCreateUserDuplicate_Fail(t *testing.T) {
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
	createUser, err = st.UserClient.CreateUser(ctx, &user)
	require.ErrorAs(t, err, &gorm.ErrDuplicatedKey)
	require.Nil(t, createUser)
}

func TestGetUser_Fail(t *testing.T) {
	ctx, st := suite.New(t)
	user := usersv3.GetUserRequest{
		Id: 0,
	}
	getUser, err := st.UserClient.GetUser(ctx, &usersv3.GetUserRequest{Id: user.Id})
	require.Error(t, err)
	require.Nil(t, getUser)
	require.ErrorAs(t, err, &gorm.ErrRecordNotFound)
}

func TestUpdateUser_Happy(t *testing.T) {
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
	updateUser, err := st.UserClient.UpdateUser(ctx, &usersv3.UpdateUserRequest{
		Id:        createUser.Id,
		Firstname: gofakeit.FirstName(),
		Lastname:  gofakeit.LastName(),
		Nickname:  gofakeit.Name(),
		Email:     gofakeit.Email(),
		Country:   1,
	})
	require.NoError(t, err)
	require.NotNil(t, updateUser)
}

func TestUpdateUser_Fail(t *testing.T) {
	ctx, st := suite.New(t)
	updateUser, err := st.UserClient.UpdateUser(ctx, &usersv3.UpdateUserRequest{
		Id:        0,
		Firstname: gofakeit.FirstName(),
		Lastname:  gofakeit.LastName(),
		Nickname:  gofakeit.Name(),
		Email:     gofakeit.Email(),
		Country:   1,
	})
	require.Error(t, err)
	require.Nil(t, updateUser)
	require.ErrorAs(t, err, &gorm.ErrRecordNotFound)
}

func TestSearchUsers_Fail(t *testing.T) {
	ctx, st := suite.New(t)
	user := usersv3.SearchUsersRequest{
		Firstname: pointercheck.ToPtr(gofakeit.FirstName()),
		Lastname:  pointercheck.ToPtr(gofakeit.LastName()),
		Nickname:  pointercheck.ToPtr(gofakeit.Name()),
		Email:     pointercheck.ToPtr(gofakeit.Email()),
	}
	u1, _ := st.UserClient.Search(ctx, &user)
	require.Empty(t, u1.Users)

	user.Lastname = nil
	u2, _ := st.UserClient.Search(ctx, &user)
	require.Empty(t, u2.Users)

	user.Email = nil
	u3, _ := st.UserClient.Search(ctx, &user)
	require.Empty(t, u3.Users)

	user.Nickname = nil
	u4, _ := st.UserClient.Search(ctx, &user)
	require.Empty(t, u4.Users)
}

func TestSearchUsers_Happy(t *testing.T) {
	ctx, st := suite.New(t)
	userC := usersv3.CreateUserRequest{
		Firstname: gofakeit.FirstName(),
		Lastname:  gofakeit.LastName(),
		Nickname:  gofakeit.Name(),
		Email:     gofakeit.Email(),
		Password:  gofakeit.Password(true, true, true, true, true, 32),
		Country:   usersv3.Country(rand.Intn(6)),
	}
	user := usersv3.SearchUsersRequest{
		Firstname: pointercheck.ToPtr(userC.Firstname),
		Lastname:  pointercheck.ToPtr(userC.Lastname),
		Nickname:  pointercheck.ToPtr(userC.Nickname),
		Email:     pointercheck.ToPtr(userC.Email),
	}
	find, err := st.UserClient.CreateUser(ctx, &userC)
	require.NoError(t, err)
	u1, _ := st.UserClient.Search(ctx, &user)
	require.Equal(t, u1.Users[0].Id, find.Id)

	user.Lastname = nil
	u2, _ := st.UserClient.Search(ctx, &user)
	require.Equal(t, u2.Users[0].Id, find.Id)

	user.Email = nil
	u3, _ := st.UserClient.Search(ctx, &user)
	require.Equal(t, u3.Users[0].Id, find.Id)

	user.Nickname = nil
	u4, _ := st.UserClient.Search(ctx, &user)
	require.Equal(t, u4.Users[0].Id, find.Id)
}
