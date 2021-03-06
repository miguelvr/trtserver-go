// Copyright (c) 2018, NVIDIA CORPORATION. All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions
// are met:
//  * Redistributions of source code must retain the above copyright
//    notice, this list of conditions and the following disclaimer.
//  * Redistributions in binary form must reproduce the above copyright
//    notice, this list of conditions and the following disclaimer in the
//    documentation and/or other materials provided with the distribution.
//  * Neither the name of NVIDIA CORPORATION nor the names of its
//    contributors may be used to endorse or promote products derived
//    from this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS ``AS IS'' AND ANY
// EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR
// PURPOSE ARE DISCLAIMED.  IN NO EVENT SHALL THE COPYRIGHT OWNER OR
// CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL,
// EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,
// PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
// PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY
// OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        (unknown)
// source: request_status.proto

package nvidia_inferenceserver

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

//@@
//@@.. cpp:enum:: RequestStatusCode
//@@
//@@   Status codes returned for inference server requests. The
//@@   :cpp:enumerator:`RequestStatusCode::SUCCESS` status code indicates
//@@   not error, all other codes indicate an error.
//@@
type RequestStatusCode int32

const (
	//@@  .. cpp:enumerator:: RequestStatusCode::INVALID = 0
	//@@
	//@@     Invalid status. Used internally but should not be returned as
	//@@     part of a :cpp:var:`RequestStatus`.
	//@@
	RequestStatusCode_INVALID RequestStatusCode = 0
	//@@  .. cpp:enumerator:: RequestStatusCode::SUCCESS = 1
	//@@
	//@@     Error code indicating success.
	//@@
	RequestStatusCode_SUCCESS RequestStatusCode = 1
	//@@  .. cpp:enumerator:: RequestStatusCode::UNKNOWN = 2
	//@@
	//@@     Error code indicating an unknown failure.
	//@@
	RequestStatusCode_UNKNOWN RequestStatusCode = 2
	//@@  .. cpp:enumerator:: RequestStatusCode::INTERNAL = 3
	//@@
	//@@     Error code indicating an internal failure.
	//@@
	RequestStatusCode_INTERNAL RequestStatusCode = 3
	//@@  .. cpp:enumerator:: RequestStatusCode::NOT_FOUND = 4
	//@@
	//@@     Error code indicating a resource or request was not found.
	//@@
	RequestStatusCode_NOT_FOUND RequestStatusCode = 4
	//@@  .. cpp:enumerator:: RequestStatusCode::INVALID_ARG = 5
	//@@
	//@@     Error code indicating a failure caused by an unknown argument or
	//@@     value.
	//@@
	RequestStatusCode_INVALID_ARG RequestStatusCode = 5
	//@@  .. cpp:enumerator:: RequestStatusCode::UNAVAILABLE = 6
	//@@
	//@@     Error code indicating an unavailable resource.
	//@@
	RequestStatusCode_UNAVAILABLE RequestStatusCode = 6
	//@@  .. cpp:enumerator:: RequestStatusCode::UNSUPPORTED = 7
	//@@
	//@@     Error code indicating an unsupported request or operation.
	//@@
	RequestStatusCode_UNSUPPORTED RequestStatusCode = 7
	//@@  .. cpp:enumerator:: RequestStatusCode::ALREADY_EXISTS = 8
	//@@
	//@@     Error code indicating an already existing resource.
	//@@
	RequestStatusCode_ALREADY_EXISTS RequestStatusCode = 8
)

// Enum value maps for RequestStatusCode.
var (
	RequestStatusCode_name = map[int32]string{
		0: "INVALID",
		1: "SUCCESS",
		2: "UNKNOWN",
		3: "INTERNAL",
		4: "NOT_FOUND",
		5: "INVALID_ARG",
		6: "UNAVAILABLE",
		7: "UNSUPPORTED",
		8: "ALREADY_EXISTS",
	}
	RequestStatusCode_value = map[string]int32{
		"INVALID":        0,
		"SUCCESS":        1,
		"UNKNOWN":        2,
		"INTERNAL":       3,
		"NOT_FOUND":      4,
		"INVALID_ARG":    5,
		"UNAVAILABLE":    6,
		"UNSUPPORTED":    7,
		"ALREADY_EXISTS": 8,
	}
)

func (x RequestStatusCode) Enum() *RequestStatusCode {
	p := new(RequestStatusCode)
	*p = x
	return p
}

func (x RequestStatusCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RequestStatusCode) Descriptor() protoreflect.EnumDescriptor {
	return file_request_status_proto_enumTypes[0].Descriptor()
}

