package nss

// swagger:model	NetworkSliceSubnetProfileReq
type NetworkSliceSubnetProfileReq struct {
	SliceProfileId string
	SNssaiList     SNssai
	PlmnIdList     PlmnId
	PerfReq        PerReq
}

type SNssai struct {
	SSt string `json: "sst"` // 8 bit
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
