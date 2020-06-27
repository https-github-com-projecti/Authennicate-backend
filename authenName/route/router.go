package route

import (
	"authenName/config"
	api "authenName/controller"
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
		AllowOrigins:     []string{"http://localhost:3001/"},
		AllowMethods:     []string{"POST, GET, OPTIONS, PUT, DELETE, UPDATE"},
		AllowHeaders:     []string{"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
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
	}

	upload := r.Group("/upload")
	{
		upload.Static("/assets", "Upload/assets")
		upload.Static("/profile", "Upload/profile")
		upload.POST("/uploadAssets", api.UploadAssets)
		upload.POST("/uploadProfile", api.UploadProfile)
		upload.DELETE("/delete/file/:id", api.DeleteFileNoUser)
	}
	r.Run(config.PortServe)
}
