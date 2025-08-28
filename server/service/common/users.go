package common

import (
	"encoding/json"
	"fmt"
	"github.com/asmcos/requests"
	"github.com/tidwall/gjson"
	"logistics/constants"
	"logistics/libs/databases"
	myerrors "logistics/libs/errors"
	"logistics/libs/logs"
	"logistics/models"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// 检查时间戳是否过期
func CheckTimestampExpire(timestamp int64) bool {
	if timestamp <= 0 {
		return false
	}
	currentTime := time.Now().Unix()
	if currentTime > timestamp {
		logs.Debug("CheckTimestampExpire timestamp is expired", zap.Int64("timestamp", timestamp), zap.Int64("current_time", currentTime))
		return false
	}
	return true
}

func CheckIsAdmin(dao databases.Dao, uid int64) (bool, *myerrors.StackError) {
	//u := models.Users{}
	//bExists, err := dao.FindById(uid, &u)
	//if err != nil || !bExists {
	//	logs.Error("CheckIsAdmin find user error", zap.Int64("uid", uid), zap.Error(err))
	//	return false, constants.ErrorUserNotFound
	//}
	//if u.IsAdmin != 1 {
	//	logs.Debug("CheckIsAdmin user is not admin", zap.Int64("uid", uid), zap.Int("role", u.Role))
	//	return false, constants.ErrorRoleMissing
	//}
	return true, nil
}

func CheckUserExistsCorporation(dao databases.Dao, uid, corporationId int64) (bool, models.Users, *myerrors.StackError) {
	u := models.Users{}
	//bExists, err := dao.FindById(uid, &u)
	//if err != nil || !bExists {
	//	logs.Error("CheckUserExistsCorporation find user error", zap.Int64("uid", uid), zap.Error(err))
	//	return false, u, constants.ErrorUserNotFound
	//}
	//if u.CorporationId != corporationId {
	//	logs.Debug("CheckUserExistsCorporation user not in this corporation", zap.Int64("uid", uid), zap.Int64("corporation_id", u.CorporationId), zap.Int64("req_corporation_id", corporationId))
	//	return false, u, constants.ErrorUserHasNoCorporation
	//}
	return true, u, nil
}

func GetUserInfo(dao databases.BaseDao, uid int64) (models.Users, *myerrors.StackError) {
	u := models.Users{}
	bExists, err := dao.FindById(uid, &u)
	if err != nil || !bExists {
		logs.Error("GetUserInfo find user error", zap.Int64("uid", uid), zap.Error(err))
		return u, constants.ErrorUserNotFound
	}
	return u, nil
}

func CheckRoleInfo(c echo.Context, role int) (bool, *myerrors.StackError) {
	roleHead := c.Request().Header.Get("role")
	if roleHead == "" {
		logs.Error("role header is missing")
		return false, constants.ErrorRoleMissing
	}
	if roleHead != "0" && roleHead != "1" {
		logs.Error("invalid role", zap.String("role", roleHead))
		return false, constants.ErrorRoleMissing
	}
	roleInt, err := strconv.Atoi(roleHead)
	if err != nil {
		logs.Error("role header is not a valid integer", zap.String("role", roleHead), zap.Error(err))
		return false, constants.ErrorDataConvert
	}
	if roleInt != role {
		logs.Error("role mismatch", zap.Int("expected_role", role), zap.Int("actual_role", roleInt))
		return false, constants.ErrorRoleMissing
	}
	return true, nil
}

func CheckIsPlatformAdmin(dao databases.Dao, uid int64) (bool, *myerrors.StackError) {
	//u := models.Users{}
	//bExists, err := dao.FindById(uid, &u)
	//if err != nil || !bExists {
	//	logs.Error("CheckIsPlatformAdmin find user error", zap.Int64("uid", uid), zap.Error(err))
	//	return false, constants.ErrorUserNotFound
	//}
	//if u.IsPlatformAdmin != 1 {
	//	logs.Debug("CheckIsPlatformAdmin user is not admin", zap.Int64("uid", uid), zap.Int("role", u.Role))
	//	return false, constants.ErrorRoleMissing
	//}
	return true, nil
}

func GetGeocoder(address string) *GeocoderAddress {
	url := fmt.Sprintf("https://apis.map.qq.com/ws/geocoder/v1/?address=%s&key=ACXBZ-DCB3C-VH724-AFFBR-DYVIZ-FFBNQ", address)
	addressReq := requests.Requests()
	response, serr := addressReq.Get(url)
	resps := &GeocoderAddress{}
	if serr != nil {
		logs.Debug("SmartAddressFill validate", zap.Error(serr))
		return resps
	}

	//var data map[string]interface{}
	//err1 := json.Unmarshal(response.Content(), &data)
	//if err1 != nil {
	//	logs.Error("SmartAddressFill unmarshal error", zap.Error(err1))
	//	return common.Response(c, nil, nil)
	//}

	respContent := response.Content()
	province := gjson.GetBytes(respContent, "result.address_components.province")
	city := gjson.GetBytes(respContent, "result.address_components.city")
	district := gjson.GetBytes(respContent, "result.address_components.district")
	street := gjson.GetBytes(respContent, "result.address_components.street")
	street_number := gjson.GetBytes(respContent, "result.address_components.street_number")
	lng := gjson.GetBytes(respContent, "result.location.lng")
	lat := gjson.GetBytes(respContent, "result.location.lat")

	if province.Exists() {
		resps.Province = province.String()
	}
	if city.Exists() {
		resps.City = city.String()
	}
	if district.Exists() {
		resps.District = district.String()
	}
	addr := ""
	if street.Exists() {
		addr += street.String()
	}
	if street_number.Exists() {
		addr += street_number.String()
	}
	if lng.Exists() {
		resps.Lng = lng.Float()
	}
	if lat.Exists() {
		resps.Lat = lat.Float()
	}
	return resps
}

func GetDistance(fromLng float64, fromLat float64, toLng float64, toLat float64) int64 {
	url := fmt.Sprintf("https://apis.map.qq.com/ws/distance/v1/matrix/?mode=driving&from=%f,%f&to=%f,%f&key=ACXBZ-DCB3C-VH724-AFFBR-DYVIZ-FFBNQ", fromLat, fromLng, toLat, toLng)
	addressReq := requests.Requests()
	response, serr := addressReq.Get(url)
	if serr != nil {
		logs.Debug("GetDistance validate", zap.Error(serr))
		return 0
	}

	var data map[string]interface{}
	err1 := json.Unmarshal(response.Content(), &data)
	if err1 != nil {
		logs.Error("GetDistance unmarshal error", zap.Error(err1))
		return 0
	}

	respContent := response.Content()
	distance := gjson.GetBytes(respContent, "result.rows.0.elements.0.distance")
	if distance.Exists() {
		return distance.Int()
	}
	return 0
}
