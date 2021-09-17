// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-rc.1
// 	protoc        v3.12.4
// source: deviceid.proto

package deviceid_io

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

type DeviceId_DeviceIdType int32

const (
	DeviceId_t_none           DeviceId_DeviceIdType = 0
	DeviceId_t_gaid           DeviceId_DeviceIdType = 1
	DeviceId_t_gaid_md5       DeviceId_DeviceIdType = 2
	DeviceId_t_gaid_sha1      DeviceId_DeviceIdType = 3
	DeviceId_t_idfa           DeviceId_DeviceIdType = 4
	DeviceId_t_idfa_md5       DeviceId_DeviceIdType = 5
	DeviceId_t_idfa_sha1      DeviceId_DeviceIdType = 6
	DeviceId_t_imei           DeviceId_DeviceIdType = 7
	DeviceId_t_imei_md5       DeviceId_DeviceIdType = 8
	DeviceId_t_imei_sha1      DeviceId_DeviceIdType = 9
	DeviceId_t_androidid      DeviceId_DeviceIdType = 10
	DeviceId_t_androidid_md5  DeviceId_DeviceIdType = 11
	DeviceId_t_androidid_sha1 DeviceId_DeviceIdType = 12
	DeviceId_t_mac_md5        DeviceId_DeviceIdType = 14
	DeviceId_t_mac_sha1       DeviceId_DeviceIdType = 15
	DeviceId_t_oaid           DeviceId_DeviceIdType = 16
	DeviceId_t_oaid_md5       DeviceId_DeviceIdType = 17
	DeviceId_t_idfv           DeviceId_DeviceIdType = 19
	DeviceId_t_ip             DeviceId_DeviceIdType = 22
	DeviceId_t_ua             DeviceId_DeviceIdType = 23
	DeviceId_t_ipua_md5       DeviceId_DeviceIdType = 24
	DeviceId_t_caid_gx        DeviceId_DeviceIdType = 31 // caid gx 广协版
	DeviceId_t_caid_ry        DeviceId_DeviceIdType = 32 // caid ry 热云版
)

// Enum value maps for DeviceId_DeviceIdType.
var (
	DeviceId_DeviceIdType_name = map[int32]string{
		0:  "t_none",
		1:  "t_gaid",
		2:  "t_gaid_md5",
		3:  "t_gaid_sha1",
		4:  "t_idfa",
		5:  "t_idfa_md5",
		6:  "t_idfa_sha1",
		7:  "t_imei",
		8:  "t_imei_md5",
		9:  "t_imei_sha1",
		10: "t_androidid",
		11: "t_androidid_md5",
		12: "t_androidid_sha1",
		14: "t_mac_md5",
		15: "t_mac_sha1",
		16: "t_oaid",
		17: "t_oaid_md5",
		19: "t_idfv",
		22: "t_ip",
		23: "t_ua",
		24: "t_ipua_md5",
		31: "t_caid_gx",
		32: "t_caid_ry",
	}
	DeviceId_DeviceIdType_value = map[string]int32{
		"t_none":           0,
		"t_gaid":           1,
		"t_gaid_md5":       2,
		"t_gaid_sha1":      3,
		"t_idfa":           4,
		"t_idfa_md5":       5,
		"t_idfa_sha1":      6,
		"t_imei":           7,
		"t_imei_md5":       8,
		"t_imei_sha1":      9,
		"t_androidid":      10,
		"t_androidid_md5":  11,
		"t_androidid_sha1": 12,
		"t_mac_md5":        14,
		"t_mac_sha1":       15,
		"t_oaid":           16,
		"t_oaid_md5":       17,
		"t_idfv":           19,
		"t_ip":             22,
		"t_ua":             23,
		"t_ipua_md5":       24,
		"t_caid_gx":        31,
		"t_caid_ry":        32,
	}
)

func (x DeviceId_DeviceIdType) Enum() *DeviceId_DeviceIdType {
	p := new(DeviceId_DeviceIdType)
	*p = x
	return p
}

func (x DeviceId_DeviceIdType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeviceId_DeviceIdType) Descriptor() protoreflect.EnumDescriptor {
	return file_deviceid_proto_enumTypes[0].Descriptor()
}

