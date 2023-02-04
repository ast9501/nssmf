package service

import (
	"github.com/ast9501/nssmf/internal/service/management"
	"github.com/gin-gonic/gin"
)

type NSSMF struct {
}

func (nssmf *NSSMF) Start() {
	// TODO: external function call for customize gin engine loger
	router := gin.New()

	//
	management.AddService(router)
	router.Run(":8000")
}
