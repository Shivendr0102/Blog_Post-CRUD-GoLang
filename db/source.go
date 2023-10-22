package db

import (
	"22nd_Oct_Antino/config"
	"22nd_Oct_Antino/db/blog"
	"database/sql"
	"fmt"
)

type Source struct {
	Blog *blog.Dao
}

func NewSource(config *config.Config) (*Source, error) {
	source := Source{}

	db, err := source.connect(config)
	if err != nil {
		return nil, err
	}

	err = source.prepare(db)
	if err != nil {
		return nil, err
	}

	return &source, err
}

func (s *Source) connect(config *config.Config) (*sql.DB, error) {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;Connection Timeout=600",
		config.SqlServer.URL, config.SqlServer.Username, config.SqlServer.Password, config.SqlServer.Port, config.SqlServer.DBName)
	return sql.Open("mssql", connString)
}

func (s *Source) prepare(db *sql.DB) error {
	var err error

	s.Blog, err = blog.NewDao(db)
	if err != nil {
		return err
	}

	return nil
}
