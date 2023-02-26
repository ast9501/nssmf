package db

import "github.com/ast9501/nssmf/pkg/logger"

// AssociateNsTb_<SliceProfileId>, mapping associate NetworkService to Slice Profile
func CreateAssocaiteNsTable(profileId string) (err error) {
	_, err = Db.Exec("CREATE TABLE AssociateNsTb_" + profileId + " (NsId varchar(20))")

	if err != nil {
		logger.AppLog.Errorln("Failed to create AssociateNsTable: ", err)
	}
	return
}

func DropAssociateNsTable(profileId string) (err error) {
	_, err = Db.Exec("DROP TABLE AssociateNsTb_" + profileId)

	return
}

// NsTable_<NsId>, keep vnf list under the NetworkService
func InitNsTable(nsName string) (err error) {

	return nil
}

func InsertNsTable(nsName string) (err error) {

	return nil
}
