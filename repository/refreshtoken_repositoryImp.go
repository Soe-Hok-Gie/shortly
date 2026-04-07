package repository

import (
	"context"
	"database/sql"
	"log"
	"shortly/model/domain"
)

type refreshtokenRepositoryImp struct {
	DB *sql.DB
}

func NewRefreshTokenRepository(DB *sql.DB) RefreshTokenRepository {
	return &refreshtokenRepositoryImp{DB: DB}

}

func (repository *refreshtokenRepositoryImp) Save(ctx context.Context, token *domain.RefreshToken) error {

	script := "INSERT INTO refresh_tokens (user_id, token_hash, expires_at) Values (?,?,?)"

	_, err := repository.DB.ExecContext(ctx, script, token.UserID, token.TokenHash, token.ExpiresAt)
	return err
}

func (repository *refreshtokenRepositoryImp) FindByHash(ctx context.Context, hash string) (*domain.RefreshToken, error) {

	script := "SELECT user_id, token_hash, expires_at FROM refresh_tokens WHERE token_hash = ? "
	row := repository.DB.QueryRowContext(ctx, script, hash)

	var token domain.RefreshToken
	err := row.Scan(
		&token.UserID,
		&token.TokenHash,
		&token.ExpiresAt,
	)
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func (repository *refreshtokenRepositoryImp) DeleteByHash(ctx context.Context, token string) error {
	script := "DELETE FROM refresh_tokens WHERE token_hash = ?"
	result, err := repository.DB.ExecContext(ctx, script, token)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	log.Println("rows deleted:", rows)

	return nil
}
