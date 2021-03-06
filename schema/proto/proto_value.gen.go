// Code generated by gen/schema_value.go. DO NOT EDIT.

package proto

// Schema.Value union type from https://github.com/capnproto/capnproto/blob/master/c++/src/capnp/schema.capnp
// Version info: 57a4ca5af5a7f55b768a9d9d6655250bffb1257f capnproto (v0.8.0)
type Value map[string]interface{}

func (v Value) Which() ValueType {
	if _, ok := v["void"]; ok {
		return ValueVoid
	} else if _, ok := v["bool"]; ok {
		return ValueBool
	} else if _, ok := v["int8"]; ok {
		return ValueInt8
	} else if _, ok := v["int16"]; ok {
		return ValueInt16
	} else if _, ok := v["int32"]; ok {
		return ValueInt32
	} else if _, ok := v["int64"]; ok {
		return ValueInt64
	} else if _, ok := v["uint8"]; ok {
		return ValueUint8
	} else if _, ok := v["uint16"]; ok {
		return ValueUint16
	} else if _, ok := v["uint32"]; ok {
		return ValueUint32
	} else if _, ok := v["uint64"]; ok {
		return ValueUint64
	} else if _, ok := v["float32"]; ok {
		return ValueFloat32
	} else if _, ok := v["float64"]; ok {
		return ValueFloat64
	} else if _, ok := v["text"]; ok {
		return ValueText
	} else if _, ok := v["data"]; ok {
		return ValueData
	} else if _, ok := v["list"]; ok {
		return ValueList
	} else if _, ok := v["enum"]; ok {
		return ValueEnum
	} else if _, ok := v["struct"]; ok {
		return ValueStruct
	} else if _, ok := v["interface"]; ok {
		return ValueInterface
	} else if _, ok := v["anyPointer"]; ok {
		return ValueAnyPointer
	} else {
		return ValueUnknown
	}
}

type ValueType uint16

const (
	ValueVoid ValueType = iota
	ValueBool
	ValueInt8
	ValueInt16
	ValueInt32
	ValueInt64
	ValueUint8
	ValueUint16
	ValueUint32
	ValueUint64
	ValueFloat32
	ValueFloat64
	ValueText
	ValueData
	ValueList
	ValueEnum
	ValueStruct
	ValueInterface
	ValueAnyPointer
	ValueUnknown
)

func (t ValueType) String() string {
	switch t {
	case ValueVoid:
		return "void"
	case ValueBool:
		return "bool"
	case ValueInt8:
		return "int8"
	case ValueInt16:
		return "int16"
	case ValueInt32:
		return "int32"
	case ValueInt64:
		return "int64"
	case ValueUint8:
		return "uint8"
	case ValueUint16:
		return "uint16"
	case ValueUint32:
		return "uint32"
	case ValueUint64:
		return "uint64"
	case ValueFloat32:
		return "float32"
	case ValueFloat64:
		return "float64"
	case ValueText:
		return "text"
	case ValueData:
		return "data"
	case ValueList:
		return "list"
	case ValueEnum:
		return "enum"
	case ValueStruct:
		return "struct"
	case ValueInterface:
		return "interface"
	case ValueAnyPointer:
		return "anyPointer"
	default:
		return ""
	}
}
