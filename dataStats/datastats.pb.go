// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: define/datastats.proto

package dataStats

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

type OptDataType int32

const (
	OptDataType_TribeIMMsg         OptDataType = 0
	OptDataType_TribeWalletBalance OptDataType = 1
	OptDataType_TribePay           OptDataType = 2
	OptDataType_TribeEnter         OptDataType = 3
)

// Enum value maps for OptDataType.
var (
	OptDataType_name = map[int32]string{
		0: "TribeIMMsg",
		1: "TribeWalletBalance",
		2: "TribePay",
		3: "TribeEnter",
	}
	OptDataType_value = map[string]int32{
		"TribeIMMsg":         0,
		"TribeWalletBalance": 1,
		"TribePay":           2,
		"TribeEnter":         3,
	}
)

func (x OptDataType) Enum() *OptDataType {
	p := new(OptDataType)
	*p = x
	return p
}

func (x OptDataType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OptDataType) Descriptor() protoreflect.EnumDescriptor {
	return file_define_datastats_proto_enumTypes[0].Descriptor()
}

func (OptDataType) Type() protoreflect.EnumType {
	return &file_define_datastats_proto_enumTypes[0]
}

func (x OptDataType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OptDataType.Descriptor instead.
func (OptDataType) EnumDescriptor() ([]byte, []int) {
	return file_define_datastats_proto_rawDescGZIP(), []int{0}
}

type TribeRankType int32

const (
	TribeRankType_ByMsgSum  TribeRankType = 0
	TribeRankType_ByPaySum  TribeRankType = 1
	TribeRankType_ByBalance TribeRankType = 2
	TribeRankType_ByActive  TribeRankType = 3
)

// Enum value maps for TribeRankType.
var (
	TribeRankType_name = map[int32]string{
		0: "ByMsgSum",
		1: "ByPaySum",
		2: "ByBalance",
		3: "ByActive",
	}
	TribeRankType_value = map[string]int32{
		"ByMsgSum":  0,
		"ByPaySum":  1,
		"ByBalance": 2,
		"ByActive":  3,
	}
)

func (x TribeRankType) Enum() *TribeRankType {
	p := new(TribeRankType)
	*p = x
	return p
}

func (x TribeRankType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TribeRankType) Descriptor() protoreflect.EnumDescriptor {
	return file_define_datastats_proto_enumTypes[1].Descriptor()
}

func (TribeRankType) Type() protoreflect.EnumType {
	return &file_define_datastats_proto_enumTypes[1]
}

func (x TribeRankType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TribeRankType.Descriptor instead.
func (TribeRankType) EnumDescriptor() ([]byte, []int) {
	return file_define_datastats_proto_rawDescGZIP(), []int{1}
}

type BaseResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret     int32  `protobuf:"varint,1,opt,name=Ret,proto3" json:"Ret,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *BaseResult) Reset() {
	*x = BaseResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_define_datastats_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseResult) ProtoMessage() {}

func (x *BaseResult) ProtoReflect() protoreflect.Message {
	mi := &file_define_datastats_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseResult.ProtoReflect.Descriptor instead.
func (*BaseResult) Descriptor() ([]byte, []int) {
	return file_define_datastats_proto_rawDescGZIP(), []int{0}
}

func (x *BaseResult) GetRet() int32 {
	if x != nil {
		return x.Ret
	}
	return 0
}

func (x *BaseResult) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ReportDataInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SetType OptDataType       `protobuf:"varint,1,opt,name=SetType,proto3,enum=dataStats.OptDataType" json:"SetType,omitempty"` //
	Time    int64             `protobuf:"varint,2,opt,name=Time,proto3" json:"Time,omitempty"`
	Ext     map[string]string `protobuf:"bytes,3,rep,name=Ext,proto3" json:"Ext,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` //扩展字段
}

func (x *ReportDataInfo) Reset() {
	*x = ReportDataInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_define_datastats_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportDataInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportDataInfo) ProtoMessage() {}

func (x *ReportDataInfo) ProtoReflect() protoreflect.Message {
	mi := &file_define_datastats_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportDataInfo.ProtoReflect.Descriptor instead.
func (*ReportDataInfo) Descriptor() ([]byte, []int) {
	return file_define_datastats_proto_rawDescGZIP(), []int{1}
}

func (x *ReportDataInfo) GetSetType() OptDataType {
	if x != nil {
		return x.SetType
	}
	return OptDataType_TribeIMMsg
}

func (x *ReportDataInfo) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *ReportDataInfo) GetExt() map[string]string {
	if x != nil {
		return x.Ext
	}
	return nil
}

type ReportDataReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info []*ReportDataInfo `protobuf:"bytes,1,rep,name=Info,proto3" json:"Info,omitempty"`
}

