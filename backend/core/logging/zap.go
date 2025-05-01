package logging

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
	"time"
)

const (
	Blue   = "\033[34m" // DEBUG
	Green  = "\033[32m" // INFO
	Yellow = "\033[33m" // WARN
	Red    = "\033[31m" // ERROR

	Reset = "\033[0m"
)

func myColoredLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	switch level {
	case zapcore.DebugLevel:
		enc.AppendString(Blue + "DEBUG" + Reset)
	case zapcore.InfoLevel:
		enc.AppendString(Green + "INFO" + Reset)
	case zapcore.WarnLevel:
		enc.AppendString(Yellow + "WARN" + Reset)
	case zapcore.ErrorLevel, zapcore.DPanicLevel, zapcore.PanicLevel, zapcore.FatalLevel:
		enc.AppendString(Red + "ERROR" + Reset)
	default:
		enc.AppendString(level.String())
	}
}

type logEncoder struct {
	zapcore.Encoder
	logDir     string
	logFile    *os.File
	errLogFile *os.File
	date       string
}

func (e *logEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	buf, err := e.Encoder.EncodeEntry(entry, fields)
	if err != nil {
		return nil, err
	}

	// Rotate if day is passed
	if today := time.Now().Format("2006-01-02"); e.date != today {
		// Close old day log files
		if e.logFile != nil {
			err := e.logFile.Close()
			if err != nil {
				return nil, err
			}
		}
		if e.errLogFile != nil {
			err := e.errLogFile.Close()
			if err != nil {
				return nil, err
			}
		}

		// Create new day directory and log files
		if err := os.MkdirAll(e.logDir+"/"+today, 0755); err != nil {
			return nil, err
		}

		// Log file
		logFilePath := e.logDir + "/" + today + "/" + "log.log"
		logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return nil, err
		}
		e.logFile = logFile

		// Error log file
		errLogFilePath := e.logDir + "/" + today + "/" + "err.log"
		errLogFile, err := os.OpenFile(errLogFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return nil, err
		}
		e.errLogFile = errLogFile

		e.date = today
	}

	// To file
	switch entry.Level {
	case zapcore.ErrorLevel:
		_, err := e.errLogFile.WriteString(buf.String())
		if err != nil {
			return nil, err
		}
	default:
		_, err := e.logFile.WriteString(buf.String())
		if err != nil {
			return nil, err
		}
	}

	// To console
	return buf, nil
}

func Init() {

	env := viper.GetString("logging.env")
	if env == "" {
		panic("The [env] is required in the [logging] section of the configuration file.")
	}

	switch {
	case strings.HasPrefix(env, "dev"):
		cfg := zap.NewDevelopmentConfig()
		cfg.EncoderConfig.EncodeLevel = myColoredLevelEncoder
		cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")

		logger, err := cfg.Build()
		if err != nil {
			panic(err)
		}
		zap.ReplaceGlobals(logger)

	case strings.HasPrefix(env, "prod"):
		dir := viper.GetString("logging.dir")
		if dir == "" {
			panic("The [dir] is required in the [logging] section of the configuration file.")
		}

		cfg := zap.NewProductionConfig()
		cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02T15:04:05.000Z07:00")

		encoder := &logEncoder{
			Encoder: zapcore.NewJSONEncoder(cfg.EncoderConfig),
			logDir:  dir,
		}
		core := zapcore.NewCore(
			encoder,
			zapcore.AddSync(os.Stdout),
			zapcore.InfoLevel,
		)

		logger := zap.New(core, zap.AddCaller())
		zap.ReplaceGlobals(logger)

	default:
		panic("Invalid configuration in logging section")
	}
}
