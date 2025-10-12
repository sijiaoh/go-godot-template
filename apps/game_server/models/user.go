package models

import (
	"github.com/sijiaoh/go-godot-template/game_server/ent"
	"github.com/sijiaoh/go-godot-template/game_server/utils"
	"github.com/sijiaoh/go-godot-template/game_server/validators"
)

type User struct {
	EntUser *ent.User

	ID   int
	Name string `validate:"min_runes=1,max_runes=12"`

	ClientSessions []*ClientSession `validate:"-"`
	TransferCode   *TransferCode    `validate:"-"` // 生成用户时创建，必须存在
}

func NewUser(name string) *User {
	return &User{
		Name: name,
	}
}

func CreateUser(deps *utils.Deps, name string) (*User, error) {
	var user *User
	err := deps.WithEntTx(func(txDeps *utils.Deps) error {
		user = NewUser(name)

		err := user.Save(deps)
		if err != nil {
			return err
		}

		_, err = CreateTransferCode(deps, user)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return user, nil
}

func NewUserFromEnt(entUser *ent.User) *User {
	return &User{
		EntUser: entUser,

		ID:   entUser.ID,
		Name: entUser.Name,
	}
}

func (u *User) GetID() int { return u.ID }

func (u *User) EntObj() *ent.User { return u.EntUser }

func (u *User) Save(deps *utils.Deps) error {
	err := validators.Validate().Struct(u)
	if err != nil {
		return err
	}

	entUser, err := utils.Save(
		deps,
		u,
		deps.EntClient.User,
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

func (u *User) LoadTransferCode(deps *utils.Deps) error {
	entTransferCode, err := deps.EntClient.User.QueryTransferCode(u.EntUser).Only(deps.Ctx)
	if err != nil {
		return err
	}

	u.TransferCode = NewTransferCodeFromEnt(entTransferCode, u)
	return nil
}
