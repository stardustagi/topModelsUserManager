export type BaseResponse<T = any> = {
  errcode: number;
  msg: string;
  data: T;
};

export type DefaultResp = {};

// 图形验证码
export type ImageCodeResp = {
  base: BaseResponse;
  code: string;
  expireTime: number;
};

export type UserInfoDTO = {
  user_id: number;
  user_name: string;
  real_name: string;
  phone: string;
  email: string;
  created_at: number;
  is_admin: number;
  id: number;
};

export type LoginResp = {
  success: boolean;
  message: string;
  user_id: number;
  token: string;
  is_new_user: boolean;
  user_info: UserInfoDTO;
};

// 公司信息
export type CompanyInfoResp = {
  company_name: string;
  license_id: string;
  address: string;
  phone: string;
  mail: string;
  is_realname_authentication: number;
};

// aksk信息
export type UserAkSkDTO = {
  node_id: string;
  access_key: string;
  secret_key: string;
  lastupdate_at: number;
};

export type UserAkSkResp = {
  keys: UserAkSkDTO[];
};
