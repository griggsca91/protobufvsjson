// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: shared.proto

package protobuf

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

type Identifiers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CardKingdomId          string `protobuf:"bytes,1,opt,name=cardKingdomId,proto3" json:"cardKingdomId,omitempty"`
	McmId                  string `protobuf:"bytes,2,opt,name=mcmId,proto3" json:"mcmId,omitempty"`
	McmMetaId              string `protobuf:"bytes,3,opt,name=mcmMetaId,proto3" json:"mcmMetaId,omitempty"`
	MtgjsonFoilVersionId   string `protobuf:"bytes,4,opt,name=mtgjsonFoilVersionId,proto3" json:"mtgjsonFoilVersionId,omitempty"`
	MtgjsonV4Id            string `protobuf:"bytes,5,opt,name=mtgjsonV4Id,proto3" json:"mtgjsonV4Id,omitempty"`
	MtgoFoilId             string `protobuf:"bytes,6,opt,name=mtgoFoilId,proto3" json:"mtgoFoilId,omitempty"`
	MtgoId                 string `protobuf:"bytes,7,opt,name=mtgoId,proto3" json:"mtgoId,omitempty"`
	MultiverseId           string `protobuf:"bytes,8,opt,name=multiverseId,proto3" json:"multiverseId,omitempty"`
	ScryfallId             string `protobuf:"bytes,9,opt,name=scryfallId,proto3" json:"scryfallId,omitempty"`
	ScryfallIllustrationId string `protobuf:"bytes,10,opt,name=scryfallIllustrationId,proto3" json:"scryfallIllustrationId,omitempty"`
	ScryfallOracleId       string `protobuf:"bytes,11,opt,name=scryfallOracleId,proto3" json:"scryfallOracleId,omitempty"`
	TcgplayerProductId     string `protobuf:"bytes,12,opt,name=tcgplayerProductId,proto3" json:"tcgplayerProductId,omitempty"`
}

func (x *Identifiers) Reset() {
	*x = Identifiers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Identifiers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Identifiers) ProtoMessage() {}

