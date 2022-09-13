// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gloo/api/v1/plugins/cloudmap/cloudmap.proto

package cloudmap

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/solo-io/protoc-gen-ext/extproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Upstream Spec for AWS CloudMap Upstreams
type UpstreamSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceName           string `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	AwsAccountId          string `protobuf:"bytes,2,opt,name=aws_account_id,json=awsAccountId,proto3" json:"aws_account_id,omitempty"`
	CloudmapNamespaceName string `protobuf:"bytes,3,opt,name=cloudmap_namespace_name,json=cloudmapNamespaceName,proto3" json:"cloudmap_namespace_name,omitempty"`
	Port                  uint32 `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
	AssumeRoleName        string `protobuf:"bytes,5,opt,name=assume_role_name,json=assumeRoleName,proto3" json:"assume_role_name,omitempty"`
}

func (x *UpstreamSpec) Reset() {
	*x = UpstreamSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_plugins_cloudmap_cloudmap_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpstreamSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpstreamSpec) ProtoMessage() {}

func (x *UpstreamSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_plugins_cloudmap_cloudmap_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpstreamSpec.ProtoReflect.Descriptor instead.
func (*UpstreamSpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_plugins_cloudmap_cloudmap_proto_rawDescGZIP(), []int{0}
}

func (x *UpstreamSpec) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *UpstreamSpec) GetAwsAccountId() string {
	if x != nil {
		return x.AwsAccountId
	}
	return ""
}

func (x *UpstreamSpec) GetCloudmapNamespaceName() string {
	if x != nil {
		return x.CloudmapNamespaceName
	}
	return ""
}

func (x *UpstreamSpec) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *UpstreamSpec) GetAssumeRoleName() string {
	if x != nil {
		return x.AssumeRoleName
	}
	return ""
}

var File_github_com_solo_io_gloo_projects_gloo_api_v1_plugins_cloudmap_cloudmap_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_v1_plugins_cloudmap_cloudmap_proto_rawDesc = []byte{
	0x0a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x61, 0x70, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x61, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x61, 0x70, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x12, 0x65,
	0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xcd, 0x01, 0x0a, 0x0c, 0x55, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x70,
	0x65, 0x63, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x61, 0x77, 0x73, 0x5f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61,
	0x77, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x17, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x6d, 0x61, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x6d, 0x61, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x61, 0x73, 0x73, 0x75, 0x6d,
	0x65, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x61, 0x73, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x42, 0x47, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x6d, 0x61, 0x70, 0xc0, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_v1_plugins_cloudmap_cloudmap_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_v1_plugins_cloudmap_cloudmap_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_v1_plugins_cloudmap_cloudmap_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_v1_plugins_cloudmap_cloudmap_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_plugins_cloudmap_cloudmap_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_plugins_cloudmap_cloudmap_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_v1_plugins_cloudmap_cloudmap_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_plugins_cloudmap_cloudmap_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_v1_plugins_cloudmap_cloudmap_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_github_com_solo_io_gloo_projects_gloo_api_v1_plugins_cloudmap_cloudmap_proto_goTypes = []interface{}{
	(*UpstreamSpec)(nil), // 0: cloudmap.options.gloo.solo.io.UpstreamSpec
}
var file_github_com_solo_io_gloo_projects_gloo_api_v1_plugins_cloudmap_cloudmap_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_gloo_projects_gloo_api_v1_plugins_cloudmap_cloudmap_proto_init() }
func file_github_com_solo_io_gloo_projects_gloo_api_v1_plugins_cloudmap_cloudmap_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_v1_plugins_cloudmap_cloudmap_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_plugins_cloudmap_cloudmap_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpstreamSpec); i {
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
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_v1_plugins_cloudmap_cloudmap_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_v1_plugins_cloudmap_cloudmap_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_v1_plugins_cloudmap_cloudmap_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_v1_plugins_cloudmap_cloudmap_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_v1_plugins_cloudmap_cloudmap_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_v1_plugins_cloudmap_cloudmap_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_plugins_cloudmap_cloudmap_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_plugins_cloudmap_cloudmap_proto_depIdxs = nil
}