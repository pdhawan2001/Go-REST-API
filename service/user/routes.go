package user

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pdhawan2001/Go-REST-API/service/auth"
	"github.com/pdhawan2001/Go-REST-API/types"
	"github.com/pdhawan2001/Go-REST-API/utils"
)

// we don't have to put store as a pointer
// because the interface already references the concrete implementation.
type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")

}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// get json payload
	var payload types.RegisterUserPayload // payload := types.RegisterUserPayload{} equivalent shorthand

	// the body is being decoded into the payload in the ParseJSON function
	if err := utils.ParseJSON(r, payload); err != nil {
		// StatusBadRequest because this is an User error
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	// check if the user exists
	// here we are are using dash (_) because we dont want to return the user, we just want to return the error here
	// because the user does not exist
	// go wants us to use every variable that's why we use something called as a dash (_) to not return that
	// but just keep in mind that it exists there
	_, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", payload.Email))
		return
	}

	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// if it doesn't create the new user
	err = h.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)

}
