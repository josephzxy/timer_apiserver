package mysql

import (
	"strconv"
	"strings"
)

type MySQLErr int

const (
	DUPLICATE_ENTRY = 1062
	UNKNOWN         = -1
)

func GetMySQLErr(err error) MySQLErr {
	errMsg := err.Error()
	code, e := strconv.Atoi(errMsg[6:strings.Index(errMsg, ":")])
	if e != nil {
		return UNKNOWN
	}
	return MySQLErr(code)
}
