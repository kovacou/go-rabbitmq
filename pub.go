// Copyright © 2020 Alexandre Kovac <contact@kovacou.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package rabbitmq

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/kovacou/go-types"
	"github.com/streadway/amqp"
)

type PubParams struct {
	Mandatory bool
	Immediate bool
}

func (c *client) getPubParams(p []PubParams) PubParams {
	if len(p) > 0 {
		return p[1]
	}

	return PubParams{
		Mandatory: c.cfg.Mandatory,
		Immediate: c.cfg.Immediate,
	}
}

func (c *client) Pub(q string, v types.Map, pp ...PubParams) (err error) {
	var e, key string

	m := amqp.Publishing{
		ContentType:  "text/plain",
		DeliveryMode: 2,
	}
	m.Body, _ = json.Marshal(v)

	if strings.HasSuffix(q, SuffixQueue) {
		key = q
	} else if strings.HasSuffix(q, SuffixExchange) {
		e = q
	} else {
		log.Panicf("The queue/exchange destination is not valid. You must respect the pattern xxx%s or xxx%s", SuffixQueue, SuffixExchange)
	}

	p := c.getPubParams(pp)
	return c.ch.Publish(e, key, p.Mandatory, p.Immediate, m)
}
