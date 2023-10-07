// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tensorflow/core/example/example_parser_configuration.proto

package tensorflow

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	framework "infer-microservices/common/tensorflow_gogofaster/core/framework"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type VarLenFeatureProto struct {
	Dtype                   framework.DataType `protobuf:"varint,1,opt,name=dtype,proto3,enum=tensorflow.DataType" json:"dtype,omitempty"`
	ValuesOutputTensorName  string             `protobuf:"bytes,2,opt,name=values_output_tensor_name,json=valuesOutputTensorName,proto3" json:"values_output_tensor_name,omitempty"`
	IndicesOutputTensorName string             `protobuf:"bytes,3,opt,name=indices_output_tensor_name,json=indicesOutputTensorName,proto3" json:"indices_output_tensor_name,omitempty"`
	ShapesOutputTensorName  string             `protobuf:"bytes,4,opt,name=shapes_output_tensor_name,json=shapesOutputTensorName,proto3" json:"shapes_output_tensor_name,omitempty"`
}

func (m *VarLenFeatureProto) Reset()         { *m = VarLenFeatureProto{} }
func (m *VarLenFeatureProto) String() string { return proto.CompactTextString(m) }
func (*VarLenFeatureProto) ProtoMessage()    {}
func (*VarLenFeatureProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_80ac05576ebb1f2e, []int{0}
}
func (m *VarLenFeatureProto) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VarLenFeatureProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VarLenFeatureProto.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VarLenFeatureProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VarLenFeatureProto.Merge(m, src)
}
func (m *VarLenFeatureProto) XXX_Size() int {
	return m.Size()
}
func (m *VarLenFeatureProto) XXX_DiscardUnknown() {
	xxx_messageInfo_VarLenFeatureProto.DiscardUnknown(m)
}

var xxx_messageInfo_VarLenFeatureProto proto.InternalMessageInfo

func (m *VarLenFeatureProto) GetDtype() framework.DataType {
	if m != nil {
		return m.Dtype
	}
	return framework.DataType_DT_INVALID
}

func (m *VarLenFeatureProto) GetValuesOutputTensorName() string {
	if m != nil {
		return m.ValuesOutputTensorName
	}
	return ""
}

func (m *VarLenFeatureProto) GetIndicesOutputTensorName() string {
	if m != nil {
		return m.IndicesOutputTensorName
	}
	return ""
}

func (m *VarLenFeatureProto) GetShapesOutputTensorName() string {
	if m != nil {
		return m.ShapesOutputTensorName
	}
	return ""
}

type FixedLenFeatureProto struct {
	Dtype                  framework.DataType          `protobuf:"varint,1,opt,name=dtype,proto3,enum=tensorflow.DataType" json:"dtype,omitempty"`
	Shape                  *framework.TensorShapeProto `protobuf:"bytes,2,opt,name=shape,proto3" json:"shape,omitempty"`
	DefaultValue           *framework.TensorProto      `protobuf:"bytes,3,opt,name=default_value,json=defaultValue,proto3" json:"default_value,omitempty"`
	ValuesOutputTensorName string                      `protobuf:"bytes,4,opt,name=values_output_tensor_name,json=valuesOutputTensorName,proto3" json:"values_output_tensor_name,omitempty"`
}

func (m *FixedLenFeatureProto) Reset()         { *m = FixedLenFeatureProto{} }
func (m *FixedLenFeatureProto) String() string { return proto.CompactTextString(m) }
func (*FixedLenFeatureProto) ProtoMessage()    {}
func (*FixedLenFeatureProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_80ac05576ebb1f2e, []int{1}
}
func (m *FixedLenFeatureProto) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FixedLenFeatureProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FixedLenFeatureProto.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FixedLenFeatureProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FixedLenFeatureProto.Merge(m, src)
}
func (m *FixedLenFeatureProto) XXX_Size() int {
	return m.Size()
}
func (m *FixedLenFeatureProto) XXX_DiscardUnknown() {
	xxx_messageInfo_FixedLenFeatureProto.DiscardUnknown(m)
}

var xxx_messageInfo_FixedLenFeatureProto proto.InternalMessageInfo

func (m *FixedLenFeatureProto) GetDtype() framework.DataType {
	if m != nil {
		return m.Dtype
	}
	return framework.DataType_DT_INVALID
}

