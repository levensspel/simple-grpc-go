// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.28.2
// source: proto/user/user.proto

package user

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

type EmailStatus int32

const (
	EmailStatus_EMAIL_UNVERIFIED EmailStatus = 0 // The first enum value must be zero for open enums.
	EmailStatus_EMAIL_VERIFIED   EmailStatus = 1
	EmailStatus_EMAIL_BLOCKED    EmailStatus = 2
)

// Enum value maps for EmailStatus.
var (
	EmailStatus_name = map[int32]string{
		0: "EMAIL_UNVERIFIED",
		1: "EMAIL_VERIFIED",
		2: "EMAIL_BLOCKED",
	}
	EmailStatus_value = map[string]int32{
		"EMAIL_UNVERIFIED": 0,
		"EMAIL_VERIFIED":   1,
		"EMAIL_BLOCKED":    2,
	}
)

func (x EmailStatus) Enum() *EmailStatus {
	p := new(EmailStatus)
	*p = x
	return p
}

func (x EmailStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EmailStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_user_user_proto_enumTypes[0].Descriptor()
}

func (EmailStatus) Type() protoreflect.EnumType {
	return &file_proto_user_user_proto_enumTypes[0]
}

func (x EmailStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EmailStatus.Descriptor instead.
func (EmailStatus) EnumDescriptor() ([]byte, []int) {
	return file_proto_user_user_proto_rawDescGZIP(), []int{0}
}

type MembershipPerks int32

const (
	MembershipPerks_PERKS_FREE_ADS             MembershipPerks = 0
	MembershipPerks_PERKS_HIGHER_VIDEO_QUALITY MembershipPerks = 1
	MembershipPerks_PERKS_VIDEO_DOWNLOAD       MembershipPerks = 2
	MembershipPerks_PERKS_BACKGROUND_PLAY      MembershipPerks = 3
)

// Enum value maps for MembershipPerks.
var (
	MembershipPerks_name = map[int32]string{
		0: "PERKS_FREE_ADS",
		1: "PERKS_HIGHER_VIDEO_QUALITY",
		2: "PERKS_VIDEO_DOWNLOAD",
		3: "PERKS_BACKGROUND_PLAY",
	}
	MembershipPerks_value = map[string]int32{
		"PERKS_FREE_ADS":             0,
		"PERKS_HIGHER_VIDEO_QUALITY": 1,
		"PERKS_VIDEO_DOWNLOAD":       2,
		"PERKS_BACKGROUND_PLAY":      3,
	}
)

func (x MembershipPerks) Enum() *MembershipPerks {
	p := new(MembershipPerks)
	*p = x
	return p
}

func (x MembershipPerks) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MembershipPerks) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_user_user_proto_enumTypes[1].Descriptor()
}

func (MembershipPerks) Type() protoreflect.EnumType {
	return &file_proto_user_user_proto_enumTypes[1]
}

func (x MembershipPerks) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MembershipPerks.Descriptor instead.
func (MembershipPerks) EnumDescriptor() ([]byte, []int) {
	return file_proto_user_user_proto_rawDescGZIP(), []int{1}
}

type RolePermission int32

const (
	RolePermission_PERMISSION_UPLOAD_HIGHER_VIDEO_QUALITY RolePermission = 0
	RolePermission_PERMISSION_MONETIZE_VIDEO              RolePermission = 1
	RolePermission_PERMISSION_REDEEM_MONEY                RolePermission = 2
	RolePermission_PERMISSION_OPEN_MEMBERSHIP             RolePermission = 3
)

// Enum value maps for RolePermission.
var (
	RolePermission_name = map[int32]string{
		0: "PERMISSION_UPLOAD_HIGHER_VIDEO_QUALITY",
		1: "PERMISSION_MONETIZE_VIDEO",
		2: "PERMISSION_REDEEM_MONEY",
		3: "PERMISSION_OPEN_MEMBERSHIP",
	}
	RolePermission_value = map[string]int32{
		"PERMISSION_UPLOAD_HIGHER_VIDEO_QUALITY": 0,
		"PERMISSION_MONETIZE_VIDEO":              1,
		"PERMISSION_REDEEM_MONEY":                2,
		"PERMISSION_OPEN_MEMBERSHIP":             3,
	}
)

func (x RolePermission) Enum() *RolePermission {
	p := new(RolePermission)
	*p = x
	return p
}

func (x RolePermission) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RolePermission) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_user_user_proto_enumTypes[2].Descriptor()
}

func (RolePermission) Type() protoreflect.EnumType {
	return &file_proto_user_user_proto_enumTypes[2]
}

