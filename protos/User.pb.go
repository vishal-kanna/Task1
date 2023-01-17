// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: protos/User.proto

package protos

import (
	_ "github.com/golang/protobuf/ptypes/timestamp"
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

type Activity int32

const (
	Activity_play  Activity = 0
	Activity_eat   Activity = 1
	Activity_sleep Activity = 2
	Activity_read  Activity = 3
)

// Enum value maps for Activity.
var (
	Activity_name = map[int32]string{
		0: "play",
		1: "eat",
		2: "sleep",
		3: "read",
	}
	Activity_value = map[string]int32{
		"play":  0,
		"eat":   1,
		"sleep": 2,
		"read":  3,
	}
)

func (x Activity) Enum() *Activity {
	p := new(Activity)
	*p = x
	return p
}

func (x Activity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Activity) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_User_proto_enumTypes[0].Descriptor()
}

func (Activity) Type() protoreflect.EnumType {
	return &file_protos_User_proto_enumTypes[0]
}

func (x Activity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Activity.Descriptor instead.
func (Activity) EnumDescriptor() ([]byte, []int) {
	return file_protos_User_proto_rawDescGZIP(), []int{0}
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_User_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_protos_User_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_protos_User_proto_rawDescGZIP(), []int{0}
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Phone int64  `protobuf:"varint,3,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_User_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_protos_User_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_protos_User_proto_rawDescGZIP(), []int{1}
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetPhone() int64 {
	if x != nil {
		return x.Phone
	}
	return 0
}

type AddActivityReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email     string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Activity  Activity `protobuf:"varint,2,opt,name=activity,proto3,enum=Activity" json:"activity,omitempty"`
	AddedTime string   `protobuf:"bytes,3,opt,name=added_time,json=addedTime,proto3" json:"added_time,omitempty"`
	Duration  int64    `protobuf:"varint,4,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (x *AddActivityReq) Reset() {
	*x = AddActivityReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_User_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddActivityReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddActivityReq) ProtoMessage() {}

func (x *AddActivityReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_User_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddActivityReq.ProtoReflect.Descriptor instead.
func (*AddActivityReq) Descriptor() ([]byte, []int) {
	return file_protos_User_proto_rawDescGZIP(), []int{2}
}

func (x *AddActivityReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AddActivityReq) GetActivity() Activity {
	if x != nil {
		return x.Activity
	}
	return Activity_play
}

func (x *AddActivityReq) GetAddedTime() string {
	if x != nil {
		return x.AddedTime
	}
	return ""
}

func (x *AddActivityReq) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

type AddActivityRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *AddActivityRes) Reset() {
	*x = AddActivityRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_User_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddActivityRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddActivityRes) ProtoMessage() {}

func (x *AddActivityRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_User_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddActivityRes.ProtoReflect.Descriptor instead.
func (*AddActivityRes) Descriptor() ([]byte, []int) {
	return file_protos_User_proto_rawDescGZIP(), []int{3}
}

func (x *AddActivityRes) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type AddUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *AddUserRequest) Reset() {
	*x = AddUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_User_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserRequest) ProtoMessage() {}

func (x *AddUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_User_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserRequest.ProtoReflect.Descriptor instead.
func (*AddUserRequest) Descriptor() ([]byte, []int) {
	return file_protos_User_proto_rawDescGZIP(), []int{4}
}

func (x *AddUserRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type AddUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *AddUserResponse) Reset() {
	*x = AddUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_User_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserResponse) ProtoMessage() {}

func (x *AddUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_User_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserResponse.ProtoReflect.Descriptor instead.
func (*AddUserResponse) Descriptor() ([]byte, []int) {
	return file_protos_User_proto_rawDescGZIP(), []int{5}
}

func (x *AddUserResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type UpdateActivityReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email     string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Activity  string `protobuf:"bytes,2,opt,name=Activity,proto3" json:"Activity,omitempty"`
	Duration  int64  `protobuf:"varint,3,opt,name=duration,proto3" json:"duration,omitempty"`
	AddedTime string `protobuf:"bytes,4,opt,name=added_time,json=addedTime,proto3" json:"added_time,omitempty"`
}

func (x *UpdateActivityReq) Reset() {
	*x = UpdateActivityReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_User_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateActivityReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateActivityReq) ProtoMessage() {}

func (x *UpdateActivityReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_User_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateActivityReq.ProtoReflect.Descriptor instead.
func (*UpdateActivityReq) Descriptor() ([]byte, []int) {
	return file_protos_User_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateActivityReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UpdateActivityReq) GetActivity() string {
	if x != nil {
		return x.Activity
	}
	return ""
}

func (x *UpdateActivityReq) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *UpdateActivityReq) GetAddedTime() string {
	if x != nil {
		return x.AddedTime
	}
	return ""
}

type UpdateActivityRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=Response,proto3" json:"Response,omitempty"`
}

func (x *UpdateActivityRes) Reset() {
	*x = UpdateActivityRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_User_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateActivityRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateActivityRes) ProtoMessage() {}

func (x *UpdateActivityRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_User_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateActivityRes.ProtoReflect.Descriptor instead.
func (*UpdateActivityRes) Descriptor() ([]byte, []int) {
	return file_protos_User_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateActivityRes) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type FindUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *FindUserReq) Reset() {
	*x = FindUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_User_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindUserReq) ProtoMessage() {}

func (x *FindUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_User_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindUserReq.ProtoReflect.Descriptor instead.
func (*FindUserReq) Descriptor() ([]byte, []int) {
	return file_protos_User_proto_rawDescGZIP(), []int{8}
}

func (x *FindUserReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type FindUserRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User       *User      `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Activities []Activity `protobuf:"varint,2,rep,packed,name=activities,proto3,enum=Activity" json:"activities,omitempty"`
}

func (x *FindUserRes) Reset() {
	*x = FindUserRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_User_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindUserRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindUserRes) ProtoMessage() {}

func (x *FindUserRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_User_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindUserRes.ProtoReflect.Descriptor instead.
func (*FindUserRes) Descriptor() ([]byte, []int) {
	return file_protos_User_proto_rawDescGZIP(), []int{9}
}

func (x *FindUserRes) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *FindUserRes) GetActivities() []Activity {
	if x != nil {
		return x.Activities
	}
	return nil
}

type Adding struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Activity  Activity `protobuf:"varint,1,opt,name=activity,proto3,enum=Activity" json:"activity,omitempty"`
	AddedTime string   `protobuf:"bytes,2,opt,name=added_time,json=addedTime,proto3" json:"added_time,omitempty"`
	Duration  int64    `protobuf:"varint,3,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (x *Adding) Reset() {
	*x = Adding{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_User_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Adding) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Adding) ProtoMessage() {}

func (x *Adding) ProtoReflect() protoreflect.Message {
	mi := &file_protos_User_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Adding.ProtoReflect.Descriptor instead.
func (*Adding) Descriptor() ([]byte, []int) {
	return file_protos_User_proto_rawDescGZIP(), []int{10}
}

func (x *Adding) GetActivity() Activity {
	if x != nil {
		return x.Activity
	}
	return Activity_play
}

func (x *Adding) GetAddedTime() string {
	if x != nil {
		return x.AddedTime
	}
	return ""
}

func (x *Adding) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

var File_protos_User_proto protoreflect.FileDescriptor

var file_protos_User_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x46, 0x0a,
	0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x88, 0x01, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x25,
	0x0a, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x09, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x08, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x64, 0x64, 0x65, 0x64, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x64, 0x64, 0x65, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x2c, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52,
	0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2b,
	0x0a, 0x0e, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x19, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x2d, 0x0a, 0x0f, 0x41,
	0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x80, 0x01, 0x0a, 0x11, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d,
	0x0a, 0x0a, 0x61, 0x64, 0x64, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x61, 0x64, 0x64, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x2f, 0x0a,
	0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52,
	0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23,
	0x0a, 0x0b, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x22, 0x53, 0x0a, 0x0b, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x12, 0x19, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x05, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x29, 0x0a,
	0x0a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0e, 0x32, 0x09, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x0a, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x22, 0x6a, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x69,
	0x6e, 0x67, 0x12, 0x25, 0x0a, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52,
	0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x64, 0x64,
	0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61,
	0x64, 0x64, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x32, 0x0a, 0x08, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x12, 0x08, 0x0a, 0x04, 0x70, 0x6c, 0x61, 0x79, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x65, 0x61,
	0x74, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x73, 0x6c, 0x65, 0x65, 0x70, 0x10, 0x02, 0x12, 0x08,
	0x0a, 0x04, 0x72, 0x65, 0x61, 0x64, 0x10, 0x03, 0x32, 0xc7, 0x01, 0x0a, 0x07, 0x54, 0x72, 0x61,
	0x63, 0x6b, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x0f, 0x2e, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x10, 0x2e, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2f, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x12, 0x0f, 0x2e, 0x41, 0x64, 0x64, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52,
	0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x41, 0x64, 0x64, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x52, 0x65, 0x73, 0x12, 0x39, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x65, 0x73, 0x12, 0x12, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x12, 0x22,
	0x0a, 0x04, 0x46, 0x69, 0x6e, 0x64, 0x12, 0x0c, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x1a, 0x0c, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x42, 0x0d, 0x5a, 0x0b, 0x54, 0x61, 0x73, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_User_proto_rawDescOnce sync.Once
	file_protos_User_proto_rawDescData = file_protos_User_proto_rawDesc
)

func file_protos_User_proto_rawDescGZIP() []byte {
	file_protos_User_proto_rawDescOnce.Do(func() {
		file_protos_User_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_User_proto_rawDescData)
	})
	return file_protos_User_proto_rawDescData
}

var file_protos_User_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protos_User_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_protos_User_proto_goTypes = []interface{}{
	(Activity)(0),             // 0: Activity
	(*Empty)(nil),             // 1: Empty
	(*User)(nil),              // 2: User
	(*AddActivityReq)(nil),    // 3: AddActivityReq
	(*AddActivityRes)(nil),    // 4: AddActivityRes
	(*AddUserRequest)(nil),    // 5: AddUserRequest
	(*AddUserResponse)(nil),   // 6: AddUserResponse
	(*UpdateActivityReq)(nil), // 7: UpdateActivityReq
	(*UpdateActivityRes)(nil), // 8: UpdateActivityRes
	(*FindUserReq)(nil),       // 9: FindUserReq
	(*FindUserRes)(nil),       // 10: FindUserRes
	(*Adding)(nil),            // 11: Adding
}
var file_protos_User_proto_depIdxs = []int32{
	0,  // 0: AddActivityReq.activity:type_name -> Activity
	2,  // 1: AddUserRequest.user:type_name -> User
	2,  // 2: FindUserRes.user:type_name -> User
	0,  // 3: FindUserRes.activities:type_name -> Activity
	0,  // 4: Adding.activity:type_name -> Activity
	5,  // 5: Tracker.AddUser:input_type -> AddUserRequest
	3,  // 6: Tracker.AddActivity:input_type -> AddActivityReq
	7,  // 7: Tracker.UpdateActivites:input_type -> UpdateActivityReq
	9,  // 8: Tracker.Find:input_type -> FindUserReq
	6,  // 9: Tracker.AddUser:output_type -> AddUserResponse
	4,  // 10: Tracker.AddActivity:output_type -> AddActivityRes
	8,  // 11: Tracker.UpdateActivites:output_type -> UpdateActivityRes
	10, // 12: Tracker.Find:output_type -> FindUserRes
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_protos_User_proto_init() }
func file_protos_User_proto_init() {
	if File_protos_User_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_User_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_protos_User_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_protos_User_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddActivityReq); i {
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
		file_protos_User_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddActivityRes); i {
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
		file_protos_User_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUserRequest); i {
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
		file_protos_User_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUserResponse); i {
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
		file_protos_User_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateActivityReq); i {
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
		file_protos_User_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateActivityRes); i {
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
		file_protos_User_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindUserReq); i {
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
		file_protos_User_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindUserRes); i {
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
		file_protos_User_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Adding); i {
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
			RawDescriptor: file_protos_User_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_User_proto_goTypes,
		DependencyIndexes: file_protos_User_proto_depIdxs,
		EnumInfos:         file_protos_User_proto_enumTypes,
		MessageInfos:      file_protos_User_proto_msgTypes,
	}.Build()
	File_protos_User_proto = out.File
	file_protos_User_proto_rawDesc = nil
	file_protos_User_proto_goTypes = nil
	file_protos_User_proto_depIdxs = nil
}
