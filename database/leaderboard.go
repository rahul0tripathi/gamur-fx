package database

type Leaderboard struct {
	Username string `json:"username"`
	Id       int    `json:"id"`
	Score    int    `json:"score"`
}

const LIMIT = 50

const GetTopPlayers = `SELECT u.username,u.id,l.score FROM Leaderboard l LEFT JOIN Users u ON u.id=l.user_id ORDER BY l.score DESC LIMIT ?`

func (d *database) GetTopPlayers() (leaderboard []Leaderboard, err error) {
	l, err := d.db.Prepare(GetTopPlayers)
	if err != nil {
		return
	}
	rows, err := l.Query(LIMIT)
	if err != nil {
		return
	}
	for rows.Next() {
		lboard := Leaderboard{}
		if err := rows.Scan(&lboard.Username, &lboard.Id, &lboard.Score); err != nil {
			d.logger.Error(err)
			continue
		}
		leaderboard = append(leaderboard, lboard)
	}
	return
}
