package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Run func registers endpoints and runs the API
func Run(address string, port int) {
	if !viper.GetBool("debug") {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	r.GET("/version", getVersion)

	v1 := r.Group("/v1")
	{
		// Registry events notification endpoint
		v1.GET("/events", getEvents)
		v1.POST("/events", postEvents)

		// Namespace endpoints
		v1.GET("/namespaces", getNamespaces)
		v1.GET("/namespace/:id", getNamespace)
		v1.PUT("/namespace", putNamespace)
		v1.DELETE("/namespace/:id", deleteNamespace)
		v1.PATCH("/namespace/:id", patchNamespace)
		v1.GET("/namespace/:id/images", getNamespaceImages)

		// Image endpoints
		v1.GET("/images", getImages)
		v1.GET("/image/:id", getImage)
		v1.GET("/image/:id/tags", getImageTags)

		// User endpoints
		v1.GET("/users", getUsers)
		v1.GET("/user", getUser)
		v1.PUT("/user", putUser)
	}

	err := r.Run(fmt.Sprintf("%s:%d", address, port))
	if err != nil {
		panic(err)
	}
}
