DROP TABLE IF EXISTS db_version;
DROP TABLE IF EXISTS camps;
DROP TABLE IF EXISTS camp_years;

CREATE TABLE IF NOT EXISTS db_version (
	major integer,
	minor integer,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO db_version (major, minor) VALUES (1, 0);

---

CREATE TABLE IF NOT EXISTS camps (
	id integer PRIMARY KEY,
	name text NOT NULL
);

CREATE TABLE IF NOT EXISTS camp_years (
	camp_id integer NOT NULL
		REFERENCES camps (id) ON DELETE CASCADE,
	year varchar(4),
	PRIMARY KEY (camp_id, year)
);

INSERT INTO camps (name) VALUES ('New Camp');
INSERT INTO camp_years (camp_id, year) VALUES (last_insert_rowid(), strftime('%Y'));

---

