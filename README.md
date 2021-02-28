# github.com/kovacou/go-rabbitmq

Personal project.
Easy & fast client for RabbitMQ allowing bindings, publishing & subscribing.

## Usages
```go
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
```

### From environment
`rabbitmq.Open()` is using default environment variables to load the configuration. 
You can also load environment with a specific prefix by using `rabbitmq.OpenEnv("RABBIT")`

### From struct
A new connection can be open from a config struct through `rabbitmq.OpenWith(cfg)`.

## Declare
### Queues
```go
func init() {
	err := rb.Queue("my_queue", "my_second_queue")
	if err != nil {
		panic(err)
	}
    
    // Equals to
    rb.MustQueue("my_queue", "my_second_queue")
}
```
### Exchanges
```go
func init() {
	err := rb.Exchange("my_exchange", "my_second_exchange")
	if err != nil {
		panic(err)
	}
    
    // Equals to
    rb.MustQueue("my_exchange", "my_second_exchange")
}
```
## Subscribing

## Publishing
