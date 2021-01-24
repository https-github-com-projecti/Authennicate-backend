package route

import (
	api "authenName/controller"
	"authenName/properties"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func Router(r *gin.Engine) {
	// CORS for https://foo.com and https://github.com origins, allowing:
	// - PUT and PATCH methods
	// - Origin header
	// - Credentials share
	// - Preflight requests cached for 12 hours
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://fourdust.kozow.com:3000", "http://localhost:3000"},
		AllowMethods:     []string{"POST, GET, OPTIONS, PUT, DELETE, UPDATE"},
		AllowHeaders:     []string{"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With", "XMLHttpRequest"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3001"
		},
		MaxAge: 12 * time.Hour,
	}))
	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	test := r.Group("/test")
	{
		test.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
	// Authorization group
	// authorized := r.Group("/", AuthRequired())
	// exactly the same as:
	authorized := r.Group("/")
	// per group middleware! in this case we use the custom created
	// AuthRequired() middleware just in the "authorized" group.
	authorized.Use()
	{
		user := authorized.Group("user")
		user.POST("/create", api.CreateUser)
		user.GET("/getId/:id", api.GetUserById)
		user.POST("/login", api.Login)
		user.GET("/image/:id", api.ImageByUserId)
		user.DELETE("/delete/:id", api.DeleteUserById)
	}

	upload := r.Group("/upload")
	{
		upload.Static("/assets", properties.Path + "/export/files/Upload/assets")
		upload.Static("/profile", properties.Path + "/export/files/Upload/profile")
		upload.Static("/qrcode",  properties.Path + "/export/files/Upload/qrcode")
		upload.POST("/uploadAssets", api.UploadAssets)
		upload.POST("/uploadProfile", api.UploadProfile)
		upload.POST("/uploadQrcode", api.UploadQrCode)
		upload.DELETE("/delete/file/:id", api.DeleteFileNoUser)
		upload.GET("/data-upload/:id", api.GetPathUploadById)
	}

	subject := r.Group("/subject")
	{
		subject.POST("/create", api.CreateSubject)
		subject.GET("/subject-all/:id", api.GetSubjectAll)
		subject.DELETE("/delete/:id", api.DeleteSubject)
		subject.GET("/subject/:id", api.GetSubject)
	}

	authen := r.Group("/authen")
	{
		authen.POST("/create", api.CreateAuthen)
		authen.GET("/subject-authen/:id", api.GetAuthenAllForSubject)
	}

	r.Run(properties.PortServe)
}
