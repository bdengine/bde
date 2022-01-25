// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bde/intrarelayer/v1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the module's genesis state.
type GenesisState struct {
	// module parameters
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// registered token pairs
	TokenPairs []TokenPair `protobuf:"bytes,2,rep,name=token_pairs,json=tokenPairs,proto3" json:"token_pairs"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab56df1cf11ba6de, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetTokenPairs() []TokenPair {
	if m != nil {
		return m.TokenPairs
	}
	return nil
}

// Params defines the intrarelayer module params
type Params struct {
	// parameter to enable the intrarelaying of Cosmos coins <--> ERC20 tokens.
	EnableIntrarelayer bool `protobuf:"varint,1,opt,name=enable_intrarelayer,json=enableIntrarelayer,proto3" json:"enable_intrarelayer,omitempty"`
	// overrides the governance voting period for token pairs proposals
	TokenPairVotingPeriod time.Duration `protobuf:"bytes,2,opt,name=token_pair_voting_period,json=tokenPairVotingPeriod,proto3,stdduration" json:"token_pair_voting_period"`
	// parameter to enable the EVM hook to convert an ERC20 token to a Cosmos
	// Coin by transferring the Tokens through a MsgEthereumTx to the
	// ModuleAddress Ethereum address.
	EnableEVMHook bool `protobuf:"varint,3,opt,name=enable_evm_hook,json=enableEvmHook,proto3" json:"enable_evm_hook,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab56df1cf11ba6de, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetEnableIntrarelayer() bool {
	if m != nil {
		return m.EnableIntrarelayer
	}
	return false
}

func (m *Params) GetTokenPairVotingPeriod() time.Duration {
	if m != nil {
		return m.TokenPairVotingPeriod
	}
	return 0
}

func (m *Params) GetEnableEVMHook() bool {
	if m != nil {
		return m.EnableEVMHook
	}
	return false
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "bde.intrarelayer.v1.GenesisState")
	proto.RegisterType((*Params)(nil), "bde.intrarelayer.v1.Params")
}

func init() {
	proto.RegisterFile("bde/intrarelayer/v1/genesis.proto", fileDescriptor_ab56df1cf11ba6de)
}

