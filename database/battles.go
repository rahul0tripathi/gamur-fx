package database

type Battle struct {
	Id         int     `json:"id"`
	EntryFee   float64 `json:"entry_fee"`
	Username   string  `json:"username"`
	ProfileImg string  `json:"profile_img"`
	Score      int     `json:"score"`
}

const (
	BattleStatusPending = "pending"
	BattleStatusCompleted = "completed"
	GET_USER_BATTLES      = `
		SELECT b.id,
			   b.entry_fee,
			   u.username,
			   p.score
		FROM   players p
			   LEFT JOIN battle b
					  ON b.id = p.battle_id
			   LEFT JOIN users u
					  ON u.id = p.player
		WHERE  p.player = ?;
`
	CREATE_NEW_BATTLE    = `INSERT INTO Battle (entry_fee,game_id,status) VALUES(?,?,'pending');`
	INSERT_PLAYERS       = `INSERT INTO Players (battle_id,player,score) VALUES(?,?,0)`
	UPDATE_BATTLE_STATUS = `UPDATE Battle SET status = ? WHERE  id=?`
	UPDATE_PLAYER_SCORE  = `UPDATE Players SET score = ? WHERE player=? and battle_id=?`
)

func (d *database) NewBattle(players []int, entryFee float64, gameId int) (err error) {
	stmnt, err := d.db.Prepare(CREATE_NEW_BATTLE)
	if err != nil {
		return
	}
	battle, err := stmnt.Exec(entryFee, gameId)
	if err != nil {
		return
	}
	battleId, err := battle.LastInsertId()
	if err != nil {
		return
	}
	insertPlayers, err := d.db.Prepare(INSERT_PLAYERS)
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
	s, err := d.db.Prepare(GET_USER_BATTLES)
	if err != nil {
		return
	}
	r, err := s.Query(user)
	if err != nil {
		return
	}
	for r.Next() {
		b := Battle{}
		if err := r.Scan(&b.Id, &b.EntryFee, &b.Username, &b.Score); err != nil {
			d.logger.Error(err)
			continue
		}
		battles = append(battles, b)
	}
	return
}

func (d *database) UpdatePlayerResult(player int, score int, battle int) (err error) {
	s, err := d.db.Prepare(UPDATE_PLAYER_SCORE)
	if err != nil {
		return
	}
	_, err = s.Exec(score, player, battle)
	if err != nil {
		return
	}
	updateBattle, err := d.db.Prepare(UPDATE_BATTLE_STATUS)
	if err != nil {
		return
	}
	_, err = updateBattle.Exec(BattleStatusCompleted, battle)
	return
}
