package pbutilmongo

import (
	"reflect"
	"time"

	durationpb2 "github.com/goinsane/pbutil/types/known/durationpb"
	timestamppb2 "github.com/goinsane/pbutil/types/known/timestamppb"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var (
	boolValueType   = reflect.TypeOf(new(wrapperspb.BoolValue))
	bytesValueType  = reflect.TypeOf(new(wrapperspb.BytesValue))
	doubleValueType = reflect.TypeOf(new(wrapperspb.DoubleValue))
	floatValueType  = reflect.TypeOf(new(wrapperspb.FloatValue))
	int32ValueType  = reflect.TypeOf(new(wrapperspb.Int32Value))
	int64ValueType  = reflect.TypeOf(new(wrapperspb.Int64Value))
	stringValueType = reflect.TypeOf(new(wrapperspb.StringValue))
	uint32ValueType = reflect.TypeOf(new(wrapperspb.UInt32Value))
	uint64ValueType = reflect.TypeOf(new(wrapperspb.UInt64Value))

	durationType  = reflect.TypeOf(new(durationpb.Duration))
	duration2Type = reflect.TypeOf(new(durationpb2.Duration))

	timestampType  = reflect.TypeOf(new(timestamppb.Timestamp))
	timestamp2Type = reflect.TypeOf(new(timestamppb2.Timestamp))

	goDurationType = reflect.TypeOf(*new(time.Duration))
	goTimeType     = reflect.TypeOf(*new(time.Time))
)

// WrappersCodec is codec for protobuf wrappers.
type WrappersCodec struct {
}

// EncodeValue encodes protobuf wrappers value to BSON value.
func (c *WrappersCodec) EncodeValue(ectx bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	val = val.Elem().FieldByName("Value")
	enc, err := ectx.LookupEncoder(val.Type())
	if err != nil {
		return err
	}
	return enc.EncodeValue(ectx, vw, val)
}

// DecodeValue decodes BSON value to protobuf wrappers value.
func (c *WrappersCodec) DecodeValue(ectx bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	val = val.Elem().FieldByName("Value")
	dec, err := ectx.LookupDecoder(val.Type())
	if err != nil {
		return err
	}
	return dec.DecodeValue(ectx, vr, val)
}

// RegisterWrappersCodec registers WrappersCodec.
func RegisterWrappersCodec(rb *bsoncodec.RegistryBuilder) *bsoncodec.RegistryBuilder {
	wrappersCodecRef := new(WrappersCodec)
	return rb.RegisterCodec(boolValueType, wrappersCodecRef).
		RegisterCodec(bytesValueType, wrappersCodecRef).
		RegisterCodec(doubleValueType, wrappersCodecRef).
		RegisterCodec(floatValueType, wrappersCodecRef).
		RegisterCodec(int32ValueType, wrappersCodecRef).
		RegisterCodec(int64ValueType, wrappersCodecRef).
		RegisterCodec(stringValueType, wrappersCodecRef).
		RegisterCodec(uint32ValueType, wrappersCodecRef).
		RegisterCodec(uint64ValueType, wrappersCodecRef)
}

// TimestampCodec is codec for protobuf type Timestamp.
type TimestampCodec struct {
}

// EncodeValue encodes protobuf type Timestamp to BSON value.
func (c *TimestampCodec) EncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	enc, err := ec.LookupEncoder(goTimeType)
	if err != nil {
		return err
	}
	var t time.Time
	var v *timestamppb.Timestamp
	v = val.Interface().(*timestamppb.Timestamp)
	if v != nil {
		if err = v.CheckValid(); err != nil {
			return err
		}
		t = v.AsTime().UTC()
	}
	return enc.EncodeValue(ec, vw, reflect.ValueOf(t))
}

// DecodeValue decodes BSON value to protobuf type Timestamp.
func (c *TimestampCodec) DecodeValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	dec, err := dc.LookupDecoder(goTimeType)
	if err != nil {
		return err
	}
	var t time.Time
	var v *timestamppb.Timestamp
	if err = dec.DecodeValue(dc, vr, reflect.ValueOf(&t).Elem()); err != nil {
		return err
	}
	if !t.IsZero() {
		v = timestamppb.New(t)
	}
	val.Set(reflect.ValueOf(v))
	return nil
}

// RegisterTimestampCodec registers TimestampCodec.
func RegisterTimestampCodec(rb *bsoncodec.RegistryBuilder) *bsoncodec.RegistryBuilder {
	timestampCodecRef := new(TimestampCodec)
	return rb.RegisterCodec(timestampType, timestampCodecRef)
}
