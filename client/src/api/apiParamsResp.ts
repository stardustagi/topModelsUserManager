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

// 获取模型token
export interface GetModelTokenResp {
  api_keys: string[];
}

// 分页数据
export interface PageDataResp {
  total_count: number;
}

// 模型市场
export interface ModelMarketDTO {
  info_address: string;
  info_id: number;
  info_model_id: string;
  info_node_id: string;
  info_name: string;
  info_api_version: string;
  info_deploy_name: string;
  info_input_price: number;
  info_output_price: number;
  info_cache_price: number;
  info_status: string;
  info_last_update: number;
  map_node_id: number;
  map_model_id: number;

  provider_id: number;
  provider_node_id: string;
  provider_provider_id: string;
  provider_type: string;
  provider_name: string;
  provider_endpoint: string;
  provider_api_type: string;
  provider_model_name: string;
  provider_input_price: number;
  provider_output_price: number;
  provider_cache_price: number;
  provider_api_keys: string;
  provider_deleted: number;
  provider_last_update: number;
}

export interface ModelMarketResp {
  errcode: number;
  msg: string;
  data: [ModelMarketDTO[], PageDataResp[]];
}

// 订阅模型
export interface SubscribeModelResp {}

// 移除模型
export interface UnSubscribeModelResp {}

// 我的模型列表
export interface UserSelectModelDTO {
  map_id: number;
  map_model_id: number;
  map_node_id: number;
  model_name: string;
  api_version: string;
  deploy_name: string;
  model_input_price: number;
  model_output_price: number;
  model_cache_price: number;
  status: string;
  last_update: number;
  domain: string;
}
export interface UserSelectModelResp {
  errcode: number;
  msg: string;
  data: {
    id: number;
    models_config: UserSelectModelDTO[];
  };
}

// 模型实体
export interface ModelEntity {
  // id: number;
  name: string;
  provider: string;
  address: string;
  input_price: number;
  output_price: number;
  cache_price: number;
  latency: number;
  health_score: number;
  last_updated: number;
  map_node_id: number;
  map_model_id: number;
}

// 用户消费记录
export interface UserConsumeRecordInfo {
  id: number;
  user_id: number;
  node_id: number;
  discount_amount: number;
  total_consumed: number;
  caller: string;
  model: string;
  model_id: number;
  actual_provider: string;
  actual_provider_id: string;
  consume_type: string;
  total_cost: number;
  created: number;
  updated: number;
}

export interface UserConsumeRecordResp {
  records: UserConsumeRecordInfo[];
  total: number;
}

// 消费明细
export interface UserConsumeDetailText {
  id: number;
  consume_id: number;

  input_tokens: number;
  output_tokens: number;
  cache_tokens: number;

  input_price: number;
  output_price: number;
  cache_price: number;

  created: number;
}
export interface UserConsumeDetailImage {
  id: number;
  consume_id: number;

  quality: string;
  size: string;

  created: number;
}
export interface UserConsumeDetailVideo {
  id: number;
  consume_id: number;

  Seconds: number; // 必须大写 S
  size: string;

  created: number;
}
export interface UserConsumeDetailResp {
  consume_record: UserConsumeRecordInfo;
  details:
    | UserConsumeDetailText[]
    | UserConsumeDetailImage[]
    | UserConsumeDetailVideo[];
  total: number;
}
