package api

import (
	"github.com/gin-gonic/gin"
	"go-community/model"
)

type Login struct {
	Email    string `form:"email" json:"email" xml:"email"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func UserCreate(c *gin.Context) {
	var json Login

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(200, gin.H{"test": "fail"})
		return
	}

	var u = model.UserService{&model.User{}, &model.User{}}

	id, err := u.Create(c.Request.Context(), "tester2", json.Email, json.Password)
	if err != nil {
		c.JSON(200, gin.H{"test": "create fail"})
		return
	}

	c.JSON(200, gin.H{"id": id, "email": json.Email, "pw": json.Password})
}
