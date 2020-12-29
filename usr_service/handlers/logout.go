package handler

import "github.com/gin-gonic/gin"

func (u *UsrServiceImpl) Logout(ctx *gin.Context) {
	ctx.JSON(200, "logout")
}
