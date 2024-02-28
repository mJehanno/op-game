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
	Logger *logger.Logger
	Data   *db.Db
}

type DifficultyLevel string

const (
	Easy   DifficultyLevel = "easy"
	Medium DifficultyLevel = "medium"
	Hard   DifficultyLevel = "hard"
)

type ScoreDB struct {
	Username   string          `db:"username" json:"username"`
	Score      int             `db:"score" json:"score"`
	Difficulty DifficultyLevel `db:"difficulty" json:"difficulty"`
	Date       string          `db:"created_at" goqu:"skipinsert"`
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
	Username   string          `db:"username" json:"username"`
	Score      int             `db:"score" json:"score"`
	Difficulty DifficultyLevel `db:"difficulty" json:"difficulty"`
	Date       time.Time       `db:"created_at" goqu:"skipinsert" json:"created_at,omitempty"`
}

type ScoreParams struct {
	fx.In
	Logger *logger.Logger
	Conn   *db.Db
}

func NewScoreService(p ScoreParams) *ScoreService {
	return &ScoreService{
		Logger: p.Logger,
		Data:   p.Conn,
	}
}

func (s *ScoreService) GetScore(difficulty DifficultyLevel) []Score {
	var (
		resDB []*ScoreDB
		res   []Score
	)
	query := s.Data.Conn.Select(
		goqu.C("username"),
		goqu.C("score"),
		goqu.C("difficulty"),
		goqu.C("created_at"),
	).From("rank").Order(goqu.C("score").Desc(), goqu.C("created_at").Desc()).
		Where(goqu.C("difficulty").Eq(difficulty)).Limit(10)

	err := query.ScanStructsContext(context.Background(), &resDB)
	if err != nil {
		s.Logger.ErrLogger.WithError(err).Error("failed to retrieve ranks from database")
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

	s.Logger.DebugLogger.Debug(sc)

	query := s.Data.Conn.Insert("rank").Prepared(true).Rows(sc).Executor()

	if _, err := query.Exec(); err != nil {
		s.Logger.ErrLogger.WithError(err).Error("failed to add rank in database")
		return err
	} else {
		return nil
	}

}
