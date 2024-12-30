package ezgo

import "go.uber.org/zap"

type Scope struct {
	loggers []*zap.Logger
}

func NewScope(
	logger *zap.Logger,
) *Scope {
	return &Scope{
		loggers: []*zap.Logger{logger},
	}
}

func NewScopeWithDefaultLogger() (*Scope, error) {
	logger, err := zap.NewProduction()
	if IsErr(err) {
		return nil, NewCause(err, "NewZapLogger")
	}
	return NewScope(logger), nil
}

func (s *Scope) GetLogger() *zap.Logger {
	return lastElelemt(s.loggers, "logger")
}

func (s *Scope) Close() {
	for _, logger := range s.loggers {
		logger.Sync()
	}
}

func (s *Scope) WithLogger(logger *zap.Logger) *Scope {
	s.loggers = append(s.loggers, logger)
	return s
}

func lastElelemt[T any](slice []*T, proprty string) *T {
	Assertf(len(slice) > 0, "No property %s", proprty)
	return slice[len(slice)-1]
}
