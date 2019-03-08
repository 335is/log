# log
Wraps the standard log package, adding levels, microsecond timestamp resolution, and other features.

## Description

### Levels

Supports the following levels
* None
* Trace
* Debug
* Info
* Warn
* Error

Default level is Info.

LOG_LEVEL environment variable sets the log level.

### Outputs

Output defaults to os.Stderr, but can be directed to any supplied io.Writer, such as a file, buffer, 3rd party log ingest, etc.

Syslog (on Linux) can be supported by setting output to a syslog io.Writer

## Usage

```go
import "github.com/335is/log"

func main() {
    log.Tracef("%s", "Low level verbose log statment")
    log.Debugf("%s", "Default configuration loaded")
    log.Infof("%s", "Starting service version 1.2.3")
    log.Warnf("%s", "Something bad may happen")
    log.Errorf("%s", "Something bad just happened!")
}
```
