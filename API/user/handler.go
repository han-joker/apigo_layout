package User

import (
	"PH_ModuleName_PH/tool"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type RequestSignIn struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

func SignIn(c *gin.Context) {
	var req RequestSignIn
	if err := c.ShouldBind(&req); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	user := &User{}
	if result := tool.DB.First(user, "username=?", req.Username); result.Error != nil {
		if errors.Is(gorm.ErrRecordNotFound, result.Error) {
			c.JSON(http.StatusOK, gin.H{
				"message": "user not exists",
			})
			return
		}
		log.Println(result.Error)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	if user.Password != fmt.Sprintf("%X", sha256.Sum256([]byte(req.Password))) {
		c.JSON(http.StatusOK, gin.H{
			"message": "user password error",
		})
		return
	}
	// Create a new token object, specifying signing method and the claims
	claims := jwt.StandardClaims{
		Audience:  "ApiGo-agent",
		ExpiresAt: time.Now().Add(30 * 24 * 3600 * time.Second).UnixMilli(),
		Id:        fmt.Sprintf("%d", user.ID),
		IssuedAt:  time.Now().UnixMilli(),
		Issuer:    "ApiGo",
		NotBefore: 0,
		Subject:   "user-authentication",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	hmacSecret := make([]byte, 8)
	if _, err := rand.Read(hmacSecret); err != nil {
	}
	tokenString, err := token.SignedString(hmacSecret)
	if err != nil {
		log.Fatalln(err)
	}
	result := tool.DB.Model(&user).Updates(&User{
		Rand:  hmacSecret,
		Token: tokenString,
	})
	log.Println(result.Error)
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

func SignUp(c *gin.Context) {

}

func SignOut(c *gin.Context) {

}

func Post(c *gin.Context) {
	var row User
	if err := c.ShouldBind(&row); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	// hash password
	row.Password = fmt.Sprintf("%X", sha256.Sum256([]byte(row.Password)))

	if result := tool.DB.Create(&row); result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusCreated, row)
}
func Delete(c *gin.Context) {
	id := c.Param("id")
	if result := tool.DB.Delete(&User{}, "id = ?", id); result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
func Put(c *gin.Context) {
	id := c.Param("id")
	var row User
	if result := tool.DB.First(&row, "id = ?", id); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": tool.MResourceNotFound,
		})
		return
	}

	if err := c.ShouldBind(&row); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if result := tool.DB.Save(&row); result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, row)
}

func List(c *gin.Context) {
	var rows []User
	if result := tool.DB.Find(&rows); result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusOK, []User{})
		return
	}
	c.JSON(http.StatusOK, rows)
}

func Retrieve(c *gin.Context) {
	id := c.Param("id")
	var row User
	if result := tool.DB.First(&row, "id = ?", id); result.Error != nil {
		log.Println(result.Error, row)
		c.JSON(http.StatusOK, nil)
		return
	}
	c.JSON(http.StatusOK, row)
}
