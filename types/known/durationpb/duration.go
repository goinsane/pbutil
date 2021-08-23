package durationpb

func (x *Duration) Duplicate() *Duration {
	if x == nil {
		return nil
	}
	return &Duration{
		Seconds: x.Seconds,
		Nanos:   x.Nanos,
	}
}
