package redis

import (
	"context"
	"errors"
	"time"

	"github.com/hashicorp/go-hclog"
	"github.com/redis/go-redis/v9"
)

var ErrorInputValuesIsNil = errors.New("input values is nil")
var ErrorResultNotOK = errors.New("result is not OK")
var ErrorResultNotTrue = errors.New("result is not true")

const RedisKeySep = ":"

var _ RedisCli = (*redisView)(nil)

type redisView struct {
	prefix string
	cmd    RedisCmd
	pubSub redis.PubSub // nolint:unsed
	logger hclog.Logger
}

func (r *redisView) XAdd(ctx context.Context, a redis.XAddArgs) *redis.StringCmd {
	return r.cmd.XAdd(ctx, &a)
}

func NewRedisView(cmd RedisCmd, prefix string, logger hclog.Logger) RedisCli {
	view := &redisView{cmd: cmd, prefix: prefix}
	if logger != nil {
		view.logger = logger.Named("redis")
	}
	return view
}

func (r *redisView) KeyPrefix() string {
	return r.prefix
}
func (r *redisView) NativeCmd() RedisCmd {
	return r.cmd
}

func (r *redisView) SetNX(ctx context.Context, key string, value []byte, duration string) (bool, error) {
	if duration != "" {
		timeout, err := time.ParseDuration(duration)
		if nil != err {
			return false, err
		}
		return r.cmd.SetNX(ctx, r.expandKey(key), value, timeout).Result()
	}
	return r.cmd.SetNX(ctx, r.expandKey(key), value, 0).Result()
}

func (r *redisView) Scan(ctx context.Context, cursor uint64, match string, count int64) ([]string, error) {
	result, _, err := r.cmd.Scan(ctx, cursor, r.expandKey(match), count).Result()
	if nil != err {
		return nil, err
	}
	return result, err
}

func (r *redisView) Get(ctx context.Context, key string) ([]byte, error) {
	if r.logger != nil && r.logger.IsTrace() {
		r.logger.Trace("get", "key", r.expandKey(key))
	}
	result, err := r.cmd.Get(ctx, r.expandKey(key)).Result()
	if nil != err {
		return nil, err
	}
	return []byte(result), nil
}

func (r *redisView) Set(ctx context.Context, key string, value []byte, duration string) error {
	if r.logger != nil && r.logger.IsTrace() {
		r.logger.Trace("set", "key", r.expandKey(key), "value", string(value))
	}

	if duration != "" {
		timeout, err := time.ParseDuration(duration)
		if nil != err {
			return err
		}
		return r.cmd.Set(ctx, r.expandKey(key), value, timeout).Err()
	}
	return r.cmd.Set(ctx, r.expandKey(key), value, 0).Err()
}

func (r *redisView) Del(ctx context.Context, keys ...string) (int64, error) {
	var all []string
	for _, key := range keys {
		all = append(all, r.expandKey(key))
	}
	return r.cmd.Del(ctx, all...).Result()
}

func (r *redisView) Expire(ctx context.Context, key string, duration string) error {
	timeout, err := time.ParseDuration(duration)
	if nil != err {
		return err
	}
	return wrapResult(func() (interface{}, error) {
		return r.cmd.Expire(ctx, r.expandKey(key), timeout).Result()
	})
}
