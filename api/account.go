package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/linhhuynhcoding/learn-go/db/accountdb"
)

// createAccountReq define Create Account Request Body
type createAccountReq struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD VND"`
}

// createAccount handles create new account logic
func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err, "Your request body invalid!"))
		return
	}

	arg := accountdb.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
	}

	account, err := server.store.Account.CreateAccount(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err, "Server can not process your request!"))
		return
	}

	ctx.JSON(http.StatusOK, apiResponse([]accountdb.Account{account}, "Create account successfully!"))
}

func (server *Server) getAccountById(ctx *gin.Context) {
	var id int64
	var err error
	var account accountdb.Account

	id, err = strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err, "The param should be integer!"))
		return
	}

	account, err = server.store.Account.GetAccount(ctx, id)

	ctx.JSON(http.StatusOK, apiResponse([]accountdb.Account{account}, "Get account successfully!"))
}
