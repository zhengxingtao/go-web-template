package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/emicklei/go-restful"
	"honor/dto"
	"honor/utils"
	"strings"
	"time"
)

var (
	SecretKey = utils.GetYmlProperties().Jwt.Secret
)

type TokenClaims struct {
	UserId string
}

func AuthJWT(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	token := req.HeaderParameter("token")

	if !validJWT(token) {
		resp.AddHeader("Content-Type", "application/json;charset=UTF-8")
		_ = resp.WriteErrorString(401, "{\"statusCode\":401,\"statusMsg\":\"令牌无效\"}")
		return
	}

	chain.ProcessFilter(req, resp)
}

func Authorize(login dto.LoginParam) (string, error) {
	user, err := QueryUserByPhone(login)
	if err != nil {
		return "", err
	}
	if login.Password == "" && login.VerificationCode == "" {
		return "", errors.New("密码或验证码至少录入一个")
	}
	//密码
	if login.Password != "" {
		pwMd5 := utils.PasswordMd5(login.Password, user.Salt)
		if pwMd5 != user.Password {
			return "", errors.New("用户名或密码错误")
		}
	}
	//验证码
	if login.VerificationCode != "" {
		key := "Authorize:" + user.Phone
		result, _ := utils.RedisClient.Get(key).Result()
		if result != login.VerificationCode {
			return "", errors.New("验证码无效")
		}
	}
	//生成token
	claims := TokenClaims{UserId: user.Id.Hex()}
	token := createToken(&claims)
	return token, nil
}

func createToken(tokenClaims *TokenClaims) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(30)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["user"] = tokenClaims.UserId
	token.Claims = claims
	tokenString, _ := token.SignedString([]byte(SecretKey))
	return tokenString
}

func validJWT(token string) bool {
	if len(token) == 0 {
		return false
	}
	parts := strings.Split(token, ".")
	if len(parts) < 3 {
		return false
	}
	err := jwt.SigningMethodHS256.Verify(strings.Join(parts[0:2], "."), parts[2], []byte(SecretKey))
	if err != nil {
		return false
	}
	return true
}
