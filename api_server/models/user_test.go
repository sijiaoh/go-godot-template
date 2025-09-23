package models_test

import (
	"testing"

	"github.com/sijiaoh/go-godot-template/api_server/models"
	"github.com/sijiaoh/go-godot-template/api_server/testutils"
)

func TestUser_Upsert_Create(t *testing.T) {
	deps := testutils.NewTestDeps()
	defer deps.Close()

	user := &models.User{Name: "Foo"}
	err := user.Upsert(&deps.Deps)
	testutils.AssertNoError(t, err)

	testutils.AssertRecordCount(t, deps.EntClient.User.Query(), deps.Ctx, 1)

	if len(user.Token) == 0 {
		t.Fatal("用户的Token生成失败")
	}
}

func TestUser_Upsert_Update(t *testing.T) {
	deps := testutils.NewTestDeps()
	defer deps.Close()

	user := CreateFooUser(t, deps)

	newName := "Bar"
	user.Name = newName

	err := user.Upsert(&deps.Deps)
	testutils.AssertNoError(t, err)

	testutils.AssertRecordCount(t, deps.EntClient.User.Query(), deps.Ctx, 1)

	newEntUser, err := deps.EntClient.User.Get(deps.Ctx, user.ID)
	testutils.AssertNoError(t, err)
	testutils.AssertEqual(t, newEntUser.Name, newName)
}

func CreateFooUser(t *testing.T, deps *testutils.TestDeps) *models.User {
	user := &models.User{Name: "Foo"}
	err := user.Upsert(&deps.Deps)
	testutils.AssertNoError(t, err)
	return user
}
