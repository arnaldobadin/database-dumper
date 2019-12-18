package storage

type Connector interface {
	Open() (error)
	Close()	(error)
	Query(query string, operator func([]string) (), args ...interface{}) (error)
}

const MYSQL_TYPE = "mysql"
// const MSSQL_TYPE = "mssql"

func NewConnector(config *Connection) Connector {
	switch config.Strg {
		case MYSQL_TYPE:
			return &mysql{Openned : false, Config : config}
		// case MSSQL_TYPE:
			// return &mssql{Openned : false, Config : config}
		default:
			panic("Wrong connection type")
	}
}