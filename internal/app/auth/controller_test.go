package auth

import (
	"bytes"
	"encoding/json"
	"go-restful/internal/app/dto"
	mock_repository "go-restful/internal/mocks/repository"
	"go-restful/internal/model"
	"go-restful/pkg/util/validator"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	// db, _   = database.NewMock()
	// r       = repository.NewUser(db)
	// c       = NewController(r)
	baseURL = "/api/auth"
)

func NewMockRepository(t *testing.T) *mock_repository.MockUser {
	ctrl := gomock.NewController(t)
	return mock_repository.NewMockUser(ctrl)
}

func TestSignUpSuccess(t *testing.T) {
	// Setup
	e := echo.New()
	e.Validator = validator.New()
	payload := dto.CreateUserRequest{
		User: model.User{
			Name:     "Fityah Salamah",
			Email:    "fityah.salamah@gmail.com",
			Password: "1234",
		},
	}
	payloadInByte, err := json.Marshal(payload)
	if err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodPost, baseURL+"/signup", bytes.NewReader(payloadInByte))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	mockR := NewMockRepository(t)
	mockR.EXPECT().Save(&payload).Return(nil)
	mockC := NewController(mockR)

	// Assertions
	if assert.NoError(t, mockC.SignUp(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestSignUpUnprocessableEntity(t *testing.T) {
	// Setup
	e := echo.New()
	e.Validator = validator.New()
	payload := dto.CreateUserRequest{
		User: model.User{
			Name:     "",
			Email:    "fityah.salamah@gmail.com",
			Password: "1234",
		},
	}
	payloadInByte, err := json.Marshal(payload)
	if err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodPost, baseURL+"/signup", bytes.NewReader(payloadInByte))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	mockR := NewMockRepository(t)
	mockC := NewController(mockR)

	// Assertions
	if assert.NoError(t, mockC.SignUp(ctx)) {
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
	}
}

func TestSignInSuccess(t *testing.T) {
	// Setup
	e := echo.New()
	e.Validator = validator.New()
	payload := dto.AuthSignInRequest{
		Email:    "fityah.salamah@gmail.com",
		Password: "1234",
	}
	payloadInByte, err := json.Marshal(payload)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(http.MethodPost, baseURL+"/signin", bytes.NewReader(payloadInByte))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	// monkey.PatchInstanceMethod(reflect.TypeOf(r), "FindByEmail",
	// 	func(userRepo *repository.UserRepo, email *string) (*model.User, error) {
	// 		user := model.User{
	// 			Name:     "Fityah Salamah",
	// 			Email:    payload.Email,
	// 			Password: payload.Password,
	// 		}
	// 		return &user, nil
	// 	})
	mockR := NewMockRepository(t)
	mockR.EXPECT().FindByEmail(&payload.Email).Return(
		&model.User{
			Name:     "Fityah Salamah",
			Email:    payload.Email,
			Password: payload.Password,
		},
		nil,
	)
	mockC := NewController(mockR)

	// Assertions
	if assert.NoError(t, mockC.SignIn(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
