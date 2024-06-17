package gapi

import (
	"fmt"

	db "github.com/techschool/simplebank/db/sqlc"
	"github.com/techschool/simplebank/pb"
	"github.com/techschool/simplebank/token"
	"github.com/techschool/simplebank/util"
)

// Server serves gRPC requests for our banking service.
type Server struct {
	pb.UnimplementedSimpleBankServer // Its main purpose is to enable forward compatibility, // Which means that the server can already accept the calls to the CreateUser and LoginUser RPCs before are actually implemented. 
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// NewServer creates a new gRPC server.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker:%w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	// there are no routes in gRPC.
	// The client will call the server by simply executing an RPC, just like it's calling a local function.
	return server, nil
}

// mustEmbedUnimplementedSimpleBankServer() -> in recent version of gRPC, a part from the server interface, 
// Protoc also generates this UnimplementedSimpleBankServer struct, 
// Where all RPC functions are already provided, But they all returns an unimplemented error.