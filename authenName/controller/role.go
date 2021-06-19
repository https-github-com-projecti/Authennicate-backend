package controller

import (
	repo "authenName/repository"

	"github.com/gin-gonic/gin"
)

func GetRoleAll(c *gin.Context) {
	role, err := repo.GetRoleAll()
	if err != nil {
		c.JSON(200, gin.H{
			"code":        405,
			"status":      "error",
			"message":     "Not found user",
			"description": "The user you find is now Error. " + err.Error(),
		})
		return
	} else {
		c.JSON(200, gin.H{
			"code":        200,
			"status":      "success",
			"message":     "Find success",
			"description": "The user you find is Success.",
			"role":        role,
		})
		return
	}
}

func GetRoleByUserId(c *gin.Context) {
	userId := c.Param("id")
	user, err := repo.GetUserById(userId)
	if err != nil {
		c.JSON(200, gin.H{
			"code":        405,
			"status":      "error",
			"message":     "Not found user",
			"description": "The user you find is now Error. " + err.Error(),
		})
		return
	} else {
		roleId := user.Role
		role, errRole := repo.GetRoleWithIDUser(roleId.Hex())
		if errRole != nil {
			c.JSON(200, gin.H{
				"code":        405,
				"status":      "error",
				"message":     "Not found user",
				"description": "The user you find is now Error. " + err.Error(),
			})
			return
		} else {
			c.JSON(200, gin.H{
				"code":        200,
				"status":      "success",
				"message":     "Find success",
				"description": "The user you find is Success.",
				"role":        role,
			})
			return
		}
	}
}
