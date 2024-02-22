package score

import (
	"context"
	"mult-game/backend/db"
	"mult-game/backend/logger"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

type ScoreService struct {
	logger *logger.Logger
	data   *db.Db
}

type ScoreDB struct {
	Username string `db:"username" json:"username"`
	Score    int    `db:"score" json:"score"`
	Date     string `db:"created_at" goqu:"skipinsert"`
}

func (s Score) Convert(dest *ScoreDB) {
	dest.Username = s.Username
	dest.Score = s.Score
	dest.Date = s.Date.String()
}

func (s ScoreDB) Convert(dest *Score) {
	var err error
	dest.Username = s.Username
	dest.Score = s.Score
	dest.Date, err = time.Parse("2006-01-02 15:04:05", s.Date)
	if err != nil {
		logrus.WithError(err).Error("failed to parse date")
	}
}

type Score struct {
	fx.In
	Username string    `db:"username" json:"username"`
	Score    int       `db:"score" json:"score"`
	Date     time.Time `db:"created_at" goqu:"skipinsert" json:"created_at"`
}

func NewScoreService(logger *logger.Logger, conn *db.Db) *ScoreService {
	return &ScoreService{
		logger: logger,
		data:   conn,
	}
}

func (s *ScoreService) GetScore() []Score {
	var (
		resDB []*ScoreDB
		res   []Score
	)
	query := s.data.Conn.Select(
		goqu.C("username"),
		goqu.C("score"),
		goqu.C("created_at"),
	).From("rank").Order(goqu.C("score").Desc()).Limit(10)

	err := query.ScanStructsContext(context.Background(), &resDB)
	if err != nil {
		s.logger.ErrLogger.WithError(err).Error("failed to retrieve ranks from database")
		return nil
	}

	for _, s := range resDB {
		var sc Score
		s.Convert(&sc)
		res = append(res, sc)
	}

	return res
}

func (s ScoreService) AddScore(sc Score) error {

	query := s.data.Conn.Insert("rank").Prepared(true).Rows(sc).Executor()

	if _, err := query.Exec(); err != nil {
		s.logger.ErrLogger.WithError(err).Error("failed to add rank in database")
		return err
	} else {
		return nil
	}

}
