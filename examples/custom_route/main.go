package main

import (
	"github.com/agxmaster/atrz"
	"github.com/agxmaster/atrz/examples/custom_route/dal"
	"github.com/agxmaster/atrz/examples/custom_route/defined"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	z := server.Default()
	dal.MysqlSetup()
	atrz.SetUp(dal.DB, z, atrz.WithCustomRoute(defined.ConfigCustomModel))

	z.Spin()
}
