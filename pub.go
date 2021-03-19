// Copyright © 2020 Alexandre KOVAC <contact@kovacou.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package rabbitmq

import (
	"encoding/json"
	"time"

	"github.com/kovacou/go-types"
	"github.com/streadway/amqp"
)

// PubParams contains parameters when publishing a new message.
type PubParams struct {
	Mandatory  bool
	Immediate  bool
	IsExchange bool // Says if the message is gonna be pushed into an exchange.
}

// getPubParams returns the first parameter found or the default one.
func (c *client) getPubParams(p []PubParams) PubParams {
	if len(p) > 0 {
		return p[1]
	}

	return PubParams{
		Mandatory: c.cfg.Mandatory,
		Immediate: c.cfg.Immediate,
	}
}

// Pub a new message to the given queue/exchange.
func (c *client) Pub(q string, v types.Map, pp ...PubParams) (err error) {
	var e, key string

	m := amqp.Publishing{
		ContentType:  "text/plain",
		DeliveryMode: 2,
	}

	// Adding date of publication (for easy debug)
	v.Set("published_at", time.Now().UTC().String())
	m.Body, _ = json.Marshal(v)

	p := c.getPubParams(pp)
	if p.IsExchange {
		e = q
	} else {
		key = q
	}

	return c.ch.Publish(e, key, p.Mandatory, p.Immediate, m)
}
