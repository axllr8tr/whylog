# whylog

A logging module for Go that's designed to look nice.

## Example
```go
package main

import "github.com/axllr8tr/whylog"

func main() {
  logger := whylog.NewSimpleLogger() // creating a logging facility

  // standard loglevels
  logger.Log(whylog.TraceLevel, "querying db for user 12..34")
  logger.Log(whylog.DebugLevel, "127.0.0.1: GET /api/v2/messages/channels/56..78")
  logger.Log(whylog.InfoLevel, "user 12..34 signed up")
  logger.Log(whylog.NoticeLevel, "backing up db...")
  logger.Log(whylog.WarnLevel, "db read latency limit exceeded (100ms limit, 580ms actual)")
  logger.Log(whylog.ErrorLevel, "failed to parse config, using defaults")
  logger.Log(whylog.CritLevel, "db integrity error: ...")
  logger.Log(whylog.AlertLevel, "out of resources: CPU load = 8.00 last 10m")
  logger.Log(whylog.EmergLevel, "shutting down...")

  // custom loglevel
  perfLevel := whylog.NewGenericLevel("PERF", "\U000026A1")
  logger.Log(perfLevel, "latency limit exceeded")

  // named logging facility
  logger2 := whylog.NewSimpleLogger()
  logger2.Name = "test"

  // with formatting
  logger2.Logf(whylog.InfoError, "2 + 2 = %d", 4)
}
```
