package raw

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/steebchen/prisma-client-go-1/engine"
	"github.com/steebchen/prisma-client-go-1/runtime/builder"
)

type Raw struct {
	Engine engine.Engine
}

func raw(engine engine.Engine, action string, query string, params ...interface{}) builder.Query {
	q := builder.NewQuery()
	q.Engine = engine
	q.Operation = "mutation"
	q.Method = action

	q.Inputs = append(q.Inputs, builder.Input{
		Name:  "query",
		Value: query,
	})

	// convert params to a string with an array
	var newParams = "["
	for i, param := range params {
		if i > 0 {
			newParams += ","
		}
		if date, ok := param.(time.Time); ok {
			data, err := json.Marshal(date)
			if err != nil {
				panic(err)
			}
			newParams += fmt.Sprintf(`{"prisma__type":"date","prisma__value":%s}`, string(data))
		} else {
			newParams += string(builder.Value(param))
		}
	}
	newParams += "]"

	q.Inputs = append(q.Inputs, builder.Input{
		Name:  "parameters",
		Value: newParams,
	})

	return q
}
