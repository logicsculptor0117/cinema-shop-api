// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: services/cinema/rpc/pb/cinema/cinema.proto

package cinema

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

type FilmListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   int32 `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Cate   int32 `protobuf:"varint,2,opt,name=cate,proto3" json:"cate,omitempty"`
	Page   int32 `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Limit  int32 `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	Status int32 `protobuf:"varint,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *FilmListRequest) Reset() {
	*x = FilmListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_cinema_rpc_pb_cinema_cinema_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilmListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilmListRequest) ProtoMessage() {}

func (x *FilmListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_cinema_rpc_pb_cinema_cinema_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilmListRequest.ProtoReflect.Descriptor instead.
func (*FilmListRequest) Descriptor() ([]byte, []int) {
	return file_services_cinema_rpc_pb_cinema_cinema_proto_rawDescGZIP(), []int{0}
}

func (x *FilmListRequest) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *FilmListRequest) GetCate() int32 {
	if x != nil {
		return x.Cate
	}
	return 0
}

func (x *FilmListRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *FilmListRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *FilmListRequest) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type FilmListInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FilmId   int32  `protobuf:"varint,1,opt,name=film_id,json=filmId,proto3" json:"film_id,omitempty"`
	Type     int32  `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	FilmName string `protobuf:"bytes,3,opt,name=film_name,json=filmName,proto3" json:"film_name,omitempty"`
	CoverPic string `protobuf:"bytes,4,opt,name=cover_pic,json=coverPic,proto3" json:"cover_pic,omitempty"`
}

func (x *FilmListInfo) Reset() {
	*x = FilmListInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_cinema_rpc_pb_cinema_cinema_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilmListInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilmListInfo) ProtoMessage() {}

func (x *FilmListInfo) ProtoReflect() protoreflect.Message {
	mi := &file_services_cinema_rpc_pb_cinema_cinema_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilmListInfo.ProtoReflect.Descriptor instead.
func (*FilmListInfo) Descriptor() ([]byte, []int) {
	return file_services_cinema_rpc_pb_cinema_cinema_proto_rawDescGZIP(), []int{1}
}

func (x *FilmListInfo) GetFilmId() int32 {
	if x != nil {
		return x.FilmId
	}
	return 0
}

func (x *FilmListInfo) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *FilmListInfo) GetFilmName() string {
	if x != nil {
		return x.FilmName
	}
	return ""
}

func (x *FilmListInfo) GetCoverPic() string {
	if x != nil {
		return x.CoverPic
	}
	return ""
}

type FilmListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  []*FilmListInfo `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Count int32           `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Page  int32           `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Limit int32           `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *FilmListResponse) Reset() {
	*x = FilmListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_cinema_rpc_pb_cinema_cinema_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilmListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilmListResponse) ProtoMessage() {}

func (x *FilmListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_cinema_rpc_pb_cinema_cinema_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilmListResponse.ProtoReflect.Descriptor instead.
func (*FilmListResponse) Descriptor() ([]byte, []int) {
	return file_services_cinema_rpc_pb_cinema_cinema_proto_rawDescGZIP(), []int{2}
}

func (x *FilmListResponse) GetData() []*FilmListInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *FilmListResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *FilmListResponse) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *FilmListResponse) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type FilmDatailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FilmId int32 `protobuf:"varint,1,opt,name=film_id,json=filmId,proto3" json:"film_id,omitempty"`
}

func (x *FilmDatailRequest) Reset() {
	*x = FilmDatailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_cinema_rpc_pb_cinema_cinema_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilmDatailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilmDatailRequest) ProtoMessage() {}

func (x *FilmDatailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_cinema_rpc_pb_cinema_cinema_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilmDatailRequest.ProtoReflect.Descriptor instead.
func (*FilmDatailRequest) Descriptor() ([]byte, []int) {
	return file_services_cinema_rpc_pb_cinema_cinema_proto_rawDescGZIP(), []int{3}
}

func (x *FilmDatailRequest) GetFilmId() int32 {
	if x != nil {
		return x.FilmId
	}
	return 0
}

type FilmDetailInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FilmId   int32  `protobuf:"varint,1,opt,name=film_id,json=filmId,proto3" json:"film_id,omitempty"`
	Cate     int32  `protobuf:"varint,2,opt,name=cate,proto3" json:"cate,omitempty"`
	FilmName string `protobuf:"bytes,3,opt,name=film_name,json=filmName,proto3" json:"film_name,omitempty"`
	CoverPic string `protobuf:"bytes,4,opt,name=cover_pic,json=coverPic,proto3" json:"cover_pic,omitempty"`
	FilmDesc string `protobuf:"bytes,5,opt,name=film_desc,json=filmDesc,proto3" json:"film_desc,omitempty"`
}

func (x *FilmDetailInfo) Reset() {
	*x = FilmDetailInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_cinema_rpc_pb_cinema_cinema_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilmDetailInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilmDetailInfo) ProtoMessage() {}

func (x *FilmDetailInfo) ProtoReflect() protoreflect.Message {
	mi := &file_services_cinema_rpc_pb_cinema_cinema_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilmDetailInfo.ProtoReflect.Descriptor instead.
func (*FilmDetailInfo) Descriptor() ([]byte, []int) {
	return file_services_cinema_rpc_pb_cinema_cinema_proto_rawDescGZIP(), []int{4}
}

func (x *FilmDetailInfo) GetFilmId() int32 {
	if x != nil {
		return x.FilmId
	}
	return 0
}

func (x *FilmDetailInfo) GetCate() int32 {
	if x != nil {
		return x.Cate
	}
	return 0
}

func (x *FilmDetailInfo) GetFilmName() string {
	if x != nil {
		return x.FilmName
	}
	return ""
}

func (x *FilmDetailInfo) GetCoverPic() string {
	if x != nil {
		return x.CoverPic
	}
	return ""
}

func (x *FilmDetailInfo) GetFilmDesc() string {
	if x != nil {
		return x.FilmDesc
	}
	return ""
}

type FilmDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *FilmDetailInfo `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *FilmDetailResponse) Reset() {
	*x = FilmDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_cinema_rpc_pb_cinema_cinema_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilmDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilmDetailResponse) ProtoMessage() {}

