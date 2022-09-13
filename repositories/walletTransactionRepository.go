package repositories

import (
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"myProject/db"
)

type WalletTransaction struct {
	Id           string `json:"id"`
	WalletId     string `json:"walletId"`
	Amount       int    `json:"amount"`
	Description  string `json:"description"`
	TrackingCode string `json:"trackingCode"`
	Type         int    `json:"type"`
}

func (instance *WalletTransaction) Insert() (string, error) {
	instance.Id = uuid.New().String()
	//db, errorDb := db.GetSqlConnection()
	//if errorDb != nil {
	//	log.Error(errorDb.Error())
	//	return "", errorDb
	//}
	//defer db.Close()

	insert, errInsert := db.SqlClient.Prepare("insert into wallet_transactions (id,walletId,amount,type,description,trackingCode) values (?,?,?,?,?,?)")
	if errInsert != nil {
		log.Error(errInsert.Error())
		return "", errInsert
	}

	insert.Exec(instance.Id, instance.WalletId, instance.Amount, instance.Type, instance.Description, instance.TrackingCode)
	defer insert.Close()
	log.WithFields(
		log.Fields{
			"id":          instance.Id,
			"walletId":    instance.WalletId,
			"amount":      instance.Amount,
			"description": instance.Description,
		},
	).Info("wallet created!")
	return instance.Id, nil

}

func (instance *WalletTransaction) CalculateCredit(walletId string) (float64, error) {
	//db, _ := db.GetSqlConnection()
	//defer db.Close()

	rows, _ := db.SqlClient.Query("select sum(amount) as credit from wallet_transactions  WHERE walletId=?", walletId)

	var credit float64
	for rows.Next() {
		rows.Scan(&credit)
	}

	return credit, nil
}
