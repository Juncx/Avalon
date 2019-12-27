package handle

import (
	"net/http"
	"strconv"

	"Avalon/avalon/items"
	"Avalon/avalon/types"

	"github.com/gin-gonic/gin"
)

func GetItemById(c *gin.Context) {
	id, _ := c.GetQuery("id")
	vid, err := strconv.Atoi(id)
	if err != nil {
		return
	}

	item, err := items.GetItemById(vid)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, item)
}

func AddItem(c *gin.Context) {
	item := &types.ItemInfo{}
	if err := c.ShouldBindJSON(item); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if err := items.AddItem(item); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

func UpdateItem(c *gin.Context) {
	item := &types.ItemInfo{}
	if err := c.ShouldBindJSON(item); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if item.Id == 0 {
		c.Status(http.StatusBadRequest)
		return
	}

	if err := items.UpdateItem(item); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

func DeleteItemById(c *gin.Context) {
	id, _ := c.GetQuery("id")
	vid, err := strconv.Atoi(id)
	if err != nil {
		return
	}

	if err := items.DeleteItemById(vid); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}
