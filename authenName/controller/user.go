package controller

import (
	"authenName/model"
	repo "authenName/repository"
	"crypto/sha512"
	b64 "encoding/base64"
	"github.com/gin-gonic/gin"
	"time"
)

func Login(c *gin.Context) {
	user := model.UserLogin{}
	c.BindJSON(&user)
	sDec, _ := b64.StdEncoding.DecodeString(user.Password)
	sha512 := sha512.New()
	sha512.Write(sDec)
	user.Password = b64.StdEncoding.EncodeToString(sha512.Sum(nil))
	userData, err := repo.Login(user.Username, user.Password)
	if err != nil || userData.Username == "" {
		c.JSON(200, gin.H{
			"code":        405,
			"status":      "error",
			"message":     "Username and Password do not match",
			"description": "The Username and Password you login The Login is now Error. " + err.Error(),
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
	c.BindJSON(&user)
	user.CreateAt = time.Now()
	sha512 := sha512.New()
	sha512.Write([]byte(user.Password))
	user.Password = b64.StdEncoding.EncodeToString(sha512.Sum(nil))
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
			"description": "The file you find is now Error. " + errFind.Error(),
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
