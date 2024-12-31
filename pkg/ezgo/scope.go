package ezgo

import "go.uber.org/zap"

type Scope struct {
	parent *Scope
	// properties of this scope that is immutable
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
	return findFirstNonNilPropoerty(s, "logger", func(s *Scope) *zap.Logger {
		return s.logger
	})
}

func (s *Scope) Close() {
	for ; s != nil; s = s.parent {
		s.logger.Sync()
	}
}

func (s *Scope) WithLogger(logger *zap.Logger) *Scope {
	return &Scope{
		parent: s,
		logger: logger,
	}
}

func findFirstNonNilPropoerty[T any](s *Scope, propertyName string, getProperty func(*Scope) *T) *T {
	for ; s != nil; s = s.parent {
		if p := getProperty(s); p != nil {
			return p
		}
	}
	AssertAlwaysf("Nil property %s", propertyName)
	return nil
}
