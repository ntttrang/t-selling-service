package response

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/ntttrang/t-selling-service/messagehandler"
)

type (

	// Map map
	Map = map[string]interface{}

	//Pagination to response the information about pagination propery
	Pagination struct {
		Skip     int `json:"skip"`
		PageSize int `json:"pageSize"`
		Total    int `json:"total"`
	}

	//HTTPResponse response structure
	HTTPResponse struct {
		Meta *messagehandler.Mess `json:"meta"`
		Data interface{}          `json:"data"`
	}
)

//Response creating and filling the `HTTPResponse` struct.
func Response(ctx echo.Context, mss *messagehandler.Mess, data interface{}) error {
	if nil == mss {
		return ctx.JSON(http.StatusBadRequest, HTTPResponse{
			Meta: messagehandler.MessCode(messagehandler.E000001).GetMessage(),
			Data: Map{},
		})
	}
	return ctx.JSON(mss.StatusCode, HTTPResponse{
		Meta: mss,
		Data: data,
	})
}

// NewPagination new pagination
func NewPagination(skip, pageSize, total int) *Pagination {
	return &Pagination{
		Skip:     skip,
		PageSize: pageSize,
		Total:    total,
	}
}
