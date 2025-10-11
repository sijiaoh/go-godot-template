package models

import (
	"github.com/google/uuid"
	"github.com/sijiaoh/go-godot-template/game_server/ent"
	"github.com/sijiaoh/go-godot-template/game_server/ent/clientsession"
	"github.com/sijiaoh/go-godot-template/game_server/utils"
)

type ClientSession struct {
	EntClientSession *ent.ClientSession

	ID    int
	Token string

	User *User `validate:"-"`
}

func CreateClientSession(deps *utils.Deps, user *User) (*ClientSession, error) {
	uuidv7, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	entCS, err := deps.EntClient.ClientSession.Create().
		SetToken(uuidv7.String()).
		SetUserID(user.ID).
		Save(deps.Ctx)
	if err != nil {
		return nil, err
	}

	cs := NewClientSessionFromEnt(entCS, user)
	user.ClientSessions = append(user.ClientSessions, cs)
	return cs, nil
}

func NewClientSessionFromEnt(ecs *ent.ClientSession, user *User) *ClientSession {
	return &ClientSession{
		EntClientSession: ecs,

		ID:    ecs.ID,
		Token: ecs.Token,

		User: user,
	}
}

func FindClientSessionByToken(deps *utils.Deps, token string) (*ClientSession, error) {
	entCS, err := deps.EntClient.ClientSession.Query().
		Where(clientsession.TokenEQ(token)).
		Only(deps.Ctx)
	if err != nil {
		return nil, err
	}

	cs := NewClientSessionFromEnt(entCS, nil)
	return cs, nil
}

func (cs *ClientSession) LoadUser(deps *utils.Deps) error {
	entUser, err := deps.EntClient.ClientSession.QueryUser(cs.EntClientSession).Only(deps.Ctx)
	if err != nil {
		return err
	}

	cs.User = NewUserFromEnt(entUser)
	return nil
}
