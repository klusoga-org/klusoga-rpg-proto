// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: proto/config/configService.proto

package config

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

type RpgConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntityServiceHost string `protobuf:"bytes,1,opt,name=EntityServiceHost,proto3" json:"EntityServiceHost,omitempty"`
	QuestServiceHost  string `protobuf:"bytes,2,opt,name=QuestServiceHost,proto3" json:"QuestServiceHost,omitempty"`
	ShopServiceHost   string `protobuf:"bytes,3,opt,name=ShopServiceHost,proto3" json:"ShopServiceHost,omitempty"`
}

func (x *RpgConfigResponse) Reset() {
	*x = RpgConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_config_configService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpgConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpgConfigResponse) ProtoMessage() {}

func (x *RpgConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_config_configService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RpgConfigResponse.ProtoReflect.Descriptor instead.
func (*RpgConfigResponse) Descriptor() ([]byte, []int) {
	return file_proto_config_configService_proto_rawDescGZIP(), []int{0}
}

func (x *RpgConfigResponse) GetEntityServiceHost() string {
	if x != nil {
		return x.EntityServiceHost
	}
	return ""
}

func (x *RpgConfigResponse) GetQuestServiceHost() string {
	if x != nil {
		return x.QuestServiceHost
	}
	return ""
}

func (x *RpgConfigResponse) GetShopServiceHost() string {
	if x != nil {
		return x.ShopServiceHost
	}
	return ""
}

type RpgConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstanceId string `protobuf:"bytes,1,opt,name=InstanceId,proto3" json:"InstanceId,omitempty"`
}

func (x *RpgConfigRequest) Reset() {
	*x = RpgConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_config_configService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpgConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpgConfigRequest) ProtoMessage() {}

func (x *RpgConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_config_configService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RpgConfigRequest.ProtoReflect.Descriptor instead.
func (*RpgConfigRequest) Descriptor() ([]byte, []int) {
	return file_proto_config_configService_proto_rawDescGZIP(), []int{1}
}

func (x *RpgConfigRequest) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

var File_proto_config_configService_proto protoreflect.FileDescriptor

var file_proto_config_configService_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x22, 0x97, 0x01, 0x0a, 0x11, 0x52, 0x70, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x11, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x51, 0x75, 0x65, 0x73, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x51, 0x75, 0x65, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x48, 0x6f, 0x73,
	0x74, 0x12, 0x28, 0x0a, 0x0f, 0x53, 0x68, 0x6f, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x48, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x53, 0x68, 0x6f, 0x70,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x22, 0x32, 0x0a, 0x10, 0x52,
	0x70, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1e, 0x0a, 0x0a, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x32,
	0x62, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x51, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x52, 0x70, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x1f, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x52, 0x70, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x20, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x52, 0x70, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6b, 0x6c, 0x75, 0x73, 0x6f, 0x67, 0x61, 0x2d, 0x6f, 0x72, 0x67, 0x2f, 0x6b, 0x6c,
	0x75, 0x73, 0x6f, 0x67, 0x61, 0x2d, 0x72, 0x70, 0x67, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_config_configService_proto_rawDescOnce sync.Once
	file_proto_config_configService_proto_rawDescData = file_proto_config_configService_proto_rawDesc
)

func file_proto_config_configService_proto_rawDescGZIP() []byte {
	file_proto_config_configService_proto_rawDescOnce.Do(func() {
		file_proto_config_configService_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_config_configService_proto_rawDescData)
	})
	return file_proto_config_configService_proto_rawDescData
}

var file_proto_config_configService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_config_configService_proto_goTypes = []interface{}{
	(*RpgConfigResponse)(nil), // 0: configService.RpgConfigResponse
	(*RpgConfigRequest)(nil),  // 1: configService.RpgConfigRequest
}
var file_proto_config_configService_proto_depIdxs = []int32{
	1, // 0: configService.ConfigService.GetRpgConfig:input_type -> configService.RpgConfigRequest
	0, // 1: configService.ConfigService.GetRpgConfig:output_type -> configService.RpgConfigResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_config_configService_proto_init() }
func file_proto_config_configService_proto_init() {
	if File_proto_config_configService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_config_configService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RpgConfigResponse); i {
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
		file_proto_config_configService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RpgConfigRequest); i {
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
			RawDescriptor: file_proto_config_configService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_config_configService_proto_goTypes,
		DependencyIndexes: file_proto_config_configService_proto_depIdxs,
		MessageInfos:      file_proto_config_configService_proto_msgTypes,
	}.Build()
	File_proto_config_configService_proto = out.File
	file_proto_config_configService_proto_rawDesc = nil
	file_proto_config_configService_proto_goTypes = nil
	file_proto_config_configService_proto_depIdxs = nil
}
