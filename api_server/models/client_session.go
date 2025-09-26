package models

import (
	"github.com/google/uuid"
	"github.com/sijiaoh/go-godot-template/api_server/ent"
	"github.com/sijiaoh/go-godot-template/api_server/utils"
)

type ClientSession struct {
	EntClientSession *ent.ClientSession

	ID    int
	Token string

	User *User
}

func NewClientSession(user *User) (*ClientSession, error) {
	cs := &ClientSession{
		User: user,
	}

	uuidv7, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	cs.Token = uuidv7.String()

	return cs, nil
}

func NewClientSessionFromEnt(ecs *ent.ClientSession, user *User) *ClientSession {
	return &ClientSession{
		EntClientSession: ecs,

		Token: ecs.Token,

		User: user,
	}
}

func (cs *ClientSession) LoadUser(deps *utils.Deps) error {
	entUser, err := deps.EntClient.ClientSession.QueryUser(cs.EntClientSession).Only(deps.Ctx)
	if err != nil {
		return err
	}

	cs.User = NewUserFromEnt(entUser)
	return nil
}
