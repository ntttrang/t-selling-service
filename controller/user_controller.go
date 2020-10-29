package controller

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo"
	"github.com/ntttrang/t-selling-service/messagehandler"
	"github.com/ntttrang/t-selling-service/request"
	"github.com/ntttrang/t-selling-service/response"
)

//GetUsers get user
func (c *EchoController) GetUsers() interface{} {
	var f echo.HandlerFunc = func(ctx echo.Context) error {
		var (
			firstName   = ctx.QueryParam("firstName")
			lastName    = ctx.QueryParam("lastName")
			countryCode = ctx.QueryParam("countryCode")
			// Pagination information
			skipS     = ctx.QueryParam("skip")
			pageSizeS = ctx.QueryParam("pageSize")
			sortBy    = ctx.QueryParam("sortBy")
			direction = ctx.QueryParam("direction")
		)

		var skip, pageSize int
		var e error
		if "" != skipS {
			skip, e = strconv.Atoi(skipS)
			if nil != e {
				fmt.Println("Convert `skip` error. Use default value `skip = 0`")
			}
		}
		if "" != pageSizeS {
			pageSize, e = strconv.Atoi(pageSizeS)
			if nil != e {
				fmt.Println("Convert `pageSize` error. Use default value `pageSize = 10`")
				pageSize = 10
			}
		}

		pagin := request.PaginationInfo{
			Skip:      skip,
			PageSize:  pageSize,
			SortBy:    sortBy,
			Direction: direction,
		}
		userSearch := request.UserSearch{
			FirstName:      firstName,
			LastName:       lastName,
			CountryCode:    countryCode,
			PaginationInfo: pagin,
		}
		users, err := c.service.GetUsers(userSearch)
		fmt.Println(users)
		if nil != err {
			return response.Response(ctx, err, response.Map{})
		}
		total := len(users)
		//Set pageSize is the length of user list if not specify pageSize.
		if "" == pageSizeS {
			pageSize = len(users)
		}
		return response.Response(ctx, messagehandler.I000008.GetMessage(), response.Map{
			"items":      users,
			"pagination": response.NewPagination(skip, pageSize, total),
		})
	}
	return f
}

// InsertUser add new user
func (c *EchoController) InsertUser() interface{} {
	var f echo.HandlerFunc = func(ctx echo.Context) error {
		var req = request.NewUserCreateReq()
		if e := ctx.Bind(&req); nil != e {
			err := messagehandler.W000005.GetMessage()
			return response.Response(ctx, err, response.Map{})
		}
		err := c.service.InsertUser(req)
		if nil != err {
			return response.Response(ctx, err, response.Map{})
		}
		return response.Response(ctx, messagehandler.I000008.GetMessage(), response.Map{})
	}
	return f
}

// UpdateUser update user
func (c *EchoController) UpdateUser() interface{} {
	var f echo.HandlerFunc = func(ctx echo.Context) error {
		var req = request.NewUserUpdateReq()
		if e := ctx.Bind(&req); nil != e {
			err := messagehandler.W000005.GetMessage()
			return response.Response(ctx, err, response.Map{})
		}
		err := c.service.UpdateUser(req)
		if nil != err {
			return response.Response(ctx, err, response.Map{})
		}
		return response.Response(ctx, messagehandler.I000008.GetMessage(), response.Map{})
	}
	return f
}

// DeleteUser delete user
func (c *EchoController) DeleteUser() interface{} {
	var f echo.HandlerFunc = func(ctx echo.Context) error {
		var userId = ctx.QueryParam("userId")
		err := c.service.DeleteUser(userId)
		if nil != err {
			return response.Response(ctx, err, response.Map{})
		}
		return response.Response(ctx, messagehandler.I000008.GetMessage(), response.Map{})
	}
	return f
}
