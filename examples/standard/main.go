package main

import (
	"github.com/agxmaster/atrz"
	"github.com/agxmaster/atrz/core"
	"github.com/agxmaster/atrz/examples/standard/dal"
	"github.com/agxmaster/atrz/examples/standard/defined"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	z := server.Default()
	dal.MysqlSetup()
	atrz.SetUp(dal.DB, z,
		atrz.WithModelConfig(defined.ConfigModelMap),
		atrz.WithRules(core.Rules{
			{Table: []string{"student"}, RouteTypes: []core.RouteType{"*"}, RuleType: core.RuleTypeAllow},
			{Table: []string{"*"}, RouteTypes: []core.RouteType{"*"}, RuleType: core.RuleTypeDeny},
		}),
	)

	z.Spin()
}
