package content

import (
	"PH_ModuleName_PH/tool"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Post(c *gin.Context) {
	var row Content
	if err := c.ShouldBind(&row); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if result := tool.DB.Create(&row); result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusCreated, row)
}
func Delete(c *gin.Context) {
	id := c.Param("id")
	if result := tool.DB.Delete(&Content{}, "id = ?", id); result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
func Put(c *gin.Context) {
	id := c.Param("id")
	var row Content
	if result := tool.DB.First(&row, "id = ?", id); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": tool.MResourceNotFound,
		})
		return
	}

	if err := c.ShouldBind(&row); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if result := tool.DB.Save(&row); result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, row)
}

func List(c *gin.Context) {
	var rows []Content
	if result := tool.DB.Find(&rows); result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusOK, []Content{})
		return
	}
	c.JSON(http.StatusOK, rows)
}

func Retrieve(c *gin.Context) {
	id := c.Param("id")
	var row Content
	if result := tool.DB.First(&row, "id = ?", id); result.Error != nil {
		log.Println(result.Error, row)
		c.JSON(http.StatusOK, nil)
		return
	}
	c.JSON(http.StatusOK, row)
}
