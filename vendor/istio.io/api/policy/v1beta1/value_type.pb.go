// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: policy/v1beta1/value_type.proto

package v1beta1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import strconv "strconv"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// ValueType describes the types that values in the Istio system can take. These
// are used to describe the type of Attributes at run time, describe the type of
// the result of evaluating an expression, and to describe the runtime type of
// fields of other descriptors.
type ValueType int32

const (
	// Invalid, default value.
	VALUE_TYPE_UNSPECIFIED ValueType = 0
	// An undiscriminated variable-length string.
	STRING ValueType = 1
	// An undiscriminated 64-bit signed integer.
	INT64 ValueType = 2
	// An undiscriminated 64-bit floating-point value.
	DOUBLE ValueType = 3
	// An undiscriminated boolean value.
	BOOL ValueType = 4
	// A point in time.
	TIMESTAMP ValueType = 5
	// An IP address.
	IP_ADDRESS ValueType = 6
	// An email address.
	EMAIL_ADDRESS ValueType = 7
	// A URI.
	URI ValueType = 8
	// A DNS name.
	DNS_NAME ValueType = 9
	// A span between two points in time.
	DURATION ValueType = 10
	// A map string -> string, typically used by headers.
	STRING_MAP ValueType = 11
)

var ValueType_name = map[int32]string{
	0:  "VALUE_TYPE_UNSPECIFIED",
	1:  "STRING",
	2:  "INT64",
	3:  "DOUBLE",
	4:  "BOOL",
	5:  "TIMESTAMP",
	6:  "IP_ADDRESS",
	7:  "EMAIL_ADDRESS",
	8:  "URI",
	9:  "DNS_NAME",
	10: "DURATION",
	11: "STRING_MAP",
}
var ValueType_value = map[string]int32{
	"VALUE_TYPE_UNSPECIFIED": 0,
	"STRING":                 1,
	"INT64":                  2,
	"DOUBLE":                 3,
	"BOOL":                   4,
	"TIMESTAMP":              5,
	"IP_ADDRESS":             6,
	"EMAIL_ADDRESS":          7,
	"URI":                    8,
	"DNS_NAME":               9,
	"DURATION":               10,
	"STRING_MAP":             11,
}

func (ValueType) EnumDescriptor() ([]byte, []int) { return fileDescriptorValueType, []int{0} }

func init() {
	proto.RegisterEnum("istio.policy.v1beta1.ValueType", ValueType_name, ValueType_value)
}
func (x ValueType) String() string {
	s, ok := ValueType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

func init() { proto.RegisterFile("policy/v1beta1/value_type.proto", fileDescriptorValueType) }

var fileDescriptorValueType = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x3f, 0x4b, 0xc3, 0x40,
	0x18, 0x87, 0x73, 0xf6, 0x6f, 0x5e, 0xad, 0x9c, 0x87, 0x38, 0x28, 0x9c, 0xbb, 0x43, 0x43, 0x51,
	0xdc, 0xaf, 0xe6, 0x94, 0x83, 0xe4, 0x12, 0x72, 0x97, 0x82, 0x2e, 0x47, 0x2b, 0x19, 0x02, 0x85,
	0x04, 0x8d, 0x85, 0x6e, 0x7e, 0x04, 0x3f, 0x86, 0xdf, 0xc1, 0x2f, 0xe0, 0xd8, 0xd1, 0xd1, 0x9e,
	0x8b, 0x63, 0x3f, 0x82, 0xa4, 0x11, 0xc1, 0xf1, 0x7d, 0x78, 0x78, 0xf9, 0xf1, 0xc0, 0x69, 0x59,
	0xcc, 0xf3, 0xfb, 0xa5, 0xb7, 0x18, 0xcd, 0xb2, 0x6a, 0x3a, 0xf2, 0x16, 0xd3, 0xf9, 0x53, 0x66,
	0xaa, 0x65, 0x99, 0x0d, 0xcb, 0x87, 0xa2, 0x2a, 0xc8, 0x61, 0xfe, 0x58, 0xe5, 0xc5, 0xb0, 0xd1,
	0x86, 0xbf, 0xda, 0xd9, 0x1b, 0x02, 0x77, 0x52, 0xab, 0x7a, 0x59, 0x66, 0xe4, 0x18, 0x8e, 0x26,
	0x2c, 0x48, 0xb9, 0xd1, 0xb7, 0x31, 0x37, 0xa9, 0x54, 0x31, 0xbf, 0x12, 0xd7, 0x82, 0xfb, 0xd8,
	0x21, 0x00, 0x5d, 0xa5, 0x13, 0x21, 0x6f, 0x30, 0x22, 0x2e, 0x74, 0x84, 0xd4, 0x97, 0x17, 0x78,
	0xa7, 0xc6, 0x7e, 0x94, 0x8e, 0x03, 0x8e, 0x5b, 0xa4, 0x0f, 0xed, 0x71, 0x14, 0x05, 0xb8, 0x4d,
	0x06, 0xe0, 0x6a, 0x11, 0x72, 0xa5, 0x59, 0x18, 0xe3, 0x0e, 0xd9, 0x07, 0x10, 0xb1, 0x61, 0xbe,
	0x9f, 0x70, 0xa5, 0x70, 0x97, 0x1c, 0xc0, 0x80, 0x87, 0x4c, 0x04, 0x7f, 0xa8, 0x47, 0x7a, 0xd0,
	0x4a, 0x13, 0x81, 0xfb, 0x64, 0x0f, 0xfa, 0xbe, 0x54, 0x46, 0xb2, 0x90, 0x63, 0x77, 0x7b, 0xa5,
	0x09, 0xd3, 0x22, 0x92, 0x18, 0xea, 0x3f, 0xcd, 0x06, 0x13, 0xb2, 0x18, 0xef, 0x8e, 0xc5, 0x6a,
	0x4d, 0x9d, 0x8f, 0x35, 0x75, 0x36, 0x6b, 0x8a, 0x9e, 0x2d, 0x45, 0xaf, 0x96, 0xa2, 0x77, 0x4b,
	0xd1, 0xca, 0x52, 0xf4, 0x69, 0x29, 0xfa, 0xb6, 0xd4, 0xd9, 0x58, 0x8a, 0x5e, 0xbe, 0xa8, 0x73,
	0x77, 0xd2, 0x14, 0xc8, 0x0b, 0x6f, 0x5a, 0xe6, 0xde, 0xff, 0x5e, 0xb3, 0xee, 0xb6, 0xd2, 0xf9,
	0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x41, 0xab, 0x82, 0x48, 0x01, 0x00, 0x00,
}
