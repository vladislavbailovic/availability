// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: probe.proto

package model

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Probe struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProbeID      int32                  `protobuf:"varint,1,opt,name=ProbeID,proto3" json:"ProbeID,omitempty"`
	SiteID       int32                  `protobuf:"varint,2,opt,name=SiteID,proto3" json:"SiteID,omitempty"`
	Recorded     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=Recorded,proto3" json:"Recorded,omitempty"`
	ResponseTime *durationpb.Duration   `protobuf:"bytes,4,opt,name=ResponseTime,proto3" json:"ResponseTime,omitempty"`
	Err          HttpErr                `protobuf:"varint,5,opt,name=Err,proto3,enum=definitions.HttpErr" json:"Err,omitempty"`
	Msg          string                 `protobuf:"bytes,6,opt,name=Msg,proto3" json:"Msg,omitempty"`
}

func (x *Probe) Reset() {
	*x = Probe{}
	if protoimpl.UnsafeEnabled {
		mi := &file_probe_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Probe) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Probe) ProtoMessage() {}

func (x *Probe) ProtoReflect() protoreflect.Message {
	mi := &file_probe_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Probe.ProtoReflect.Descriptor instead.
func (*Probe) Descriptor() ([]byte, []int) {
	return file_probe_proto_rawDescGZIP(), []int{0}
}

func (x *Probe) GetProbeID() int32 {
	if x != nil {
		return x.ProbeID
	}
	return 0
}

func (x *Probe) GetSiteID() int32 {
	if x != nil {
		return x.SiteID
	}
	return 0
}

func (x *Probe) GetRecorded() *timestamppb.Timestamp {
	if x != nil {
		return x.Recorded
	}
	return nil
}

func (x *Probe) GetResponseTime() *durationpb.Duration {
	if x != nil {
		return x.ResponseTime
	}
	return nil
}

func (x *Probe) GetErr() HttpErr {
	if x != nil {
		return x.Err
	}
	return HttpErr_HTTPERR_NONE
}

func (x *Probe) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type ProbeRef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProbeID int32   `protobuf:"varint,1,opt,name=ProbeID,proto3" json:"ProbeID,omitempty"`
	Err     HttpErr `protobuf:"varint,2,opt,name=Err,proto3,enum=definitions.HttpErr" json:"Err,omitempty"`
}

func (x *ProbeRef) Reset() {
	*x = ProbeRef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_probe_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProbeRef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProbeRef) ProtoMessage() {}

func (x *ProbeRef) ProtoReflect() protoreflect.Message {
	mi := &file_probe_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProbeRef.ProtoReflect.Descriptor instead.
func (*ProbeRef) Descriptor() ([]byte, []int) {
	return file_probe_proto_rawDescGZIP(), []int{1}
}

func (x *ProbeRef) GetProbeID() int32 {
	if x != nil {
		return x.ProbeID
	}
	return 0
}

func (x *ProbeRef) GetErr() HttpErr {
	if x != nil {
		return x.Err
	}
	return HttpErr_HTTPERR_NONE
}

var File_probe_proto protoreflect.FileDescriptor

var file_probe_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x64,
	0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x68, 0x74, 0x74,
	0x70, 0x65, 0x72, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xea, 0x01, 0x0a, 0x05, 0x50,
	0x72, 0x6f, 0x62, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x49, 0x44, 0x12, 0x16,
	0x0a, 0x06, 0x53, 0x69, 0x74, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x53, 0x69, 0x74, 0x65, 0x49, 0x44, 0x12, 0x36, 0x0a, 0x08, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x64, 0x12, 0x3d,
	0x0a, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x26, 0x0a,
	0x03, 0x45, 0x72, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x45, 0x72, 0x72,
	0x52, 0x03, 0x45, 0x72, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x22, 0x4c, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x62, 0x65,
	0x52, 0x65, 0x66, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x49, 0x44, 0x12, 0x26, 0x0a,
	0x03, 0x45, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x45, 0x72, 0x72,
	0x52, 0x03, 0x45, 0x72, 0x72, 0x42, 0x10, 0x5a, 0x0e, 0x70, 0x6b, 0x67, 0x2f, 0x64, 0x61, 0x74,
	0x61, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_probe_proto_rawDescOnce sync.Once
	file_probe_proto_rawDescData = file_probe_proto_rawDesc
)

func file_probe_proto_rawDescGZIP() []byte {
	file_probe_proto_rawDescOnce.Do(func() {
		file_probe_proto_rawDescData = protoimpl.X.CompressGZIP(file_probe_proto_rawDescData)
	})
	return file_probe_proto_rawDescData
}

var file_probe_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_probe_proto_goTypes = []interface{}{
	(*Probe)(nil),                 // 0: definitions.Probe
	(*ProbeRef)(nil),              // 1: definitions.ProbeRef
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),   // 3: google.protobuf.Duration
	(HttpErr)(0),                  // 4: definitions.HttpErr
}
var file_probe_proto_depIdxs = []int32{
	2, // 0: definitions.Probe.Recorded:type_name -> google.protobuf.Timestamp
	3, // 1: definitions.Probe.ResponseTime:type_name -> google.protobuf.Duration
	4, // 2: definitions.Probe.Err:type_name -> definitions.HttpErr
	4, // 3: definitions.ProbeRef.Err:type_name -> definitions.HttpErr
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_probe_proto_init() }
func file_probe_proto_init() {
	if File_probe_proto != nil {
		return
	}
	file_httperr_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_probe_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Probe); i {
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
		file_probe_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProbeRef); i {
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
			RawDescriptor: file_probe_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_probe_proto_goTypes,
		DependencyIndexes: file_probe_proto_depIdxs,
		MessageInfos:      file_probe_proto_msgTypes,
	}.Build()
	File_probe_proto = out.File
	file_probe_proto_rawDesc = nil
	file_probe_proto_goTypes = nil
	file_probe_proto_depIdxs = nil
}
