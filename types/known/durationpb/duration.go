package durationpb

import (
	"math"
	"math/big"
)

// Duplicate duplicates the Duration without any check.
func (x *Duration) Duplicate() *Duration {
	if x == nil {
		return nil
	}
	return &Duration{
		Seconds: x.Seconds,
		Nanos:   x.Nanos,
	}
}

// AsNanos returns nanoseconds as big.Int.
func (x *Duration) AsNanos() *big.Int {
	result := big.NewInt(x.GetSeconds())
	result.Mul(result, big.NewInt(1e9))
	result.Add(result, big.NewInt(int64(x.GetNanos())))
	return result
}

func (x *Duration) asAnyseconds(divider int64) int64 {
	b := x.AsNanos()
	b.Quo(b, big.NewInt(divider))
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

// AsNanoseconds returns nanoseconds as int64.
// If the result is out of range, it returns math.MaxInt64 or math.MinInt64.
func (x *Duration) AsNanoseconds() int64 {
	return x.asAnyseconds(1)
}

// AsMicroseconds returns microseconds as int64.
// If the result is out of range, it returns math.MaxInt64 or math.MinInt64.
func (x *Duration) AsMicroseconds() int64 {
	return x.asAnyseconds(1e3)
}

// AsMilliseconds returns milliseconds as int64.
// If the result is out of range, it returns math.MaxInt64 or math.MinInt64.
func (x *Duration) AsMilliseconds() int64 {
	return x.asAnyseconds(1e6)
}

// AsSeconds returns seconds as int64.
// If the result is out of range, it returns math.MaxInt64 or math.MinInt64.
func (x *Duration) AsSeconds() int64 {
	return x.asAnyseconds(1e9)
}

func newByAnyseconds(d int64, multiplier int64) *Duration {
	b := big.NewInt(d)
	b = b.Mul(b, big.NewInt(multiplier))
	b, m := b.QuoRem(b, big.NewInt(1e9), new(big.Int))
	secs, nanos := b.Int64(), m.Int64()
	if !b.IsInt64() {
		if b.Sign() >= 0 {
			secs = math.MaxInt64
		} else {
			secs = math.MinInt64
		}
	}
	return &Duration{
		Seconds: secs,
		Nanos:   int32(nanos),
	}
}

// NewByNanoseconds constructs a new Duration from the provided int64 by nanoseconds.
func NewByNanoseconds(d int64) *Duration {
	return newByAnyseconds(d, 1)
}

// NewByFromMicroseconds constructs a new Duration from the provided int64 by microseconds.
func NewByFromMicroseconds(d int64) *Duration {
	return newByAnyseconds(d, 1e3)
}

// NewByMilliseconds constructs a new Duration from the provided int64 by milliseconds.
func NewByMilliseconds(d int64) *Duration {
	return newByAnyseconds(d, 1e6)
}

// NewBySeconds constructs a new Duration from the provided int64 by seconds.
func NewBySeconds(d int64) *Duration {
	return newByAnyseconds(d, 1e9)
}
