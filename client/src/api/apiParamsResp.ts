export type BaseResponse<T=any> = {
    errcode: number;
    msg: string;
    data: T;
}

export type DefaultResp = {}

// 图形验证码
export type ImageCodeResp = {
  base: BaseResponse;
  code: string;
  expireTime: number;
};