func (DeviceId_DeviceIdType) Type() protoreflect.EnumType {
	return &file_deviceid_proto_enumTypes[0]
}

func (x DeviceId_DeviceIdType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeviceId_DeviceIdType.Descriptor instead.
func (DeviceId_DeviceIdType) EnumDescriptor() ([]byte, []int) {
	return file_deviceid_proto_rawDescGZIP(), []int{1, 0}
}

type DeviceIdCollector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceIds string `protobuf:"bytes,1,opt,name=device_ids,json=deviceIds,proto3" json:"device_ids,omitempty"` // DeviceId pb
}

func (x *DeviceIdCollector) Reset() {
	*x = DeviceIdCollector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deviceid_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceIdCollector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceIdCollector) ProtoMessage() {}

func (x *DeviceIdCollector) ProtoReflect() protoreflect.Message {
	mi := &file_deviceid_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceIdCollector.ProtoReflect.Descriptor instead.
func (*DeviceIdCollector) Descriptor() ([]byte, []int) {
	return file_deviceid_proto_rawDescGZIP(), []int{0}
}

func (x *DeviceIdCollector) GetDeviceIds() string {
	if x != nil {
		return x.DeviceIds
	}
	return ""
}

type DeviceId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gaid          string `protobuf:"bytes,1,opt,name=gaid,proto3" json:"gaid,omitempty"`
	GaidMd5       string `protobuf:"bytes,2,opt,name=gaid_md5,json=gaidMd5,proto3" json:"gaid_md5,omitempty"`
	GaidSha1      string `protobuf:"bytes,3,opt,name=gaid_sha1,json=gaidSha1,proto3" json:"gaid_sha1,omitempty"`
	Idfa          string `protobuf:"bytes,4,opt,name=idfa,proto3" json:"idfa,omitempty"`
	IdfaMd5       string `protobuf:"bytes,5,opt,name=idfa_md5,json=idfaMd5,proto3" json:"idfa_md5,omitempty"`
	IdfaSha1      string `protobuf:"bytes,6,opt,name=idfa_sha1,json=idfaSha1,proto3" json:"idfa_sha1,omitempty"`
	Imei          string `protobuf:"bytes,7,opt,name=imei,proto3" json:"imei,omitempty"`
	ImeiMd5       string `protobuf:"bytes,8,opt,name=imei_md5,json=imeiMd5,proto3" json:"imei_md5,omitempty"`
	ImeiSha1      string `protobuf:"bytes,9,opt,name=imei_sha1,json=imeiSha1,proto3" json:"imei_sha1,omitempty"`
	Androidid     string `protobuf:"bytes,10,opt,name=androidid,proto3" json:"androidid,omitempty"`
	AndroididMd5  string `protobuf:"bytes,11,opt,name=androidid_md5,json=androididMd5,proto3" json:"androidid_md5,omitempty"`
	AndroididSha1 string `protobuf:"bytes,12,opt,name=androidid_sha1,json=androididSha1,proto3" json:"androidid_sha1,omitempty"`
	MacMd5        string `protobuf:"bytes,14,opt,name=mac_md5,json=macMd5,proto3" json:"mac_md5,omitempty"`
	MacSha1       string `protobuf:"bytes,15,opt,name=mac_sha1,json=macSha1,proto3" json:"mac_sha1,omitempty"`
	Oaid          string `protobuf:"bytes,16,opt,name=oaid,proto3" json:"oaid,omitempty"`
	OaidMd5       string `protobuf:"bytes,17,opt,name=oaid_md5,json=oaidMd5,proto3" json:"oaid_md5,omitempty"`
	Idfv          string `protobuf:"bytes,19,opt,name=idfv,proto3" json:"idfv,omitempty"`
	Ip            string `protobuf:"bytes,22,opt,name=ip,proto3" json:"ip,omitempty"`
	Ua            string `protobuf:"bytes,23,opt,name=ua,proto3" json:"ua,omitempty"`
	IpuaMd5       string `protobuf:"bytes,24,opt,name=ipua_md5,json=ipuaMd5,proto3" json:"ipua_md5,omitempty"`
	CaidGx        string `protobuf:"bytes,31,opt,name=caid_gx,json=caidGx,proto3" json:"caid_gx,omitempty"` // caid gx 广协版
	CaidRy        string `protobuf:"bytes,32,opt,name=caid_ry,json=caidRy,proto3" json:"caid_ry,omitempty"` // caid ry 热云版
}

