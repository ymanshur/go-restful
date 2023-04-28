package repository

import (
	"go-restful/database"
	"go-restful/internal/dto"
	"go-restful/internal/model"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestUserSave(t *testing.T) {
	// Setup
	db, mock := database.NewMock()
	r := NewUser(db)
	payload := dto.CreateUserRequest{
		User: model.User{
			Name:     "Fityah Salamah",
			Email:    "fityah.salamah@gmail.com",
			Password: "1234",
		},
	}
	// mock.
	// 	ExpectQuery(regexp.QuoteMeta(`
	// 		INSERT into "users" ("name","email","password")
	// 		VALUES (?,?,?,?)
	// 	`)).
	// 	WillReturnRows()
	mock.ExpectBegin()
	mock.
		// ExpectExec(regexp.QuoteMeta(`
		// 	INSERT INTO "users" ("created_at","updated_at","name","email","password")
		// 	VALUES (?,?,?,?,?)
		// `)).
		ExpectExec("INSERT INTO users").
		WithArgs(time.Now(), time.Now(), payload.Name, payload.Email, payload.Password).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Assertions
	assert.NoError(t, r.Save(&payload))
}
