// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: models/v1/item_type.proto

package v1

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ItemType int32

const (
	ItemType_Drink ItemType = 0
	ItemType_Food  ItemType = 1
	ItemType_Both  ItemType = 2
	ItemType_Other ItemType = 3
)

// Enum value maps for ItemType.
var (
	ItemType_name = map[int32]string{
		0: "Drink",
		1: "Food",
		2: "Both",
		3: "Other",
	}
	ItemType_value = map[string]int32{
		"Drink": 0,
		"Food":  1,
		"Both":  2,
		"Other": 3,
	}
)

func (x ItemType) Enum() *ItemType {
	p := new(ItemType)
	*p = x
	return p
}

func (x ItemType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ItemType) Descriptor() protoreflect.EnumDescriptor {
	return file_models_v1_item_type_proto_enumTypes[0].Descriptor()
}

func (ItemType) Type() protoreflect.EnumType {
	return &file_models_v1_item_type_proto_enumTypes[0]
}

func (x ItemType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ItemType.Descriptor instead.
func (ItemType) EnumDescriptor() ([]byte, []int) {
	return file_models_v1_item_type_proto_rawDescGZIP(), []int{0}
}

var File_models_v1_item_type_proto protoreflect.FileDescriptor

var file_models_v1_item_type_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x74, 0x65, 0x6d,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x54, 0x61, 0x70,
	0x4d, 0x61, 0x72, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2a, 0x34,
	0x0a, 0x08, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x72,
	0x69, 0x6e, 0x6b, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x6f, 0x6f, 0x64, 0x10, 0x01, 0x12,
	0x08, 0x0a, 0x04, 0x42, 0x6f, 0x74, 0x68, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x74, 0x68,
	0x65, 0x72, 0x10, 0x03, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x54, 0x61, 0x70, 0x4d, 0x61, 0x72, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_v1_item_type_proto_rawDescOnce sync.Once
	file_models_v1_item_type_proto_rawDescData = file_models_v1_item_type_proto_rawDesc
)

func file_models_v1_item_type_proto_rawDescGZIP() []byte {
	file_models_v1_item_type_proto_rawDescOnce.Do(func() {
		file_models_v1_item_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_v1_item_type_proto_rawDescData)
	})
	return file_models_v1_item_type_proto_rawDescData
}

var file_models_v1_item_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_models_v1_item_type_proto_goTypes = []interface{}{
	(ItemType)(0), // 0: TapMars.models.v1.ItemType
}
var file_models_v1_item_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_models_v1_item_type_proto_init() }
func file_models_v1_item_type_proto_init() {
	if File_models_v1_item_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_models_v1_item_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_v1_item_type_proto_goTypes,
		DependencyIndexes: file_models_v1_item_type_proto_depIdxs,
		EnumInfos:         file_models_v1_item_type_proto_enumTypes,
	}.Build()
	File_models_v1_item_type_proto = out.File
	file_models_v1_item_type_proto_rawDesc = nil
	file_models_v1_item_type_proto_goTypes = nil
	file_models_v1_item_type_proto_depIdxs = nil
}
