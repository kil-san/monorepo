package connection

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kil-san/micro-serv/pkg/model"
)

func NewSqlDbConnection(config model.DbConfig) (*sql.DB, error) {
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.User, config.Password, config.Host, config.Port, config.DbName)
	db, err := sql.Open("mysql", conn)
	if err != nil {
		return db, err
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	createNoteTableSQL := `CREATE TABLE IF NOT EXISTS note (
		id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,		
		title TEXT,
		status TEXT
	  );`

	stmt, err := db.Prepare(createNoteTableSQL)
	if err != nil {
		return nil, err
	}
	stmt.Exec()

	return db, nil
}
