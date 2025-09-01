package pkg

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "log"
)

func InitDB(filepath string) *sql.DB {
    db, err := sql.Open("sqlite3", filepath)
    if err != nil {
        log.Fatal(err)
    }

    createSpots := `
    CREATE TABLE IF NOT EXISTS spots (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        photo_url TEXT,
        name TEXT,
        location_id INTEGER,
        difficulty INTEGER,
        surf_breaks TEXT,
        season_start TEXT,
        season_end TEXT,
        FOREIGN KEY(location_id) REFERENCES locations(id)
    );`
    _, err = db.Exec(createSpots)
    if err != nil {
        log.Fatal(err)
    }

    createLocations := `
    CREATE TABLE IF NOT EXISTS locations (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        country TEXT,
        city TEXT,
        lat INTEGER,
        long INTEGER
    );`
    _, err = db.Exec(createLocations)
    if err != nil {
        log.Fatal(err)
    }

    return db
}
