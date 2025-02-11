package configuration
import "fmt"

const (
	LogLevelDebug = "debug"
	LogLevelInfo = "info"
	LogLevelWarn = "warn"
	LogLevelError = "error"

	LogFormatText = "text"
	LogFormatJson = "json"


)

type LogConfig struct {
	Level  string
	Format string
}

func (lc *LogConfig) String() string {
	return fmt.Sprintf("Level: %s, Format: %s", lc.Level, lc.Format)

}
