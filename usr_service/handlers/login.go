package handler

import "github.com/gin-gonic/gin"

func (u *UsrServiceImpl) Login(ctx *gin.Context) {
	ctx.JSON(200, "login")
}
