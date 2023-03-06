package logger

import (
	"context"
	"os"
	"religion/config"
	"runtime"
	"strconv"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm/logger"
)

type Logger interface {
	InitLogger()
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
	Info(context.Context, string, ...interface{})
	Infof(template string, args ...interface{})
	Warn(context.Context, string, ...interface{})
	Warnf(template string, args ...interface{})
	Error(context.Context, string, ...interface{})
	Errorf(template string, args ...interface{})
	DPanic(args ...interface{})
	DPanicf(template string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(template string, args ...interface{})
}

// Logger
type apiLogger struct {
	cfg         *config.Config
	sugarLogger *zap.SugaredLogger
}

// App Logger constructor
func NewApiLogger(cfg *config.Config) Logger {
	return &apiLogger{cfg: cfg}
}

// For mapping config logger to the service logger levels
var loggerLevelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func (l *apiLogger) getLoggerLevel(cfg *config.Config) zapcore.Level {
	level, exist := loggerLevelMap[cfg.Logger.Level]
	if !exist {
		return zapcore.DebugLevel
	}

	return level
}

// Init logger
func (l *apiLogger) InitLogger() {
	logLevel := l.getLoggerLevel(l.cfg)

	logWriter := zapcore.AddSync(os.Stderr)

	var encoderCfg zapcore.EncoderConfig
	if l.cfg.HttpServer.Mode == "Development" {
		encoderCfg = zap.NewDevelopmentEncoderConfig()
	} else {
		encoderCfg = zap.NewProductionEncoderConfig()
	}

	var encoder zapcore.Encoder
	encoderCfg.LevelKey = "LEVEL"
	encoderCfg.CallerKey = "CALLER"
	encoderCfg.TimeKey = "TIME"
	encoderCfg.NameKey = "NAME"
	encoderCfg.MessageKey = "MESSAGE"

	if l.cfg.Logger.Encoding == "console" {
		encoder = zapcore.NewConsoleEncoder(encoderCfg)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	}

	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(encoder, logWriter, zap.NewAtomicLevelAt(logLevel))
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	l.sugarLogger = logger.Sugar()
	if err := l.sugarLogger.Sync(); err != nil {
		l.sugarLogger.Error(err)
	}
}

// Logger methods

func (l *apiLogger) Debug(args ...interface{}) {
	l.sugarLogger.Debug(args...)
}

func (l *apiLogger) Debugf(template string, args ...interface{}) {
	l.sugarLogger.Debugf(template, args...)
}

// func (l *apiLogger) Info(args ...interface{}) {
// 	l.sugarLogger.Info(args...)
// }

func (l *apiLogger) Infof(template string, args ...interface{}) {
	l.sugarLogger.Infof(template, args...)
}

func (l *apiLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	logArgs := makeLogArgs(msg, data)
	l.sugarLogger.Info(logArgs...)
}

func (l *apiLogger) Warnf(template string, args ...interface{}) {
	l.sugarLogger.Warnf(template, args...)
}

func (l *apiLogger) Error(ctx context.Context, msg string, args ...interface{}) {
	logArgs := makeLogArgs(msg, args)
	l.sugarLogger.Error(logArgs...)
}

func (l *apiLogger) Errorf(template string, args ...interface{}) {
	l.sugarLogger.Errorf(template, args...)
}

func (l *apiLogger) DPanic(args ...interface{}) {
	l.sugarLogger.DPanic(args...)
}

func (l *apiLogger) DPanicf(template string, args ...interface{}) {
	l.sugarLogger.DPanicf(template, args...)
}

func (l *apiLogger) Panic(args ...interface{}) {
	l.sugarLogger.Panic(args...)
}

func (l *apiLogger) Panicf(template string, args ...interface{}) {
	l.sugarLogger.Panicf(template, args...)
}

func (l *apiLogger) Fatal(args ...interface{}) {
	l.sugarLogger.Fatal(args...)
}

func (l *apiLogger) Fatalf(template string, args ...interface{}) {
	l.sugarLogger.Fatalf(template, args...)
}

func (l *apiLogger) LogMode(logger.LogLevel) logger.Interface {
	return l
}

func (l *apiLogger) Info(ctx context.Context, s string, args ...interface{}) {
	logArgs := makeLogArgs(s, args)
	l.sugarLogger.Info(logArgs...)
}

func (l *apiLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()
	if rows == -1 {
		l.Infof("[trace] %s, %v, %s, %s", FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
	} else {
		l.Infof("[trace] %s, %v, %d, %s", FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
	}
}

func makeLogArgs(msg string, data ...interface{}) []interface{} {
	logArgs := make([]interface{}, 0, 1+len(data))
	logArgs = append(logArgs, msg, FileWithLineNum(), data)
	return logArgs
}

// FileWithLineNum return the file name and line number of the current file
func FileWithLineNum() string {
	// the second caller usually from gorm internal, so set i start from 2
	for i := 2; i < 15; i++ {
		_, file, line, ok := runtime.Caller(i)
		if ok {
			return file + ":" + strconv.FormatInt(int64(line), 10)
		}
	}

	return ""
}