func (m *FixedLenFeatureProto) GetShape() *framework.TensorShapeProto {
	if m != nil {
		return m.Shape
	}
	return nil
}

func (m *FixedLenFeatureProto) GetDefaultValue() *framework.TensorProto {
	if m != nil {
		return m.DefaultValue
	}
	return nil
}

func (m *FixedLenFeatureProto) GetValuesOutputTensorName() string {
	if m != nil {
		return m.ValuesOutputTensorName
	}
	return ""
}

type FeatureConfiguration struct {
	// Types that are valid to be assigned to Config:
	//	*FeatureConfiguration_FixedLenFeature
	//	*FeatureConfiguration_VarLenFeature
	Config isFeatureConfiguration_Config `protobuf_oneof:"config"`
}

func (m *FeatureConfiguration) Reset()         { *m = FeatureConfiguration{} }
func (m *FeatureConfiguration) String() string { return proto.CompactTextString(m) }
func (*FeatureConfiguration) ProtoMessage()    {}
func (*FeatureConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_80ac05576ebb1f2e, []int{2}
}
func (m *FeatureConfiguration) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FeatureConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FeatureConfiguration.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FeatureConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeatureConfiguration.Merge(m, src)
}
func (m *FeatureConfiguration) XXX_Size() int {
	return m.Size()
}
func (m *FeatureConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_FeatureConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_FeatureConfiguration proto.InternalMessageInfo

type isFeatureConfiguration_Config interface {
	isFeatureConfiguration_Config()
	MarshalTo([]byte) (int, error)
	Size() int
}

type FeatureConfiguration_FixedLenFeature struct {
	FixedLenFeature *FixedLenFeatureProto `protobuf:"bytes,1,opt,name=fixed_len_feature,json=fixedLenFeature,proto3,oneof" json:"fixed_len_feature,omitempty"`
}
type FeatureConfiguration_VarLenFeature struct {
	VarLenFeature *VarLenFeatureProto `protobuf:"bytes,2,opt,name=var_len_feature,json=varLenFeature,proto3,oneof" json:"var_len_feature,omitempty"`
}

func (*FeatureConfiguration_FixedLenFeature) isFeatureConfiguration_Config() {}
func (*FeatureConfiguration_VarLenFeature) isFeatureConfiguration_Config()   {}

func (m *FeatureConfiguration) GetConfig() isFeatureConfiguration_Config {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *FeatureConfiguration) GetFixedLenFeature() *FixedLenFeatureProto {
	if x, ok := m.GetConfig().(*FeatureConfiguration_FixedLenFeature); ok {
		return x.FixedLenFeature
	}
	return nil
}

func (m *FeatureConfiguration) GetVarLenFeature() *VarLenFeatureProto {
	if x, ok := m.GetConfig().(*FeatureConfiguration_VarLenFeature); ok {
		return x.VarLenFeature
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FeatureConfiguration) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FeatureConfiguration_FixedLenFeature)(nil),
		(*FeatureConfiguration_VarLenFeature)(nil),
	}
}

