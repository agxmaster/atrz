package main

import (
	"github.com/agxmaster/atrz"
	"github.com/agxmaster/atrz/examples/custom_handler/dal"
	"github.com/agxmaster/atrz/examples/custom_handler/defined"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	z := server.Default()
	dal.MysqlSetup()
	atrz.SetUp(dal.DB, z, atrz.WithCustomHandler(defined.Custom{}))

	z.Spin()
}
