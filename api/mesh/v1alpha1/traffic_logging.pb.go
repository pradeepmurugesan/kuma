// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mesh/v1alpha1/traffic_logging.proto

package v1alpha1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type TrafficLogging struct {
	Rules                []*TrafficLogging_Rule `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *TrafficLogging) Reset()         { *m = TrafficLogging{} }
func (m *TrafficLogging) String() string { return proto.CompactTextString(m) }
func (*TrafficLogging) ProtoMessage()    {}
func (*TrafficLogging) Descriptor() ([]byte, []int) {
	return fileDescriptor_5dbea858e8ab6868, []int{0}
}
func (m *TrafficLogging) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TrafficLogging) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TrafficLogging.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TrafficLogging) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrafficLogging.Merge(m, src)
}
func (m *TrafficLogging) XXX_Size() int {
	return m.Size()
}
func (m *TrafficLogging) XXX_DiscardUnknown() {
	xxx_messageInfo_TrafficLogging.DiscardUnknown(m)
}

var xxx_messageInfo_TrafficLogging proto.InternalMessageInfo

func (m *TrafficLogging) GetRules() []*TrafficLogging_Rule {
	if m != nil {
		return m.Rules
	}
	return nil
}

type TrafficLogging_Rule struct {
	Sources              []*Selector               `protobuf:"bytes,1,rep,name=sources,proto3" json:"sources,omitempty"`
	Destinations         []*Selector               `protobuf:"bytes,2,rep,name=destinations,proto3" json:"destinations,omitempty"`
	Conf                 *TrafficLogging_Rule_Conf `protobuf:"bytes,3,opt,name=conf,proto3" json:"conf,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *TrafficLogging_Rule) Reset()         { *m = TrafficLogging_Rule{} }
func (m *TrafficLogging_Rule) String() string { return proto.CompactTextString(m) }
func (*TrafficLogging_Rule) ProtoMessage()    {}
func (*TrafficLogging_Rule) Descriptor() ([]byte, []int) {
	return fileDescriptor_5dbea858e8ab6868, []int{0, 0}
}
func (m *TrafficLogging_Rule) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TrafficLogging_Rule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TrafficLogging_Rule.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TrafficLogging_Rule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrafficLogging_Rule.Merge(m, src)
}
func (m *TrafficLogging_Rule) XXX_Size() int {
	return m.Size()
}
func (m *TrafficLogging_Rule) XXX_DiscardUnknown() {
	xxx_messageInfo_TrafficLogging_Rule.DiscardUnknown(m)
}

var xxx_messageInfo_TrafficLogging_Rule proto.InternalMessageInfo

func (m *TrafficLogging_Rule) GetSources() []*Selector {
	if m != nil {
		return m.Sources
	}
	return nil
}

func (m *TrafficLogging_Rule) GetDestinations() []*Selector {
	if m != nil {
		return m.Destinations
	}
	return nil
}

func (m *TrafficLogging_Rule) GetConf() *TrafficLogging_Rule_Conf {
	if m != nil {
		return m.Conf
	}
	return nil
}

