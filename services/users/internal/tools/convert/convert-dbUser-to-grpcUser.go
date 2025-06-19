package convert

import (
	usersv3 "crm/proto/gen/go/users/v3"
	databaseusers "crm/services/users/database"
)

func UserDbtoGrpcUser(user *databaseusers.Users) *usersv3.User {
	return &usersv3.User{
		Id:        user.ID,
		Firstname: user.FirstName,
		Lastname:  user.LastName,
		Nickname:  user.Nickname,
		Email:     user.Email,
		Country:   usersv3.Country(user.Country),
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}
}

func UserGrpctoDbUser(user *usersv3.User) *databaseusers.Users {
	return &databaseusers.Users{
		ID:        user.Id,
		FirstName: user.Firstname,
		LastName:  user.Lastname,
		Nickname:  user.Nickname,
		Email:     user.Email,
		Country:   int(user.Country),
	}
}
