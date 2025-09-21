package models_test

import (
	"testing"

	"github.com/sijiaoh/go-godot-template/api_server/models"
	"github.com/sijiaoh/go-godot-template/api_server/testutils"
)

func TestUser_ApplyUpdate(t *testing.T) {
	deps := testutils.NewTestDeps()
	defer deps.Close()

	user := CreateUser(t, deps)
	newName := "Bar"
	user.Name = &newName

	err := user.ApplyUpdate(&deps.Deps)
	testutils.AssertNoError(t, err)

	testutils.AssertRecordCount(t, deps.EntClient.User.Query(), deps.Ctx, 1)

	newEntUser, err := deps.EntClient.User.Get(deps.Ctx, user.ID)
	testutils.AssertNoError(t, err)
	testutils.AssertStrPtrEqual(t, newEntUser.Name, &newName)
}

func CreateUser(t *testing.T, deps *testutils.TestDeps) *models.User {
	entUser, err := deps.EntClient.User.Create().SetName("Foo").Save(deps.Ctx)
	testutils.AssertNoError(t, err)
	return models.NewUserFromEnt(entUser)
}
