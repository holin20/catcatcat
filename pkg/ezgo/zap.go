package ezgo

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// CloneLogger clones a new logger from base logger with new name and additional output path
// relative to zap output root.
func CloneLogger(baseLogger *zap.Logger, name string, outputPath string) *zap.Logger {
	derived := baseLogger
	if name != "" {
		derived = baseLogger.Named(name)
	}
	if outputPath != "" {
		derived = derived.WithOptions(zap.WrapCore(func(core zapcore.Core) zapcore.Core {
			// Create a new file core
			file, err := os.OpenFile(getZapLogPathRoot()+"/"+outputPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
			AssertNoErrorf(err, "Failed to create output path for new logger %s", name)
			fileSyncer := zapcore.AddSync(file)
			fileCore := zapcore.NewCore(
				zapcore.NewJSONEncoder(*getJsonEncoder()),
				fileSyncer,
				zapcore.InfoLevel,
			)
			// Combine the original core with the new file core
			return zapcore.NewTee(core, fileCore)
		}))
	}

	return derived
}

func CreateDefaultLogger(loggerName string) *zap.Logger {
	loc, _ := time.LoadLocation("America/Los_Angeles")
	logFilePath := fmt.Sprintf(
		"%s/%s_%s.txt",
		getZapLogPathRoot(),
		loggerName,
		time.Now().In(loc).Format("2006-01-02"),
	)
	config := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		//		Encoding:         "console",
		Encoding:         "json",
		EncoderConfig:    *getJsonEncoder(),
		OutputPaths:      []string{"stdout", logFilePath},
		ErrorOutputPaths: []string{"stderr", logFilePath},
	}

	return Must(config.Build())
}

func getZapLogPathRoot() string {
	return NonEmptyOr(os.Getenv("CATCATCAT_ZAP_OUTPUT_ROOT"), "logs")
}

func getJsonEncoder() *zapcore.EncoderConfig {
	location, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		panic(err)
	}

	// Custom Time Encoder for PST
	pstTimeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.In(location).Format(time.RFC3339))
	}

	return &zapcore.EncoderConfig{
		TimeKey:    "ts",
		LevelKey:   "level",
		MessageKey: "msg",
		NameKey:    "logger",
		// StacktraceKey:    "stacktrace",
		EncodeTime:       pstTimeEncoder, // Use the custom PST time encoder
		EncodeLevel:      zapcore.CapitalLevelEncoder,
		EncodeDuration:   zapcore.StringDurationEncoder,
		ConsoleSeparator: " | ",
	}
}
