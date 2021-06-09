package foo

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message    string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	ReturnCode int32  `protobuf:"varint,2,opt,name=return_code,json=returnCode,proto3" json:"return_code,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foo_foo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_foo_foo_proto_msgTypes[0]
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
	return file_foo_foo_proto_rawDescGZIP(), []int{0}
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Response) GetReturnCode() int32 {
	if x != nil {
		return x.ReturnCode
	}
	return 0
}

var File_foo_foo_proto protoreflect.FileDescriptor

var file_foo_foo_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x66, 0x6f, 0x6f, 0x2f, 0x66, 0x6f, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x66, 0x6f, 0x6f, 0x22, 0x45, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x16, 0x5a, 0x14, 0x67,
	0x72, 0x69, 0x70, 0x6d, 0x6f, 0x63, 0x6b, 0x2f, 0x70, 0x72, 0x6a, 0x2d, 0x66, 0x6f, 0x6f, 0x2f,
	0x66, 0x6f, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_foo_foo_proto_rawDescOnce sync.Once
	file_foo_foo_proto_rawDescData = file_foo_foo_proto_rawDesc
)

func file_foo_foo_proto_rawDescGZIP() []byte {
	file_foo_foo_proto_rawDescOnce.Do(func() {
		file_foo_foo_proto_rawDescData = protoimpl.X.CompressGZIP(file_foo_foo_proto_rawDescData)
	})
	return file_foo_foo_proto_rawDescData
}

var file_foo_foo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_foo_foo_proto_goTypes = []interface{}{
	(*Response)(nil), // 0: foo.Response
}
var file_foo_foo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_foo_foo_proto_init() }
func file_foo_foo_proto_init() {
	if File_foo_foo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_foo_foo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_foo_foo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_foo_foo_proto_goTypes,
		DependencyIndexes: file_foo_foo_proto_depIdxs,
		MessageInfos:      file_foo_foo_proto_msgTypes,
	}.Build()
	File_foo_foo_proto = out.File
	file_foo_foo_proto_rawDesc = nil
	file_foo_foo_proto_goTypes = nil
	file_foo_foo_proto_depIdxs = nil
}
