package main

import (
	"github.com/gin-gonic/gin"
	"jwttest/cache"
	"jwttest/controllers"
	"jwttest/middlewares"
	"jwttest/models"
	_ "net/http"
)

func main() {
	models.ConnectDataBase()
	cache.ConnectToRedis()
	r := gin.Default()
	
	public := r.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login",controllers.Login)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user",controllers.CurrentUser)
	public.GET("/users",controllers.Allusers)
	public.GET("/users/:id",controllers.GetUser)
	public.DELETE("/users/:id",controllers.DeleteUser)
	r.Run(":8000")

}