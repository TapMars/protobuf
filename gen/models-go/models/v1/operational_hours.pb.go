// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: models/v1/operational_hours.proto

package v1

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

type OperationalHours struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start string `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	End   string `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *OperationalHours) Reset() {
	*x = OperationalHours{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_v1_operational_hours_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationalHours) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationalHours) ProtoMessage() {}

func (x *OperationalHours) ProtoReflect() protoreflect.Message {
	mi := &file_models_v1_operational_hours_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationalHours.ProtoReflect.Descriptor instead.
func (*OperationalHours) Descriptor() ([]byte, []int) {
	return file_models_v1_operational_hours_proto_rawDescGZIP(), []int{0}
}

func (x *OperationalHours) GetStart() string {
	if x != nil {
		return x.Start
	}
	return ""
}

func (x *OperationalHours) GetEnd() string {
	if x != nil {
		return x.End
	}
	return ""
}

var File_models_v1_operational_hours_proto protoreflect.FileDescriptor

var file_models_v1_operational_hours_proto_rawDesc = []byte{
	0x0a, 0x21, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x11, 0x54, 0x61, 0x70, 0x4d, 0x61, 0x72, 0x73, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x22, 0x3a, 0x0a, 0x10, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65,
	0x6e, 0x64, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x54, 0x61, 0x70, 0x4d, 0x61, 0x72, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_v1_operational_hours_proto_rawDescOnce sync.Once
	file_models_v1_operational_hours_proto_rawDescData = file_models_v1_operational_hours_proto_rawDesc
)

func file_models_v1_operational_hours_proto_rawDescGZIP() []byte {
	file_models_v1_operational_hours_proto_rawDescOnce.Do(func() {
		file_models_v1_operational_hours_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_v1_operational_hours_proto_rawDescData)
	})
	return file_models_v1_operational_hours_proto_rawDescData
}

var file_models_v1_operational_hours_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_v1_operational_hours_proto_goTypes = []interface{}{
	(*OperationalHours)(nil), // 0: TapMars.models.v1.OperationalHours
}
var file_models_v1_operational_hours_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_models_v1_operational_hours_proto_init() }
func file_models_v1_operational_hours_proto_init() {
	if File_models_v1_operational_hours_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_v1_operational_hours_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationalHours); i {
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
			RawDescriptor: file_models_v1_operational_hours_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_v1_operational_hours_proto_goTypes,
		DependencyIndexes: file_models_v1_operational_hours_proto_depIdxs,
		MessageInfos:      file_models_v1_operational_hours_proto_msgTypes,
	}.Build()
	File_models_v1_operational_hours_proto = out.File
	file_models_v1_operational_hours_proto_rawDesc = nil
	file_models_v1_operational_hours_proto_goTypes = nil
	file_models_v1_operational_hours_proto_depIdxs = nil
}
