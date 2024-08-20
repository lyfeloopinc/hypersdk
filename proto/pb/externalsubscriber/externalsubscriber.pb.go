// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: externalsubscriber/externalsubscriber.proto

package externalsubscriber

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockData []byte `protobuf:"bytes,1,opt,name=block_data,json=blockData,proto3" json:"block_data,omitempty"`
	Results   []byte `protobuf:"bytes,2,opt,name=results,proto3" json:"results,omitempty"`
}

func (x *BlockRequest) Reset() {
	*x = BlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_externalsubscriber_externalsubscriber_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockRequest) ProtoMessage() {}

func (x *BlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_externalsubscriber_externalsubscriber_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockRequest.ProtoReflect.Descriptor instead.
func (*BlockRequest) Descriptor() ([]byte, []int) {
	return file_externalsubscriber_externalsubscriber_proto_rawDescGZIP(), []int{0}
}

func (x *BlockRequest) GetBlockData() []byte {
	if x != nil {
		return x.BlockData
	}
	return nil
}

func (x *BlockRequest) GetResults() []byte {
	if x != nil {
		return x.Results
	}
	return nil
}

type InitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NetworkId uint32 `protobuf:"varint,1,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
	ChainId   []byte `protobuf:"bytes,2,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	Genesis   []byte `protobuf:"bytes,3,opt,name=genesis,proto3" json:"genesis,omitempty"`
}

func (x *InitRequest) Reset() {
	*x = InitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_externalsubscriber_externalsubscriber_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitRequest) ProtoMessage() {}

func (x *InitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_externalsubscriber_externalsubscriber_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitRequest.ProtoReflect.Descriptor instead.
func (*InitRequest) Descriptor() ([]byte, []int) {
	return file_externalsubscriber_externalsubscriber_proto_rawDescGZIP(), []int{1}
}

func (x *InitRequest) GetNetworkId() uint32 {
	if x != nil {
		return x.NetworkId
	}
	return 0
}

func (x *InitRequest) GetChainId() []byte {
	if x != nil {
		return x.ChainId
	}
	return nil
}

func (x *InitRequest) GetGenesis() []byte {
	if x != nil {
		return x.Genesis
	}
	return nil
}

var File_externalsubscriber_externalsubscriber_proto protoreflect.FileDescriptor

var file_externalsubscriber_externalsubscriber_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x72, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x73, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x47,
	0x0a, 0x0c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a,
	0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0x61, 0x0a, 0x0b, 0x49, 0x6e, 0x69, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x07, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x32, 0xa5, 0x01, 0x0a, 0x12, 0x45,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x72, 0x12, 0x48, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x12, 0x20, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x45, 0x0a, 0x0a, 0x49,
	0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x12, 0x1f, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x2e, 0x49,
	0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x76, 0x61, 0x2d, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x68, 0x79, 0x70, 0x65, 0x72, 0x73,
	0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_externalsubscriber_externalsubscriber_proto_rawDescOnce sync.Once
	file_externalsubscriber_externalsubscriber_proto_rawDescData = file_externalsubscriber_externalsubscriber_proto_rawDesc
)

func file_externalsubscriber_externalsubscriber_proto_rawDescGZIP() []byte {
	file_externalsubscriber_externalsubscriber_proto_rawDescOnce.Do(func() {
		file_externalsubscriber_externalsubscriber_proto_rawDescData = protoimpl.X.CompressGZIP(file_externalsubscriber_externalsubscriber_proto_rawDescData)
	})
	return file_externalsubscriber_externalsubscriber_proto_rawDescData
}

var file_externalsubscriber_externalsubscriber_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_externalsubscriber_externalsubscriber_proto_goTypes = []interface{}{
	(*BlockRequest)(nil),  // 0: externalsubscriber.BlockRequest
	(*InitRequest)(nil),   // 1: externalsubscriber.InitRequest
	(*emptypb.Empty)(nil), // 2: google.protobuf.Empty
}
var file_externalsubscriber_externalsubscriber_proto_depIdxs = []int32{
	0, // 0: externalsubscriber.ExternalSubscriber.ProcessBlock:input_type -> externalsubscriber.BlockRequest
	1, // 1: externalsubscriber.ExternalSubscriber.Initialize:input_type -> externalsubscriber.InitRequest
	2, // 2: externalsubscriber.ExternalSubscriber.ProcessBlock:output_type -> google.protobuf.Empty
	2, // 3: externalsubscriber.ExternalSubscriber.Initialize:output_type -> google.protobuf.Empty
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_externalsubscriber_externalsubscriber_proto_init() }
func file_externalsubscriber_externalsubscriber_proto_init() {
	if File_externalsubscriber_externalsubscriber_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_externalsubscriber_externalsubscriber_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockRequest); i {
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
		file_externalsubscriber_externalsubscriber_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitRequest); i {
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
			RawDescriptor: file_externalsubscriber_externalsubscriber_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_externalsubscriber_externalsubscriber_proto_goTypes,
		DependencyIndexes: file_externalsubscriber_externalsubscriber_proto_depIdxs,
		MessageInfos:      file_externalsubscriber_externalsubscriber_proto_msgTypes,
	}.Build()
	File_externalsubscriber_externalsubscriber_proto = out.File
	file_externalsubscriber_externalsubscriber_proto_rawDesc = nil
	file_externalsubscriber_externalsubscriber_proto_goTypes = nil
	file_externalsubscriber_externalsubscriber_proto_depIdxs = nil
}
