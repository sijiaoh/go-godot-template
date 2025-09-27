package models_test

import (
	"testing"

	"github.com/sijiaoh/go-godot-template/api_server/models"
	"github.com/sijiaoh/go-godot-template/api_server/testutils"
)

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

	user := CreateFooUser(t, deps)

	newName := "Bar"
	user.Name = newName

	err := user.Save(&deps.Deps)
	testutils.AssertNoError(t, err)

	testutils.AssertEqual(t, deps.EntClient.User.Query().CountX(deps.Ctx), 1)

	newEntUser, err := deps.EntClient.User.Get(deps.Ctx, user.ID)
	testutils.AssertNoError(t, err)
	testutils.AssertEqual(t, newEntUser.Name, newName)
}

func CreateFooUser(t *testing.T, deps *testutils.TestDeps) *models.User {
	user := models.NewUser("Foo")
	err := user.Save(&deps.Deps)
	testutils.AssertNoError(t, err)
	return user
}
