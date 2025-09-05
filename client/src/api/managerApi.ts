import { http } from "@/utils/http";
import { BaseResponse, DefaultResp, ImageCodeResp } from "./apiParamsResp";
import { baseUrlApi } from "./utils";

/** 图形验证码 */
export const getImageCodeApi = (data?: object) => {
  return http.request<BaseResponse<ImageCodeResp>>("post", baseUrlApi("/system/graphCode"), { data } );
};

// 手机验证码
export const getPhoneCodeApi = (data?: object) => {
  return http.request<BaseResponse<DefaultResp>>("post", baseUrlApi("/system/phoneCode"), { data })
}

// 邮箱验证码
export const getMailCodeApi = (data?: object) => {
  return http.request<BaseResponse<DefaultResp>>("post", baseUrlApi("/system/mailCode"),{ data })
}

// 账号密码注册登录
export const accountLoginApi = (data?: object) => {
  console.log("data ==== ", data);
  return http.request<BaseResponse<DefaultResp>>("post", baseUrlApi("/system/usernameLogin"), { data })
}

// 手机号注册登录
export const phoneLoginApi = (data?: object) => {
  return http.request<BaseResponse<DefaultResp>>("post", baseUrlApi("/system/phoneLogin"),{ data })
}

// 邮箱注册登录
export const mailLoginApi = (data?: object) => {
  return http.request<BaseResponse<DefaultResp>>("post", baseUrlApi("/system/mailLogin"),{ data })
}