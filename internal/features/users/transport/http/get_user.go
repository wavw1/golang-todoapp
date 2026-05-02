package users_transport_http

import (
	"net/http"

	core_logger "github.com/wavw1/golang-todoapp/internal/core/logger"
	core_http_response "github.com/wavw1/golang-todoapp/internal/core/transport/http/response"
	core_http_utils "github.com/wavw1/golang-todoapp/internal/core/transport/http/utils"
)

type GetUserResponse UserDTOResponse

// GetUser 		godoc
// @Summary 	Получение пользователя
// @Description Получение конкретного пользователя по его ID
// @Tags 		users
// @Produce		json
// @Param 		id path int true "ID получаемого пользователя"
// @Success 	200   {object} GetUserResponse "Пользователь успешно найден"
// @Failure 	400   {object} core_http_response.ErrorResponse "Bad request"
// @Failure 	404   {object} core_http_response.ErrorResponse "User not found"
// @Failure 	500   {object} core_http_response.ErrorResponse "Internal server error"
// @Router 		/users/{id} [get]
func (h *UsersHTTPHandler) GetUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, err := core_http_utils.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get userID path value",
		)

		return
	}

	user, err := h.UsersService.GetUser(ctx, userID)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get user",
		)

		return
	}

	response := GetUserResponse(userDTOFromDomain(user))

	responseHandler.JSONResponse(response, http.StatusOK)
}
