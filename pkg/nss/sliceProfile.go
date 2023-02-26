package nss

import (
	"math/rand"
	"time"

	"github.com/ast9501/nssmf/pkg/logger"

	"github.com/ast9501/nssmf/pkg/db"
	"github.com/ast9501/nssmf/pkg/nservice"
)

// swagger:model	NetworkSliceSubnetProfileReq
type NetworkSliceSubnetProfileReq struct {
	SNssaiList []SNssai
	PlmnIdList []PlmnId
	PerfReq    []PerReq
}

type NetworkSliceSubnetProfile struct {
	SliceProfileId string
	SNssaiList     []SNssai
	PlmnIdList     []PlmnId
	PerfReq        []PerReq
}

type SNssai struct {
	Sst string `json: "sst"` // 8 bit
	Sd  string `json: "sd"`  // 24 bit
}

type PlmnId struct {
	Mnc string `json: "mnc"` // 3 digit
	Mcc string `json: "mcc"` // 3 digit
}

// eMBB PerfReq
type PerReq struct {
	ExpDataRateDl      int
	ExpDataRateUl      int
	AreaTrafficCapDl   int
	AreaTrafficCapUl   int
	OverallUserDensity int
}

var letterRunes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// TODO: check if exist network service fit the requirements
func NsProfileLookup() {
	nservice.GetAllNetworkService()
}

// TODO: assocaite network service to network slice profile
func NsProfileAssociate() {

}

// TODO: use type assertions to fit different type (SNssai, PlmnId, ...)
// https://go.dev/ref/spec#Type_assertions
func CreateNssProfile(snssaiList []SNssai, plmnList []PlmnId, perList []PerReq) (profileId string, err error) {
	profileId = GenId(8)
	err = db.InsertNssProfileTable(profileId)
	if err != nil {
		profileId = "Failed to create NssProfileTable, please check system logs"
		return
	}
	// TODO: create related table
	err = createSnssaiList(profileId, snssaiList)
	if err != nil {
		profileId = "Failed to create SnssaiListTable, please check system logs"
		return
	}

	err = createPlmnList(profileId, plmnList)
	if err != nil {
		profileId = "Failed to create PlmnListTable, please check system logs"
		return
	}

	err = createPerreqList(profileId, perList)
	if err != nil {
		profileId = "Failed to create PerreqListTable, please check system logs"
		return
	}

	err = db.CreateAssocaiteNsTable(profileId)
	if err != nil {
		profileId = "Failed to create assocaiteNsTable for NSS profile, please check system logs"
		return
	}

	return
}

func createSnssaiList(profileId string, s []SNssai) (err error) {
	for _, e := range s {
		err = db.InsertNssaiTable(profileId, e.Sst, e.Sd)
		if err != nil {
			return
		}
	}
	return nil
}

func createPlmnList(profileId string, p []PlmnId) (err error) {
	for _, e := range p {
		err = db.InsertPlmnTable(profileId, e.Mnc, e.Mcc)
		if err != nil {
			return
		}
	}
	return nil
}

func createPerreqList(profileId string, p []PerReq) (err error) {
	for _, e := range p {
		err = db.InsertPerreqTable(profileId, e.AreaTrafficCapDl, e.AreaTrafficCapUl, e.ExpDataRateDl, e.ExpDataRateUl, e.OverallUserDensity)
		if err != nil {
			return
		}
	}
	return nil
}

func DeleteNssProfile(profileId string) (err error) {
	err = db.DropNssProfile(profileId)
	if err != nil {
		logger.AppLog.Errorln("Failed to Delete NssProfile: ")
		return
	}

	err = db.DropAssociateNsTable(profileId)
	if err != nil {
		logger.AppLog.Errorln("Failed to Delete NssProfile: ")
		return
	}
	err = db.CleanAssociateTable(profileId)
	if err != nil {
		logger.AppLog.Errorln("Failed to Delete NssProfile: ")
		return
	}
	return
}

func GenId(length int) (id string) {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	id = string(b)
	return
}
