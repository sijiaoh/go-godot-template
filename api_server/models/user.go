package models

import (
	"github.com/google/uuid"
	"github.com/sijiaoh/go-godot-template/api_server/ent"
	ent_user "github.com/sijiaoh/go-godot-template/api_server/ent/user"
	"github.com/sijiaoh/go-godot-template/api_server/utils"
)

type User struct {
	ID    int
	Name  string `validate:"min_runes=1,max_runes=12"`
	Token string
}

func NewUserFromEnt(entUser *ent.User) *User {
	return &User{
		ID:    entUser.ID,
		Name:  entUser.Name,
		Token: entUser.Token,
	}
}

func FindUserByToken(deps *utils.Deps, token string) (*User, error) {
	entUser, err := deps.EntClient.User.Query().Where(ent_user.TokenEQ(token)).Only(deps.Ctx)
	if err != nil {
		return nil, err
	}

	return NewUserFromEnt(entUser), nil
}

func (u *User) Upsert(deps *utils.Deps) error {
	if u.Token == "" {
		uuidv7, err := uuid.NewV7()
		if err != nil {
			return err
		}
		u.Token = uuidv7.String()
	}

	id, err := deps.EntClient.User.Create().
		SetName(u.Name).
		SetToken(u.Token).
		OnConflict().
		UpdateNewValues().
		ID(deps.Ctx)
	if err != nil {
		return err
	}
	u.ID = id

	return nil
}
