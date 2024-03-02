package core

import (
	"context"
	"errors"
	"fmt"
	tools "github.com/agxmaster/atrz/util"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"regexp"
	"sync"
)

type Iatr interface {
	Info(ctx context.Context, c *app.RequestContext)
	List(ctx context.Context, c *app.RequestContext)
	Delete(ctx context.Context, c *app.RequestContext)
	Create(ctx context.Context, c *app.RequestContext)
	BatchCreate(ctx context.Context, c *app.RequestContext)
	Update(ctx context.Context, c *app.RequestContext)
}

type RouteType string
type Method string

const (
	PageKey  = "page"
	IdKey    = "id"
	ModelKey = "model"

	RouterPrefix       = "/atr"
	CustomRouterPrefix = "/custom"

	RouteTypeInfo       RouteType = "INFO"
	RouteTypeList       RouteType = "LIST"
	RouteTypeCreate     RouteType = "CREATE"
	RouteTypeUpdate     RouteType = "UPDATE"
	RouteTypeDelete     RouteType = "DELETE"
	RouteTypeCrateBatch RouteType = "CREATE_BATCH"

	MethodPost Method = "POST"
	MethodGet  Method = "GET"
)

var RouteBindInstance RouteBind

type RouteBind struct {
	RouteMap   map[RouteType]app.HandlerFunc
	RegPathMap map[string]RouteType
	mu         sync.Mutex
}

func initRoute() {
	SetDefaultRoute()
	RouteBindInstance.SetRouteBind()
}

var defaultRoute []CustomRoute

func (r *RouteBind) SetRouteBind() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.RouteMap = make(map[RouteType]app.HandlerFunc)
	r.RegPathMap = make(map[string]RouteType)
	for _, route := range defaultRoute {
		r.RouteMap[route.RouteType] = route.Handler
		r.RegPathMap[GetRouteRegKey(route.Method, fmt.Sprintf("%s%s", Mp.RoutePrefix, route.RoutePath))] = route.RouteType
	}
}

func (r *RouteBind) AddBind(routeType RouteType, HandlerFunc app.HandlerFunc) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.RouteMap[routeType] = HandlerFunc
}

func (r *RouteBind) GetDefaultHandler(routeType RouteType) app.HandlerFunc {
	return r.RouteMap[routeType]
}

func AtrRouter(r *server.Hertz) {

	initRoute()
	dr := r.Group(Mp.RoutePrefix, func(c context.Context, ctx *app.RequestContext) {
		for reg, routeType := range RouteBindInstance.RegPathMap {
			if ok, _ := regexp.MatchString(reg, GetRouteKey(Method(ctx.Method()), tools.BytesToString(ctx.Path()))); ok {
				if !Mp.Rules.Allow(routeType, ctx.Param(ModelKey)) {
					Error(ctx, errors.New("model not support this route"))
					ctx.Abort()
				} else {
					return
				}
			}
		}
	})

	for _, route := range defaultRoute {
		if route.Method == MethodGet {
			dr.GET(route.RoutePath, route.Handler)
		}
		if route.Method == MethodPost {
			dr.POST(route.RoutePath, route.Handler)
		}
	}

	cr := r.Group(Mp.CustomRouterPrefix, func(c context.Context, ctx *app.RequestContext) {
		if models := (&MpModel{}).GetCustomMpModel(Method(ctx.Method()), tools.BytesToString(ctx.Path()), ctx.Param(ModelKey)); models == nil {
			Error(ctx, errors.New("model not support this route"))
			ctx.Abort()
		}
	})

	customConf := *GetCustomConf()
	for _, confs := range customConf {
		for _, confOne := range confs {
			if confOne.MatchRouters == nil || len(confOne.MatchRouters) == 0 {
				return
			}
			for _, routeConfig := range confOne.MatchRouters {

				handler := routeConfig.Handler
				if handler == nil {
					handler = RouteBindInstance.GetDefaultHandler(routeConfig.RouteType)
				}

				if routeConfig.Method == MethodGet {
					cr.GET(routeConfig.RoutePath, handler)
				}
				if routeConfig.Method == MethodPost {
					cr.POST(routeConfig.RoutePath, handler)
				}
			}

		}
	}

}

func SetDefaultRoute() {
	defaultRoute = []CustomRoute{
		{
			RouteType: RouteTypeInfo,
			RoutePath: "/info/:model/:id",
			Handler:   Mp.Iatr.Info,
			Method:    MethodGet,
		},
		{
			RouteType: RouteTypeList,
			RoutePath: "/list/:model/:page",
			Handler:   Mp.Iatr.List,
			Method:    MethodGet,
		},
		{
			RouteType: RouteTypeCreate,
			RoutePath: "/create/:model",
			Handler:   Mp.Iatr.Create,
			Method:    MethodPost,
		},
		{
			RouteType: RouteTypeDelete,
			RoutePath: "/delete/:model/:id",
			Handler:   Mp.Iatr.Delete,
			Method:    MethodPost,
		},
		{
			RouteType: RouteTypeCrateBatch,
			RoutePath: "/batch/create/:model",
			Handler:   Mp.Iatr.BatchCreate,
			Method:    MethodPost,
		},
		{
			RouteType: RouteTypeUpdate,
			RoutePath: "/update/:model/:id",
			Handler:   Mp.Iatr.Update,
			Method:    MethodPost,
		},
	}
}
