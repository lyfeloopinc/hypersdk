// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: dsmr/dsmr.proto

package dsmr

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

type Chunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Producer     []byte         `protobuf:"bytes,1,opt,name=producer,proto3" json:"producer,omitempty"`
	Expiry       uint64         `protobuf:"varint,2,opt,name=expiry,proto3" json:"expiry,omitempty"`
	Beneficiary  []byte         `protobuf:"bytes,3,opt,name=beneficiary,proto3" json:"beneficiary,omitempty"`
	Transactions []*Transaction `protobuf:"bytes,4,rep,name=transactions,proto3" json:"transactions,omitempty"`
	Signer       []byte         `protobuf:"bytes,5,opt,name=signer,proto3" json:"signer,omitempty"`
	Signature    []byte         `protobuf:"bytes,6,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *Chunk) Reset() {
	*x = Chunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dsmr_dsmr_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chunk) ProtoMessage() {}

func (x *Chunk) ProtoReflect() protoreflect.Message {
	mi := &file_dsmr_dsmr_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chunk.ProtoReflect.Descriptor instead.
func (*Chunk) Descriptor() ([]byte, []int) {
	return file_dsmr_dsmr_proto_rawDescGZIP(), []int{0}
}

func (x *Chunk) GetProducer() []byte {
	if x != nil {
		return x.Producer
	}
	return nil
}

func (x *Chunk) GetExpiry() uint64 {
	if x != nil {
		return x.Expiry
	}
	return 0
}

func (x *Chunk) GetBeneficiary() []byte {
	if x != nil {
		return x.Beneficiary
	}
	return nil
}

func (x *Chunk) GetTransactions() []*Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

func (x *Chunk) GetSigner() []byte {
	if x != nil {
		return x.Signer
	}
	return nil
}

func (x *Chunk) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bytes []byte `protobuf:"bytes,1,opt,name=bytes,proto3" json:"bytes,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dsmr_dsmr_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_dsmr_dsmr_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_dsmr_dsmr_proto_rawDescGZIP(), []int{1}
}

func (x *Transaction) GetBytes() []byte {
	if x != nil {
		return x.Bytes
	}
	return nil
}

type ChunkSignature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChunkId   []byte `protobuf:"bytes,1,opt,name=chunk_id,json=chunkId,proto3" json:"chunk_id,omitempty"`
	Producer  []byte `protobuf:"bytes,2,opt,name=producer,proto3" json:"producer,omitempty"`
	Expiry    uint64 `protobuf:"varint,3,opt,name=expiry,proto3" json:"expiry,omitempty"`
	Signer    []byte `protobuf:"bytes,4,opt,name=signer,proto3" json:"signer,omitempty"` // bls public key
	Signature []byte `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *ChunkSignature) Reset() {
	*x = ChunkSignature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dsmr_dsmr_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChunkSignature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChunkSignature) ProtoMessage() {}

func (x *ChunkSignature) ProtoReflect() protoreflect.Message {
	mi := &file_dsmr_dsmr_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChunkSignature.ProtoReflect.Descriptor instead.
func (*ChunkSignature) Descriptor() ([]byte, []int) {
	return file_dsmr_dsmr_proto_rawDescGZIP(), []int{2}
}

func (x *ChunkSignature) GetChunkId() []byte {
	if x != nil {
		return x.ChunkId
	}
	return nil
}

func (x *ChunkSignature) GetProducer() []byte {
	if x != nil {
		return x.Producer
	}
	return nil
}

func (x *ChunkSignature) GetExpiry() uint64 {
	if x != nil {
		return x.Expiry
	}
	return 0
}

func (x *ChunkSignature) GetSigner() []byte {
	if x != nil {
		return x.Signer
	}
	return nil
}

func (x *ChunkSignature) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type ChunkCertificate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChunkId   []byte `protobuf:"bytes,1,opt,name=chunk_id,json=chunkId,proto3" json:"chunk_id,omitempty"`
	Producer  []byte `protobuf:"bytes,2,opt,name=producer,proto3" json:"producer,omitempty"`
	Expiry    uint64 `protobuf:"varint,3,opt,name=expiry,proto3" json:"expiry,omitempty"`
	Signers   []byte `protobuf:"bytes,4,opt,name=signers,proto3" json:"signers,omitempty"` // bitset
	Signature []byte `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *ChunkCertificate) Reset() {
	*x = ChunkCertificate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dsmr_dsmr_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChunkCertificate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChunkCertificate) ProtoMessage() {}

