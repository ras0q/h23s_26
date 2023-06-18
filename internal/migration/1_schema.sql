-- +goose Up
CREATE TABLE IF NOT EXISTS users (
	id VARCHAR(36) NOT NULL,
	PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS missions (
	id CHAR(36) NOT NULL,
	name VARCHAR(255) NOT NULL,
	description VARCHAR(255) NOT NULL,
	creator_id VARCHAR(36) NOT NULL,
	PRIMARY KEY (id),
	FOREIGN KEY (creator_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS user_mission_relations (
	id CHAR(36) NOT NULL,
	user_id VARCHAR(36) NOT NULL,
	mission_id CHAR(36) NOT NULL,
	PRIMARY KEY (id),
	FOREIGN KEY (user_id) REFERENCES users(id),
	FOREIGN KEY (mission_id) REFERENCES missions(id),
	UNIQUE mission_id_user_id (mission_id, user_id)
);

-- +goose Down
DROP TABLE IF EXISTS user_mission_relations;
DROP TABLE IF EXISTS missions;
DROP TABLE IF EXISTS users;
