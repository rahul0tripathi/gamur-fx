package database

type Battle struct {
	Id         int     `json:"id"`
	EntryFee   float64 `json:"entry_fee"`
	Username   string  `json:"username"`
	ProfileImg string  `json:"profile_img"`
	Score      int     `json:"score"`
	Status     string  `json:"status"`
}

const (
	BattleStatusPending   = "pending"
	BattleStatusCompleted = "completed"
	GetUserBattles        = `
		SELECT b.id,
			   b.entry_fee,
			   u.username,
			   p.score,
			   b.status
		FROM   players p
			   LEFT JOIN battle b
					  ON b.id = p.battle_id
			   LEFT JOIN users u
					  ON u.id = p.player
		WHERE  p.player = ?;
`
	CreateNewBattle    = `INSERT INTO Battle (entry_fee,game_id,status) VALUES(?,?,'pending');`
	InsertPlayers      = `INSERT INTO Players (battle_id,player,score) VALUES(?,?,0)`
	UpdateBattleStatus = `UPDATE Battle SET status = ? WHERE  id=?`
	UpdatePlayerScore  = `UPDATE Players SET score = ? WHERE player=? and battle_id=?`
)

func (d *database) NewBattle(players []int, entryFee float64, gameId int) (battleId int64, err error) {
	stmnt, err := d.db.Prepare(CreateNewBattle)
	if err != nil {
		return
	}
	battle, err := stmnt.Exec(entryFee, gameId)
	if err != nil {
		return
	}
	battleId, err = battle.LastInsertId()
	if err != nil {
		return
	}
	insertPlayers, err := d.db.Prepare(InsertPlayers)
	if err != nil {
		return
	}
	for _, v := range players {
		_, err := insertPlayers.Exec(battleId, v)
		if err != nil {
			d.logger.Error(err)
		}
	}
	return
}

func (d *database) GetUserBattles(user int) (battles []Battle, err error) {
	s, err := d.db.Prepare(GetUserBattles)
	if err != nil {
		return
	}
	r, err := s.Query(user)
	if err != nil {
		return
	}
	for r.Next() {
		b := Battle{}
		if err := r.Scan(&b.Id, &b.EntryFee, &b.Username, &b.Score, &b.Status); err != nil {
			d.logger.Error(err)
			continue
		}
		battles = append(battles, b)
	}
	return
}

func (d *database) UpdatePlayerResult(player int, score int, battle int) (err error) {
	s, err := d.db.Prepare(UpdatePlayerScore)
	if err != nil {
		return
	}
	_, err = s.Exec(score, player, battle)
	if err != nil {
		return
	}
	updateBattle, err := d.db.Prepare(UpdateBattleStatus)
	if err != nil {
		return
	}
	_, err = updateBattle.Exec(BattleStatusCompleted, battle)
	return
}
