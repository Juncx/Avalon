package dbconn

import (
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var DBEngine *xorm.Engine

func NewEngine() (*xorm.Engine, error) {
	dbDriver := "mysql"
	dbInfo := "192.168.2.150:3306"
	dbUser := "root"
	dbPasswd := "daemon"
	dbName := "avalon"

	dbConn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", dbUser, dbPasswd, dbInfo, dbName)
	engine, err := xorm.NewEngine(dbDriver, dbConn)
	if err != nil {
		panic(err)
	}
	return engine, err
}

func GetDBEngine() (*xorm.Engine, error) {
	var err error
	if DBEngine == nil {
		err = errors.New("error db engine")
	}
	return DBEngine, err
}

func keepAlive(engine *xorm.Engine) error {
	return nil
}

func close(engine *xorm.Engine) error {
	return engine.Close()
}

func init() {
	var err error
	DBEngine, err = NewEngine()
	if err != nil {
		panic(err)
	}
	keepAlive(DBEngine)
}
