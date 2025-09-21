package models

import (
	"errors"

	"github.com/sijiaoh/go-godot-template/api_server/ent"
	"github.com/sijiaoh/go-godot-template/api_server/utils"
)

type User struct {
	ID      int
	Name    *string
	entUser *ent.User
}

func FromEnt(entUser *ent.User) User {
	return User{
		ID:      entUser.ID,
		Name:    entUser.Name,
		entUser: entUser,
	}
}

func (u *User) ApplyUpdate(deps utils.Deps) error {
	if u.entUser == nil {
		return errors.New("user.entUser is nil")
	}

	update := deps.EntClient.User.UpdateOne(u.entUser)
	if utils.StrPtrEq(u.Name, u.entUser.Name) {
		if u.Name != nil {
			update.SetName(*u.Name)
		} else {
			update.ClearName()
		}
	}

	entUser, err := update.Save(deps.Ctx)
	if err != nil {
		return err
	}

	u.entUser = entUser
	return nil
}
