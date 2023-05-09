package logger

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"runtime/debug"
	"sync"
	"time"

	"github.com/fatih/color"
)

var (
	red    = color.New(color.FgHiRed, color.Bold).SprintFunc()
	yellow = color.New(color.FgHiYellow, color.Bold).SprintFunc()
)

type Level int8

const (
	LevelAll Level = iota
	LevelInfo
	LevelWarning
	LevelError
	LevelFatal
	LevelOff
)

func (l Level) String() string {
	switch l {
	case LevelInfo:
		return "INFO"
	case LevelWarning:
		return "WARNING"
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	default:
		return ""
	}
}

type Logger struct {
	out      io.Writer
	minLevel Level
	useJSON  bool
	colorize bool
	mu       sync.Mutex
}

func New(out io.Writer, minLevel Level, colorize bool) *Logger {
	return &Logger{
		out:      out,
		minLevel: minLevel,
		colorize: colorize,
	}
}

func NewJSON(out io.Writer, minLevel Level) *Logger {
	return &Logger{
		out:      out,
		minLevel: minLevel,
		useJSON:  true,
	}
}

func (l *Logger) Info(format string, v ...any) {
	l.print(LevelInfo, fmt.Sprintf(format, v...))
}

func (l *Logger) Warning(format string, v ...any) {
	l.print(LevelWarning, fmt.Sprintf(format, v...))
}

func (l *Logger) Error(err error) {
	l.print(LevelError, err.Error())
}

func (l *Logger) Fatal(err error) {
	l.print(LevelFatal, err.Error())
	os.Exit(1)
}

func (l *Logger) print(level Level, message string) {
	if level < l.minLevel {
		return
	}

	var line string
	if l.useJSON {
		line = jsonLine(level, message)
	} else {
		line = textLine(level, message, l.colorize)
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	fmt.Fprintln(l.out, line)
}

func textLine(level Level, message string, colorize bool) string {
	t := time.Now().Format(time.RFC3339)
	line := fmt.Sprintf("level=%q time=%q message=%q", level, t, message)

	if colorize {
		switch level {
		case LevelError, LevelFatal:
			line = red(line)
		case LevelWarning:
			line = yellow(line)
		}
	}

	if level >= LevelError {
		line += fmt.Sprintf("\n%s", string(debug.Stack()))
	}

	return line
}

type jsonMessage struct {
	Level   string `json:"level"`
	Time    string `json:"time"`
	Message string `json:"message"`
	Trace   string `json:"trace,omitempty"`
}

func jsonLine(level Level, message string) string {
	data := jsonMessage{
		Level:   level.String(),
		Time:    time.Now().UTC().Format(time.RFC3339),
		Message: message,
	}

	if level >= LevelError {
		data.Trace = string(debug.Stack())
	}

	var line []byte

	line, err := json.Marshal(data)
	if err != nil {
		return fmt.Sprintf("%s: unable to marshal log message: %s", LevelError.String(), err.Error())
	}

	return string(line)
}
