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
	camp_id integer NOT NULL,
	year integer NOT NULL,
	PRIMARY KEY (camp_id, year),
	FOREIGN KEY (camp_id)
		REFERENCES camps (id) ON DELETE CASCADE
);

INSERT INTO camps (name) VALUES ('New Camp');
INSERT INTO camp_years (camp_id, year) VALUES (last_insert_rowid(), strftime('%Y'));

---

CREATE TABLE IF NOT EXISTS registrations (
    id integer PRIMARY KEY,
    for_camp integer NOT NULL,
    camp_year integer NOT NULL,
    created_at text,
    updated_at text,
    FOREIGN KEY (for_camp, camp_year)
        REFERENCES camp_years (camp_id, year)
        ON DELETE RESTRICT
);

CREATE TRIGGER IF NOT EXISTS registrations_set_created_at
AFTER INSERT ON registrations
FOR EACH ROW
WHEN NEW.created_at IS NULL
BEGIN
    UPDATE registrations SET created_at = strftime('%Y-%m-%d %H:%M:%f', 'now')
    WHERE id = NEW.id;
END;

CREATE TRIGGER IF NOT EXISTS registrations_update_updated_at
AFTER UPDATE ON registrations
FOR EACH ROW
BEGIN
    UPDATE registrations SET updated_at = strftime('%Y-%m-%d %H:%M:%f', 'now')
    WHERE id = OLD.id;
END;

CREATE TRIGGER IF NOT EXISTS registrations_set_updated_at
AFTER INSERT ON registrations
FOR EACH ROW
WHEN NEW.updated_at IS NULL
BEGIN
    UPDATE registrations SET updated_at = strftime('%Y-%m-%d %H:%M:%f', 'now')
    WHERE id = NEW.id;
END;


