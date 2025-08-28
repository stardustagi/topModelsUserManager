package wechat

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"logistics/constants"
	"logistics/libs/databases"
	"logistics/libs/jwt"
	"logistics/libs/logs"
	"logistics/libs/redis"
	"logistics/libs/uuid"
	"logistics/libs/wechat"
	"logistics/models"
	"logistics/service/common"
	"time"
)

// WxCode2Session handles the WeChat login process by exchanging the code for a session
// @Summary WeChat Code to Session
// @Description Exchanges WeChat code for session and user info
// @Tags WeChat
// @Accept json
// @Produce json
// @Param request body AuthRequest true "WeChat Auth Request"
// @Success 200 {object} common.BaseResponse{data=AuthResponse} "Successful response with user info"
// @Failure 400 {object} common.BaseResponse "Invalid request"
// @Failure 500 {object} common.BaseResponse "Internal server error"
// @Router /wechat/code2session [post]
func WxCode2Session(c echo.Context) error {
	req, err := common.ValidateRequest[AuthRequest](c)
	if err != nil {
		logs.Debug("WxCode2Session validate", zap.Error(err))
		return common.Response(c, constants.ErrorInvalidRequest, nil)
	}
	ctx := c.Request().Context()

	session := wechat.WxClient.AuthCode2Session(ctx, req.JsCode)
	if session.OpenID == "" {
		logs.Info("session openid null", zap.String("js code", req.JsCode))
		return common.Response(c, constants.ErrorPasswordSalt, nil)
	}

	red := redis.GetRedisDb()
	// 取IP和时间
	ip := echo.ExtractIPFromXFFHeader()(c.Request())
	_t := time.Now().Unix()

	dbSession := databases.GetDao().NewSession()
	defer dbSession.Close()

	// 准备salt和password
	salt := uuid.GenString(8)
	pwd, jwtEncErr := jwt.EncryptWithFixedSalt(constants.UserDefaultPassword, constants.ObtenationIterations, session.OpenID, salt)
	if jwtEncErr != nil {
		logs.Info("jwt EncryptWithFixedSalt error", zap.Error(err))
		return common.Response(c, constants.ErrorPasswordSalt, nil)
	}

	// 查账号是否注册过
	u := models.Users{
		OpenId: session.OpenID,
	}
	var dbErr error
	var userExist bool
	if userExist, dbErr = dbSession.FindOne(&u); dbErr != nil {
		logs.Info("find one by openid number", zap.Error(err))
		return common.Response(c, constants.ErrDatabaseQuery, nil)
	}
	if !userExist {
		u = models.Users{
			OpenId:      session.OpenID,
			UnionId:     session.UnionID,
			LastLoginIp: ip,
			Password:    pwd,
			Salt:        salt,
			CreatedAt:   _t,
		}
		// 新增用户
		_, dbErr = dbSession.InsertOne(&u)
		if dbErr != nil {
			logs.Info("insert one", zap.Error(dbErr))
			return common.Response(c, constants.ErrDatabaseCommit, nil)
		}
	}

	// 登录发jwt
	if pnum, err := jwt.DecryptWithFixedSalt(constants.UserDefaultPassword, constants.ObtenationIterations, u.Password, u.Salt); err != nil || pnum != u.OpenId {
		logs.Info("DecryptWithFixedSalt error", zap.Error(err))
		return common.Response(c, constants.ErrorPasswordSalt, nil)
	}
	u.LastLoginIp = ip
	u.LastLoginTime = _t

	if _, dbErr = dbSession.UpdateById(u.Id, &u); dbErr != nil {
		logs.Debug(fmt.Sprintf("update user %s", u.OpenId), zap.Error(dbErr))
		return common.Response(c, constants.ErrDatabaseCommit, nil)
	}

	k := jwt.UserTokenKey(fmt.Sprintf("%d", u.Id))
	token := jwt.GenString(32)

	//if result, err := red.Set(context.TODO(), k, token, time.Hour*5).Result(); err != nil {
	if result, err := red.Set(context.TODO(), k, token, 0).Result(); err != nil {
		return common.Response(c, constants.ErrRedisConnection, nil)
	} else {
		logs.Debug("redis set success", zap.String("result", result))
	}

	logs.Info("account login info", zap.String("token", token), zap.String("key", k))
	jwtKey := fmt.Sprintf("%s-%s", constants.AppName, constants.AppVersion)
	jwtString := jwt.JWTEncrypt(fmt.Sprintf("%d", u.Id), token, jwtKey)
	c.Response().Header().Add("jwt", jwtString)

	resp := &AuthResponse{
		UserId: u.Id,
	}
	return common.Response(c, nil, resp)

	//u := models.Users{}
	//dao := databases.GetDao()
	//if bExists, dbErr := dao.FindById(req.Role, &u); dbErr != nil || !bExists {
	//	logs.Error("AccountApply find user error", zap.Int64("uid", 1), zap.Error(dbErr))
	//	return common.Response(c, constants.ErrorUserNotFound, nil)
	//}
	//
	////c.Response().Header().Add("jwt", jwtString)
	//var resp AuthResponse
	//resp.UserId = u.Id
	//resp.Role = u.Role
	//resp.Success = true
	//resp.IsAdmin = u.IsAdmin
	//resp.IsAudit = u.IsAudit
	//resp.IsAuth = u.IsImproveUserInfo
	//resp.Name = u.Name
	//resp.CompanyName = "企业1"
	//resp.Phone = u.PhoneNumber
	//resp.CorporationId = u.CorporationId
	//resp.Expire = u.Expire
	//resp.IsPlatformAdmin = u.IsPlatformAdmin
	//
	//return common.Response(c, nil, resp)
}
