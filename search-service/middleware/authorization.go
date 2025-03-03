package middleware

import (
	"context"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strings"

	"search-service/models"
	"search-service/utils"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/golang-jwt/jwt/v5"
)

type CognitoClient struct {
	UserPoolID     string
	Region         string
	AdminOnlyPaths []models.AdminPath
	CognitoClient  *cognitoidentityprovider.Client
}

func (m *CognitoClient) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			utils.WriteErrorResponse(w, http.StatusUnauthorized, "Authorization header required", nil)
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		claims, err := m.validateToken(tokenString)

		if err != nil {
			utils.WriteErrorResponse(w, http.StatusUnauthorized, "Invalid Token", nil)
			return
		}

		userRole, ok := claims.CustomAttributes["role"].(string)

		if !ok {
			userRole = "user" // Default to user
		}

		currentPath := r.URL.Path
		currentMethod := r.Method

		requiresAdmin := false

		for _, adminPath := range m.AdminOnlyPaths {
			if strings.HasPrefix(currentPath, adminPath.Path) && (adminPath.Method == "*" || adminPath.Method == currentMethod) {
				requiresAdmin = true
				break
			}
		}

		if requiresAdmin && userRole != "admin" {
			utils.WriteErrorResponse(w, http.StatusForbidden, "Unauthorized Access", nil)
			return
		}

		ctx := context.WithValue(r.Context(), "claims", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (m *CognitoClient) validateToken(tokenString string) (*models.Claims, error) {
	// Fetch the JWKS from the Cognito User Pool
	jwksURL := fmt.Sprintf("https://cognito-idp.%s.amazonaws.com/%s/.well-known/jwks.json", m.Region, m.UserPoolID)
	resp, err := http.Get(jwksURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch JWKS: %v", err)
	}
	defer resp.Body.Close()

	var jwks struct {
		Keys []struct {
			Kid string `json:"kid"`
			N   string `json:"n"`
			E   string `json:"e"`
		} `json:"keys"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&jwks); err != nil {
		return nil, fmt.Errorf("failed to decode JWKS: %v", err)
	}

	claims := &models.Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		kid, ok := token.Header["kid"].(string)
		if !ok {
			return nil, fmt.Errorf("kid not found in token header")
		}

		for _, key := range jwks.Keys {
			if key.Kid == kid {
				nBytes, err := base64.RawURLEncoding.DecodeString(key.N)
				if err != nil {
					return nil, fmt.Errorf("failed to decode modulus: %v", err)
				}

				eBytes, err := base64.RawURLEncoding.DecodeString(key.E)
				if err != nil {
					return nil, fmt.Errorf("failed to decode exponent: %v", err)
				}

				n := new(big.Int).SetBytes(nBytes)
				e := new(big.Int).SetBytes(eBytes).Int64()

				return &rsa.PublicKey{
					N: n,
					E: int(e),
				}, nil
			}
		}

		return nil, fmt.Errorf("matching key not found")
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
