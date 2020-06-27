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
	userData := model.User{}
	c.BindJSON(&user)
	sDec, _ := b64.StdEncoding.DecodeString(user.Password)
	sha_512 := sha512.New()
	sha_512.Write(sDec)
	user.Password = b64.StdEncoding.EncodeToString(sha_512.Sum(nil))
	getUserData, err := repo.Login(user.Username, user.Password)
	if err != nil || getUserData.Username == "" {
		c.JSON(200, gin.H{
			"status":  "POST",
			"message": "Fail",
		})
	} else {
		c.JSON(200, gin.H{
			"status":  "POST",
			"message": "Success",
			"user_id": userData.ID,
		})
	}
}

func GetUserById(c *gin.Context) {
	user := model.User{}
	id := c.Param("id")
	getUser, err2 := repo.GetUserById(id)
	if err2 != nil || getUser.Username == "" {
		c.JSON(200, gin.H{
			"status":  "GET",
			"message": "Fail",
		})
	} else {
		c.JSON(200, gin.H{
			"status":  "GET",
			"message": user,
		})
	}
}

func CreateUser(c *gin.Context) {
	user := model.User{}
	c.BindJSON(&user)
	user.CreateAt = time.Now()
	sha_512 := sha512.New()
	sha_512.Write([]byte(user.Password))
	user.Password = b64.StdEncoding.EncodeToString(sha_512.Sum(nil))
	res := repo.CreateUser(user)
	if res != nil {
		c.JSON(200, gin.H{
			"status":  "POST",
			"message": res,
		})
	}
	c.JSON(200, gin.H{
		"status":  "POST",
		"message": "Success",
	})
}
