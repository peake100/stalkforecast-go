// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.4
// source: stalk-proto/service.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Enum of the price patterns that can occur in animal crossing.
type PricePatterns int32

const (
	PricePatterns_FLUCTUATING PricePatterns = 0
	PricePatterns_BIGSPIKE    PricePatterns = 1
	PricePatterns_DECREASING  PricePatterns = 2
	PricePatterns_SMALLSPIKE  PricePatterns = 3
	PricePatterns_UNKNOWN     PricePatterns = 4
)

// Enum value maps for PricePatterns.
var (
	PricePatterns_name = map[int32]string{
		0: "FLUCTUATING",
		1: "BIGSPIKE",
		2: "DECREASING",
		3: "SMALLSPIKE",
		4: "UNKNOWN",
	}
	PricePatterns_value = map[string]int32{
		"FLUCTUATING": 0,
		"BIGSPIKE":    1,
		"DECREASING":  2,
		"SMALLSPIKE":  3,
		"UNKNOWN":     4,
	}
)

func (x PricePatterns) Enum() *PricePatterns {
	p := new(PricePatterns)
	*p = x
	return p
}

func (x PricePatterns) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PricePatterns) Descriptor() protoreflect.EnumDescriptor {
	return file_stalk_proto_service_proto_enumTypes[0].Descriptor()
}

func (PricePatterns) Type() protoreflect.EnumType {
	return &file_stalk_proto_service_proto_enumTypes[0]
}

func (x PricePatterns) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PricePatterns.Descriptor instead.
func (PricePatterns) EnumDescriptor() ([]byte, []int) {
	return file_stalk_proto_service_proto_rawDescGZIP(), []int{0}
}

// Island price data sent when requesting a prediction
type Ticker struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The purchase price of the turnips on sunday.
	PurchasePrice int32 `protobuf:"varint,1,opt,name=purchase_price,json=purchasePrice,proto3" json:"purchase_price,omitempty"`
	// The pattern of the previous week's prices for this island.
	PreviousPattern PricePatterns `protobuf:"varint,2,opt,name=previous_pattern,json=previousPattern,proto3,enum=proto.PricePatterns" json:"previous_pattern,omitempty"`
	// The known nook prices for bells. This is expected to be a 12-value array with
	// a price of '0' standing in for "unknown". If this field contains less than
	// 12 values, it is assumed the remaining values are 0. The server will reject a
	// ticker with less than 12 values
	Prices []int32 `protobuf:"varint,3,rep,packed,name=prices,proto3" json:"prices,omitempty"`
}

func (x *Ticker) Reset() {
	*x = Ticker{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stalk_proto_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ticker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ticker) ProtoMessage() {}

func (x *Ticker) ProtoReflect() protoreflect.Message {
	mi := &file_stalk_proto_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ticker.ProtoReflect.Descriptor instead.
func (*Ticker) Descriptor() ([]byte, []int) {
	return file_stalk_proto_service_proto_rawDescGZIP(), []int{0}
}

func (x *Ticker) GetPurchasePrice() int32 {
	if x != nil {
		return x.PurchasePrice
	}
	return 0
}

func (x *Ticker) GetPreviousPattern() PricePatterns {
	if x != nil {
		return x.PreviousPattern
	}
	return PricePatterns_FLUCTUATING
}

func (x *Ticker) GetPrices() []int32 {
	if x != nil {
		return x.Prices
	}
	return nil
}

// Price info for a single price period.
type PricePeriod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The minimum price that might occur during this period.
	Min int32 `protobuf:"varint,1,opt,name=min,proto3" json:"min,omitempty"`
	// The maximum price that might occur during this period.
	Max     int32 `protobuf:"varint,2,opt,name=max,proto3" json:"max,omitempty"`
	IsSpike bool  `protobuf:"varint,3,opt,name=is_spike,json=isSpike,proto3" json:"is_spike,omitempty"`
}

func (x *PricePeriod) Reset() {
	*x = PricePeriod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stalk_proto_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PricePeriod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PricePeriod) ProtoMessage() {}

func (x *PricePeriod) ProtoReflect() protoreflect.Message {
	mi := &file_stalk_proto_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PricePeriod.ProtoReflect.Descriptor instead.
func (*PricePeriod) Descriptor() ([]byte, []int) {
	return file_stalk_proto_service_proto_rawDescGZIP(), []int{1}
}

func (x *PricePeriod) GetMin() int32 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *PricePeriod) GetMax() int32 {
	if x != nil {
		return x.Max
	}
	return 0
}

func (x *PricePeriod) GetIsSpike() bool {
	if x != nil {
		return x.IsSpike
	}
	return false
}

// Price info for a series of prices.
type PricesSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The minimum guaranteed price for the parent. This is the highest minimum price
	// we can say will happen with 100% certainty.
	Min int32 `protobuf:"varint,1,opt,name=min,proto3" json:"min,omitempty"`
	// The potential maximum price of the parent object.
	Max int32 `protobuf:"varint,2,opt,name=max,proto3" json:"max,omitempty"`
	// The price periods during which ``min`` might occur.
	MinPeriods []int32 `protobuf:"varint,3,rep,packed,name=min_periods,json=minPeriods,proto3" json:"min_periods,omitempty"`
	// The price periods during which ``max`` might occur.
	MaxPeriods []int32 `protobuf:"varint,4,rep,packed,name=max_periods,json=maxPeriods,proto3" json:"max_periods,omitempty"`
}

