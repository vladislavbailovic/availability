// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: incident.proto

package model

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type Incident struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IncidentID  int32 `protobuf:"varint,1,opt,name=IncidentID,proto3" json:"IncidentID,omitempty"`
	SiteID      int32 `protobuf:"varint,2,opt,name=SiteID,proto3" json:"SiteID,omitempty"`
	DownProbeID int32 `protobuf:"varint,3,opt,name=DownProbeID,proto3" json:"DownProbeID,omitempty"`
	UpProbeID   int32 `protobuf:"varint,4,opt,name=UpProbeID,proto3" json:"UpProbeID,omitempty"`
}

func (x *Incident) Reset() {
	*x = Incident{}
	if protoimpl.UnsafeEnabled {
		mi := &file_incident_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Incident) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Incident) ProtoMessage() {}

func (x *Incident) ProtoReflect() protoreflect.Message {
	mi := &file_incident_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Incident.ProtoReflect.Descriptor instead.
func (*Incident) Descriptor() ([]byte, []int) {
	return file_incident_proto_rawDescGZIP(), []int{0}
}

func (x *Incident) GetIncidentID() int32 {
	if x != nil {
		return x.IncidentID
	}
	return 0
}

func (x *Incident) GetSiteID() int32 {
	if x != nil {
		return x.SiteID
	}
	return 0
}

func (x *Incident) GetDownProbeID() int32 {
	if x != nil {
		return x.DownProbeID
	}
	return 0
}

func (x *Incident) GetUpProbeID() int32 {
	if x != nil {
		return x.UpProbeID
	}
	return 0
}

type IncidentReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IncidentID int32                  `protobuf:"varint,1,opt,name=IncidentID,json=incident_id,proto3" json:"IncidentID,omitempty"`
	SiteID     int32                  `protobuf:"varint,2,opt,name=SiteID,json=site_id,proto3" json:"SiteID,omitempty"`
	URL        string                 `protobuf:"bytes,3,opt,name=URL,json=url,proto3" json:"URL,omitempty"`
	Started    *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=Started,json=started,proto3" json:"Started,omitempty"`
	Err        HttpErr                `protobuf:"varint,6,opt,name=Err,json=err,proto3,enum=definitions.HttpErr" json:"Err,omitempty"`
	Msg        string                 `protobuf:"bytes,7,opt,name=Msg,json=msg,proto3" json:"Msg,omitempty"`
	Ended      *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=Ended,json=ended,proto3" json:"Ended,omitempty"`
}

func (x *IncidentReport) Reset() {
	*x = IncidentReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_incident_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncidentReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncidentReport) ProtoMessage() {}

func (x *IncidentReport) ProtoReflect() protoreflect.Message {
	mi := &file_incident_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncidentReport.ProtoReflect.Descriptor instead.
func (*IncidentReport) Descriptor() ([]byte, []int) {
	return file_incident_proto_rawDescGZIP(), []int{1}
}

func (x *IncidentReport) GetIncidentID() int32 {
	if x != nil {
		return x.IncidentID
	}
	return 0
}

func (x *IncidentReport) GetSiteID() int32 {
	if x != nil {
		return x.SiteID
	}
	return 0
}

func (x *IncidentReport) GetURL() string {
	if x != nil {
		return x.URL
	}
	return ""
}

func (x *IncidentReport) GetStarted() *timestamppb.Timestamp {
	if x != nil {
		return x.Started
	}
	return nil
}

func (x *IncidentReport) GetErr() HttpErr {
	if x != nil {
		return x.Err
	}
	return HttpErr_HTTPERR_NONE
}

func (x *IncidentReport) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *IncidentReport) GetEnded() *timestamppb.Timestamp {
	if x != nil {
		return x.Ended
	}
	return nil
}

type PeriodicIncidentReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start     *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=Start,json=start,proto3" json:"Start,omitempty"`
	End       *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=End,json=end,proto3,oneof" json:"End,omitempty"`
	Incidents []*IncidentReport      `protobuf:"bytes,3,rep,name=Incidents,json=incidents,proto3" json:"Incidents,omitempty"`
}

func (x *PeriodicIncidentReport) Reset() {
	*x = PeriodicIncidentReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_incident_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeriodicIncidentReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeriodicIncidentReport) ProtoMessage() {}

