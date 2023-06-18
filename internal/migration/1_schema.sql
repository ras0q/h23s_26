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

INSERT INTO users (id) VALUES
	("user1"),
	("user2");
INSERT INTO missions (id, name, description, creator_id) VALUES
	("fdc91790-41a1-4e3a-a5ba-af4600efedec", "mission1", "mission1", "user1"),
	("101a5594-7ead-4255-8084-b245b69c0f0d", "mission2", "mission2", "user2");
INSERT INTO user_mission_relations (id, user_id, mission_id) VALUES
	("8b420cdf-a8e1-4317-b9a2-1fc43059a66f", "user1", "fdc91790-41a1-4e3a-a5ba-af4600efedec"),
	("b1b2b3b4-b5b6-b7b8-b9ba-bcbdbebf0000", "user2", "fdc91790-41a1-4e3a-a5ba-af4600efedec"),
	("b1b2b3b4-b5b6-b7b8-b9ba-bcbdbebf0001", "user1", "101a5594-7ead-4255-8084-b245b69c0f0d");

-- +goose Down
DROP TABLE IF EXISTS user_mission_relations;
DROP TABLE IF EXISTS missions;
DROP TABLE IF EXISTS users;
