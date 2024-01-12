/*
Package pocketlog exposes an API to log your work.

Instantiate a logger with a pocketlog.New, giving it a threshold level
[Debug, Info, Warn, Error, Fatal].  Messages with lower criticality
than the input threshold will not be logged.

Sharing the logger is the responsibility of the caller
*/
package pocketlog
