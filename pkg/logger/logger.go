package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"sync"
)

var (
	_lock   sync.Mutex
	_logger = map[string]*zap.SugaredLogger{}
)

var levelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func getLoggerLevel(lvl string) zapcore.Level {
	if level, ok := levelMap[lvl]; ok {
		return level
	}
	return zapcore.InfoLevel
}

// 获得默认Logger对象
func GetDefaultLogger() *zap.SugaredLogger {
	return GetLogger("", "")
}

// 获得Logger对象
func GetLogger(filePath string, level string) *zap.SugaredLogger {

	if filePath == "" {
		filePath = "run.log"
	}

	if logger, ok := _logger[filePath]; ok {
		return logger
	}

	return getLogger(filePath, level)
}

func getLogger(filePath string, lvl string) *zap.SugaredLogger {

	_lock.Lock()
	defer _lock.Unlock()

	if logger, ok := _logger[filePath]; ok {
		return logger
	}

	level := getLoggerLevel(lvl)
	syncWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   filePath, // ⽇志⽂件路径
		MaxSize:    1024,     // megabytes
		MaxBackups: 20,       //最多保留20个备份
		LocalTime:  true,
		Compress:   true, // 是否压缩 disabled by default
	})
	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoder), syncWriter, zap.NewAtomicLevelAt(level))
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(0))

	_logger[filePath] = logger.Sugar()

	return _logger[filePath]
}

//func main()  {
//	logger := GetLogger("/Users/tree/work/99_tree/03_github/log.log", "info")
//
//	logger.Info("log info ")
//
//	GetDefaultLogger().Info("aaaaaaaaaaa")
//}
