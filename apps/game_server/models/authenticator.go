package models

import (
	"github.com/sijiaoh/go-godot-template/game_server/utils"
)

type Authenticator struct{}

func NewAuthenticator() *Authenticator {
	return &Authenticator{}
}

func (a *Authenticator) Login(d *utils.Deps, transferCode string) (*ClientSession, error) {
	tc, err := FindTransferCodeByCode(d, transferCode)
	if err != nil {
		return nil, err
	}

	err = tc.LoadUser(d)
	if err != nil {
		return nil, err
	}

	// NOTE: 根据需要，可以删除现有的ClientSession保持登录的唯一性
	cs, err := CreateClientSession(d, tc.User)
	if err != nil {
		return nil, err
	}

	return cs, nil
}
