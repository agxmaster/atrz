package main

import (
	"github.com/agxmaster/atrz"
	"github.com/agxmaster/atrz/examples/bind_model/dal"
	"github.com/agxmaster/atrz/examples/bind_model/defined"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	z := server.Default()
	dal.MysqlSetup()
	atrz.SetUp(dal.DB, z, atrz.WithModelConfig(defined.ConfigModelMap))

	z.Spin()
}
