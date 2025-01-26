package middleware

import (
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Ambil token dari header Authorization
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Authorization header missing"})
		}

		// Validasi token
		secret := os.Getenv("JWT_SECRET")
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid or expired token"})
		}

		// Ekstrak klaim token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid token claims"})
		}

		// Simpan klaim di konteks
		c.Set("user", claims)
		return next(c)
	}
}

// func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		// Ambil token langsung dari header Authorization
// 		tokenString := c.Request().Header.Get("Authorization")
// 		if tokenString == "" {
// 			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Authorization header missing"})
// 		}

// 		// Validasi token
// 		secret := os.Getenv("JWT_SECRET")
// 		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
// 			// Pastikan token menggunakan metode signing yang benar
// 			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
// 				return nil, jwt.ErrSignatureInvalid
// 			}
// 			return []byte(secret), nil
// 		})

// 		if err != nil || !token.Valid {
// 			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid or expired token"})
// 		}

// 		// Ekstrak data dari klaim token
// 		claims, ok := token.Claims.(jwt.MapClaims)
// 		if !ok {
// 			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid token claims"})
// 		}

// 		// Simpan klaim ke dalam konteks untuk digunakan di handler
// 		c.Set("user", claims)

// 		// Lanjutkan ke handler berikutnya
// 		return next(c)
// 	}
// }

// package middleware

// import (
// 	"net/http"
// 	"os"
// 	"strings"

// 	"github.com/golang-jwt/jwt/v5"
// 	"github.com/labstack/echo/v4"
// )

// func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		// Ambil token dari header Authorization
// 		authHeader := c.Request().Header.Get("Authorization")
// 		if authHeader == "" {
// 			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Authorization header missing"})
// 		}

// 		// Pastikan token diawali dengan "Bearer"
// 		tokenParts := strings.Split(authHeader, " ")
// 		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
// 			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid authorization format"})
// 		}

// 		tokenString := tokenParts[1]

// 		// Validasi token
// 		secret := os.Getenv("JWT_SECRET")
// 		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
// 			// Pastikan token menggunakan metode signing yang benar
// 			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
// 				return nil, jwt.ErrSignatureInvalid
// 			}
// 			return []byte(secret), nil
// 		})

// 		if err != nil || !token.Valid {
// 			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid or expired token"})
// 		}

// 		// Ekstrak data dari klaim token
// 		claims, ok := token.Claims.(jwt.MapClaims)
// 		if !ok {
// 			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid token claims"})
// 		}

// 		// Simpan klaim ke dalam konteks untuk digunakan di handler
// 		c.Set("user", claims)

// 		// Lanjutkan ke handler berikutnya
// 		return next(c)
// 	}
// }
