// 处理 JWT 认证

package jwt

import (
	"errors"
	"os"
	"strconv"

	// "strings"
	"time"

	// "github.com/gin-gonic/gin"
	jwtpkg "github.com/golang-jwt/jwt"
	"github.com/sjxiang/gohub/conf"
	"github.com/sjxiang/gohub/pkg/logger"
	"github.com/sjxiang/gohub/pkg/util"
)


var (
	ErrTokenExpired            error = errors.New("令牌已过期")
	ErrTokenExpiredMaxRefresh  error = errors.New("令牌已过最大刷新时间")
	ErrTokenMalformed          error = errors.New("请求令牌格式有误")
	ErrTokenInvalid            error = errors.New("请求令牌无效")
	ErrHeaderEmpty             error = errors.New("需要认证才能访问")
	ErrHeaderMalformed         error = errors.New("请求头中 Authorization 格式有误")
)


// JWT 定义一个 jwt 对象
type JWT struct {

	// 密钥，用以加密 JWT，读取配置消息 APP_KEY
	SignKey []byte

	// 刷新 Token 的最大过期时间
	MaxRefresh time.Duration

}


// JWTCustomClaims 自定义载荷（payload）
type JWTCustomClaims struct {
	UserID 			string `json:"user_id"`
	UserName 		string `json:"user_name"`
	ExpireAtTime 	int64  `json:"expire_time"`



	/* 
	
	StandardClaims 结构体实现了 Claims 接口继承了 Valid() 方法
	
	JWT 规定了 7 个官方字段，提供使用：

	- iss（issuer）：发布者
	- sub（suject）：主题
	- iat（Issued At）：生成签名的时间
	- exp（expiration time）：签名过期时间
	- aud（audience）：观众，相当于接收者
	- nbf（Not Before）：生效时间
	- jti（JWT ID）：编号
	
	*/
	jwtpkg.StandardClaims
}


func NewJWT() *JWT {

	refreshTime, _ := strconv.Atoi(os.Getenv("JWT_MAX_REFRESH_TIME"))

	return &JWT{
		SignKey: []byte(os.Getenv("APP_KEY")),
		MaxRefresh: time.Duration(refreshTime) * time.Minute,
	}
}


// GenToken 生成 JWT Token，在登录成功时调用
func (jwt *JWT) GenToken(userID string, username string) string {

	var err error

    // 1. 创建自定义的声明

    expireAtTime := jwt.expireAtTime()
    claims := JWTCustomClaims{
		// 自定义字段
    	UserID: userID,
        UserName: username,
        ExpireAtTime: expireAtTime,
		// 原生
        StandardClaims: jwtpkg.StandardClaims{
            NotBefore: util.TimenowInTimezone().Unix(), // 签名生效时间
            IssuedAt:  util.TimenowInTimezone().Unix(), // 首次签名时间（后续刷新 Token 不会更新）
            ExpiresAt: expireAtTime,                   // 签名过期时间
            Issuer:    os.Getenv("APP_NAME"),   // 签发者
        },
    }

	// 2. 使用指定的签名方法，创建签名对象 
	token := jwtpkg.NewWithClaims(jwtpkg.SigningMethodHS256, claims)
    
	// 3. 使用指定的『私钥』签名，获得完整的编码后的字符串，即 token
	result, err := token.SignedString(jwt.SignKey)
	if err != nil {
		logger.LogIf(err)
		return ""
	}

	return result
}


// expireAtTime 过期时间
func (jwt *JWT) expireAtTime() int64 {
	timeNow := util.TimenowInTimezone()
	expireTime, _ := strconv.Atoi(os.Getenv("JWT_EXPIRE_TIME"))

	if conf.IsLocal() {
		expireTime, _ = strconv.Atoi(os.Getenv("JWT_DEBUG_EXPIRE_TIME"))	
	}

	expire := time.Duration(expireTime) * time.Minute
	return timeNow.Add(expire).Unix()  // 1661160279 从格林威治时间1970年01月01日00时00分00秒起至现在的总秒数。
}



