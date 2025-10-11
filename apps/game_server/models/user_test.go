package models_test

import (
	"testing"

	"github.com/sijiaoh/go-godot-template/game_server/models"
	"github.com/sijiaoh/go-godot-template/game_server/testutils"
)

func TestUser_Create(t *testing.T) {
	deps := testutils.NewTestDeps()
	defer deps.Close()

	user, err := models.CreateUser(&deps.Deps, "Foo")
	testutils.AssertNoError(t, err)
	testutils.AssertEqual(t, deps.EntClient.User.Query().CountX(deps.Ctx), 1)

	if user.TransferCode == nil {
		t.Fatal("创建用户时，需要同时创建转移码")
	}
	testutils.AssertEqual(t, deps.EntClient.TransferCode.Query().CountX(deps.Ctx), 1)
}

func TestUser_Save_Create(t *testing.T) {
	deps := testutils.NewTestDeps()
	defer deps.Close()

	user := models.NewUser("Foo")
	err := user.Save(&deps.Deps)
	testutils.AssertNoError(t, err)

	testutils.AssertEqual(t, deps.EntClient.User.Query().CountX(deps.Ctx), 1)
}

func TestUser_Save_Update(t *testing.T) {
	deps := testutils.NewTestDeps()
	defer deps.Close()

	user, err := models.CreateUser(&deps.Deps, "Foo")
	testutils.AssertNoError(t, err)

	newName := "Bar"
	user.Name = newName

	err = user.Save(&deps.Deps)
	testutils.AssertNoError(t, err)

	testutils.AssertEqual(t, deps.EntClient.User.Query().CountX(deps.Ctx), 1)

	newEntUser, err := deps.EntClient.User.Get(deps.Ctx, user.ID)
	testutils.AssertNoError(t, err)
	testutils.AssertEqual(t, newEntUser.Name, newName)
}
