package logics

import (
	"fmt"
	"myProject/repositories"
)

func AddWalletLogic(walletInstance repositories.Wallet) (repositories.Wallet, error) {
	isWallet, errGet := walletInstance.Exists()
	if errGet != nil {
		return walletInstance, fmt.Errorf("error in db")
	}
	if isWallet {
		return walletInstance, fmt.Errorf("wallet exists")
	}
	id, errInsert := walletInstance.Insert()
	if errInsert != nil {
		return walletInstance, fmt.Errorf("wallet exists")
	}

	walletInstance.Id = id

	return walletInstance, nil
}