func (x *PeriodicIncidentReport) ProtoReflect() protoreflect.Message {
	mi := &file_incident_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeriodicIncidentReport.ProtoReflect.Descriptor instead.
func (*PeriodicIncidentReport) Descriptor() ([]byte, []int) {
	return file_incident_proto_rawDescGZIP(), []int{2}
}

func (x *PeriodicIncidentReport) GetStart() *timestamppb.Timestamp {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *PeriodicIncidentReport) GetEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.End
	}
	return nil
}

func (x *PeriodicIncidentReport) GetIncidents() []*IncidentReport {
	if x != nil {
		return x.Incidents
	}
	return nil
}

var File_incident_proto protoreflect.FileDescriptor

var file_incident_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0b, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d,
	0x68, 0x74, 0x74, 0x70, 0x65, 0x72, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x01,
	0x0a, 0x08, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x49, 0x6e,
	0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x69,
	0x74, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x53, 0x69, 0x74, 0x65,
	0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x6f, 0x77, 0x6e, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x49,
	0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x44, 0x6f, 0x77, 0x6e, 0x50, 0x72, 0x6f,
	0x62, 0x65, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x70, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x49,
	0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x55, 0x70, 0x50, 0x72, 0x6f, 0x62, 0x65,
	0x49, 0x44, 0x22, 0xfe, 0x01, 0x0a, 0x0e, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1f, 0x0a, 0x0a, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x69, 0x6e, 0x63, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x06, 0x53, 0x69, 0x74, 0x65, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x69, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x55, 0x52, 0x4c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x12, 0x34, 0x0a, 0x07, 0x53, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x12, 0x26, 0x0a, 0x03, 0x45, 0x72, 0x72, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x45, 0x72, 0x72, 0x52, 0x03, 0x65, 0x72, 0x72, 0x12,
	0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x12, 0x30, 0x0a, 0x05, 0x45, 0x6e, 0x64, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x65, 0x6e,
	0x64, 0x65, 0x64, 0x22, 0xc0, 0x01, 0x0a, 0x16, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x69, 0x63,
	0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x30,
	0x0a, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x12, 0x31, 0x0a, 0x03, 0x45, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x03, 0x65, 0x6e, 0x64,
	0x88, 0x01, 0x01, 0x12, 0x39, 0x0a, 0x09, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x52, 0x09, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x06,
	0x0a, 0x04, 0x5f, 0x45, 0x6e, 0x64, 0x42, 0x10, 0x5a, 0x0e, 0x70, 0x6b, 0x67, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_incident_proto_rawDescOnce sync.Once
	file_incident_proto_rawDescData = file_incident_proto_rawDesc
)

func file_incident_proto_rawDescGZIP() []byte {
	file_incident_proto_rawDescOnce.Do(func() {
		file_incident_proto_rawDescData = protoimpl.X.CompressGZIP(file_incident_proto_rawDescData)
	})
	return file_incident_proto_rawDescData
}

var file_incident_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_incident_proto_goTypes = []interface{}{
	(*Incident)(nil),               // 0: definitions.Incident
	(*IncidentReport)(nil),         // 1: definitions.IncidentReport
	(*PeriodicIncidentReport)(nil), // 2: definitions.PeriodicIncidentReport
	(*timestamppb.Timestamp)(nil),  // 3: google.protobuf.Timestamp
	(HttpErr)(0),                   // 4: definitions.HttpErr
}
var file_incident_proto_depIdxs = []int32{
	3, // 0: definitions.IncidentReport.Started:type_name -> google.protobuf.Timestamp
	4, // 1: definitions.IncidentReport.Err:type_name -> definitions.HttpErr
	3, // 2: definitions.IncidentReport.Ended:type_name -> google.protobuf.Timestamp
	3, // 3: definitions.PeriodicIncidentReport.Start:type_name -> google.protobuf.Timestamp
	3, // 4: definitions.PeriodicIncidentReport.End:type_name -> google.protobuf.Timestamp
	1, // 5: definitions.PeriodicIncidentReport.Incidents:type_name -> definitions.IncidentReport
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_incident_proto_init() }
func file_incident_proto_init() {
	if File_incident_proto != nil {
		return
	}
	file_httperr_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_incident_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Incident); i {
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
		file_incident_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncidentReport); i {
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
		file_incident_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeriodicIncidentReport); i {
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
	file_incident_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_incident_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_incident_proto_goTypes,
		DependencyIndexes: file_incident_proto_depIdxs,
		MessageInfos:      file_incident_proto_msgTypes,
	}.Build()
	File_incident_proto = out.File
	file_incident_proto_rawDesc = nil
	file_incident_proto_goTypes = nil
	file_incident_proto_depIdxs = nil
}
