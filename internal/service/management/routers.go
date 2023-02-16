package management

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
	group := engine.Group("/nssmf/v1")

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

// @Summary		root path
// @Description	health check for api server
// @Success		200
// @Router		/nssmf/v1/ [get]
func Index(C *gin.Context) {
	C.String(http.StatusOK, "Hello User!")
	logger.HandlerLog.Infoln("Receive Health Check Request From: ", C.ClientIP())
}

var routes = Routes{
	{
		"Index",
		"GET",
		"/",
		Index,
	},
}
