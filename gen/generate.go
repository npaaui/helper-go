package gen

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/npaaui/helper-go/db"
)

func genModelFile(render *template.Template, dbName, tableName string) {
	tableSchema := make([]TableSchema, 0)
	err := db.EngineIns.SQL(
		"show full columns from " + tableName + " from " + dbName).Find(&tableSchema)

	if err != nil {
		fmt.Println(err)
		return
	}
	if db.ConfIns.Prefix != "" {
		tableName = tableName[len(db.ConfIns.Prefix):]
	}
	fileName := ConfIns.ModelFolder + strings.ToLower(tableName) + "Model.go"
	_ = os.Remove(fileName)
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	model := &ModelInfo{
		PackageName:     "model",
		BDName:          dbName,
		TablePrefixName: db.ConfIns.Prefix + tableName,
		TableName:       tableName,
		ModelName:       tableName,
		TableSchema:     &tableSchema,
	}
	if err := render.Execute(f, model); err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileName)
	cmd := exec.Command("imports", "-w", fileName)
	_ = cmd.Run()
}

func GenerateModelFile() {
	logDir, _ := filepath.Abs(ConfIns.ModelFolder)
	if _, err := os.Stat(logDir); err != nil {
		_ = os.Mkdir(logDir, os.ModePerm)
	}
	fmt.Println(ConfIns.ModelFolder)

	data, err := ioutil.ReadFile(ConfIns.TplFile)
	if nil != err {
		fmt.Printf("%v\n", err)
		return
	}

	render := template.Must(template.New("model").
		Funcs(template.FuncMap{
			"FirstCharUpper":       FirstCharUpper,
			"TypeConvert":          TypeConvert,
			"Tags":                 Tags,
			"ExportColumn":         ExportColumn,
			"Join":                 Join,
			"MakeQuestionMarkList": MakeQuestionMarkList,
			"ColumnAndType":        ColumnAndType,
			"ColumnWithPostfix":    ColumnWithPostfix,
			"FormatCamelcase":      FormatCamelcase,
		}).Parse(string(data)))

	tableName := ConfIns.TableNames
	if tableName != "" {
		tableNameSlice := strings.Split(tableName, ",")
		for _, v := range tableNameSlice {
			if db.ConfIns.Prefix != "" {
				v = db.ConfIns.Prefix + v
			}
			genModelFile(render, ConfIns.DbName, v)
		}
	} else {
		getTablesNameSql := "show tables from " + ConfIns.DbName
		tablaNames, err := db.EngineIns.QueryString(getTablesNameSql)
		if err != nil {
			fmt.Println(err)
		}
		for _, table := range tablaNames {
			tableCol := "Tables_in_" + ConfIns.DbName
			tablePrefixName := table[tableCol]
			genModelFile(render, ConfIns.DbName, tablePrefixName)
		}
	}
}
