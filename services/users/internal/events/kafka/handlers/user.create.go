package handlers

import (
	"context"
	databaseusers "crm/services/users/database"
	events "crm/services/users/internal/events/kafka"
	"crm/services/users/internal/services/users"
	"encoding/json"
)

func NewUserCreatedHandler(userService *users.User) func(ctx context.Context, payload []byte) error {
	return func(ctx context.Context, payload []byte) error {
		var ev events.CreateUserEvent
		if err := json.Unmarshal(payload, &ev); err != nil {
			return err
		}
		user := &databaseusers.Users{
			Email:     ev.Email,
			Password:  ev.Password,
			Nickname:  ev.Nickname,
			FirstName: ev.FirstName,
			LastName:  ev.LastName,
			Country:   ev.Country,
		}
		_, err := userService.Create(ctx, user)
		return err
	}
}
