package api

import (
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/linhhuynhcoding/learn-go/db"
)

// Server serves HTTP requests for our bank service
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccountById)
	// router.GET("/accounts", server.getAccounts)
	// router.PATCH("/accounts/:id", server.updateAccount)

	// add routes to router
	server.router = router
	return server
}

// Start runs HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

type ErrorResponse struct {
	Success    bool   `json:"success"`
	ErrorCode  string `json:"error_code"`
	Message    string `json:"message"`
	Error      string `json:"error"`
	Timestampz string `json:"timestampz"`
}

// errorResponse returns Error Response
func errorResponse(err error, mess string) ErrorResponse {
	return ErrorResponse{
		Success:    false,
		ErrorCode:  "0",
		Message:    mess,
		Error:      err.Error(),
		Timestampz: time.Now().UTC().Format(time.RFC3339),
	}
}

type ApiResponse[T any] struct {
	Success    bool   `json:"success"`
	Message    string `json:"message"`
	Data       []T    `json:"data"`
	Timestampz string `json:"timestampz"`
}

// apiResponse returns Api Response
func apiResponse[T any](data []T, mess string) ApiResponse[T] {
	return ApiResponse[T]{
		Success:    true,
		Message:    mess,
		Data:       data,
		Timestampz: time.Now().UTC().Format(time.RFC3339),
	}
}
