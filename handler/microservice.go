// Code generated by "grpc-kit-cli/0.3.5". DO NOT EDIT.

package handler

import (
	"github.com/grpc-kit/pkg/cfg"
	"github.com/grpc-kit/pkg/rpc"
	"github.com/sirupsen/logrus"

	"github.com/mingqing/xdserver/modeler"
)

// Microservice 该微服务的结构
type Microservice struct {
	code    string                  // 服务代码
	server  *rpc.Server             // 服务定义
	client  *rpc.Client             // 服务调用
	logger  *logrus.Entry           // 全局日志
	baseCfg *cfg.LocalConfig        // 基础配置
	thisCfg *modeler.IndependentCfg // 个性配置
}

// NewMicroservice 全局只实例化一次
func NewMicroservice(lc *cfg.LocalConfig) (*Microservice, error) {
	// 基础配置初始化
	if err := lc.Init(); err != nil {
		return nil, err
	}

	m := &Microservice{
		code:    lc.Services.ServiceCode,
		logger:  lc.GetLogger(),
		baseCfg: lc,
		thisCfg: &modeler.IndependentCfg{},
	}

	// 自定义配置实例化
	if err := lc.GetIndependent(m.thisCfg); err != nil {
		return m, err
	}

	// RPC客户端服务端
	client, err := lc.GetRPCClient()
	if err != nil {
		return m, err
	}
	server, err := lc.GetRPCServer()
	if err != nil {
		return m, err
	}
	m.client = client
	m.server = server

	// 其他个性扩展逻辑
	if err := m.privateExtended(); err != nil {
		return m, err
	}

	// 自定义配置初始化
	if err := m.thisCfg.Init(); err != nil {
		return m, err
	}

	return m, nil
}