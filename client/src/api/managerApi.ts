import { http } from "@/utils/http";
import {
  BaseResponse,
  CompanyInfoResp,
  DayReportResp,
  DefaultResp,
  GetModelTokenResp,
  ImageCodeResp,
  LoginResp,
  ModelMarketResp,
  SubscribeModelResp,
  UserAkSkResp,
  UserConsumeDetailResp,
  UserConsumeRecordResp,
  UserSelectModelResp
} from "./apiParamsResp";
import { baseUrlApi } from "./utils";
import { object } from "vue-types";

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

// 模型市场
export const getModelMarketApi = (data?: object) => {
  return http.request<ModelMarketResp>("post", baseUrlApi("/system/query"), {
    data
  });
};

// 获取模型的延迟
export const getModelDelayApi = (data?: object) => {
  return http.request<BaseResponse<DefaultResp>>(
    "post",
    baseUrlApi("/user/userGetNodesStatus"),
    { data }
  );
};

// 订阅模型
export const subscribeModelApi = (data?: object) => {
  return http.request<BaseResponse<SubscribeModelResp>>(
    "post",
    baseUrlApi("/user/userSaveLLMInfo"),
    { data }
  );
};

// 移除模型
export const unsubscribeModelApi = (data?: object) => {
  return http.request<BaseResponse<SubscribeModelResp>>(
    "post",
    baseUrlApi("/user/userSaveLLMInfo"),
    { data }
  );
};

// 获取模型
export const userGetSelectLLMInfo = (data?: object) => {
  return http.request<UserSelectModelResp>(
    "post",
    baseUrlApi("/user/userGetSelectLLMInfo"),
    { data }
  );
};

// 获取模型token
export const getModelKeysApi = (data?: object) => {
  return http.request<BaseResponse<GetModelTokenResp>>(
    "post",
    baseUrlApi("/user/userGetApiKeys"),
    { data }
  );
};
// 设置模型token
export const setModelKeysApi = (data?: object) => {
  return http.request<BaseResponse<DefaultResp>>(
    "post",
    baseUrlApi("/user/userSaveApiKeys"),
    { data }
  );
};

// apikey 登录
export const apikeyLoginApi = (data?: object) => {
  return http.request<BaseResponse<LoginResp>>(
    "post",
    baseUrlApi("/system/loginFormUserApiKey"),
    { data }
  );
};

// 模型分组
export const upsertUserModeGroupApi = (data?: object) => {
  return http.request<BaseResponse<DefaultResp>>(
    "post",
    baseUrlApi("/user/upsertUserModeGroup"),
    { data }
  );
};

// 获取模型分组
export const getUserModeGroupsApi = (data?: object) => {
  return http.request<BaseResponse<DefaultResp>>(
    "post",
    baseUrlApi("/user/getUserModeGroups"),
    { data }
  );
};

// 删除模型分组
export const deleteUserModeGroupApi = (data?: object) => {
  return http.request<BaseResponse<DefaultResp>>(
    "post",
    baseUrlApi("/user/deleteUserModeGroup"),
    { data }
  );
};

// 统计相关

// 实时用量查询
export const realtimeApi = (data?: object) => {
  return http.request<BaseResponse<DefaultResp>>(
    "get",
    baseUrlApi("/v1/usage/realtime", "stats"),
    { data }
  );
};

// 历史趋势查询
export const trendsApi = (data?: object) => {
  return http.request<BaseResponse<DefaultResp>>(
    "get",
    baseUrlApi("/v1/usage/trends", "stats"),
    { data }
  );
};

// 模型占比统计
export const modelProportionApi = (data?: object) => {
  return http.request<BaseResponse<DefaultResp>>(
    "get",
    baseUrlApi("/v1/usage/model-proportion", "stats"),
    { data }
  );
};

// 预算预警状态查询
export const budgetAlertApi = (data?: object) => {
  return http.request<BaseResponse<DefaultResp>>(
    "get",
    baseUrlApi("/v1/usage/budget-alert", "stats"),
    { data }
  );
};

// 明细导出数据查询
export const detailExportApi = (data?: object) => {
  return http.request<BaseResponse<DefaultResp>>(
    "get",
    baseUrlApi("/v1/usage/detail-export", "stats"),
    { data }
  );
};

// 多APIKey聚合统计
export const apiKeyAggregateApi = (data?: object) => {
  return http.request<BaseResponse<DefaultResp>>(
    "get",
    baseUrlApi("/v1/usage/api-key-aggregate", "stats"),
    { data }
  );
};

// 用户消费清单
export const getUserConsumeRecordApi = (data?: object) => {
  return http.request<BaseResponse<UserConsumeRecordResp>>(
    "post",
    baseUrlApi("/user/getUserConsumeRecord"),
    { data }
  );
};

// 明细数据
export const getUserConsumeDetailApi = (data?: object) => {
  return http.request<BaseResponse<UserConsumeDetailResp>>(
    "post",
    baseUrlApi("/user/getUserConsumeDetail"),
    { data }
  );
};

// 日报
export const getDailyReportApi = (data?: object) => {
  return http.request<BaseResponse<DayReportResp>>(
    "post",
    baseUrlApi("/user/getUsersDayReport"),
    { data }
  );
};
