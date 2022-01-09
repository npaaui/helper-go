package db

import (
	"errors"
	"fmt"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

var engineIns *xorm.Engine

var SetDbEngineOnce sync.Once

func GetDbEngineIns() *xorm.Engine {
	SetDbEngineOnce.Do(SetDbEngine)
	return engineIns
}

func SetDbEngine() {
	if ConfIns == nil {
		panic(errors.New("[danger] DbConfIns is nil"))
	}

	DbEngine, err := xorm.NewEngine(ConfIns.DriverName, ConfIns.Conn.GetDataSourceName())
	if err != nil {
		panic(fmt.Errorf("[danger] NewEngine error: %w", err))
	}

	if ConfIns.ConnMaxLifetime > 0 {
		DbEngine.DB().SetConnMaxLifetime(time.Duration(ConfIns.ConnMaxLifetime) * time.Second)
	}

	if ConfIns.Prefix != "" {
		tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, ConfIns.Prefix)
		DbEngine.SetTableMapper(tbMapper)
	}

	engineIns = DbEngine
}
