package controller

import (
	"authenName/model"
	repo "authenName/repository"
	"github.com/gin-gonic/gin"
	"time"
)

func CreateAuthen(c *gin.Context) {
	authen := model.Authen{}
	_ = c.BindJSON(&authen)
	authen.CreateAt = time.Now()
	authen.CreateEnd = time.Now()

	res := repo.CreateAuthen(authen)
	if res != nil {
		c.JSON(200, gin.H{
			"code":        405,
			"status":      "error",
			"message":     "Cannot create authen",
			"description": "The authen you create is now Error. " + res.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":        200,
		"status":      "success",
		"message":     "Create authen Success",
		"description": "The authen you create is Success.",
	})
}

func GetAuthenAllForSubject(c *gin.Context){
	id := c.Param("id")
	authen, err := repo.FindAuthenBySubjectId(id)
	if err != nil {
		c.JSON(200, gin.H{
			"code":        405,
			"status":      "error",
			"message":     "Not found authen",
			"description": "The authen is not find " + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":        200,
		"status":      "success",
		"message":     "Fine success",
		"description": "The authen you fine is Success.",
		"authen":     authen,
	})
	return
}