func (x *ChunkCertificate) ProtoReflect() protoreflect.Message {
	mi := &file_dsmr_dsmr_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChunkCertificate.ProtoReflect.Descriptor instead.
func (*ChunkCertificate) Descriptor() ([]byte, []int) {
	return file_dsmr_dsmr_proto_rawDescGZIP(), []int{3}
}

func (x *ChunkCertificate) GetChunkId() []byte {
	if x != nil {
		return x.ChunkId
	}
	return nil
}

func (x *ChunkCertificate) GetProducer() []byte {
	if x != nil {
		return x.Producer
	}
	return nil
}

func (x *ChunkCertificate) GetExpiry() uint64 {
	if x != nil {
		return x.Expiry
	}
	return 0
}

func (x *ChunkCertificate) GetSigners() []byte {
	if x != nil {
		return x.Signers
	}
	return nil
}

func (x *ChunkCertificate) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type ExecutedChunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChunkId      []byte         `protobuf:"bytes,1,opt,name=chunk_id,json=chunkId,proto3" json:"chunk_id,omitempty"`
	Beneficiary  []byte         `protobuf:"bytes,2,opt,name=beneficiary,proto3" json:"beneficiary,omitempty"`
	Transactions []*Transaction `protobuf:"bytes,3,rep,name=transactions,proto3" json:"transactions,omitempty"`
	WarpResults  []byte         `protobuf:"bytes,4,opt,name=warp_results,json=warpResults,proto3" json:"warp_results,omitempty"` // bitset
}

func (x *ExecutedChunk) Reset() {
	*x = ExecutedChunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dsmr_dsmr_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecutedChunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutedChunk) ProtoMessage() {}

func (x *ExecutedChunk) ProtoReflect() protoreflect.Message {
	mi := &file_dsmr_dsmr_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutedChunk.ProtoReflect.Descriptor instead.
func (*ExecutedChunk) Descriptor() ([]byte, []int) {
	return file_dsmr_dsmr_proto_rawDescGZIP(), []int{4}
}

func (x *ExecutedChunk) GetChunkId() []byte {
	if x != nil {
		return x.ChunkId
	}
	return nil
}

func (x *ExecutedChunk) GetBeneficiary() []byte {
	if x != nil {
		return x.Beneficiary
	}
	return nil
}

func (x *ExecutedChunk) GetTransactions() []*Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

func (x *ExecutedChunk) GetWarpResults() []byte {
	if x != nil {
		return x.WarpResults
	}
	return nil
}

type GetChunkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChunkId []byte `protobuf:"bytes,1,opt,name=chunk_id,json=chunkId,proto3" json:"chunk_id,omitempty"`
}

func (x *GetChunkRequest) Reset() {
	*x = GetChunkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dsmr_dsmr_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChunkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChunkRequest) ProtoMessage() {}

