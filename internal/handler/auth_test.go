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

			name: "Kerzhakov",
			input: models.User{
				PhoneNumber: 89991259178,
				FIO:         "Kerzhakov Nikita Alexeevich",
			},
			mockReturn:   "jwt_example",
			mockError:    nil,
			expectedCode: 200,
			expectedBody: `{"status":"success", "token": "jwt_example"}`,
		},
		{
			name: "Losevsky",
			input: models.User{
				PhoneNumber: 79614632626,
				FIO:         "Losevsky Ivan Zaharovich",
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
