# logruswr - a Golang Logrus Wrapper

This wrapper enhances Logrus logging in Go applications, providing structured logging, customizable log levels, output formats, log rotation, and the ability to extend functionality with custom hooks.

## Features

- **Structured Logging**: Log messages with contextual information as key-value pairs.
- **Log Levels**: Control log output verbosity from Debug to Fatal.
- **Output Formats**: Choose between JSON and plain text for log output.
- **Log Rotation**: Automatically manage log file size and retention.
- **Custom Hooks**: Extend logging functionality by adding hooks.

## Installation

Use `go get` to install the Logrus wrapper:

```sh
go get github.com/thiagozs/go-logruswr
```
## Usage

### Basic Logging

Initialize the logger and log a simple message:

```go
package main

import (
    "github.com/thiagozs/go-logruswr"
)

func main() {
    logger, _ := logruswr.New()
    logger.Info("Application started successfully")
}
```

### Structured Logging

Log messages with additional context:

```go
logger.WithFields(logruswr.Fields{"user_id": 42}).Info("User logged in")
```

### Error Logging

Log errors with contextual information:

```go
err := errors.New("failed to connect to database")
logger.WithError(err).Error("Database connection error")
```

### Log Rotation

Configure log rotation to manage log files:

```go
logger, _ := logruswr.New(
    logruswr.WithLogFilePath("/var/log/myapp.log"),
    logruswr.WithMaxLogSize(5), // in MB
    logruswr.WithMaxBackups(3),
    logruswr.WithMaxAge(7), // in days
    logruswr.WithCompressLogs(true),
)
```

### Adding Custom Hooks

Extend logging functionality by implementing and adding custom hooks:

```go
type MyCustomHook struct {
    // Custom hook implementation
}

func (hook *MyCustomHook) Levels() []logrus.Level {
    return []logrus.Level{logrus.InfoLevel}
}

func (hook *MyCustomHook) Fire(entry *logrus.Entry) error {
    // Hook logic here
    return nil
}

// Usage
logger.AddHook(&MyCustomHook{})
```

With this setup, your application can leverage powerful logging capabilities with minimal effort. Customize the logger according to your needs, and take advantage of structured logging, log rotation, and extensibility with hooks.

-----

## Versioning and license

Our version numbers follow the [semantic versioning specification](http://semver.org/). You can see the available versions by checking the [tags on this repository](https://github.com/thiagozs/go-logruswr/tags). For more details about our license model, please take a look at the [LICENSE](LICENSE.md) file.

**2024**, thiagozs.
