package utils_test

import (
	"github.com/sijiaoh/go-godot-template/game_server/ent"
	"github.com/sijiaoh/go-godot-template/game_server/models"
	"github.com/sijiaoh/go-godot-template/game_server/testutils"
	"github.com/sijiaoh/go-godot-template/game_server/utils"
	"github.com/sijiaoh/go-godot-template/game_server/validators"
)

func ExampleSave() {
	deps := testutils.NewTestDeps()
	defer deps.Close()

	user := models.NewUser("Foo")

	err := validators.Validate().Struct(user)
	if err != nil {
		panic(err)
	}

	entUser, err := utils.Save(
		&deps.Deps,
		user,
		deps.EntClient.User,
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
