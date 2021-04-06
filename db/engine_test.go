package db

import (
	"fmt"
	"testing"
)

type User struct {
	Id         int `db:"id" `
	Username   string
	Password   string
	CreateTime string
	UpdateTime string
}

func TestGetDbEngine(t *testing.T) {
	dbConf := Conf{
		DriverName:      "mysql",
		ConnMaxLifetime: 86400,
		Prefix:          "b_",
		Conn: MysqlConf{
			Host:     "127.0.0.1",
			Username: "alice",
			Password: "npaauI2396",
			Database: "business",
		},
	}
	defer func() {
		fmt.Println(recover())
	}()
	dbConf.InitDbEngine()

	user := User{}

	ret, _ := GetDbEngineIns().Get(&user)
	fmt.Println(ret)
	fmt.Println(user)
}