func (x *ReportDataReq) Reset() {
	*x = ReportDataReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_define_datastats_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportDataReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportDataReq) ProtoMessage() {}

func (x *ReportDataReq) ProtoReflect() protoreflect.Message {
	mi := &file_define_datastats_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportDataReq.ProtoReflect.Descriptor instead.
func (*ReportDataReq) Descriptor() ([]byte, []int) {
	return file_define_datastats_proto_rawDescGZIP(), []int{2}
}

func (x *ReportDataReq) GetInfo() []*ReportDataInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type ReportDataResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *BaseResult `protobuf:"bytes,1,opt,name=Result,proto3" json:"Result,omitempty"`
}

func (x *ReportDataResp) Reset() {
	*x = ReportDataResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_define_datastats_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportDataResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportDataResp) ProtoMessage() {}

func (x *ReportDataResp) ProtoReflect() protoreflect.Message {
	mi := &file_define_datastats_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportDataResp.ProtoReflect.Descriptor instead.
func (*ReportDataResp) Descriptor() ([]byte, []int) {
	return file_define_datastats_proto_rawDescGZIP(), []int{3}
}

func (x *ReportDataResp) GetResult() *BaseResult {
	if x != nil {
		return x.Result
	}
	return nil
}

type TribeStat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time     int32 `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	Balances int64 `protobuf:"varint,2,opt,name=Balances,proto3" json:"Balances,omitempty"` //过去30天的余额
	Members  int64 `protobuf:"varint,3,opt,name=Members,proto3" json:"Members,omitempty"`   //过去30天的成员数
	Pays     int64 `protobuf:"varint,4,opt,name=Pays,proto3" json:"Pays,omitempty"`         //累计支付额度
}

func (x *TribeStat) Reset() {
	*x = TribeStat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_define_datastats_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TribeStat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TribeStat) ProtoMessage() {}

func (x *TribeStat) ProtoReflect() protoreflect.Message {
	mi := &file_define_datastats_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TribeStat.ProtoReflect.Descriptor instead.
func (*TribeStat) Descriptor() ([]byte, []int) {
	return file_define_datastats_proto_rawDescGZIP(), []int{4}
}

func (x *TribeStat) GetTime() int32 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *TribeStat) GetBalances() int64 {
	if x != nil {
		return x.Balances
	}
	return 0
}

func (x *TribeStat) GetMembers() int64 {
	if x != nil {
		return x.Members
	}
	return 0
}

func (x *TribeStat) GetPays() int64 {
	if x != nil {
		return x.Pays
	}
	return 0
}

type TribeStatsData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TribeID string       `protobuf:"bytes,1,opt,name=TribeID,proto3" json:"TribeID,omitempty"`
	Stats   []*TribeStat `protobuf:"bytes,2,rep,name=Stats,proto3" json:"Stats,omitempty"`
}

func (x *TribeStatsData) Reset() {
	*x = TribeStatsData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_define_datastats_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TribeStatsData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TribeStatsData) ProtoMessage() {}

func (x *TribeStatsData) ProtoReflect() protoreflect.Message {
	mi := &file_define_datastats_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TribeStatsData.ProtoReflect.Descriptor instead.
func (*TribeStatsData) Descriptor() ([]byte, []int) {
	return file_define_datastats_proto_rawDescGZIP(), []int{5}
}

func (x *TribeStatsData) GetTribeID() string {
	if x != nil {
		return x.TribeID
	}
	return ""
}

func (x *TribeStatsData) GetStats() []*TribeStat {
	if x != nil {
		return x.Stats
	}
	return nil
}

type TribeMemberData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID      string `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Msg7Count   int32  `protobuf:"varint,2,opt,name=Msg7Count,proto3" json:"Msg7Count,omitempty"`     //过去7 30天消息
	Msg30Count  int32  `protobuf:"varint,3,opt,name=Msg30Count,proto3" json:"Msg30Count,omitempty"`   //过去7 30天消息
	Balance     int64  `protobuf:"varint,4,opt,name=Balance,proto3" json:"Balance,omitempty"`         //账户余额
	Pay7Sum     int64  `protobuf:"varint,5,opt,name=Pay7Sum,proto3" json:"Pay7Sum,omitempty"`         //支付
	Pay30Sum    int64  `protobuf:"varint,6,opt,name=Pay30Sum,proto3" json:"Pay30Sum,omitempty"`       //支付
	LastEnter   int64  `protobuf:"varint,7,opt,name=LastEnter,proto3" json:"LastEnter,omitempty"`     //最后一次来部落
	LastMsgTime int64  `protobuf:"varint,8,opt,name=LastMsgTime,proto3" json:"LastMsgTime,omitempty"` //最后一次在部落发消息
}

