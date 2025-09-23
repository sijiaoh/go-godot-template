package models_test

import (
	"testing"

	"github.com/sijiaoh/go-godot-template/api_server/models"
	"github.com/sijiaoh/go-godot-template/api_server/testutils"
)

func TestCreateUser(t *testing.T) {
	deps := testutils.NewTestDeps()
	defer deps.Close()

	user, err := models.CreateUser(&deps.Deps, "Foo")
	testutils.AssertNoError(t, err)

	if len(user.Token) == 0 {
		t.Fatal("用户的Token生成失败")
	}
}

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
	user, err := models.CreateUser(&deps.Deps, "Foo")
	testutils.AssertNoError(t, err)
	return user
}
