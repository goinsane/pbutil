package pbutil

import (
	"math"
	"math/big"
	"time"

	"google.golang.org/protobuf/types/known/durationpb"
)

// DuplicateDuration duplicates the Duration without any check.
func DuplicateDuration(x *durationpb.Duration) *durationpb.Duration {
	if x == nil {
		return nil
	}
	return &durationpb.Duration{
		Seconds: x.Seconds,
		Nanos:   x.Nanos,
	}
}

// DurationAsNanos returns nanoseconds as big.Int.
func DurationAsNanos(x *durationpb.Duration) *big.Int {
	result := big.NewInt(x.GetSeconds())
	result.Mul(result, big.NewInt(1e9))
	result.Add(result, big.NewInt(int64(x.GetNanos())))
	return result
}

func durationAsAnyseconds(x *durationpb.Duration, nanosDivider int64) int64 {
	b := DurationAsNanos(x)
	b.Quo(b, big.NewInt(nanosDivider))
	result := b.Int64()
	if !b.IsInt64() {
		if b.Sign() >= 0 {
			result = math.MaxInt64
		} else {
			result = math.MinInt64
		}
	}
	return result
}

// DurationAsNanoseconds returns nanoseconds as int64.
// If the result is out of range, it returns math.MaxInt64 or math.MinInt64.
func DurationAsNanoseconds(x *durationpb.Duration) int64 {
	return durationAsAnyseconds(x, 1)
}

// DurationAsMicroseconds returns microseconds as int64.
// If the result is out of range, it returns math.MaxInt64 or math.MinInt64.
func DurationAsMicroseconds(x *durationpb.Duration) int64 {
	return durationAsAnyseconds(x, 1e3)
}

// DurationAsMilliseconds returns milliseconds as int64.
// If the result is out of range, it returns math.MaxInt64 or math.MinInt64.
func DurationAsMilliseconds(x *durationpb.Duration) int64 {
	return durationAsAnyseconds(x, 1e6)
}

// DurationAsSeconds returns seconds as int64.
// If the result is out of range, it returns math.MaxInt64 or math.MinInt64.
func DurationAsSeconds(x *durationpb.Duration) int64 {
	return durationAsAnyseconds(x, 1e9)
}

// NewDuration constructs a new Duration from the provided duration d.
func NewDuration(d time.Duration) *durationpb.Duration {
	return durationpb.New(d)
}

func newDurationByAnyseconds(d int64, nanosMultiplier int64) *durationpb.Duration {
	b := big.NewInt(d)
	b = b.Mul(b, big.NewInt(nanosMultiplier))
	b, m := b.QuoRem(b, big.NewInt(1e9), new(big.Int))
	secs, nanos := b.Int64(), m.Int64()
	if !b.IsInt64() {
		if b.Sign() >= 0 {
			secs = math.MaxInt64
		} else {
			secs = math.MinInt64
		}
	}
	return &durationpb.Duration{
		Seconds: secs,
		Nanos:   int32(nanos),
	}
}

// NewDurationByNanoseconds constructs a new Duration from the provided int64 by nanoseconds.
func NewDurationByNanoseconds(d int64) *durationpb.Duration {
	return newDurationByAnyseconds(d, 1)
}

// NewDurationByMicroseconds constructs a new Duration from the provided int64 by microseconds.
func NewDurationByMicroseconds(d int64) *durationpb.Duration {
	return newDurationByAnyseconds(d, 1e3)
}

// NewDurationByMilliseconds constructs a new Duration from the provided int64 by milliseconds.
func NewDurationByMilliseconds(d int64) *durationpb.Duration {
	return newDurationByAnyseconds(d, 1e6)
}

// NewDurationBySeconds constructs a new Duration from the provided int64 by seconds.
func NewDurationBySeconds(d int64) *durationpb.Duration {
	return newDurationByAnyseconds(d, 1e9)
}
