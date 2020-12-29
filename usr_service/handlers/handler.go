package handler

import "github.com/gin-gonic/gin"

type UsrService interface {
	Login(ctx *gin.Context)
	Logout(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type UsrServiceImpl struct {
}

func NewUsrServiceImpl() *UsrServiceImpl {
	return new(UsrServiceImpl)
}
