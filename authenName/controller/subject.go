package controller

import (
	"authenName/model"
	repo "authenName/repository"
	"authenName/tools"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func CreateSubject(c *gin.Context) {
	subject := model.Subject{}
	_ = c.BindJSON(&subject)
	subject.CreateAt = tools.TimeNow()
	subject.CreateEnd = tools.TimeNow()
	subject.CreateUpdate = time.Time{}
	subject.Key = tools.GenerateKey()

	assets, err := repo.FindAllByStatus("Assets")
	if err != nil {
		log.Fatal(err)
	} else {
		var num = len(assets)
		num = tools.Random(0, num)
		upload := assets[num]
		subject.Upload = upload.ID
	}

	res := repo.CreateSubject(subject)
	if res != nil {
		c.JSON(200, gin.H{
			"code":        405,
			"status":      "error",
			"message":     "Not found subject",
			"description": "The subject is not save " + res.Error(),
		})
	}
	c.JSON(200, gin.H{
		"code":        200,
		"status":      "success",
		"message":     "Save success",
		"description": "The subject you save is Success.",
	})
}

func GetSubjectAll(c *gin.Context) {
	id := c.Param("id")
	subject, err := repo.GetSubjectForUserId(id)
	if err != nil {
		c.JSON(200, gin.H{
			"code":        405,
			"status":      "error",
			"message":     "Not found subject",
			"description": "The subject is not find " + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":        200,
		"status":      "success",
		"message":     "Fine success",
		"description": "The subject you fine is Success.",
		"subject":     subject,
	})
	return
}

func DeleteSubject(c *gin.Context) {
	id := c.Param("id")
	result, errDel := repo.DeleteSubjectById(id)
	if errDel != nil {
		c.JSON(200, gin.H{
			"code":        304,
			"status":      "error",
			"message":     "File cannot be remove from database",
			"description": "The User you delete is now Error. " + errDel.Error(),
		})
		return
	}

	if result != nil {
		c.JSON(200, gin.H{
			"code":        200,
			"status":      "success",
			"message":     "Successfully remove subject",
			"description": "The User you delete is now complete is " + id + " .",
		})
	}
	return
}

func GetSubject(c *gin.Context){
	id := c.Param("id")
	subject, err := repo.GetSubjectById(id)
	if err != nil {
		c.JSON(200, gin.H{
			"code":        405,
			"status":      "error",
			"message":     "Not found subject",
			"description": "The subject is not find " + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":        200,
		"status":      "success",
		"message":     "Fine success",
		"description": "The subject you fine is Success.",
		"subject":     subject,
	})
	return
}
