package db

import (
	"strconv"

	"github.com/ast9501/nssmf/pkg/logger"
)

func InsertNssProfileTable(profileId string) (err error) {
	_, err = Db.Exec("INSERT INTO SLICEPROFILELISTS (SliceProfileId) VALUES ('" + profileId + "')")

	if err != nil {
		logger.AppLog.Errorln("Failed to create SliceProfile: ", err)
	}
	return
}

func InsertNssaiTable(profileId, sst, sd string) (err error) {
	_, err = Db.Exec("INSERT INTO SNSSAILISTS (SliceProfileId, Sst, Sd) VALUES ('" + profileId + "', '" + sst + "', '" + sd + "')")
	if err != nil {
		logger.AppLog.Errorln("Failed to create SNssaiTable: ", err)
	}
	return
}

func DeleteNssaiTable(profileId string) (err error) {
	_, err = Db.Exec("DELETE FROM SNSSAILISTS WHERE SliceProfileId='" + profileId + "'")
	if err != nil {
		logger.AppLog.Errorln("Failed to delete SNssaiTable (profileId: ", profileId+"): ", err)
	}
	return
}

func InsertPlmnTable(profileId, mnc, mcc string) (err error) {
	_, err = Db.Exec("INSERT INTO PLMNIDLISTS (SliceProfileId, Mnc, Mcc) VALUES ('" + profileId + "', '" + mnc + "', '" + mcc + "')")
	if err != nil {
		logger.AppLog.Errorln("Failed to create PlmnTable: ", err)
	}
	return
}

func DeletePlmnTable(profileId string) (err error) {
	_, err = Db.Exec("DELETE FROM PLMNIDLISTS WHERE SliceProfileId='" + profileId + "'")
	if err != nil {
		logger.AppLog.Errorln("Failed to delete PlmnTable (profileId: ", profileId+"): ", err)
	}
	return
}

func DeletePerreqTable(profileId string) (err error) {
	_, err = Db.Exec("DELETE FROM PERREQLISTS WHERE SliceProfileId='" + profileId + "'")
	if err != nil {
		logger.AppLog.Errorln("Failed to delete PerReqTable (profileId: ", profileId+"): ", err)
	}
	return
}

func InsertPerreqTable(profileId string, atCapDl, atCapUl, edrDl, edrUl, oUserDsty int) (err error) {
	_, err = Db.Exec("INSERT INTO PERREQLISTS (SliceProfileId, ExpDataRateDl, ExpDataRateUl, AreaTrafficCapDl, AreaTrafficCapUl, OverallUserDensity) VALUES ( '" + profileId + "', " + strconv.Itoa(edrDl) + ", " + strconv.Itoa(edrUl) + ", " + strconv.Itoa(atCapDl) + ", " + strconv.Itoa(atCapUl) + ", " + strconv.Itoa(oUserDsty) + ")")
	if err != nil {
		logger.AppLog.Errorln("Failed to create PerReqTable: ", err)
		logger.AppLog.Errorln("Exec: " + "INSERT INTO PERREQLISTS (SliceProfileId, ExpDataRateDl, ExpDataRateUl, AreaTrafficCapDl, AreaTrafficCapUl, OverallUserDensity) VALUES ( '" + profileId + "', " + strconv.Itoa(edrDl) + ", " + strconv.Itoa(edrUl) + ", " + strconv.Itoa(atCapDl) + ", " + strconv.Itoa(atCapUl) + ", " + strconv.Itoa(oUserDsty) + ")")
	}
	return
}

func DropNssProfile(profileId string) (err error) {
	_, err = Db.Exec("DELETE FROM SLICEPROFILELISTS WHERE SliceProfileId='" + profileId + "'")
	if err != nil {
		logger.AppLog.Errorln("Failed to drop NssProfile table: ", err)
	}
	return
}

func CleanAssociateTable(profileId string) (err error) {
	err = DeleteNssaiTable(profileId)
	if err != nil {
		return
	}

	err = DeletePerreqTable(profileId)
	if err != nil {
		return
	}

	err = DeletePlmnTable(profileId)
	return
}

func GetNssProfile(profileId string) (err error, status int) {
	rows, err := Db.Query("SELECT Status FROM SLICEPROFILELISTS WHERE SliceProfileId = '" + profileId + "'")
	if err != nil {
		logger.AppLog.Errorln("Failed to filter SliceProfileId from sliceProfileList table: ", err)
		status = 503
		return
	}

	for rows.Next() {
		var s int
		err = rows.Scan(&s)
		if err != nil {
			logger.AppLog.Errorln("Failed at parsing SliceProfileId query: ", err)
		}
		status = s
		return
	}
	status = -2
	return
}

func GetAllNssProfile() {

}

func DeleteNssProfile() {

}
