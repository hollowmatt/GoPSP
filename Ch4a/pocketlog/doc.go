/*
Package pocketlog exposes an API to log your work.

Instantiate a logger with pocketlog.New, giving a threshold level.
Only meessages at that level or higher will be logged.

# Sharing the logger is the responsibility of the caller

The logger can be called with the following levels:
  - Debug
  - Info
  - Warn
  - Error
  - Fatal
*/
package pocketlog
