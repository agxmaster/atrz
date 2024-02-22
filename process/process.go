package process

import (
	"context"
	"github.com/agxmaster/atm/clause"
	"github.com/agxmaster/atrz/adapter"
	"github.com/agxmaster/atrz/core"
	"github.com/agxmaster/atrz/parse"
	"github.com/agxmaster/atrz/util"
)

type Process interface {
	Info(id int) (interface{}, error)
	List(params map[string]interface{}, page int) (interface{}, error)
	Delete(id int) error
	Create() error
	BatchCreate() error
	Update(id int) error
}

func ProcessFactory(ctx context.Context, c adapter.HertzCtxCore, modelName string) Process {

	var modelConf = (&core.MpModel{}).GetMpModel(core.Method(c.Method()), util.BytesToString(c.Path()), modelName)

	if modelConf == nil {
		modelConf = &core.MpModel{}
	}
	if modelConf.Model != nil {
		return &AtmWithModel{BaseStore{ModelConf: modelConf, ModelName: modelName, Ctx: ctx, C: c}}
	}
	return &AtmWithOutModel{BaseStore{ModelConf: modelConf, ModelName: modelName, Ctx: ctx, C: c}}
}

type BaseStore struct {
	ModelName string
	ModelConf *core.MpModel
	Ctx       context.Context
	C         adapter.HertzCtxCore
}

func (b BaseStore) ProcessList(params map[string]interface{}, page int) (clause.Clauses, error) {
	params = b.ModelConf.QueryChange(b.Ctx, params)
	return (&parse.ClauseParse{ModelType: b.ModelConf}).Parse(&params, page).GetClause()
}
