package main

import (
	"fmt"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/service"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("mysql init failed with %+v", err))
	}

	r := gin.Default()
	r.POST("/api/message", service.HandleInteraction)
	r.Run(":80")
}
