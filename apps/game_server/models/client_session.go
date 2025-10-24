package models

import (
	"github.com/google/uuid"
	"github.com/sijiaoh/go-godot-template/game_server/ent"
	"github.com/sijiaoh/go-godot-template/game_server/ent/clientsession"
	"github.com/sijiaoh/go-godot-template/game_server/utils"
)

type ClientSessionPreload int

const (
	ClientSessionPreloadUser ClientSessionPreload = iota
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
	return cs, nil
}

func NewClientSessionFromEnt(ecs *ent.ClientSession, user *User) *ClientSession {
	cs := &ClientSession{
		EntClientSession: ecs,

		ID:    ecs.ID,
		Token: ecs.Token,
	}

	if user != nil {
		cs.User = user
	} else if ecs.Edges.User != nil {
		cs.User = NewUserFromEnt(ecs.Edges.User)
	}

	return cs
}

func FindClientSessionByToken(deps *utils.Deps, token string, preload ...ClientSessionPreload) (*ClientSession, error) {
	q := deps.EntClient.ClientSession.Query().Where(clientsession.TokenEQ(token))

	for _, p := range preload {
		switch p {
		case ClientSessionPreloadUser:
			q.WithUser()
		}
	}

	entCS, err := q.Only(deps.Ctx)
	if err != nil {
		return nil, err
	}

	cs := NewClientSessionFromEnt(entCS, nil)
	return cs, nil
}
