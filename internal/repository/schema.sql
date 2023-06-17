CREATE TABLE IF NOT EXISTS users (
	id VARCHAR(36) NOT NULL,
	PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS missions (
	id CHAR(36) NOT NULL,
	name VARCHAR(255) NOT NULL,
	description VARCHAR(255) NOT NULL,
	creator_id VARCHAR(36) NOT NULL,
	created_at DATETIME NOT NULL,
	updated_at DATETIME NOT NULL,
	deleted_at DATETIME,
	PRIMARY KEY (id),
	FOREIGN KEY (creator_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS user_mission_relations (
	id CHAR(36) NOT NULL,
	user_id VARCHAR(36) NOT NULL,
	mission_id CHAR(36) NOT NULL,
	PRIMARY KEY (id),
	FOREIGN KEY (user_id) REFERENCES users(id),
	FOREIGN KEY (mission_id) REFERENCES missions(id)
);
