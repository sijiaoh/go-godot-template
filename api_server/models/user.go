package models

import (
	"errors"

	"github.com/google/uuid"
	"github.com/sijiaoh/go-godot-template/api_server/ent"
	ent_user "github.com/sijiaoh/go-godot-template/api_server/ent/user"
	"github.com/sijiaoh/go-godot-template/api_server/utils"
)

type User struct {
	entUser *ent.User

	ID    int
	Name  string `validate:"min_runes=1,max_runes=12"`
	Token string
}

func NewUserFromEnt(entUser *ent.User) *User {
	return &User{
		entUser: entUser,

		ID:    entUser.ID,
		Name:  entUser.Name,
		Token: entUser.Token,
	}
}

func CreateUser(deps *utils.Deps, name string) (*User, error) {
	uuidv7, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	entUser, err := deps.EntClient.User.Create().
		SetName(name).
		SetToken(uuidv7.String()).
		Save(deps.Ctx)
	if err != nil {
		return nil, err
	}

	return NewUserFromEnt(entUser), nil
}

func FindUserByToken(deps *utils.Deps, token string) (*User, error) {
	entUser, err := deps.EntClient.User.Query().Where(ent_user.TokenEQ(token)).Only(deps.Ctx)
	if err != nil {
		return nil, err
	}

	return NewUserFromEnt(entUser), nil
}

func (u *User) ApplyUpdate(deps *utils.Deps) error {
	if u.entUser == nil {
		return errors.New("user.entUser is nil")
	}

	update := deps.EntClient.User.UpdateOne(u.entUser)

	if u.Name != u.entUser.Name {
		update.SetName(u.Name)
	}
	if u.Token != u.entUser.Token {
		update.SetToken(u.Token)
	}

	entUser, err := update.Save(deps.Ctx)
	if err != nil {
		return err
	}

	u.entUser = entUser
	return nil
}
