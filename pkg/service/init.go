package service

import (
	"github.com/ast9501/nssmf/internal/service/management"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type NSSMF struct {
}

// schemes http
func (nssmf *NSSMF) Start() {
	// TODO: external function call for customize gin engine loger
	router := gin.New()

	// Add service to router
	management.AddService(router)

	// api server by swagger in debug mode
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// TODO: Load server binding port from conifg
	router.RunTLS(":8000", "../config/TLS/nssmf.nycu.crt", "../config/TLS/nssmf.nycu.key")
}
