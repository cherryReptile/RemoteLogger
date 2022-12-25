// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.10
// source: proto/auth.proto

package api

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

type AppRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=Email,proto3" json:"Email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *AppRequest) Reset() {
	*x = AppRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppRequest) ProtoMessage() {}

func (x *AppRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppRequest.ProtoReflect.Descriptor instead.
func (*AppRequest) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{0}
}

func (x *AppRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AppRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type AddAppRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request  *AppRequest `protobuf:"bytes,1,opt,name=Request,proto3" json:"Request,omitempty"`
	UserUUID string      `protobuf:"bytes,2,opt,name=UserUUID,proto3" json:"UserUUID,omitempty"`
}

func (x *AddAppRequest) Reset() {
	*x = AddAppRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddAppRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddAppRequest) ProtoMessage() {}

func (x *AddAppRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddAppRequest.ProtoReflect.Descriptor instead.
func (*AddAppRequest) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{1}
}

func (x *AddAppRequest) GetRequest() *AppRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *AddAppRequest) GetUserUUID() string {
	if x != nil {
		return x.UserUUID
	}
	return ""
}

type GitHubRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login string `protobuf:"bytes,1,opt,name=Login,proto3" json:"Login,omitempty"`
	Data  []byte `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *GitHubRequest) Reset() {
	*x = GitHubRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GitHubRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GitHubRequest) ProtoMessage() {}

func (x *GitHubRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GitHubRequest.ProtoReflect.Descriptor instead.
func (*GitHubRequest) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{2}
}

func (x *GitHubRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *GitHubRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type AddGitHubRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request  *GitHubRequest `protobuf:"bytes,1,opt,name=Request,proto3" json:"Request,omitempty"`
	UserUUID string         `protobuf:"bytes,2,opt,name=UserUUID,proto3" json:"UserUUID,omitempty"`
}

func (x *AddGitHubRequest) Reset() {
	*x = AddGitHubRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddGitHubRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddGitHubRequest) ProtoMessage() {}

func (x *AddGitHubRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddGitHubRequest.ProtoReflect.Descriptor instead.
func (*AddGitHubRequest) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{3}
}

func (x *AddGitHubRequest) GetRequest() *GitHubRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *AddGitHubRequest) GetUserUUID() string {
	if x != nil {
		return x.UserUUID
	}
	return ""
}

type GoogleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=Email,proto3" json:"Email,omitempty"`
	Data  []byte `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *GoogleRequest) Reset() {
	*x = GoogleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoogleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoogleRequest) ProtoMessage() {}

func (x *GoogleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoogleRequest.ProtoReflect.Descriptor instead.
func (*GoogleRequest) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{4}
}

func (x *GoogleRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GoogleRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type AddGoogleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request  *GoogleRequest `protobuf:"bytes,1,opt,name=Request,proto3" json:"Request,omitempty"`
	UserUUID string         `protobuf:"bytes,2,opt,name=UserUUID,proto3" json:"UserUUID,omitempty"`
}

func (x *AddGoogleRequest) Reset() {
	*x = AddGoogleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddGoogleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddGoogleRequest) ProtoMessage() {}

func (x *AddGoogleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddGoogleRequest.ProtoReflect.Descriptor instead.
func (*AddGoogleRequest) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{5}
}

func (x *AddGoogleRequest) GetRequest() *GoogleRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *AddGoogleRequest) GetUserUUID() string {
	if x != nil {
		return x.UserUUID
	}
	return ""
}

type TelegramRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
}

func (x *TelegramRequest) Reset() {
	*x = TelegramRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TelegramRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TelegramRequest) ProtoMessage() {}

func (x *TelegramRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TelegramRequest.ProtoReflect.Descriptor instead.
func (*TelegramRequest) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{6}
}

func (x *TelegramRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type AppResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Struct   *User  `protobuf:"bytes,1,opt,name=Struct,proto3" json:"Struct,omitempty"`
	TokenStr string `protobuf:"bytes,2,opt,name=TokenStr,proto3" json:"TokenStr,omitempty"`
}

func (x *AppResponse) Reset() {
	*x = AppResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppResponse) ProtoMessage() {}

func (x *AppResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppResponse.ProtoReflect.Descriptor instead.
func (*AppResponse) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{7}
}

func (x *AppResponse) GetStruct() *User {
	if x != nil {
		return x.Struct
	}
	return nil
}

func (x *AppResponse) GetTokenStr() string {
	if x != nil {
		return x.TokenStr
	}
	return ""
}

type AddedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Struct  *User  `protobuf:"bytes,1,opt,name=Struct,proto3" json:"Struct,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *AddedResponse) Reset() {
	*x = AddedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddedResponse) ProtoMessage() {}

func (x *AddedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddedResponse.ProtoReflect.Descriptor instead.
func (*AddedResponse) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{8}
}

func (x *AddedResponse) GetStruct() *User {
	if x != nil {
		return x.Struct
	}
	return nil
}

