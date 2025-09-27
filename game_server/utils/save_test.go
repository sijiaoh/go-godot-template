package utils_test

import (
	"github.com/sijiaoh/go-godot-template/api_server/ent"
	"github.com/sijiaoh/go-godot-template/api_server/models"
	"github.com/sijiaoh/go-godot-template/api_server/testutils"
	"github.com/sijiaoh/go-godot-template/api_server/utils"
	"github.com/sijiaoh/go-godot-template/api_server/validators"
)

func ExampleSave() {
	deps := testutils.NewTestDeps()
	user := models.NewUser("Foo")

	err := validators.Validate().Struct(user)
	if err != nil {
		panic(err)
	}

	entUser, err := utils.Save(
		&deps.Deps,
		user.EntUser,
		func() *ent.UserCreate { return deps.EntClient.User.Create() },
		func() *ent.UserUpdateOne { return user.EntUser.Update() },
		func(mutation *ent.UserMutation) {
			if user.EntUser == nil || user.Name != user.EntUser.Name {
				mutation.SetName(user.Name)
			}
		},
	)
	if err != nil {
		panic(err)
	}

	user.ID = entUser.ID
	user.EntUser = entUser
}
