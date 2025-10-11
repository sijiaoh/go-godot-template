package utils_test

import (
	"github.com/sijiaoh/go-godot-template/game_server/testutils"
	"github.com/sijiaoh/go-godot-template/game_server/utils"
)

func ExampleDeps_WithEntTx() {
	testDeps := testutils.NewTestDeps()
	deps := testDeps.Deps

	err := deps.WithEntTx(func(txDeps *utils.Deps) error {
		_, err := txDeps.EntClient.User.Create().SetName("Foo").Save(txDeps.Ctx)
		if err != nil {
			panic(err)
		}

		_, err = txDeps.EntClient.User.Create().SetName("Foo").Save(txDeps.Ctx)
		if err != nil {
			panic(err)
		}

		return nil
	})
	if err != nil {
		panic(err)
	}
}

func ExampleDeps_WithEntTx_nest() {
	testDeps := testutils.NewTestDeps()
	deps := testDeps.Deps

	err := deps.WithEntTx(func(txDeps *utils.Deps) error {
		_, err := txDeps.EntClient.User.Create().SetName("Foo").Save(txDeps.Ctx)
		if err != nil {
			panic(err)
		}

		txDeps.WithEntTx(func(txDeps2 *utils.Deps) error {
			_, err = txDeps2.EntClient.User.Create().SetName("Foo").Save(txDeps2.Ctx)
			if err != nil {
				panic(err)
			}
			return nil
		})

		return nil
	})
	if err != nil {
		panic(err)
	}
}
