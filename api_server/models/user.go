package models

import (
	"github.com/sijiaoh/go-godot-template/api_server/ent"
	"github.com/sijiaoh/go-godot-template/api_server/utils"
	"github.com/sijiaoh/go-godot-template/api_server/validators"
)

type User struct {
	EntUser *ent.User

	ID   int
	Name string `validate:"min_runes=1,max_runes=12"`

	ClientSessions []*ClientSession
}

func NewUser(name string) *User {
	return &User{
		Name: name,
	}
}

func NewUserFromEnt(entUser *ent.User) *User {
	return &User{
		EntUser: entUser,

		ID:   entUser.ID,
		Name: entUser.Name,
	}
}

func (u *User) Save(deps *utils.Deps) error {
	err := validators.Validate().Struct(u)
	if err != nil {
		return err
	}

	entUser, err := utils.Save(
		deps,
		u.EntUser,
		func() *ent.UserCreate { return deps.EntClient.User.Create() },
		func() *ent.UserUpdateOne { return u.EntUser.Update() },
		func(mutation *ent.UserMutation) {
			if u.EntUser == nil || u.Name != u.EntUser.Name {
				mutation.SetName(u.Name)
			}
		},
	)
	if err != nil {
		return err
	}

	u.ID = entUser.ID
	u.EntUser = entUser

	return nil
}

func (u *User) LoadClientSessions(deps *utils.Deps) error {
	entClientSessions, err := deps.EntClient.User.QueryClientSessions(u.EntUser).All(deps.Ctx)
	if err != nil {
		return err
	}

	for _, ecs := range entClientSessions {
		cs := NewClientSessionFromEnt(ecs, u)
		u.ClientSessions = append(u.ClientSessions, cs)
	}

	return nil
}
