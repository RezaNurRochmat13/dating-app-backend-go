package tests

import (
	"bytes"
	"dating-app-backend/config"
	"dating-app-backend/entities"
	"dating-app-backend/service"
	"dating-app-backend/utils"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	// Use SQLite for in-memory test DB
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	// Auto-migrate the schema for testing
	db.AutoMigrate(&entities.User{}, &entities.Profile{}, &entities.Matches{}, &entities.Message{}, &entities.Like{})
	return db
}

// Mock Echo context setup
func setupEcho(method, path string, body []byte) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(method, path, bytes.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return c, rec
}

// Test user registration
func TestRegisterUser(t *testing.T) {
	DB := setupTestDB()
	config.DB = DB // Assign mock DB to global config

	// Insert a test user
	password := "password123"
	hashedPassword, _ := utils.HashPassword(password)

	// Test payload
	payload := &entities.User{
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Password: hashedPassword,
	}
	body, _ := json.Marshal(payload)

	// Setup Echo context
	c, rec := setupEcho(http.MethodPost, "/register", body)

	// Call the handler
	err := service.UserAuthRegister(payload, c)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)

	// Validate database insertion
	var user entities.User
	DB.First(&user)
	assert.NotEmpty(t, user.Password) // Ensure password is hashed
}

// Test user login
func TestLogin(t *testing.T) {
	DB := setupTestDB()
	config.DB = DB // Assign mock DB to global config

	// Insert a test user
	password := "password123"
	hashedPassword, _ := utils.HashPassword(password)
	user := &entities.User{
		Name:     "Jane Doe",
		Email:    "jane.doe@example.com",
		Password: hashedPassword,
	}
	DB.Create(&user)

	// Test payload
	payload := map[string]interface{}{
		"email":    user.Email,
		"password": password,
	}
	body, _ := json.Marshal(payload)

	// Setup Echo context
	c, rec := setupEcho(http.MethodPost, "/login", body)

	// Call the handler
	_, err := service.UserAuthLogin(user, c)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	// Validate JWT response
	response := make(map[string]interface{})
	json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NotEmpty(t, response["token"]) // Ensure token is present
}

