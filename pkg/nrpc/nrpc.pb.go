// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.5
// source: proto/nrpc.proto

package nrpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RequestWrapper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *anypb.Any `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *RequestWrapper) Reset() {
	*x = RequestWrapper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_nrpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestWrapper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestWrapper) ProtoMessage() {}

func (x *RequestWrapper) ProtoReflect() protoreflect.Message {
	mi := &file_proto_nrpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestWrapper.ProtoReflect.Descriptor instead.
func (*RequestWrapper) Descriptor() ([]byte, []int) {
	return file_proto_nrpc_proto_rawDescGZIP(), []int{0}
}

func (x *RequestWrapper) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

type ResponseWrapper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//	*ResponseWrapper_Ok
	//	*ResponseWrapper_Err
	Data isResponseWrapper_Data `protobuf_oneof:"data"`
}

func (x *ResponseWrapper) Reset() {
	*x = ResponseWrapper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_nrpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseWrapper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseWrapper) ProtoMessage() {}

func (x *ResponseWrapper) ProtoReflect() protoreflect.Message {
	mi := &file_proto_nrpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseWrapper.ProtoReflect.Descriptor instead.
func (*ResponseWrapper) Descriptor() ([]byte, []int) {
	return file_proto_nrpc_proto_rawDescGZIP(), []int{1}
}

func (m *ResponseWrapper) GetData() isResponseWrapper_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *ResponseWrapper) GetOk() *anypb.Any {
	if x, ok := x.GetData().(*ResponseWrapper_Ok); ok {
		return x.Ok
	}
	return nil
}

func (x *ResponseWrapper) GetErr() string {
	if x, ok := x.GetData().(*ResponseWrapper_Err); ok {
		return x.Err
	}
	return ""
}

type isResponseWrapper_Data interface {
	isResponseWrapper_Data()
}

type ResponseWrapper_Ok struct {
	Ok *anypb.Any `protobuf:"bytes,1,opt,name=ok,proto3,oneof"`
}

type ResponseWrapper_Err struct {
	Err string `protobuf:"bytes,3,opt,name=err,proto3,oneof"`
}

func (*ResponseWrapper_Ok) isResponseWrapper_Data() {}

func (*ResponseWrapper_Err) isResponseWrapper_Data() {}

var File_proto_nrpc_proto protoreflect.FileDescriptor

var file_proto_nrpc_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3a, 0x0a,
	0x0e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x12,
	0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x41, 0x6e, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x55, 0x0a, 0x0f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x02,
	0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x48, 0x00,
	0x52, 0x02, 0x6f, 0x6b, 0x12, 0x12, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x03, 0x65, 0x72, 0x72, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65,
	0x6d, 0x6d, 0x30, 0x33, 0x35, 0x2f, 0x6e, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6e,
	0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_nrpc_proto_rawDescOnce sync.Once
	file_proto_nrpc_proto_rawDescData = file_proto_nrpc_proto_rawDesc
)

func file_proto_nrpc_proto_rawDescGZIP() []byte {
	file_proto_nrpc_proto_rawDescOnce.Do(func() {
		file_proto_nrpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_nrpc_proto_rawDescData)
	})
	return file_proto_nrpc_proto_rawDescData
}

var file_proto_nrpc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_nrpc_proto_goTypes = []interface{}{
	(*RequestWrapper)(nil),  // 0: RequestWrapper
	(*ResponseWrapper)(nil), // 1: ResponseWrapper
	(*anypb.Any)(nil),       // 2: google.protobuf.Any
}
var file_proto_nrpc_proto_depIdxs = []int32{
	2, // 0: RequestWrapper.data:type_name -> google.protobuf.Any
	2, // 1: ResponseWrapper.ok:type_name -> google.protobuf.Any
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_nrpc_proto_init() }
func file_proto_nrpc_proto_init() {
	if File_proto_nrpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_nrpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestWrapper); i {
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
		file_proto_nrpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseWrapper); i {
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
	file_proto_nrpc_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*ResponseWrapper_Ok)(nil),
		(*ResponseWrapper_Err)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_nrpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_nrpc_proto_goTypes,
		DependencyIndexes: file_proto_nrpc_proto_depIdxs,
		MessageInfos:      file_proto_nrpc_proto_msgTypes,
	}.Build()
	File_proto_nrpc_proto = out.File
	file_proto_nrpc_proto_rawDesc = nil
	file_proto_nrpc_proto_goTypes = nil
	file_proto_nrpc_proto_depIdxs = nil
}
