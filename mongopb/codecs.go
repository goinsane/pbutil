package mongopb

import (
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var (
	primitiveNullType     = reflect.TypeOf(*new(primitive.Null))
	primitiveObjectIDType = reflect.TypeOf(primitive.ObjectID{})

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
	timestampType = reflect.TypeOf(new(timestamppb.Timestamp))

	objectIDType = reflect.TypeOf(new(ObjectID))

	goDurationType = reflect.TypeOf(*new(time.Duration))
	goTimeType     = reflect.TypeOf(*new(time.Time))
)

// RegisterAllCodecs registers all of implemented codecs.
func RegisterAllCodecs(rb *bsoncodec.RegistryBuilder) *bsoncodec.RegistryBuilder {
	rb = RegisterWrappersCodec(rb)
	rb = RegisterDurationCodec(rb)
	rb = RegisterTimestampCodec(rb)
	rb = RegisterObjectIDCodec(rb)
	return rb
}

func encodeNull(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter) error {
	enc, err := ec.LookupEncoder(primitiveNullType)
	if err != nil {
		return err
	}
	return enc.EncodeValue(ec, vw, reflect.New(primitiveNullType).Elem())
}

func decodeNull(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader) (bool, error) {
	dec, err := dc.LookupDecoder(primitiveNullType)
	if err != nil {
		return false, err
	}
	return dec.DecodeValue(dc, vr, reflect.New(primitiveNullType).Elem()) == nil, nil
}

// WrappersCodec is codec for protobuf wrappers.
type WrappersCodec struct {
}

// EncodeValue encodes protobuf wrappers value to BSON value.
func (c *WrappersCodec) EncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	if val.IsNil() {
		return encodeNull(ec, vw)
	}
	val = val.Elem().FieldByName("Value")
	enc, err := ec.LookupEncoder(val.Type())
	if err != nil {
		return err
	}
	return enc.EncodeValue(ec, vw, val)
}

// DecodeValue decodes BSON value to protobuf wrappers value.
func (c *WrappersCodec) DecodeValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	if vr.ReadNull() == nil {
		val.Set(reflect.New(val.Type()).Elem())
		return nil
	}
	val = val.Elem().FieldByName("Value")
	dec, err := dc.LookupDecoder(val.Type())
	if err != nil {
		return err
	}
	return dec.DecodeValue(dc, vr, val)
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

// DurationCodec is codec for protobuf type Duration.
type DurationCodec struct {
}

// EncodeValue encodes protobuf type Duration to BSON value.
func (c *DurationCodec) EncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	if val.IsNil() {
		return encodeNull(ec, vw)
	}
	enc, err := ec.LookupEncoder(goDurationType)
	if err != nil {
		return err
	}
	v := val.Interface().(*durationpb.Duration)
	if err = v.CheckValid(); err != nil {
		return err
	}
	return enc.EncodeValue(ec, vw, reflect.ValueOf(v.AsDuration()))
}

// DecodeValue decodes BSON value to protobuf type Duration.
func (c *DurationCodec) DecodeValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	if vr.ReadNull() == nil {
		val.Set(reflect.New(val.Type()).Elem())
		return nil
	}
	dec, err := dc.LookupDecoder(goDurationType)
	if err != nil {
		return err
	}
	var d time.Duration
	if err = dec.DecodeValue(dc, vr, reflect.ValueOf(&d).Elem()); err != nil {
		return err
	}
	val.Set(reflect.ValueOf(durationpb.New(d)))
	return nil
}

// RegisterDurationCodec registers DurationCodec.
func RegisterDurationCodec(rb *bsoncodec.RegistryBuilder) *bsoncodec.RegistryBuilder {
	durationCodecRef := new(DurationCodec)
	return rb.RegisterCodec(durationType, durationCodecRef)
}

// TimestampCodec is codec for protobuf type Timestamp.
type TimestampCodec struct {
}

// EncodeValue encodes protobuf type Timestamp to BSON value.
func (c *TimestampCodec) EncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	if val.IsNil() {
		return encodeNull(ec, vw)
	}
	enc, err := ec.LookupEncoder(goTimeType)
	if err != nil {
		return err
	}
	v := val.Interface().(*timestamppb.Timestamp)
	if err = v.CheckValid(); err != nil {
		return err
	}
	return enc.EncodeValue(ec, vw, reflect.ValueOf(v.AsTime().UTC()))
}

// DecodeValue decodes BSON value to protobuf type Timestamp.
func (c *TimestampCodec) DecodeValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	if vr.ReadNull() == nil {
		val.Set(reflect.New(val.Type()).Elem())
		return nil
	}
	dec, err := dc.LookupDecoder(goTimeType)
	if err != nil {
		return err
	}
	var t time.Time
	if err = dec.DecodeValue(dc, vr, reflect.ValueOf(&t).Elem()); err != nil {
		return err
	}
	var v *timestamppb.Timestamp
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

// ObjectIDCodec is codec for protobuf type ObjectID.
type ObjectIDCodec struct {
}

// EncodeValue encodes protobuf type ObjectID to BSON value.
func (c *ObjectIDCodec) EncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	if val.IsNil() {
		return encodeNull(ec, vw)
	}
	enc, err := ec.LookupEncoder(primitiveObjectIDType)
	if err != nil {
		return err
	}
	v := val.Interface().(*ObjectID)
	oid, err := primitive.ObjectIDFromHex(v.Value)
	if err != nil {
		return err
	}
	return enc.EncodeValue(ec, vw, reflect.ValueOf(oid))
}

// DecodeValue decodes BSON value to protobuf type ObjectID.
func (c *ObjectIDCodec) DecodeValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	if vr.ReadNull() == nil {
		val.Set(reflect.New(val.Type()).Elem())
		return nil
	}
	dec, err := dc.LookupDecoder(primitiveObjectIDType)
	if err != nil {
		return err
	}
	var oid primitive.ObjectID
	if err = dec.DecodeValue(dc, vr, reflect.ValueOf(&oid).Elem()); err != nil {
		return err
	}
	val.Set(reflect.ValueOf(&ObjectID{Value: oid.Hex()}))
	return nil
}

// RegisterObjectIDCodec registers ObjectIDCodec.
func RegisterObjectIDCodec(rb *bsoncodec.RegistryBuilder) *bsoncodec.RegistryBuilder {
	objectIDCodecRef := new(ObjectIDCodec)
	return rb.RegisterCodec(objectIDType, objectIDCodecRef)
}
