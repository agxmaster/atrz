package main

import (
	"github.com/agxmaster/atrz"
	"github.com/agxmaster/atrz/examples/simple/dal"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	z := server.Default()
	dal.MysqlSetup()
	atrz.SetUp(dal.DB, z)

	z.Spin()
}
