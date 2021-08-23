package timestamppb

func (x *Timestamp) Duplicate() *Timestamp {
	if x == nil {
		return nil
	}
	if x.IsZero() {
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
	if x.IsValid() && x.AsTime().IsZero() {
		return true
	}
	return false
}
