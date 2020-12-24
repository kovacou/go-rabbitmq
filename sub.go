// Copyright © 2020 Alexandre Kovac <contact@kovacou.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package rabbitmq

import (
	"encoding/json"

	"github.com/kovacou/go-types"
	"github.com/streadway/amqp"
)

type SubParams struct {
	Consumer      string
	AutoAck       bool
	Exclusive     bool
	NoLocal       bool
	NoWait        bool
	Multiple      bool
	RequeueOnNack bool
}

func (c *client) getSubParams(p []SubParams) SubParams {
	if len(p) > 0 {
		return p[1]
	}

	return SubParams{
		Consumer:  c.cfg.Consumer,
		AutoAck:   c.cfg.AutoAck,
		Exclusive: c.cfg.Exclusive,
		NoLocal:   c.cfg.NoLocal,
		NoWait:    c.cfg.NoWait,
	}
}

func (c *client) Sub(q string, cb func(amqp.Delivery, types.Map) bool, sp ...SubParams) (err error) {
	c.MustQueue(q)
	p := c.getSubParams(sp)
	ch, err := c.ch.Consume(q, p.Consumer, p.AutoAck, p.Exclusive, p.NoLocal, p.NoWait, nil)
	if err != nil {
		return
	}

	for m := range ch {
		payload := types.Map{}
		_ = json.Unmarshal(m.Body, &payload)
		if cb(m, payload) {
			m.Ack(p.Multiple)
		} else {
			m.Nack(p.Multiple, p.RequeueOnNack)
		}
	}

	return
}
