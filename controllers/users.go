package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"jwttest/cache"
	"jwttest/models"
	"net/http"
)

func Allusers(c *gin.Context)  {
	users :=[]models.User{}
	models.DB.Find(&users)
	c.JSON(http.StatusOK,users)
}

func GetUser(c *gin.Context)  {
	id := c.Param("id")
	user :=cache.GetItem("user."+id)
	if(user!=""){
		fmt.Println("read from redis!!!")
		c.JSON(http.StatusOK,user)
		return
	}
	u :=models.User{}
	result := models.DB.Find(&u ,id)
	if result.RowsAffected==0 {
		c.JSON(http.StatusNotFound,"not found")
		return
	}
	cache.SetItem("user."+id,models.ToString(u),0)
	fmt.Println("read from db!!!")
	c.JSON(http.StatusOK,u)
}

func DeleteUser(c *gin.Context)  {
	id := c.Param("id")
	result := models.DB.Delete(&models.User{} ,id)
	if result.RowsAffected==0 {
		c.JSON(http.StatusNotFound,"not found")
		return
	}
	c.JSON(http.StatusOK,"success delete")
}