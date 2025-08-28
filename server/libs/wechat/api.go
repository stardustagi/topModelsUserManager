package wechat

import (
	"context"
	"encoding/json"
	"github.com/silenceper/wechat/cache"
	"strconv"

	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/miniprogram"
	"github.com/silenceper/wechat/v2/miniprogram/auth"

	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	mpContext "github.com/silenceper/wechat/v2/miniprogram/context"
)

type wechatClient struct {
	wc *wechat.Wechat
	mp *miniprogram.MiniProgram
}

var WxClient *wechatClient

func Init(cfg []byte) {
	var wxConfig miniConfig.Config
	err := json.Unmarshal(cfg, &wxConfig)
	if err != nil {
		panic("Failed to parse WeChat configuration: " + err.Error())
	}
	WxClient = InitWeichatClient(context.Background(), wxConfig)
	if WxClient == nil {
		panic("Failed to initialize WeChat client")
	}
	// return WxClient
}

func InitWeichatClient(ctx context.Context, miniConfig miniConfig.Config) *wechatClient {
	wc := wechat.NewWechat()
	memory := cache.NewMemory()
	wc.SetCache(memory)
	// Set the official account configuration

	mp := wc.GetMiniProgram(&miniConfig)
	return &wechatClient{
		wc: wc,
		mp: mp,
	}
}

// AuthCode2Session 用于将小程序的 jsCode 转换为 Session
func (w *wechatClient) AuthCode2Session(ctx context.Context, jsCode string) auth.ResCode2Session {
	session, err := w.mp.GetAuth().Code2Session(jsCode)
	if err != nil {
		return auth.ResCode2Session{}
	}
	return session
}

func (w *wechatClient) GetContext() *mpContext.Context {
	return w.mp.GetContext()
}

func (w *wechatClient) DecryptUserInfo(SessionKey, encryptedData, iv string) (map[string]string, error) {
	data, err := w.mp.GetEncryptor().Decrypt(SessionKey, encryptedData, iv)
	if err != nil {
		return nil, err
	}
	var userInfo = make(map[string]string)
	userInfo["OpenId"] = data.OpenID
	userInfo["UnionId"] = data.UnionID
	userInfo["NickName"] = data.NickName
	userInfo["Avatar"] = data.AvatarURL
	userInfo["Gender"] = strconv.Itoa(data.Gender)
	userInfo["City"] = data.City
	userInfo["Province"] = data.Province
	userInfo["Country"] = data.Country
	userInfo["Language"] = data.Language
	userInfo["PhoneNumber"] = data.PhoneNumber
	userInfo["PurePhoneNumber"] = data.PurePhoneNumber
	userInfo["CountryCode"] = data.CountryCode
	userInfo["AppId"] = data.Watermark.AppID
	userInfo["Timestamp"] = strconv.FormatInt(data.Watermark.Timestamp, 10)
	return userInfo, nil
}
