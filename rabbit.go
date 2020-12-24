// Copyright © 2020 Alexandre Kovac <contact@kovacou.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package rabbitmq

import (
	"fmt"

	"github.com/kovacou/go-types"
	"github.com/streadway/amqp"
)

var (
	SuffixQueue    = "_q"
	SuffixExchange = "_e"
)

// Connection is a client for RabbitMQ.
type Connection interface {
	Close() error
	Bind(string, ...string) error
	Exchange(...string) error
	MustQueue(...string)
	MustExchange(...string)
	Pub(string, types.Map, ...PubParams) error
	Queue(...string) error
	Sub(string, func(d amqp.Delivery, m types.Map) (ack bool), ...SubParams) error
}

// Open a new connection to rabbitmq with default config.
func Open() (Connection, error) {
	return OpenWith(cfgEnviron)
}

// OpenEnv opens a new connection from environment.
func OpenEnv() {
}

// OpenWith opens a new connection with the given config.
func OpenWith(cfg Config) (con Connection, err error) {
	amqpc, err := amqp.Dial(cfg.URI())
	if err != nil {
		return
	}

	ch, err := amqpc.Channel()
	if err != nil {
		return
	}

	if err = ch.Qos(cfg.QOSPrefetchCount, cfg.QOSPrefetchSize, cfg.QOSGlobal); err != nil {
		return
	}

	return &client{Connection: amqpc, ch: ch}, nil
}

// parseQueue add the queue suffix to the given name.
func parseQueue(name string) string {
	return fmt.Sprintf("%s%s", name, SuffixQueue)
}

// parseExchange add the exchange suffix to the given name.
func parseExchange(name string) string {
	return fmt.Sprintf("%s%s", name, SuffixExchange)
}
