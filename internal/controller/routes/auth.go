package routes

import "github.com/gin-gonic/gin"

func AuthRoutes(r *gin.Engine, handlers ...gin.HandlerFunc) {
	v1 := r.Group("/auth")
	{
		v1.GET("profile", handlers...)
	}
}
