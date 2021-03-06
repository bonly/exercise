// Code generated by protoc-gen-go.
// source: pokemon.proto
// DO NOT EDIT!

/*
Package pk is a generated protocol buffer package.

It is generated from these files:
	pokemon.proto

It has these top-level messages:
	RequestEnvelop
	ResponseEnvelop
*/
package pk

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RequestEnvelop struct {
	Unknown1         *int32                     `protobuf:"varint,1,req,name=unknown1" json:"unknown1,omitempty"`
	RpcId            *int64                     `protobuf:"varint,3,opt,name=rpc_id" json:"rpc_id,omitempty"`
	Requests         []*RequestEnvelop_Requests `protobuf:"bytes,4,rep,name=requests" json:"requests,omitempty"`
	Unknown6         *RequestEnvelop_Unknown6   `protobuf:"bytes,6,opt,name=unknown6" json:"unknown6,omitempty"`
	Latitude         *uint64                    `protobuf:"fixed64,7,opt,name=latitude" json:"latitude,omitempty"`
	Longitude        *uint64                    `protobuf:"fixed64,8,opt,name=longitude" json:"longitude,omitempty"`
	Altitude         *uint64                    `protobuf:"fixed64,9,opt,name=altitude" json:"altitude,omitempty"`
	Auth             *RequestEnvelop_AuthInfo   `protobuf:"bytes,10,opt,name=auth" json:"auth,omitempty"`
	Unknown12        *int64                     `protobuf:"varint,12,opt,name=unknown12" json:"unknown12,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *RequestEnvelop) Reset()                    { *m = RequestEnvelop{} }
func (m *RequestEnvelop) String() string            { return proto.CompactTextString(m) }
func (*RequestEnvelop) ProtoMessage()               {}
func (*RequestEnvelop) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RequestEnvelop) GetUnknown1() int32 {
	if m != nil && m.Unknown1 != nil {
		return *m.Unknown1
	}
	return 0
}

func (m *RequestEnvelop) GetRpcId() int64 {
	if m != nil && m.RpcId != nil {
		return *m.RpcId
	}
	return 0
}

func (m *RequestEnvelop) GetRequests() []*RequestEnvelop_Requests {
	if m != nil {
		return m.Requests
	}
	return nil
}

func (m *RequestEnvelop) GetUnknown6() *RequestEnvelop_Unknown6 {
	if m != nil {
		return m.Unknown6
	}
	return nil
}

func (m *RequestEnvelop) GetLatitude() uint64 {
	if m != nil && m.Latitude != nil {
		return *m.Latitude
	}
	return 0
}

func (m *RequestEnvelop) GetLongitude() uint64 {
	if m != nil && m.Longitude != nil {
		return *m.Longitude
	}
	return 0
}

func (m *RequestEnvelop) GetAltitude() uint64 {
	if m != nil && m.Altitude != nil {
		return *m.Altitude
	}
	return 0
}

func (m *RequestEnvelop) GetAuth() *RequestEnvelop_AuthInfo {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *RequestEnvelop) GetUnknown12() int64 {
	if m != nil && m.Unknown12 != nil {
		return *m.Unknown12
	}
	return 0
}

type RequestEnvelop_Requests struct {
	Type             *int32                   `protobuf:"varint,1,req,name=type" json:"type,omitempty"`
	Message          *RequestEnvelop_Unknown3 `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (m *RequestEnvelop_Requests) Reset()                    { *m = RequestEnvelop_Requests{} }
func (m *RequestEnvelop_Requests) String() string            { return proto.CompactTextString(m) }
func (*RequestEnvelop_Requests) ProtoMessage()               {}
func (*RequestEnvelop_Requests) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *RequestEnvelop_Requests) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *RequestEnvelop_Requests) GetMessage() *RequestEnvelop_Unknown3 {
	if m != nil {
		return m.Message
	}
	return nil
}

