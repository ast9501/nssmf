package service

import (
	"github.com/ast9501/nssmf/internal/service/management"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type NSSMF struct {
}

//	@title			O-RAN NSSMF api doc
//	@version		1.0
//	@description	NSSMF api doc

//	@contact.name	ast9501
//	@contact.email	ast9501.cs10@nycu.edu.tw

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host			192.168.0.166:8000
//
// schemes http
func (nssmf *NSSMF) Start() {
	// TODO: external function call for customize gin engine loger
	router := gin.New()

	// Add service to router
	management.AddService(router)

	// api server by swagger in debug mode
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// TODO: Load server binding port from conifg
	router.Run(":8000")
}
