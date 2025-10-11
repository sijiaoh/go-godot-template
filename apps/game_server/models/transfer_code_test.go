package models_test

import (
	"testing"

	"github.com/sijiaoh/go-godot-template/game_server/models"
	"github.com/sijiaoh/go-godot-template/game_server/testutils"
)

func TestTransferCode_Rotate(t *testing.T) {
	deps := testutils.NewTestDeps()
	defer deps.Close()

	user, err := models.CreateUser(&deps.Deps, "Foo")
	testutils.AssertNoError(t, err)

	tc := user.TransferCode
	oldCode := tc.Code

	err = tc.Rotate(&deps.Deps)
	testutils.AssertNoError(t, err)

	testutils.AssertEqual(t, deps.EntClient.TransferCode.Query().CountX(deps.Ctx), 1)
	testutils.AssertNotEqual(t, oldCode, tc.Code)
	testutils.AssertEqual(t, deps.EntClient.TransferCode.GetX(deps.Ctx, tc.ID).Code, tc.Code)
}

func CreateTransferCode(t *testing.T, deps *testutils.TestDeps, user *models.User) *models.TransferCode {
	tc, err := models.CreateTransferCode(&deps.Deps, user)
	testutils.AssertNoError(t, err)
	testutils.AssertEqual(t, deps.EntClient.TransferCode.Query().CountX(deps.Ctx), 1)
	return tc
}
