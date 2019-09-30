package service

import (
	"github.com/sirupsen/logrus"
)

type Config struct {
	Name    string `envconfig:"-"`
	Version string `envconfig:"-"`

	RPCAddr string `default:"0.0.0.0:5004" envconfig:"RPC_ADDR"`

	DBURI string `default:"root:root@/videocoin?charset=utf8&parseTime=True&loc=Local" envconfig:"DBURI"`

	Logger *logrus.Entry `envconfig:"-"`
}
