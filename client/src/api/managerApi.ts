import { http } from "@/utils/http";
import {
  BaseResponse,
  CompanyInfoResp,
  DefaultResp,
  ImageCodeResp,
  LoginResp
} from "./apiParamsResp";
import { baseUrlApi } from "./utils";

/** 图形验证码 */
export const getImageCodeApi = (data?: object) => {
  return http.request<BaseResponse<ImageCodeResp>>(
    "post",
    baseUrlApi("/system/graphCode"),
    { data }
  );
};

// 手机验证码
export const getPhoneCodeApi = (data?: object) => {
  return http.request<BaseResponse<DefaultResp>>(
    "post",
    baseUrlApi("/system/phoneCode"),
    { data }
  );
};

// 邮箱验证码
export const getMailCodeApi = (data?: object) => {
  return http.request<BaseResponse<DefaultResp>>(
    "post",
    baseUrlApi("/system/mailCode"),
    { data }
  );
};

// 账号密码注册登录
export const accountLoginApi = (data?: object) => {
  return http.request<BaseResponse<LoginResp>>(
    "post",
    baseUrlApi("/system/usernameLogin"),
    { data }
  );
};

// 手机号注册登录
export const phoneLoginApi = (data?: object) => {
  return http.request<BaseResponse<DefaultResp>>(
    "post",
    baseUrlApi("/system/phoneLogin"),
    { data }
  );
};

// 邮箱注册登录
export const mailLoginApi = (data?: object) => {
  return http.request<BaseResponse<DefaultResp>>(
    "post",
    baseUrlApi("/system/mailLogin"),
    { data }
  );
};

// 绑定公司信息
export const bindCompanyInfoApi = (data?: object) => {
  return http.request<BaseResponse<DefaultResp>>(
    "post",
    baseUrlApi("/user/userBindCompanyInfo"),
    { data }
  );
};

// 获取公司信息
export const getCompanyInfoApi = (data?: object) => {
  return http.request<BaseResponse<CompanyInfoResp>>(
    "post",
    baseUrlApi("/user/userGetCompanyInfo"),
    { data }
  );
};

// 绑定手机和邮件
export const bindPhoneAndMailApi = (data?: object) => {
  return http.request<BaseResponse<DefaultResp>>(
    "post",
    baseUrlApi("/user/userBindPhoneAndMail"),
    { data }
  );
};

// 实名认证
export const realNameApi = (data?: object) => {
  return http.request<BaseResponse<DefaultResp>>(
    "post",
    baseUrlApi("/user/userRealName"),
    { data }
  );
};

// 管理员充值
export const adminPaymentApi = (data?: object) => {
  return http.request<BaseResponse<DefaultResp>>(
    "post",
    baseUrlApi("/user/userAdminPayment"),
    { data }
  );
};

// 修改密码
export const changePasswordApi = (data?: object) => {
  return http.request<BaseResponse<DefaultResp>>(
    "post",
    baseUrlApi("/user/userChangePassword"),
    { data }
  );
};

// 节点用户邮箱注册
export const nodeUserEmailRegisterApi = (data?: object) => {
  return http.request<BaseResponse<DefaultResp>>(
    "post",
    baseUrlApi("/system/nodeUserRegister"),
    { data }
  );
};

// 后台邮件登录
export const nodeUserEmailLoginApi = (data?: object) => {
  return http.request<BaseResponse<LoginResp>>(
    "post",
    baseUrlApi("/system/nodeUserEmailLogin"),
    { data }
  );
};

// 申请AKSK
export const nodeUserAddAccessKeyAndSecurityKeyApi = (data?: object) => {
  return http.request<BaseResponse<LoginResp>>(
    "post",
    baseUrlApi("/llmUser/nodeUserAddAccessKeyAndSecurityKey"),
    { data }
  );
};
