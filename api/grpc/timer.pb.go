// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/grpc/timer.proto

package grpc

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

type TimerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	TriggerAt string `protobuf:"bytes,2,opt,name=trigger_at,json=triggerAt,proto3" json:"trigger_at,omitempty"`
}

func (x *TimerInfo) Reset() {
	*x = TimerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_timer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimerInfo) ProtoMessage() {}

func (x *TimerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_timer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimerInfo.ProtoReflect.Descriptor instead.
func (*TimerInfo) Descriptor() ([]byte, []int) {
	return file_api_grpc_timer_proto_rawDescGZIP(), []int{0}
}

func (x *TimerInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TimerInfo) GetTriggerAt() string {
	if x != nil {
		return x.TriggerAt
	}
	return ""
}

type GetAllPendingTimersReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllPendingTimersReq) Reset() {
	*x = GetAllPendingTimersReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_timer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllPendingTimersReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllPendingTimersReq) ProtoMessage() {}

func (x *GetAllPendingTimersReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_timer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllPendingTimersReq.ProtoReflect.Descriptor instead.
func (*GetAllPendingTimersReq) Descriptor() ([]byte, []int) {
	return file_api_grpc_timer_proto_rawDescGZIP(), []int{1}
}

type GetAllPendingTimersResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*TimerInfo `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *GetAllPendingTimersResp) Reset() {
	*x = GetAllPendingTimersResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_timer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllPendingTimersResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllPendingTimersResp) ProtoMessage() {}

func (x *GetAllPendingTimersResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_timer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllPendingTimersResp.ProtoReflect.Descriptor instead.
func (*GetAllPendingTimersResp) Descriptor() ([]byte, []int) {
	return file_api_grpc_timer_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllPendingTimersResp) GetItems() []*TimerInfo {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_api_grpc_timer_proto protoreflect.FileDescriptor

var file_api_grpc_timer_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x22, 0x3e, 0x0a,
	0x09, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x41, 0x74, 0x22, 0x18, 0x0a,
	0x16, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x54, 0x69,
	0x6d, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x22, 0x41, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x26, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0x5f, 0x0a, 0x05, 0x54, 0x69,
	0x6d, 0x65, 0x72, 0x12, 0x56, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x65, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x73, 0x12, 0x1d, 0x2e, 0x74, 0x69, 0x6d,
	0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x54, 0x69, 0x6d, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x74, 0x69, 0x6d, 0x65,
	0x72, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x54,
	0x69, 0x6d, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x2f, 0x5a, 0x2d, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x6f, 0x73, 0x65, 0x70, 0x68,
	0x7a, 0x78, 0x79, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_grpc_timer_proto_rawDescOnce sync.Once
	file_api_grpc_timer_proto_rawDescData = file_api_grpc_timer_proto_rawDesc
)

func file_api_grpc_timer_proto_rawDescGZIP() []byte {
	file_api_grpc_timer_proto_rawDescOnce.Do(func() {
		file_api_grpc_timer_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_grpc_timer_proto_rawDescData)
	})
	return file_api_grpc_timer_proto_rawDescData
}

var file_api_grpc_timer_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_grpc_timer_proto_goTypes = []interface{}{
	(*TimerInfo)(nil),               // 0: timer.TimerInfo
	(*GetAllPendingTimersReq)(nil),  // 1: timer.GetAllPendingTimersReq
	(*GetAllPendingTimersResp)(nil), // 2: timer.GetAllPendingTimersResp
}
var file_api_grpc_timer_proto_depIdxs = []int32{
	0, // 0: timer.GetAllPendingTimersResp.items:type_name -> timer.TimerInfo
	1, // 1: timer.Timer.GetAllPendingTimers:input_type -> timer.GetAllPendingTimersReq
	2, // 2: timer.Timer.GetAllPendingTimers:output_type -> timer.GetAllPendingTimersResp
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_grpc_timer_proto_init() }
func file_api_grpc_timer_proto_init() {
	if File_api_grpc_timer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_grpc_timer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimerInfo); i {
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
		file_api_grpc_timer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllPendingTimersReq); i {
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
		file_api_grpc_timer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllPendingTimersResp); i {
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
			RawDescriptor: file_api_grpc_timer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_grpc_timer_proto_goTypes,
		DependencyIndexes: file_api_grpc_timer_proto_depIdxs,
		MessageInfos:      file_api_grpc_timer_proto_msgTypes,
	}.Build()
	File_api_grpc_timer_proto = out.File
	file_api_grpc_timer_proto_rawDesc = nil
	file_api_grpc_timer_proto_goTypes = nil
	file_api_grpc_timer_proto_depIdxs = nil
}
