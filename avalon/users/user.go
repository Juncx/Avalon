package users

import (
	"Avalon/avalon/types"
	"Avalon/dbconn"

	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func userTablesCheck(engine *xorm.Engine) error {
	user := &types.UserInfo{}
	exist, err := engine.IsTableExist(user)
	if err != nil {
		panic(err)
	}
	if !exist {
		return engine.CreateTables(user)
	}

	return nil
}

func init() {
	var err error
	engine, err = dbconn.GetDBEngine()
	if err != nil {
		panic(err)
	}

	err = userTablesCheck(engine)
	if err != nil {
		panic(err)
	}
}

func GetUserById(id int) (*types.UserInfo, error) {
	user := &types.UserInfo{}
	_, err := engine.Id(id).Get(user)
	return user, err
}

func AddUser(user *types.UserInfo) error {
	_, err := engine.Insert(user)
	return err
}

func UpdateUser(user *types.UserInfo) error {
	_, err := engine.Update(user)
	return err
}

func DeleteUserById(id int) error {
	user := &types.UserInfo{}
	_, err := engine.Id(id).Delete(user)
	return err
}
