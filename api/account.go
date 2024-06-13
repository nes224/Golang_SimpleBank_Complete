package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	request "github.com/techschool/simplebank/api/request"
	db "github.com/techschool/simplebank/db/sqlc"
)

func (server *Server) createAccount(ctx *gin.Context) { // server * Server pointer receiver.
	// ctx *gin.Context -> *Context input -> Basically, when using gin, everything we do inside a handler will involve this context object.
	// It provides a lot of convenient methods to read input parameters and write out responses.
	// ShouldBindJSON function to parse the input data from HTTP request body and Gin will validate the output object to make sure
	// it satisfy the conditions we specified in the binding tag.
	var req request.CreateAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

func (server *Server) getAccount(ctx *gin.Context) {
	var req request.GetAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	account, err := server.store.GetAccountForUpdate(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, account)
}


func (server *Server) listAccount(ctx *gin.Context) {
	var req request.ListAccountRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListAccountsParams{
		Limit: req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	account, err := server.store.ListAccounts(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// account = db.Account{} // change this account to an empty object 

	ctx.JSON(http.StatusOK, account)
}