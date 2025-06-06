package repository

import (
	"context"
	"log/slog"

	sq "github.com/Masterminds/squirrel"
	"github.com/brkss/dextrace-server/internal/adapter/storage/postgres"
	"github.com/brkss/dextrace-server/internal/core/domain"
	"github.com/jackc/pgx/v5"
)

/*
	UserRepository implement port.UserRepository
	and provide access to the postgres database.
*/

type UserRepository struct {
	DB *postgres.DB
}

func NewUserRepository(db *postgres.DB) (*UserRepository) {
	return &UserRepository{	
		db,
	}
}

func (ur *UserRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	query := ur.DB.QueryBuilder.Insert("users").
		Columns("id", "name", "email", "password").
		Values(user.ID, user.Name, user.Email, user.Password).
		Suffix("RETURNING *")
	sql, args, err := query.ToSql()
	if err != nil {
		slog.Error("Failed to build SQL query", "error", err)
		return nil, domain.ErrInternal
	} 

	err = ur.DB.QueryRow(ctx, sql, args...).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	
	if err != nil {
		if errCode := ur.DB.ErrorCode(err); errCode == "23505" {
			return nil, domain.ErrConflictingData
		}
		slog.Error("Database error while creating user", "error", err)
		return nil, domain.ErrInternal
	}

	return user, nil
}

func (ur *UserRepository) GetUserById(ctx context.Context, id string) (*domain.User, error) {

	var user domain.User

	query := ur.DB.QueryBuilder.Select("*").From("users").Where(sq.Eq{"id": id}).Limit(1)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	err = ur.DB.QueryRow(ctx, sql, args...).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, domain.ErrNoDataFound
		}
		return nil, err;
	}

	return &user, nil
}

func (ur *UserRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {

	var user domain.User

	query := ur.DB.QueryBuilder.Select("*").From("users").Where(sq.Eq{"email": email}).Limit(1)
	sql, args, err := query.ToSql()

	if err != nil {
		return nil, err;
	}

	err = ur.DB.QueryRow(ctx, sql, args...).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, domain.ErrNoDataFound
		}
		return nil, err
	}

	return &user, nil
}