type ExampleParserConfiguration struct {
	FeatureMap map[string]*FeatureConfiguration `protobuf:"bytes,1,rep,name=feature_map,json=featureMap,proto3" json:"feature_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *ExampleParserConfiguration) Reset()         { *m = ExampleParserConfiguration{} }
func (m *ExampleParserConfiguration) String() string { return proto.CompactTextString(m) }
func (*ExampleParserConfiguration) ProtoMessage()    {}
func (*ExampleParserConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_80ac05576ebb1f2e, []int{3}
}
func (m *ExampleParserConfiguration) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExampleParserConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExampleParserConfiguration.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExampleParserConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExampleParserConfiguration.Merge(m, src)
}
func (m *ExampleParserConfiguration) XXX_Size() int {
	return m.Size()
}
func (m *ExampleParserConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_ExampleParserConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_ExampleParserConfiguration proto.InternalMessageInfo

func (m *ExampleParserConfiguration) GetFeatureMap() map[string]*FeatureConfiguration {
	if m != nil {
		return m.FeatureMap
	}
	return nil
}

func init() {
	proto.RegisterType((*VarLenFeatureProto)(nil), "tensorflow.VarLenFeatureProto")
	proto.RegisterType((*FixedLenFeatureProto)(nil), "tensorflow.FixedLenFeatureProto")
	proto.RegisterType((*FeatureConfiguration)(nil), "tensorflow.FeatureConfiguration")
	proto.RegisterType((*ExampleParserConfiguration)(nil), "tensorflow.ExampleParserConfiguration")
	proto.RegisterMapType((map[string]*FeatureConfiguration)(nil), "tensorflow.ExampleParserConfiguration.FeatureMapEntry")
}

func init() {
	proto.RegisterFile("tensorflow/core/example/example_parser_configuration.proto", fileDescriptor_80ac05576ebb1f2e)
}

var fileDescriptor_80ac05576ebb1f2e = []byte{
	// 522 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xb3, 0x49, 0x53, 0xd1, 0x31, 0x25, 0x60, 0x55, 0x6d, 0x6a, 0x21, 0xcb, 0x8a, 0x04,
	0x8a, 0x10, 0x4a, 0x24, 0x23, 0x55, 0xb4, 0x70, 0x0a, 0xb4, 0xca, 0x01, 0x4a, 0x64, 0xaa, 0xf6,
	0xb8, 0x5a, 0x92, 0x71, 0xb1, 0x6a, 0x7b, 0xad, 0xb5, 0x9d, 0x36, 0x6f, 0xc1, 0xf3, 0xf0, 0x04,
	0x1c, 0x7b, 0x42, 0x1c, 0x51, 0xf2, 0x00, 0x1c, 0xb8, 0x70, 0x44, 0xde, 0xb5, 0x5a, 0x27, 0x71,
	0xc3, 0x81, 0x53, 0xa2, 0x9d, 0xef, 0x9f, 0xd9, 0xff, 0xcf, 0x4e, 0xe0, 0x20, 0xc1, 0x30, 0xe6,
	0xc2, 0xf5, 0xf9, 0x65, 0x77, 0xc8, 0x05, 0x76, 0xf1, 0x8a, 0x05, 0x91, 0x7f, 0xf3, 0x49, 0x23,
	0x26, 0x62, 0x14, 0x74, 0xc8, 0x43, 0xd7, 0x3b, 0x4f, 0x05, 0x4b, 0x3c, 0x1e, 0x76, 0x22, 0xc1,
	0x13, 0xae, 0xc3, 0xad, 0xd6, 0x78, 0xbe, 0xd8, 0xc7, 0x15, 0x2c, 0xc0, 0x4b, 0x2e, 0x2e, 0xba,
	0xaa, 0x42, 0xe3, 0xcf, 0x2c, 0x42, 0xa5, 0x34, 0x9e, 0xfe, 0x8b, 0xce, 0xb9, 0x27, 0x2b, 0xb8,
	0x49, 0x84, 0xb1, 0xc2, 0x5a, 0xbf, 0x09, 0xe8, 0xa7, 0x4c, 0xbc, 0xc3, 0xf0, 0x08, 0x59, 0x92,
	0x0a, 0x1c, 0xc8, 0xfb, 0x3d, 0x83, 0xfa, 0x28, 0xc3, 0x9a, 0xc4, 0x22, 0xed, 0x07, 0xf6, 0x56,
	0xe7, 0xb6, 0x5b, 0xe7, 0x2d, 0x4b, 0xd8, 0xc9, 0x24, 0x42, 0x47, 0x21, 0xfa, 0x3e, 0xec, 0x8e,
	0x99, 0x9f, 0x62, 0x4c, 0x79, 0x9a, 0x44, 0x69, 0x42, 0xf3, 0x5b, 0x87, 0x2c, 0xc0, 0x66, 0xd5,
	0x22, 0xed, 0x0d, 0x67, 0x5b, 0x01, 0x1f, 0x64, 0xfd, 0x44, 0x96, 0x8f, 0x59, 0x80, 0xfa, 0x2b,
	0x30, 0xbc, 0x70, 0xe4, 0x0d, 0xcb, 0xb5, 0x35, 0xa9, 0xdd, 0xc9, 0x89, 0x25, 0xf1, 0x3e, 0xec,
	0xca, 0x60, 0x4a, 0xb5, 0x6b, 0x6a, 0xae, 0x02, 0x16, 0xa5, 0xad, 0x5f, 0x04, 0xb6, 0x8e, 0xbc,
	0x2b, 0x1c, 0xfd, 0x8f, 0x6f, 0x1b, 0xea, 0xb2, 0xbd, 0xf4, 0xa8, 0xd9, 0x8f, 0x8b, 0xac, 0x9a,
	0xf5, 0x31, 0x2b, 0xcb, 0xc6, 0x8e, 0x42, 0xf5, 0xd7, 0xb0, 0x39, 0x42, 0x97, 0xa5, 0x7e, 0x42,
	0x65, 0x24, 0xd2, 0xa3, 0x66, 0xef, 0x2c, 0x6b, 0x95, 0xec, 0x7e, 0x4e, 0x9f, 0x66, 0xf0, 0xea,
	0xa4, 0xd7, 0x56, 0x25, 0xdd, 0xfa, 0x9a, 0x39, 0x56, 0x4e, 0xdf, 0x14, 0xdf, 0xa3, 0x7e, 0x0c,
	0x8f, 0xdc, 0x2c, 0x09, 0xea, 0x63, 0x48, 0x5d, 0x45, 0x48, 0xf7, 0x9a, 0x6d, 0x15, 0x6f, 0x55,
	0x16, 0x57, 0xbf, 0xe2, 0x34, 0xdc, 0xf9, 0x73, 0xbd, 0x0f, 0x8d, 0x31, 0x13, 0x73, 0xdd, 0x54,
	0x3e, 0x66, 0xb1, 0xdb, 0xf2, 0x93, 0xeb, 0x57, 0x9c, 0xcd, 0x71, 0xf1, 0xb4, 0x77, 0x0f, 0xd6,
	0xd5, 0xea, 0xb4, 0xbe, 0x13, 0x30, 0x0e, 0xd5, 0x52, 0x0d, 0xe4, 0x4e, 0xcd, 0x5b, 0x38, 0x03,
	0x2d, 0x1f, 0x45, 0x03, 0x16, 0x35, 0x89, 0x55, 0x6b, 0x6b, 0xf6, 0x5e, 0x71, 0xdc, 0xdd, 0xe2,
	0x4e, 0x3e, 0xed, 0x3d, 0x8b, 0x0e, 0xc3, 0x44, 0x4c, 0x1c, 0x70, 0x6f, 0x0e, 0x0c, 0x0a, 0x8d,
	0x85, 0xb2, 0xfe, 0x10, 0x6a, 0x17, 0x38, 0x91, 0x01, 0x6d, 0x38, 0xd9, 0x57, 0x7d, 0x0f, 0xea,
	0xea, 0xa7, 0xac, 0x96, 0x84, 0x56, 0x92, 0xb8, 0xa3, 0xf0, 0x83, 0xea, 0x4b, 0xd2, 0x3b, 0xfb,
	0x36, 0x35, 0xc9, 0xf5, 0xd4, 0x24, 0x3f, 0xa7, 0x26, 0xf9, 0x32, 0x33, 0x2b, 0xd7, 0x33, 0xb3,
	0xf2, 0x63, 0x66, 0x56, 0x60, 0x9b, 0x8b, 0xf3, 0x62, 0xa7, 0xfc, 0x3f, 0xa5, 0x67, 0xdd, 0x6d,
	0x45, 0xc6, 0x18, 0x0f, 0xc8, 0x1f, 0x42, 0x3e, 0xad, 0xcb, 0xed, 0x7e, 0xf1, 0x37, 0x00, 0x00,
	0xff, 0xff, 0x41, 0xd0, 0xb0, 0x84, 0xa4, 0x04, 0x00, 0x00,
}

func (m *VarLenFeatureProto) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VarLenFeatureProto) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VarLenFeatureProto) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ShapesOutputTensorName) > 0 {
		i -= len(m.ShapesOutputTensorName)
		copy(dAtA[i:], m.ShapesOutputTensorName)
		i = encodeVarintExampleParserConfiguration(dAtA, i, uint64(len(m.ShapesOutputTensorName)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.IndicesOutputTensorName) > 0 {
		i -= len(m.IndicesOutputTensorName)
		copy(dAtA[i:], m.IndicesOutputTensorName)
		i = encodeVarintExampleParserConfiguration(dAtA, i, uint64(len(m.IndicesOutputTensorName)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ValuesOutputTensorName) > 0 {
		i -= len(m.ValuesOutputTensorName)
		copy(dAtA[i:], m.ValuesOutputTensorName)
		i = encodeVarintExampleParserConfiguration(dAtA, i, uint64(len(m.ValuesOutputTensorName)))
		i--
		dAtA[i] = 0x12
	}
	if m.Dtype != 0 {
		i = encodeVarintExampleParserConfiguration(dAtA, i, uint64(m.Dtype))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *FixedLenFeatureProto) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FixedLenFeatureProto) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FixedLenFeatureProto) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ValuesOutputTensorName) > 0 {
		i -= len(m.ValuesOutputTensorName)
		copy(dAtA[i:], m.ValuesOutputTensorName)
		i = encodeVarintExampleParserConfiguration(dAtA, i, uint64(len(m.ValuesOutputTensorName)))
		i--
		dAtA[i] = 0x22
	}
	if m.DefaultValue != nil {
		{
			size, err := m.DefaultValue.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintExampleParserConfiguration(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Shape != nil {
		{
			size, err := m.Shape.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintExampleParserConfiguration(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Dtype != 0 {
		i = encodeVarintExampleParserConfiguration(dAtA, i, uint64(m.Dtype))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *FeatureConfiguration) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FeatureConfiguration) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FeatureConfiguration) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Config != nil {
		{
			size := m.Config.Size()
			i -= size
			if _, err := m.Config.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *FeatureConfiguration_FixedLenFeature) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FeatureConfiguration_FixedLenFeature) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.FixedLenFeature != nil {
		{
			size, err := m.FixedLenFeature.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintExampleParserConfiguration(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *FeatureConfiguration_VarLenFeature) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FeatureConfiguration_VarLenFeature) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.VarLenFeature != nil {
		{
			size, err := m.VarLenFeature.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintExampleParserConfiguration(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *ExampleParserConfiguration) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExampleParserConfiguration) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExampleParserConfiguration) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FeatureMap) > 0 {
		for k := range m.FeatureMap {
			v := m.FeatureMap[k]
			baseI := i
			if v != nil {
				{
					size, err := v.MarshalToSizedBuffer(dAtA[:i])
					if err != nil {
						return 0, err
					}
					i -= size
					i = encodeVarintExampleParserConfiguration(dAtA, i, uint64(size))
				}
				i--
				dAtA[i] = 0x12
			}
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintExampleParserConfiguration(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintExampleParserConfiguration(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintExampleParserConfiguration(dAtA []byte, offset int, v uint64) int {
	offset -= sovExampleParserConfiguration(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *VarLenFeatureProto) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Dtype != 0 {
		n += 1 + sovExampleParserConfiguration(uint64(m.Dtype))
	}
	l = len(m.ValuesOutputTensorName)
	if l > 0 {
		n += 1 + l + sovExampleParserConfiguration(uint64(l))
	}
	l = len(m.IndicesOutputTensorName)
	if l > 0 {
		n += 1 + l + sovExampleParserConfiguration(uint64(l))
	}
	l = len(m.ShapesOutputTensorName)
	if l > 0 {
		n += 1 + l + sovExampleParserConfiguration(uint64(l))
	}
	return n
}

func (m *FixedLenFeatureProto) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Dtype != 0 {
		n += 1 + sovExampleParserConfiguration(uint64(m.Dtype))
	}
	if m.Shape != nil {
		l = m.Shape.Size()
		n += 1 + l + sovExampleParserConfiguration(uint64(l))
	}
	if m.DefaultValue != nil {
		l = m.DefaultValue.Size()
		n += 1 + l + sovExampleParserConfiguration(uint64(l))
	}
	l = len(m.ValuesOutputTensorName)
	if l > 0 {
		n += 1 + l + sovExampleParserConfiguration(uint64(l))
	}
	return n
}

func (m *FeatureConfiguration) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Config != nil {
		n += m.Config.Size()
	}
	return n
}

func (m *FeatureConfiguration_FixedLenFeature) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FixedLenFeature != nil {
		l = m.FixedLenFeature.Size()
		n += 1 + l + sovExampleParserConfiguration(uint64(l))
	}
	return n
}
func (m *FeatureConfiguration_VarLenFeature) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.VarLenFeature != nil {
		l = m.VarLenFeature.Size()
		n += 1 + l + sovExampleParserConfiguration(uint64(l))
	}
	return n
}
func (m *ExampleParserConfiguration) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.FeatureMap) > 0 {
		for k, v := range m.FeatureMap {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovExampleParserConfiguration(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovExampleParserConfiguration(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovExampleParserConfiguration(uint64(mapEntrySize))
		}
	}
	return n
}

func sovExampleParserConfiguration(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozExampleParserConfiguration(x uint64) (n int) {
	return sovExampleParserConfiguration(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *VarLenFeatureProto) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExampleParserConfiguration
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: VarLenFeatureProto: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VarLenFeatureProto: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dtype", wireType)
			}
			m.Dtype = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExampleParserConfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Dtype |= framework.DataType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValuesOutputTensorName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExampleParserConfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthExampleParserConfiguration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExampleParserConfiguration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValuesOutputTensorName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IndicesOutputTensorName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExampleParserConfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthExampleParserConfiguration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExampleParserConfiguration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IndicesOutputTensorName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShapesOutputTensorName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExampleParserConfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthExampleParserConfiguration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExampleParserConfiguration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ShapesOutputTensorName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExampleParserConfiguration(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExampleParserConfiguration
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *FixedLenFeatureProto) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExampleParserConfiguration
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FixedLenFeatureProto: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FixedLenFeatureProto: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dtype", wireType)
			}
			m.Dtype = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExampleParserConfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Dtype |= framework.DataType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Shape", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExampleParserConfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthExampleParserConfiguration
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthExampleParserConfiguration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Shape == nil {
				m.Shape = &framework.TensorShapeProto{}
			}
			if err := m.Shape.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultValue", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExampleParserConfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthExampleParserConfiguration
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthExampleParserConfiguration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DefaultValue == nil {
				m.DefaultValue = &framework.TensorProto{}
			}
			if err := m.DefaultValue.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValuesOutputTensorName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExampleParserConfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthExampleParserConfiguration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExampleParserConfiguration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValuesOutputTensorName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExampleParserConfiguration(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExampleParserConfiguration
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *FeatureConfiguration) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExampleParserConfiguration
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FeatureConfiguration: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FeatureConfiguration: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FixedLenFeature", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExampleParserConfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthExampleParserConfiguration
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthExampleParserConfiguration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &FixedLenFeatureProto{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Config = &FeatureConfiguration_FixedLenFeature{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VarLenFeature", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExampleParserConfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthExampleParserConfiguration
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthExampleParserConfiguration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &VarLenFeatureProto{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Config = &FeatureConfiguration_VarLenFeature{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExampleParserConfiguration(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExampleParserConfiguration
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ExampleParserConfiguration) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExampleParserConfiguration
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ExampleParserConfiguration: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExampleParserConfiguration: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeatureMap", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExampleParserConfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthExampleParserConfiguration
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthExampleParserConfiguration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.FeatureMap == nil {
				m.FeatureMap = make(map[string]*FeatureConfiguration)
			}
			var mapkey string
			var mapvalue *FeatureConfiguration
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowExampleParserConfiguration
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowExampleParserConfiguration
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthExampleParserConfiguration
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthExampleParserConfiguration
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowExampleParserConfiguration
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthExampleParserConfiguration
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthExampleParserConfiguration
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &FeatureConfiguration{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipExampleParserConfiguration(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthExampleParserConfiguration
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.FeatureMap[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExampleParserConfiguration(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExampleParserConfiguration
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipExampleParserConfiguration(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowExampleParserConfiguration
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowExampleParserConfiguration
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowExampleParserConfiguration
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthExampleParserConfiguration
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupExampleParserConfiguration
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthExampleParserConfiguration
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthExampleParserConfiguration        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowExampleParserConfiguration          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupExampleParserConfiguration = fmt.Errorf("proto: unexpected end of group")
)
