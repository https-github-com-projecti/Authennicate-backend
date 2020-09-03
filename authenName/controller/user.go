package controller

import (
	"authenName/model"
	repo "authenName/repository"
	"authenName/tools"
	"crypto/sha512"
	b64 "encoding/base64"
	"github.com/gin-gonic/gin"
	"time"
)

func Login(c *gin.Context) {
	user := model.UserLogin{}
	_ = c.BindJSON(&user)
	sDec, _ := b64.StdEncoding.DecodeString(user.Password)
	sha512_ := sha512.New()
	sha512_.Write(sDec)
	user.Password = b64.StdEncoding.EncodeToString(sha512_.Sum(nil))
	userData, errRepo := repo.Login(user.Username, user.Password)
	if errRepo != nil || userData.Username == "" {
		c.JSON(200, gin.H{
			"code":        405,
			"status":      "error",
			"message":     "Username and Password do not match",
			"description": "The Username and Password you login The Login is now Error. " + errRepo.Error(),
		})
		return
	} else {
		c.JSON(200, gin.H{
			"code":        200,
			"status":      "success",
			"message":     "Login Success",
			"description": "The Username and Password you login The Login is Success.",
			"id":          userData.ID,
		})
		return
	}
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	user, errFind := repo.GetUserById(id)
	if errFind != nil || user.Username == "" {
		c.JSON(200, gin.H{
			"code":        405,
			"status":      "error",
			"message":     "Not found user",
			"description": "The user you find is now Error. " + errFind.Error(),
		})
		return
	} else {
		c.JSON(200, gin.H{
			"code":        200,
			"status":      "success",
			"message":     "Find success",
			"description": "The user you find is Success.",
			"user":        user,
		})
	}
}

func CreateUser(c *gin.Context) {
	user := model.User{}
	_ = c.BindJSON(&user)
	sha512_ := sha512.New()
	sha512_.Write([]byte(user.Password))
	user.Password = b64.StdEncoding.EncodeToString(sha512_.Sum(nil))
	user.CreateAt = tools.TimeNow()
	user.CreateEnd = tools.TimeNow()
	user.CreateUpdate = time.Time{}

	res := repo.CreateUser(user)
	if res != nil {
		c.JSON(200, gin.H{
			"code":        405,
			"status":      "error",
			"message":     "Cannot create user",
			"description": "The user you create is now Error. " + res.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":        200,
		"status":      "success",
		"message":     "Create User Success",
		"description": "The user you create is Success.",
	})
}

func ImageByUserId(c *gin.Context) {
	id := c.Param("id")
	user, errFind := repo.GetUserById(id)
	if errFind != nil || user.Username == "" {
		c.JSON(200, gin.H{
			"code":        405,
			"status":      "error",
			"message":     "Not found user",
			"description": "The user you find is now Error. " + errFind.Error(),
		})
		return
	}
	stringObjectID := user.Upload.Hex()
	upload, errUp := repo.FindById(stringObjectID)
	if errUp != nil {
		c.JSON(200, gin.H{
			"code":        405,
			"status":      "error",
			"message":     "Not found file",
			"description": "The file you find is now Error. " + errUp.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":        200,
		"status":      "success",
		"message":     "Find Path Success",
		"description": "The path you find is Success.",
		"path":        upload.Path,
	})
}

func DeleteUserById(c *gin.Context) {
	id := c.Param("id")
	result, errDel := repo.DeleteUserById(id)
	if errDel != nil {
		c.JSON(200, gin.H{
			"code":        304,
			"status":      "error",
			"message":     "File cannot be remove from database",
			"description": "The User you delete The delete is now Error. " + errDel.Error(),
		})
		return
	}

	if result != nil {
		c.JSON(200, gin.H{
			"code":        200,
			"status":      "success",
			"message":     "Successfully remove file",
			"description": "The User you delete The delete is now complete is " + id + " .",
		})
	}
	return
}

func UpdateUserById(c *gin.Context) {
	//user := model.User

}
