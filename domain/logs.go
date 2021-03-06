package domain

import "fmt"

const (
	// LogTypeOut for stdout
	LogTypeOut LogType = "STDOUT"

	// LogTypeErr for stderr
	LogTypeErr LogType = "STDERR"
)

// LogType stdout or stderr
type LogType string

// LogItem line of log output
type LogItem struct {
	CmdID   string
	Type    LogType
	Content string
	Number  int64
}

func (item LogItem) String() string {
	return fmt.Sprintf("%s#%s#%d#%s", item.CmdID, item.Type, item.Number, item.Content)
}