func (x *FilmDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_cinema_rpc_pb_cinema_cinema_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilmDetailResponse.ProtoReflect.Descriptor instead.
func (*FilmDetailResponse) Descriptor() ([]byte, []int) {
	return file_services_cinema_rpc_pb_cinema_cinema_proto_rawDescGZIP(), []int{5}
}

func (x *FilmDetailResponse) GetData() *FilmDetailInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

//根据地理位置获取影院信息
type CinemaInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CityCode string `protobuf:"bytes,1,opt,name=city_code,json=cityCode,proto3" json:"city_code,omitempty"`
	AreaCode string `protobuf:"bytes,2,opt,name=area_code,json=areaCode,proto3" json:"area_code,omitempty"`
}

func (x *CinemaInfoRequest) Reset() {
	*x = CinemaInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_cinema_rpc_pb_cinema_cinema_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CinemaInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CinemaInfoRequest) ProtoMessage() {}

func (x *CinemaInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_cinema_rpc_pb_cinema_cinema_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CinemaInfoRequest.ProtoReflect.Descriptor instead.
func (*CinemaInfoRequest) Descriptor() ([]byte, []int) {
	return file_services_cinema_rpc_pb_cinema_cinema_proto_rawDescGZIP(), []int{6}
}

func (x *CinemaInfoRequest) GetCityCode() string {
	if x != nil {
		return x.CityCode
	}
	return ""
}

func (x *CinemaInfoRequest) GetAreaCode() string {
	if x != nil {
		return x.AreaCode
	}
	return ""
}

type CinemaInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CinemaId   int32  `protobuf:"varint,1,opt,name=cinema_id,json=cinemaId,proto3" json:"cinema_id,omitempty"`
	CinemaName string `protobuf:"bytes,2,opt,name=cinema_name,json=cinemaName,proto3" json:"cinema_name,omitempty"`
	Place      string `protobuf:"bytes,3,opt,name=place,proto3" json:"place,omitempty"`
	Score      int64  `protobuf:"varint,4,opt,name=score,proto3" json:"score,omitempty"`
	Tags       string `protobuf:"bytes,5,opt,name=tags,proto3" json:"tags,omitempty"`
}