func (x *GetChunkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dsmr_dsmr_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChunkRequest.ProtoReflect.Descriptor instead.
func (*GetChunkRequest) Descriptor() ([]byte, []int) {
	return file_dsmr_dsmr_proto_rawDescGZIP(), []int{5}
}

func (x *GetChunkRequest) GetChunkId() []byte {
	if x != nil {
		return x.ChunkId
	}
	return nil
}

var File_dsmr_dsmr_proto protoreflect.FileDescriptor

var file_dsmr_dsmr_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x64, 0x73, 0x6d, 0x72, 0x2f, 0x64, 0x73, 0x6d, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x64, 0x73, 0x6d, 0x72, 0x22, 0xca, 0x01, 0x0a, 0x05, 0x43, 0x68, 0x75, 0x6e,
	0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x12, 0x16, 0x0a,
	0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x63,
	0x69, 0x61, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x62, 0x65, 0x6e, 0x65,
	0x66, 0x69, 0x63, 0x69, 0x61, 0x72, 0x79, 0x12, 0x35, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x64, 0x73, 0x6d, 0x72, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06,
	0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x22, 0x23, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x22, 0x95, 0x01, 0x0a, 0x0e, 0x43, 0x68,
	0x75, 0x6e, 0x6b, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07,
	0x63, 0x68, 0x75, 0x6e, 0x6b, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x69, 0x67, 0x6e, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x69, 0x67,
	0x6e, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x22, 0x99, 0x01, 0x0a, 0x10, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x43, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x12, 0x16, 0x0a,
	0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x73, 0x12,
	0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0xa6, 0x01,
	0x0a, 0x0d, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x64, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12,
	0x19, 0x0a, 0x08, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x07, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x65,
	0x6e, 0x65, 0x66, 0x69, 0x63, 0x69, 0x61, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0b, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x63, 0x69, 0x61, 0x72, 0x79, 0x12, 0x35, 0x0a, 0x0c,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x64, 0x73, 0x6d, 0x72, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x61, 0x72, 0x70, 0x5f, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x77, 0x61, 0x72, 0x70, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0x2c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x68, 0x75,
	0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x75,
	0x6e, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x68, 0x75,
	0x6e, 0x6b, 0x49, 0x64, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x61, 0x76, 0x61, 0x2d, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x68, 0x79, 0x70, 0x65,
	0x72, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x64, 0x73,
	0x6d, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dsmr_dsmr_proto_rawDescOnce sync.Once
	file_dsmr_dsmr_proto_rawDescData = file_dsmr_dsmr_proto_rawDesc
)

func file_dsmr_dsmr_proto_rawDescGZIP() []byte {
	file_dsmr_dsmr_proto_rawDescOnce.Do(func() {
		file_dsmr_dsmr_proto_rawDescData = protoimpl.X.CompressGZIP(file_dsmr_dsmr_proto_rawDescData)
	})
	return file_dsmr_dsmr_proto_rawDescData
}

var file_dsmr_dsmr_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_dsmr_dsmr_proto_goTypes = []interface{}{
	(*Chunk)(nil),            // 0: dsmr.Chunk
	(*Transaction)(nil),      // 1: dsmr.Transaction
	(*ChunkSignature)(nil),   // 2: dsmr.ChunkSignature
	(*ChunkCertificate)(nil), // 3: dsmr.ChunkCertificate
	(*ExecutedChunk)(nil),    // 4: dsmr.ExecutedChunk
	(*GetChunkRequest)(nil),  // 5: dsmr.GetChunkRequest
}
var file_dsmr_dsmr_proto_depIdxs = []int32{
	1, // 0: dsmr.Chunk.transactions:type_name -> dsmr.Transaction
	1, // 1: dsmr.ExecutedChunk.transactions:type_name -> dsmr.Transaction
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_dsmr_dsmr_proto_init() }
func file_dsmr_dsmr_proto_init() {
	if File_dsmr_dsmr_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dsmr_dsmr_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chunk); i {
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
		file_dsmr_dsmr_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
		file_dsmr_dsmr_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChunkSignature); i {
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
		file_dsmr_dsmr_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChunkCertificate); i {
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
		file_dsmr_dsmr_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecutedChunk); i {
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
		file_dsmr_dsmr_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChunkRequest); i {
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
			RawDescriptor: file_dsmr_dsmr_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_dsmr_dsmr_proto_goTypes,
		DependencyIndexes: file_dsmr_dsmr_proto_depIdxs,
		MessageInfos:      file_dsmr_dsmr_proto_msgTypes,
	}.Build()
	File_dsmr_dsmr_proto = out.File
	file_dsmr_dsmr_proto_rawDesc = nil
	file_dsmr_dsmr_proto_goTypes = nil
	file_dsmr_dsmr_proto_depIdxs = nil
}
