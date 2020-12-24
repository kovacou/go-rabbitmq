// Copyright © 2020 Alexandre Kovac <contact@kovacou.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

type client struct {
	*amqp.Connection

	cfg Config
	ch  *amqp.Channel
}

func (c *client) Queue(queues ...string) (err error) {
	for _, q := range queues {
		_, err = c.ch.QueueDeclare(q, true, false, false, false, nil)
		if err != nil {
			break
		}
	}

	return
}

func (c *client) MustQueue(queues ...string) {
	if err := c.Queue(queues...); err != nil {
		log.Panic(err)
	}
}

func (c *client) Exchange(exchanges ...string) (err error) {
	for _, ex := range exchanges {
		err = c.ch.ExchangeDeclare(ex, "fanout", true, false, false, false, nil)
		if err != nil {
			break
		}
	}

	return
}

func (c *client) MustExchange(exchanges ...string) {
	if err := c.Exchange(exchanges...); err != nil {
		log.Panic(err)
	}
}

func (c *client) Bind(ex string, queues ...string) (err error) {
	err = c.Exchange(ex)
	if err != nil {
		return
	}

	for _, q := range queues {
		err = c.Queue(q)
		if err != nil {
			break
		}

		err = c.ch.QueueBind(q, "", ex, false, nil)
		if err != nil {
			break
		}
	}

	return
}

func (c *client) MustBind() {
}
