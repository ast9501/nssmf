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
	C.String(http.StatusOK, "Hello User!")
	logger.HandlerLog.Infoln("Receive Health Check Request From: ", C.ClientIP())
}

// @Summary		deallocate nssi
// @Tags		NSSI
// @Description	deallocate nssi
// @param		SliceProfileId path string true "SliceProfileId for deleation"
// @Success		200
// @Router		/ObjectManagement/NSS/SliceProfiles/{SliceProfileId} [delete]
func DeallocateNssi(C *gin.Context) {
	sliceProfileId := C.Param("SliceProfileId")
	C.String(http.StatusOK, "Deallocate Nssi, "+sliceProfileId)
	logger.HandlerLog.Infoln("Receive Health Check Request From: ", C.ClientIP())
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