type RequestEnvelop_Unknown3 struct {
	Unknown4         *string `protobuf:"bytes,1,req,name=unknown4" json:"unknown4,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RequestEnvelop_Unknown3) Reset()                    { *m = RequestEnvelop_Unknown3{} }
func (m *RequestEnvelop_Unknown3) String() string            { return proto.CompactTextString(m) }
func (*RequestEnvelop_Unknown3) ProtoMessage()               {}
func (*RequestEnvelop_Unknown3) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

func (m *RequestEnvelop_Unknown3) GetUnknown4() string {
	if m != nil && m.Unknown4 != nil {
		return *m.Unknown4
	}
	return ""
}

type RequestEnvelop_Unknown6 struct {
	Unknown1         *int32                            `protobuf:"varint,1,req,name=unknown1" json:"unknown1,omitempty"`
	Unknown2         *RequestEnvelop_Unknown6_Unknown2 `protobuf:"bytes,2,req,name=unknown2" json:"unknown2,omitempty"`
	XXX_unrecognized []byte                            `json:"-"`
}

func (m *RequestEnvelop_Unknown6) Reset()                    { *m = RequestEnvelop_Unknown6{} }
func (m *RequestEnvelop_Unknown6) String() string            { return proto.CompactTextString(m) }
func (*RequestEnvelop_Unknown6) ProtoMessage()               {}
func (*RequestEnvelop_Unknown6) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 2} }

func (m *RequestEnvelop_Unknown6) GetUnknown1() int32 {
	if m != nil && m.Unknown1 != nil {
		return *m.Unknown1
	}
	return 0
}

func (m *RequestEnvelop_Unknown6) GetUnknown2() *RequestEnvelop_Unknown6_Unknown2 {
	if m != nil {
		return m.Unknown2
	}
	return nil
}

type RequestEnvelop_Unknown6_Unknown2 struct {
	Unknown1         []byte `protobuf:"bytes,1,req,name=unknown1" json:"unknown1,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RequestEnvelop_Unknown6_Unknown2) Reset()         { *m = RequestEnvelop_Unknown6_Unknown2{} }
func (m *RequestEnvelop_Unknown6_Unknown2) String() string { return proto.CompactTextString(m) }
func (*RequestEnvelop_Unknown6_Unknown2) ProtoMessage()    {}
func (*RequestEnvelop_Unknown6_Unknown2) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 2, 0}
}

func (m *RequestEnvelop_Unknown6_Unknown2) GetUnknown1() []byte {
	if m != nil {
		return m.Unknown1
	}
	return nil
}

type RequestEnvelop_AuthInfo struct {
	Provider         *string                      `protobuf:"bytes,1,req,name=provider" json:"provider,omitempty"`
	Token            *RequestEnvelop_AuthInfo_JWT `protobuf:"bytes,2,req,name=token" json:"token,omitempty"`
	XXX_unrecognized []byte                       `json:"-"`
}

func (m *RequestEnvelop_AuthInfo) Reset()                    { *m = RequestEnvelop_AuthInfo{} }
func (m *RequestEnvelop_AuthInfo) String() string            { return proto.CompactTextString(m) }
func (*RequestEnvelop_AuthInfo) ProtoMessage()               {}
func (*RequestEnvelop_AuthInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 3} }

func (m *RequestEnvelop_AuthInfo) GetProvider() string {
	if m != nil && m.Provider != nil {
		return *m.Provider
	}
	return ""
}

func (m *RequestEnvelop_AuthInfo) GetToken() *RequestEnvelop_AuthInfo_JWT {
	if m != nil {
		return m.Token
	}
	return nil
}

