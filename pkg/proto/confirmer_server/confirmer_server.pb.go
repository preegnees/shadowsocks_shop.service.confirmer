// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: pkg/proto/confirmer_server/confirmer_server.proto

package __

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

type Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Label string `protobuf:"bytes,1,opt,name=Label,proto3" json:"Label,omitempty"`
	Id    string `protobuf:"bytes,2,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *Req) Reset() {
	*x = Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_confirmer_server_confirmer_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Req) ProtoMessage() {}

func (x *Req) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_confirmer_server_confirmer_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Req.ProtoReflect.Descriptor instead.
func (*Req) Descriptor() ([]byte, []int) {
	return file_pkg_proto_confirmer_server_confirmer_server_proto_rawDescGZIP(), []int{0}
}

func (x *Req) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *Req) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Resp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=Ok,proto3" json:"Ok,omitempty"`
}

func (x *Resp) Reset() {
	*x = Resp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_confirmer_server_confirmer_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resp) ProtoMessage() {}

func (x *Resp) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_confirmer_server_confirmer_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resp.ProtoReflect.Descriptor instead.
func (*Resp) Descriptor() ([]byte, []int) {
	return file_pkg_proto_confirmer_server_confirmer_server_proto_rawDescGZIP(), []int{1}
}

func (x *Resp) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

var File_pkg_proto_confirmer_server_confirmer_server_proto protoreflect.FileDescriptor

var file_pkg_proto_confirmer_server_confirmer_server_proto_rawDesc = []byte{
	0x0a, 0x31, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x72, 0x6d, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x72, 0x6d, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x11, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x72, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x2b, 0x0a, 0x03, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a,
	0x05, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x49, 0x64, 0x22, 0x16, 0x0a, 0x04, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x4f,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x4f, 0x6b, 0x32, 0xc7, 0x01, 0x0a, 0x10,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x38, 0x0a, 0x03, 0x50, 0x75, 0x74, 0x12, 0x16, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72,
	0x6d, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x1a,
	0x17, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x06, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x72,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x72, 0x6d, 0x12, 0x16, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x72, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x72, 0x6d, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_confirmer_server_confirmer_server_proto_rawDescOnce sync.Once
	file_pkg_proto_confirmer_server_confirmer_server_proto_rawDescData = file_pkg_proto_confirmer_server_confirmer_server_proto_rawDesc
)

func file_pkg_proto_confirmer_server_confirmer_server_proto_rawDescGZIP() []byte {
	file_pkg_proto_confirmer_server_confirmer_server_proto_rawDescOnce.Do(func() {
		file_pkg_proto_confirmer_server_confirmer_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_confirmer_server_confirmer_server_proto_rawDescData)
	})
	return file_pkg_proto_confirmer_server_confirmer_server_proto_rawDescData
}

var file_pkg_proto_confirmer_server_confirmer_server_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_proto_confirmer_server_confirmer_server_proto_goTypes = []interface{}{
	(*Req)(nil),  // 0: confirmer_service.Req
	(*Resp)(nil), // 1: confirmer_service.Resp
}
var file_pkg_proto_confirmer_server_confirmer_server_proto_depIdxs = []int32{
	0, // 0: confirmer_service.ConfirmerService.Put:input_type -> confirmer_service.Req
	0, // 1: confirmer_service.ConfirmerService.Delete:input_type -> confirmer_service.Req
	0, // 2: confirmer_service.ConfirmerService.Confirm:input_type -> confirmer_service.Req
	1, // 3: confirmer_service.ConfirmerService.Put:output_type -> confirmer_service.Resp
	1, // 4: confirmer_service.ConfirmerService.Delete:output_type -> confirmer_service.Resp
	1, // 5: confirmer_service.ConfirmerService.Confirm:output_type -> confirmer_service.Resp
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_proto_confirmer_server_confirmer_server_proto_init() }
func file_pkg_proto_confirmer_server_confirmer_server_proto_init() {
	if File_pkg_proto_confirmer_server_confirmer_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_confirmer_server_confirmer_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Req); i {
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
		file_pkg_proto_confirmer_server_confirmer_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resp); i {
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
			RawDescriptor: file_pkg_proto_confirmer_server_confirmer_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_proto_confirmer_server_confirmer_server_proto_goTypes,
		DependencyIndexes: file_pkg_proto_confirmer_server_confirmer_server_proto_depIdxs,
		MessageInfos:      file_pkg_proto_confirmer_server_confirmer_server_proto_msgTypes,
	}.Build()
	File_pkg_proto_confirmer_server_confirmer_server_proto = out.File
	file_pkg_proto_confirmer_server_confirmer_server_proto_rawDesc = nil
	file_pkg_proto_confirmer_server_confirmer_server_proto_goTypes = nil
	file_pkg_proto_confirmer_server_confirmer_server_proto_depIdxs = nil
}