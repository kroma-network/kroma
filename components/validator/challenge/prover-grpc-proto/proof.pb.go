// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: proof.proto

package l2output

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

type ProofResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FinalPair []byte `protobuf:"bytes,1,opt,name=final_pair,json=finalPair,proto3" json:"final_pair,omitempty"`
	Proof     []byte `protobuf:"bytes,2,opt,name=proof,proto3" json:"proof,omitempty"`
}

func (x *ProofResponse) Reset() {
	*x = ProofResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proof_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProofResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProofResponse) ProtoMessage() {}

func (x *ProofResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proof_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProofResponse.ProtoReflect.Descriptor instead.
func (*ProofResponse) Descriptor() ([]byte, []int) {
	return file_proof_proto_rawDescGZIP(), []int{0}
}

func (x *ProofResponse) GetFinalPair() []byte {
	if x != nil {
		return x.FinalPair
	}
	return nil
}

func (x *ProofResponse) GetProof() []byte {
	if x != nil {
		return x.Proof
	}
	return nil
}

type ProofRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockNumberHex string `protobuf:"bytes,1,opt,name=block_number_hex,json=blockNumberHex,proto3" json:"block_number_hex,omitempty"`
}

func (x *ProofRequest) Reset() {
	*x = ProofRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proof_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProofRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProofRequest) ProtoMessage() {}

func (x *ProofRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proof_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProofRequest.ProtoReflect.Descriptor instead.
func (*ProofRequest) Descriptor() ([]byte, []int) {
	return file_proof_proto_rawDescGZIP(), []int{1}
}

func (x *ProofRequest) GetBlockNumberHex() string {
	if x != nil {
		return x.BlockNumberHex
	}
	return ""
}

var File_proof_proto protoreflect.FileDescriptor

var file_proof_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x6f, 0x66, 0x22, 0x44, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x70,
	0x61, 0x69, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x66, 0x69, 0x6e, 0x61, 0x6c,
	0x50, 0x61, 0x69, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x22, 0x38, 0x0a, 0x0c, 0x50, 0x72,
	0x6f, 0x6f, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x68, 0x65, 0x78, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x48, 0x65, 0x78, 0x32, 0x3d, 0x0a, 0x05, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x12, 0x34, 0x0a,
	0x05, 0x50, 0x72, 0x6f, 0x76, 0x65, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x2e, 0x50,
	0x72, 0x6f, 0x6f, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72,
	0x6f, 0x6f, 0x66, 0x2e, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x6c, 0x32, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proof_proto_rawDescOnce sync.Once
	file_proof_proto_rawDescData = file_proof_proto_rawDesc
)

func file_proof_proto_rawDescGZIP() []byte {
	file_proof_proto_rawDescOnce.Do(func() {
		file_proof_proto_rawDescData = protoimpl.X.CompressGZIP(file_proof_proto_rawDescData)
	})
	return file_proof_proto_rawDescData
}

var file_proof_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proof_proto_goTypes = []interface{}{
	(*ProofResponse)(nil), // 0: proof.ProofResponse
	(*ProofRequest)(nil),  // 1: proof.ProofRequest
}
var file_proof_proto_depIdxs = []int32{
	1, // 0: proof.Proof.Prove:input_type -> proof.ProofRequest
	0, // 1: proof.Proof.Prove:output_type -> proof.ProofResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proof_proto_init() }
func file_proof_proto_init() {
	if File_proof_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proof_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProofResponse); i {
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
		file_proof_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProofRequest); i {
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
			RawDescriptor: file_proof_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proof_proto_goTypes,
		DependencyIndexes: file_proof_proto_depIdxs,
		MessageInfos:      file_proof_proto_msgTypes,
	}.Build()
	File_proof_proto = out.File
	file_proof_proto_rawDesc = nil
	file_proof_proto_goTypes = nil
	file_proof_proto_depIdxs = nil
}
