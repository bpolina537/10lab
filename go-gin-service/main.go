package main

import (
    "github.com/gin-gonic/gin"
    "go-gin-service/middleware"
)

type ProcessRequest struct {
    Data string `json:"data"`
}

type ProcessResponse struct {
    Status   string `json:"status"`
    Received string `json:"received"`
}

func main() {
    r := gin.Default()
    r.Use(middleware.Logger())

    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "pong"})
    })

    r.POST("/process", func(c *gin.Context) {
        var req ProcessRequest
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(400, gin.H{"error": "invalid request"})
            return
        }
        c.JSON(200, ProcessResponse{
            Status:   "ok",
            Received: req.Data,
        })
    })

    r.Run(":8080")
}