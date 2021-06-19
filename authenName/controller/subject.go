package controller

import (
	"authenName/model"
	"authenName/pojo"
	repo "authenName/repository"
	"authenName/tools"
	"strconv"

	"github.com/gin-gonic/gin"

	// "log"
	"fmt"
	"time"
)

func CreateSubject(c *gin.Context) {
	subject := model.Subject{}
	_ = c.BindJSON(&subject)
	subject.CreateAt = time.Now()
	subject.CreateEnd = time.Now()
	subject.CreateUpdate = time.Time{}
	subject.Key = tools.GenerateKey()

	assets, _ := repo.FindAllByStatus("Assets")
	if len(assets) != 0 {
		fmt.Println("Random Picture")
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

func CreateJoinSubject(c *gin.Context) {
	subject := model.Subject{}
	_ = c.BindJSON(&subject)
	subject.CreateAt = time.Now()
	subject.CreateEnd = time.Now()
	subject.CreateUpdate = time.Time{}

	assets, _ := repo.FindAllByStatus("Assets")
	if len(assets) != 0 {
		fmt.Println("Random Picture")
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

func GetSubject(c *gin.Context) {
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

func GetSubjectBySubjectKey(c *gin.Context) {
	subKey := c.Param("id")
	n, err := strconv.ParseInt(subKey, 10, 64)
	if err == nil {
		fmt.Printf("%d of type %T", n, n)
	}
	subject, err := repo.GetSubjectBySubKey(n)
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

func CheckPasswordSubject(c *gin.Context) {
	checkpass := pojo.CheckPass{}
	_ = c.BindJSON(&checkpass)
	fmt.Println(checkpass)
	subject, err := repo.GetSubjectById(checkpass.ID)
	if err != nil {
		c.JSON(200, gin.H{
			"code":        405,
			"status":      "error",
			"message":     "Not found subject",
			"description": "The subject is not find " + err.Error(),
		})
		return
	}
	if subject.Password != checkpass.Password {
		c.JSON(200, gin.H{
			"code":        405,
			"status":      "error",
			"message":     "Not found subject Password",
			"description": "The subject password is note equal",
		})
		return
	} else {
		c.JSON(200, gin.H{
			"code":        200,
			"status":      "success",
			"message":     "Fine success",
			"description": "The subject you fine is Success.",
		})
		return
	}
}
