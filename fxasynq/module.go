package fxasynq

import (
	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
)

// ModuleName is the module name.
const ModuleName = "asynq"

// FxAsynq is the [Fx] asynq module.
//
// [Fx]: https://github.com/uber-go/fx
var FxAsynqModule = fx.Module(
	ModuleName,
	fx.Provide(NewFxAsynqClient),
)

// FxWorkerPoolParam allows injection of the required dependencies in [NewFxAsynqClient].
type FxAsynqClientParam struct {
	fx.In
	LifeCycle fx.Lifecycle
	client *redis.Client
}

// NewFxAsynqClient returns a new [asynq.Client].
func NewFxAsynqClient(p FxAsynqClientParam) *asynq.Client {
	return asynq.NewClientFromRedisClient(p.client)
}
