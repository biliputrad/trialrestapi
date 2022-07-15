package Helper

import (
	"RestAPI-GETNPOST/dtos"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GeneratePaginationRequest(c *gin.Context) *dtos.Pagination {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	sort := c.DefaultQuery("sort", "user_name desc")

	return &dtos.Pagination{Limit: limit, Page: page, Sort: sort}
}
