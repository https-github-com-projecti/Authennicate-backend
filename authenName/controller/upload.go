package controller

import (
	"authenName/model"
	repo "authenName/repository"
	"authenName/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
	"time"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func UploadAssets(c *gin.Context) {
	err := os.MkdirAll("Upload/assets", 0444)
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
		res, errFind := repo.FindByPathAndStatus("upload/profile/"+filename, "Profile")
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
			if err := c.SaveUploadedFile(file, "upload/assets/"+filename); err != nil {
				c.JSON(200, gin.H{
					"code":        304,
					"status":      "error",
					"message":     "File cannot be saved",
					"description": "The image you uploaded The upload is now Error. " + err.Error(),
				})
				return
			}
			uploads.Path = "upload/assets/" + filename
			uploads.Status = "Assets"
			uploads.Name = filename
			uploads.CreateAt = tools.TimeNow()
			uploads.CreateEnd = tools.TimeNow()
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
		time.Sleep(500 * time.Millisecond)
	}
}

func UploadProfile(c *gin.Context) {
	err := os.MkdirAll("Upload/profile", 0444)
	check(err)

	// single file
	file, _ := c.FormFile("file")
	filename := c.PostForm("filename")

	uploads := model.Upload{}
	res, errFind := repo.FindByPathAndStatus("upload/profile/"+filename, "Profile")
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
		if err := c.SaveUploadedFile(file, "upload/profile/"+filename); err != nil {
			c.JSON(200, gin.H{
				"code":        304,
				"status":      "error",
				"message":     "File cannot be saved",
				"description": "The image you uploaded The upload is now Error. " + err.Error(),
			})
			return
		}
		uploads.Path = "upload/profile/" + filename
		uploads.Status = "Profile"
		uploads.Name = filename
		uploads.CreateAt = tools.TimeNow()
		uploads.CreateEnd = tools.TimeNow()
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
			"description": "The image you delete The delete is now Error. " + errDel.Error(),
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

func UploadQrCode(c *gin.Context) {
	err := os.MkdirAll("Upload/qrcode", 0444)
	check(err)

	// single file
	file, _ := c.FormFile("file")
	filename := c.PostForm("filename")

	uploads := model.Upload{}
	res, errFind := repo.FindByPathAndStatus("upload/qrcode/"+filename, "Qrcode")
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
		if err := c.SaveUploadedFile(file, "upload/qrcode/"+filename); err != nil {
			c.JSON(200, gin.H{
				"code":        304,
				"status":      "error",
				"message":     "File cannot be saved",
				"description": "The image you uploaded The upload is now Error. " + err.Error(),
			})
			return
		}
		uploads.Path = "upload/qrcode/" + filename
		uploads.Status = "Qrcode"
		uploads.Name = filename
		uploads.CreateAt = tools.TimeNow()
		uploads.CreateEnd = tools.TimeNow()
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

func GetPathUploadById(c *gin.Context){
	id := c.Param("id")
	upload, errFind := repo.FindByUserId(id)
	if errFind != nil || upload.Name == "" {
		c.JSON(200, gin.H{
			"code":        405,
			"status":      "error",
			"message":     "Not found user",
			"description": "The upload you find is now Error. " + errFind.Error(),
		})
		return
	} else {
		c.JSON(200, gin.H{
			"code":        200,
			"status":      "success",
			"message":     "Find success",
			"description": "The upload you find is Success.",
			"upload":        upload,
		})
	}
}
