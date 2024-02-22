package handler

import (
	"context"
	"github.com/agxmaster/atrz/core"
	"github.com/agxmaster/atrz/process"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
)

type Atr struct {
}

func (a *Atr) Info(ctx context.Context, c *app.RequestContext) {

	id, err := strconv.Atoi(c.Param(core.IdKey))
	if err != nil {
		core.Error(c, err)
		return
	}

	res, err := process.ProcessFactory(ctx, c, c.Param(core.ModelKey)).Info(id)
	if err != nil {
		core.Error(c, err)
		return
	}

	core.Success(c, res)
}

func (a *Atr) List(ctx context.Context, c *app.RequestContext) {

	var params map[string]interface{}

	page, err := strconv.Atoi(c.Param(core.PageKey))

	if err != nil {
		core.Error(c, err)
		return
	}

	err = c.BindQuery(&params)
	if err != nil {
		core.Error(c, err)
		return
	}
	res, err := process.ProcessFactory(ctx, c, c.Param(core.ModelKey)).List(params, page)

	if err != nil {
		core.Error(c, err)
		return
	}
	core.Success(c, res)

}

func (a *Atr) Delete(ctx context.Context, c *app.RequestContext) {

	id, err := strconv.Atoi(c.Param(core.IdKey))
	if err != nil {
		core.Error(c, err)
		return
	}

	err = process.ProcessFactory(ctx, c, c.Param(core.ModelKey)).Delete(id)
	if err != nil {
		core.Error(c, err)
		return
	}
	core.Success(c, nil)
}

func (a *Atr) Create(ctx context.Context, c *app.RequestContext) {
	//curl -X POST 127.0.0.1:8888/create/student -H 'Content-Type: application/json' -d "{\"name\":\"name1\",\"gender\":100,\"age\":100,\"class\":\"aaa\"}"

	err := process.ProcessFactory(ctx, c, c.Param(core.ModelKey)).Create()

	if err != nil {
		core.Error(c, err)
		return
	}
	core.Success(c, nil)

}

func (a *Atr) BatchCreate(ctx context.Context, c *app.RequestContext) {

	//curl -X POST 127.0.0.1:8888/batch/create/student -H 'Content-Type: application/json' -d "[{\"name\":\"name1\",\"gender\":100,\"age\":100,\"class\":\"aaa\"},{\"name\":\"name2\",\"gender\":100,\"age\":100,\"class\":\"aaa\"}]"
	err := process.ProcessFactory(ctx, c, c.Param(core.ModelKey)).BatchCreate()
	if err != nil {
		core.Error(c, err)
		return
	}
	core.Success(c, nil)

}

func (a *Atr) Update(ctx context.Context, c *app.RequestContext) {

	id, err := strconv.Atoi(c.Param(core.IdKey))
	if err != nil {
		core.Error(c, err)
		return
	}

	err = process.ProcessFactory(ctx, c, c.Param(core.ModelKey)).Update(id)
	if err != nil {
		core.Error(c, err)
		return
	}
	core.Success(c, nil)

}