type RequestEnvelop_AuthInfo_JWT struct {
	Contents         *string `protobuf:"bytes,1,req,name=contents" json:"contents,omitempty"`
	Unknown13        *int32  `protobuf:"varint,2,req,name=unknown13" json:"unknown13,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RequestEnvelop_AuthInfo_JWT) Reset()         { *m = RequestEnvelop_AuthInfo_JWT{} }
func (m *RequestEnvelop_AuthInfo_JWT) String() string { return proto.CompactTextString(m) }
func (*RequestEnvelop_AuthInfo_JWT) ProtoMessage()    {}
func (*RequestEnvelop_AuthInfo_JWT) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 3, 0}
}

func (m *RequestEnvelop_AuthInfo_JWT) GetContents() string {
	if m != nil && m.Contents != nil {
		return *m.Contents
	}
	return ""
}

func (m *RequestEnvelop_AuthInfo_JWT) GetUnknown13() int32 {
	if m != nil && m.Unknown13 != nil {
		return *m.Unknown13
	}
	return 0
}

type ResponseEnvelop struct {
	Unknown1         *int32                     `protobuf:"varint,1,req,name=unknown1" json:"unknown1,omitempty"`
	Unknown2         *int64                     `protobuf:"varint,2,opt,name=unknown2" json:"unknown2,omitempty"`
	ApiUrl           *string                    `protobuf:"bytes,3,opt,name=api_url" json:"api_url,omitempty"`
	Unknown6         *ResponseEnvelop_Unknown6  `protobuf:"bytes,6,opt,name=unknown6" json:"unknown6,omitempty"`
	Unknown7         *ResponseEnvelop_Unknown7  `protobuf:"bytes,7,opt,name=unknown7" json:"unknown7,omitempty"`
	Payload          []*ResponseEnvelop_Payload `protobuf:"bytes,100,rep,name=payload" json:"payload,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *ResponseEnvelop) Reset()                    { *m = ResponseEnvelop{} }
func (m *ResponseEnvelop) String() string            { return proto.CompactTextString(m) }
func (*ResponseEnvelop) ProtoMessage()               {}
func (*ResponseEnvelop) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ResponseEnvelop) GetUnknown1() int32 {
	if m != nil && m.Unknown1 != nil {
		return *m.Unknown1
	}
	return 0
}

func (m *ResponseEnvelop) GetUnknown2() int64 {
	if m != nil && m.Unknown2 != nil {
		return *m.Unknown2
	}
	return 0
}

func (m *ResponseEnvelop) GetApiUrl() string {
	if m != nil && m.ApiUrl != nil {
		return *m.ApiUrl
	}
	return ""
}

func (m *ResponseEnvelop) GetUnknown6() *ResponseEnvelop_Unknown6 {
	if m != nil {
		return m.Unknown6
	}
	return nil
}

func (m *ResponseEnvelop) GetUnknown7() *ResponseEnvelop_Unknown7 {
	if m != nil {
		return m.Unknown7
	}
	return nil
}

func (m *ResponseEnvelop) GetPayload() []*ResponseEnvelop_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

type ResponseEnvelop_Unknown6 struct {
	Unknown1         *int32                             `protobuf:"varint,1,req,name=unknown1" json:"unknown1,omitempty"`
	Unknown2         *ResponseEnvelop_Unknown6_Unknown2 `protobuf:"bytes,2,req,name=unknown2" json:"unknown2,omitempty"`
	XXX_unrecognized []byte                             `json:"-"`
}

func (m *ResponseEnvelop_Unknown6) Reset()                    { *m = ResponseEnvelop_Unknown6{} }
func (m *ResponseEnvelop_Unknown6) String() string            { return proto.CompactTextString(m) }
func (*ResponseEnvelop_Unknown6) ProtoMessage()               {}
func (*ResponseEnvelop_Unknown6) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *ResponseEnvelop_Unknown6) GetUnknown1() int32 {
	if m != nil && m.Unknown1 != nil {
		return *m.Unknown1
	}
	return 0
}

func (m *ResponseEnvelop_Unknown6) GetUnknown2() *ResponseEnvelop_Unknown6_Unknown2 {
	if m != nil {
		return m.Unknown2
	}
	return nil
}

type ResponseEnvelop_Unknown6_Unknown2 struct {
	Unknown1         []byte `protobuf:"bytes,1,req,name=unknown1" json:"unknown1,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ResponseEnvelop_Unknown6_Unknown2) Reset()         { *m = ResponseEnvelop_Unknown6_Unknown2{} }
func (m *ResponseEnvelop_Unknown6_Unknown2) String() string { return proto.CompactTextString(m) }
func (*ResponseEnvelop_Unknown6_Unknown2) ProtoMessage()    {}
func (*ResponseEnvelop_Unknown6_Unknown2) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 0, 0}
}

func (m *ResponseEnvelop_Unknown6_Unknown2) GetUnknown1() []byte {
	if m != nil {
		return m.Unknown1
	}
	return nil
}

type ResponseEnvelop_Unknown7 struct {
	Unknown71        []byte `protobuf:"bytes,1,opt,name=unknown71" json:"unknown71,omitempty"`
	Unknown72        *int64 `protobuf:"varint,2,opt,name=unknown72" json:"unknown72,omitempty"`
	Unknown73        []byte `protobuf:"bytes,3,opt,name=unknown73" json:"unknown73,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ResponseEnvelop_Unknown7) Reset()                    { *m = ResponseEnvelop_Unknown7{} }
func (m *ResponseEnvelop_Unknown7) String() string            { return proto.CompactTextString(m) }
func (*ResponseEnvelop_Unknown7) ProtoMessage()               {}
func (*ResponseEnvelop_Unknown7) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 1} }

