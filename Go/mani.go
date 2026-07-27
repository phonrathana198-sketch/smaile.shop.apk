package main

import (
    "net/http"
    "gin-gonic/Goo"
)

func main() {
    r := gin.Default()

    // Route ចម្បង
    r.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello Gin!",
        })
    })

    // Route DNS query simulation
    r.GET("/dns-query", func(c *gin.Context) {
        domain := c.Query("name")
        c.JSON(http.StatusOK, gin.H{
            "Status": 0,
            "Question": []gin.H{
                {"name": domain, "type": 1},
            },
            "Answer": []gin.H{
                {"name": domain, "type": 1, "TTL": 3599, "data": "93.184.216.34"},
            },
        })
    })

    // Run server on port 8080
    r.Run(":8080")
}
