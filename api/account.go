package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/techschool/simplebank/db/sqlc"
)

type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`    // golang to json
	Currency string `json:"currency" binding:"required,oneof=USD EUR"` // golang to json
}


func (server *Server) createAccount(ctx *gin.Context) { // server * Server pointer receiver.
	// ctx *gin.Context -> *Context input -> Basically, when using gin, everything we do inside a handler will involve this context object.
	// It provides a lot of convenient methods to read input parameters and write out responses.
	// ShouldBindJSON function to parse the input data from HTTP request body and Gin will validate the output object to make sure
	// it satisfy the conditions we specified in the binding tag.
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateAccountParams{
		Owner: req.Owner,
		Currency: req.Currency,
		Balance: 0,
	}
	
	account, err := server.store.CreateAccount(ctx,arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