type TrafficLogging_Rule_Conf struct {
	Backend              string   `protobuf:"bytes,1,opt,name=backend,proto3" json:"backend,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TrafficLogging_Rule_Conf) Reset()         { *m = TrafficLogging_Rule_Conf{} }
func (m *TrafficLogging_Rule_Conf) String() string { return proto.CompactTextString(m) }
func (*TrafficLogging_Rule_Conf) ProtoMessage()    {}
func (*TrafficLogging_Rule_Conf) Descriptor() ([]byte, []int) {
	return fileDescriptor_5dbea858e8ab6868, []int{0, 0, 0}
}
func (m *TrafficLogging_Rule_Conf) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TrafficLogging_Rule_Conf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TrafficLogging_Rule_Conf.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TrafficLogging_Rule_Conf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrafficLogging_Rule_Conf.Merge(m, src)
}
func (m *TrafficLogging_Rule_Conf) XXX_Size() int {
	return m.Size()
}
func (m *TrafficLogging_Rule_Conf) XXX_DiscardUnknown() {
	xxx_messageInfo_TrafficLogging_Rule_Conf.DiscardUnknown(m)
}

var xxx_messageInfo_TrafficLogging_Rule_Conf proto.InternalMessageInfo

func (m *TrafficLogging_Rule_Conf) GetBackend() string {
	if m != nil {
		return m.Backend
	}
	return ""
}

func init() {
	proto.RegisterType((*TrafficLogging)(nil), "kuma.mesh.v1alpha1.TrafficLogging")
	proto.RegisterType((*TrafficLogging_Rule)(nil), "kuma.mesh.v1alpha1.TrafficLogging.Rule")
	proto.RegisterType((*TrafficLogging_Rule_Conf)(nil), "kuma.mesh.v1alpha1.TrafficLogging.Rule.Conf")
}

func init() {
	proto.RegisterFile("mesh/v1alpha1/traffic_logging.proto", fileDescriptor_5dbea858e8ab6868)
}

var fileDescriptor_5dbea858e8ab6868 = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xce, 0x4d, 0x2d, 0xce,
	0xd0, 0x2f, 0x33, 0x4c, 0xcc, 0x29, 0xc8, 0x48, 0x34, 0xd4, 0x2f, 0x29, 0x4a, 0x4c, 0x4b, 0xcb,
	0x4c, 0x8e, 0xcf, 0xc9, 0x4f, 0x4f, 0xcf, 0xcc, 0x4b, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x12, 0xca, 0x2e, 0xcd, 0x4d, 0xd4, 0x03, 0xa9, 0xd4, 0x83, 0xa9, 0x94, 0x92, 0x41, 0xd5, 0x58,
	0x9c, 0x9a, 0x93, 0x9a, 0x5c, 0x92, 0x5f, 0x04, 0xd1, 0xa1, 0xb4, 0x8d, 0x89, 0x8b, 0x2f, 0x04,
	0x62, 0x96, 0x0f, 0xc4, 0x28, 0x21, 0x5b, 0x2e, 0xd6, 0xa2, 0xd2, 0x9c, 0xd4, 0x62, 0x09, 0x46,
	0x05, 0x66, 0x0d, 0x6e, 0x23, 0x75, 0x3d, 0x4c, 0x43, 0xf5, 0x50, 0xb5, 0xe8, 0x05, 0x95, 0xe6,
	0xa4, 0x06, 0x41, 0x74, 0x49, 0x3d, 0x61, 0xe4, 0x62, 0x01, 0xf1, 0x85, 0xcc, 0xb8, 0xd8, 0x8b,
	0xf3, 0x4b, 0x8b, 0x92, 0xe1, 0x26, 0xc9, 0x60, 0x33, 0x29, 0x18, 0xea, 0x9e, 0x20, 0x98, 0x62,
	0x21, 0x07, 0x2e, 0x9e, 0x94, 0xd4, 0xe2, 0x92, 0xcc, 0xbc, 0xc4, 0x92, 0xcc, 0xfc, 0xbc, 0x62,
	0x09, 0x26, 0x22, 0x34, 0xa3, 0xe8, 0x10, 0x72, 0xe0, 0x62, 0x49, 0xce, 0xcf, 0x4b, 0x93, 0x60,
	0x56, 0x60, 0xd4, 0xe0, 0x36, 0xd2, 0x21, 0xd2, 0x03, 0x7a, 0xce, 0xf9, 0x79, 0x69, 0x41, 0x60,
	0x9d, 0x52, 0x0a, 0x5c, 0x2c, 0x20, 0x9e, 0x90, 0x04, 0x17, 0x7b, 0x52, 0x62, 0x72, 0x76, 0x6a,
	0x5e, 0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x8c, 0xeb, 0x24, 0x76, 0xe2, 0x91, 0x1c,
	0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x46, 0x71, 0xc0, 0x0c, 0x4e, 0x62, 0x03,
	0x87, 0xab, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xda, 0x79, 0x39, 0xea, 0xb0, 0x01, 0x00, 0x00,
}

func (m *TrafficLogging) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TrafficLogging) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Rules) > 0 {
		for _, msg := range m.Rules {
			dAtA[i] = 0xa
			i++
			i = encodeVarintTrafficLogging(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *TrafficLogging_Rule) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TrafficLogging_Rule) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Sources) > 0 {
		for _, msg := range m.Sources {
			dAtA[i] = 0xa
			i++
			i = encodeVarintTrafficLogging(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Destinations) > 0 {
		for _, msg := range m.Destinations {
			dAtA[i] = 0x12
			i++
			i = encodeVarintTrafficLogging(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.Conf != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintTrafficLogging(dAtA, i, uint64(m.Conf.Size()))
		n1, err := m.Conf.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *TrafficLogging_Rule_Conf) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TrafficLogging_Rule_Conf) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Backend) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTrafficLogging(dAtA, i, uint64(len(m.Backend)))
		i += copy(dAtA[i:], m.Backend)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintTrafficLogging(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *TrafficLogging) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Rules) > 0 {
		for _, e := range m.Rules {
			l = e.Size()
			n += 1 + l + sovTrafficLogging(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TrafficLogging_Rule) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Sources) > 0 {
		for _, e := range m.Sources {
			l = e.Size()
			n += 1 + l + sovTrafficLogging(uint64(l))
		}
	}
	if len(m.Destinations) > 0 {
		for _, e := range m.Destinations {
			l = e.Size()
			n += 1 + l + sovTrafficLogging(uint64(l))
		}
	}
	if m.Conf != nil {
		l = m.Conf.Size()
		n += 1 + l + sovTrafficLogging(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TrafficLogging_Rule_Conf) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Backend)
	if l > 0 {
		n += 1 + l + sovTrafficLogging(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTrafficLogging(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTrafficLogging(x uint64) (n int) {
	return sovTrafficLogging(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TrafficLogging) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTrafficLogging
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TrafficLogging: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TrafficLogging: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rules", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrafficLogging
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTrafficLogging
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTrafficLogging
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rules = append(m.Rules, &TrafficLogging_Rule{})
			if err := m.Rules[len(m.Rules)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTrafficLogging(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTrafficLogging
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTrafficLogging
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TrafficLogging_Rule) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTrafficLogging
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Rule: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Rule: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sources", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrafficLogging
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTrafficLogging
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTrafficLogging
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sources = append(m.Sources, &Selector{})
			if err := m.Sources[len(m.Sources)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Destinations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrafficLogging
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTrafficLogging
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTrafficLogging
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Destinations = append(m.Destinations, &Selector{})
			if err := m.Destinations[len(m.Destinations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Conf", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrafficLogging
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTrafficLogging
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTrafficLogging
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Conf == nil {
				m.Conf = &TrafficLogging_Rule_Conf{}
			}
			if err := m.Conf.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTrafficLogging(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTrafficLogging
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTrafficLogging
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TrafficLogging_Rule_Conf) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTrafficLogging
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Conf: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Conf: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Backend", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrafficLogging
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTrafficLogging
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTrafficLogging
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Backend = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTrafficLogging(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTrafficLogging
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTrafficLogging
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTrafficLogging(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTrafficLogging
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTrafficLogging
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTrafficLogging
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTrafficLogging
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthTrafficLogging
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTrafficLogging
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipTrafficLogging(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthTrafficLogging
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthTrafficLogging = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTrafficLogging   = fmt.Errorf("proto: integer overflow")
)