func (x *TribeMemberData) Reset() {
	*x = TribeMemberData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_define_datastats_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TribeMemberData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TribeMemberData) ProtoMessage() {}

func (x *TribeMemberData) ProtoReflect() protoreflect.Message {
	mi := &file_define_datastats_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TribeMemberData.ProtoReflect.Descriptor instead.
func (*TribeMemberData) Descriptor() ([]byte, []int) {
	return file_define_datastats_proto_rawDescGZIP(), []int{6}
}

func (x *TribeMemberData) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *TribeMemberData) GetMsg7Count() int32 {
	if x != nil {
		return x.Msg7Count
	}
	return 0
}

func (x *TribeMemberData) GetMsg30Count() int32 {
	if x != nil {
		return x.Msg30Count
	}
	return 0
}

func (x *TribeMemberData) GetBalance() int64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *TribeMemberData) GetPay7Sum() int64 {
	if x != nil {
		return x.Pay7Sum
	}
	return 0
}

func (x *TribeMemberData) GetPay30Sum() int64 {
	if x != nil {
		return x.Pay30Sum
	}
	return 0
}

func (x *TribeMemberData) GetLastEnter() int64 {
	if x != nil {
		return x.LastEnter
	}
	return 0
}

func (x *TribeMemberData) GetLastMsgTime() int64 {
	if x != nil {
		return x.LastMsgTime
	}
	return 0
}

type GetTribeStatsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TribeID string `protobuf:"bytes,1,opt,name=TribeID,proto3" json:"TribeID,omitempty"`
	Count   int32  `protobuf:"varint,2,opt,name=Count,proto3" json:"Count,omitempty"` //  最长30天的数据，如果只看最新，可以指定1天，默认1天
}

func (x *GetTribeStatsReq) Reset() {
	*x = GetTribeStatsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_define_datastats_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTribeStatsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTribeStatsReq) ProtoMessage() {}

func (x *GetTribeStatsReq) ProtoReflect() protoreflect.Message {
	mi := &file_define_datastats_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTribeStatsReq.ProtoReflect.Descriptor instead.
func (*GetTribeStatsReq) Descriptor() ([]byte, []int) {
	return file_define_datastats_proto_rawDescGZIP(), []int{7}
}

func (x *GetTribeStatsReq) GetTribeID() string {
	if x != nil {
		return x.TribeID
	}
	return ""
}

func (x *GetTribeStatsReq) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type GetTribeStatResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result    *BaseResult     `protobuf:"bytes,1,opt,name=Result,proto3" json:"Result,omitempty"`
	StatsData *TribeStatsData `protobuf:"bytes,2,opt,name=StatsData,proto3" json:"StatsData,omitempty"` //部落趋势数据
}

func (x *GetTribeStatResp) Reset() {
	*x = GetTribeStatResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_define_datastats_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTribeStatResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTribeStatResp) ProtoMessage() {}

func (x *GetTribeStatResp) ProtoReflect() protoreflect.Message {
	mi := &file_define_datastats_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTribeStatResp.ProtoReflect.Descriptor instead.
func (*GetTribeStatResp) Descriptor() ([]byte, []int) {
	return file_define_datastats_proto_rawDescGZIP(), []int{8}
}

func (x *GetTribeStatResp) GetResult() *BaseResult {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *GetTribeStatResp) GetStatsData() *TribeStatsData {
	if x != nil {
		return x.StatsData
	}
	return nil
}

type GetTribeRankReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TribeID  string        `protobuf:"bytes,1,opt,name=TribeID,proto3" json:"TribeID,omitempty"`
	RankType TribeRankType `protobuf:"varint,2,opt,name=RankType,proto3,enum=dataStats.TribeRankType" json:"RankType,omitempty"`
	Pages    int32         `protobuf:"varint,3,opt,name=Pages,proto3" json:"Pages,omitempty"` //默认第一页
	Count    int32         `protobuf:"varint,4,opt,name=Count,proto3" json:"Count,omitempty"` //默认过去30天的
}

