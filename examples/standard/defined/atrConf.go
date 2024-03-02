package defined

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/agxmaster/atm/clause"
	"github.com/agxmaster/atrz/core"
	"github.com/agxmaster/atrz/examples/standard/model"
	"github.com/agxmaster/atrz/util"
	"gorm.io/gorm"
	"reflect"
)

type StudentCustomScope struct {
	MinAge int `json:"minAge"`
	MaxAge int `json:"maxAge"`
}

func ConfigModelMap() map[string]core.MpModel {
	return map[string]core.MpModel{
		"student": {
			Model: reflect.TypeOf(model.Student{}),
			ColumnMapping: map[string]string{
				"gender": "sex",
			},
			//Select: []string{"name"},
			Select:              []string{"name", "sex", "class", "created_at", "updated_at"},
			Hidden:              []string{"updated_at"},
			AddColumns:          StudentAddColumns,
			ListWithCustomScope: StudentListWithCustomScope,
			CreateParamsHandler: StudentCreateParamsHandler,
			UpdateParamsHandler: StudentUpdateParamsHandler,
			//UpdateParamsHandler: nil,
			//Aggregations:        []core.Aggregation{{core.AggregationTypeCount, "age", "count_name"}},
			CanGroupColumn: []string{"name", "class"},
		},
	}
}

func StudentListWithCustomScope(ctx context.Context, customJson []byte) (clause.Scope, error) {
	var studentCustomStruct StudentCustomScope
	err := json.Unmarshal(customJson, &studentCustomStruct)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("custom scope Unmarshal error %+v", err))
	}

	if studentCustomStruct.MinAge >= studentCustomStruct.MaxAge {
		return nil, errors.New("custom scope validate error min age can't >= max age")
	}

	return func(db *gorm.DB) *gorm.DB {
		return db.Where("age >= ?", studentCustomStruct.MinAge).Where("age <= ?", studentCustomStruct.MaxAge)
	}, nil
}

func StudentUpdateParamsHandler(ctx context.Context, params map[string]interface{}) (map[string]interface{}, error) {
	modelParams, err := util.InterfaceToAny[model.Student](params)
	if err != nil {
		return nil, err
	}
	if modelParams.Age < 5 {
		return nil, errors.New("age 5 not allowed")
	} else if modelParams.Age < 6 {
		params["age"] = 8
	}
	return params, nil
}

func StudentCreateParamsHandler(ctx context.Context, params map[string]interface{}) (map[string]interface{}, error) {
	if age, ok := params["age"]; ok {
		if ageInt, ok := age.(int); ok {
			if ageInt < 5 {
				return nil, errors.New("age 5 not allowed")
			} else if ageInt < 6 {
				params["age"] = 8
			}
		}
	}
	return params, nil
}

var StudentAddColumns = []map[string]core.ColumnAddFunc{
	{
		"className": func(ctx context.Context, line map[string]interface{}, addKey string) interface{} {
			if class, ok := line["class"]; ok {
				if classId, ok := class.(string); ok {
					if classId == "3" {
						return "3班"
					}
				}
			}
			return "其他班级"
		},
		"name": func(ctx context.Context, line map[string]interface{}, addKey string) interface{} {
			if name, ok := line["name"]; ok {
				if nameCode, ok := name.(string); ok {
					if nameCode == "2" {
						return "小2"
					}
				}
			}
			return line["name"]
		},
	},
}
