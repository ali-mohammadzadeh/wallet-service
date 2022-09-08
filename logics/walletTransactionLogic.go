package logics

import (
	"fmt"
	"myProject/repositories"
)

func AddTransactionLogic(walletInstance repositories.Wallet, walletTransactionInstance repositories.WalletTransaction) (repositories.WalletTransaction, error) {
	result, errWal := walletInstance.FindById(walletInstance.UserName)
	if errWal != nil {
		return walletTransactionInstance, fmt.Errorf("wallet not exists")
	}
	walletTransactionInstance.WalletId = result.Id
	id, errInsert := walletTransactionInstance.Insert()
	if errInsert != nil {
		return walletTransactionInstance, fmt.Errorf("transaction error!")
	}

	walletTransactionInstance.Id = id

	return walletTransactionInstance, nil
}
