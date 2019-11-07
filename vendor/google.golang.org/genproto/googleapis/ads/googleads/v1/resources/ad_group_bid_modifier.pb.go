// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/ad_group_bid_modifier.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "google.golang.org/genproto/googleapis/ads/googleads/v1/common"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Represents an ad group bid modifier.
type AdGroupBidModifier struct {
	// The resource name of the ad group bid modifier.
	// Ad group bid modifier resource names have the form:
	//
	// `customers/{customer_id}/adGroupBidModifiers/{ad_group_id}~{criterion_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ad group to which this criterion belongs.
	AdGroup *wrappers.StringValue `protobuf:"bytes,2,opt,name=ad_group,json=adGroup,proto3" json:"ad_group,omitempty"`
	// The ID of the criterion to bid modify.
	//
	// This field is ignored for mutates.
	CriterionId *wrappers.Int64Value `protobuf:"bytes,3,opt,name=criterion_id,json=criterionId,proto3" json:"criterion_id,omitempty"`
	// The modifier for the bid when the criterion matches. The modifier must be
	// in the range: 0.1 - 10.0. The range is 1.0 - 6.0 for PreferredContent.
	// Use 0 to opt out of a Device type.
	BidModifier *wrappers.DoubleValue `protobuf:"bytes,4,opt,name=bid_modifier,json=bidModifier,proto3" json:"bid_modifier,omitempty"`
	// The base ad group from which this draft/trial adgroup bid modifier was
	// created. If ad_group is a base ad group then this field will be equal to
	// ad_group. If the ad group was created in the draft or trial and has no
	// corresponding base ad group, then this field will be null.
	// This field is readonly.
	BaseAdGroup *wrappers.StringValue `protobuf:"bytes,9,opt,name=base_ad_group,json=baseAdGroup,proto3" json:"base_ad_group,omitempty"`
	// Bid modifier source.
	BidModifierSource enums.BidModifierSourceEnum_BidModifierSource `protobuf:"varint,10,opt,name=bid_modifier_source,json=bidModifierSource,proto3,enum=google.ads.googleads.v1.enums.BidModifierSourceEnum_BidModifierSource" json:"bid_modifier_source,omitempty"`
	// The criterion of this ad group bid modifier.
	//
	// Types that are valid to be assigned to Criterion:
	//	*AdGroupBidModifier_HotelDateSelectionType
	//	*AdGroupBidModifier_HotelAdvanceBookingWindow
	//	*AdGroupBidModifier_HotelLengthOfStay
	//	*AdGroupBidModifier_HotelCheckInDay
	//	*AdGroupBidModifier_Device
	//	*AdGroupBidModifier_PreferredContent
	Criterion            isAdGroupBidModifier_Criterion `protobuf_oneof:"criterion"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *AdGroupBidModifier) Reset()         { *m = AdGroupBidModifier{} }
func (m *AdGroupBidModifier) String() string { return proto.CompactTextString(m) }
func (*AdGroupBidModifier) ProtoMessage()    {}
func (*AdGroupBidModifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_14c1022b135ec410, []int{0}
}

func (m *AdGroupBidModifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupBidModifier.Unmarshal(m, b)
}
func (m *AdGroupBidModifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupBidModifier.Marshal(b, m, deterministic)
}
func (m *AdGroupBidModifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupBidModifier.Merge(m, src)
}
func (m *AdGroupBidModifier) XXX_Size() int {
	return xxx_messageInfo_AdGroupBidModifier.Size(m)
}
func (m *AdGroupBidModifier) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupBidModifier.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupBidModifier proto.InternalMessageInfo

func (m *AdGroupBidModifier) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *AdGroupBidModifier) GetAdGroup() *wrappers.StringValue {
	if m != nil {
		return m.AdGroup
	}
	return nil
}

func (m *AdGroupBidModifier) GetCriterionId() *wrappers.Int64Value {
	if m != nil {
		return m.CriterionId
	}
	return nil
}

func (m *AdGroupBidModifier) GetBidModifier() *wrappers.DoubleValue {
	if m != nil {
		return m.BidModifier
	}
	return nil
}

func (m *AdGroupBidModifier) GetBaseAdGroup() *wrappers.StringValue {
	if m != nil {
		return m.BaseAdGroup
	}
	return nil
}

func (m *AdGroupBidModifier) GetBidModifierSource() enums.BidModifierSourceEnum_BidModifierSource {
	if m != nil {
		return m.BidModifierSource
	}
	return enums.BidModifierSourceEnum_UNSPECIFIED
}

type isAdGroupBidModifier_Criterion interface {
	isAdGroupBidModifier_Criterion()
}

type AdGroupBidModifier_HotelDateSelectionType struct {
	HotelDateSelectionType *common.HotelDateSelectionTypeInfo `protobuf:"bytes,5,opt,name=hotel_date_selection_type,json=hotelDateSelectionType,proto3,oneof"`
}

type AdGroupBidModifier_HotelAdvanceBookingWindow struct {
	HotelAdvanceBookingWindow *common.HotelAdvanceBookingWindowInfo `protobuf:"bytes,6,opt,name=hotel_advance_booking_window,json=hotelAdvanceBookingWindow,proto3,oneof"`
}

type AdGroupBidModifier_HotelLengthOfStay struct {
	HotelLengthOfStay *common.HotelLengthOfStayInfo `protobuf:"bytes,7,opt,name=hotel_length_of_stay,json=hotelLengthOfStay,proto3,oneof"`
}

type AdGroupBidModifier_HotelCheckInDay struct {
	HotelCheckInDay *common.HotelCheckInDayInfo `protobuf:"bytes,8,opt,name=hotel_check_in_day,json=hotelCheckInDay,proto3,oneof"`
}

type AdGroupBidModifier_Device struct {
	Device *common.DeviceInfo `protobuf:"bytes,11,opt,name=device,proto3,oneof"`
}

type AdGroupBidModifier_PreferredContent struct {
	PreferredContent *common.PreferredContentInfo `protobuf:"bytes,12,opt,name=preferred_content,json=preferredContent,proto3,oneof"`
}

func (*AdGroupBidModifier_HotelDateSelectionType) isAdGroupBidModifier_Criterion() {}

func (*AdGroupBidModifier_HotelAdvanceBookingWindow) isAdGroupBidModifier_Criterion() {}

func (*AdGroupBidModifier_HotelLengthOfStay) isAdGroupBidModifier_Criterion() {}

func (*AdGroupBidModifier_HotelCheckInDay) isAdGroupBidModifier_Criterion() {}

func (*AdGroupBidModifier_Device) isAdGroupBidModifier_Criterion() {}

func (*AdGroupBidModifier_PreferredContent) isAdGroupBidModifier_Criterion() {}

func (m *AdGroupBidModifier) GetCriterion() isAdGroupBidModifier_Criterion {
	if m != nil {
		return m.Criterion
	}
	return nil
}

func (m *AdGroupBidModifier) GetHotelDateSelectionType() *common.HotelDateSelectionTypeInfo {
	if x, ok := m.GetCriterion().(*AdGroupBidModifier_HotelDateSelectionType); ok {
		return x.HotelDateSelectionType
	}
	return nil
}

func (m *AdGroupBidModifier) GetHotelAdvanceBookingWindow() *common.HotelAdvanceBookingWindowInfo {
	if x, ok := m.GetCriterion().(*AdGroupBidModifier_HotelAdvanceBookingWindow); ok {
		return x.HotelAdvanceBookingWindow
	}
	return nil
}

func (m *AdGroupBidModifier) GetHotelLengthOfStay() *common.HotelLengthOfStayInfo {
	if x, ok := m.GetCriterion().(*AdGroupBidModifier_HotelLengthOfStay); ok {
		return x.HotelLengthOfStay
	}
	return nil
}

func (m *AdGroupBidModifier) GetHotelCheckInDay() *common.HotelCheckInDayInfo {
	if x, ok := m.GetCriterion().(*AdGroupBidModifier_HotelCheckInDay); ok {
		return x.HotelCheckInDay
	}
	return nil
}

func (m *AdGroupBidModifier) GetDevice() *common.DeviceInfo {
	if x, ok := m.GetCriterion().(*AdGroupBidModifier_Device); ok {
		return x.Device
	}
	return nil
}

func (m *AdGroupBidModifier) GetPreferredContent() *common.PreferredContentInfo {
	if x, ok := m.GetCriterion().(*AdGroupBidModifier_PreferredContent); ok {
		return x.PreferredContent
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AdGroupBidModifier) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AdGroupBidModifier_HotelDateSelectionType)(nil),
		(*AdGroupBidModifier_HotelAdvanceBookingWindow)(nil),
		(*AdGroupBidModifier_HotelLengthOfStay)(nil),
		(*AdGroupBidModifier_HotelCheckInDay)(nil),
		(*AdGroupBidModifier_Device)(nil),
		(*AdGroupBidModifier_PreferredContent)(nil),
	}
}

func init() {
	proto.RegisterType((*AdGroupBidModifier)(nil), "google.ads.googleads.v1.resources.AdGroupBidModifier")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/ad_group_bid_modifier.proto", fileDescriptor_14c1022b135ec410)
}

var fileDescriptor_14c1022b135ec410 = []byte{
	// 693 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdd, 0x6e, 0xd3, 0x48,
	0x14, 0xde, 0xa4, 0xbb, 0xfd, 0x99, 0xa4, 0xbb, 0x5b, 0xef, 0x6a, 0xd7, 0x94, 0x0a, 0xb5, 0xa0,
	0x4a, 0x15, 0x12, 0xb6, 0xd2, 0x16, 0x2a, 0x19, 0x15, 0x48, 0x1a, 0x68, 0x83, 0xf8, 0xa9, 0x12,
	0x14, 0x24, 0x14, 0x69, 0x34, 0xf6, 0x9c, 0x38, 0xa3, 0xc6, 0x33, 0x96, 0x3d, 0x49, 0x94, 0x3b,
	0x2e, 0x78, 0x12, 0x2e, 0xb9, 0xe1, 0x3d, 0x78, 0x14, 0x9e, 0x02, 0x79, 0x66, 0x6c, 0x85, 0xb6,
	0xa1, 0xb9, 0x3b, 0x3e, 0xe7, 0xfb, 0x39, 0xe7, 0x8c, 0x67, 0xd0, 0x71, 0x28, 0x44, 0x38, 0x04,
	0x97, 0xd0, 0xd4, 0xd5, 0x61, 0x16, 0x8d, 0x6b, 0x6e, 0x02, 0xa9, 0x18, 0x25, 0x01, 0xa4, 0x2e,
	0xa1, 0x38, 0x4c, 0xc4, 0x28, 0xc6, 0x3e, 0xa3, 0x38, 0x12, 0x94, 0xf5, 0x19, 0x24, 0x4e, 0x9c,
	0x08, 0x29, 0xac, 0x1d, 0xcd, 0x71, 0x08, 0x4d, 0x9d, 0x82, 0xee, 0x8c, 0x6b, 0x4e, 0x41, 0xdf,
	0x7c, 0x30, 0xcf, 0x21, 0x10, 0x51, 0x24, 0xb8, 0x1b, 0x24, 0x4c, 0x42, 0xc2, 0x88, 0x56, 0xdc,
	0x3c, 0x9a, 0x07, 0x07, 0x3e, 0x8a, 0x52, 0x77, 0xb6, 0x07, 0xac, 0x2d, 0x0c, 0x71, 0x2b, 0x27,
	0xc6, 0xcc, 0x25, 0x9c, 0x0b, 0x49, 0x24, 0x13, 0x3c, 0x35, 0xd5, 0x3b, 0xa6, 0xaa, 0xbe, 0xfc,
	0x51, 0xdf, 0x9d, 0x24, 0x24, 0x8e, 0x21, 0x31, 0xf5, 0xbb, 0x5f, 0x57, 0x91, 0x55, 0xa7, 0xa7,
	0xd9, 0x9c, 0x0d, 0x46, 0x5f, 0x1b, 0x07, 0xeb, 0x1e, 0x5a, 0xcf, 0x27, 0xc1, 0x9c, 0x44, 0x60,
	0x97, 0xb6, 0x4b, 0x7b, 0x6b, 0xed, 0x6a, 0x9e, 0x7c, 0x43, 0x22, 0xb0, 0x8e, 0xd0, 0x6a, 0xbe,
	0x23, 0xbb, 0xbc, 0x5d, 0xda, 0xab, 0xec, 0x6f, 0x99, 0x65, 0x38, 0xb9, 0x9d, 0xd3, 0x91, 0x09,
	0xe3, 0x61, 0x97, 0x0c, 0x47, 0xd0, 0x5e, 0x21, 0xda, 0xc8, 0x7a, 0x82, 0xaa, 0x66, 0x7a, 0xc1,
	0x31, 0xa3, 0xf6, 0x92, 0x22, 0xdf, 0xbe, 0x42, 0x6e, 0x71, 0xf9, 0xe8, 0x50, 0x73, 0x2b, 0x05,
	0xa1, 0x45, 0xad, 0xa7, 0xa8, 0x3a, 0xbb, 0x0f, 0xfb, 0xf7, 0x39, 0xe6, 0x4d, 0x31, 0xf2, 0x87,
	0x60, 0x04, 0xfc, 0x99, 0xf1, 0x9e, 0xa1, 0x75, 0x9f, 0xa4, 0x80, 0x8b, 0xf6, 0xd7, 0x16, 0x68,
	0xbf, 0x92, 0x51, 0xcc, 0xae, 0xac, 0x31, 0xfa, 0xe7, 0x9a, 0x23, 0xb1, 0xd1, 0x76, 0x69, 0xef,
	0xcf, 0xfd, 0x17, 0xce, 0xbc, 0xdf, 0x43, 0x1d, 0xa6, 0x33, 0xb3, 0xe9, 0x8e, 0xe2, 0x3d, 0xe7,
	0xa3, 0xe8, 0x6a, 0xb6, 0xbd, 0xe1, 0x5f, 0x4e, 0x59, 0x13, 0x74, 0x6b, 0x20, 0x24, 0x0c, 0x31,
	0x25, 0x12, 0x70, 0x0a, 0x43, 0x08, 0xb2, 0xe3, 0xc6, 0x72, 0x1a, 0x83, 0xfd, 0x87, 0x9a, 0xc2,
	0x9b, 0xeb, 0xae, 0xff, 0x3c, 0xe7, 0x2c, 0x13, 0x68, 0x12, 0x09, 0x9d, 0x9c, 0xfe, 0x6e, 0x1a,
	0x43, 0x8b, 0xf7, 0xc5, 0xd9, 0x6f, 0xed, 0xff, 0x06, 0xd7, 0x56, 0xad, 0x8f, 0x25, 0xb4, 0xa5,
	0x9d, 0x09, 0x1d, 0x13, 0x1e, 0x00, 0xf6, 0x85, 0xb8, 0x60, 0x3c, 0xc4, 0x13, 0xc6, 0xa9, 0x98,
	0xd8, 0xcb, 0xca, 0xfc, 0x78, 0x21, 0xf3, 0xba, 0x96, 0x68, 0x68, 0x85, 0xf7, 0x4a, 0xc0, 0xf8,
	0xeb, 0xf1, 0xae, 0x03, 0x58, 0x03, 0xf4, 0xaf, 0xee, 0x60, 0x08, 0x3c, 0x94, 0x03, 0x2c, 0xfa,
	0x38, 0x95, 0x64, 0x6a, 0xaf, 0x28, 0xe7, 0x87, 0x0b, 0x39, 0xbf, 0x52, 0xd4, 0xb7, 0xfd, 0x8e,
	0x24, 0x53, 0xe3, 0xb8, 0x31, 0xb8, 0x5c, 0xb0, 0x7c, 0x64, 0x69, 0xa7, 0x60, 0x00, 0xc1, 0x05,
	0x66, 0x1c, 0x53, 0x32, 0xb5, 0x57, 0x95, 0xcf, 0xc1, 0x42, 0x3e, 0x27, 0x19, 0xb1, 0xc5, 0x9b,
	0x85, 0xcb, 0x5f, 0x83, 0x9f, 0xd3, 0x56, 0x13, 0x2d, 0x53, 0x18, 0xb3, 0x00, 0xec, 0x8a, 0xd2,
	0xbd, 0x7f, 0x93, 0x6e, 0x53, 0xa1, 0x8d, 0x9c, 0xe1, 0x5a, 0x01, 0xda, 0x88, 0x13, 0xe8, 0x43,
	0x92, 0x00, 0xc5, 0x81, 0xe0, 0x12, 0xb8, 0xb4, 0xab, 0x4a, 0xf0, 0xf0, 0x26, 0xc1, 0xf3, 0x9c,
	0x78, 0xa2, 0x79, 0x46, 0xfa, 0xef, 0xf8, 0x52, 0xbe, 0x51, 0x41, 0x6b, 0xc5, 0xf5, 0x6b, 0x7c,
	0x2a, 0xa3, 0xdd, 0x40, 0x44, 0xce, 0x8d, 0x2f, 0x60, 0xe3, 0xff, 0xab, 0x0f, 0xcb, 0x79, 0x76,
	0xb1, 0xce, 0x4b, 0x1f, 0x5e, 0x1a, 0x76, 0x28, 0x86, 0x84, 0x87, 0x8e, 0x48, 0x42, 0x37, 0x04,
	0xae, 0xae, 0x5d, 0xfe, 0xfa, 0xc5, 0x2c, 0xfd, 0xc5, 0xeb, 0xfc, 0xb8, 0x88, 0x3e, 0x97, 0x97,
	0x4e, 0xeb, 0xf5, 0x2f, 0xe5, 0x9d, 0x53, 0x2d, 0x59, 0xa7, 0xa9, 0xa3, 0xc3, 0x2c, 0xea, 0xd6,
	0x9c, 0x76, 0x8e, 0xfc, 0x96, 0x63, 0x7a, 0x75, 0x9a, 0xf6, 0x0a, 0x4c, 0xaf, 0x5b, 0xeb, 0x15,
	0x98, 0xef, 0xe5, 0x5d, 0x5d, 0xf0, 0xbc, 0x3a, 0x4d, 0x3d, 0xaf, 0x40, 0x79, 0x5e, 0xb7, 0xe6,
	0x79, 0x05, 0xce, 0x5f, 0x56, 0xcd, 0x1e, 0xfc, 0x08, 0x00, 0x00, 0xff, 0xff, 0x36, 0x2c, 0x67,
	0x25, 0x49, 0x06, 0x00, 0x00,
}