func (x RolePermission) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RolePermission.Descriptor instead.
func (RolePermission) EnumDescriptor() ([]byte, []int) {
	return file_proto_user_user_proto_rawDescGZIP(), []int{2}
}

type RegisterUserRequest struct {
	state    protoimpl.MessageState `protogen:"open.v1"`
	Name     string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Username string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email    string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password string                 `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Settings map[string]string      `protobuf:"bytes,5,rep,name=settings,proto3" json:"settings,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Types that are valid to be assigned to Role:
	//
	//	*RegisterUserRequest_UserRole
	//	*RegisterUserRequest_MembershipRole
	//	*RegisterUserRequest_CreatorRole
	Role          isRegisterUserRequest_Role `protobuf_oneof:"role"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterUserRequest) Reset() {
	*x = RegisterUserRequest{}
	mi := &file_proto_user_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterUserRequest) ProtoMessage() {}

func (x *RegisterUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterUserRequest.ProtoReflect.Descriptor instead.
func (*RegisterUserRequest) Descriptor() ([]byte, []int) {
	return file_proto_user_user_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterUserRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RegisterUserRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RegisterUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegisterUserRequest) GetSettings() map[string]string {
	if x != nil {
		return x.Settings
	}
	return nil
}

func (x *RegisterUserRequest) GetRole() isRegisterUserRequest_Role {
	if x != nil {
		return x.Role
	}
	return nil
}

func (x *RegisterUserRequest) GetUserRole() *UserRole {
	if x != nil {
		if x, ok := x.Role.(*RegisterUserRequest_UserRole); ok {
			return x.UserRole
		}
	}
	return nil
}

func (x *RegisterUserRequest) GetMembershipRole() *MembershipRole {
	if x != nil {
		if x, ok := x.Role.(*RegisterUserRequest_MembershipRole); ok {
			return x.MembershipRole
		}
	}
	return nil
}

func (x *RegisterUserRequest) GetCreatorRole() *CreatorRole {
	if x != nil {
		if x, ok := x.Role.(*RegisterUserRequest_CreatorRole); ok {
			return x.CreatorRole
		}
	}
	return nil
}

type isRegisterUserRequest_Role interface {
	isRegisterUserRequest_Role()
}

type RegisterUserRequest_UserRole struct {
	UserRole *UserRole `protobuf:"bytes,6,opt,name=user_role,json=userRole,proto3,oneof"`
}

type RegisterUserRequest_MembershipRole struct {
	MembershipRole *MembershipRole `protobuf:"bytes,7,opt,name=membership_role,json=membershipRole,proto3,oneof"`
}

type RegisterUserRequest_CreatorRole struct {
	CreatorRole *CreatorRole `protobuf:"bytes,8,opt,name=creator_role,json=creatorRole,proto3,oneof"`
}

func (*RegisterUserRequest_UserRole) isRegisterUserRequest_Role() {}

func (*RegisterUserRequest_MembershipRole) isRegisterUserRequest_Role() {}

func (*RegisterUserRequest_CreatorRole) isRegisterUserRequest_Role() {}

type RegisterUserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        uint32                 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	EmailStatus   EmailStatus            `protobuf:"varint,2,opt,name=email_status,json=emailStatus,proto3,enum=user.EmailStatus" json:"email_status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterUserResponse) Reset() {
	*x = RegisterUserResponse{}
	mi := &file_proto_user_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterUserResponse) ProtoMessage() {}

func (x *RegisterUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterUserResponse.ProtoReflect.Descriptor instead.
func (*RegisterUserResponse) Descriptor() ([]byte, []int) {
	return file_proto_user_user_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterUserResponse) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *RegisterUserResponse) GetEmailStatus() EmailStatus {
	if x != nil {
		return x.EmailStatus
	}
	return EmailStatus_EMAIL_UNVERIFIED
}

type UserRole struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccountLevel  uint32                 `protobuf:"varint,1,opt,name=account_level,json=accountLevel,proto3" json:"account_level,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserRole) Reset() {
	*x = UserRole{}
	mi := &file_proto_user_user_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserRole) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRole) ProtoMessage() {}

func (x *UserRole) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_user_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRole.ProtoReflect.Descriptor instead.
func (*UserRole) Descriptor() ([]byte, []int) {
	return file_proto_user_user_proto_rawDescGZIP(), []int{2}
}

func (x *UserRole) GetAccountLevel() uint32 {
	if x != nil {
		return x.AccountLevel
	}
	return 0
}

type MembershipRole struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccountLevel  uint32                 `protobuf:"varint,1,opt,name=account_level,json=accountLevel,proto3" json:"account_level,omitempty"`
	Perks         []MembershipPerks      `protobuf:"varint,2,rep,packed,name=perks,proto3,enum=user.MembershipPerks" json:"perks,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MembershipRole) Reset() {
	*x = MembershipRole{}
	mi := &file_proto_user_user_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MembershipRole) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MembershipRole) ProtoMessage() {}