var fileDescriptor_ab56df1cf11ba6de = []byte{
	// 387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xc1, 0xae, 0xd2, 0x40,
	0x14, 0x86, 0x3b, 0x60, 0x08, 0x19, 0x24, 0xc6, 0x2a, 0x49, 0x25, 0xb1, 0x10, 0xdc, 0xb0, 0x71,
	0x26, 0xe0, 0xca, 0xb8, 0x6b, 0x24, 0xe8, 0xc2, 0x84, 0x54, 0xc3, 0xc2, 0x98, 0x34, 0x53, 0x19,
	0xcb, 0x04, 0xda, 0xd3, 0xcc, 0x0c, 0x8d, 0xbc, 0x85, 0x0b, 0x17, 0x3e, 0x12, 0x4b, 0xe2, 0xca,
	0x15, 0xde, 0x94, 0x17, 0xb9, 0xe9, 0xb4, 0xdc, 0x0b, 0x09, 0x77, 0x37, 0xe7, 0xfc, 0xff, 0xf9,
	0xfb, 0xf5, 0x1c, 0xfc, 0x8a, 0x67, 0x31, 0x28, 0x2a, 0x12, 0x2d, 0x99, 0xe4, 0x6b, 0xb6, 0xe5,
	0x92, 0x66, 0x23, 0x1a, 0xf1, 0x84, 0x2b, 0xa1, 0x48, 0x2a, 0x41, 0x83, 0xdd, 0x31, 0x26, 0x72,
	0x6e, 0x22, 0xd9, 0xa8, 0x3b, 0xbc, 0x3e, 0x7b, 0x61, 0x33, 0x01, 0xdd, 0xe7, 0x11, 0x44, 0x60,
	0x9e, 0xb4, 0x78, 0x55, 0x5d, 0x37, 0x02, 0x88, 0xd6, 0x9c, 0x9a, 0x2a, 0xdc, 0xfc, 0xa0, 0x8b,
	0x8d, 0x64, 0x5a, 0x40, 0x52, 0xea, 0x83, 0xdf, 0x08, 0x3f, 0x9e, 0x96, 0x20, 0x9f, 0x35, 0xd3,
	0xdc, 0x7e, 0x87, 0x1b, 0x29, 0x93, 0x2c, 0x56, 0x0e, 0xea, 0xa3, 0x61, 0x6b, 0xfc, 0x92, 0x5c,
	0x05, 0x23, 0x33, 0x63, 0xf2, 0x1e, 0xed, 0x0e, 0x3d, 0xcb, 0xaf, 0x46, 0xec, 0x29, 0x6e, 0x69,
	0x58, 0xf1, 0x24, 0x48, 0x99, 0x90, 0xca, 0xa9, 0xf5, 0xeb, 0xc3, 0xd6, 0xb8, 0xff, 0x40, 0xc2,
	0x97, 0xc2, 0x39, 0x63, 0x42, 0x56, 0x21, 0x58, 0x9f, 0x1a, 0x6a, 0xf0, 0x17, 0xe1, 0x46, 0xf9,
	0x05, 0x9b, 0xe2, 0x67, 0x3c, 0x61, 0xe1, 0x9a, 0x07, 0xe7, 0x01, 0x86, 0xae, 0xe9, 0xdb, 0xa5,
	0xf4, 0xf1, 0x4c, 0xb1, 0xbf, 0x61, 0xe7, 0x1e, 0x22, 0xc8, 0x40, 0x8b, 0x24, 0x0a, 0x52, 0x2e,
	0x05, 0x2c, 0x9c, 0x9a, 0xf9, 0xa7, 0x17, 0xa4, 0xdc, 0x0a, 0x39, 0x6d, 0x85, 0xbc, 0xaf, 0xb6,
	0xe2, 0x35, 0x0b, 0x94, 0x3f, 0xff, 0x7b, 0xc8, 0xef, 0xdc, 0xe1, 0xcc, 0x4d, 0xc4, 0xcc, 0x24,
	0xd8, 0x6f, 0xf1, 0x93, 0x0a, 0x87, 0x67, 0x71, 0xb0, 0x04, 0x58, 0x39, 0xf5, 0x02, 0xc5, 0x7b,
	0x9a, 0x1f, 0x7a, 0xed, 0x89, 0x91, 0x26, 0xf3, 0x4f, 0x1f, 0x00, 0x56, 0x7e, 0xbb, 0x74, 0x4e,
	0xb2, 0xb8, 0x28, 0xbd, 0xe9, 0x2e, 0x77, 0xd1, 0x3e, 0x77, 0xd1, 0x4d, 0xee, 0xa2, 0x5f, 0x47,
	0xd7, 0xda, 0x1f, 0x5d, 0xeb, 0xdf, 0xd1, 0xb5, 0xbe, 0xbe, 0x8e, 0x84, 0x5e, 0x6e, 0x42, 0xf2,
	0x1d, 0x62, 0xaa, 0x97, 0x4c, 0x2a, 0xa1, 0x68, 0x79, 0xf8, 0x9f, 0x97, 0xa7, 0xd7, 0xdb, 0x94,
	0xab, 0xb0, 0x61, 0xb8, 0xdf, 0xdc, 0x06, 0x00, 0x00, 0xff, 0xff, 0xb2, 0x2e, 0xc1, 0x49, 0x59,
	0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TokenPairs) > 0 {
		for iNdEx := len(m.TokenPairs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TokenPairs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EnableEVMHook {
		i--
		if m.EnableEVMHook {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	n2, err2 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.TokenPairVotingPeriod, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.TokenPairVotingPeriod):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintGenesis(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x12
	if m.EnableIntrarelayer {
		i--
		if m.EnableIntrarelayer {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.TokenPairs) > 0 {
		for _, e := range m.TokenPairs {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EnableIntrarelayer {
		n += 2
	}
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.TokenPairVotingPeriod)
	n += 1 + l + sovGenesis(uint64(l))
	if m.EnableEVMHook {
		n += 2
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenPairs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenPairs = append(m.TokenPairs, TokenPair{})
			if err := m.TokenPairs[len(m.TokenPairs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnableIntrarelayer", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.EnableIntrarelayer = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenPairVotingPeriod", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.TokenPairVotingPeriod, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnableEVMHook", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.EnableEVMHook = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
