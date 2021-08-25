package timestamppb

import (
	"math"
	"math/big"
)

// Duplicate duplicates the Timestamp without any check.
func (x *Timestamp) Duplicate() *Timestamp {
	if x == nil {
		return nil
	}
	return &Timestamp{
		Seconds: x.Seconds,
		Nanos:   x.Nanos,
	}
}

// IsZero checks whether the Timestamp is zero.
// If the Timestamp is nil or zero as time.Time, it returns true. Otherwise, returns false.
func (x *Timestamp) IsZero() bool {
	if x == nil {
		return true
	}
	if !x.IsValid() || !x.AsTime().IsZero() {
		return false
	}
	return true
}

// AsNanos returns unix timestamp by nanoseconds as big.Int.
func (x *Timestamp) AsNanos() *big.Int {
	result := big.NewInt(x.GetSeconds())
	result.Mul(result, big.NewInt(1e9))
	result.Add(result, big.NewInt(int64(x.GetNanos())))
	return result
}

func (x *Timestamp) asAnyseconds(divider int64) int64 {
	b := x.AsNanos()
	b.Div(b, big.NewInt(divider))
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

// AsNanoseconds returns unix timestamp by nanoseconds as int64.
// If the result is out of range, it returns math.MaxInt64 or math.MinInt64.
func (x *Timestamp) AsNanoseconds() int64 {
	return x.asAnyseconds(1)
}

// AsMicroseconds returns unix timestamp by microseconds as int64.
// If the result is out of range, it returns math.MaxInt64 or math.MinInt64.
func (x *Timestamp) AsMicroseconds() int64 {
	return x.asAnyseconds(1e3)
}

// AsMilliseconds returns unix timestamp by milliseconds as int64.
// If the result is out of range, it returns math.MaxInt64 or math.MinInt64.
func (x *Timestamp) AsMilliseconds() int64 {
	return x.asAnyseconds(1e6)
}

// AsSeconds returns unix timestamp by seconds as int64.
// If the result is out of range, it returns math.MaxInt64 or math.MinInt64.
func (x *Timestamp) AsSeconds() int64 {
	return x.asAnyseconds(1e9)
}

func newByAnyseconds(d int64, multiplier int64) *Timestamp {
	b := big.NewInt(d)
	b = b.Mul(b, big.NewInt(multiplier))
	b, m := b.DivMod(b, big.NewInt(1e9), new(big.Int))
	secs, nanos := b.Int64(), m.Int64()
	if !b.IsInt64() {
		if b.Sign() >= 0 {
			secs = math.MaxInt64
		} else {
			secs = math.MinInt64
		}
	}
	return &Timestamp{
		Seconds: secs,
		Nanos:   int32(nanos),
	}
}

// NewByNanoseconds constructs a new Timestamp from the provided int64 unix timestamp by nanoseconds.
func NewByNanoseconds(d int64) *Timestamp {
	return newByAnyseconds(d, 1)
}

// NewByMicroseconds constructs a new Timestamp from the provided int64 unix timestamp by microseconds.
func NewByMicroseconds(d int64) *Timestamp {
	return newByAnyseconds(d, 1e3)
}

// NewByMilliseconds constructs a new Timestamp from the provided int64 unix timestamp by milliseconds.
func NewByMilliseconds(d int64) *Timestamp {
	return newByAnyseconds(d, 1e6)
}

// NewBySeconds constructs a new Timestamp from the provided int64 unix timestamp by seconds.
func NewBySeconds(d int64) *Timestamp {
	return newByAnyseconds(d, 1e9)
}
