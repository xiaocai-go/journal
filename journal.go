// 日志格式化记录器

package journal

import (
	"errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
)

var UnknownEncoder error = errors.New("unknown encoder")
var UnknownLevel error = errors.New("unknown level")

// Level 日志等级
type Level int8

// Encoder 编码器类型
type Encoder int8

const (
	DebugLevel Level = iota - 1
	InfoLevel
	WarnLevel
	ErrorLevel
	DPanicLevel
	PanicLevel
	FatalLevel
)

const (
	ConsoleEncoder Encoder = iota
	JSONEncoder
)

type Options struct {
	Writer io.Writer
	// 编码器类型
	Encoder Encoder
	// 日志级别
	Level Level
	// 跳过行结尾（默认：\n 结尾）
	SkipLineEnding bool
}

func NewOptions(w io.Writer) *Options {
	return &Options{
		Writer:         w,
		Encoder:        ConsoleEncoder,
		Level:          DebugLevel,
		SkipLineEnding: false,
	}
}

type Journal struct {
	o *Options
	*zap.Logger
}

func New(o *Options) *Journal {
	return &Journal{
		o:      o,
		Logger: instance(o),
	}
}

// 编码器配置
func encoderConfig() *zapcore.EncoderConfig {
	conf := zap.NewProductionEncoderConfig()
	conf.TimeKey = "time"
	conf.MessageKey = "message"
	conf.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	return &conf
}

// 编码器
func encoder(e Encoder, c *zapcore.EncoderConfig) zapcore.Encoder {
	switch e {
	case ConsoleEncoder:
		return zapcore.NewConsoleEncoder(*c)
	case JSONEncoder:
		return zapcore.NewJSONEncoder(*c)
	default:
		panic(UnknownEncoder)
	}
}

// 写入器
func writer(o *Options) zapcore.WriteSyncer {
	return zapcore.AddSync(o.Writer)
}

// instance 日志实例
func instance(o *Options) *zap.Logger {
	if o.Level < DebugLevel || o.Level > FatalLevel {
		panic(UnknownLevel)
	}
	encoderConf := encoderConfig()
	encoderConf.SkipLineEnding = o.SkipLineEnding
	core := zapcore.NewCore(encoder(o.Encoder, encoderConf), writer(o), zapcore.Level(o.Level))
	return zap.New(core)
}
