package api

import "github.com/gin-gonic/gin"

//Server serves https request for service
type Server struct {
	store  *db.store
	router *gin.Engine
}

//NewServer Creats a new http server and setup routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)

	server.router = router
	return server
}

//Starts runs the http server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
