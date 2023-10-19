package main

import (
	"go-jwt-api/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)


func main(){
	router := gin.Default()

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Welcome Home")
	})

	server := &http.Server{
		Addr: ":8000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}