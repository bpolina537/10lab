package main

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
)

func TestPingEndpoint(t *testing.T) {
    router := gin.Default()
    router.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "pong"})
    })

    req, _ := http.NewRequest("GET", "/ping", nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    if w.Code != 200 {
        t.Errorf("Expected 200, got %d", w.Code)
    }

    var response map[string]string
    json.Unmarshal(w.Body.Bytes(), &response)
    if response["message"] != "pong" {
        t.Errorf("Expected pong, got %s", response["message"])
    }
}

func TestProcessEndpoint(t *testing.T) {
    router := gin.Default()
    router.POST("/process", func(c *gin.Context) {
        var req ProcessRequest
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(400, gin.H{"error": "invalid request"})
            return
        }
        c.JSON(200, ProcessResponse{Status: "ok", Received: req.Data})
    })

    body := `{"data":"hello"}`
    req, _ := http.NewRequest("POST", "/process", bytes.NewBufferString(body))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    if w.Code != 200 {
        t.Errorf("Expected 200, got %d", w.Code)
    }

    var response ProcessResponse
    json.Unmarshal(w.Body.Bytes(), &response)
    if response.Status != "ok" || response.Received != "hello" {
        t.Errorf("Expected ok/hello, got %s/%s", response.Status, response.Received)
    }
}

func TestProcessEndpointInvalidJSON(t *testing.T) {
    router := gin.Default()
    router.POST("/process", func(c *gin.Context) {
        var req ProcessRequest
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(400, gin.H{"error": "invalid request"})
            return
        }
        c.JSON(200, ProcessResponse{Status: "ok", Received: req.Data})
    })

    body := `invalid json`
    req, _ := http.NewRequest("POST", "/process", bytes.NewBufferString(body))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    if w.Code != 400 {
        t.Errorf("Expected 400, got %d", w.Code)
    }
}

func TestProcessEndpointEmptyData(t *testing.T) {
    router := gin.Default()
    router.POST("/process", func(c *gin.Context) {
        var req ProcessRequest
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(400, gin.H{"error": "invalid request"})
            return
        }
        c.JSON(200, ProcessResponse{Status: "ok", Received: req.Data})
    })

    body := `{"data":""}`
    req, _ := http.NewRequest("POST", "/process", bytes.NewBufferString(body))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    if w.Code != 200 {
        t.Errorf("Expected 200, got %d", w.Code)
    }
}

func TestProcessEndpointLargeData(t *testing.T) {
    router := gin.Default()
    router.POST("/process", func(c *gin.Context) {
        var req ProcessRequest
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(400, gin.H{"error": "invalid request"})
            return
        }
        c.JSON(200, ProcessResponse{Status: "ok", Received: req.Data})
    })

    // Создаём строку из читаемых символов длиной 10000
    largeStr := ""
    for i := 0; i < 10000; i++ {
        largeStr += "a"
    }
    body := `{"data":"` + largeStr + `"}`
    req, _ := http.NewRequest("POST", "/process", bytes.NewBufferString(body))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    if w.Code != 200 {
        t.Errorf("Expected 200, got %d", w.Code)
    }
}


func TestMiddlewareLogging(t *testing.T) {
    router := gin.New()
    router.Use(loggingMiddleware())
    router.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "pong"})
    })

    req, _ := http.NewRequest("GET", "/ping", nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    if w.Code != 200 {
        t.Errorf("Expected 200, got %d", w.Code)
    }
}