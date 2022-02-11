package database

const LeaderBoardTrigger = `
			CREATE DEFINER = CURRENT_USER TRIGGER gamur.Players_AFTER_UPDATE AFTER UPDATE ON Players FOR EACH ROW
			BEGIN
			IF (new.score >
			(
				   SELECT score
				   FROM   gamur.leaderboard
				   WHERE  user_id = new.player)
			OR
			NOT EXISTS
			(
				   SELECT score
				   FROM   gamur.leaderboard
				   WHERE  user_id = new.player) ) THEN
			INSERT INTO gamur.leaderboard
						(
									user_id,
									score
						)
						VALUES
						(
									new.player,
									new.score
						)
			ON duplicate KEY
			UPDATE score = new.score;
			end IF;
			END
`