func (m *ResponseEnvelop_Unknown7) GetUnknown71() []byte {
	if m != nil {
		return m.Unknown71
	}
	return nil
}

func (m *ResponseEnvelop_Unknown7) GetUnknown72() int64 {
	if m != nil && m.Unknown72 != nil {
		return *m.Unknown72
	}
	return 0
}

func (m *ResponseEnvelop_Unknown7) GetUnknown73() []byte {
	if m != nil {
		return m.Unknown73
	}
	return nil
}

type ResponseEnvelop_Payload struct {
	Unknown1         *int32                   `protobuf:"varint,1,req,name=unknown1" json:"unknown1,omitempty"`
	Profile          *ResponseEnvelop_Profile `protobuf:"bytes,2,opt,name=profile" json:"profile,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (m *ResponseEnvelop_Payload) Reset()                    { *m = ResponseEnvelop_Payload{} }
func (m *ResponseEnvelop_Payload) String() string            { return proto.CompactTextString(m) }
func (*ResponseEnvelop_Payload) ProtoMessage()               {}
func (*ResponseEnvelop_Payload) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 2} }

func (m *ResponseEnvelop_Payload) GetUnknown1() int32 {
	if m != nil && m.Unknown1 != nil {
		return *m.Unknown1
	}
	return 0
}

func (m *ResponseEnvelop_Payload) GetProfile() *ResponseEnvelop_Profile {
	if m != nil {
		return m.Profile
	}
	return nil
}

type ResponseEnvelop_Profile struct {
	CreationTime     *int64                                 `protobuf:"varint,1,req,name=creation_time" json:"creation_time,omitempty"`
	Username         *string                                `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Team             *int32                                 `protobuf:"varint,5,opt,name=team" json:"team,omitempty"`
	Tutorial         []byte                                 `protobuf:"bytes,7,opt,name=tutorial" json:"tutorial,omitempty"`
	Avatar           *ResponseEnvelop_Profile_AvatarDetails `protobuf:"bytes,8,opt,name=avatar" json:"avatar,omitempty"`
	PokeStorage      *int32                                 `protobuf:"varint,9,opt,name=poke_storage" json:"poke_storage,omitempty"`
	ItemStorage      *int32                                 `protobuf:"varint,10,opt,name=item_storage" json:"item_storage,omitempty"`
	DailyBonus       *ResponseEnvelop_Profile_DailyBonus    `protobuf:"bytes,11,opt,name=daily_bonus" json:"daily_bonus,omitempty"`
	Unknown12        []byte                                 `protobuf:"bytes,12,opt,name=unknown12" json:"unknown12,omitempty"`
	Unknown13        []byte                                 `protobuf:"bytes,13,opt,name=unknown13" json:"unknown13,omitempty"`
	Currency         []*ResponseEnvelop_Profile_Currency    `protobuf:"bytes,14,rep,name=currency" json:"currency,omitempty"`
	XXX_unrecognized []byte                                 `json:"-"`
}

func (m *ResponseEnvelop_Profile) Reset()                    { *m = ResponseEnvelop_Profile{} }
func (m *ResponseEnvelop_Profile) String() string            { return proto.CompactTextString(m) }
func (*ResponseEnvelop_Profile) ProtoMessage()               {}
func (*ResponseEnvelop_Profile) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 3} }

func (m *ResponseEnvelop_Profile) GetCreationTime() int64 {
	if m != nil && m.CreationTime != nil {
		return *m.CreationTime
	}
	return 0
}

func (m *ResponseEnvelop_Profile) GetUsername() string {
	if m != nil && m.Username != nil {
		return *m.Username
	}
	return ""
}

