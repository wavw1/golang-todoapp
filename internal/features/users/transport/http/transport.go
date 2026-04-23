package users_transport_http

import (
	"context"
	"net/http"

	"github.com/wavw1/golang-todoapp/internal/core/domain"
	core_http_server "github.com/wavw1/golang-todoapp/internal/core/transport/http/server"
)

type UsersHTTPHandler struct {
	UsersService UsersService
}

type UsersService interface {
	CreateUser(
		ctx context.Context,
		user domain.User,
	) (domain.User, error)
}

func NewUsersHTTPHandler(
	usersService UsersService,
) *UsersHTTPHandler {
	return &UsersHTTPHandler{
		UsersService: usersService,
	}
}

func (h *UsersHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/users",
			Handler: h.CreateUser,
		},
	}
}