func (x *GetTribeRankReq) Reset() {
	*x = GetTribeRankReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_define_datastats_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTribeRankReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTribeRankReq) ProtoMessage() {}

func (x *GetTribeRankReq) ProtoReflect() protoreflect.Message {
	mi := &file_define_datastats_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTribeRankReq.ProtoReflect.Descriptor instead.
func (*GetTribeRankReq) Descriptor() ([]byte, []int) {
	return file_define_datastats_proto_rawDescGZIP(), []int{9}
}

func (x *GetTribeRankReq) GetTribeID() string {
	if x != nil {
		return x.TribeID
	}
	return ""
}

func (x *GetTribeRankReq) GetRankType() TribeRankType {
	if x != nil {
		return x.RankType
	}
	return TribeRankType_ByMsgSum
}

func (x *GetTribeRankReq) GetPages() int32 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *GetTribeRankReq) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type GetTribeRankResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result  *BaseResult        `protobuf:"bytes,1,opt,name=Result,proto3" json:"Result,omitempty"`
	Members []*TribeMemberData `protobuf:"bytes,2,rep,name=Members,proto3" json:"Members,omitempty"` //排行榜
}

func (x *GetTribeRankResp) Reset() {
	*x = GetTribeRankResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_define_datastats_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTribeRankResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTribeRankResp) ProtoMessage() {}

