// Copyright © 2020 Alexandre Kovac <contact@kovacou.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package rabbitmq

import "fmt"

// Config of the client.
type Config struct {
	User string `env:"RMQ_USER"`
	Pass string `env:"RMQ_PASS"`
	Host string `env:"RMQ_HOST"`
	Port int8   `env:"RMQ_PORT"`

	Consumer         string
	Immediate        bool
	Mandatory        bool
	QOSPrefetchCount int
	QOSPrefetchSize  int
	QOSGlobal        bool
	AutoAck          bool
	Exclusive        bool
	NoLocal          bool
	NoWait           bool
}

// URI returns the connection URI based on the configuration.
func (cfg Config) URI() string {
	return fmt.Sprintf("")
}