// ParseToken 解析 JWT token 
func (jwt *JWT)ParseToken(tokenString string) (*JWTCustomClaims, error) {

	// 1. 解析 token 
	token, err := jwtpkg.ParseWithClaims(
		tokenString, 
		&JWTCustomClaims{}, 
		func(token *jwtpkg.Token) (interface{}, error) {
        	return jwt.SignKey, nil
    	},
	)

	// 解析出错
	if err != nil {
		validationErr, ok := err.(*jwtpkg.ValidationError)
		if ok {
			if validationErr.Errors == jwtpkg.ValidationErrorMalformed {
				return nil, ErrHeaderMalformed
			} else if validationErr.Errors == jwtpkg.ValidationErrorExpired {
				return nil, ErrTokenExpired
			}
		}

		return nil, ErrTokenInvalid
	}

	// 3. 将 token 中的 claims 信息解析出来和 JWTCustomClaims 数据结构进行校验
	if claims, ok := token.Claims.(*JWTCustomClaims); ok && token.Valid {
		return claims, nil 
	}

	return nil, ErrTokenInvalid
}



// e.g. Authorization:Bearer xxxx
// func (jwt *JWT) getTokenFromHeader(c *gin.Context) (string, error) {
// 	authHeader := c.Request.Header.Get("Authorization")
// 	if authHeader == "" {
// 		return "", ErrHeaderEmpty
// 	}

// 	// 按空格分隔
// 	parts := strings.SplitN(authHeader, " ", 2)
// 	if !(len(parts) == 2 && parts[0] == "Bearer") {
// 		return "", ErrHeaderMalformed
// 	}

// 	return parts[1], nil
// }

// // ParseToken 解析 JWT Token，中间件中调用
// func (jwt *JWT) ParserToken(c *gin.Context) (*JWTCustomClaims, error) {

// 	tokenString, parseErr := jwt.getTokenFromHeader(c)
// 	if parseErr != nil {
// 		return nil, parseErr
// 	}

// 	// 1. 调用 jwt 库，解析用户传参的 Token
// 	token, err := jwt.parseTokenString(tokenString)
	
// 	// 2. 解析出错
// 	if err != nil {
// 		validationErr, ok := err.(*jwtpkg.ValidationError)
// 		if ok {
// 			if validationErr.Errors == jwtpkg.ValidationErrorMalformed {
// 				return nil, ErrHeaderMalformed
// 			} else if validationErr.Errors == jwtpkg.ValidationErrorExpired {
// 				return nil, ErrTokenExpired
// 			}
// 		}

// 		return nil, ErrTokenInvalid
// 	}

// 	// 3. 将 token 中的 claims 信息解析出来和 JWTCustomClaims 数据结构进行校验
// 	if claims, ok := token.Claims.(*JWTCustomClaims); ok && token.Valid {
// 		return claims, nil 
// 	}

// 	return nil, ErrTokenInvalid
// }





// // RefreshToken 更新 Token，用以提供 refresh token 接口
// func (jwt *JWT) RefreshToken(c *gin.Context) (string, error) {

// 	// 1. 从 Header 里获取 token
// 	tokenString, parseErr := jwt.getTokenFromHeader(c)
// 	if parseErr != nil {
// 		return "", parseErr
// 	}

// 	// 2. 调用 jwt 库，解析用户传参的 Token
// 	token, err := jwt.parseTokenString(tokenString)

// 	// 3. 解析出错，未出错证明是合法的 Token（甚至未到过期时间）
// 	if err != nil {
// 		validationErr, ok := err.(*jwtpkg.ValidationError)

// 		// 满足 refresh 的条件：只是单一的报错 ValidationErrorExpired
// 		if !ok || validationErr.Errors != jwtpkg.ValidationErrorExpired {
// 			return "", err
// 		}
// 	}

// 	// 4. 解析 JWTCustomClaims 的数据
// 	claims := token.Claims.(*JWTCustomClaims)

// 	// 5. 检查是否错过了『最大允许刷新的时间』
// 	x := util.TimenowInTimezone().Add(-jwt.MaxRefresh).Unix()
// 	if claims.IssuedAt > x {
// 		claims.StandardClaims.ExpiresAt = jwt.expireAtTime()
// 		return jwt.createToken(*claims)
// 	}

// 	return "", ErrTokenExpiredMaxRefresh
// }

