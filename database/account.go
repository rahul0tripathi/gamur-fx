package database

import (
	"context"
)

const (
	TransactionStatusFailed  = "failed"
	TransactionStatusSuccess = "successful"
	TransactionTypeDeducted  = "deducted"
	TransactionTypeAdded     = "added"
)

type Txn struct {
	Id     int     `json:"id"`
	Title  string  `json:"title"`
	Amount float64 `json:"amount"`
	Status string  `json:"status"`
	Type   string  `json:"type"`
}

const UPDATE_BALANCE = `UPDATE Users SET balance = balance - ? WHERE id=?`
const UPDATE_TXN = `UPDATE Transactions SET status = 'successful' WHERE id=?;`
const CREATE_TXN = `
	INSERT INTO TRANSACTIONS (amount,type,status,user_id,title) VALUES (?,?,'failed',?,?);
`
const GET_ALL_USER_TXN = `SELECT id,amount,type,status,title FROM Transactions WHERE user_id = ?`

func (d *database) DeductBalance(amount float64, userId int, title string) (err error) {
	defer func() {
		if err != nil {
			d.logger.Error("Failed to deduct amount ", err)
		} else {
			d.logger.Infof("successfully deducted %f from user %d", amount, userId)
		}

	}()
	create, err := d.db.Prepare(CREATE_TXN)
	if err != nil {
		return
	}
	resp, err := create.Exec(amount, TransactionTypeDeducted, userId, title)
	if err != nil {
		return
	}
	txnId, err := resp.LastInsertId()
	if err != nil {
		return
	}
	ctx := context.Background()
	txn, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		return
	}
	_, err = txn.ExecContext(ctx, UPDATE_BALANCE, amount, userId)
	if err != nil {
		txn.Rollback()
		return
	}
	_, err = txn.ExecContext(ctx, UPDATE_TXN, txnId)
	if err != nil {
		txn.Rollback()
		return
	}
	err = txn.Commit()
	return

}

func (d *database) GetAllTransactions(user int) (txn []Txn, err error) {
	txns, err := d.db.Prepare(GET_ALL_USER_TXN)
	if err != nil {
		return
	}
	rows, err := txns.Query(user)
	defer rows.Close()
	for rows.Next() {
		t := Txn{}
		if err := rows.Scan(&t.Id, &t.Amount, &t.Type, &t.Status, &t.Title); err != nil {
			d.logger.Error(err)
			continue
		}
		txn = append(txn, t)
	}
	return
}
