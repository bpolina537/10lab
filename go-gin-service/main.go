package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/gorilla/websocket"
    "go-gin-service/middleware"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool { return true },
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)

type Message struct {
    Username string `json:"username"`
    Text     string `json:"text"`
}

type ProcessRequest struct {
    Data string `json:"data"`
}

type ProcessResponse struct {
    Status   string `json:"status"`
    Received string `json:"received"`
}

func handleWebSocket(c *gin.Context) {
    conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        return
    }
    defer conn.Close()
    clients[conn] = true

    for {
        var msg Message
        err := conn.ReadJSON(&msg)
        if err != nil {
            delete(clients, conn)
            break
        }
        broadcast <- msg
    }
}

func handleMessages() {
    for {
        msg := <-broadcast
        for client := range clients {
            err := client.WriteJSON(msg)
            if err != nil {
                client.Close()
                delete(clients, client)
            }
        }
    }
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

    r.GET("/ws", handleWebSocket)

    go handleMessages()
    r.Run(":8080")
}