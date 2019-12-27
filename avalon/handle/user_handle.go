package handle

import (
	"Avalon/avalon/types"
	"Avalon/avalon/users"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserById(c *gin.Context) {
	id, _ := c.GetQuery("id")
	vid, err := strconv.Atoi(id)
	if err != nil {
		return
	}

	item, err := users.GetUserById(vid)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, item)
}

func AddUser(c *gin.Context) {
	user := &types.UserInfo{}
	if err := c.ShouldBindJSON(user); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if err := users.AddUser(user); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

func UpdateUser(c *gin.Context) {
	user := &types.UserInfo{}
	if err := c.ShouldBindJSON(user); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if user.Id == 0 {
		c.Status(http.StatusBadRequest)
		return
	}

	if err := users.UpdateUser(user); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

func DeleteUserById(c *gin.Context) {
	id, _ := c.GetQuery("id")
	vid, err := strconv.Atoi(id)
	if err != nil {
		return
	}

	if err := users.DeleteUserById(vid); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}
