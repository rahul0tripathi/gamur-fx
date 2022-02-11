package database

const CreateUsersSchema = `
	CREATE TABLE IF NOT EXISTS Users(
			id INT NOT NULL AUTO_INCREMENT,
			balance DECIMAL(10,2) NOT NULL,
			username VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL,
			PRIMARY KEY(id)
	);
`
const AddBalanceCheckConstraint = `
	ALTER TABLE Users 
	ADD CONSTRAINT BalanceConstraint CHECK (balance > 0.00)
`
const CreateTxnSchema = `
	CREATE TABLE IF NOT EXISTS Transactions(
			id INT NOT NULL AUTO_INCREMENT,
			amount DECIMAL(10,2) NOT NULL,
			title VARCHAR(255),
			type ENUM('added','deducted') NOT NULL,
			status ENUM('successful','failed') NOT NULL,
			user_id INT NOT NULL,
			PRIMARY KEY(id),
			FOREIGN KEY (user_id) REFERENCES Users(id)
	);`
const CreateGameSchema = `
	CREATE TABLE IF NOT EXISTS Games(
			id INT NOT NULL AUTO_INCREMENT,
			title VARCHAR(255) NOT NULL,
			assets VARCHAR(255) NOT NULL,
			PRIMARY KEY(id)
	);
`
const CreateBattleSchema = `
	CREATE TABLE IF NOT EXISTS Battle(
			id INT NOT NULL AUTO_INCREMENT,
			entry_fee DECIMAL(10,2) NOT NULL,
			game_id INT NOT NULL,
			status ENUM('pending','completed'),
			PRIMARY KEY(id),
			FOREIGN KEY (game_id) REFERENCES Games(id)
	);
`
const CreatePlayersSchema = `
	CREATE TABLE IF NOT EXISTS Players(
			battle_id INT NOT NULL,
			player INT NOT NULL,
			score INT NOT NULL,
			PRIMARY KEY(battle_id,player),
			FOREIGN KEY (battle_id) REFERENCES Battle(id),
			FOREIGN KEY (player) REFERENCES Users(id)
	);
`
const CreateLeaderBoardSchema = `
	CREATE TABLE IF NOT EXISTS Leaderboard(
			user_id INT NOT NULL,
			score INT NOT NULL,
			PRIMARY KEY(user_id),
			FOREIGN KEY (user_id) REFERENCES Users(id)
	);
`

var SchemaList = []string{CreateUsersSchema, AddBalanceCheckConstraint, CreateGameSchema, CreateTxnSchema, CreateBattleSchema, CreatePlayersSchema, CreateLeaderBoardSchema, LeaderBoardTrigger}
