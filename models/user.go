package models

import (
	"database/sql"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/younesious/events/db"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
}

var jwtKey = []byte("younesious_key")

type Claims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (u *User) CreateUser() error {
	db := db.GetDB()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.CreatedAt = time.Now()
	query := `INSERT INTO users (username, email, password, created_at) VALUES (?, ?, ?, ?)`
	res, err := db.Exec(query, u.Username, u.Email, hashedPassword, u.CreatedAt)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = id

	return nil
}

func AuthenticateUser(username, password string) (*User, error) {
	db := db.GetDB()
	var user User
	query := `SELECT id, username, password, created_at FROM users WHERE username = ?`
	err := db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("incorrect password")
	}

	return &user, nil
}

func GenerateJWT(userID int64, username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, errors.New("invalid token signature")
		}
		return nil, errors.New("invalid token")
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
