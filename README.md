# RLimiter

:star: Star us on GitHub — it motivates us a lot!

RLimiter is a limiter for sending with worker and maximum limit.

## Table of content

- [RLimiter](#rlimiter)
  - [Table of content](#table-of-content)
  - [Installation](#installation)
  - [Example](#example)
  - [Authors](#authors)

## Installation

`go get github.com/rahman-teja/rlimiter`

## Example

```go
import "github.com/rahman-teja/rlimiter"

mg, _ := rlimiter.NewManager(
		rlimiter.NewConfig().
			SetSender(sender).
			SetMaxLimit(5).
			SetWorker(5),
	)
	
mg.Start()

mg.Send("Your Message")
```

or can view full example `example/main.go`

## Authors

**Rahman Teja**
- [Github](https://github.com/rahman-teja)