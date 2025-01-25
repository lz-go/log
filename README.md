# log

A strongly opinionated Go logging

## Installation

```shell
go get github.com/lz-go/log
```

## Getting started

### Prepare config file

`log` uses these following `config.yaml` configuration:

```yaml
log:
  level: debug # Valid values are "debug", "info", "warn", "error". Default to "debug".
  encoding: json # Valid values are "console", "json". Default to "json".
```

Or, override with environment variables:

```text
LOG_LEVEL=info
LOG_ENCODING=console
```

### Use the log

Anywhere within your application, the logger can be used as:

```go

package user

import "github.com/lz-go/log"

func getUser() {
	logger := log.NewLogger()
	logger.Debug("getUser is being called")
}
```

The logger can also be used with custom field, e.g.

```go
package user

import (
	"context"
	
	"github.com/lz-go/log"

	"github.com/example/example"
)

func getUser(ctx context.Context, id string) (*example.User, error) {
	logger := log.NewLogger(
		log.WithField("func", "getUser"),
		log.WithField("id", id),
    )
	
	user, err := example.GetUserById(ctx, id)
	if err != nil {
		logger.WithField("error", err)
		logger.Errorf("example.GetUserById failed with error: %s", err.Error())
		return nil, err
    }
	
	return user, nil
}
```

### Correlation context

The logger injects the `cxid` field into the output, if the provided context has, e.g.

```go
package main

import (
	"context"
	
	"github.com/lz-go/log"
	
	"github.com/example/example"
)

func main () {
	ctx := log.CorrelationContext(context.Background())
	
	users := getUsers(ctx)
}

func getUsers(ctx context.Context) ([]example.User) {
	logger := log.NewLogger(log.WithContext(ctx))
	logger.Debug("This log has the `cxid` field injected automatically.")

	// ...
}
```