func (x *AddedResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type TokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (x *TokenRequest) Reset() {
	*x = TokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenRequest) ProtoMessage() {}

func (x *TokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenRequest.ProtoReflect.Descriptor instead.
func (*TokenRequest) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{9}
}

func (x *TokenRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type LogoutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *LogoutResponse) Reset() {
	*x = LogoutResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogoutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutResponse) ProtoMessage() {}

func (x *LogoutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutResponse.ProtoReflect.Descriptor instead.
func (*LogoutResponse) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{10}
}

func (x *LogoutResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CheckAuthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserUUID string `protobuf:"bytes,1,opt,name=UserUUID,proto3" json:"UserUUID,omitempty"`
}

func (x *CheckAuthResponse) Reset() {
	*x = CheckAuthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckAuthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckAuthResponse) ProtoMessage() {}

func (x *CheckAuthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckAuthResponse.ProtoReflect.Descriptor instead.
func (*CheckAuthResponse) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{11}
}

func (x *CheckAuthResponse) GetUserUUID() string {
	if x != nil {
		return x.UserUUID
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UUID      string `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
	Login     string `protobuf:"bytes,2,opt,name=Login,proto3" json:"Login,omitempty"`
	CreatedAt string `protobuf:"bytes,5,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[12]
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
	return file_proto_auth_proto_rawDescGZIP(), []int{12}
}

func (x *User) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

func (x *User) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *User) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

var File_proto_auth_proto protoreflect.FileDescriptor

var file_proto_auth_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x09, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x22, 0x3e, 0x0a,
	0x0a, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x5c, 0x0a,
	0x0d, 0x41, 0x64, 0x64, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f,
	0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x22, 0x39, 0x0a, 0x0d, 0x47,
	0x69, 0x74, 0x48, 0x75, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x62, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x47, 0x69, 0x74,
	0x48, 0x75, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x07, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6c, 0x6f,
	0x67, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x69, 0x74, 0x48, 0x75, 0x62, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x22, 0x39, 0x0a, 0x0d, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x62, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x07, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6c, 0x6f, 0x67,
	0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x55, 0x73, 0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x55, 0x73, 0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x22, 0x2d, 0x0a, 0x0f, 0x54, 0x65, 0x6c,
	0x65, 0x67, 0x72, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x52, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x06, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x74, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x74, 0x72, 0x22, 0x52, 0x0a, 0x0d,
	0x41, 0x64, 0x64, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a,
	0x06, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x06,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x24, 0x0a, 0x0c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x2a, 0x0a, 0x0e, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x2f, 0x0a, 0x11, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x75, 0x74, 0x68, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x55,
	0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x55,
	0x55, 0x49, 0x44, 0x22, 0x4e, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x55,
	0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x55, 0x49, 0x44, 0x12,
	0x14, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x32, 0xcb, 0x01, 0x0a, 0x0e, 0x41, 0x75, 0x74, 0x68, 0x41, 0x70, 0x70, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x12, 0x15, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x70, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6c, 0x6f, 0x67, 0x67,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x15, 0x2e, 0x6c,
	0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a,
	0x0a, 0x41, 0x64, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x2e, 0x6c, 0x6f,
	0x67, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x41, 0x70, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x64, 0x64, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x32, 0x97, 0x01, 0x0a, 0x11, 0x41, 0x75, 0x74, 0x68, 0x47, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x12, 0x18, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x69, 0x74,
	0x48, 0x75, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6c, 0x6f, 0x67,
	0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x1b, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x64, 0x64, 0x47, 0x69, 0x74, 0x48, 0x75, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x65,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0x97, 0x01, 0x0a, 0x11,
	0x41, 0x75, 0x74, 0x68, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3b, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x18, 0x2e, 0x6c, 0x6f, 0x67,
	0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45,
	0x0a, 0x0a, 0x41, 0x64, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x2e, 0x6c,
	0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6c, 0x6f, 0x67, 0x67,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0x54, 0x0a, 0x13, 0x41, 0x75, 0x74, 0x68, 0x54, 0x65, 0x6c,
	0x65, 0x67, 0x72, 0x61, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x05,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0x58, 0x0a, 0x10, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x44, 0x0a, 0x09, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x75, 0x74, 0x68, 0x12, 0x17, 0x2e, 0x6c,
	0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0x4f, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74,
	0x12, 0x17, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6c, 0x6f, 0x67, 0x67,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_auth_proto_rawDescOnce sync.Once
	file_proto_auth_proto_rawDescData = file_proto_auth_proto_rawDesc
)

func file_proto_auth_proto_rawDescGZIP() []byte {
	file_proto_auth_proto_rawDescOnce.Do(func() {
		file_proto_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_auth_proto_rawDescData)
	})
	return file_proto_auth_proto_rawDescData
}

