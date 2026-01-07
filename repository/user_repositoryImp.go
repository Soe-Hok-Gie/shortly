package repository

import (
	"context"
	"database/sql"
	"errors"
	"shortly/model/domain"

	"github.com/go-sql-driver/mysql"
)

type userRepositoryImp struct {
	DB *sql.DB
}

func (repository *userRepositoryImp) Save(ctx context.Context, user domain.User) (domain.User, error) {
	script := "INSERT INTO users (username,password) VALUES (?,?)"
	result, err := repository.DB.ExecContext(ctx, script, user.Username, user.Password)
	if err != nil {
		if isDuplicateKeyError(err) {
			return user, errors.New("username already exists")
		}
		return user, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return user, err
	}

	user.Id = id
	return user, nil
}

// Helper untuk cek duplicate key MySQL
func isDuplicateKeyError(err error) bool {
	if me, ok := err.(*mysql.MySQLError); ok {
		return me.Number == 1062
	}
	return false
}
