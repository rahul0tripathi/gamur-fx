package database

import (
	"crypto/md5"
	"encoding/hex"
)

const INITIAL_BALANCE = 100.00
const GetUser = `SELECT id,balance,username FROM Users WHERE id=?`
const UpdateUserBalance = `UPDATE Users SET balance = ? WHERE id=?`
const CreateUser = `INSERT INTO Users(username,balance,password) VALUES(?,?,?)`
const GetUserPassword = `SELECT password FROM Users WHERE id=?`

type User struct {
	Id         int     `json:"id"`
	Username   string  `json:"username"`
	ProfileImg string  `json:"profile_img"`
	Balance    float64 `json:"balance"`
}

func (d *database) GetUser(id int) (user User, err error) {
	stmt, err := d.db.Prepare(GetUser)
	if err != nil {
		d.logger.Error(err)
		return user, err
	}
	result := stmt.QueryRow(id)
	if err = result.Scan(&user.Id, &user.Balance, &user.Username); err != nil {
		d.logger.Error(err)
	}
	return
}

func (d *database) CreateUser(username string, password string) (err error) {
	stmt, err := d.db.Prepare(CreateUser)
	if err != nil {
		return err
	}
	h := md5.Sum([]byte(password))
	_, err = stmt.Exec(username, INITIAL_BALANCE, hex.EncodeToString(h[:]))
	return
}

func (d *database) VerifyUserPassword(user int, password string) (verified bool) {
	passQuery, err := d.db.Prepare(GetUserPassword)
	if err != nil {
		d.logger.Error(err)
		return
	}
	pass := passQuery.QueryRow(user)
	dbPassword := ""
	err = pass.Scan(&dbPassword)
	if err != nil {
		d.logger.Error(err)
		return
	}
	return dbPassword == password
}
