package models_test

import (
	"testing"

	"github.com/sijiaoh/go-godot-template/game_server/models"
	"github.com/sijiaoh/go-godot-template/game_server/testutils"
)

func TestAuthenticator_Login(t *testing.T) {
	deps := testutils.NewTestDeps()
	defer deps.Close()

	entClient := deps.EntClient

	user, err := models.CreateUser(&deps.Deps, "Foo")
	testutils.AssertNoError(t, err)

	testutils.AssertEqual(t, entClient.ClientSession.Query().CountX(deps.Ctx), 0)

	auth := models.NewAuthenticator()
	cs, err := auth.Login(&deps.Deps, user.TransferCode.Code)
	testutils.AssertNoError(t, err)
	testutils.AssertEqual(t, cs.User.ID, user.ID)

	testutils.AssertEqual(t, entClient.ClientSession.Query().CountX(deps.Ctx), 1)
}

func TestAuthenticator_Login_InvalidCode(t *testing.T) {
	deps := testutils.NewTestDeps()
	defer deps.Close()

	entClient := deps.EntClient

	_, err := models.CreateUser(&deps.Deps, "Foo")
	testutils.AssertNoError(t, err)

	testutils.AssertEqual(t, entClient.ClientSession.Query().CountX(deps.Ctx), 0)

	auth := models.NewAuthenticator()
	_, err = auth.Login(&deps.Deps, "Foo")
	testutils.AssertWithError(t, err)

	testutils.AssertEqual(t, entClient.ClientSession.Query().CountX(deps.Ctx), 0)
}