func (x *Identifiers) ProtoReflect() protoreflect.Message {
	mi := &file_shared_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Identifiers.ProtoReflect.Descriptor instead.
func (*Identifiers) Descriptor() ([]byte, []int) {
	return file_shared_proto_rawDescGZIP(), []int{0}
}

func (x *Identifiers) GetCardKingdomId() string {
	if x != nil {
		return x.CardKingdomId
	}
	return ""
}

func (x *Identifiers) GetMcmId() string {
	if x != nil {
		return x.McmId
	}
	return ""
}

func (x *Identifiers) GetMcmMetaId() string {
	if x != nil {
		return x.McmMetaId
	}
	return ""
}

func (x *Identifiers) GetMtgjsonFoilVersionId() string {
	if x != nil {
		return x.MtgjsonFoilVersionId
	}
	return ""
}

func (x *Identifiers) GetMtgjsonV4Id() string {
	if x != nil {
		return x.MtgjsonV4Id
	}
	return ""
}

func (x *Identifiers) GetMtgoFoilId() string {
	if x != nil {
		return x.MtgoFoilId
	}
	return ""
}

func (x *Identifiers) GetMtgoId() string {
	if x != nil {
		return x.MtgoId
	}
	return ""
}

func (x *Identifiers) GetMultiverseId() string {
	if x != nil {
		return x.MultiverseId
	}
	return ""
}

func (x *Identifiers) GetScryfallId() string {
	if x != nil {
		return x.ScryfallId
	}
	return ""
}

func (x *Identifiers) GetScryfallIllustrationId() string {
	if x != nil {
		return x.ScryfallIllustrationId
	}
	return ""
}

func (x *Identifiers) GetScryfallOracleId() string {
	if x != nil {
		return x.ScryfallOracleId
	}
	return ""
}

func (x *Identifiers) GetTcgplayerProductId() string {
	if x != nil {
		return x.TcgplayerProductId
	}
	return ""
}

type Foreigndata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FlavorText   string `protobuf:"bytes,1,opt,name=flavorText,proto3" json:"flavorText,omitempty"`
	Language     string `protobuf:"bytes,2,opt,name=language,proto3" json:"language,omitempty"`
	MultiverseId uint32 `protobuf:"varint,3,opt,name=multiverseId,proto3" json:"multiverseId,omitempty"`
	Name         string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Text         string `protobuf:"bytes,5,opt,name=text,proto3" json:"text,omitempty"`
	Type         string `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *Foreigndata) Reset() {
	*x = Foreigndata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Foreigndata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Foreigndata) ProtoMessage() {}

func (x *Foreigndata) ProtoReflect() protoreflect.Message {
	mi := &file_shared_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Foreigndata.ProtoReflect.Descriptor instead.
func (*Foreigndata) Descriptor() ([]byte, []int) {
	return file_shared_proto_rawDescGZIP(), []int{1}
}

func (x *Foreigndata) GetFlavorText() string {
	if x != nil {
		return x.FlavorText
	}
	return ""
}

func (x *Foreigndata) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *Foreigndata) GetMultiverseId() uint32 {
	if x != nil {
		return x.MultiverseId
	}
	return 0
}

func (x *Foreigndata) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Foreigndata) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Foreigndata) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type Purchaseurls struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CardKingdom string `protobuf:"bytes,1,opt,name=cardKingdom,proto3" json:"cardKingdom,omitempty"`
	Cardmarket  string `protobuf:"bytes,2,opt,name=cardmarket,proto3" json:"cardmarket,omitempty"`
	Tcgplayer   string `protobuf:"bytes,3,opt,name=tcgplayer,proto3" json:"tcgplayer,omitempty"`
}

func (x *Purchaseurls) Reset() {
	*x = Purchaseurls{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Purchaseurls) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Purchaseurls) ProtoMessage() {}

func (x *Purchaseurls) ProtoReflect() protoreflect.Message {
	mi := &file_shared_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Purchaseurls.ProtoReflect.Descriptor instead.
func (*Purchaseurls) Descriptor() ([]byte, []int) {
	return file_shared_proto_rawDescGZIP(), []int{2}
}

func (x *Purchaseurls) GetCardKingdom() string {
	if x != nil {
		return x.CardKingdom
	}
	return ""
}

func (x *Purchaseurls) GetCardmarket() string {
	if x != nil {
		return x.Cardmarket
	}
	return ""
}

func (x *Purchaseurls) GetTcgplayer() string {
	if x != nil {
		return x.Tcgplayer
	}
	return ""
}

type Sourceproducts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nonfoil []string `protobuf:"bytes,1,rep,name=nonfoil,proto3" json:"nonfoil,omitempty"`
}

func (x *Sourceproducts) Reset() {
	*x = Sourceproducts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sourceproducts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sourceproducts) ProtoMessage() {}

func (x *Sourceproducts) ProtoReflect() protoreflect.Message {
	mi := &file_shared_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sourceproducts.ProtoReflect.Descriptor instead.
func (*Sourceproducts) Descriptor() ([]byte, []int) {
	return file_shared_proto_rawDescGZIP(), []int{3}
}

func (x *Sourceproducts) GetNonfoil() []string {
	if x != nil {
		return x.Nonfoil
	}
	return nil
}

var File_shared_proto protoreflect.FileDescriptor

var file_shared_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x22, 0xcd, 0x03, 0x0a, 0x0b, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x61, 0x72, 0x64,
	0x4b, 0x69, 0x6e, 0x67, 0x64, 0x6f, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x63, 0x61, 0x72, 0x64, 0x4b, 0x69, 0x6e, 0x67, 0x64, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x6d, 0x63, 0x6d, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d,
	0x63, 0x6d, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x63, 0x6d, 0x4d, 0x65, 0x74, 0x61, 0x49,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x63, 0x6d, 0x4d, 0x65, 0x74, 0x61,
	0x49, 0x64, 0x12, 0x32, 0x0a, 0x14, 0x6d, 0x74, 0x67, 0x6a, 0x73, 0x6f, 0x6e, 0x46, 0x6f, 0x69,
	0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x14, 0x6d, 0x74, 0x67, 0x6a, 0x73, 0x6f, 0x6e, 0x46, 0x6f, 0x69, 0x6c, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x74, 0x67, 0x6a, 0x73, 0x6f,
	0x6e, 0x56, 0x34, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x74, 0x67,
	0x6a, 0x73, 0x6f, 0x6e, 0x56, 0x34, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x74, 0x67, 0x6f,
	0x46, 0x6f, 0x69, 0x6c, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x74,
	0x67, 0x6f, 0x46, 0x6f, 0x69, 0x6c, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x74, 0x67, 0x6f,
	0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x74, 0x67, 0x6f, 0x49, 0x64,
	0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x49, 0x64,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x76, 0x65, 0x72,
	0x73, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x63, 0x72, 0x79, 0x66, 0x61, 0x6c, 0x6c,
	0x49, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x63, 0x72, 0x79, 0x66, 0x61,
	0x6c, 0x6c, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x16, 0x73, 0x63, 0x72, 0x79, 0x66, 0x61, 0x6c, 0x6c,
	0x49, 0x6c, 0x6c, 0x75, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x73, 0x63, 0x72, 0x79, 0x66, 0x61, 0x6c, 0x6c, 0x49, 0x6c,
	0x6c, 0x75, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x10,
	0x73, 0x63, 0x72, 0x79, 0x66, 0x61, 0x6c, 0x6c, 0x4f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x49, 0x64,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x63, 0x72, 0x79, 0x66, 0x61, 0x6c, 0x6c,
	0x4f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x12, 0x74, 0x63, 0x67, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x74, 0x63, 0x67, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x22, 0xa9, 0x01, 0x0a, 0x0b, 0x46, 0x6f, 0x72,
	0x65, 0x69, 0x67, 0x6e, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x6c, 0x61, 0x76,
	0x6f, 0x72, 0x54, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x6c,
	0x61, 0x76, 0x6f, 0x72, 0x54, 0x65, 0x78, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x76, 0x65, 0x72,
	0x73, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x6d, 0x75, 0x6c, 0x74,
	0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x22, 0x6e, 0x0a, 0x0c, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65,
	0x75, 0x72, 0x6c, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x61, 0x72, 0x64, 0x4b, 0x69, 0x6e, 0x67,
	0x64, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x61, 0x72, 0x64, 0x4b,
	0x69, 0x6e, 0x67, 0x64, 0x6f, 0x6d, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x72, 0x64, 0x6d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x72, 0x64,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x63, 0x67, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x63, 0x67, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x22, 0x2a, 0x0a, 0x0e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x6f, 0x6e, 0x66, 0x6f, 0x69,
	0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x6f, 0x6e, 0x66, 0x6f, 0x69, 0x6c,
	0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shared_proto_rawDescOnce sync.Once
	file_shared_proto_rawDescData = file_shared_proto_rawDesc
)

func file_shared_proto_rawDescGZIP() []byte {
	file_shared_proto_rawDescOnce.Do(func() {
		file_shared_proto_rawDescData = protoimpl.X.CompressGZIP(file_shared_proto_rawDescData)
	})
	return file_shared_proto_rawDescData
}

var file_shared_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_shared_proto_goTypes = []interface{}{
	(*Identifiers)(nil),    // 0: protobuf.Identifiers
	(*Foreigndata)(nil),    // 1: protobuf.Foreigndata
	(*Purchaseurls)(nil),   // 2: protobuf.Purchaseurls
	(*Sourceproducts)(nil), // 3: protobuf.Sourceproducts
}
var file_shared_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_shared_proto_init() }
func file_shared_proto_init() {
	if File_shared_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shared_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Identifiers); i {
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
		file_shared_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Foreigndata); i {
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
		file_shared_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Purchaseurls); i {
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
		file_shared_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sourceproducts); i {
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
			RawDescriptor: file_shared_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_shared_proto_goTypes,
		DependencyIndexes: file_shared_proto_depIdxs,
		MessageInfos:      file_shared_proto_msgTypes,
	}.Build()
	File_shared_proto = out.File
	file_shared_proto_rawDesc = nil
	file_shared_proto_goTypes = nil
	file_shared_proto_depIdxs = nil
}
