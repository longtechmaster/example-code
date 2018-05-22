// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: media_image.proto

package media

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type UploadImageRequest struct {
	// Link image
	Image *Image `protobuf:"bytes,1,opt,name=Image" json:"Image,omitempty"`
}

func (m *UploadImageRequest) Reset()                    { *m = UploadImageRequest{} }
func (*UploadImageRequest) ProtoMessage()               {}
func (*UploadImageRequest) Descriptor() ([]byte, []int) { return fileDescriptorMediaImage, []int{0} }

func (m *UploadImageRequest) GetImage() *Image {
	if m != nil {
		return m.Image
	}
	return nil
}

type GetImageByIdRequest struct {
	ImageId int32 `protobuf:"varint,1,opt,name=image_id,json=imageId,proto3" json:"image_id,omitempty"`
}

func (m *GetImageByIdRequest) Reset()                    { *m = GetImageByIdRequest{} }
func (*GetImageByIdRequest) ProtoMessage()               {}
func (*GetImageByIdRequest) Descriptor() ([]byte, []int) { return fileDescriptorMediaImage, []int{1} }

func (m *GetImageByIdRequest) GetImageId() int32 {
	if m != nil {
		return m.ImageId
	}
	return 0
}

type GetImageByIdResponse struct {
	Image *Image `protobuf:"bytes,1,opt,name=Image" json:"Image,omitempty"`
}

func (m *GetImageByIdResponse) Reset()                    { *m = GetImageByIdResponse{} }
func (*GetImageByIdResponse) ProtoMessage()               {}
func (*GetImageByIdResponse) Descriptor() ([]byte, []int) { return fileDescriptorMediaImage, []int{2} }

func (m *GetImageByIdResponse) GetImage() *Image {
	if m != nil {
		return m.Image
	}
	return nil
}

