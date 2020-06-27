package controller

import (
	"authenName/model"
	repo "authenName/repository"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func UploadAssets(c *gin.Context) {
	err := os.MkdirAll("Upload/assets", 0755)
	check(err)
	// Multipart form
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(200, gin.H{
			"code":        405,
			"status":      "error",
			"message":     "Techniques do not match",
			"description": "The image you uploaded The upload is now Error. " + err.Error(),
		})
		return
	}
	files := form.File["files"]
	for _, file := range files {
		filename := filepath.Base(file.Filename)
		uploads := model.Upload{}
		res, errFind := repo.FindByPathAndStatus("Upload/profile/"+filename, "Profile")
		if errFind == nil {
			c.JSON(200, gin.H{
				"code":        200,
				"status":      "warning",
				"message":     "This file already exists",
				"description": "The image you uploaded The image already exists in the database.",
				"id":          res.ID,
			})
			return
		} else {
			if err := c.SaveUploadedFile(file, "Upload/assets/"+filename); err != nil {
				c.JSON(200, gin.H{
					"code":        304,
					"status":      "error",
					"message":     "File cannot be saved",
					"description": "The image you uploaded The upload is now Error. " + err.Error(),
				})
				return
			}
			uploads.Path = "Upload/profile/" + filename
			uploads.Status = "Assets"
			uploads.Name = filename
			res, err := repo.InsertUpload(uploads)
			if err != nil {
				c.JSON(200, gin.H{
					"code":        304,
					"status":      "error",
					"message":     "File cannot be saved to database",
					"description": "The image you uploaded The upload is now Error. " + err.Error(),
				})
				return
			}
			c.JSON(200, gin.H{
				"code":        200,
				"status":      "success",
				"message":     "Successfully saved data",
				"description": "The image you uploaded The upload is now complete is " + filename + " files.",
				"id":          res.InsertedID,
			})
		}
	}
}

func UploadProfile(c *gin.Context) {
	err := os.MkdirAll("Upload/profile", 0755)
	check(err)

	// single file
	file, _ := c.FormFile("file")
	log.Println(file.Filename)
	filename := c.PostForm("filename")

	uploads := model.Upload{}
	res, errFind := repo.FindByPathAndStatus("Upload/profile/"+filename, "Profile")
	fmt.Println("err ", err)
	if errFind == nil {
		c.JSON(200, gin.H{
			"code":        200,
			"status":      "warning",
			"message":     "This file already exists",
			"description": "The image you uploaded The image already exists in the database.",
			"id":          res.ID,
		})
		return
	} else {
		// Upload the file to specific dst.
		if err := c.SaveUploadedFile(file, "Upload/profile/"+filename); err != nil {
			c.JSON(200, gin.H{
				"code":        304,
				"status":      "error",
				"message":     "File cannot be saved",
				"description": "The image you uploaded The upload is now Error. " + err.Error(),
			})
			return
		}
		uploads.Path = "Upload/profile/" + filename
		uploads.Status = "Profile"
		uploads.Name = filename
		res, err := repo.InsertUpload(uploads)
		if err != nil {
			c.JSON(200, gin.H{
				"code":        304,
				"status":      "error",
				"message":     "File cannot be saved to database",
				"description": "The image you uploaded The upload is now Error. " + err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"code":        200,
			"status":      "success",
			"message":     "Successfully saved data",
			"description": "The image you uploaded The upload is now complete is " + filename + " files.",
			"id":          res.InsertedID,
		})
		return
	}
}

func DeleteFileNoUser(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("id : ", id)
	file, err := repo.FindById(id)
	if err != nil {
		c.JSON(200, gin.H{
			"code":        404,
			"status":      "error",
			"message":     "Cannot find file",
			"description": "The image you delete The delete is now Error. " + err.Error(),
		})
		return
	}
	if err = os.Remove(file.Path); err != nil {
		c.JSON(200, gin.H{
			"code":        404,
			"status":      "error",
			"message":     "Cannot remove file",
			"description": "The image you delete The delete is now Error. " + err.Error(),
		})
		return
	}

	_, errDel := repo.DeleteById(id)
	if errDel != nil {
		c.JSON(200, gin.H{
			"code":        304,
			"status":      "error",
			"message":     "File cannot be remove from database",
			"description": "The image you delete The delete is now Error. " + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":        200,
		"status":      "success",
		"message":     "Successfully remove file",
		"description": "The image you delete The delete is now complete is " + file.Name + " file.",
	})
	return
}