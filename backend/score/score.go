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

type GameEnum string

const (
	Mult GameEnum = "mult"
)

type ScoreDB struct {
	ID         *int64          `db:"id" json:"id,omitempty"`
	Username   string          `db:"username" json:"username"`
	Score      int             `db:"score" json:"score"`
	Difficulty DifficultyLevel `db:"difficulty" json:"difficulty"`
	Game       GameEnum        `db:"game"`
	Date       string          `db:"created_at" goqu:"skipinsert"`
}

func (s Score) Convert(dest *ScoreDB) {
	dest.Username = s.Username
	dest.Score = s.Score
	dest.Date = s.Date.String()
	dest.Difficulty = s.Difficulty
	dest.Game = s.Game
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
	ID         int64           `json:"id,omitempty"`
	Username   string          `json:"username"`
	Score      int             `json:"score"`
	Difficulty DifficultyLevel `json:"difficulty"`
	Game       GameEnum        `json:"game"`
	Date       time.Time       `json:"created_at,omitempty"`
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

func (s *ScoreService) GetScore(game GameEnum, difficulty DifficultyLevel) []Score {
	var (
		resDB []*ScoreDB
		res   []Score
	)
	query := s.Data.Conn.Select(
		goqu.C("id"),
		goqu.C("username"),
		goqu.C("score"),
		goqu.C("difficulty"),
		goqu.C("created_at"),
	).From("rank").Order(goqu.C("score").Desc(), goqu.C("created_at").Desc()).
		Where(goqu.And(goqu.C("difficulty").Eq(difficulty), goqu.C("game").Eq(game))).Limit(10)

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
	var dbVal ScoreDB
	sc.Convert(&dbVal)
	s.Logger.DebugLogger.WithFields(logrus.Fields{
		"from_front": sc,
		"to_db":      dbVal,
	}).Info("trying to insert a new score")
	query := s.Data.Conn.Insert("rank").Prepared(true).Rows(dbVal).Executor()

	if _, err := query.Exec(); err != nil {
		s.Logger.ErrLogger.WithError(err).Error("failed to add rank in database")
		return err
	} else {
		return nil
	}
}
