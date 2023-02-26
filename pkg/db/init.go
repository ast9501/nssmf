package db

import (
	"database/sql"

	"github.com/ast9501/nssmf/pkg/logger"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db *sql.DB = nil
)

func Init(url, dbName, user, passwd string) {
	db, err := sql.Open("mysql", user+":"+passwd+"@tcp("+url+")/"+dbName)
	if err != nil {
		logger.DbLog.Errorln("Failed to open the connection to DB: ", err)
	} else if err = db.Ping(); err != nil {
		logger.DbLog.Errorln("Falied to connect DB, is DB alived?")
	}
	Db = db
	logger.DbLog.Infoln("Connection established: ", url)
	logger.DbLog.Infoln("Using database: ", dbName)
	createSnssaiLists()
	createPlmnIdLists()
	createPerReqLists()
	createSliceProfileLists()
}

func createSnssaiLists() {
	_, err := Db.Exec("CREATE TABLE IF NOT EXISTS SNSSAILISTS (SliceProfileId varchar(20), Sst varchar(8), Sd varchar(24))")
	if err != nil {
		logger.DbLog.Errorln("Failed to create SNSSAILISTS table: ", err)
	}
}

func createPlmnIdLists() {
	_, err := Db.Exec("CREATE TABLE IF NOT EXISTS PLMNIDLISTS (SliceProfileId varchar(20), Mnc varchar(3), Mcc varchar(3))")
	if err != nil {
		logger.DbLog.Errorln("Failed to create PLMNISLISTS table: ", err)
	}
}

func createPerReqLists() {
	_, err := Db.Exec("CREATE TABLE IF NOT EXISTS PERREQLISTS (SliceProfileId varchar(20), ExpDataRateDl integer, ExpDataRateUl integer, AreaTrafficCapDl integer, AreaTrafficCapUl integer, OverallUserDensity integer)")
	if err != nil {
		logger.DbLog.Errorln("Failed to create PERREQLISTS table: ", err)
	}
}

func createSliceProfileLists() {
	// Status: 1 = prepared, 0 = processing, -1 = unknown
	_, err := Db.Exec("CREATE TABLE IF NOT EXISTS SLICEPROFILELISTS (SliceProfileId varchar(20), Status integer)")
	if err != nil {
		logger.DbLog.Errorln("Failed to create SLICEPROFILELISTS table: ", err)
	}
}