func (x *DeviceId) Reset() {
	*x = DeviceId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deviceid_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceId) ProtoMessage() {}

func (x *DeviceId) ProtoReflect() protoreflect.Message {
	mi := &file_deviceid_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceId.ProtoReflect.Descriptor instead.
func (*DeviceId) Descriptor() ([]byte, []int) {
	return file_deviceid_proto_rawDescGZIP(), []int{1}
}

func (x *DeviceId) GetGaid() string {
	if x != nil {
		return x.Gaid
	}
	return ""
}

func (x *DeviceId) GetGaidMd5() string {
	if x != nil {
		return x.GaidMd5
	}
	return ""
}

func (x *DeviceId) GetGaidSha1() string {
	if x != nil {
		return x.GaidSha1
	}
	return ""
}

func (x *DeviceId) GetIdfa() string {
	if x != nil {
		return x.Idfa
	}
	return ""
}

func (x *DeviceId) GetIdfaMd5() string {
	if x != nil {
		return x.IdfaMd5
	}
	return ""
}

func (x *DeviceId) GetIdfaSha1() string {
	if x != nil {
		return x.IdfaSha1
	}
	return ""
}

func (x *DeviceId) GetImei() string {
	if x != nil {
		return x.Imei
	}
	return ""
}

func (x *DeviceId) GetImeiMd5() string {
	if x != nil {
		return x.ImeiMd5
	}
	return ""
}

func (x *DeviceId) GetImeiSha1() string {
	if x != nil {
		return x.ImeiSha1
	}
	return ""
}

func (x *DeviceId) GetAndroidid() string {
	if x != nil {
		return x.Androidid
	}
	return ""
}

func (x *DeviceId) GetAndroididMd5() string {
	if x != nil {
		return x.AndroididMd5
	}
	return ""
}

func (x *DeviceId) GetAndroididSha1() string {
	if x != nil {
		return x.AndroididSha1
	}
	return ""
}

func (x *DeviceId) GetMacMd5() string {
	if x != nil {
		return x.MacMd5
	}
	return ""
}

func (x *DeviceId) GetMacSha1() string {
	if x != nil {
		return x.MacSha1
	}
	return ""
}

func (x *DeviceId) GetOaid() string {
	if x != nil {
		return x.Oaid
	}
	return ""
}

func (x *DeviceId) GetOaidMd5() string {
	if x != nil {
		return x.OaidMd5
	}
	return ""
}

func (x *DeviceId) GetIdfv() string {
	if x != nil {
		return x.Idfv
	}
	return ""
}

func (x *DeviceId) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *DeviceId) GetUa() string {
	if x != nil {
		return x.Ua
	}
	return ""
}

func (x *DeviceId) GetIpuaMd5() string {
	if x != nil {
		return x.IpuaMd5
	}
	return ""
}

func (x *DeviceId) GetCaidGx() string {
	if x != nil {
		return x.CaidGx
	}
	return ""
}

func (x *DeviceId) GetCaidRy() string {
	if x != nil {
		return x.CaidRy
	}
	return ""
}

var File_deviceid_proto protoreflect.FileDescriptor

