package controllers

import (
	"fmt"
	"go-gorm/database"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte("secret")

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func generateToken(email string) (string, error) {
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func login(ctx *gin.Context) (string, error) {
	db := database.GetDB()

	var userLogin UserLogin

	if err := ctx.ShouldBindJSON(&userLogin); err != nil {
		return "", err
	}

	var user User
	result := db.Where("email = ?", userLogin.Email).First(&user)
	if result.Error != nil {
		return "", fmt.Errorf("Email tidak ditemukan")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLogin.Password)); err != nil {
		return "", fmt.Errorf("Password tidak valid")
	}

	// Jika email dan password valid, generate token
	token, err := generateToken(userLogin.Email)
	if err != nil {
		return "", fmt.Errorf("Gagal membuat token")
	}

	return token, nil
}

func LoginHandler(ctx *gin.Context) {
	token, err := login(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
