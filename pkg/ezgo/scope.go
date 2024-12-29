package ezgo

import "go.uber.org/zap"

type Scope struct {
	logger *zap.Logger
}

func NewScope(
	logger *zap.Logger,
) *Scope {
	return &Scope{
		logger: logger,
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
	return s.logger
}

func (s *Scope) Close() {
	s.logger.Sync()
}