func (x *PricesSummary) Reset() {
	*x = PricesSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stalk_proto_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PricesSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PricesSummary) ProtoMessage() {}

func (x *PricesSummary) ProtoReflect() protoreflect.Message {
	mi := &file_stalk_proto_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PricesSummary.ProtoReflect.Descriptor instead.
func (*PricesSummary) Descriptor() ([]byte, []int) {
	return file_stalk_proto_service_proto_rawDescGZIP(), []int{2}
}

func (x *PricesSummary) GetMin() int32 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *PricesSummary) GetMax() int32 {
	if x != nil {
		return x.Max
	}
	return 0
}

func (x *PricesSummary) GetMinPeriods() []int32 {
	if x != nil {
		return x.MinPeriods
	}
	return nil
}

func (x *PricesSummary) GetMaxPeriods() []int32 {
	if x != nil {
		return x.MaxPeriods
	}
	return nil
}

// Spike info for weeks, weeks and patterns.
type SpikeRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether the parent has this type of spike.
	Has bool `protobuf:"varint,1,opt,name=has,proto3" json:"has,omitempty"`
	// The price period (0-11) the spike starts on.
	Start int32 `protobuf:"varint,2,opt,name=start,proto3" json:"start,omitempty"`
	// The price period (0-11) the spike ends on.
	End int32 `protobuf:"varint,3,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *SpikeRange) Reset() {
	*x = SpikeRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stalk_proto_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpikeRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpikeRange) ProtoMessage() {}

func (x *SpikeRange) ProtoReflect() protoreflect.Message {
	mi := &file_stalk_proto_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpikeRange.ProtoReflect.Descriptor instead.
func (*SpikeRange) Descriptor() ([]byte, []int) {
	return file_stalk_proto_service_proto_rawDescGZIP(), []int{3}
}

func (x *SpikeRange) GetHas() bool {
	if x != nil {
		return x.Has
	}
	return false
}

func (x *SpikeRange) GetStart() int32 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *SpikeRange) GetEnd() int32 {
	if x != nil {
		return x.End
	}
	return 0
}

// Describes a potential price fluctuation for a week of a given price pattern.
type PotentialWeek struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The chance this price fluctuation will occur.
	Chance float64 `protobuf:"fixed64,1,opt,name=chance,proto3" json:"chance,omitempty"`
	// Price information for each price period. Will always have 12 items.
	Prices []*PricePeriod `protobuf:"bytes,3,rep,name=prices,proto3" json:"prices,omitempty"`
	// A summary of the prices in this fluctuation.
	PricesSummary *PricesSummary `protobuf:"bytes,4,opt,name=prices_summary,json=pricesSummary,proto3" json:"prices_summary,omitempty"`
	// A summary of whether there will be a spike and what periods it will occur on.
	Spike *SpikeRange `protobuf:"bytes,5,opt,name=spike,proto3" json:"spike,omitempty"`
}

func (x *PotentialWeek) Reset() {
	*x = PotentialWeek{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stalk_proto_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PotentialWeek) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PotentialWeek) ProtoMessage() {}

func (x *PotentialWeek) ProtoReflect() protoreflect.Message {
	mi := &file_stalk_proto_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PotentialWeek.ProtoReflect.Descriptor instead.
func (*PotentialWeek) Descriptor() ([]byte, []int) {
	return file_stalk_proto_service_proto_rawDescGZIP(), []int{4}
}

func (x *PotentialWeek) GetChance() float64 {
	if x != nil {
		return x.Chance
	}
	return 0
}

func (x *PotentialWeek) GetPrices() []*PricePeriod {
	if x != nil {
		return x.Prices
	}
	return nil
}

func (x *PotentialWeek) GetPricesSummary() *PricesSummary {
	if x != nil {
		return x.PricesSummary
	}
	return nil
}

func (x *PotentialWeek) GetSpike() *SpikeRange {
	if x != nil {
		return x.Spike
	}
	return nil
}

// Describes the potential course a price pattern may take given the user's price
// ticker.
type PotentialPattern struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pattern        PricePatterns    `protobuf:"varint,1,opt,name=pattern,proto3,enum=proto.PricePatterns" json:"pattern,omitempty"`
	Chance         float64          `protobuf:"fixed64,2,opt,name=chance,proto3" json:"chance,omitempty"`
	PricesSummary  *PricesSummary   `protobuf:"bytes,3,opt,name=prices_summary,json=pricesSummary,proto3" json:"prices_summary,omitempty"`
	Spike          *SpikeRange      `protobuf:"bytes,4,opt,name=spike,proto3" json:"spike,omitempty"`
	PotentialWeeks []*PotentialWeek `protobuf:"bytes,5,rep,name=potential_weeks,json=potentialWeeks,proto3" json:"potential_weeks,omitempty"`
}

func (x *PotentialPattern) Reset() {
	*x = PotentialPattern{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stalk_proto_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PotentialPattern) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PotentialPattern) ProtoMessage() {}

func (x *PotentialPattern) ProtoReflect() protoreflect.Message {
	mi := &file_stalk_proto_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PotentialPattern.ProtoReflect.Descriptor instead.
func (*PotentialPattern) Descriptor() ([]byte, []int) {
	return file_stalk_proto_service_proto_rawDescGZIP(), []int{5}
}

func (x *PotentialPattern) GetPattern() PricePatterns {
	if x != nil {
		return x.Pattern
	}
	return PricePatterns_FLUCTUATING
}

func (x *PotentialPattern) GetChance() float64 {
	if x != nil {
		return x.Chance
	}
	return 0
}

func (x *PotentialPattern) GetPricesSummary() *PricesSummary {
	if x != nil {
		return x.PricesSummary
	}
	return nil
}

func (x *PotentialPattern) GetSpike() *SpikeRange {
	if x != nil {
		return x.Spike
	}
	return nil
}

func (x *PotentialPattern) GetPotentialWeeks() []*PotentialWeek {
	if x != nil {
		return x.PotentialWeeks
	}
	return nil
}

// Spike info for the overall prediction.
type SpikeChances struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether the parent has this type of spike.
	Has bool `protobuf:"varint,1,opt,name=has,proto3" json:"has,omitempty"`
	// The price period (0-11) the spike might start on.
	Start int32 `protobuf:"varint,2,opt,name=start,proto3" json:"start,omitempty"`
	// The price period (0-11) the spike might end on.
	End int32 `protobuf:"varint,3,opt,name=end,proto3" json:"end,omitempty"`
	// The overall chance this type of spike will occur
	Chance float64 `protobuf:"fixed64,4,opt,name=chance,proto3" json:"chance,omitempty"`
	// A period-by-period breakdown of the chances this type of spike will occur on the
	// given price period.
	Breakdown []float64 `protobuf:"fixed64,5,rep,packed,name=breakdown,proto3" json:"breakdown,omitempty"`
}

func (x *SpikeChances) Reset() {
	*x = SpikeChances{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stalk_proto_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpikeChances) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpikeChances) ProtoMessage() {}

func (x *SpikeChances) ProtoReflect() protoreflect.Message {
	mi := &file_stalk_proto_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpikeChances.ProtoReflect.Descriptor instead.
func (*SpikeChances) Descriptor() ([]byte, []int) {
	return file_stalk_proto_service_proto_rawDescGZIP(), []int{6}
}

func (x *SpikeChances) GetHas() bool {
	if x != nil {
		return x.Has
	}
	return false
}

func (x *SpikeChances) GetStart() int32 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *SpikeChances) GetEnd() int32 {
	if x != nil {
		return x.End
	}
	return 0
}

func (x *SpikeChances) GetChance() float64 {
	if x != nil {
		return x.Chance
	}
	return 0
}

func (x *SpikeChances) GetBreakdown() []float64 {
	if x != nil {
		return x.Breakdown
	}
	return nil
}

// Contains the overall chances of a price spike occurring this week, as well as a
// breakdown of likelihood for each price period, by spike type.
type ForecastSpikes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The chances of a small spike happening this week.
	Small *SpikeChances `protobuf:"bytes,1,opt,name=small,proto3" json:"small,omitempty"`
	// The chances of a big spike happening this week.
	Big *SpikeChances `protobuf:"bytes,2,opt,name=big,proto3" json:"big,omitempty"`
	// The chances of any spike happening this week.
	Any *SpikeChances `protobuf:"bytes,3,opt,name=any,proto3" json:"any,omitempty"`
}

func (x *ForecastSpikes) Reset() {
	*x = ForecastSpikes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stalk_proto_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForecastSpikes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForecastSpikes) ProtoMessage() {}

func (x *ForecastSpikes) ProtoReflect() protoreflect.Message {
	mi := &file_stalk_proto_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForecastSpikes.ProtoReflect.Descriptor instead.
func (*ForecastSpikes) Descriptor() ([]byte, []int) {
	return file_stalk_proto_service_proto_rawDescGZIP(), []int{7}
}

func (x *ForecastSpikes) GetSmall() *SpikeChances {
	if x != nil {
		return x.Small
	}
	return nil
}

func (x *ForecastSpikes) GetBig() *SpikeChances {
	if x != nil {
		return x.Big
	}
	return nil
}

func (x *ForecastSpikes) GetAny() *SpikeChances {
	if x != nil {
		return x.Any
	}
	return nil
}

// A forecast for prices on your island
type Forecast struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PricesSummary *PricesSummary      `protobuf:"bytes,1,opt,name=prices_summary,json=pricesSummary,proto3" json:"prices_summary,omitempty"`
	Spikes        *ForecastSpikes     `protobuf:"bytes,2,opt,name=spikes,proto3" json:"spikes,omitempty"`
	Patterns      []*PotentialPattern `protobuf:"bytes,3,rep,name=patterns,proto3" json:"patterns,omitempty"`
}

func (x *Forecast) Reset() {
	*x = Forecast{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stalk_proto_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Forecast) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Forecast) ProtoMessage() {}

func (x *Forecast) ProtoReflect() protoreflect.Message {
	mi := &file_stalk_proto_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Forecast.ProtoReflect.Descriptor instead.
func (*Forecast) Descriptor() ([]byte, []int) {
	return file_stalk_proto_service_proto_rawDescGZIP(), []int{8}
}

func (x *Forecast) GetPricesSummary() *PricesSummary {
	if x != nil {
		return x.PricesSummary
	}
	return nil
}

func (x *Forecast) GetSpikes() *ForecastSpikes {
	if x != nil {
		return x.Spikes
	}
	return nil
}

func (x *Forecast) GetPatterns() []*PotentialPattern {
	if x != nil {
		return x.Patterns
	}
	return nil
}

var File_stalk_proto_service_proto protoreflect.FileDescriptor

var file_stalk_proto_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x74, 0x61, 0x6c, 0x6b, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x28, 0x73, 0x74, 0x61, 0x6c, 0x6b, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x01, 0x0a,
	0x06, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x75, 0x72, 0x63, 0x68,
	0x61, 0x73, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0d, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x3f,
	0x0a, 0x10, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x70, 0x61, 0x74, 0x74, 0x65,
	0x72, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x52, 0x0f,
	0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52,
	0x06, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x22, 0x4c, 0x0a, 0x0b, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x6d, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73,
	0x5f, 0x73, 0x70, 0x69, 0x6b, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73,
	0x53, 0x70, 0x69, 0x6b, 0x65, 0x22, 0x75, 0x0a, 0x0d, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x53,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x6d, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x69,
	0x6e, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52,
	0x0a, 0x6d, 0x69, 0x6e, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x6d,
	0x61, 0x78, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x05,
	0x52, 0x0a, 0x6d, 0x61, 0x78, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x73, 0x22, 0x46, 0x0a, 0x0a,
	0x53, 0x70, 0x69, 0x6b, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x68, 0x61,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x68, 0x61, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x65, 0x6e, 0x64, 0x22, 0xb9, 0x01, 0x0a, 0x0d, 0x50, 0x6f, 0x74, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x57, 0x65, 0x65, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x68, 0x61, 0x6e, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x63, 0x68, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x2a,
	0x0a, 0x06, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x50, 0x65, 0x72, 0x69,
	0x6f, 0x64, 0x52, 0x06, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x12, 0x3b, 0x0a, 0x0e, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x73, 0x5f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x73, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x0d, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73,
	0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x27, 0x0a, 0x05, 0x73, 0x70, 0x69, 0x6b, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53,
	0x70, 0x69, 0x6b, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x05, 0x73, 0x70, 0x69, 0x6b, 0x65,
	0x22, 0xff, 0x01, 0x0a, 0x10, 0x50, 0x6f, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x50, 0x61,
	0x74, 0x74, 0x65, 0x72, 0x6e, 0x12, 0x2e, 0x0a, 0x07, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x52, 0x07, 0x70, 0x61,
	0x74, 0x74, 0x65, 0x72, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x68, 0x61, 0x6e, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x63, 0x68, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x3b, 0x0a,
	0x0e, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x5f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x73, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x0d, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x73, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x27, 0x0a, 0x05, 0x73, 0x70,
	0x69, 0x6b, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x53, 0x70, 0x69, 0x6b, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x05, 0x73, 0x70,
	0x69, 0x6b, 0x65, 0x12, 0x3d, 0x0a, 0x0f, 0x70, 0x6f, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x5f, 0x77, 0x65, 0x65, 0x6b, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6f, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x57, 0x65,
	0x65, 0x6b, 0x52, 0x0e, 0x70, 0x6f, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x57, 0x65, 0x65,
	0x6b, 0x73, 0x22, 0x7e, 0x0a, 0x0c, 0x53, 0x70, 0x69, 0x6b, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x63,
	0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x68, 0x61, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x03, 0x68, 0x61, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x63, 0x68, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x63, 0x68,
	0x61, 0x6e, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x64, 0x6f, 0x77,
	0x6e, 0x18, 0x05, 0x20, 0x03, 0x28, 0x01, 0x52, 0x09, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x64, 0x6f,
	0x77, 0x6e, 0x22, 0x89, 0x01, 0x0a, 0x0e, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x53,
	0x70, 0x69, 0x6b, 0x65, 0x73, 0x12, 0x29, 0x0a, 0x05, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x70, 0x69,
	0x6b, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x05, 0x73, 0x6d, 0x61, 0x6c, 0x6c,
	0x12, 0x25, 0x0a, 0x03, 0x62, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x70, 0x69, 0x6b, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x63,
	0x65, 0x73, 0x52, 0x03, 0x62, 0x69, 0x67, 0x12, 0x25, 0x0a, 0x03, 0x61, 0x6e, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x70, 0x69,
	0x6b, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x03, 0x61, 0x6e, 0x79, 0x22, 0xab,
	0x01, 0x0a, 0x08, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x0e, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x73, 0x5f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x73, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x0d, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x73, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x2d, 0x0a, 0x06, 0x73, 0x70, 0x69, 0x6b,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x53, 0x70, 0x69, 0x6b, 0x65, 0x73, 0x52,
	0x06, 0x73, 0x70, 0x69, 0x6b, 0x65, 0x73, 0x12, 0x33, 0x0a, 0x08, 0x70, 0x61, 0x74, 0x74, 0x65,
	0x72, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x50, 0x6f, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x50, 0x61, 0x74, 0x74, 0x65,
	0x72, 0x6e, 0x52, 0x08, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x2a, 0x5b, 0x0a, 0x0d,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x12, 0x0f, 0x0a,
	0x0b, 0x46, 0x4c, 0x55, 0x43, 0x54, 0x55, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0c,
	0x0a, 0x08, 0x42, 0x49, 0x47, 0x53, 0x50, 0x49, 0x4b, 0x45, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a,
	0x44, 0x45, 0x43, 0x52, 0x45, 0x41, 0x53, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a,
	0x53, 0x4d, 0x41, 0x4c, 0x4c, 0x53, 0x50, 0x49, 0x4b, 0x45, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x04, 0x32, 0x59, 0x0a, 0x0f, 0x53, 0x74, 0x61,
	0x6c, 0x6b, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x65, 0x72, 0x12, 0x46, 0x0a, 0x0e,
	0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x12, 0x0d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x1a, 0x0f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x22, 0x14,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x22, 0x09, 0x2f, 0x66, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73,
	0x74, 0x3a, 0x01, 0x2a, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x61, 0x6b, 0x65, 0x31, 0x30, 0x30, 0x2f, 0x73, 0x74, 0x61, 0x6c,
	0x6b, 0x66, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2d, 0x67, 0x6f, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stalk_proto_service_proto_rawDescOnce sync.Once
	file_stalk_proto_service_proto_rawDescData = file_stalk_proto_service_proto_rawDesc
)

func file_stalk_proto_service_proto_rawDescGZIP() []byte {
	file_stalk_proto_service_proto_rawDescOnce.Do(func() {
		file_stalk_proto_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_stalk_proto_service_proto_rawDescData)
	})
	return file_stalk_proto_service_proto_rawDescData
}

var file_stalk_proto_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_stalk_proto_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_stalk_proto_service_proto_goTypes = []interface{}{
	(PricePatterns)(0),       // 0: proto.PricePatterns
	(*Ticker)(nil),           // 1: proto.Ticker
	(*PricePeriod)(nil),      // 2: proto.PricePeriod
	(*PricesSummary)(nil),    // 3: proto.PricesSummary
	(*SpikeRange)(nil),       // 4: proto.SpikeRange
	(*PotentialWeek)(nil),    // 5: proto.PotentialWeek
	(*PotentialPattern)(nil), // 6: proto.PotentialPattern
	(*SpikeChances)(nil),     // 7: proto.SpikeChances
	(*ForecastSpikes)(nil),   // 8: proto.ForecastSpikes
	(*Forecast)(nil),         // 9: proto.Forecast
}
var file_stalk_proto_service_proto_depIdxs = []int32{
	0,  // 0: proto.Ticker.previous_pattern:type_name -> proto.PricePatterns
	2,  // 1: proto.PotentialWeek.prices:type_name -> proto.PricePeriod
	3,  // 2: proto.PotentialWeek.prices_summary:type_name -> proto.PricesSummary
	4,  // 3: proto.PotentialWeek.spike:type_name -> proto.SpikeRange
	0,  // 4: proto.PotentialPattern.pattern:type_name -> proto.PricePatterns
	3,  // 5: proto.PotentialPattern.prices_summary:type_name -> proto.PricesSummary
	4,  // 6: proto.PotentialPattern.spike:type_name -> proto.SpikeRange
	5,  // 7: proto.PotentialPattern.potential_weeks:type_name -> proto.PotentialWeek
	7,  // 8: proto.ForecastSpikes.small:type_name -> proto.SpikeChances
	7,  // 9: proto.ForecastSpikes.big:type_name -> proto.SpikeChances
	7,  // 10: proto.ForecastSpikes.any:type_name -> proto.SpikeChances
	3,  // 11: proto.Forecast.prices_summary:type_name -> proto.PricesSummary
	8,  // 12: proto.Forecast.spikes:type_name -> proto.ForecastSpikes
	6,  // 13: proto.Forecast.patterns:type_name -> proto.PotentialPattern
	1,  // 14: proto.StalkForecaster.ForecastPrices:input_type -> proto.Ticker
	9,  // 15: proto.StalkForecaster.ForecastPrices:output_type -> proto.Forecast
	15, // [15:16] is the sub-list for method output_type
	14, // [14:15] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_stalk_proto_service_proto_init() }
func file_stalk_proto_service_proto_init() {
	if File_stalk_proto_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stalk_proto_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ticker); i {
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
		file_stalk_proto_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PricePeriod); i {
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
		file_stalk_proto_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PricesSummary); i {
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
		file_stalk_proto_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpikeRange); i {
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
		file_stalk_proto_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PotentialWeek); i {
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
		file_stalk_proto_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PotentialPattern); i {
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
		file_stalk_proto_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpikeChances); i {
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
		file_stalk_proto_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForecastSpikes); i {
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
		file_stalk_proto_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Forecast); i {
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
			RawDescriptor: file_stalk_proto_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stalk_proto_service_proto_goTypes,
		DependencyIndexes: file_stalk_proto_service_proto_depIdxs,
		EnumInfos:         file_stalk_proto_service_proto_enumTypes,
		MessageInfos:      file_stalk_proto_service_proto_msgTypes,
	}.Build()
	File_stalk_proto_service_proto = out.File
	file_stalk_proto_service_proto_rawDesc = nil
	file_stalk_proto_service_proto_goTypes = nil
	file_stalk_proto_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StalkForecasterClient is the client API for StalkForecaster service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StalkForecasterClient interface {
	// Returns a turnip price forecast for supplied price ticker.
	ForecastPrices(ctx context.Context, in *Ticker, opts ...grpc.CallOption) (*Forecast, error)
}

type stalkForecasterClient struct {
	cc grpc.ClientConnInterface
}

func NewStalkForecasterClient(cc grpc.ClientConnInterface) StalkForecasterClient {
	return &stalkForecasterClient{cc}
}

func (c *stalkForecasterClient) ForecastPrices(ctx context.Context, in *Ticker, opts ...grpc.CallOption) (*Forecast, error) {
	out := new(Forecast)
	err := c.cc.Invoke(ctx, "/proto.StalkForecaster/ForecastPrices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StalkForecasterServer is the server API for StalkForecaster service.
type StalkForecasterServer interface {
	// Returns a turnip price forecast for supplied price ticker.
	ForecastPrices(context.Context, *Ticker) (*Forecast, error)
}

// UnimplementedStalkForecasterServer can be embedded to have forward compatible implementations.
type UnimplementedStalkForecasterServer struct {
}

func (*UnimplementedStalkForecasterServer) ForecastPrices(context.Context, *Ticker) (*Forecast, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForecastPrices not implemented")
}

func RegisterStalkForecasterServer(s *grpc.Server, srv StalkForecasterServer) {
	s.RegisterService(&_StalkForecaster_serviceDesc, srv)
}

func _StalkForecaster_ForecastPrices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ticker)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StalkForecasterServer).ForecastPrices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.StalkForecaster/ForecastPrices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StalkForecasterServer).ForecastPrices(ctx, req.(*Ticker))
	}
	return interceptor(ctx, in, info, handler)
}

var _StalkForecaster_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.StalkForecaster",
	HandlerType: (*StalkForecasterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ForecastPrices",
			Handler:    _StalkForecaster_ForecastPrices_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stalk-proto/service.proto",
}
