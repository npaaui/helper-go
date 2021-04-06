package {{.PackageName}}
{{$exportModelName := .ModelName | FormatCamelcase}}
import (
    . "business/common"
)

/**{{range .TableSchema}}
 * {{.Field}} {{.Type | TypeConvert}} {{.Comment}} {{end}}
 */

type {{$exportModelName}} struct {
{{range .TableSchema}}    {{.Field | ExportColumn | FormatCamelcase}} {{.Type | TypeConvert}} {{.Field | Tags}}
{{end}}}

func New{{$exportModelName}}Model() *{{$exportModelName}} {
	return &{{$exportModelName}}{}
}

func (m *{{$exportModelName}}) Info() bool {
	has, err := DbEngine.Get(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return has
}

func (m *{{$exportModelName}}) Insert() int64 {
	row, err := DbEngine.Insert(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *{{$exportModelName}}) Update(arg *{{$exportModelName}}) int64 {
	row, err := DbEngine.Update(arg, m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *{{$exportModelName}}) Delete() int64 {
	row, err := DbEngine.Delete(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

{{range .TableSchema}}func (m *{{$exportModelName}}) Set{{.Field | FormatCamelcase}}(arg {{.Type | TypeConvert}}) *{{$exportModelName}} {
	m.{{.Field | FormatCamelcase}} = arg
	return m
}
{{end}}