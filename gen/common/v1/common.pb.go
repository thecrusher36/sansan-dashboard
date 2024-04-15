// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: common/v1/common.proto

package commonv1

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

type StandardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status       string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Code         uint64 `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	ErrorMessage string `protobuf:"bytes,3,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
}

func (x *StandardResponse) Reset() {
	*x = StandardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_v1_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StandardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StandardResponse) ProtoMessage() {}

func (x *StandardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_v1_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StandardResponse.ProtoReflect.Descriptor instead.
func (*StandardResponse) Descriptor() ([]byte, []int) {
	return file_common_v1_common_proto_rawDescGZIP(), []int{0}
}

func (x *StandardResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *StandardResponse) GetCode() uint64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *StandardResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

type ErrorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message    string            `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	HttpStatus *StandardResponse `protobuf:"bytes,2,opt,name=http_status,json=httpStatus,proto3" json:"http_status,omitempty"`
}

func (x *ErrorResponse) Reset() {
	*x = ErrorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_v1_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorResponse) ProtoMessage() {}

func (x *ErrorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_v1_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorResponse.ProtoReflect.Descriptor instead.
func (*ErrorResponse) Descriptor() ([]byte, []int) {
	return file_common_v1_common_proto_rawDescGZIP(), []int{1}
}

func (x *ErrorResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ErrorResponse) GetHttpStatus() *StandardResponse {
	if x != nil {
		return x.HttpStatus
	}
	return nil
}

var File_common_v1_common_proto protoreflect.FileDescriptor

var file_common_v1_common_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x22, 0x63, 0x0a, 0x10, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x67, 0x0a, 0x0d, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0a, 0x68, 0x74, 0x74, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x42, 0xa4, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x42, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x61, 0x6e, 0x64, 0x69, 0x73, 0x75, 0x72, 0x79, 0x61, 0x64, 0x69, 0x33, 0x36, 0x2f, 0x73, 0x61,
	0x6e, 0x73, 0x61, 0x6e, 0x2d, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_v1_common_proto_rawDescOnce sync.Once
	file_common_v1_common_proto_rawDescData = file_common_v1_common_proto_rawDesc
)

func file_common_v1_common_proto_rawDescGZIP() []byte {
	file_common_v1_common_proto_rawDescOnce.Do(func() {
		file_common_v1_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_v1_common_proto_rawDescData)
	})
	return file_common_v1_common_proto_rawDescData
}

var file_common_v1_common_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_common_v1_common_proto_goTypes = []interface{}{
	(*StandardResponse)(nil), // 0: common.v1.StandardResponse
	(*ErrorResponse)(nil),    // 1: common.v1.ErrorResponse
}
var file_common_v1_common_proto_depIdxs = []int32{
	0, // 0: common.v1.ErrorResponse.http_status:type_name -> common.v1.StandardResponse
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_common_v1_common_proto_init() }
func file_common_v1_common_proto_init() {
	if File_common_v1_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_v1_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StandardResponse); i {
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
		file_common_v1_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_v1_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_v1_common_proto_goTypes,
		DependencyIndexes: file_common_v1_common_proto_depIdxs,
		MessageInfos:      file_common_v1_common_proto_msgTypes,
	}.Build()
	File_common_v1_common_proto = out.File
	file_common_v1_common_proto_rawDesc = nil
	file_common_v1_common_proto_goTypes = nil
	file_common_v1_common_proto_depIdxs = nil
}
