// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: ydb_scripting_v1.proto

package Ydb_Scripting_V1

import (
	Ydb_Scripting "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_Scripting"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_ydb_scripting_v1_proto protoreflect.FileDescriptor

var file_ydb_scripting_v1_proto_rawDesc = []byte{
	0x0a, 0x16, 0x79, 0x64, 0x62, 0x5f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6e, 0x67, 0x5f,
	0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x59, 0x64, 0x62, 0x2e, 0x53, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x31, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2f, 0x79, 0x64, 0x62, 0x5f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x9a, 0x02, 0x0a, 0x10, 0x53, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x0a, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x59, 0x71, 0x6c, 0x12, 0x20, 0x2e, 0x59, 0x64, 0x62, 0x2e,
	0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x65, 0x59, 0x71, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x59, 0x64,
	0x62, 0x2e, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x65, 0x59, 0x71, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x60,
	0x0a, 0x10, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x59,
	0x71, 0x6c, 0x12, 0x20, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x59, 0x71, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x53, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x59, 0x71, 0x6c, 0x50,
	0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01,
	0x12, 0x51, 0x0a, 0x0a, 0x45, 0x78, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x59, 0x71, 0x6c, 0x12, 0x20,
	0x2e, 0x59, 0x64, 0x62, 0x2e, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x45,
	0x78, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x59, 0x71, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x45, 0x78, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x59, 0x71, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x57, 0x0a, 0x1b, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x79, 0x64, 0x62, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79,
	0x64, 0x62, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x79, 0x64, 0x62, 0x2d,
	0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x59, 0x64, 0x62, 0x5f,
	0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_ydb_scripting_v1_proto_goTypes = []interface{}{
	(*Ydb_Scripting.ExecuteYqlRequest)(nil),         // 0: Ydb.Scripting.ExecuteYqlRequest
	(*Ydb_Scripting.ExplainYqlRequest)(nil),         // 1: Ydb.Scripting.ExplainYqlRequest
	(*Ydb_Scripting.ExecuteYqlResponse)(nil),        // 2: Ydb.Scripting.ExecuteYqlResponse
	(*Ydb_Scripting.ExecuteYqlPartialResponse)(nil), // 3: Ydb.Scripting.ExecuteYqlPartialResponse
	(*Ydb_Scripting.ExplainYqlResponse)(nil),        // 4: Ydb.Scripting.ExplainYqlResponse
}
var file_ydb_scripting_v1_proto_depIdxs = []int32{
	0, // 0: Ydb.Scripting.V1.ScriptingService.ExecuteYql:input_type -> Ydb.Scripting.ExecuteYqlRequest
	0, // 1: Ydb.Scripting.V1.ScriptingService.StreamExecuteYql:input_type -> Ydb.Scripting.ExecuteYqlRequest
	1, // 2: Ydb.Scripting.V1.ScriptingService.ExplainYql:input_type -> Ydb.Scripting.ExplainYqlRequest
	2, // 3: Ydb.Scripting.V1.ScriptingService.ExecuteYql:output_type -> Ydb.Scripting.ExecuteYqlResponse
	3, // 4: Ydb.Scripting.V1.ScriptingService.StreamExecuteYql:output_type -> Ydb.Scripting.ExecuteYqlPartialResponse
	4, // 5: Ydb.Scripting.V1.ScriptingService.ExplainYql:output_type -> Ydb.Scripting.ExplainYqlResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ydb_scripting_v1_proto_init() }
func file_ydb_scripting_v1_proto_init() {
	if File_ydb_scripting_v1_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ydb_scripting_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ydb_scripting_v1_proto_goTypes,
		DependencyIndexes: file_ydb_scripting_v1_proto_depIdxs,
	}.Build()
	File_ydb_scripting_v1_proto = out.File
	file_ydb_scripting_v1_proto_rawDesc = nil
	file_ydb_scripting_v1_proto_goTypes = nil
	file_ydb_scripting_v1_proto_depIdxs = nil
}