func (m *ResponseEnvelop_Profile) GetTeam() int32 {
	if m != nil && m.Team != nil {
		return *m.Team
	}
	return 0
}

func (m *ResponseEnvelop_Profile) GetTutorial() []byte {
	if m != nil {
		return m.Tutorial
	}
	return nil
}

func (m *ResponseEnvelop_Profile) GetAvatar() *ResponseEnvelop_Profile_AvatarDetails {
	if m != nil {
		return m.Avatar
	}
	return nil
}

func (m *ResponseEnvelop_Profile) GetPokeStorage() int32 {
	if m != nil && m.PokeStorage != nil {
		return *m.PokeStorage
	}
	return 0
}

func (m *ResponseEnvelop_Profile) GetItemStorage() int32 {
	if m != nil && m.ItemStorage != nil {
		return *m.ItemStorage
	}
	return 0
}

func (m *ResponseEnvelop_Profile) GetDailyBonus() *ResponseEnvelop_Profile_DailyBonus {
	if m != nil {
		return m.DailyBonus
	}
	return nil
}

func (m *ResponseEnvelop_Profile) GetUnknown12() []byte {
	if m != nil {
		return m.Unknown12
	}
	return nil
}

func (m *ResponseEnvelop_Profile) GetUnknown13() []byte {
	if m != nil {
		return m.Unknown13
	}
	return nil
}

func (m *ResponseEnvelop_Profile) GetCurrency() []*ResponseEnvelop_Profile_Currency {
	if m != nil {
		return m.Currency
	}
	return nil
}

type ResponseEnvelop_Profile_AvatarDetails struct {
	Unknown2         *int32 `protobuf:"varint,2,opt,name=unknown2" json:"unknown2,omitempty"`
	Unknown3         *int32 `protobuf:"varint,3,opt,name=unknown3" json:"unknown3,omitempty"`
	Unknown9         *int32 `protobuf:"varint,9,opt,name=unknown9" json:"unknown9,omitempty"`
	Unknown10        *int32 `protobuf:"varint,10,opt,name=unknown10" json:"unknown10,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ResponseEnvelop_Profile_AvatarDetails) Reset()         { *m = ResponseEnvelop_Profile_AvatarDetails{} }
func (m *ResponseEnvelop_Profile_AvatarDetails) String() string { return proto.CompactTextString(m) }
func (*ResponseEnvelop_Profile_AvatarDetails) ProtoMessage()    {}
func (*ResponseEnvelop_Profile_AvatarDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 3, 0}
}

func (m *ResponseEnvelop_Profile_AvatarDetails) GetUnknown2() int32 {
	if m != nil && m.Unknown2 != nil {
		return *m.Unknown2
	}
	return 0
}

func (m *ResponseEnvelop_Profile_AvatarDetails) GetUnknown3() int32 {
	if m != nil && m.Unknown3 != nil {
		return *m.Unknown3
	}
	return 0
}

func (m *ResponseEnvelop_Profile_AvatarDetails) GetUnknown9() int32 {
	if m != nil && m.Unknown9 != nil {
		return *m.Unknown9
	}
	return 0
}

func (m *ResponseEnvelop_Profile_AvatarDetails) GetUnknown10() int32 {
	if m != nil && m.Unknown10 != nil {
		return *m.Unknown10
	}
	return 0
}

type ResponseEnvelop_Profile_DailyBonus struct {
	NextCollectTimestampMs              *int64 `protobuf:"varint,1,opt,name=NextCollectTimestampMs" json:"NextCollectTimestampMs,omitempty"`
	NextDefenderBonusCollectTimestampMs *int64 `protobuf:"varint,2,opt,name=NextDefenderBonusCollectTimestampMs" json:"NextDefenderBonusCollectTimestampMs,omitempty"`
	XXX_unrecognized                    []byte `json:"-"`
}

func (m *ResponseEnvelop_Profile_DailyBonus) Reset()         { *m = ResponseEnvelop_Profile_DailyBonus{} }
func (m *ResponseEnvelop_Profile_DailyBonus) String() string { return proto.CompactTextString(m) }
func (*ResponseEnvelop_Profile_DailyBonus) ProtoMessage()    {}
func (*ResponseEnvelop_Profile_DailyBonus) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 3, 1}
}

func (m *ResponseEnvelop_Profile_DailyBonus) GetNextCollectTimestampMs() int64 {
	if m != nil && m.NextCollectTimestampMs != nil {
		return *m.NextCollectTimestampMs
	}
	return 0
}

func (m *ResponseEnvelop_Profile_DailyBonus) GetNextDefenderBonusCollectTimestampMs() int64 {
	if m != nil && m.NextDefenderBonusCollectTimestampMs != nil {
		return *m.NextDefenderBonusCollectTimestampMs
	}
	return 0
}

type ResponseEnvelop_Profile_Currency struct {
	Type             *string `protobuf:"bytes,1,req,name=type" json:"type,omitempty"`
	Amount           *int32  `protobuf:"varint,2,opt,name=amount" json:"amount,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ResponseEnvelop_Profile_Currency) Reset()         { *m = ResponseEnvelop_Profile_Currency{} }
func (m *ResponseEnvelop_Profile_Currency) String() string { return proto.CompactTextString(m) }
func (*ResponseEnvelop_Profile_Currency) ProtoMessage()    {}
func (*ResponseEnvelop_Profile_Currency) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 3, 2}
}

