package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/linhhuynhcoding/learn-go/db/accountdb"
)

type createAccountReq struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD VND"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	arg := accountdb.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
	}

	account, err := server.store.Account.CreateAccount(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}