func (x *CinemaInfo) Reset() {
	*x = CinemaInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_cinema_rpc_pb_cinema_cinema_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CinemaInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CinemaInfo) ProtoMessage() {}

func (x *CinemaInfo) ProtoReflect() protoreflect.Message {
	mi := &file_services_cinema_rpc_pb_cinema_cinema_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CinemaInfo.ProtoReflect.Descriptor instead.
func (*CinemaInfo) Descriptor() ([]byte, []int) {
	return file_services_cinema_rpc_pb_cinema_cinema_proto_rawDescGZIP(), []int{7}
}

func (x *CinemaInfo) GetCinemaId() int32 {
	if x != nil {
		return x.CinemaId
	}
	return 0
}

func (x *CinemaInfo) GetCinemaName() string {
	if x != nil {
		return x.CinemaName
	}
	return ""
}

func (x *CinemaInfo) GetPlace() string {
	if x != nil {
		return x.Place
	}
	return ""
}

func (x *CinemaInfo) GetScore() int64 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *CinemaInfo) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

type CinemaInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*CinemaInfo `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *CinemaInfoResponse) Reset() {
	*x = CinemaInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_cinema_rpc_pb_cinema_cinema_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CinemaInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CinemaInfoResponse) ProtoMessage() {}

func (x *CinemaInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_cinema_rpc_pb_cinema_cinema_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CinemaInfoResponse.ProtoReflect.Descriptor instead.
func (*CinemaInfoResponse) Descriptor() ([]byte, []int) {
	return file_services_cinema_rpc_pb_cinema_cinema_proto_rawDescGZIP(), []int{8}
}

func (x *CinemaInfoResponse) GetData() []*CinemaInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

//根据影片和日期和影院ID获取排片信息
type ScreenCinemaInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FilmId    int32   `protobuf:"varint,1,opt,name=film_id,json=filmId,proto3" json:"film_id,omitempty"`
	TDate     string  `protobuf:"bytes,2,opt,name=t_date,json=tDate,proto3" json:"t_date,omitempty"`
	CinemaIds []int32 `protobuf:"varint,3,rep,packed,name=cinema_ids,json=cinemaIds,proto3" json:"cinema_ids,omitempty"`
}

func (x *ScreenCinemaInfoRequest) Reset() {
	*x = ScreenCinemaInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_cinema_rpc_pb_cinema_cinema_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScreenCinemaInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScreenCinemaInfoRequest) ProtoMessage() {}

func (x *ScreenCinemaInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_cinema_rpc_pb_cinema_cinema_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScreenCinemaInfoRequest.ProtoReflect.Descriptor instead.
func (*ScreenCinemaInfoRequest) Descriptor() ([]byte, []int) {
	return file_services_cinema_rpc_pb_cinema_cinema_proto_rawDescGZIP(), []int{9}
}

func (x *ScreenCinemaInfoRequest) GetFilmId() int32 {
	if x != nil {
		return x.FilmId
	}
	return 0
}

func (x *ScreenCinemaInfoRequest) GetTDate() string {
	if x != nil {
		return x.TDate
	}
	return ""
}

func (x *ScreenCinemaInfoRequest) GetCinemaIds() []int32 {
	if x != nil {
		return x.CinemaIds
	}
	return nil
}

type ScreenCinemaInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CinemaId  int32  `protobuf:"varint,1,opt,name=cinema_id,json=cinemaId,proto3" json:"cinema_id,omitempty"`
	Price     int32  `protobuf:"varint,2,opt,name=price,proto3" json:"price,omitempty"`
	StartTime string `protobuf:"bytes,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
}

func (x *ScreenCinemaInfo) Reset() {
	*x = ScreenCinemaInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_cinema_rpc_pb_cinema_cinema_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScreenCinemaInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScreenCinemaInfo) ProtoMessage() {}

