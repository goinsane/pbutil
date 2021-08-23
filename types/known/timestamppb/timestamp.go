package timestamppb

import (
	"math"
	"math/big"
)

func (x *Timestamp) Duplicate() *Timestamp {
	if x == nil {
		return nil
	}
	return &Timestamp{
		Seconds: x.Seconds,
		Nanos:   x.Nanos,
	}
}

func (x *Timestamp) IsZero() bool {
	if x == nil {
		return true
	}
	if !x.IsValid() || !x.AsTime().IsZero() {
		return false
	}
	return true
}

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

func (x *Timestamp) AsNanoseconds() int64 {
	return x.asAnyseconds(1)
}

func (x *Timestamp) AsMicroseconds() int64 {
	return x.asAnyseconds(1e3)
}

func (x *Timestamp) AsMilliseconds() int64 {
	return x.asAnyseconds(1e6)
}

func (x *Timestamp) AsSeconds() int64 {
	return x.asAnyseconds(1e9)
}

func fromAnyseconds(d int64, multiplier int64) *Timestamp {
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

func FromNanoseconds(d int64) *Timestamp {
	return fromAnyseconds(d, 1)
}

func FromMicroseconds(d int64) *Timestamp {
	return fromAnyseconds(d, 1e3)
}

func FromMilliseconds(d int64) *Timestamp {
	return fromAnyseconds(d, 1e6)
}

func FromSeconds(d int64) *Timestamp {
	return fromAnyseconds(d, 1e9)
}
