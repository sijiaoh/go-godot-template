package models

import (
	"github.com/google/uuid"
	"github.com/sijiaoh/go-godot-template/game_server/ent"
	"github.com/sijiaoh/go-godot-template/game_server/ent/transfercode"
	"github.com/sijiaoh/go-godot-template/game_server/utils"
)

type TransferCode struct {
	EntTransferCode *ent.TransferCode

	ID   int
	Code string

	User *User `validate:"-"`
}

func CreateTransferCode(deps *utils.Deps, user *User) (*TransferCode, error) {
	code, err := GenerateTransferCode()
	if err != nil {
		return nil, err
	}

	entTC, err := deps.EntClient.TransferCode.Create().
		SetCode(code).
		SetUserID(user.ID).
		Save(deps.Ctx)
	if err != nil {
		return nil, err
	}

	tc := NewTransferCodeFromEnt(entTC, user)
	return tc, nil
}

func NewTransferCodeFromEnt(etc *ent.TransferCode, user *User) *TransferCode {
	return &TransferCode{
		EntTransferCode: etc,

		ID:   etc.ID,
		Code: etc.Code,

		User: user,
	}
}

func FindTransferCodeByCode(deps *utils.Deps, code string) (*TransferCode, error) {
	entTC, err := deps.EntClient.TransferCode.Query().
		Where(transfercode.CodeEQ(code)).
		Only(deps.Ctx)
	if err != nil {
		return nil, err
	}

	tc := NewTransferCodeFromEnt(entTC, nil)
	return tc, nil
}

func (tc *TransferCode) LoadUser(deps *utils.Deps) error {
	entUser, err := deps.EntClient.TransferCode.QueryUser(tc.EntTransferCode).Only(deps.Ctx)
	if err != nil {
		return err
	}

	tc.User = NewUserFromEnt(entUser)
	return nil
}

func (tc *TransferCode) Rotate(deps *utils.Deps) error {
	code, err := GenerateTransferCode()
	if err != nil {
		return err
	}

	etc, err := tc.EntTransferCode.Update().SetCode(code).Save(deps.Ctx)
	if err != nil {
		return err
	}

	tc.EntTransferCode = etc
	tc.Code = etc.Code
	return nil
}

func GenerateTransferCode() (string, error) {
	uuidv7, err := uuid.NewV7()
	if err != nil {
		return "", err
	}
	return uuidv7.String(), nil
}
