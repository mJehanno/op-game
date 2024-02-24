package db

import (
	"database/sql"
	_ "embed"
	"fmt"
	"mult-game/backend/logger"
	"os"
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

	conf, erra := os.UserConfigDir()
	if erra != nil {
		conf, err = os.UserHomeDir()
		if err != nil {
			logger.ErrLogger.WithError(fmt.Errorf("%w and %w happened", erra, err)).Error("failed to access config or home directory")
			return nil
		}
	}

	confFolder := path.Join(conf, "mult-game")
	err = os.MkdirAll(confFolder, 0775)
	if err != nil {
		logger.ErrLogger.WithError(err).Error("failed to create folder in config or home directory")
		return nil
	}

	if conn == nil {
		db, err := sql.Open("sqlite", path.Join(confFolder, "score.db"))
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