func (x *GetTribeRankResp) ProtoReflect() protoreflect.Message {
	mi := &file_define_datastats_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTribeRankResp.ProtoReflect.Descriptor instead.
func (*GetTribeRankResp) Descriptor() ([]byte, []int) {
	return file_define_datastats_proto_rawDescGZIP(), []int{10}
}

func (x *GetTribeRankResp) GetResult() *BaseResult {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *GetTribeRankResp) GetMembers() []*TribeMemberData {
	if x != nil {
		return x.Members
	}
	return nil
}

var File_define_datastats_proto protoreflect.FileDescriptor

var file_define_datastats_proto_rawDesc = []byte{
	0x0a, 0x16, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x61,
	0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x22, 0x38, 0x0a, 0x0a, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x52, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x52, 0x65, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xc4, 0x01,
	0x0a, 0x0e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x30, 0x0a, 0x07, 0x53, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x16, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x4f, 0x70,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x53, 0x65, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x03, 0x45, 0x78, 0x74, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x61, 0x74, 0x73, 0x2e,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x45,
	0x78, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x45, 0x78, 0x74, 0x1a, 0x36, 0x0a, 0x08,
	0x45, 0x78, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x3e, 0x0a, 0x0d, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x71, 0x12, 0x2d, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x61, 0x74, 0x73, 0x2e,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04,
	0x49, 0x6e, 0x66, 0x6f, 0x22, 0x3f, 0x0a, 0x0e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2d, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x69, 0x0a, 0x09, 0x54, 0x72, 0x69, 0x62, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x50, 0x61, 0x79, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x50, 0x61, 0x79, 0x73,
	0x22, 0x56, 0x0a, 0x0e, 0x54, 0x72, 0x69, 0x62, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x72, 0x69, 0x62, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x54, 0x72, 0x69, 0x62, 0x65, 0x49, 0x44, 0x12, 0x2a, 0x0a, 0x05,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x53, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x54, 0x72, 0x69, 0x62, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x52, 0x05, 0x53, 0x74, 0x61, 0x74, 0x73, 0x22, 0xf7, 0x01, 0x0a, 0x0f, 0x54, 0x72, 0x69,
	0x62, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x4d, 0x73, 0x67, 0x37, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x4d, 0x73, 0x67, 0x37, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x4d, 0x73, 0x67, 0x33, 0x30, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x4d, 0x73, 0x67, 0x33, 0x30, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x50, 0x61, 0x79, 0x37, 0x53, 0x75, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x50,
	0x61, 0x79, 0x37, 0x53, 0x75, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x79, 0x33, 0x30, 0x53,
	0x75, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x50, 0x61, 0x79, 0x33, 0x30, 0x53,
	0x75, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x4c, 0x61, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x4c, 0x61, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x65, 0x72,
	0x12, 0x20, 0x0a, 0x0b, 0x4c, 0x61, 0x73, 0x74, 0x4d, 0x73, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x4c, 0x61, 0x73, 0x74, 0x4d, 0x73, 0x67, 0x54, 0x69,
	0x6d, 0x65, 0x22, 0x42, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x54, 0x72, 0x69, 0x62, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x72, 0x69, 0x62, 0x65, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x54, 0x72, 0x69, 0x62, 0x65, 0x49, 0x44,
	0x12, 0x14, 0x0a, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x7a, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x54, 0x72, 0x69,
	0x62, 0x65, 0x53, 0x74, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2d, 0x0a, 0x06, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x53, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x52, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x37, 0x0a, 0x09, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x53, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x54, 0x72, 0x69, 0x62, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x44, 0x61, 0x74, 0x61, 0x52, 0x09, 0x53, 0x74, 0x61, 0x74, 0x73, 0x44, 0x61,
	0x74, 0x61, 0x22, 0x8d, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54, 0x72, 0x69, 0x62, 0x65, 0x52,
	0x61, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x72, 0x69, 0x62, 0x65, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x54, 0x72, 0x69, 0x62, 0x65, 0x49, 0x44,
	0x12, 0x34, 0x0a, 0x08, 0x52, 0x61, 0x6e, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x18, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x54,
	0x72, 0x69, 0x62, 0x65, 0x52, 0x61, 0x6e, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x52, 0x61,
	0x6e, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x61, 0x67, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x50, 0x61, 0x67, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x77, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x54, 0x72, 0x69, 0x62, 0x65, 0x52, 0x61,
	0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2d, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x34, 0x0a, 0x07, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x2e, 0x54, 0x72, 0x69, 0x62, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x07, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2a, 0x53, 0x0a, 0x0b, 0x4f,
	0x70, 0x74, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x72,
	0x69, 0x62, 0x65, 0x49, 0x4d, 0x4d, 0x73, 0x67, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x54, 0x72,
	0x69, 0x62, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x72, 0x69, 0x62, 0x65, 0x50, 0x61, 0x79, 0x10, 0x02,
	0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x72, 0x69, 0x62, 0x65, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x10, 0x03,
	0x2a, 0x48, 0x0a, 0x0d, 0x54, 0x72, 0x69, 0x62, 0x65, 0x52, 0x61, 0x6e, 0x6b, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0c, 0x0a, 0x08, 0x42, 0x79, 0x4d, 0x73, 0x67, 0x53, 0x75, 0x6d, 0x10, 0x00, 0x12,
	0x0c, 0x0a, 0x08, 0x42, 0x79, 0x50, 0x61, 0x79, 0x53, 0x75, 0x6d, 0x10, 0x01, 0x12, 0x0d, 0x0a,
	0x09, 0x42, 0x79, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08,
	0x42, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x10, 0x03, 0x32, 0xe6, 0x01, 0x0a, 0x09, 0x44,
	0x61, 0x74, 0x61, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x45, 0x0a, 0x0e, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x54, 0x72, 0x69, 0x62, 0x65, 0x4d, 0x73, 0x67, 0x12, 0x18, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x53, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x49, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x54, 0x72, 0x69, 0x62, 0x65, 0x53, 0x74, 0x61, 0x74, 0x61,
	0x12, 0x1b, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74,
	0x54, 0x72, 0x69, 0x62, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x69,
	0x62, 0x65, 0x53, 0x74, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x47, 0x0a, 0x0c, 0x47, 0x65,
	0x74, 0x54, 0x72, 0x69, 0x62, 0x65, 0x52, 0x61, 0x6e, 0x6b, 0x12, 0x1a, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x53, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x69, 0x62, 0x65, 0x52,
	0x61, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x69, 0x62, 0x65, 0x52, 0x61, 0x6e, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_define_datastats_proto_rawDescOnce sync.Once
	file_define_datastats_proto_rawDescData = file_define_datastats_proto_rawDesc
)

func file_define_datastats_proto_rawDescGZIP() []byte {
	file_define_datastats_proto_rawDescOnce.Do(func() {
		file_define_datastats_proto_rawDescData = protoimpl.X.CompressGZIP(file_define_datastats_proto_rawDescData)
	})
	return file_define_datastats_proto_rawDescData
}

var file_define_datastats_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_define_datastats_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_define_datastats_proto_goTypes = []interface{}{
	(OptDataType)(0),         // 0: dataStats.OptDataType
	(TribeRankType)(0),       // 1: dataStats.TribeRankType
	(*BaseResult)(nil),       // 2: dataStats.BaseResult
	(*ReportDataInfo)(nil),   // 3: dataStats.ReportDataInfo
	(*ReportDataReq)(nil),    // 4: dataStats.ReportDataReq
	(*ReportDataResp)(nil),   // 5: dataStats.ReportDataResp
	(*TribeStat)(nil),        // 6: dataStats.TribeStat
	(*TribeStatsData)(nil),   // 7: dataStats.TribeStatsData
	(*TribeMemberData)(nil),  // 8: dataStats.TribeMemberData
	(*GetTribeStatsReq)(nil), // 9: dataStats.GetTribeStatsReq
	(*GetTribeStatResp)(nil), // 10: dataStats.GetTribeStatResp
	(*GetTribeRankReq)(nil),  // 11: dataStats.GetTribeRankReq
	(*GetTribeRankResp)(nil), // 12: dataStats.GetTribeRankResp
	nil,                      // 13: dataStats.ReportDataInfo.ExtEntry
}
var file_define_datastats_proto_depIdxs = []int32{
	0,  // 0: dataStats.ReportDataInfo.SetType:type_name -> dataStats.OptDataType
	13, // 1: dataStats.ReportDataInfo.Ext:type_name -> dataStats.ReportDataInfo.ExtEntry
	3,  // 2: dataStats.ReportDataReq.Info:type_name -> dataStats.ReportDataInfo
	2,  // 3: dataStats.ReportDataResp.Result:type_name -> dataStats.BaseResult
	6,  // 4: dataStats.TribeStatsData.Stats:type_name -> dataStats.TribeStat
	2,  // 5: dataStats.GetTribeStatResp.Result:type_name -> dataStats.BaseResult
	7,  // 6: dataStats.GetTribeStatResp.StatsData:type_name -> dataStats.TribeStatsData
	1,  // 7: dataStats.GetTribeRankReq.RankType:type_name -> dataStats.TribeRankType
	2,  // 8: dataStats.GetTribeRankResp.Result:type_name -> dataStats.BaseResult
	8,  // 9: dataStats.GetTribeRankResp.Members:type_name -> dataStats.TribeMemberData
	4,  // 10: dataStats.DataStats.ReportTribeMsg:input_type -> dataStats.ReportDataReq
	9,  // 11: dataStats.DataStats.GetTribeStata:input_type -> dataStats.GetTribeStatsReq
	11, // 12: dataStats.DataStats.GetTribeRank:input_type -> dataStats.GetTribeRankReq
	5,  // 13: dataStats.DataStats.ReportTribeMsg:output_type -> dataStats.ReportDataResp
	10, // 14: dataStats.DataStats.GetTribeStata:output_type -> dataStats.GetTribeStatResp
	12, // 15: dataStats.DataStats.GetTribeRank:output_type -> dataStats.GetTribeRankResp
	13, // [13:16] is the sub-list for method output_type
	10, // [10:13] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_define_datastats_proto_init() }
func file_define_datastats_proto_init() {
	if File_define_datastats_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_define_datastats_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseResult); i {
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
		file_define_datastats_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportDataInfo); i {
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
		file_define_datastats_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportDataReq); i {
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
		file_define_datastats_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportDataResp); i {
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
		file_define_datastats_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TribeStat); i {
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
		file_define_datastats_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TribeStatsData); i {
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
		file_define_datastats_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TribeMemberData); i {
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
		file_define_datastats_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTribeStatsReq); i {
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
		file_define_datastats_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTribeStatResp); i {
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
		file_define_datastats_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTribeRankReq); i {
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
		file_define_datastats_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTribeRankResp); i {
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
			RawDescriptor: file_define_datastats_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_define_datastats_proto_goTypes,
		DependencyIndexes: file_define_datastats_proto_depIdxs,
		EnumInfos:         file_define_datastats_proto_enumTypes,
		MessageInfos:      file_define_datastats_proto_msgTypes,
	}.Build()
	File_define_datastats_proto = out.File
	file_define_datastats_proto_rawDesc = nil
	file_define_datastats_proto_goTypes = nil
	file_define_datastats_proto_depIdxs = nil
}