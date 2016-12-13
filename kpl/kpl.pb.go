// Code generated by protoc-gen-go.
// source: kpl.proto
// DO NOT EDIT!

/*
Package kpl is a generated protocol buffer package.

It is generated from these files:
	kpl.proto

It has these top-level messages:
	AggregatedRecord
	Tag
	Record
*/
package kpl

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AggregatedRecord struct {
	PartitionKeyTable    []string  `protobuf:"bytes,1,rep,name=partition_key_table,json=partitionKeyTable" json:"partition_key_table,omitempty"`
	ExplicitHashKeyTable []string  `protobuf:"bytes,2,rep,name=explicit_hash_key_table,json=explicitHashKeyTable" json:"explicit_hash_key_table,omitempty"`
	Records              []*Record `protobuf:"bytes,3,rep,name=records" json:"records,omitempty"`
	XXX_unrecognized     []byte    `json:"-"`
}

func (m *AggregatedRecord) Reset()                    { *m = AggregatedRecord{} }
func (m *AggregatedRecord) String() string            { return proto.CompactTextString(m) }
func (*AggregatedRecord) ProtoMessage()               {}
func (*AggregatedRecord) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AggregatedRecord) GetPartitionKeyTable() []string {
	if m != nil {
		return m.PartitionKeyTable
	}
	return nil
}

func (m *AggregatedRecord) GetExplicitHashKeyTable() []string {
	if m != nil {
		return m.ExplicitHashKeyTable
	}
	return nil
}

func (m *AggregatedRecord) GetRecords() []*Record {
	if m != nil {
		return m.Records
	}
	return nil
}

type Tag struct {
	Key              *string `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	Value            *string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Tag) Reset()                    { *m = Tag{} }
func (m *Tag) String() string            { return proto.CompactTextString(m) }
func (*Tag) ProtoMessage()               {}
func (*Tag) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Tag) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *Tag) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

type Record struct {
	PartitionKeyIndex    *uint64 `protobuf:"varint,1,req,name=partition_key_index,json=partitionKeyIndex" json:"partition_key_index,omitempty"`
	ExplicitHashKeyIndex *uint64 `protobuf:"varint,2,opt,name=explicit_hash_key_index,json=explicitHashKeyIndex" json:"explicit_hash_key_index,omitempty"`
	Data                 []byte  `protobuf:"bytes,3,req,name=data" json:"data,omitempty"`
	Tags                 []*Tag  `protobuf:"bytes,4,rep,name=tags" json:"tags,omitempty"`
	XXX_unrecognized     []byte  `json:"-"`
}

func (m *Record) Reset()                    { *m = Record{} }
func (m *Record) String() string            { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()               {}
func (*Record) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Record) GetPartitionKeyIndex() uint64 {
	if m != nil && m.PartitionKeyIndex != nil {
		return *m.PartitionKeyIndex
	}
	return 0
}

func (m *Record) GetExplicitHashKeyIndex() uint64 {
	if m != nil && m.ExplicitHashKeyIndex != nil {
		return *m.ExplicitHashKeyIndex
	}
	return 0
}

func (m *Record) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Record) GetTags() []*Tag {
	if m != nil {
		return m.Tags
	}
	return nil
}

func init() {
	proto.RegisterType((*AggregatedRecord)(nil), "AggregatedRecord")
	proto.RegisterType((*Tag)(nil), "Tag")
	proto.RegisterType((*Record)(nil), "Record")
}

func init() { proto.RegisterFile("kpl.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x8f, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0xc9, 0x26, 0xba, 0x74, 0xf4, 0xb0, 0xc6, 0x05, 0x73, 0xac, 0x3d, 0xe5, 0x62, 0x0f,
	0x82, 0x0f, 0xe0, 0x4d, 0xf1, 0x16, 0x7a, 0x2f, 0xe3, 0x76, 0x48, 0x43, 0xcb, 0xb6, 0xa4, 0x51,
	0xb6, 0xef, 0xa2, 0xef, 0x2a, 0x49, 0xd9, 0x45, 0x51, 0x6f, 0x93, 0xf9, 0xf9, 0xf2, 0xcf, 0x07,
	0x59, 0x37, 0xf6, 0xe5, 0xe8, 0x87, 0x30, 0x14, 0x1f, 0x0c, 0x36, 0x8f, 0xd6, 0x7a, 0xb2, 0x18,
	0xa8, 0x31, 0xb4, 0x1b, 0x7c, 0x23, 0x4b, 0xb8, 0x1e, 0xd1, 0x07, 0x17, 0xdc, 0xb0, 0xaf, 0x3b,
	0x9a, 0xeb, 0x80, 0xaf, 0x3d, 0x29, 0x96, 0x73, 0x9d, 0x99, 0xab, 0x53, 0xf4, 0x42, 0x73, 0x15,
	0x03, 0xf9, 0x00, 0x37, 0x74, 0x18, 0x7b, 0xb7, 0x73, 0xa1, 0x6e, 0x71, 0x6a, 0xbf, 0x31, 0xab,
	0xc4, 0x6c, 0x8f, 0xf1, 0x13, 0x4e, 0xed, 0x09, 0xbb, 0x85, 0xb5, 0x4f, 0x85, 0x93, 0xe2, 0x39,
	0xd7, 0x17, 0xf7, 0xeb, 0x72, 0x39, 0xc0, 0x1c, 0xf7, 0xc5, 0x1d, 0xf0, 0x0a, 0xad, 0xdc, 0x00,
	0xef, 0x68, 0x56, 0x2c, 0x5f, 0xe9, 0xcc, 0xc4, 0x51, 0x6e, 0xe1, 0xec, 0x1d, 0xfb, 0xb7, 0x58,
	0xc0, 0x74, 0x66, 0x96, 0x47, 0xf1, 0xc9, 0xe0, 0xfc, 0x3f, 0x07, 0xb7, 0x6f, 0xe8, 0x90, 0xbe,
	0x10, 0x3f, 0x1d, 0x9e, 0x63, 0xf0, 0xb7, 0xc3, 0xc2, 0xc4, 0x0a, 0xf1, 0xcb, 0x61, 0xc1, 0x24,
	0x88, 0x06, 0x03, 0x2a, 0x9e, 0xaf, 0xf4, 0xa5, 0x49, 0xb3, 0x54, 0x20, 0x02, 0xda, 0x49, 0x89,
	0x24, 0x25, 0xca, 0x0a, 0xad, 0x49, 0x9b, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4d, 0x53, 0xbe,
	0x48, 0x79, 0x01, 0x00, 0x00,
}