var file_deviceid_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x69, 0x64, 0x22, 0x32, 0x0a, 0x11, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12,
	0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x73, 0x22, 0xa5,
	0x07, 0x0a, 0x08, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x67,
	0x61, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x67, 0x61, 0x69, 0x64, 0x12,
	0x19, 0x0a, 0x08, 0x67, 0x61, 0x69, 0x64, 0x5f, 0x6d, 0x64, 0x35, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x67, 0x61, 0x69, 0x64, 0x4d, 0x64, 0x35, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x61,
	0x69, 0x64, 0x5f, 0x73, 0x68, 0x61, 0x31, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67,
	0x61, 0x69, 0x64, 0x53, 0x68, 0x61, 0x31, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x64, 0x66, 0x61, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x64, 0x66, 0x61, 0x12, 0x19, 0x0a, 0x08, 0x69,
	0x64, 0x66, 0x61, 0x5f, 0x6d, 0x64, 0x35, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69,
	0x64, 0x66, 0x61, 0x4d, 0x64, 0x35, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x64, 0x66, 0x61, 0x5f, 0x73,
	0x68, 0x61, 0x31, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x64, 0x66, 0x61, 0x53,
	0x68, 0x61, 0x31, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6d, 0x65, 0x69, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x69, 0x6d, 0x65, 0x69, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x6d, 0x65, 0x69, 0x5f,
	0x6d, 0x64, 0x35, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x6d, 0x65, 0x69, 0x4d,
	0x64, 0x35, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x65, 0x69, 0x5f, 0x73, 0x68, 0x61, 0x31, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x65, 0x69, 0x53, 0x68, 0x61, 0x31, 0x12,
	0x1c, 0x0a, 0x09, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x69, 0x64, 0x12, 0x23, 0x0a,
	0x0d, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x69, 0x64, 0x5f, 0x6d, 0x64, 0x35, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x69, 0x64, 0x4d,
	0x64, 0x35, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x69, 0x64, 0x5f,
	0x73, 0x68, 0x61, 0x31, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x6e, 0x64, 0x72,
	0x6f, 0x69, 0x64, 0x69, 0x64, 0x53, 0x68, 0x61, 0x31, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x61, 0x63,
	0x5f, 0x6d, 0x64, 0x35, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x61, 0x63, 0x4d,
	0x64, 0x35, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x61, 0x63, 0x5f, 0x73, 0x68, 0x61, 0x31, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x61, 0x63, 0x53, 0x68, 0x61, 0x31, 0x12, 0x12, 0x0a,
	0x04, 0x6f, 0x61, 0x69, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6f, 0x61, 0x69,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x61, 0x69, 0x64, 0x5f, 0x6d, 0x64, 0x35, 0x18, 0x11, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x61, 0x69, 0x64, 0x4d, 0x64, 0x35, 0x12, 0x12, 0x0a, 0x04,
	0x69, 0x64, 0x66, 0x76, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x64, 0x66, 0x76,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70,
	0x12, 0x0e, 0x0a, 0x02, 0x75, 0x61, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x75, 0x61,
	0x12, 0x19, 0x0a, 0x08, 0x69, 0x70, 0x75, 0x61, 0x5f, 0x6d, 0x64, 0x35, 0x18, 0x18, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x69, 0x70, 0x75, 0x61, 0x4d, 0x64, 0x35, 0x12, 0x17, 0x0a, 0x07, 0x63,
	0x61, 0x69, 0x64, 0x5f, 0x67, 0x78, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x61,
	0x69, 0x64, 0x47, 0x78, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x61, 0x69, 0x64, 0x5f, 0x72, 0x79, 0x18,
	0x20, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x61, 0x69, 0x64, 0x52, 0x79, 0x22, 0xe6, 0x02,
	0x0a, 0x0c, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a,
	0x0a, 0x06, 0x74, 0x5f, 0x6e, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x74, 0x5f,
	0x67, 0x61, 0x69, 0x64, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x74, 0x5f, 0x67, 0x61, 0x69, 0x64,
	0x5f, 0x6d, 0x64, 0x35, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x74, 0x5f, 0x67, 0x61, 0x69, 0x64,
	0x5f, 0x73, 0x68, 0x61, 0x31, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x74, 0x5f, 0x69, 0x64, 0x66,
	0x61, 0x10, 0x04, 0x12, 0x0e, 0x0a, 0x0a, 0x74, 0x5f, 0x69, 0x64, 0x66, 0x61, 0x5f, 0x6d, 0x64,
	0x35, 0x10, 0x05, 0x12, 0x0f, 0x0a, 0x0b, 0x74, 0x5f, 0x69, 0x64, 0x66, 0x61, 0x5f, 0x73, 0x68,
	0x61, 0x31, 0x10, 0x06, 0x12, 0x0a, 0x0a, 0x06, 0x74, 0x5f, 0x69, 0x6d, 0x65, 0x69, 0x10, 0x07,
	0x12, 0x0e, 0x0a, 0x0a, 0x74, 0x5f, 0x69, 0x6d, 0x65, 0x69, 0x5f, 0x6d, 0x64, 0x35, 0x10, 0x08,
	0x12, 0x0f, 0x0a, 0x0b, 0x74, 0x5f, 0x69, 0x6d, 0x65, 0x69, 0x5f, 0x73, 0x68, 0x61, 0x31, 0x10,
	0x09, 0x12, 0x0f, 0x0a, 0x0b, 0x74, 0x5f, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x69, 0x64,
	0x10, 0x0a, 0x12, 0x13, 0x0a, 0x0f, 0x74, 0x5f, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x69,
	0x64, 0x5f, 0x6d, 0x64, 0x35, 0x10, 0x0b, 0x12, 0x14, 0x0a, 0x10, 0x74, 0x5f, 0x61, 0x6e, 0x64,
	0x72, 0x6f, 0x69, 0x64, 0x69, 0x64, 0x5f, 0x73, 0x68, 0x61, 0x31, 0x10, 0x0c, 0x12, 0x0d, 0x0a,
	0x09, 0x74, 0x5f, 0x6d, 0x61, 0x63, 0x5f, 0x6d, 0x64, 0x35, 0x10, 0x0e, 0x12, 0x0e, 0x0a, 0x0a,
	0x74, 0x5f, 0x6d, 0x61, 0x63, 0x5f, 0x73, 0x68, 0x61, 0x31, 0x10, 0x0f, 0x12, 0x0a, 0x0a, 0x06,
	0x74, 0x5f, 0x6f, 0x61, 0x69, 0x64, 0x10, 0x10, 0x12, 0x0e, 0x0a, 0x0a, 0x74, 0x5f, 0x6f, 0x61,
	0x69, 0x64, 0x5f, 0x6d, 0x64, 0x35, 0x10, 0x11, 0x12, 0x0a, 0x0a, 0x06, 0x74, 0x5f, 0x69, 0x64,
	0x66, 0x76, 0x10, 0x13, 0x12, 0x08, 0x0a, 0x04, 0x74, 0x5f, 0x69, 0x70, 0x10, 0x16, 0x12, 0x08,
	0x0a, 0x04, 0x74, 0x5f, 0x75, 0x61, 0x10, 0x17, 0x12, 0x0e, 0x0a, 0x0a, 0x74, 0x5f, 0x69, 0x70,
	0x75, 0x61, 0x5f, 0x6d, 0x64, 0x35, 0x10, 0x18, 0x12, 0x0d, 0x0a, 0x09, 0x74, 0x5f, 0x63, 0x61,
	0x69, 0x64, 0x5f, 0x67, 0x78, 0x10, 0x1f, 0x12, 0x0d, 0x0a, 0x09, 0x74, 0x5f, 0x63, 0x61, 0x69,
	0x64, 0x5f, 0x72, 0x79, 0x10, 0x20, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x2d, 0x30, 0x30, 0x37,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x5f, 0x69, 0x6f, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x69, 0x64, 0x5f, 0x69, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_deviceid_proto_rawDescOnce sync.Once
	file_deviceid_proto_rawDescData = file_deviceid_proto_rawDesc
)

