// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/automl/v1beta1/text.proto

package automl

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Dataset metadata for classification.
type TextClassificationDatasetMetadata struct {
	// Required.
	// Type of the classification problem.
	ClassificationType   ClassificationType `protobuf:"varint,1,opt,name=classification_type,json=classificationType,proto3,enum=google.cloud.automl.v1beta1.ClassificationType" json:"classification_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *TextClassificationDatasetMetadata) Reset()         { *m = TextClassificationDatasetMetadata{} }
func (m *TextClassificationDatasetMetadata) String() string { return proto.CompactTextString(m) }
func (*TextClassificationDatasetMetadata) ProtoMessage()    {}
func (*TextClassificationDatasetMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1121cf231f416fd, []int{0}
}

func (m *TextClassificationDatasetMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextClassificationDatasetMetadata.Unmarshal(m, b)
}
func (m *TextClassificationDatasetMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextClassificationDatasetMetadata.Marshal(b, m, deterministic)
}
func (m *TextClassificationDatasetMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextClassificationDatasetMetadata.Merge(m, src)
}
func (m *TextClassificationDatasetMetadata) XXX_Size() int {
	return xxx_messageInfo_TextClassificationDatasetMetadata.Size(m)
}
func (m *TextClassificationDatasetMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_TextClassificationDatasetMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_TextClassificationDatasetMetadata proto.InternalMessageInfo

func (m *TextClassificationDatasetMetadata) GetClassificationType() ClassificationType {
	if m != nil {
		return m.ClassificationType
	}
	return ClassificationType_CLASSIFICATION_TYPE_UNSPECIFIED
}

// Model metadata that is specific to text classification.
type TextClassificationModelMetadata struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TextClassificationModelMetadata) Reset()         { *m = TextClassificationModelMetadata{} }
func (m *TextClassificationModelMetadata) String() string { return proto.CompactTextString(m) }
func (*TextClassificationModelMetadata) ProtoMessage()    {}
func (*TextClassificationModelMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1121cf231f416fd, []int{1}
}

func (m *TextClassificationModelMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextClassificationModelMetadata.Unmarshal(m, b)
}
func (m *TextClassificationModelMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextClassificationModelMetadata.Marshal(b, m, deterministic)
}
func (m *TextClassificationModelMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextClassificationModelMetadata.Merge(m, src)
}
func (m *TextClassificationModelMetadata) XXX_Size() int {
	return xxx_messageInfo_TextClassificationModelMetadata.Size(m)
}
func (m *TextClassificationModelMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_TextClassificationModelMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_TextClassificationModelMetadata proto.InternalMessageInfo

// Dataset metadata that is specific to text extraction
type TextExtractionDatasetMetadata struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TextExtractionDatasetMetadata) Reset()         { *m = TextExtractionDatasetMetadata{} }
func (m *TextExtractionDatasetMetadata) String() string { return proto.CompactTextString(m) }
func (*TextExtractionDatasetMetadata) ProtoMessage()    {}
func (*TextExtractionDatasetMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1121cf231f416fd, []int{2}
}

func (m *TextExtractionDatasetMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextExtractionDatasetMetadata.Unmarshal(m, b)
}
func (m *TextExtractionDatasetMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextExtractionDatasetMetadata.Marshal(b, m, deterministic)
}
func (m *TextExtractionDatasetMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextExtractionDatasetMetadata.Merge(m, src)
}
func (m *TextExtractionDatasetMetadata) XXX_Size() int {
	return xxx_messageInfo_TextExtractionDatasetMetadata.Size(m)
}
func (m *TextExtractionDatasetMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_TextExtractionDatasetMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_TextExtractionDatasetMetadata proto.InternalMessageInfo

// Model metadata that is specific to text extraction.
type TextExtractionModelMetadata struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TextExtractionModelMetadata) Reset()         { *m = TextExtractionModelMetadata{} }
func (m *TextExtractionModelMetadata) String() string { return proto.CompactTextString(m) }
func (*TextExtractionModelMetadata) ProtoMessage()    {}
func (*TextExtractionModelMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1121cf231f416fd, []int{3}
}

func (m *TextExtractionModelMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextExtractionModelMetadata.Unmarshal(m, b)
}
func (m *TextExtractionModelMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextExtractionModelMetadata.Marshal(b, m, deterministic)
}
func (m *TextExtractionModelMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextExtractionModelMetadata.Merge(m, src)
}
func (m *TextExtractionModelMetadata) XXX_Size() int {
	return xxx_messageInfo_TextExtractionModelMetadata.Size(m)
}
func (m *TextExtractionModelMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_TextExtractionModelMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_TextExtractionModelMetadata proto.InternalMessageInfo

// Dataset metadata for text sentiment.
type TextSentimentDatasetMetadata struct {
	// Required.
	// A sentiment is expressed as an integer ordinal, where higher value
	// means a more positive sentiment. The range of sentiments that will be used
	// is between 0 and sentiment_max (inclusive on both ends), and all the values
	// in the range must be represented in the dataset before a model can be
	// created.
	// sentiment_max value must be between 1 and 10 (inclusive).
	SentimentMax         int32    `protobuf:"varint,1,opt,name=sentiment_max,json=sentimentMax,proto3" json:"sentiment_max,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TextSentimentDatasetMetadata) Reset()         { *m = TextSentimentDatasetMetadata{} }
func (m *TextSentimentDatasetMetadata) String() string { return proto.CompactTextString(m) }
func (*TextSentimentDatasetMetadata) ProtoMessage()    {}
func (*TextSentimentDatasetMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1121cf231f416fd, []int{4}
}

func (m *TextSentimentDatasetMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextSentimentDatasetMetadata.Unmarshal(m, b)
}
func (m *TextSentimentDatasetMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextSentimentDatasetMetadata.Marshal(b, m, deterministic)
}
func (m *TextSentimentDatasetMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextSentimentDatasetMetadata.Merge(m, src)
}
func (m *TextSentimentDatasetMetadata) XXX_Size() int {
	return xxx_messageInfo_TextSentimentDatasetMetadata.Size(m)
}
func (m *TextSentimentDatasetMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_TextSentimentDatasetMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_TextSentimentDatasetMetadata proto.InternalMessageInfo

func (m *TextSentimentDatasetMetadata) GetSentimentMax() int32 {
	if m != nil {
		return m.SentimentMax
	}
	return 0
}

// Model metadata that is specific to text sentiment.
type TextSentimentModelMetadata struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TextSentimentModelMetadata) Reset()         { *m = TextSentimentModelMetadata{} }
func (m *TextSentimentModelMetadata) String() string { return proto.CompactTextString(m) }
func (*TextSentimentModelMetadata) ProtoMessage()    {}
func (*TextSentimentModelMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1121cf231f416fd, []int{5}
}

func (m *TextSentimentModelMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextSentimentModelMetadata.Unmarshal(m, b)
}
func (m *TextSentimentModelMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextSentimentModelMetadata.Marshal(b, m, deterministic)
}
func (m *TextSentimentModelMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextSentimentModelMetadata.Merge(m, src)
}
func (m *TextSentimentModelMetadata) XXX_Size() int {
	return xxx_messageInfo_TextSentimentModelMetadata.Size(m)
}
func (m *TextSentimentModelMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_TextSentimentModelMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_TextSentimentModelMetadata proto.InternalMessageInfo

func init() {
	proto.RegisterType((*TextClassificationDatasetMetadata)(nil), "google.cloud.automl.v1beta1.TextClassificationDatasetMetadata")
	proto.RegisterType((*TextClassificationModelMetadata)(nil), "google.cloud.automl.v1beta1.TextClassificationModelMetadata")
	proto.RegisterType((*TextExtractionDatasetMetadata)(nil), "google.cloud.automl.v1beta1.TextExtractionDatasetMetadata")
	proto.RegisterType((*TextExtractionModelMetadata)(nil), "google.cloud.automl.v1beta1.TextExtractionModelMetadata")
	proto.RegisterType((*TextSentimentDatasetMetadata)(nil), "google.cloud.automl.v1beta1.TextSentimentDatasetMetadata")
	proto.RegisterType((*TextSentimentModelMetadata)(nil), "google.cloud.automl.v1beta1.TextSentimentModelMetadata")
}

func init() {
	proto.RegisterFile("google/cloud/automl/v1beta1/text.proto", fileDescriptor_c1121cf231f416fd)
}

var fileDescriptor_c1121cf231f416fd = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x89, 0xa0, 0xe0, 0xa2, 0x1e, 0xea, 0x45, 0xfa, 0x87, 0xda, 0x08, 0xe2, 0x69, 0xd7,
	0xea, 0xd1, 0x53, 0x5b, 0xc5, 0x53, 0xa0, 0xd4, 0xe2, 0x41, 0x0a, 0x75, 0x9a, 0x8e, 0x21, 0xb0,
	0xd9, 0x09, 0xcd, 0x54, 0xd2, 0x0f, 0xe0, 0xd9, 0xef, 0xe5, 0xa7, 0x92, 0xec, 0xb6, 0xc2, 0xda,
	0xd2, 0x63, 0xf6, 0xfd, 0xde, 0xcb, 0x9b, 0x19, 0x71, 0x9d, 0x10, 0x25, 0x1a, 0x55, 0xac, 0x69,
	0x39, 0x57, 0xb0, 0x64, 0xca, 0xb4, 0xfa, 0xec, 0xce, 0x90, 0xa1, 0xab, 0x18, 0x4b, 0x96, 0xf9,
	0x82, 0x98, 0x6a, 0x0d, 0xc7, 0x49, 0xcb, 0x49, 0xc7, 0xc9, 0x35, 0x57, 0xbf, 0xdd, 0x17, 0x12,
	0x6b, 0x28, 0x8a, 0xf4, 0x23, 0x8d, 0x81, 0x53, 0x32, 0x2e, 0xae, 0xde, 0x5c, 0x3b, 0x20, 0x4f,
	0x15, 0x18, 0x43, 0x6c, 0xc5, 0xc2, 0xa9, 0xe1, 0x57, 0x20, 0x3a, 0x63, 0x2c, 0x79, 0xe0, 0x59,
	0x1f, 0x81, 0xa1, 0x40, 0x8e, 0x90, 0x61, 0x0e, 0x0c, 0xb5, 0x77, 0x71, 0xee, 0x67, 0x4f, 0x79,
	0x95, 0xe3, 0x45, 0x70, 0x19, 0xdc, 0x9c, 0xdd, 0x29, 0xb9, 0xa7, 0xb0, 0xf4, 0x83, 0xc7, 0xab,
	0x1c, 0x47, 0xb5, 0x78, 0xeb, 0x2d, 0xec, 0x88, 0xf6, 0x76, 0x8d, 0x88, 0xe6, 0xa8, 0x37, 0x25,
	0xc2, 0xb6, 0x68, 0x55, 0xc8, 0x53, 0xc9, 0x0b, 0x88, 0x77, 0xb4, 0x0c, 0x5b, 0xa2, 0xe1, 0x03,
	0xbe, 0x7f, 0x20, 0x9a, 0x95, 0xfc, 0x82, 0x86, 0xd3, 0x0c, 0x0d, 0xff, 0x1f, 0xf2, 0x4a, 0x9c,
	0x16, 0x1b, 0x6d, 0x9a, 0x41, 0x69, 0xc7, 0x3b, 0x1c, 0x9d, 0xfc, 0x3d, 0x46, 0x50, 0x86, 0x4d,
	0x51, 0xf7, 0x42, 0xbc, 0x5f, 0xf4, 0xbf, 0x03, 0xd1, 0x8e, 0x29, 0xdb, 0xb7, 0x90, 0xfe, 0x71,
	0xe5, 0x1f, 0x56, 0xcb, 0x1f, 0x06, 0x6f, 0xbd, 0x35, 0x99, 0x90, 0x06, 0x93, 0x48, 0x5a, 0x24,
	0x2a, 0x41, 0x63, 0x4f, 0xa3, 0x9c, 0x04, 0x79, 0x5a, 0xec, 0xbc, 0xf6, 0x83, 0xfb, 0xfc, 0x39,
	0x68, 0x3c, 0x5b, 0x70, 0x32, 0xa8, 0xa0, 0x49, 0x6f, 0xc9, 0x14, 0xe9, 0xc9, 0xab, 0x83, 0x66,
	0x47, 0x36, 0xeb, 0xfe, 0x37, 0x00, 0x00, 0xff, 0xff, 0xf1, 0x97, 0xbc, 0x37, 0x7d, 0x02, 0x00,
	0x00,
}
