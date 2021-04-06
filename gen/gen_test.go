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
		Prefix:          "b_",
		Conn: db.MysqlConf{
			Host:     "127.0.0.1",
			Username: "alice",
			Password: "npaauI2396",
			Database: "business",
		},
	}).InitDbEngine()

	(&Conf{
		ModelFolder: "model/",
		TplFile:     "model.tpl",
		TableNames:  "user",
		DbName:      "business",
	}).InitGenConf()
	GenerateModelFile()
}
