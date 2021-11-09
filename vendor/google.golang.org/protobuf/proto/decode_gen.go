// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by generate-types. DO NOT EDIT.

package proto

import (
	"math"
	"unicode/utf8"

	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/internal/errors"
	"google.golang.org/protobuf/internal/strs"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// unmarshalScalar decodes a value of the given kind.
//
// Message values are decoded into a []byte which aliases the input data.
func (o UnmarshalOptions) unmarshalScalar(b []byte, wtyp protowire.Type, fd protoreflect.FieldDescriptor) (val protoreflect.Value, n int, err error) {
	switch fd.Kind() {
	case protoreflect.BoolKind:
		if wtyp != protowire.VarintType {
			return val, 0, errUnknown
		}
		v, n := protowire.ConsumeVarint(b)
		if n < 0 {
			return val, 0, errDecode
		}
		return protoreflect.ValueOfBool(protowire.DecodeBool(v)), n, nil
	case protoreflect.EnumKind:
		if wtyp != protowire.VarintType {
			return val, 0, errUnknown
		}
		v, n := protowire.ConsumeVarint(b)
		if n < 0 {
			return val, 0, errDecode
		}
		return protoreflect.ValueOfEnum(protoreflect.EnumNumber(v)), n, nil
	case protoreflect.Int32Kind:
		if wtyp != protowire.VarintType {
			return val, 0, errUnknown
		}
		v, n := protowire.ConsumeVarint(b)
		if n < 0 {
			return val, 0, errDecode
		}
		return protoreflect.ValueOfInt32(int32(v)), n, nil
	case protoreflect.Sint32Kind:
		if wtyp != protowire.VarintType {
			return val, 0, errUnknown
		}
		v, n := protowire.ConsumeVarint(b)
		if n < 0 {
			return val, 0, errDecode
		}
		return protoreflect.ValueOfInt32(int32(protowire.DecodeZigZag(v & math.MaxUint32))), n, nil
	case protoreflect.Uint32Kind:
		if wtyp != protowire.VarintType {
			return val, 0, errUnknown
		}
		v, n := protowire.ConsumeVarint(b)
		if n < 0 {
			return val, 0, errDecode
		}
		return protoreflect.ValueOfUint32(uint32(v)), n, nil
	case protoreflect.Int64Kind:
		if wtyp != protowire.VarintType {
			return val, 0, errUnknown
		}
		v, n := protowire.ConsumeVarint(b)
		if n < 0 {
			return val, 0, errDecode
		}
		return protoreflect.ValueOfInt64(int64(v)), n, nil
	case protoreflect.Sint64Kind:
		if wtyp != protowire.VarintType {
			return val, 0, errUnknown
		}
		v, n := protowire.ConsumeVarint(b)
		if n < 0 {
			return val, 0, errDecode
		}
		return protoreflect.ValueOfInt64(protowire.DecodeZigZag(v)), n, nil
	case protoreflect.Uint64Kind:
		if wtyp != protowire.VarintType {
			return val, 0, errUnknown
		}
		v, n := protowire.ConsumeVarint(b)
		if n < 0 {
			return val, 0, errDecode
		}
		return protoreflect.ValueOfUint64(v), n, nil
	case protoreflect.Sfixed32Kind:
		if wtyp != protowire.Fixed32Type {
			return val, 0, errUnknown
		}
		v, n := protowire.ConsumeFixed32(b)
		if n < 0 {
			return val, 0, errDecode
		}
		return protoreflect.ValueOfInt32(int32(v)), n, nil
	case protoreflect.Fixed32Kind:
		if wtyp != protowire.Fixed32Type {
			return val, 0, errUnknown
		}
		v, n := protowire.ConsumeFixed32(b)
		if n < 0 {
			return val, 0, errDecode
		}
		return protoreflect.ValueOfUint32(uint32(v)), n, nil
	case protoreflect.FloatKind:
		if wtyp != protowire.Fixed32Type {
			return val, 0, errUnknown
		}
		v, n := protowire.ConsumeFixed32(b)
		if n < 0 {
			return val, 0, errDecode
		}
		return protoreflect.ValueOfFloat32(math.Float32frombits(uint32(v))), n, nil
	case protoreflect.Sfixed64Kind:
		if wtyp != protowire.Fixed64Type {
			return val, 0, errUnknown
		}
		v, n := protowire.ConsumeFixed64(b)
		if n < 0 {
			return val, 0, errDecode
		}
		return protoreflect.ValueOfInt64(int64(v)), n, nil
	case protoreflect.Fixed64Kind:
		if wtyp != protowire.Fixed64Type {
			return val, 0, errUnknown
		}
		v, n := protowire.ConsumeFixed64(b)
		if n < 0 {
			return val, 0, errDecode
		}
		return protoreflect.ValueOfUint64(v), n, nil
	case protoreflect.DoubleKind:
		if wtyp != protowire.Fixed64Type {
			return val, 0, errUnknown
		}
		v, n := protowire.ConsumeFixed64(b)
		if n < 0 {
			return val, 0, errDecode
		}
		return protoreflect.ValueOfFloat64(math.Float64frombits(v)), n, nil
	case protoreflect.StringKind:
		if wtyp != protowire.BytesType {
			return val, 0, errUnknown
		}
		v, n := protowire.ConsumeBytes(b)
		if n < 0 {
			return val, 0, errDecode
		}
		if strs.EnforceUTF8(fd) && !utf8.Valid(v) {
			return protoreflect.Value{}, 0, errors.InvalidUTF8(string(fd.FullName()))
		}
		return protoreflect.ValueOfString(string(v)), n, nil
	case protoreflect.BytesKind:
		if wtyp != protowire.BytesType {
			return val, 0, errUnknown
		}
		v, n := protowire.ConsumeBytes(b)
		if n < 0 {
			return val, 0, errDecode
		}
		return protoreflect.ValueOfBytes(append(emptyBuf[:], v...)), n, nil
	case protoreflect.MessageKind:
		if wtyp != protowire.BytesType {
			return val, 0, errUnknown
		}
		v, n := protowire.ConsumeBytes(b)
		if n < 0 {
			return val, 0, errDecode
		}
		return protoreflect.ValueOfBytes(v), n, nil
	case protoreflect.GroupKind:
		if wtyp != protowire.StartGroupType {
			return val, 0, errUnknown
		}
		v, n := protowire.ConsumeGroup(fd.Number(), b)
		if n < 0 {
			return val, 0, errDecode
		}
		return protoreflect.ValueOfBytes(v), n, nil
	default:
		return val, 0, errUnknown
	}
}

func (o UnmarshalOptions) unmarshalList(b []byte, wtyp protowire.Type, list protoreflect.List, fd protoreflect.FieldDescriptor) (n int, err error) {
	switch fd.Kind() {
	case protoreflect.BoolKind:
		if wtyp == protowire.BytesType {
			buf, n := protowire.ConsumeBytes(b)
			if n < 0 {
				return 0, errDecode
			}
			for len(buf) > 0 {
				v, n := protowire.ConsumeVarint(buf)
				if n < 0 {
					return 0, errDecode
				}
				buf = buf[n:]
				list.Append(protoreflect.ValueOfBool(protowire.DecodeBool(v)))
			}
			return n, nil
		}
		if wtyp != protowire.VarintType {
			return 0, errUnknown
		}
		v, n := protowire.ConsumeVarint(b)
		if n < 0 {
			return 0, errDecode
		}
		list.Append(protoreflect.ValueOfBool(protowire.DecodeBool(v)))
		return n, nil
	case protoreflect.EnumKind:
		if wtyp == protowire.BytesType {
			buf, n := protowire.ConsumeBytes(b)
			if n < 0 {
				return 0, errDecode
			}
			for len(buf) > 0 {
				v, n := protowire.ConsumeVarint(buf)
				if n < 0 {
					return 0, errDecode
				}
				buf = buf[n:]
				list.Append(protoreflect.ValueOfEnum(protoreflect.EnumNumber(v)))
			}
			return n, nil
		}
		if wtyp != protowire.VarintType {
			return 0, errUnknown
		}
		v, n := protowire.ConsumeVarint(b)
		if n < 0 {
			return 0, errDecode
		}
		list.Append(protoreflect.ValueOfEnum(protoreflect.EnumNumber(v)))
		return n, nil
	case protoreflect.Int32Kind:
		if wtyp == protowire.BytesType {
			buf, n := protowire.ConsumeBytes(b)
			if n < 0 {
				return 0, errDecode
			}
			for len(buf) > 0 {
				v, n := protowire.ConsumeVarint(buf)
				if n < 0 {
					return 0, errDecode
				}
				buf = buf[n:]
				list.Append(protoreflect.ValueOfInt32(int32(v)))
			}
			return n, nil
		}
		if wtyp != protowire.VarintType {
			return 0, errUnknown
		}
		v, n := protowire.ConsumeVarint(b)
		if n < 0 {
			return 0, errDecode
		}
		list.Append(protoreflect.ValueOfInt32(int32(v)))
		return n, nil
	case protoreflect.Sint32Kind:
		if wtyp == protowire.BytesType {
			buf, n := protowire.ConsumeBytes(b)
			if n < 0 {
				return 0, errDecode
			}
			for len(buf) > 0 {
				v, n := protowire.ConsumeVarint(buf)
				if n < 0 {
					return 0, errDecode
				}
				buf = buf[n:]
				list.Append(protoreflect.ValueOfInt32(int32(protowire.DecodeZigZag(v & math.MaxUint32))))
			}
			return n, nil
		}
		if wtyp != protowire.VarintType {
			return 0, errUnknown
		}
		v, n := protowire.ConsumeVarint(b)
		if n < 0 {
			return 0, errDecode
		}
		list.Append(protoreflect.ValueOfInt32(int32(protowire.DecodeZigZag(v & math.MaxUint32))))
		return n, nil
	case protoreflect.Uint32Kind:
		if wtyp == protowire.BytesType {
			buf, n := protowire.ConsumeBytes(b)
			if n < 0 {
				return 0, errDecode
			}
			for len(buf) > 0 {
				v, n := protowire.ConsumeVarint(buf)
				if n < 0 {
					return 0, errDecode
				}
				buf = buf[n:]
				list.Append(protoreflect.ValueOfUint32(uint32(v)))
			}
			return n, nil
		}
		if wtyp != protowire.VarintType {
			return 0, errUnknown
		}
		v, n := protowire.ConsumeVarint(b)
		if n < 0 {
			return 0, errDecode
		}
		list.Append(protoreflect.ValueOfUint32(uint32(v)))
		return n, nil
	case protoreflect.Int64Kind:
		if wtyp == protowire.BytesType {
			buf, n := protowire.ConsumeBytes(b)
			if n < 0 {
				return 0, errDecode
			}
			for len(buf) > 0 {
				v, n := protowire.ConsumeVarint(buf)
				if n < 0 {
					return 0, errDecode
				}
				buf = buf[n:]
				list.Append(protoreflect.ValueOfInt64(int64(v)))
			}
			return n, nil
		}
		if wtyp != protowire.VarintType {
			return 0, errUnknown
		}
		v, n := protowire.ConsumeVarint(b)
		if n < 0 {
			return 0, errDecode
		}
		list.Append(protoreflect.ValueOfInt64(int64(v)))
		return n, nil
	case protoreflect.Sint64Kind:
		if wtyp == protowire.BytesType {
			buf, n := protowire.ConsumeBytes(b)
			if n < 0 {
				return 0, errDecode
			}
			for len(buf) > 0 {
				v, n := protowire.ConsumeVarint(buf)
				if n < 0 {
					return 0, errDecode
				}
				buf = buf[n:]
				list.Append(protoreflect.ValueOfInt64(protowire.DecodeZigZag(v)))
			}
			return n, nil
		}
		if wtyp != protowire.VarintType {
			return 0, errUnknown
		}
		v, n := protowire.ConsumeVarint(b)
		if n < 0 {
			return 0, errDecode
		}
		list.Append(protoreflect.ValueOfInt64(protowire.DecodeZigZag(v)))
		return n, nil
	case protoreflect.Uint64Kind:
		if wtyp == protowire.BytesType {
			buf, n := protowire.ConsumeBytes(b)
			if n < 0 {
				return 0, errDecode
			}
			for len(buf) > 0 {
				v, n := protowire.ConsumeVarint(buf)
				if n < 0 {
					return 0, errDecode
				}
				buf = buf[n:]
				list.Append(protoreflect.ValueOfUint64(v))
			}
			return n, nil
		}
		if wtyp != protowire.VarintType {
			return 0, errUnknown
		}
		v, n := protowire.ConsumeVarint(b)
		if n < 0 {
			return 0, errDecode
		}
		list.Append(protoreflect.ValueOfUint64(v))
		return n, nil
	case protoreflect.Sfixed32Kind:
		if wtyp == protowire.BytesType {
			buf, n := protowire.ConsumeBytes(b)
			if n < 0 {
				return 0, errDecode
			}
			for len(buf) > 0 {
				v, n := protowire.ConsumeFixed32(buf)
				if n < 0 {
					return 0, errDecode
				}
				buf = buf[n:]
				list.Append(protoreflect.ValueOfInt32(int32(v)))
			}
			return n, nil
		}
		if wtyp != protowire.Fixed32Type {
			return 0, errUnknown
		}
		v, n := protowire.ConsumeFixed32(b)
		if n < 0 {
			return 0, errDecode
		}
		list.Append(protoreflect.ValueOfInt32(int32(v)))
		return n, nil
	case protoreflect.Fixed32Kind:
		if wtyp == protowire.BytesType {
			buf, n := protowire.ConsumeBytes(b)
			if n < 0 {
				return 0, errDecode
			}
			for len(buf) > 0 {
				v, n := protowire.ConsumeFixed32(buf)
				if n < 0 {
					return 0, errDecode
				}
				buf = buf[n:]
				list.Append(protoreflect.ValueOfUint32(uint32(v)))
			}
			return n, nil
		}
		if wtyp != protowire.Fixed32Type {
			return 0, errUnknown
		}
		v, n := protowire.ConsumeFixed32(b)
		if n < 0 {
			return 0, errDecode
		}
		list.Append(protoreflect.ValueOfUint32(uint32(v)))
		return n, nil
	case protoreflect.FloatKind:
		if wtyp == protowire.BytesType {
			buf, n := protowire.ConsumeBytes(b)
			if n < 0 {
				return 0, errDecode
			}
			for len(buf) > 0 {
				v, n := protowire.ConsumeFixed32(buf)
				if n < 0 {
					return 0, errDecode
				}
				buf = buf[n:]
				list.Append(protoreflect.ValueOfFloat32(math.Float32frombits(uint32(v))))
			}
			return n, nil
		}
		if wtyp != protowire.Fixed32Type {
			return 0, errUnknown
		}
		v, n := protowire.ConsumeFixed32(b)
		if n < 0 {
			return 0, errDecode
		}
		list.Append(protoreflect.ValueOfFloat32(math.Float32frombits(uint32(v))))
		return n, nil
	case protoreflect.Sfixed64Kind:
		if wtyp == protowire.BytesType {
			buf, n := protowire.ConsumeBytes(b)
			if n < 0 {
				return 0, errDecode
			}
			for len(buf) > 0 {
				v, n := protowire.ConsumeFixed64(buf)
				if n < 0 {
					return 0, errDecode
				}
				buf = buf[n:]
				list.Append(protoreflect.ValueOfInt64(int64(v)))
			}
			return n, nil
		}
		if wtyp != protowire.Fixed64Type {
			return 0, errUnknown
		}
		v, n := protowire.ConsumeFixed64(b)
		if n < 0 {
			return 0, errDecode
		}
		list.Append(protoreflect.ValueOfInt64(int64(v)))
		return n, nil
	case protoreflect.Fixed64Kind:
		if wtyp == protowire.BytesType {
			buf, n := protowire.ConsumeBytes(b)
			if n < 0 {
				return 0, errDecode
			}
			for len(buf) > 0 {
				v, n := protowire.ConsumeFixed64(buf)
				if n < 0 {
					return 0, errDecode
				}
				buf = buf[n:]
				list.Append(protoreflect.ValueOfUint64(v))
			}
			return n, nil
		}
		if wtyp != protowire.Fixed64Type {
			return 0, errUnknown
		}
		v, n := protowire.ConsumeFixed64(b)
		if n < 0 {
			return 0, errDecode
		}
		list.Append(protoreflect.ValueOfUint64(v))
		return n, nil
	case protoreflect.DoubleKind:
		if wtyp == protowire.BytesType {
			buf, n := protowire.ConsumeBytes(b)
			if n < 0 {
				return 0, errDecode
			}
			for len(buf) > 0 {
				v, n := protowire.ConsumeFixed64(buf)
				if n < 0 {
					return 0, errDecode
				}
				buf = buf[n:]
				list.Append(protoreflect.ValueOfFloat64(math.Float64frombits(v)))
			}
			return n, nil
		}
		if wtyp != protowire.Fixed64Type {
			return 0, errUnknown
		}
		v, n := protowire.ConsumeFixed64(b)
		if n < 0 {
			return 0, errDecode
		}
		list.Append(protoreflect.ValueOfFloat64(math.Float64frombits(v)))
		return n, nil
	case protoreflect.StringKind:
		if wtyp != protowire.BytesType {
			return 0, errUnknown
		}
		v, n := protowire.ConsumeBytes(b)
		if n < 0 {
			return 0, errDecode
		}
		if strs.EnforceUTF8(fd) && !utf8.Valid(v) {
			return 0, errors.InvalidUTF8(string(fd.FullName()))
		}
		list.Append(protoreflect.ValueOfString(string(v)))
		return n, nil
	case protoreflect.BytesKind:
		if wtyp != protowire.BytesType {
			return 0, errUnknown
		}
		v, n := protowire.ConsumeBytes(b)
		if n < 0 {
			return 0, errDecode
		}
		list.Append(protoreflect.ValueOfBytes(append(emptyBuf[:], v...)))
		return n, nil
	case protoreflect.MessageKind:
		if wtyp != protowire.BytesType {
			return 0, errUnknown
		}
		v, n := protowire.ConsumeBytes(b)
		if n < 0 {
			return 0, errDecode
		}
		m := list.NewElement()
		if err := o.unmarshalMessage(v, m.Message()); err != nil {
			return 0, err
		}
		list.Append(m)
		return n, nil
	case protoreflect.GroupKind:
		if wtyp != protowire.StartGroupType {
			return 0, errUnknown
		}
		v, n := protowire.ConsumeGroup(fd.Number(), b)
		if n < 0 {
			return 0, errDecode
		}
		m := list.NewElement()
		if err := o.unmarshalMessage(v, m.Message()); err != nil {
			return 0, err
		}
		list.Append(m)
		return n, nil
	default:
		return 0, errUnknown
	}
}

// We append to an empty array rather than a nil []byte to get non-nil zero-length byte slices.
var emptyBuf [0]byte
