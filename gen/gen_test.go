package gen

import (
	"fmt"
	"testing"

	"github.com/npaaui/helper-go/db"
)

func TestGenModelFile(t *testing.T) {
	defer func() {
		fmt.Println(recover())
	}()

	(&db.Conf{
		DriverName:      "mysql",
		ConnMaxLifetime: 86400,
		Prefix:          "m_",
		Conn: db.MysqlConf{
			Host:     "127.0.0.1",
			Username: "Username",
			Password: "Password",
			Database: "mall_goods",
		},
	}).InitDbConf()

	(&Conf{
		ModelFolder: "./model/",
		TplFile:     "model.tpl",
		TableNames:  "goods",
		DbName:      "mall_goods",
	}).InitGenConf()
	GenerateModelFile()
}
