// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: social/social_service.proto

package social

import (
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

var File_social_social_service_proto protoreflect.FileDescriptor

var file_social_social_service_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x2f, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x2f, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x32, 0x84, 0x03, 0x0a, 0x0d, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x6f, 0x63, 0x69, 0x61, 0x6c, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c,
	0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45,
	0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x12, 0x1b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x42, 0x79, 0x49,
	0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53,
	0x6f, 0x63, 0x69, 0x61, 0x6c, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x4b, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c,
	0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53,
	0x6f, 0x63, 0x69, 0x61, 0x6c, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x42,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a,
	0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x12, 0x1e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6f, 0x63, 0x69,
	0x61, 0x6c, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x42, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x73, 0x61, 0x74,
	0x61, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73,
	0x74, 0x75, 0x62, 0x73, 0x2f, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_social_social_service_proto_goTypes = []any{
	(*CreateSocialRequest)(nil),     // 0: proto.CreateSocialRequest
	(*GetSocialByIDRequest)(nil),    // 1: proto.GetSocialByIDRequest
	(*GetAllSocialRequest)(nil),     // 2: proto.GetAllSocialRequest
	(*UpdateSocialByIDRequest)(nil), // 3: proto.UpdateSocialByIDRequest
	(*DeleteSocialByIDRequest)(nil), // 4: proto.DeleteSocialByIDRequest
	(*SocialBaseResponse)(nil),      // 5: proto.SocialBaseResponse
	(*GetAllSocialResponse)(nil),    // 6: proto.GetAllSocialResponse
}
var file_social_social_service_proto_depIdxs = []int32{
	0, // 0: proto.SocialService.CreateSocial:input_type -> proto.CreateSocialRequest
	1, // 1: proto.SocialService.GetSocial:input_type -> proto.GetSocialByIDRequest
	2, // 2: proto.SocialService.GetAllSocial:input_type -> proto.GetAllSocialRequest
	3, // 3: proto.SocialService.UpdateSocial:input_type -> proto.UpdateSocialByIDRequest
	4, // 4: proto.SocialService.DeleteSocial:input_type -> proto.DeleteSocialByIDRequest
	5, // 5: proto.SocialService.CreateSocial:output_type -> proto.SocialBaseResponse
	5, // 6: proto.SocialService.GetSocial:output_type -> proto.SocialBaseResponse
	6, // 7: proto.SocialService.GetAllSocial:output_type -> proto.GetAllSocialResponse
	5, // 8: proto.SocialService.UpdateSocial:output_type -> proto.SocialBaseResponse
	5, // 9: proto.SocialService.DeleteSocial:output_type -> proto.SocialBaseResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_social_social_service_proto_init() }
func file_social_social_service_proto_init() {
	if File_social_social_service_proto != nil {
		return
	}
	file_social_payload_messages_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_social_social_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_social_social_service_proto_goTypes,
		DependencyIndexes: file_social_social_service_proto_depIdxs,
	}.Build()
	File_social_social_service_proto = out.File
	file_social_social_service_proto_rawDesc = nil
	file_social_social_service_proto_goTypes = nil
	file_social_social_service_proto_depIdxs = nil
}