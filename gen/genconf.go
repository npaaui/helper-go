package gen

type Conf struct {
	ModelFolder string
	TplFile     string
	TableNames  string
	DbName      string
}

var ConfIns *Conf

func (g *Conf) InitGenConf() {
	ConfIns = g
}
