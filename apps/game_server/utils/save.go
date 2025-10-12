package utils

import (
	"context"

	"github.com/sijiaoh/go-godot-template/game_server/ent"
)

type Model[T any] interface {
	GetID() int
	EntObj() *T
}

type SaveBuilder[EntObj any, Mut ent.Mutation] interface {
	Mutation() Mut
	Save(ctx context.Context) (*EntObj, error)
}

type EntityClient[
	EntObj any,
	Mut ent.Mutation,
	CreateBuilder SaveBuilder[EntObj, Mut],
	UpdateBuilder SaveBuilder[EntObj, Mut],
] interface {
	Create() CreateBuilder
	UpdateOneID(id int) UpdateBuilder
}

// 用于实现模型Save函数的辅助函数
func Save[
	EntObj any,
	Mut ent.Mutation,
	CreateBuilder SaveBuilder[EntObj, Mut],
	UpdateBuilder SaveBuilder[EntObj, Mut],
](
	deps *Deps,
	model Model[EntObj],
	entityClient EntityClient[EntObj, Mut, CreateBuilder, UpdateBuilder],
	mutate func(m Mut),
) (*EntObj, error) {
	if model.EntObj() == nil {
		c := entityClient.Create()
		mutate(c.Mutation())
		return c.Save(deps.Ctx)
	} else {
		u := entityClient.UpdateOneID(model.GetID())
		mutate(u.Mutation())
		return u.Save(deps.Ctx)
	}
}