func init() {
	proto.RegisterType((*UploadImageRequest)(nil), "media.UploadImageRequest")
	proto.RegisterType((*GetImageByIdRequest)(nil), "media.GetImageByIdRequest")
	proto.RegisterType((*GetImageByIdResponse)(nil), "media.GetImageByIdResponse")
}
func (this *UploadImageRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UploadImageRequest)
	if !ok {
		that2, ok := that.(UploadImageRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Image.Equal(that1.Image) {
		return false
	}
	return true
}
func (this *GetImageByIdRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetImageByIdRequest)
	if !ok {
		that2, ok := that.(GetImageByIdRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ImageId != that1.ImageId {
		return false
	}
	return true
}
func (this *GetImageByIdResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetImageByIdResponse)
	if !ok {
		that2, ok := that.(GetImageByIdResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Image.Equal(that1.Image) {
		return false
	}
	return true
}
func (this *UploadImageRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&media.UploadImageRequest{")
	if this.Image != nil {
		s = append(s, "Image: "+fmt.Sprintf("%#v", this.Image)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GetImageByIdRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&media.GetImageByIdRequest{")
	s = append(s, "ImageId: "+fmt.Sprintf("%#v", this.ImageId)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GetImageByIdResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&media.GetImageByIdResponse{")
	if this.Image != nil {
		s = append(s, "Image: "+fmt.Sprintf("%#v", this.Image)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringMediaImage(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *UploadImageRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UploadImageRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Image != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMediaImage(dAtA, i, uint64(m.Image.Size()))
		n1, err := m.Image.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *GetImageByIdRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetImageByIdRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ImageId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintMediaImage(dAtA, i, uint64(m.ImageId))
	}
	return i, nil
}

func (m *GetImageByIdResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetImageByIdResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Image != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMediaImage(dAtA, i, uint64(m.Image.Size()))
		n2, err := m.Image.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func encodeVarintMediaImage(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *UploadImageRequest) Size() (n int) {
	var l int
	_ = l
	if m.Image != nil {
		l = m.Image.Size()
		n += 1 + l + sovMediaImage(uint64(l))
	}
	return n
}

func (m *GetImageByIdRequest) Size() (n int) {
	var l int
	_ = l
	if m.ImageId != 0 {
		n += 1 + sovMediaImage(uint64(m.ImageId))
	}
	return n
}

func (m *GetImageByIdResponse) Size() (n int) {
	var l int
	_ = l
	if m.Image != nil {
		l = m.Image.Size()
		n += 1 + l + sovMediaImage(uint64(l))
	}
	return n
}

func sovMediaImage(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMediaImage(x uint64) (n int) {
	return sovMediaImage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *UploadImageRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&UploadImageRequest{`,
		`Image:` + strings.Replace(fmt.Sprintf("%v", this.Image), "Image", "Image", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GetImageByIdRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetImageByIdRequest{`,
		`ImageId:` + fmt.Sprintf("%v", this.ImageId) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GetImageByIdResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetImageByIdResponse{`,
		`Image:` + strings.Replace(fmt.Sprintf("%v", this.Image), "Image", "Image", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringMediaImage(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *UploadImageRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMediaImage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UploadImageRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UploadImageRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Image", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMediaImage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMediaImage
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Image == nil {
				m.Image = &Image{}
			}
			if err := m.Image.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMediaImage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMediaImage
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
func (m *GetImageByIdRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMediaImage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetImageByIdRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetImageByIdRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ImageId", wireType)
			}
			m.ImageId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMediaImage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ImageId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMediaImage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMediaImage
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
func (m *GetImageByIdResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMediaImage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetImageByIdResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetImageByIdResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Image", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMediaImage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMediaImage
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Image == nil {
				m.Image = &Image{}
			}
			if err := m.Image.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMediaImage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMediaImage
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
func skipMediaImage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMediaImage
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
					return 0, ErrIntOverflowMediaImage
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
					return 0, ErrIntOverflowMediaImage
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthMediaImage
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMediaImage
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
				next, err := skipMediaImage(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthMediaImage = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMediaImage   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("media_image.proto", fileDescriptorMediaImage) }

var fileDescriptorMediaImage = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcc, 0x4d, 0x4d, 0xc9,
	0x4c, 0x8c, 0xcf, 0xcc, 0x4d, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05,
	0x0b, 0x49, 0x71, 0x25, 0x25, 0x16, 0x43, 0x85, 0x94, 0x2c, 0xb8, 0x84, 0x42, 0x0b, 0x72, 0xf2,
	0x13, 0x53, 0x3c, 0x41, 0xea, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x94, 0xb8, 0x58,
	0xc1, 0x7c, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x1e, 0x3d, 0xb0, 0x46, 0x3d, 0x88, 0x1a,
	0x88, 0x94, 0x92, 0x01, 0x97, 0xb0, 0x7b, 0x6a, 0x09, 0x98, 0xed, 0x54, 0xe9, 0x99, 0x02, 0xd3,
	0x2a, 0xc9, 0xc5, 0x01, 0xb6, 0x32, 0x3e, 0x33, 0x05, 0xac, 0x9b, 0x35, 0x88, 0x1d, 0xcc, 0xf7,
	0x4c, 0x51, 0xb2, 0xe2, 0x12, 0x41, 0xd5, 0x51, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x4a, 0x8c, 0x6d,
	0x4e, 0x3a, 0x17, 0x1e, 0xca, 0x31, 0xdc, 0x78, 0x28, 0xc7, 0xf0, 0xe1, 0xa1, 0x1c, 0x63, 0xc3,
	0x23, 0x39, 0xc6, 0x15, 0x8f, 0xe4, 0x18, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1,
	0xc1, 0x23, 0x39, 0xc6, 0x17, 0x8f, 0xe4, 0x18, 0x3e, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e,
	0x21, 0x89, 0x0d, 0xec, 0x39, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x78, 0x03, 0x42, 0xe6,
	0x04, 0x01, 0x00, 0x00,
}
