package handle

import "github.com/gin-gonic/gin"

func Version(c *gin.Context) {
	version := "v0.0.1"
	c.JSON(200, gin.H{
		"version": version,
	})
}

func ItemRegister(r *gin.Engine) {
	r.GET("/item/get", GetItemById)
	r.POST("/item/add", AddItem)
	r.POST("/item/update", UpdateItem)
	r.DELETE("/item/delete", DeleteItemById)
}

func UserRegister(r *gin.Engine) {
	r.GET("/user/get", GetUserById)
	r.POST("/user/add", AddUser)
	r.POST("/user/update", UpdateUser)
	r.DELETE("/user/delete", DeleteUserById)
}

func Router(r *gin.Engine) {
	r.GET("/version", Version)

	// registe item http handle
	ItemRegister(r)

	// registe user http handle
	UserRegister(r)
}
