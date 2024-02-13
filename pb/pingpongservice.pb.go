// Copyright (c) 2024 g41797
// SPDX-License-Identifier: MIT

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: pingpongservice.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Part struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Content []byte  `protobuf:"bytes,2,opt,name=content,proto3,oneof" json:"content,omitempty"`
	Text    *string `protobuf:"bytes,3,opt,name=text,proto3,oneof" json:"text,omitempty"`
}

func (x *Part) Reset() {
	*x = Part{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pingpongservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Part) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Part) ProtoMessage() {}

func (x *Part) ProtoReflect() protoreflect.Message {
	mi := &file_pingpongservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Part.ProtoReflect.Descriptor instead.
func (*Part) Descriptor() ([]byte, []int) {
	return file_pingpongservice_proto_rawDescGZIP(), []int{0}
}

func (x *Part) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Part) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *Part) GetText() string {
	if x != nil && x.Text != nil {
		return *x.Text
	}
	return ""
}

type Ball struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parts []*Part `protobuf:"bytes,1,rep,name=parts,proto3" json:"parts,omitempty"`
}

func (x *Ball) Reset() {
	*x = Ball{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pingpongservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ball) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ball) ProtoMessage() {}

func (x *Ball) ProtoReflect() protoreflect.Message {
	mi := &file_pingpongservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ball.ProtoReflect.Descriptor instead.
func (*Ball) Descriptor() ([]byte, []int) {
	return file_pingpongservice_proto_rawDescGZIP(), []int{1}
}

func (x *Ball) GetParts() []*Part {
	if x != nil {
		return x.Parts
	}
	return nil
}

var File_pingpongservice_proto protoreflect.FileDescriptor

var file_pingpongservice_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6e, 0x67, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x67, 0x0a, 0x04, 0x50,
	0x61, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x88, 0x01, 0x01, 0x42,
	0x0a, 0x0a, 0x08, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x07, 0x0a, 0x05, 0x5f,
	0x74, 0x65, 0x78, 0x74, 0x22, 0x26, 0x0a, 0x04, 0x42, 0x61, 0x6c, 0x6c, 0x12, 0x1e, 0x0a, 0x05,
	0x70, 0x61, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62,
	0x2e, 0x50, 0x61, 0x72, 0x74, 0x52, 0x05, 0x70, 0x61, 0x72, 0x74, 0x73, 0x32, 0x2f, 0x0a, 0x0f,
	0x50, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x1c, 0x0a, 0x04, 0x50, 0x6c, 0x61, 0x79, 0x12, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x61, 0x6c,
	0x6c, 0x1a, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x61, 0x6c, 0x6c, 0x22, 0x00, 0x42, 0x24, 0x5a,
	0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x34, 0x31, 0x37,
	0x39, 0x37, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x70, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6e, 0x67,
	0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pingpongservice_proto_rawDescOnce sync.Once
	file_pingpongservice_proto_rawDescData = file_pingpongservice_proto_rawDesc
)

func file_pingpongservice_proto_rawDescGZIP() []byte {
	file_pingpongservice_proto_rawDescOnce.Do(func() {
		file_pingpongservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_pingpongservice_proto_rawDescData)
	})
	return file_pingpongservice_proto_rawDescData
}

var file_pingpongservice_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pingpongservice_proto_goTypes = []interface{}{
	(*Part)(nil), // 0: pb.Part
	(*Ball)(nil), // 1: pb.Ball
}
var file_pingpongservice_proto_depIdxs = []int32{
	0, // 0: pb.Ball.parts:type_name -> pb.Part
	1, // 1: pb.PingPongService.Play:input_type -> pb.Ball
	1, // 2: pb.PingPongService.Play:output_type -> pb.Ball
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pingpongservice_proto_init() }
func file_pingpongservice_proto_init() {
	if File_pingpongservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pingpongservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Part); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pingpongservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ball); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_pingpongservice_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pingpongservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pingpongservice_proto_goTypes,
		DependencyIndexes: file_pingpongservice_proto_depIdxs,
		MessageInfos:      file_pingpongservice_proto_msgTypes,
	}.Build()
	File_pingpongservice_proto = out.File
	file_pingpongservice_proto_rawDesc = nil
	file_pingpongservice_proto_goTypes = nil
	file_pingpongservice_proto_depIdxs = nil
}
