package defined

import (
	"context"
	"github.com/agxmaster/atrz/core"
	"github.com/agxmaster/atrz/handler"
	"github.com/cloudwego/hertz/pkg/app"
)

type Custom struct {
	*handler.Atr
}

func (a Custom) List(ctx context.Context, c *app.RequestContext) {
	core.Success(c, "Custom")
}
