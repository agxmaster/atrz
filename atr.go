package atrz

import (
	"github.com/agxmaster/atrz/core"
	"github.com/agxmaster/atrz/handler"
	"github.com/cloudwego/hertz/pkg/app/server"
	"gorm.io/gorm"
)

func SetUp(db *gorm.DB, z *server.Hertz, opts ...core.Option) {

	core.Mp = &core.Config{
		Iatr:               &handler.Atr{},
		RoutePrefix:        core.RouterPrefix,
		CustomRouterPrefix: core.CustomRouterPrefix,
		SuccessCode:        core.CodeSuccess,
		ErrorCode:          core.CodeAtrError,
	}

	core.Db = db
	core.Mp.Apply(opts)
	core.CompleteConfig()
	core.AtrRouter(z)
}

func WithCustomHandler(handlerStruct core.Iatr) core.Option {
	return core.Option{Fn: func(o *core.Config) {
		o.Iatr = handlerStruct
	}}
}

func WithModelConfig(fn core.SetModel) core.Option {
	return core.Option{Fn: func(o *core.Config) {
		o.ModelMap = fn()
	}}
}

func WithCustomRoute(fn core.SetCustomModel) core.Option {
	return core.Option{Fn: func(o *core.Config) {
		o.CustomModelMap = fn()
		core.SetCustomModelCache()
	}}
}

func WithSuccessCode(code int) core.Option {
	return core.Option{Fn: func(o *core.Config) {
		o.SuccessCode = core.ResponseCode(code)
	}}
}

func WithErrorCode(code int) core.Option {
	return core.Option{Fn: func(o *core.Config) {
		o.ErrorCode = core.ResponseCode(code)
	}}
}

func WithCustomRouterPrefix(routerPrefix string) core.Option {
	return core.Option{Fn: func(o *core.Config) {
		o.CustomRouterPrefix = routerPrefix
	}}
}

func WithRouterPrefix(routerPrefix string) core.Option {
	return core.Option{Fn: func(o *core.Config) {
		o.RoutePrefix = routerPrefix
	}}
}

func WithRules(rules core.Rules) core.Option {
	return core.Option{Fn: func(o *core.Config) {
		o.Rules = rules
	}}
}