func (x *ScreenCinemaInfo) ProtoReflect() protoreflect.Message {
	mi := &file_services_cinema_rpc_pb_cinema_cinema_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScreenCinemaInfo.ProtoReflect.Descriptor instead.
func (*ScreenCinemaInfo) Descriptor() ([]byte, []int) {
	return file_services_cinema_rpc_pb_cinema_cinema_proto_rawDescGZIP(), []int{10}
}

func (x *ScreenCinemaInfo) GetCinemaId() int32 {
	if x != nil {
		return x.CinemaId
	}
	return 0
}

func (x *ScreenCinemaInfo) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *ScreenCinemaInfo) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

type ScreenCinemaInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*ScreenCinemaInfo `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ScreenCinemaInfoResponse) Reset() {
	*x = ScreenCinemaInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_cinema_rpc_pb_cinema_cinema_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScreenCinemaInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScreenCinemaInfoResponse) ProtoMessage() {}

func (x *ScreenCinemaInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_cinema_rpc_pb_cinema_cinema_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScreenCinemaInfoResponse.ProtoReflect.Descriptor instead.
func (*ScreenCinemaInfoResponse) Descriptor() ([]byte, []int) {
	return file_services_cinema_rpc_pb_cinema_cinema_proto_rawDescGZIP(), []int{11}
}

func (x *ScreenCinemaInfoResponse) GetData() []*ScreenCinemaInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_services_cinema_rpc_pb_cinema_cinema_proto protoreflect.FileDescriptor

var file_services_cinema_rpc_pb_cinema_cinema_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x69, 0x6e, 0x65, 0x6d,
	0x61, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x69, 0x6e, 0x65, 0x6d, 0x61, 0x2f,
	0x63, 0x69, 0x6e, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x69,
	0x6e, 0x65, 0x6d, 0x61, 0x22, 0x7b, 0x0a, 0x0f, 0x46, 0x69, 0x6c, 0x6d, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x61, 0x74, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x75, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x6d, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x66, 0x69, 0x6c, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x5f, 0x70, 0x69, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x50, 0x69, 0x63, 0x22, 0x7c, 0x0a, 0x10, 0x46, 0x69, 0x6c, 0x6d,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x69, 0x6e,
	0x65, 0x6d, 0x61, 0x2e, 0x46, 0x69, 0x6c, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x2c, 0x0a, 0x11, 0x46, 0x69, 0x6c, 0x6d, 0x44, 0x61,
	0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x66,
	0x69, 0x6c, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x66, 0x69,
	0x6c, 0x6d, 0x49, 0x64, 0x22, 0x94, 0x01, 0x0a, 0x0e, 0x46, 0x69, 0x6c, 0x6d, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x6d, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x6d, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x63, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x6d, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x6d, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x70, 0x69, 0x63, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x50, 0x69, 0x63, 0x12, 0x1b,
	0x0a, 0x09, 0x66, 0x69, 0x6c, 0x6d, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x6d, 0x44, 0x65, 0x73, 0x63, 0x22, 0x40, 0x0a, 0x12, 0x46,
	0x69, 0x6c, 0x6d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x63, 0x69, 0x6e, 0x65, 0x6d, 0x61, 0x2e, 0x46, 0x69, 0x6c, 0x6d, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x4d, 0x0a,
	0x11, 0x43, 0x69, 0x6e, 0x65, 0x6d, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x61, 0x72, 0x65, 0x61, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x8a, 0x01, 0x0a,
	0x0a, 0x43, 0x69, 0x6e, 0x65, 0x6d, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x63,
	0x69, 0x6e, 0x65, 0x6d, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x63, 0x69, 0x6e, 0x65, 0x6d, 0x61, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x69, 0x6e, 0x65,
	0x6d, 0x61, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63,
	0x69, 0x6e, 0x65, 0x6d, 0x61, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x3c, 0x0a, 0x12, 0x43, 0x69, 0x6e,
	0x65, 0x6d, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x26, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x63, 0x69, 0x6e, 0x65, 0x6d, 0x61, 0x2e, 0x43, 0x69, 0x6e, 0x65, 0x6d, 0x61, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x68, 0x0a, 0x17, 0x53, 0x63, 0x72, 0x65, 0x65,
	0x6e, 0x43, 0x69, 0x6e, 0x65, 0x6d, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x6d, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x74,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x69, 0x6e, 0x65, 0x6d, 0x61, 0x5f, 0x69, 0x64, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52, 0x09, 0x63, 0x69, 0x6e, 0x65, 0x6d, 0x61, 0x49, 0x64,
	0x73, 0x22, 0x64, 0x0a, 0x10, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x43, 0x69, 0x6e, 0x65, 0x6d,
	0x61, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x69, 0x6e, 0x65, 0x6d, 0x61, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x69, 0x6e, 0x65, 0x6d, 0x61,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x48, 0x0a, 0x18, 0x53, 0x63, 0x72, 0x65, 0x65,
	0x6e, 0x43, 0x69, 0x6e, 0x65, 0x6d, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x63, 0x69, 0x6e, 0x65, 0x6d, 0x61, 0x2e, 0x53, 0x63, 0x72, 0x65, 0x65,
	0x6e, 0x43, 0x69, 0x6e, 0x65, 0x6d, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x32, 0xa8, 0x02, 0x0a, 0x06, 0x43, 0x69, 0x6e, 0x65, 0x6d, 0x61, 0x12, 0x3d, 0x0a, 0x08,
	0x46, 0x69, 0x6c, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x17, 0x2e, 0x63, 0x69, 0x6e, 0x65, 0x6d,
	0x61, 0x2e, 0x46, 0x69, 0x6c, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x63, 0x69, 0x6e, 0x65, 0x6d, 0x61, 0x2e, 0x46, 0x69, 0x6c, 0x6d, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0a, 0x46,
	0x69, 0x6c, 0x6d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x19, 0x2e, 0x63, 0x69, 0x6e, 0x65,
	0x6d, 0x61, 0x2e, 0x46, 0x69, 0x6c, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x69, 0x6e, 0x65, 0x6d, 0x61, 0x2e, 0x46, 0x69,
	0x6c, 0x6d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x43, 0x0a, 0x0a, 0x43, 0x69, 0x6e, 0x65, 0x6d, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19,
	0x2e, 0x63, 0x69, 0x6e, 0x65, 0x6d, 0x61, 0x2e, 0x43, 0x69, 0x6e, 0x65, 0x6d, 0x61, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x69, 0x6e, 0x65,
	0x6d, 0x61, 0x2e, 0x43, 0x69, 0x6e, 0x65, 0x6d, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x10, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x43,
	0x69, 0x6e, 0x65, 0x6d, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x2e, 0x63, 0x69, 0x6e, 0x65,
	0x6d, 0x61, 0x2e, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x43, 0x69, 0x6e, 0x65, 0x6d, 0x61, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x63, 0x69, 0x6e,
	0x65, 0x6d, 0x61, 0x2e, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x43, 0x69, 0x6e, 0x65, 0x6d, 0x61,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x63, 0x69, 0x6e, 0x65, 0x6d, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_cinema_rpc_pb_cinema_cinema_proto_rawDescOnce sync.Once
	file_services_cinema_rpc_pb_cinema_cinema_proto_rawDescData = file_services_cinema_rpc_pb_cinema_cinema_proto_rawDesc
)

func file_services_cinema_rpc_pb_cinema_cinema_proto_rawDescGZIP() []byte {
	file_services_cinema_rpc_pb_cinema_cinema_proto_rawDescOnce.Do(func() {
		file_services_cinema_rpc_pb_cinema_cinema_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_cinema_rpc_pb_cinema_cinema_proto_rawDescData)
	})
	return file_services_cinema_rpc_pb_cinema_cinema_proto_rawDescData
}

var file_services_cinema_rpc_pb_cinema_cinema_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_services_cinema_rpc_pb_cinema_cinema_proto_goTypes = []interface{}{
	(*FilmListRequest)(nil),          // 0: cinema.FilmListRequest
	(*FilmListInfo)(nil),             // 1: cinema.FilmListInfo
	(*FilmListResponse)(nil),         // 2: cinema.FilmListResponse
	(*FilmDatailRequest)(nil),        // 3: cinema.FilmDatailRequest
	(*FilmDetailInfo)(nil),           // 4: cinema.FilmDetailInfo
	(*FilmDetailResponse)(nil),       // 5: cinema.FilmDetailResponse
	(*CinemaInfoRequest)(nil),        // 6: cinema.CinemaInfoRequest
	(*CinemaInfo)(nil),               // 7: cinema.CinemaInfo
	(*CinemaInfoResponse)(nil),       // 8: cinema.CinemaInfoResponse
	(*ScreenCinemaInfoRequest)(nil),  // 9: cinema.ScreenCinemaInfoRequest
	(*ScreenCinemaInfo)(nil),         // 10: cinema.ScreenCinemaInfo
	(*ScreenCinemaInfoResponse)(nil), // 11: cinema.ScreenCinemaInfoResponse
}
var file_services_cinema_rpc_pb_cinema_cinema_proto_depIdxs = []int32{
	1,  // 0: cinema.FilmListResponse.data:type_name -> cinema.FilmListInfo
	4,  // 1: cinema.FilmDetailResponse.data:type_name -> cinema.FilmDetailInfo
	7,  // 2: cinema.CinemaInfoResponse.data:type_name -> cinema.CinemaInfo
	10, // 3: cinema.ScreenCinemaInfoResponse.data:type_name -> cinema.ScreenCinemaInfo
	0,  // 4: cinema.Cinema.FilmList:input_type -> cinema.FilmListRequest
	3,  // 5: cinema.Cinema.FilmDetail:input_type -> cinema.FilmDatailRequest
	6,  // 6: cinema.Cinema.CinemaInfo:input_type -> cinema.CinemaInfoRequest
	9,  // 7: cinema.Cinema.ScreenCinemaInfo:input_type -> cinema.ScreenCinemaInfoRequest
	2,  // 8: cinema.Cinema.FilmList:output_type -> cinema.FilmListResponse
	5,  // 9: cinema.Cinema.FilmDetail:output_type -> cinema.FilmDetailResponse
	8,  // 10: cinema.Cinema.CinemaInfo:output_type -> cinema.CinemaInfoResponse
	11, // 11: cinema.Cinema.ScreenCinemaInfo:output_type -> cinema.ScreenCinemaInfoResponse
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_services_cinema_rpc_pb_cinema_cinema_proto_init() }
func file_services_cinema_rpc_pb_cinema_cinema_proto_init() {
	if File_services_cinema_rpc_pb_cinema_cinema_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_cinema_rpc_pb_cinema_cinema_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilmListRequest); i {
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
		file_services_cinema_rpc_pb_cinema_cinema_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilmListInfo); i {
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
		file_services_cinema_rpc_pb_cinema_cinema_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilmListResponse); i {
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
		file_services_cinema_rpc_pb_cinema_cinema_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilmDatailRequest); i {
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
		file_services_cinema_rpc_pb_cinema_cinema_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilmDetailInfo); i {
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
		file_services_cinema_rpc_pb_cinema_cinema_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilmDetailResponse); i {
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
		file_services_cinema_rpc_pb_cinema_cinema_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CinemaInfoRequest); i {
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
		file_services_cinema_rpc_pb_cinema_cinema_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CinemaInfo); i {
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
		file_services_cinema_rpc_pb_cinema_cinema_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CinemaInfoResponse); i {
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
		file_services_cinema_rpc_pb_cinema_cinema_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScreenCinemaInfoRequest); i {
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
		file_services_cinema_rpc_pb_cinema_cinema_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScreenCinemaInfo); i {
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
		file_services_cinema_rpc_pb_cinema_cinema_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScreenCinemaInfoResponse); i {
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
			RawDescriptor: file_services_cinema_rpc_pb_cinema_cinema_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_cinema_rpc_pb_cinema_cinema_proto_goTypes,
		DependencyIndexes: file_services_cinema_rpc_pb_cinema_cinema_proto_depIdxs,
		MessageInfos:      file_services_cinema_rpc_pb_cinema_cinema_proto_msgTypes,
	}.Build()
	File_services_cinema_rpc_pb_cinema_cinema_proto = out.File
	file_services_cinema_rpc_pb_cinema_cinema_proto_rawDesc = nil
	file_services_cinema_rpc_pb_cinema_cinema_proto_goTypes = nil
	file_services_cinema_rpc_pb_cinema_cinema_proto_depIdxs = nil
}
