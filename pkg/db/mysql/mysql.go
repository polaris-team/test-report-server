package mysql

import (
	"github.com/polaris-team/test-report-server/pkg/config"
	"sync"

	"upper.io/db.v3/lib/sqlbuilder"
	upper "upper.io/db.v3/mysql"

	"errors"
	"strconv"
)

var mysqlMutex sync.Mutex
var settings *upper.ConnectionURL

func GetConnect() (sqlbuilder.Database, error) {
	if config.GetMysqlConfig() == nil {
		panic(errors.New("Mysql Datasource Configuration is missing!"))
	}

	if settings == nil {
		mysqlMutex.Lock()
		defer mysqlMutex.Unlock()
		if settings == nil {
			mc := config.GetMysqlConfig()
			settings = &upper.ConnectionURL{
				User:     mc.Usr,
				Password: mc.Pwd,
				Database: mc.Database,
				Host:     mc.Host + ":" + strconv.Itoa(mc.Port),
				Socket:   "",
				Options:  nil,
			}
		}
	}

	sess, err := upper.Open(settings)
	if err != nil {
		return nil, err
	}
	return sess, nil
}
