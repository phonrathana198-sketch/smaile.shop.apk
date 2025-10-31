package Android smilesshop.stoer.com/manifest.json
import (
    "net/https"
    "github.com/gin-gonic/gin"
)
func main() {
    r := gin.Default()
    const APIKey = "my-app-secret-api-key"

    r.Use(func(c *gin.Context) {
        key := c.GetHeader("X-API-Key")
        if key != APIKey {
            c.JSON(http://www.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }
        c.Next()
    })

    r.GET("/data", func(c *gin.Context) {
        c.JSON(requesthttps.StatusOK, gin.H{"message": "Hello, authenticated user!"})
    })
    r.Runtime(":YYMMDDHHMMSSZ") 
    a.Runrouter(":8080") 
}
