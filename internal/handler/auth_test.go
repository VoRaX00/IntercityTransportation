package handler

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"kursachDB/internal/domain/models"
	"kursachDB/internal/handler/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuth_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockAuthService := new(mocks.Auth)

	h := &Handler{
		services: &Service{Auth: mockAuthService},
	}

	cases := []struct {
		name         string
		input        models.User
		mockReturn   interface{}
		mockError    error
		expectedCode int
		expectedBody string
	}{
		{
			name: "Русское ФИО",
			input: models.User{
				PhoneNumber: 89991259178,
				FIO:         "Кержаков Никита Алексеевич",
			},
			mockReturn:   "jwt_example",
			mockError:    nil,
			expectedCode: 200,
			expectedBody: `{"status":"success", "token": "jwt_example"}`,
		},
		{
			name: "English full name",
			input: models.User{
				PhoneNumber: 79614632626,
				FIO:         "Losevskiy Ivan Zaharovich",
			},
			mockReturn:   "jwt_example",
			mockError:    nil,
			expectedCode: 200,
			expectedBody: `{"status":"success", "token": "jwt_example"}`,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			mockAuthService.ExpectedCalls = nil
			mockAuthService.On("Login", c.input).Return(c.mockReturn, c.mockError)

			reqBody, _ := json.Marshal(c.input)
			req, _ := http.NewRequest(http.MethodPost, "https://localhost:7080/auth/login", bytes.NewBuffer(reqBody))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = req

			h.Login(ctx)

			assert.Equal(t, c.expectedCode, w.Code)
			assert.JSONEq(t, c.expectedBody, w.Body.String())

			mockAuthService.AssertExpectations(t)
		})

	}
}

func TestAuth_Fail(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockAuthService := new(mocks.Auth)

	h := &Handler{
		services: &Service{Auth: mockAuthService},
	}

	cases := []struct {
		name         string
		input        models.User
		expectedCode int
		expectedBody string
	}{
		{
			name: "Неверное ФИО с неверным количеством заглавных букв",
			input: models.User{
				PhoneNumber: 89991259178,
				FIO:         "КеРжаков Никита Алексеевич",
			},
			expectedCode: http.StatusBadRequest,
			expectedBody: `{"message":"invalid arguments"}`,
		},
		{
			name: "Номер телефона меньше необходимого",
			input: models.User{
				PhoneNumber: 7961463262,
				FIO:         "Losevskiy Ivan Zaharovich",
			},
			expectedCode: http.StatusBadRequest,
			expectedBody: `{"message":"invalid arguments"}`,
		},
		{
			name: "Номер телефона больше необходимого",
			input: models.User{
				PhoneNumber: 8999125917,
				FIO:         "Losevskiy Ivan Zaharovich",
			},
			expectedCode: http.StatusBadRequest,
			expectedBody: `{"message":"invalid arguments"}`,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			mockAuthService.ExpectedCalls = nil

			reqBody, _ := json.Marshal(c.input)
			req, _ := http.NewRequest(http.MethodPost, "/auth/login", bytes.NewBuffer(reqBody))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = req

			h.Login(ctx)

			assert.Equal(t, c.expectedCode, w.Code)
			assert.JSONEq(t, c.expectedBody, w.Body.String())
		})
	}
}
