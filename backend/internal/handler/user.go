package handler

import (
	"errors"
	"net/http"

	"saffronreal.estate/backend/internal/repository"
	"saffronreal.estate/backend/internal/response"
	"saffronreal.estate/backend/internal/service"
	"saffronreal.estate/backend/internal/util"
)

type UserHandler struct {
	users *service.UserService
}

func NewUserHandler(
	users *service.UserService,
) *UserHandler {
	return &UserHandler{
		users: users,
	}
}

type createUserRequest struct {
	Email string `json:"email"`
}

type getUserRequest struct {
	Email string `json:"email"`
}

func (h *UserHandler) Create(
	w http.ResponseWriter,
	r *http.Request,
) {
	var request createUserRequest

	err := util.DecodeJSON(w, r, &request)
	if err != nil {
		writeUserError(w, err)
		return
	}

	user, err := h.users.Create(
		r.Context(),
		request.Email,
	)
	if err != nil {
		writeUserError(w, err)
		return
	}

	response.JSON(w, http.StatusCreated, user)
}

func (h *UserHandler) GetByEmail(
	w http.ResponseWriter,
	r *http.Request,
) {
	var request getUserRequest

	err := util.DecodeJSON(w, r, &request)
	if err != nil {
		writeUserError(w, err)
		return
	}

	user, err := h.users.Get(r.Context(), request.Email)
	if err != nil {
		writeUserError(w, err)
		return
	}

	response.JSON(w, http.StatusOK, user)
}

func writeUserError(
	w http.ResponseWriter,
	err error,
) {
	switch {
	case errors.Is(err, repository.ErrUserNotFound):
		response.Error(
			w,
			http.StatusNotFound,
			"user_not_found",
			"user not found",
		)

	case errors.Is(err, repository.ErrEmailConflict):
		response.Error(
			w,
			http.StatusConflict,
			"email_conflict",
			"email is already in use",
		)

	case errors.Is(err, service.ErrInvalidEmail):
		response.Error(
			w,
			http.StatusBadRequest,
			"invalid_email",
			"email address is invalid",
		)

	default:
		response.Error(
			w,
			http.StatusInternalServerError,
			"internal_error",
			"an internal server error occurred)",
		)
	}
}
