package test

import (
	"context"
	"regexp"
	"testing"
	"time"

	"chekisvc/internal/domain/entity"
	"chekisvc/internal/infrastructure/database"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := database.NewUserRepository(db)

	user := &entity.UserRegisterRequest{
		Name:     "test",
		Email:    "test@example.com",
		Password: "password",
	}

	query := `INSERT INTO users (name, email, password, created_at, updated_at) VALUES ($1, $2, $3, NOW(), NOW())`
	mock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(user.Name, user.Email, user.Password).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Create(context.Background(), user)
	assert.NoError(t, err)
}

func TestUserRepository_GetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := database.NewUserRepository(db)

	rows := sqlmock.NewRows([]string{"id", "name", "email", "created_at", "updated_at"}).
		AddRow(1, "test", "test@example.com", time.Now(), time.Now())

	query := `SELECT id, name, email, created_at, updated_at FROM users WHERE id = $1`
	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(1).
		WillReturnRows(rows)

	user, err := repo.GetByID(context.Background(), 1)
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestUserRepository_GetByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := database.NewUserRepository(db)

	rows := sqlmock.NewRows([]string{"id", "name", "email", "password", "created_at", "updated_at"}).
		AddRow(1, "test", "test@example.com", "password", time.Now(), time.Now())

	query := `SELECT id, name, email, password, created_at, updated_at FROM users WHERE email = $1`
	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs("test@example.com").
		WillReturnRows(rows)

	user, err := repo.GetByEmail(context.Background(), "test@example.com")
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestUserRepository_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := database.NewUserRepository(db)

	user := &entity.UpdateUserRequest{
		Name:  "test updated",
		Email: "testupdated@example.com",
	}

	query := `UPDATE users SET name = $1, email = $2, updated_at = NOW() WHERE id = $3`
	mock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(user.Name, user.Email, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Update(context.Background(), user, 1)
	assert.NoError(t, err)
}

func TestUserRepository_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := database.NewUserRepository(db)

	query := `DELETE FROM users WHERE id = $1`
	mock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Delete(context.Background(), 1)
	assert.NoError(t, err)
}

func TestUserRepository_List(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := database.NewUserRepository(db)

	rows := sqlmock.NewRows([]string{"id", "name", "email", "created_at", "updated_at"}).
		AddRow(1, "test1", "test1@example.com", time.Now(), time.Now()).
		AddRow(2, "test2", "test2@example.com", time.Now(), time.Now())

	query := `SELECT id, name, email, created_at, updated_at FROM users LIMIT $1 OFFSET $2`
	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(10, 0).
		WillReturnRows(rows)

	users, err := repo.List(context.Background(), 10, 0)
	assert.NoError(t, err)
	assert.Len(t, users, 2)
}