func (m *ResponseEnvelop_Profile_Currency) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *ResponseEnvelop_Profile_Currency) GetAmount() int32 {
	if m != nil && m.Amount != nil {
		return *m.Amount
	}
	return 0
}

func init() {
	proto.RegisterType((*RequestEnvelop)(nil), "RequestEnvelop")
	proto.RegisterType((*RequestEnvelop_Requests)(nil), "RequestEnvelop.Requests")
	proto.RegisterType((*RequestEnvelop_Unknown3)(nil), "RequestEnvelop.Unknown3")
	proto.RegisterType((*RequestEnvelop_Unknown6)(nil), "RequestEnvelop.Unknown6")
	proto.RegisterType((*RequestEnvelop_Unknown6_Unknown2)(nil), "RequestEnvelop.Unknown6.Unknown2")
	proto.RegisterType((*RequestEnvelop_AuthInfo)(nil), "RequestEnvelop.AuthInfo")
	proto.RegisterType((*RequestEnvelop_AuthInfo_JWT)(nil), "RequestEnvelop.AuthInfo.JWT")
	proto.RegisterType((*ResponseEnvelop)(nil), "ResponseEnvelop")
	proto.RegisterType((*ResponseEnvelop_Unknown6)(nil), "ResponseEnvelop.Unknown6")
	proto.RegisterType((*ResponseEnvelop_Unknown6_Unknown2)(nil), "ResponseEnvelop.Unknown6.Unknown2")
	proto.RegisterType((*ResponseEnvelop_Unknown7)(nil), "ResponseEnvelop.Unknown7")
	proto.RegisterType((*ResponseEnvelop_Payload)(nil), "ResponseEnvelop.Payload")
	proto.RegisterType((*ResponseEnvelop_Profile)(nil), "ResponseEnvelop.Profile")
	proto.RegisterType((*ResponseEnvelop_Profile_AvatarDetails)(nil), "ResponseEnvelop.Profile.AvatarDetails")
	proto.RegisterType((*ResponseEnvelop_Profile_DailyBonus)(nil), "ResponseEnvelop.Profile.DailyBonus")
	proto.RegisterType((*ResponseEnvelop_Profile_Currency)(nil), "ResponseEnvelop.Profile.Currency")
}