var file_proto_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_proto_auth_proto_goTypes = []interface{}{
	(*AppRequest)(nil),        // 0: logger.v1.AppRequest
	(*AddAppRequest)(nil),     // 1: logger.v1.AddAppRequest
	(*GitHubRequest)(nil),     // 2: logger.v1.GitHubRequest
	(*AddGitHubRequest)(nil),  // 3: logger.v1.AddGitHubRequest
	(*GoogleRequest)(nil),     // 4: logger.v1.GoogleRequest
	(*AddGoogleRequest)(nil),  // 5: logger.v1.AddGoogleRequest
	(*TelegramRequest)(nil),   // 6: logger.v1.TelegramRequest
	(*AppResponse)(nil),       // 7: logger.v1.AppResponse
	(*AddedResponse)(nil),     // 8: logger.v1.AddedResponse
	(*TokenRequest)(nil),      // 9: logger.v1.TokenRequest
	(*LogoutResponse)(nil),    // 10: logger.v1.LogoutResponse
	(*CheckAuthResponse)(nil), // 11: logger.v1.CheckAuthResponse
	(*User)(nil),              // 12: logger.v1.User
}
var file_proto_auth_proto_depIdxs = []int32{
	0,  // 0: logger.v1.AddAppRequest.Request:type_name -> logger.v1.AppRequest
	2,  // 1: logger.v1.AddGitHubRequest.Request:type_name -> logger.v1.GitHubRequest
	4,  // 2: logger.v1.AddGoogleRequest.Request:type_name -> logger.v1.GoogleRequest
	12, // 3: logger.v1.AppResponse.Struct:type_name -> logger.v1.User
	12, // 4: logger.v1.AddedResponse.Struct:type_name -> logger.v1.User
	0,  // 5: logger.v1.AuthAppService.Register:input_type -> logger.v1.AppRequest
	0,  // 6: logger.v1.AuthAppService.Login:input_type -> logger.v1.AppRequest
	1,  // 7: logger.v1.AuthAppService.AddAccount:input_type -> logger.v1.AddAppRequest
	2,  // 8: logger.v1.AuthGithubService.Login:input_type -> logger.v1.GitHubRequest
	3,  // 9: logger.v1.AuthGithubService.AddAccount:input_type -> logger.v1.AddGitHubRequest
	4,  // 10: logger.v1.AuthGoogleService.Login:input_type -> logger.v1.GoogleRequest
	5,  // 11: logger.v1.AuthGoogleService.AddAccount:input_type -> logger.v1.AddGoogleRequest
	6,  // 12: logger.v1.AuthTelegramService.Login:input_type -> logger.v1.TelegramRequest
	9,  // 13: logger.v1.CheckAuthService.CheckAuth:input_type -> logger.v1.TokenRequest
	9,  // 14: logger.v1.LogoutService.Logout:input_type -> logger.v1.TokenRequest
	7,  // 15: logger.v1.AuthAppService.Register:output_type -> logger.v1.AppResponse
	7,  // 16: logger.v1.AuthAppService.Login:output_type -> logger.v1.AppResponse
	8,  // 17: logger.v1.AuthAppService.AddAccount:output_type -> logger.v1.AddedResponse
	7,  // 18: logger.v1.AuthGithubService.Login:output_type -> logger.v1.AppResponse
	8,  // 19: logger.v1.AuthGithubService.AddAccount:output_type -> logger.v1.AddedResponse
	7,  // 20: logger.v1.AuthGoogleService.Login:output_type -> logger.v1.AppResponse
	8,  // 21: logger.v1.AuthGoogleService.AddAccount:output_type -> logger.v1.AddedResponse
	7,  // 22: logger.v1.AuthTelegramService.Login:output_type -> logger.v1.AppResponse
	11, // 23: logger.v1.CheckAuthService.CheckAuth:output_type -> logger.v1.CheckAuthResponse
	10, // 24: logger.v1.LogoutService.Logout:output_type -> logger.v1.LogoutResponse
	15, // [15:25] is the sub-list for method output_type
	5,  // [5:15] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_proto_auth_proto_init() }
func file_proto_auth_proto_init() {
	if File_proto_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppRequest); i {
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
		file_proto_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddAppRequest); i {
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
		file_proto_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GitHubRequest); i {
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
		file_proto_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddGitHubRequest); i {
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
		file_proto_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoogleRequest); i {
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
		file_proto_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddGoogleRequest); i {
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
		file_proto_auth_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TelegramRequest); i {
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
		file_proto_auth_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppResponse); i {
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
		file_proto_auth_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddedResponse); i {
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
		file_proto_auth_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenRequest); i {
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
		file_proto_auth_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogoutResponse); i {
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
		file_proto_auth_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckAuthResponse); i {
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
		file_proto_auth_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   6,
		},
		GoTypes:           file_proto_auth_proto_goTypes,
		DependencyIndexes: file_proto_auth_proto_depIdxs,
		MessageInfos:      file_proto_auth_proto_msgTypes,
	}.Build()
	File_proto_auth_proto = out.File
	file_proto_auth_proto_rawDesc = nil
	file_proto_auth_proto_goTypes = nil
	file_proto_auth_proto_depIdxs = nil
}
