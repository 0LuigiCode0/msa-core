// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: proto/msa_service.proto

package msa_service

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

type RequestCall struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data     []byte            `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
	FuncName string            `protobuf:"bytes,2,opt,name=FuncName,proto3" json:"FuncName,omitempty"`
	Args     map[string]string `protobuf:"bytes,3,rep,name=Args,proto3" json:"Args,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *RequestCall) Reset() {
	*x = RequestCall{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_msa_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestCall) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestCall) ProtoMessage() {}

func (x *RequestCall) ProtoReflect() protoreflect.Message {
	mi := &file_proto_msa_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestCall.ProtoReflect.Descriptor instead.
func (*RequestCall) Descriptor() ([]byte, []int) {
	return file_proto_msa_service_proto_rawDescGZIP(), []int{0}
}

func (x *RequestCall) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *RequestCall) GetFuncName() string {
	if x != nil {
		return x.FuncName
	}
	return ""
}

func (x *RequestCall) GetArgs() map[string]string {
	if x != nil {
		return x.Args
	}
	return nil
}

type RequestAddService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host      string `protobuf:"bytes,1,opt,name=Host,proto3" json:"Host,omitempty"`
	Port      int32  `protobuf:"varint,2,opt,name=Port,proto3" json:"Port,omitempty"`
	GroupType string `protobuf:"bytes,3,opt,name=GroupType,proto3" json:"GroupType,omitempty"`
	Key       string `protobuf:"bytes,4,opt,name=Key,proto3" json:"Key,omitempty"`
}

func (x *RequestAddService) Reset() {
	*x = RequestAddService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_msa_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestAddService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestAddService) ProtoMessage() {}

func (x *RequestAddService) ProtoReflect() protoreflect.Message {
	mi := &file_proto_msa_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestAddService.ProtoReflect.Descriptor instead.
func (*RequestAddService) Descriptor() ([]byte, []int) {
	return file_proto_msa_service_proto_rawDescGZIP(), []int{1}
}

func (x *RequestAddService) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *RequestAddService) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *RequestAddService) GetGroupType() string {
	if x != nil {
		return x.GroupType
	}
	return ""
}

func (x *RequestAddService) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type RequestDelService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupType string `protobuf:"bytes,1,opt,name=GroupType,proto3" json:"GroupType,omitempty"`
	Key       string `protobuf:"bytes,2,opt,name=Key,proto3" json:"Key,omitempty"`
}

func (x *RequestDelService) Reset() {
	*x = RequestDelService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_msa_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestDelService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestDelService) ProtoMessage() {}

func (x *RequestDelService) ProtoReflect() protoreflect.Message {
	mi := &file_proto_msa_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestDelService.ProtoReflect.Descriptor instead.
func (*RequestDelService) Descriptor() ([]byte, []int) {
	return file_proto_msa_service_proto_rawDescGZIP(), []int{2}
}

func (x *RequestDelService) GetGroupType() string {
	if x != nil {
		return x.GroupType
	}
	return ""
}

func (x *RequestDelService) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type ResponseCall struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []byte `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
}

func (x *ResponseCall) Reset() {
	*x = ResponseCall{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_msa_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseCall) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseCall) ProtoMessage() {}

func (x *ResponseCall) ProtoReflect() protoreflect.Message {
	mi := &file_proto_msa_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseCall.ProtoReflect.Descriptor instead.
func (*ResponseCall) Descriptor() ([]byte, []int) {
	return file_proto_msa_service_proto_rawDescGZIP(), []int{3}
}

func (x *ResponseCall) GetResult() []byte {
	if x != nil {
		return x.Result
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_msa_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_msa_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proto_msa_service_proto_rawDescGZIP(), []int{4}
}

func (x *Response) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_proto_msa_service_proto protoreflect.FileDescriptor

var file_proto_msa_service_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x73, 0x61, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6d, 0x73, 0x61, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0xae, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x75,
	0x6e, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x75,
	0x6e, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x04, 0x41, 0x72, 0x67, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6d, 0x73, 0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x2e, 0x41,
	0x72, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x37,
	0x0a, 0x09, 0x41, 0x72, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x6b, 0x0a, 0x11, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x41, 0x64, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x48, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x48, 0x6f, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x50, 0x6f, 0x72, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x4b, 0x65, 0x79, 0x22, 0x43, 0x0a, 0x11, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44,
	0x65, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x22, 0x26, 0x0a, 0x0c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x22, 0x24, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xd3, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x04, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x18, 0x2e, 0x6d, 0x73,
	0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x43, 0x61, 0x6c, 0x6c, 0x1a, 0x19, 0x2e, 0x6d, 0x73, 0x61, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x61, 0x6c, 0x6c,
	0x12, 0x43, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1e,
	0x2e, 0x6d, 0x73, 0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x41, 0x64, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x15,
	0x2e, 0x6d, 0x73, 0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1e, 0x2e, 0x6d, 0x73, 0x61, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x65, 0x6c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x15, 0x2e, 0x6d, 0x73, 0x61, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x12, 0x5a,
	0x10, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x73, 0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_msa_service_proto_rawDescOnce sync.Once
	file_proto_msa_service_proto_rawDescData = file_proto_msa_service_proto_rawDesc
)

func file_proto_msa_service_proto_rawDescGZIP() []byte {
	file_proto_msa_service_proto_rawDescOnce.Do(func() {
		file_proto_msa_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_msa_service_proto_rawDescData)
	})
	return file_proto_msa_service_proto_rawDescData
}

var file_proto_msa_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_msa_service_proto_goTypes = []interface{}{
	(*RequestCall)(nil),       // 0: msa_service.RequestCall
	(*RequestAddService)(nil), // 1: msa_service.RequestAddService
	(*RequestDelService)(nil), // 2: msa_service.RequestDelService
	(*ResponseCall)(nil),      // 3: msa_service.ResponseCall
	(*Response)(nil),          // 4: msa_service.Response
	nil,                       // 5: msa_service.RequestCall.ArgsEntry
}
var file_proto_msa_service_proto_depIdxs = []int32{
	5, // 0: msa_service.RequestCall.Args:type_name -> msa_service.RequestCall.ArgsEntry
	0, // 1: msa_service.Service.Call:input_type -> msa_service.RequestCall
	1, // 2: msa_service.Service.AddService:input_type -> msa_service.RequestAddService
	2, // 3: msa_service.Service.DeleteService:input_type -> msa_service.RequestDelService
	3, // 4: msa_service.Service.Call:output_type -> msa_service.ResponseCall
	4, // 5: msa_service.Service.AddService:output_type -> msa_service.Response
	4, // 6: msa_service.Service.DeleteService:output_type -> msa_service.Response
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_msa_service_proto_init() }
func file_proto_msa_service_proto_init() {
	if File_proto_msa_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_msa_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestCall); i {
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
		file_proto_msa_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestAddService); i {
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
		file_proto_msa_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestDelService); i {
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
		file_proto_msa_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseCall); i {
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
		file_proto_msa_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_proto_msa_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_msa_service_proto_goTypes,
		DependencyIndexes: file_proto_msa_service_proto_depIdxs,
		MessageInfos:      file_proto_msa_service_proto_msgTypes,
	}.Build()
	File_proto_msa_service_proto = out.File
	file_proto_msa_service_proto_rawDesc = nil
	file_proto_msa_service_proto_goTypes = nil
	file_proto_msa_service_proto_depIdxs = nil
}
