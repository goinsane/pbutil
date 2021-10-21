package pbutilmongo

import (
	"reflect"
	"time"

	durationpb2 "github.com/goinsane/pbutil/types/known/durationpb"
	timestamppb2 "github.com/goinsane/pbutil/types/known/timestamppb"
	"github.com/golang/protobuf/ptypes/wrappers"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	boolValueType   = reflect.TypeOf(wrappers.BoolValue{})
	bytesValueType  = reflect.TypeOf(wrappers.BytesValue{})
	doubleValueType = reflect.TypeOf(wrappers.DoubleValue{})
	floatValueType  = reflect.TypeOf(wrappers.FloatValue{})
	int32ValueType  = reflect.TypeOf(wrappers.Int32Value{})
	int64ValueType  = reflect.TypeOf(wrappers.Int64Value{})
	stringValueType = reflect.TypeOf(wrappers.StringValue{})
	uint32ValueType = reflect.TypeOf(wrappers.UInt32Value{})
	uint64ValueType = reflect.TypeOf(wrappers.UInt64Value{})

	timestampType  = reflect.TypeOf(timestamppb.Timestamp{})
	timestamp2Type = reflect.TypeOf(timestamppb2.Timestamp{})
	goTimeType     = reflect.TypeOf(*new(time.Time))

	durationType   = reflect.TypeOf(durationpb.Duration{})
	duration2Type  = reflect.TypeOf(durationpb2.Duration{})
	goDurationType = reflect.TypeOf(*new(time.Duration))
)

type TimestampCodec struct {
}

// EncodeValue encodes Timestamp value to BSON value
func (e *TimestampCodec) EncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {

	v := val.Interface().(timestamp.Timestamp)
	t, err := ptypes.Timestamp(&v)
	if err != nil {
		return err
	}
	enc, err := ec.LookupEncoder(timeType)
	if err != nil {
		return err
	}
	return enc.EncodeValue(ec, vw, reflect.ValueOf(t.In(time.UTC)))
}

// DecodeValue decodes BSON value to Timestamp value
func (e *TimestampCodec) DecodeValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	enc, err := dc.LookupDecoder(timeType)
	if err != nil {
		return err
	}
	var t time.Time
	if err = enc.DecodeValue(dc, vr, reflect.ValueOf(&t).Elem()); err != nil {
		return err
	}
	ts, err := ptypes.TimestampProto(t.In(time.UTC))
	if err != nil {
		return err
	}
	val.Set(reflect.ValueOf(*ts))
	return nil
}

func panicInvalidCodecType() {
	panic("invalid codec type")
}
