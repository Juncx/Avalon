package items

import (
	"Avalon/avalon/types"
	"Avalon/dbconn"

	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func itemTablesCheck(engine *xorm.Engine) error {
	item := &types.ItemInfo{}
	exist, err := engine.IsTableExist(item)
	if err != nil {
		panic(err)
	}
	if !exist {
		return engine.CreateTables(item)
	}

	return nil
}

func init() {
	var err error
	engine, err = dbconn.GetDBEngine()
	if err != nil {
		panic(err)
	}

	err = itemTablesCheck(engine)
	if err != nil {
		panic(err)
	}
}