func file_deviceid_proto_rawDescGZIP() []byte {
	file_deviceid_proto_rawDescOnce.Do(func() {
		file_deviceid_proto_rawDescData = protoimpl.X.CompressGZIP(file_deviceid_proto_rawDescData)
	})
	return file_deviceid_proto_rawDescData
}

var file_deviceid_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_deviceid_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_deviceid_proto_goTypes = []interface{}{
	(DeviceId_DeviceIdType)(0), // 0: deviceid.DeviceId.DeviceIdType
	(*DeviceIdCollector)(nil),  // 1: deviceid.DeviceIdCollector
	(*DeviceId)(nil),           // 2: deviceid.DeviceId
}
var file_deviceid_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_deviceid_proto_init() }
func file_deviceid_proto_init() {
	if File_deviceid_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_deviceid_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceIdCollector); i {
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
		file_deviceid_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceId); i {
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
			RawDescriptor: file_deviceid_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_deviceid_proto_goTypes,
		DependencyIndexes: file_deviceid_proto_depIdxs,
		EnumInfos:         file_deviceid_proto_enumTypes,
		MessageInfos:      file_deviceid_proto_msgTypes,
	}.Build()
	File_deviceid_proto = out.File
	file_deviceid_proto_rawDesc = nil
	file_deviceid_proto_goTypes = nil
	file_deviceid_proto_depIdxs = nil
}
