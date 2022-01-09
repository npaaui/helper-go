package db

type Conf struct {
	DriverName      string
	ConnMaxLifetime int64
	Prefix          string
	Conn
}
type Conn interface {
	GetDataSourceName() string
}

var ConfIns *Conf

func (d *Conf) InitDbConf() {
	ConfIns = d
}

/**
 * mysql
 */
type MysqlConf struct {
	Host     string
	Username string
	Password string
	Database string
}

func (c MysqlConf) GetDataSourceName() (dataSourceName string) {
	dataSourceName = c.Username + ":" + c.Password + "@(" + c.Host + ")/" + c.Database + "?charset=utf8&loc=Local"
	return
}
