# Journal Tool

## Installation
1. You can use the below Go command to install Journal tool.

```sh
$ go get -u github.com/xiaocai-go/journal
```

2. Import it in your code:

```go
import "github.com/xiaocai-go/journal"
```

## Quick start
```go
package main

import (
	"github.com/xiaocai-go/journal"
	"os"
)

func main() {
	logger := journal.New(&journal.Option{
		Writer:         os.Stdout,
		Encoder:        journal.ConsoleEncoder,
		Level:          journal.InfoLevel,
		SkipLineEnding: false,
	})

	logger.Debug("this is message")
	logger.Info("this is message")
	logger.Warn("this is message")
	logger.Error("this is message")
	logger.Fatal("this is message")
}
```