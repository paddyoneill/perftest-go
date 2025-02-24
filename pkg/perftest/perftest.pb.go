// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.21.12
// source: perftest.proto

package perftest

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PerftestPortRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PerftestPortRequest) Reset() {
	*x = PerftestPortRequest{}
	mi := &file_perftest_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PerftestPortRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PerftestPortRequest) ProtoMessage() {}

func (x *PerftestPortRequest) ProtoReflect() protoreflect.Message {
	mi := &file_perftest_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PerftestPortRequest.ProtoReflect.Descriptor instead.
func (*PerftestPortRequest) Descriptor() ([]byte, []int) {
	return file_perftest_proto_rawDescGZIP(), []int{0}
}

type PerftestPortResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         int32                  `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PerftestPortResponse) Reset() {
	*x = PerftestPortResponse{}
	mi := &file_perftest_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PerftestPortResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PerftestPortResponse) ProtoMessage() {}

func (x *PerftestPortResponse) ProtoReflect() protoreflect.Message {
	mi := &file_perftest_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PerftestPortResponse.ProtoReflect.Descriptor instead.
func (*PerftestPortResponse) Descriptor() ([]byte, []int) {
	return file_perftest_proto_rawDescGZIP(), []int{1}
}

func (x *PerftestPortResponse) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_perftest_proto protoreflect.FileDescriptor

var file_perftest_proto_rawDesc = string([]byte{
	0x0a, 0x0e, 0x70, 0x65, 0x72, 0x66, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x70, 0x65, 0x72, 0x66, 0x74, 0x65, 0x73, 0x74, 0x22, 0x15, 0x0a, 0x13, 0x50, 0x65,
	0x72, 0x66, 0x74, 0x65, 0x73, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x2c, 0x0a, 0x14, 0x50, 0x65, 0x72, 0x66, 0x74, 0x65, 0x73, 0x74, 0x50, 0x6f, 0x72,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32,
	0x56, 0x0a, 0x08, 0x50, 0x65, 0x72, 0x66, 0x74, 0x65, 0x73, 0x74, 0x12, 0x4a, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x1d, 0x2e, 0x70, 0x65, 0x72, 0x66, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x50, 0x65, 0x72, 0x66, 0x74, 0x65, 0x73, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x65, 0x72, 0x66, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x50, 0x65, 0x72, 0x66, 0x74, 0x65, 0x73, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x61, 0x64, 0x64, 0x79, 0x6f, 0x6e, 0x65, 0x69, 0x6c,
	0x6c, 0x2f, 0x70, 0x65, 0x72, 0x66, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x70, 0x65, 0x72, 0x66, 0x74, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
})

var (
	file_perftest_proto_rawDescOnce sync.Once
	file_perftest_proto_rawDescData []byte
)

func file_perftest_proto_rawDescGZIP() []byte {
	file_perftest_proto_rawDescOnce.Do(func() {
		file_perftest_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_perftest_proto_rawDesc), len(file_perftest_proto_rawDesc)))
	})
	return file_perftest_proto_rawDescData
}

var file_perftest_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_perftest_proto_goTypes = []any{
	(*PerftestPortRequest)(nil),  // 0: perftest.PerftestPortRequest
	(*PerftestPortResponse)(nil), // 1: perftest.PerftestPortResponse
}
var file_perftest_proto_depIdxs = []int32{
	0, // 0: perftest.Perftest.GetPort:input_type -> perftest.PerftestPortRequest
	1, // 1: perftest.Perftest.GetPort:output_type -> perftest.PerftestPortResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_perftest_proto_init() }
func file_perftest_proto_init() {
	if File_perftest_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_perftest_proto_rawDesc), len(file_perftest_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_perftest_proto_goTypes,
		DependencyIndexes: file_perftest_proto_depIdxs,
		MessageInfos:      file_perftest_proto_msgTypes,
	}.Build()
	File_perftest_proto = out.File
	file_perftest_proto_goTypes = nil
	file_perftest_proto_depIdxs = nil
}
