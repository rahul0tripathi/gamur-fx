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

const UpdateBalance = `UPDATE Users SET balance = balance - ? WHERE id=?`
const UpdateTxn = `UPDATE Transactions SET status = 'successful' WHERE id=?;`
const CreateTxn = `
	INSERT INTO TRANSACTIONS (amount,type,status,user_id,title) VALUES (?,?,'failed',?,?);
`
const GetAllUserTxn = `SELECT id,amount,type,status,title FROM Transactions WHERE user_id = ?`

func (d *database) DeductBalance(amount float64, userId int, title string) (err error) {
	defer func() {
		if err != nil {
			d.logger.Error("Failed to deduct amount ", err)
		} else {
			d.logger.Infof("successfully deducted %f from user %d", amount, userId)
		}

	}()
	create, err := d.db.Prepare(CreateTxn)
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
	_, err = txn.ExecContext(ctx, UpdateBalance, amount, userId)
	if err != nil {
		txn.Rollback()
		return
	}
	_, err = txn.ExecContext(ctx, UpdateTxn, txnId)
	if err != nil {
		txn.Rollback()
		return
	}
	err = txn.Commit()
	return

}

func (d *database) GetAllTransactions(user int) (txn []Txn, err error) {
	txns, err := d.db.Prepare(GetAllUserTxn)
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
