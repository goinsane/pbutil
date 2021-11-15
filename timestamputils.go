package pbutil

import (
	"math"
	"math/big"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// DuplicateTimestamp duplicates the Timestamp without any check.
func DuplicateTimestamp(x *timestamppb.Timestamp) *timestamppb.Timestamp {
	if x == nil {
		return nil
	}
	return &timestamppb.Timestamp{
		Seconds: x.Seconds,
		Nanos:   x.Nanos,
	}
}

// IsTimestampZero checks whether the Timestamp is zero.
// If the Timestamp is nil or zero as time.Time, it returns true. Otherwise, returns false.
func IsTimestampZero(x *timestamppb.Timestamp) bool {
	if x == nil {
		return true
	}
	if x.IsValid() && x.AsTime().IsZero() {
		return true
	}
	return false
}

// TimestampAsNanos returns unix timestamp by nanoseconds as big.Int.
func TimestampAsNanos(x *timestamppb.Timestamp) *big.Int {
	result := big.NewInt(x.GetSeconds())
	result.Mul(result, big.NewInt(1e9))
	result.Add(result, big.NewInt(int64(x.GetNanos())))
	return result
}

func timestampAsAnyseconds(x *timestamppb.Timestamp, nanosDivider int64) int64 {
	b := TimestampAsNanos(x)
	b.Div(b, big.NewInt(nanosDivider))
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

// TimestampAsNanoseconds returns unix timestamp by nanoseconds as int64.
// If the result is out of range, it returns math.MaxInt64 or math.MinInt64.
func TimestampAsNanoseconds(x *timestamppb.Timestamp) int64 {
	return timestampAsAnyseconds(x, 1)
}

// TimestampAsMicroseconds returns unix timestamp by microseconds as int64.
// If the result is out of range, it returns math.MaxInt64 or math.MinInt64.
func TimestampAsMicroseconds(x *timestamppb.Timestamp) int64 {
	return timestampAsAnyseconds(x, 1e3)
}

// TimestampAsMilliseconds returns unix timestamp by milliseconds as int64.
// If the result is out of range, it returns math.MaxInt64 or math.MinInt64.
func TimestampAsMilliseconds(x *timestamppb.Timestamp) int64 {
	return timestampAsAnyseconds(x, 1e6)
}

// TimestampAsSeconds returns unix timestamp by seconds as int64.
// If the result is out of range, it returns math.MaxInt64 or math.MinInt64.
func TimestampAsSeconds(x *timestamppb.Timestamp) int64 {
	return timestampAsAnyseconds(x, 1e9)
}

// NewTimestamp constructs a new Timestamp from the provided time t.
// If t is zero, it returns nil.
func NewTimestamp(t time.Time) *timestamppb.Timestamp {
	if t.IsZero() {
		return nil
	}
	return timestamppb.New(t)
}

func newTimestampByAnyseconds(d int64, nanosMultiplier int64) *timestamppb.Timestamp {
	b := big.NewInt(d)
	b = b.Mul(b, big.NewInt(nanosMultiplier))
	b, m := b.DivMod(b, big.NewInt(1e9), new(big.Int))
	secs, nanos := b.Int64(), m.Int64()
	if !b.IsInt64() {
		if b.Sign() >= 0 {
			secs = math.MaxInt64
		} else {
			secs = math.MinInt64
		}
	}
	return &timestamppb.Timestamp{
		Seconds: secs,
		Nanos:   int32(nanos),
	}
}

// NewTimestampByNanoseconds constructs a new Timestamp from the provided int64 unix timestamp by nanoseconds.
func NewTimestampByNanoseconds(d int64) *timestamppb.Timestamp {
	return newTimestampByAnyseconds(d, 1)
}

// NewTimestampByMicroseconds constructs a new Timestamp from the provided int64 unix timestamp by microseconds.
func NewTimestampByMicroseconds(d int64) *timestamppb.Timestamp {
	return newTimestampByAnyseconds(d, 1e3)
}

// NewTimestampByMilliseconds constructs a new Timestamp from the provided int64 unix timestamp by milliseconds.
func NewTimestampByMilliseconds(d int64) *timestamppb.Timestamp {
	return newTimestampByAnyseconds(d, 1e6)
}

// NewTimestampBySeconds constructs a new Timestamp from the provided int64 unix timestamp by seconds.
func NewTimestampBySeconds(d int64) *timestamppb.Timestamp {
	return newTimestampByAnyseconds(d, 1e9)
}
