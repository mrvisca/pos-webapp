package helpers

import (
	"github.com/gin-gonic/gin"
)

func ElorResponse(c *gin.Context, message string) {
	c.JSON(400, gin.H{
		"status":  "Elor",
		"message": message,
	})
}

func SuksesWithDataResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(201, gin.H{
		"status":  "Sukses",
		"message": message,
		"data":    data,
	})
}

func DataResponse(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"data": data,
	})
}

func SuksesResponse(c *gin.Context, message string) {
	c.JSON(201, gin.H{
		"status":  "Sukses",
		"message": message,
	})
}

func SuksesLogin(c *gin.Context, message string, token string, roleid int64) {
	c.JSON(201, gin.H{
		"status":  "Sukses",
		"message": message,
		"token":   token,
		"role_id": roleid,
	})
}

func ElorWithData(c *gin.Context, message string, elor interface{}) {
	c.JSON(401, gin.H{
		"status":  "Elor",
		"message": message,
		"elor":    elor,
	})
}

func ListPaginate(c *gin.Context, data interface{}, totalRecords interface{}, draw int) {
	c.JSON(200, gin.H{
		"draw":            draw,
		"recordsTotal":    totalRecords,
		"recordsFiltered": totalRecords,
		"data":            data,
	})
}
