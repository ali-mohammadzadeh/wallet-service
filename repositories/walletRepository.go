package repositories

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"myProject/db"
)

type Wallet struct {
	Id         string  `json:"id"`
	UserName   string  `json:"username"`
	CurrencyId string  `json:"currencyId"`
	Credit     float64 `json:"credit"`
}

func (walletInstance *Wallet) Insert() (string, error) {
	walletInstance.Id = uuid.New().String()
	//db=db.SqlClient
	//if errorDb != nil {
	//	log.Error(errorDb.Error())
	//	return "", errorDb
	//}
	//defer db.SqlClient.Close()

	insert, errInsert := db.SqlClient.Prepare("insert into wallet (id,username,currencyId) values (?,?,?)")
	if errInsert != nil {
		log.Error(errInsert.Error())
		return "", errInsert
	}

	insert.Exec(walletInstance.Id, walletInstance.UserName, walletInstance.CurrencyId)
	//defer insert.Close()
	log.WithFields(
		log.Fields{
			"id":         walletInstance.Id,
			"username":   walletInstance.UserName,
			"currencyId": walletInstance.CurrencyId,
		},
	).Info("wallet created!")
	return walletInstance.Id, nil

}

func (walletInstance *Wallet) Exists() (bool, error) {
	//db, errorDb := db.GetSqlConnection()
	//if errorDb != nil {
	//	log.Error(errorDb.Error())
	//	return false, errorDb
	//}
	//defer db.Close()

	rows, errInsert := db.SqlClient.Query("SELECT * FROM wallet WHERE username = ?", walletInstance.UserName)
	if errInsert != nil {
		log.Error(errInsert.Error())
		return false, errInsert
	}
	var wallets []Wallet
	for rows.Next() {
		var wallet Wallet
		if err := rows.Scan(&wallet.Id, &wallet.UserName, &wallet.CurrencyId); err != nil {
			return false, err
		}
		wallets = append(wallets, wallet)
	}
	if err := rows.Err(); err != nil {
		return false, err
	}

	if len(wallets) >= 1 {
		return true, nil
	}
	//defer rows.Close()
	log.Info("data was")
	return false, nil

}

func (walletInstance *Wallet) FindById(username string) (Wallet, error) {
	var wallet Wallet
	//db, errorDb := db.GetSqlConnection()
	//if errorDb != nil {
	//	log.Error(errorDb.Error())
	//	return wallet, errorDb
	//}
	//defer db.Close()

	row := db.SqlClient.QueryRow("SELECT * FROM wallet WHERE username = ?", username)
	if err := row.Scan(&wallet.Id, &wallet.UserName, &wallet.CurrencyId); err != nil {
		if err == sql.ErrNoRows {
			return wallet, fmt.Errorf("albumsById %d: no such album", username)
		}
		return wallet, fmt.Errorf("albumsById %d: %v", username, err)
	}
	return wallet, nil

}
