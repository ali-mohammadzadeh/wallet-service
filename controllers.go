package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"myProject/helpers"
	"myProject/logics"
	"myProject/repositories"
	"net/http"
)

var errorInTransaction helpers.ErrorStruct

func add(writer http.ResponseWriter, request *http.Request) {

	var walletInstance repositories.Wallet
	err := json.NewDecoder(request.Body).Decode(&walletInstance)
	if err != nil {
		helpers.ErrorResponse(writer, err.Error(), http.StatusBadRequest)
		return
	}

	result, isErr := logics.AddWalletLogic(walletInstance)
	if isErr != nil {
		helpers.ErrorResponse(writer, isErr.Error(), http.StatusBadRequest)
		return
	}
	helpers.SuccessResponse(writer, result, http.StatusCreated)
}

func addTransaction(writer http.ResponseWriter, request *http.Request) {

	var walletInstance repositories.Wallet
	var walletTransactionInstance repositories.WalletTransaction
	inputs := mux.Vars(request)
	err := json.NewDecoder(request.Body).Decode(&walletTransactionInstance)
	if err != nil {
		errorInTransaction.Message = err.Error()
		helpers.SuccessResponse(writer, errorInTransaction, http.StatusInternalServerError)
		return
	}
	walletInstance.UserName = inputs["username"]

	result, errInsert := logics.AddTransactionLogic(walletInstance, walletTransactionInstance)
	if errInsert != nil {
		errorInTransaction.Message = errInsert.Error()
		helpers.SuccessResponse(writer, errorInTransaction, http.StatusInternalServerError)
		return
	}

	helpers.SuccessResponse(writer, result, http.StatusCreated)
}

func get(writer http.ResponseWriter, request *http.Request) {
	inputs := mux.Vars(request)

	var wallet repositories.Wallet
	var walletTransaction repositories.WalletTransaction
	result, errWal := wallet.FindById(inputs["username"])
	if errWal != nil {

		helpers.ErrorResponse(writer, "wallet not exists", http.StatusNotFound)
		return
	}
	credit, _ := walletTransaction.CalculateCredit(result.Id)

	result.Credit = credit
	helpers.SuccessResponse(writer, result, http.StatusCreated)

}
