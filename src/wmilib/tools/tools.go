package tools

import (
	"github.com/StackExchange/wmi"
)

func ExecuteQuery(query interface{}, condition string) (err error) {
	queryString := wmi.CreateQuery(query, condition)
	err = wmi.Query(queryString, query)
	if err != nil {
		return err
	}
	return nil
}
