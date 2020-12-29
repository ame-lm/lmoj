package handler

import "github.com/gin-gonic/gin"

func (u *UsrServiceImpl) Register(ctx *gin.Context) {
	ctx.JSON(200, "register")
}