func (RequestStatusCode) Type() protoreflect.EnumType {
	return &file_request_status_proto_enumTypes[0]
}

func (x RequestStatusCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RequestStatusCode.Descriptor instead.
func (RequestStatusCode) EnumDescriptor() ([]byte, []int) {
	return file_request_status_proto_rawDescGZIP(), []int{0}
}

//@@
//@@.. cpp:var:: message RequestStatus
//@@
//@@   Status returned for all inference server requests. The
//@@   RequestStatus provides a :cpp:enum:`RequestStatusCode`, an
//@@   optional status message, and server and request IDs.
//@@
type RequestStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//@@  .. cpp:var:: RequestStatusCode code
	//@@
	//@@     The status code.
	//@@
	Code RequestStatusCode `protobuf:"varint,1,opt,name=code,proto3,enum=nvidia.inferenceserver.RequestStatusCode" json:"code,omitempty"`
	//@@  .. cpp:var:: string msg
	//@@
	//@@     The optional status message.
	//@@
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	//@@  .. cpp:var:: string server_id
	//@@
	//@@     The identifying string for the server that is returning
	//@@     this status.
	//@@
	ServerId string `protobuf:"bytes,3,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	//@@  .. cpp:var:: string request_id
	//@@
	//@@     Unique identifier for the request assigned by the inference
	//@@     server. Value 0 (zero) indicates the request ID is not known.
	//@@
	RequestId uint64 `protobuf:"varint,4,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
}

func (x *RequestStatus) Reset() {
	*x = RequestStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_status_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestStatus) ProtoMessage() {}

func (x *RequestStatus) ProtoReflect() protoreflect.Message {
	mi := &file_request_status_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestStatus.ProtoReflect.Descriptor instead.
func (*RequestStatus) Descriptor() ([]byte, []int) {
	return file_request_status_proto_rawDescGZIP(), []int{0}
}

func (x *RequestStatus) GetCode() RequestStatusCode {
	if x != nil {
		return x.Code
	}
	return RequestStatusCode_INVALID
}

func (x *RequestStatus) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *RequestStatus) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *RequestStatus) GetRequestId() uint64 {
	if x != nil {
		return x.RequestId
	}
	return 0
}

var File_request_status_proto protoreflect.FileDescriptor

var file_request_status_proto_rawDesc = []byte{
	0x0a, 0x14, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x6e, 0x76, 0x69, 0x64, 0x69, 0x61, 0x2e, 0x69,
	0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x9c,
	0x01, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x3d, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29,
	0x2e, 0x6e, 0x76, 0x69, 0x64, 0x69, 0x61, 0x2e, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x2a, 0x9e, 0x01,
	0x0a, 0x11, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x0b, 0x0a,
	0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x4e,
	0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x4f, 0x54, 0x5f,
	0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x04, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x4e, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x5f, 0x41, 0x52, 0x47, 0x10, 0x05, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x41, 0x56,
	0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x06, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53,
	0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x10, 0x07, 0x12, 0x12, 0x0a, 0x0e, 0x41, 0x4c,
	0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x10, 0x08, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_request_status_proto_rawDescOnce sync.Once
	file_request_status_proto_rawDescData = file_request_status_proto_rawDesc
)

func file_request_status_proto_rawDescGZIP() []byte {
	file_request_status_proto_rawDescOnce.Do(func() {
		file_request_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_request_status_proto_rawDescData)
	})
	return file_request_status_proto_rawDescData
}

var file_request_status_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_request_status_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_request_status_proto_goTypes = []interface{}{
	(RequestStatusCode)(0), // 0: nvidia.inferenceserver.RequestStatusCode
	(*RequestStatus)(nil),  // 1: nvidia.inferenceserver.RequestStatus
}
var file_request_status_proto_depIdxs = []int32{
	0, // 0: nvidia.inferenceserver.RequestStatus.code:type_name -> nvidia.inferenceserver.RequestStatusCode
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_request_status_proto_init() }
func file_request_status_proto_init() {
	if File_request_status_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_request_status_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestStatus); i {
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
			RawDescriptor: file_request_status_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_request_status_proto_goTypes,
		DependencyIndexes: file_request_status_proto_depIdxs,
		EnumInfos:         file_request_status_proto_enumTypes,
		MessageInfos:      file_request_status_proto_msgTypes,
	}.Build()
	File_request_status_proto = out.File
	file_request_status_proto_rawDesc = nil
	file_request_status_proto_goTypes = nil
	file_request_status_proto_depIdxs = nil
}