func init() { proto.RegisterFile("pokemon.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 675 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x55, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0x55, 0xfe, 0x9d, 0x89, 0xd3, 0xe6, 0xb3, 0xfa, 0xa1, 0xc5, 0xaa, 0x90, 0xdb, 0x4a, 0x95,
	0x69, 0xa5, 0x08, 0x9c, 0xaa, 0x81, 0xcb, 0x92, 0x82, 0x44, 0x25, 0x10, 0xaa, 0x8a, 0x0a, 0x57,
	0xd1, 0x12, 0x6f, 0x5b, 0xab, 0xeb, 0x5d, 0xb3, 0xbb, 0x2e, 0xe4, 0x95, 0x78, 0x11, 0x1e, 0x89,
	0x5b, 0xb4, 0x6b, 0xc7, 0x75, 0xd2, 0xb8, 0xe2, 0xce, 0x99, 0x39, 0x3e, 0x33, 0x67, 0xe6, 0x8c,
	0x03, 0xfd, 0x84, 0xdf, 0x92, 0x98, 0xb3, 0x61, 0x22, 0xb8, 0xe2, 0xbb, 0xbf, 0x9b, 0xb0, 0x71,
	0x4e, 0xbe, 0xa7, 0x44, 0xaa, 0xb7, 0xec, 0x8e, 0x50, 0x9e, 0x38, 0x03, 0xb0, 0x52, 0x76, 0xcb,
	0xf8, 0x0f, 0xf6, 0x12, 0xd5, 0xbc, 0xba, 0xdf, 0x72, 0x36, 0xa0, 0x2d, 0x92, 0xd9, 0x34, 0x0a,
	0x51, 0xc3, 0xab, 0xf9, 0x0d, 0xe7, 0x00, 0x2c, 0x91, 0xbd, 0x23, 0x51, 0xd3, 0x6b, 0xf8, 0xbd,
	0x00, 0x0d, 0x97, 0x49, 0x16, 0x3f, 0xa5, 0xc6, 0xe6, 0x6c, 0xc7, 0xa8, 0xed, 0xd5, 0xd6, 0x61,
	0x3f, 0xe7, 0x79, 0x5d, 0x99, 0x62, 0x15, 0xa9, 0x34, 0x24, 0xa8, 0xe3, 0xd5, 0xfc, 0xb6, 0xf3,
	0x1f, 0x74, 0x29, 0x67, 0xd7, 0x59, 0xc8, 0x32, 0xa1, 0x01, 0x58, 0x98, 0xe6, 0xa0, 0xae, 0x89,
	0xec, 0x43, 0x13, 0xa7, 0xea, 0x06, 0xc1, 0x7a, 0xfa, 0x93, 0x54, 0xdd, 0xbc, 0x67, 0x57, 0x5c,
	0x93, 0x2d, 0x84, 0x05, 0xc8, 0xd6, 0x4a, 0xdc, 0x09, 0x58, 0x45, 0xa7, 0x36, 0x34, 0xd5, 0x3c,
	0x21, 0xb9, 0xe6, 0xe7, 0xd0, 0x89, 0x89, 0x94, 0xf8, 0x9a, 0xa0, 0xfa, 0xa3, 0x6d, 0x8f, 0xdc,
	0x6d, 0xb0, 0x16, 0xcf, 0xa5, 0xe1, 0x1d, 0x19, 0xa2, 0xae, 0xcb, 0x8b, 0xec, 0xf1, 0x9a, 0xd1,
	0x8e, 0x8a, 0x48, 0x80, 0xea, 0x5e, 0xdd, 0xef, 0x05, 0x3b, 0x55, 0xe3, 0x59, 0x3c, 0x04, 0xa5,
	0x82, 0xc1, 0x03, 0x4a, 0xdb, 0x4d, 0xc1, 0x2a, 0x24, 0x0f, 0xc0, 0x4a, 0x04, 0xbf, 0x8b, 0x42,
	0x22, 0xb2, 0x76, 0x9c, 0x43, 0x68, 0x29, 0x7e, 0x4b, 0x58, 0x5e, 0x6d, 0xbb, 0x6a, 0x5a, 0xc3,
	0xb3, 0xcb, 0x0b, 0xf7, 0x00, 0x1a, 0x67, 0x97, 0x17, 0x9a, 0x65, 0xc6, 0x99, 0x22, 0x4c, 0xc9,
	0x9c, 0xa5, 0x34, 0xca, 0x91, 0x61, 0x6a, 0xed, 0xfe, 0xe9, 0xc0, 0xe6, 0x39, 0x91, 0x09, 0x67,
	0x92, 0x54, 0x5b, 0x69, 0xb0, 0xa4, 0x57, 0x9b, 0x69, 0x13, 0x3a, 0x38, 0x89, 0xa6, 0xa9, 0xa0,
	0xc6, 0x5d, 0xba, 0xc3, 0x55, 0xc7, 0x3c, 0x1d, 0xae, 0x10, 0xdf, 0x5b, 0xe6, 0x1e, 0x3c, 0x36,
	0x96, 0x79, 0x04, 0x3c, 0xd6, 0x3b, 0x4d, 0xf0, 0x9c, 0x72, 0x1c, 0xa2, 0xb0, 0xb0, 0xed, 0x32,
	0xf6, 0x53, 0x96, 0x77, 0x93, 0x47, 0xb7, 0x76, 0xf4, 0x60, 0x6b, 0xbb, 0x95, 0x2d, 0xfe, 0xeb,
	0xda, 0x26, 0x45, 0x76, 0x5c, 0x1a, 0xef, 0x58, 0xa7, 0x6b, 0xbe, 0x5d, 0x0e, 0x2d, 0x26, 0x57,
	0x0a, 0x8d, 0xcc, 0xec, 0x6c, 0xf7, 0x1d, 0x74, 0x72, 0x05, 0x6b, 0xba, 0xd6, 0xf2, 0x05, 0xbf,
	0x8a, 0x68, 0xd9, 0xd2, 0x2b, 0xf2, 0xb3, 0xbc, 0xfb, 0xab, 0x09, 0x9d, 0xfc, 0xd9, 0xf9, 0x1f,
	0xfa, 0x33, 0x41, 0xb0, 0x8a, 0x38, 0x9b, 0xaa, 0x28, 0xce, 0x0e, 0xa4, 0x61, 0xf8, 0x25, 0x11,
	0x0c, 0xc7, 0x19, 0x5d, 0xd7, 0x1c, 0x10, 0xc1, 0x31, 0x6a, 0x79, 0xb5, 0x6c, 0xd3, 0x2a, 0x55,
	0x5c, 0x44, 0x98, 0x9a, 0xcd, 0xd8, 0xce, 0x31, 0xb4, 0xf1, 0x1d, 0x56, 0x58, 0x98, 0x4b, 0xee,
	0x05, 0xfb, 0x55, 0xe5, 0x87, 0x27, 0x06, 0x76, 0x4a, 0x14, 0x8e, 0xa8, 0x74, 0xb6, 0xc0, 0xd6,
	0x1f, 0xad, 0xa9, 0x54, 0x5c, 0xe8, 0x7b, 0xec, 0x1a, 0xfe, 0x2d, 0xb0, 0x23, 0x45, 0xe2, 0x22,
	0x0a, 0x26, 0xfa, 0x0a, 0x7a, 0x21, 0x8e, 0xe8, 0x7c, 0xfa, 0x8d, 0xb3, 0x54, 0xa2, 0x9e, 0x29,
	0xb4, 0x57, 0x59, 0xe8, 0x54, 0x63, 0xdf, 0x68, 0xe8, 0xc3, 0xaf, 0x83, 0xbd, 0xec, 0xf2, 0xbe,
	0x09, 0x8d, 0xc0, 0x9a, 0xa5, 0x42, 0x10, 0x36, 0x9b, 0xa3, 0x0d, 0xe3, 0xa1, 0x9d, 0x4a, 0xf2,
	0x49, 0x0e, 0x74, 0xbf, 0x40, 0x7f, 0x59, 0xd1, 0xea, 0x15, 0x94, 0xef, 0x22, 0x5b, 0x65, 0x39,
	0xf2, 0x3a, 0x57, 0x5c, 0x6a, 0xe7, 0x45, 0x26, 0xd7, 0xfd, 0x0a, 0x50, 0x92, 0xf0, 0x0c, 0x9e,
	0x7c, 0x24, 0x3f, 0xd5, 0x84, 0x53, 0x4a, 0x66, 0xea, 0x22, 0x8a, 0x89, 0x54, 0x38, 0x4e, 0x3e,
	0x48, 0xe3, 0xa1, 0x86, 0x73, 0x08, 0x7b, 0x3a, 0x7f, 0x4a, 0xae, 0x08, 0x0b, 0x89, 0x30, 0x2f,
	0xad, 0x01, 0x1b, 0x77, 0xb9, 0x3e, 0x58, 0x0b, 0x01, 0x4b, 0x9f, 0xc6, 0xae, 0xfe, 0x3b, 0xc0,
	0x31, 0x4f, 0x99, 0xca, 0x7a, 0xff, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x4d, 0xaf, 0x3a, 0x97, 0x53,
	0x06, 0x00, 0x00,
}
