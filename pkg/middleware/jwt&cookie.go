package middleware

import (
	"errors"
	"fmt"
	"gateway/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func UserRetriveCookie(c *gin.Context) {

	valid := ValidateCookie(c)
	if valid == false {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not logged in"})
		c.Abort()
	} else {
		userId, Phone, role, err := RetriveJwtToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "cookie retriving failed"})
			c.Abort()
		} else {
			c.Set("userID", userId)
			c.Set("phoneNumber", Phone)
		}
		if role != "user" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "role mismatching"})
			c.Abort()
		}
	}
	c.Next()
}
func AdminRetriveCookie(c *gin.Context) {

	valid := ValidateCookie(c)
	if valid == false {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not logged in"})
		c.Abort()
	} else {
		userId, Phone, role, err := RetriveJwtToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "cookie retriving failed"})
			c.Abort()
		} else {
			c.Set("userID", userId)
			c.Set("phoneNumber", Phone)
		}
		if role != "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "role mismatching"})
			c.Abort()
		} else {
			c.Next()
		}
	}

}

func CreateJwtCookie(userId int, userPhone string, role string, c *gin.Context) {
	config, _ := utils.LoadConfig("./")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid": userId,
		"phone":  userPhone,
		"role":   role,
	})
	tokenString, err := token.SignedString([]byte(config.JWTKEY))

	if err == nil {
		fmt.Println("token created")
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorise", tokenString, 3600, "", "", false, true)
}

func ValidateCookie(c *gin.Context) bool {
	cookie, _ := c.Cookie("Authorise")
	if cookie == "" {
		fmt.Println("cookie not found")
		return false
	} else {
		return true
	}

}

func RetriveJwtToken(c *gin.Context) (int, string, string, error) {
	config, _ := utils.LoadConfig("./")
	cookie, _ := c.Cookie("Authorise")
	if cookie == "" {
		return 0, "", "", errors.New("cookie not found")
	} else {
		token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(config.JWTKEY), nil
		})

		if err != nil {
			return 0, "", "", err
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userId := claims["userid"].(float64)
			userPhone := claims["phone"].(string)
			role := claims["role"].(string)
			return int(userId), userPhone, role, nil
		} else {
			return 0, "", "", fmt.Errorf("invalid token")
		}

	}
}

func DeleteCookie(c *gin.Context) error {
	c.SetCookie("Authorise", "", 0, "", "", true, true)
	fmt.Println("cookie deleted")
	return nil
}
