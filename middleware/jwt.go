package middleware

import (
	"ginblogMine/utils"
	"ginblogMine/utils/errmsg"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

var JwtKey = []byte(utils.JwtKey)
type MyClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

var code int
// 生成token
func SetToken(username string)(string, int){
	expireTieme := time.Now().Add(10 * time.Hour)
	SetClaims := MyClaims{
		Username:username,
		StandardClaims:jwt.StandardClaims{
			ExpiresAt: expireTieme.Unix(),
			Issuer: "ginblog",
		},

	}
	reqClaims := jwt.NewWithClaims(jwt.SigningMethodHS256,SetClaims)
	token,err := reqClaims.SignedString(JwtKey)
	if err != nil{
		return err.Error(),errmsg.ERROR
	}
	return token,errmsg.SUCCSE
}

// 验证token
func CheckToken(token string)(*MyClaims,int){
	setToken, _ := jwt.ParseWithClaims(token,&MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey,nil
	})
	if key, code := setToken.Claims.(*MyClaims); code && setToken.Valid{
		return key, errmsg.SUCCSE
	}else {
		return nil, errmsg.ERROR
	}
}


// jwt中间件
func JwtToken() gin.HandlerFunc{
	return func(c *gin.Context) {
		tokenHandler := c.Request.Header.Get("Authorization")

		if tokenHandler == ""{
			code = errmsg.ERROR_TOKEN_EXIST
		}
		checkToken := strings.SplitN(tokenHandler," ", 2)
		if len(checkToken) !=2 && checkToken[0] != "Bearer"{
			code = errmsg.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		key, tCode := CheckToken(checkToken[1])
		if tCode == errmsg.ERROR{
			code = errmsg.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		if time.Now().Unix() > key.ExpiresAt{
			code = errmsg.ERROR_TOKEN_RUNTIME
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		c.Set("username",key.Username)
		c.Next()


	}
}















