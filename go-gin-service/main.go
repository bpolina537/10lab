package main

import (
    "github.com/gin-gonic/gin"
    "go-gin-service/middleware"
)

func main() {
    r := gin.Default()

    // Подключаем middleware глобально
    r.Use(middleware.Logger())

    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

    r.Run(":8080")
}