func (x *MembershipRole) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_user_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MembershipRole.ProtoReflect.Descriptor instead.
func (*MembershipRole) Descriptor() ([]byte, []int) {
	return file_proto_user_user_proto_rawDescGZIP(), []int{3}
}

func (x *MembershipRole) GetAccountLevel() uint32 {
	if x != nil {
		return x.AccountLevel
	}
	return 0
}

func (x *MembershipRole) GetPerks() []MembershipPerks {
	if x != nil {
		return x.Perks
	}
	return nil
}

type CreatorRole struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Permissions   []RolePermission       `protobuf:"varint,1,rep,packed,name=permissions,proto3,enum=user.RolePermission" json:"permissions,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreatorRole) Reset() {
	*x = CreatorRole{}
	mi := &file_proto_user_user_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreatorRole) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatorRole) ProtoMessage() {}

func (x *CreatorRole) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_user_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatorRole.ProtoReflect.Descriptor instead.
func (*CreatorRole) Descriptor() ([]byte, []int) {
	return file_proto_user_user_proto_rawDescGZIP(), []int{4}
}

func (x *CreatorRole) GetPermissions() []RolePermission {
	if x != nil {
		return x.Permissions
	}
	return nil
}

var File_proto_user_user_proto protoreflect.FileDescriptor

var file_proto_user_user_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0xa9, 0x03,
	0x0a, 0x13, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x43, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x2d, 0x0a, 0x09,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x48,
	0x00, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x3f, 0x0a, 0x0f, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x52, 0x6f, 0x6c, 0x65, 0x48, 0x00, 0x52, 0x0e, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x36, 0x0a, 0x0c,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x6f,
	0x72, 0x52, 0x6f, 0x6c, 0x65, 0x48, 0x00, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72,
	0x52, 0x6f, 0x6c, 0x65, 0x1a, 0x3b, 0x0a, 0x0d, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x42, 0x06, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x65, 0x0a, 0x14, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x0c, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x0b, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x2f, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x22, 0x62, 0x0a, 0x0e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x52,
	0x6f, 0x6c, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x2b, 0x0a, 0x05, 0x70, 0x65, 0x72, 0x6b,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x50, 0x65, 0x72, 0x6b, 0x73, 0x52, 0x05,
	0x70, 0x65, 0x72, 0x6b, 0x73, 0x22, 0x45, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72,
	0x52, 0x6f, 0x6c, 0x65, 0x12, 0x36, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2a, 0x4a, 0x0a, 0x0b,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x10, 0x45,
	0x4d, 0x41, 0x49, 0x4c, 0x5f, 0x55, 0x4e, 0x56, 0x45, 0x52, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x12, 0x0a, 0x0e, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x5f, 0x56, 0x45, 0x52, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x5f, 0x42,
	0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10, 0x02, 0x2a, 0x7a, 0x0a, 0x0f, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x50, 0x65, 0x72, 0x6b, 0x73, 0x12, 0x12, 0x0a, 0x0e, 0x50,
	0x45, 0x52, 0x4b, 0x53, 0x5f, 0x46, 0x52, 0x45, 0x45, 0x5f, 0x41, 0x44, 0x53, 0x10, 0x00, 0x12,
	0x1e, 0x0a, 0x1a, 0x50, 0x45, 0x52, 0x4b, 0x53, 0x5f, 0x48, 0x49, 0x47, 0x48, 0x45, 0x52, 0x5f,
	0x56, 0x49, 0x44, 0x45, 0x4f, 0x5f, 0x51, 0x55, 0x41, 0x4c, 0x49, 0x54, 0x59, 0x10, 0x01, 0x12,
	0x18, 0x0a, 0x14, 0x50, 0x45, 0x52, 0x4b, 0x53, 0x5f, 0x56, 0x49, 0x44, 0x45, 0x4f, 0x5f, 0x44,
	0x4f, 0x57, 0x4e, 0x4c, 0x4f, 0x41, 0x44, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x50, 0x45, 0x52,
	0x4b, 0x53, 0x5f, 0x42, 0x41, 0x43, 0x4b, 0x47, 0x52, 0x4f, 0x55, 0x4e, 0x44, 0x5f, 0x50, 0x4c,
	0x41, 0x59, 0x10, 0x03, 0x2a, 0x98, 0x01, 0x0a, 0x0e, 0x52, 0x6f, 0x6c, 0x65, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x26, 0x50, 0x45, 0x52, 0x4d, 0x49,
	0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x50, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x48, 0x49, 0x47,
	0x48, 0x45, 0x52, 0x5f, 0x56, 0x49, 0x44, 0x45, 0x4f, 0x5f, 0x51, 0x55, 0x41, 0x4c, 0x49, 0x54,
	0x59, 0x10, 0x00, 0x12, 0x1d, 0x0a, 0x19, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f,
	0x4e, 0x5f, 0x4d, 0x4f, 0x4e, 0x45, 0x54, 0x49, 0x5a, 0x45, 0x5f, 0x56, 0x49, 0x44, 0x45, 0x4f,
	0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e,
	0x5f, 0x52, 0x45, 0x44, 0x45, 0x45, 0x4d, 0x5f, 0x4d, 0x4f, 0x4e, 0x45, 0x59, 0x10, 0x02, 0x12,
	0x1e, 0x0a, 0x1a, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x4f, 0x50,
	0x45, 0x4e, 0x5f, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x53, 0x48, 0x49, 0x50, 0x10, 0x03, 0x32,
	0x56, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47,
	0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x12, 0x19,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x65, 0x76, 0x65, 0x6e, 0x73, 0x73, 0x70, 0x65, 0x6c,
	0x2f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x6f, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_user_user_proto_rawDescOnce sync.Once
	file_proto_user_user_proto_rawDescData = file_proto_user_user_proto_rawDesc
)

func file_proto_user_user_proto_rawDescGZIP() []byte {
	file_proto_user_user_proto_rawDescOnce.Do(func() {
		file_proto_user_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_user_user_proto_rawDescData)
	})
	return file_proto_user_user_proto_rawDescData
}

var file_proto_user_user_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_proto_user_user_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_user_user_proto_goTypes = []any{
	(EmailStatus)(0),             // 0: user.EmailStatus
	(MembershipPerks)(0),         // 1: user.MembershipPerks
	(RolePermission)(0),          // 2: user.RolePermission
	(*RegisterUserRequest)(nil),  // 3: user.RegisterUserRequest
	(*RegisterUserResponse)(nil), // 4: user.RegisterUserResponse
	(*UserRole)(nil),             // 5: user.UserRole
	(*MembershipRole)(nil),       // 6: user.MembershipRole
	(*CreatorRole)(nil),          // 7: user.CreatorRole
	nil,                          // 8: user.RegisterUserRequest.SettingsEntry
}
var file_proto_user_user_proto_depIdxs = []int32{
	8, // 0: user.RegisterUserRequest.settings:type_name -> user.RegisterUserRequest.SettingsEntry
	5, // 1: user.RegisterUserRequest.user_role:type_name -> user.UserRole
	6, // 2: user.RegisterUserRequest.membership_role:type_name -> user.MembershipRole
	7, // 3: user.RegisterUserRequest.creator_role:type_name -> user.CreatorRole
	0, // 4: user.RegisterUserResponse.email_status:type_name -> user.EmailStatus
	1, // 5: user.MembershipRole.perks:type_name -> user.MembershipPerks
	2, // 6: user.CreatorRole.permissions:type_name -> user.RolePermission
	3, // 7: user.UserService.RegisterUser:input_type -> user.RegisterUserRequest
	4, // 8: user.UserService.RegisterUser:output_type -> user.RegisterUserResponse
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_proto_user_user_proto_init() }
func file_proto_user_user_proto_init() {
	if File_proto_user_user_proto != nil {
		return
	}
	file_proto_user_user_proto_msgTypes[0].OneofWrappers = []any{
		(*RegisterUserRequest_UserRole)(nil),
		(*RegisterUserRequest_MembershipRole)(nil),
		(*RegisterUserRequest_CreatorRole)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_user_user_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_user_user_proto_goTypes,
		DependencyIndexes: file_proto_user_user_proto_depIdxs,
		EnumInfos:         file_proto_user_user_proto_enumTypes,
		MessageInfos:      file_proto_user_user_proto_msgTypes,
	}.Build()
	File_proto_user_user_proto = out.File
	file_proto_user_user_proto_rawDesc = nil
	file_proto_user_user_proto_goTypes = nil
	file_proto_user_user_proto_depIdxs = nil
}
