package db

import (
	"database/sql"
	_ "embed"
	"mult-game/backend/conf"
	"mult-game/backend/logger"
	"path"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/sqlite3"
	_ "modernc.org/sqlite"
)

var conn *goqu.Database

//go:embed schema.sql
var schema string

type Db struct {
	Conn   *goqu.Database
	logger *logger.Logger
}

func GetDbConnection(logger *logger.Logger) *Db {
	var err error
	dialect := goqu.Dialect("sqlite3")
	var confFolder *string

	confFolder, err = conf.GetConfFolder()
	if err != nil {
		logger.ErrLogger.WithError(err).Error("failed to create config folder")
	}

	if conn == nil {
		db, err := sql.Open("sqlite", path.Join(*confFolder, "score.db"))
		if err != nil {
			logger.ErrLogger.WithError(err).Error("failed to open db")
		}

		conn = dialect.DB(db)
	}

	res := &Db{
		Conn:   conn,
		logger: logger,
	}

	res.createTable()

	return res
}

func (db *Db) createTable() {
	_, err := db.Conn.Exec(schema)
	if err != nil {
		db.logger.ErrLogger.WithError(err).Error("failed to create table rank in database")
	}
}
