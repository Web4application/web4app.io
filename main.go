package main

import (
    "my_project/handlers"
    "my_project/services"

    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // WebSocket Endpoint
    r.GET("/ws", func(c *gin.Context) {
        handlers.WebSocketHandler(c.Writer, c.Request)
    })

    // Logging Middleware
    r.Use(func(c *gin.Context) {
        handlers.LoggingMiddleware(c.Writer, c.Request)
        c.Next()
    })

    // Fraud Detection Endpoint
    r.POST("/fraud", func(c *gin.Context) {
        type Request struct {
            Amount float64 `json:"amount"`
            IP     string  `json:"ip"`
        }
        var req Request
        if err := c.BindJSON(&req); err != nil {
            c.JSON(400, gin.H{"error": err.Error()})
            return
        }
        isFraud := services.IsFraudulentTransaction(req.Amount, req.IP)
        c.JSON(200, gin.H{"isFraud": isFraud})
    })

    // Recommendations Endpoint
    r.GET("/recommendations/:userID", func(c *gin.Context) {
        userID := c.Param("userID")
        recommendations := services.GetRecommendations(userID)
        c.JSON(200, gin.H{"recommendations": recommendations})
    })

    // Sentiment Analysis Endpoint
    r.POST("/sentiment", func(c *gin.Context) {
        type Request struct {
            Text string `json:"text"`
        }
        var req Request
        if err := c.BindJSON(&req); err != nil {
            c.JSON(400, gin.H{"error": err.Error()})
            return
        }
        sentiment := services.AnalyzeSentiment(req.Text)
        c.JSON(200, gin.H{"sentiment": sentiment})
    })

    r.Run(":5432") // Run on port 3000
}
