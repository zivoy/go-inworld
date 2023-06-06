// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: voices.proto

package voices

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

type TTSType int32

const (
	TTSType_TTS_TYPE_STANDARD TTSType = 0
	TTSType_TTS_TYPE_ADVANCED TTSType = 1
)

// Enum value maps for TTSType.
var (
	TTSType_name = map[int32]string{
		0: "TTS_TYPE_STANDARD",
		1: "TTS_TYPE_ADVANCED",
	}
	TTSType_value = map[string]int32{
		"TTS_TYPE_STANDARD": 0,
		"TTS_TYPE_ADVANCED": 1,
	}
)

func (x TTSType) Enum() *TTSType {
	p := new(TTSType)
	*p = x
	return p
}

func (x TTSType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TTSType) Descriptor() protoreflect.EnumDescriptor {
	return file_voices_proto_enumTypes[0].Descriptor()
}

func (TTSType) Type() protoreflect.EnumType {
	return &file_voices_proto_enumTypes[0]
}

func (x TTSType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TTSType.Descriptor instead.
func (TTSType) EnumDescriptor() ([]byte, []int) {
	return file_voices_proto_rawDescGZIP(), []int{0}
}

type Gender int32

const (
	Gender_VOICE_GENDER_UNSPECIFIED Gender = 0
	Gender_VOICE_GENDER_MALE        Gender = 1
	Gender_VOICE_GENDER_FEMALE      Gender = 2
	Gender_VOICE_GENDER_NEUTRAL     Gender = 3
)

// Enum value maps for Gender.
var (
	Gender_name = map[int32]string{
		0: "VOICE_GENDER_UNSPECIFIED",
		1: "VOICE_GENDER_MALE",
		2: "VOICE_GENDER_FEMALE",
		3: "VOICE_GENDER_NEUTRAL",
	}
	Gender_value = map[string]int32{
		"VOICE_GENDER_UNSPECIFIED": 0,
		"VOICE_GENDER_MALE":        1,
		"VOICE_GENDER_FEMALE":      2,
		"VOICE_GENDER_NEUTRAL":     3,
	}
)

func (x Gender) Enum() *Gender {
	p := new(Gender)
	*p = x
	return p
}

func (x Gender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Gender) Descriptor() protoreflect.EnumDescriptor {
	return file_voices_proto_enumTypes[1].Descriptor()
}

func (Gender) Type() protoreflect.EnumType {
	return &file_voices_proto_enumTypes[1]
}

func (x Gender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Gender.Descriptor instead.
func (Gender) EnumDescriptor() ([]byte, []int) {
	return file_voices_proto_rawDescGZIP(), []int{1}
}

type Voice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Basename                string  `protobuf:"bytes,1,opt,name=basename,proto3" json:"basename,omitempty"`
	TtsType                 TTSType `protobuf:"varint,2,opt,name=ttsType,proto3,enum=ai.inworld.voices.TTSType" json:"ttsType,omitempty"`
	Gender                  Gender  `protobuf:"varint,3,opt,name=gender,proto3,enum=ai.inworld.voices.Gender" json:"gender,omitempty"`
	Pitch                   float64 `protobuf:"fixed64,4,opt,name=pitch,proto3" json:"pitch,omitempty"`
	SpeakingRate            float64 `protobuf:"fixed64,5,opt,name=speakingRate,proto3" json:"speakingRate,omitempty"`
	RoboticVoiceFilterLevel float64 `protobuf:"fixed64,6,opt,name=roboticVoiceFilterLevel,proto3" json:"roboticVoiceFilterLevel,omitempty"`
}

func (x *Voice) Reset() {
	*x = Voice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voices_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Voice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Voice) ProtoMessage() {}

func (x *Voice) ProtoReflect() protoreflect.Message {
	mi := &file_voices_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Voice.ProtoReflect.Descriptor instead.
func (*Voice) Descriptor() ([]byte, []int) {
	return file_voices_proto_rawDescGZIP(), []int{0}
}

func (x *Voice) GetBasename() string {
	if x != nil {
		return x.Basename
	}
	return ""
}

func (x *Voice) GetTtsType() TTSType {
	if x != nil {
		return x.TtsType
	}
	return TTSType_TTS_TYPE_STANDARD
}

func (x *Voice) GetGender() Gender {
	if x != nil {
		return x.Gender
	}
	return Gender_VOICE_GENDER_UNSPECIFIED
}

func (x *Voice) GetPitch() float64 {
	if x != nil {
		return x.Pitch
	}
	return 0
}

func (x *Voice) GetSpeakingRate() float64 {
	if x != nil {
		return x.SpeakingRate
	}
	return 0
}

func (x *Voice) GetRoboticVoiceFilterLevel() float64 {
	if x != nil {
		return x.RoboticVoiceFilterLevel
	}
	return 0
}

var File_voices_proto protoreflect.FileDescriptor

var file_voices_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11,
	0x61, 0x69, 0x2e, 0x69, 0x6e, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x6f, 0x69, 0x63, 0x65,
	0x73, 0x22, 0x80, 0x02, 0x0a, 0x05, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x62,
	0x61, 0x73, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62,
	0x61, 0x73, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x74, 0x74, 0x73, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x61, 0x69, 0x2e, 0x69, 0x6e,
	0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x54, 0x53,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x74, 0x74, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x31, 0x0a,
	0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e,
	0x61, 0x69, 0x2e, 0x69, 0x6e, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x6f, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x69, 0x74, 0x63, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x05, 0x70, 0x69, 0x74, 0x63, 0x68, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x69,
	0x6e, 0x67, 0x52, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x73, 0x70,
	0x65, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x61, 0x74, 0x65, 0x12, 0x38, 0x0a, 0x17, 0x72, 0x6f,
	0x62, 0x6f, 0x74, 0x69, 0x63, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x17, 0x72, 0x6f, 0x62,
	0x6f, 0x74, 0x69, 0x63, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x2a, 0x37, 0x0a, 0x07, 0x54, 0x54, 0x53, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x15, 0x0a, 0x11, 0x54, 0x54, 0x53, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x4e,
	0x44, 0x41, 0x52, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x54, 0x54, 0x53, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x41, 0x44, 0x56, 0x41, 0x4e, 0x43, 0x45, 0x44, 0x10, 0x01, 0x2a, 0x70, 0x0a,
	0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x18, 0x56, 0x4f, 0x49, 0x43, 0x45,
	0x5f, 0x47, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x56, 0x4f, 0x49, 0x43, 0x45, 0x5f, 0x47,
	0x45, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13,
	0x56, 0x4f, 0x49, 0x43, 0x45, 0x5f, 0x47, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x46, 0x45, 0x4d,
	0x41, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x56, 0x4f, 0x49, 0x43, 0x45, 0x5f, 0x47,
	0x45, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x4e, 0x45, 0x55, 0x54, 0x52, 0x41, 0x4c, 0x10, 0x03, 0x42,
	0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x69,
	0x76, 0x6f, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x69, 0x6e, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x42, 0x75, 0x66, 0x2f, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_voices_proto_rawDescOnce sync.Once
	file_voices_proto_rawDescData = file_voices_proto_rawDesc
)

func file_voices_proto_rawDescGZIP() []byte {
	file_voices_proto_rawDescOnce.Do(func() {
		file_voices_proto_rawDescData = protoimpl.X.CompressGZIP(file_voices_proto_rawDescData)
	})
	return file_voices_proto_rawDescData
}

var file_voices_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_voices_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_voices_proto_goTypes = []interface{}{
	(TTSType)(0),  // 0: ai.inworld.voices.TTSType
	(Gender)(0),   // 1: ai.inworld.voices.Gender
	(*Voice)(nil), // 2: ai.inworld.voices.Voice
}
var file_voices_proto_depIdxs = []int32{
	0, // 0: ai.inworld.voices.Voice.ttsType:type_name -> ai.inworld.voices.TTSType
	1, // 1: ai.inworld.voices.Voice.gender:type_name -> ai.inworld.voices.Gender
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_voices_proto_init() }
func file_voices_proto_init() {
	if File_voices_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_voices_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Voice); i {
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
			RawDescriptor: file_voices_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_voices_proto_goTypes,
		DependencyIndexes: file_voices_proto_depIdxs,
		EnumInfos:         file_voices_proto_enumTypes,
		MessageInfos:      file_voices_proto_msgTypes,
	}.Build()
	File_voices_proto = out.File
	file_voices_proto_rawDesc = nil
	file_voices_proto_goTypes = nil
	file_voices_proto_depIdxs = nil
}