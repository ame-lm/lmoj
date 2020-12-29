package main

import (
	handler "github.com/ame-lm/lmoj/usr_service/handlers"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	var h handler.UsrService = handler.NewUsrServiceImpl()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, "pong")
	})
	r.POST("/login", h.Login)
	r.POST("/logout", h.Logout)
	r.POST("/register", h.Register)
	log.Fatalln(r.Run())
}
