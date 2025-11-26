package user

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pdhawan2001/Go-REST-API/types"
	"github.com/pdhawan2001/Go-REST-API/utils"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
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

	// if it doesn't create the new user

}
