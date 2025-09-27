package utils

import (
	"context"

	"github.com/sijiaoh/go-godot-template/game_server/ent"
)

type SaveBuilder[EntObj any, M ent.Mutation] interface {
	Mutation() M
	Save(ctx context.Context) (*EntObj, error)
}

// 用于实现模型Save函数的辅助函数
func Save[EntObj any, M ent.Mutation, CreateBuilder SaveBuilder[EntObj, M], UpdateBuilder SaveBuilder[EntObj, M]](
	deps *Deps,
	entObj *EntObj,
	createFn func() CreateBuilder,
	updateFn func() UpdateBuilder,
	mutate func(m M),
) (*EntObj, error) {
	if entObj == nil {
		c := createFn()
		mutate(c.Mutation())
		return c.Save(deps.Ctx)
	} else {
		u := updateFn()
		mutate(u.Mutation())
		return u.Save(deps.Ctx)
	}
}
