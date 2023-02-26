package nss

import (
	"net/http"

	"github.com/ast9501/nssmf/pkg/logger"
	"github.com/gin-gonic/gin"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method string
	// Pattern is the pattern of the URI.
	Pattern string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

type Routes []Route

// Register service to router
func AddService(engine *gin.Engine) *gin.RouterGroup {
	group := engine.Group("/ObjectManagement/NSS")

	for _, route := range routes {
		switch route.Method {
		case "GET":
			group.GET(route.Pattern, route.HandlerFunc)
		case "POST":
			group.POST(route.Pattern, route.HandlerFunc)
		case "PUT":
			group.PUT(route.Pattern, route.HandlerFunc)
		case "DELETE":
			group.DELETE(route.Pattern, route.HandlerFunc)
		case "PATCH":
			group.PATCH(route.Pattern, route.HandlerFunc)
		}
	}

	return group
}

// @Summary		allocate nssi
// @Tags		NSSI
// @Description	allocate nssi
// @Success		201
// @param		SliceProfileId body NetworkSliceSubnetProfileReq true "SliceProfile"
// @Router		/ObjectManagement/NSS/SliceProfiles [post]
func AllocateNssi(C *gin.Context) {
	var requestBody NetworkSliceSubnetProfileReq
	if err := C.BindJSON(&requestBody); err != nil {
		logger.HandlerLog.Errorln("Failed to bind AllocateNssi.Request to &SliceProfile: ", err)
	}
	id, err := CreateNssProfile(requestBody.SNssaiList, requestBody.PlmnIdList, requestBody.PerfReq)
	if err != nil {
		C.JSON(http.StatusInternalServerError, CreateNssProfileRes{NssProfileId: id})
	} else {
		C.JSON(http.StatusCreated, CreateNssProfileRes{NssProfileId: id})
	}
}

// @Summary		deallocate nssi
// @Tags		NSSI
// @Description	deallocate nssi
// @param		SliceProfileId path string true "SliceProfileId for deleation"
// @Success		200
// @Router		/ObjectManagement/NSS/SliceProfiles/{SliceProfileId} [delete]
func DeallocateNssi(C *gin.Context) {
	sliceProfileId := C.Param("SliceProfileId")
	err := DeleteNssProfile(sliceProfileId)
	if err != nil {
		C.Status(http.StatusInternalServerError)
	} else {
		C.Status(http.StatusOK)
	}
}

var routes = Routes{
	{
		"AllocateNssi",
		"POST",
		"/SliceProfiles",
		AllocateNssi,
	},
	{
		"DeallocateNssi",
		"DELETE",
		"/SliceProfiles/:SliceProfileId",
		DeallocateNssi,
	},
}
