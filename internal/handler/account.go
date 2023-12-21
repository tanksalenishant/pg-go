package handler

import (
	"net/http"
	"rest1/internal/domain"
	"rest1/internal/usecases"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)


type AccountHandler struct {
	UseCase *usecases.AccountUsecase
	conn *pgx.Conn
	logger *zap.Logger
}

func NewAccountHandler(useCase *usecases.AccountUsecase , conn *pgx.Conn, logger *zap.Logger) *AccountHandler{
	return &AccountHandler{
		UseCase: useCase,
		conn: conn,
		logger: logger,
	}
}


// Create Account route
// http://localhost:3000/{userid}/account/create
func (h *AccountHandler) CreateAccountHandler(w http.ResponseWriter, r *http.Request){
	var account domain.Account
	idStr := chi.URLParam(r, "userid")
	userId, err := strconv.Atoi(idStr)

	// check if user from req.body is valid
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		h.logger.Error("Failed to Create Account Handler", zap.Error(err))
		return 
	}

	_, err = h.UseCase.CreateAccount(userId, h.conn)
	
	respondWithJSON(w, http.StatusOK, account)
}


// http://localhost:3000/account/{accoundId}
func (h *AccountHandler) GetByAccountNoHandler(w http.ResponseWriter, r *http.Request){
	// get ID from url parameters
	idStr := chi.URLParam(r, "accoundId")
	accountId, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		h.logger.Error("Failed to get account by ID at Handler layer", zap.Error(err))
		return 
	}


	acc, err := h.UseCase.GetByAccountNo(accountId, h.conn)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	respondWithJSON(w, http.StatusOK, acc)
}