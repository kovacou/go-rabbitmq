// Copyright © 2020 Alexandre Kovac <contact@kovacou.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package rabbitmq

import (
	"fmt"

	"github.com/kovacou/go-env"
)

// cfgEnviron contains the loaded configuration from environment.
var cfgEnviron Config

// init loads the default configuration from environment.
func init() {
	env.Unmarshal(&cfgEnviron)
}

// Config of the client for the connection to RabbitMQ service.
type Config struct {
	User             string `env:"RMQ_USER"`
	Pass             string `env:"RMQ_PASS"`
	Host             string `env:"RMQ_HOST"`
	Port             int16  `env:"RMQ_PORT"`
	Consumer         string `env:"RMQ_CONSUMER"`
	Immediate        bool   `env:"RMQ_IMMEDIATE"`
	Mandatory        bool   `env:"RMQ_MANDATORY"`
	QOSPrefetchCount int    `env:"RMQ_QOS_PREFETCH_COUNT"`
	QOSPrefetchSize  int    `env:"RMQ_QOS_PREFETCH_SIZE"`
	QOSGlobal        bool   `env:"RMQ_QOS_GLOBAL"`
	AutoAck          bool   `env:"RMQ_AUTO_ACK"`
	Exclusive        bool   `env:"RMQ_EXCLUSIVE"`
	NoLocal          bool   `env:"RMQ_NO_LOCAL"`
	NoWait           bool   `env:"RMQ_NO_WAIT"`
}

// URI returns the connection URI based on the configuration.
func (cfg Config) URI() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d", cfg.User, cfg.Pass, cfg.Host, cfg.Port)
}
