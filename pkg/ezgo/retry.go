package ezgo

import "time"

func Retry[T any](
	fn func() (T, error),
	shouldRetry func(T, error) bool,
	maxAttempts int,
	retryInterval time.Duration,
) (T, error) {
	var lastError error
	for i := 0; i < maxAttempts; i++ {
		result, lastError := fn()
		if shouldRetry(result, lastError) {
			time.Sleep(retryInterval)
			continue
		}
		return result, lastError
	}
	var zero T
	return zero, NewCausef(lastError, "maxAttempts(%d) reached", maxAttempts)
}

func RetryOnErr[T any](
	fn func() (T, error),
	maxAttempts int,
	retryInterval time.Duration,
) (T, error) {
	return Retry(
		fn,
		func(_ T, err error) bool {
			return IsErr(err)
		},
		maxAttempts,
		retryInterval,
	)
}

func RetryNoDelay[T any](
	fn func() (T, error),
	shouldRetry func(T, error) bool,
	maxAttempts int,
) (T, error) {
	return Retry(fn, shouldRetry, maxAttempts, 0)
}
