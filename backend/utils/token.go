package utils

import (
	"errors"
	"fmt"
	"github.com/KokoiRuby/rbac-based-management-system/backend/config/runtime"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/model"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type ClaimMeta struct {
	UserID   uint   `json:"userID"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	//RoleList []uint `json:"roleList"`
}

type CustomClaims struct {
	ClaimMeta
	jwt.RegisteredClaims
}

// TODO: function-ize jwt.NewWithClaims?!

func CreateAccessToken(user *model.User, cfg runtime.JWT) (accessToken string, err error) {
	meta := ClaimMeta{
		UserID:   user.ID,
		Username: user.Username,
		Email:    user.Email,
		//RoleList: user.RoleList,
	}
	claims := CustomClaims{
		ClaimMeta: meta,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(cfg.Expire) * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    cfg.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err = token.SignedString([]byte(cfg.SecretKey))
	if err != nil {
		return "", err
	}
	return
}

func CreateRefreshToken(user *model.User, cfg runtime.JWT) (accessToken string, err error) {
	meta := ClaimMeta{
		UserID: user.ID,
		//RoleList: user.RoleList,
	}
	claims := CustomClaims{
		ClaimMeta: meta,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(cfg.RefreshExpire) * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    cfg.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err = token.SignedString([]byte(cfg.SecretKey))
	if err != nil {
		return "", err
	}
	return
}

func CreateConfirmToken(req *model.SignupConfirmRequest, cfg runtime.JWT) (confirmToken string, err error) {
	meta := ClaimMeta{
		Email:    req.Email,
		Password: req.HashedPassword,
	}
	claims := CustomClaims{
		ClaimMeta: meta,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(cfg.ConfirmExpire) * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    cfg.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	confirmToken, err = token.SignedString([]byte(cfg.SecretKey))
	if err != nil {
		return "", err
	}
	return
}

func ParseToken(tokenString string, cfg runtime.JWT) (claims *CustomClaims, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(cfg.SecretKey), nil
	})
	fmt.Println(token, err)
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// TODO: generic-ize extract field from token via reflection

func ExtractIDFromToken(tokenString string, cfg runtime.JWT) (uint, error) {
	claims, err := ParseToken(tokenString, cfg)
	if err != nil {
		return 0, err
	}
	return claims.UserID, nil
}

func ExtractExpireAtFromToken(tokenString string, cfg runtime.JWT) (*jwt.NumericDate, error) {
	claims, err := ParseToken(tokenString, cfg)
	if err != nil {
		return nil, err
	}
	return claims.ExpiresAt, nil
}

func ExtractCredFromToken(tokenString string, cfg runtime.JWT) (*model.SignupRequest, error) {
	claims, err := ParseToken(tokenString, cfg)
	if err != nil {
		return nil, err
	}
	return &model.SignupRequest{
		Email:    claims.Email,
		Password: claims.Password,
	}, nil
}
