// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.20.3
// source: gc/ssl_gc_api.proto

package gc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Message format that is pushed from the GC to the client
type Output struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The current match state
	MatchState *State `protobuf:"bytes,1,opt,name=match_state,json=matchState" json:"match_state,omitempty"`
	// The current GC state
	GcState *GcState `protobuf:"bytes,2,opt,name=gc_state,json=gcState" json:"gc_state,omitempty"`
	// The protocol
	Protocol *Protocol `protobuf:"bytes,3,opt,name=protocol" json:"protocol,omitempty"`
	// The engine config
	Config *Config `protobuf:"bytes,4,opt,name=config" json:"config,omitempty"`
}

func (x *Output) Reset() {
	*x = Output{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gc_ssl_gc_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Output) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Output) ProtoMessage() {}

func (x *Output) ProtoReflect() protoreflect.Message {
	mi := &file_gc_ssl_gc_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Output.ProtoReflect.Descriptor instead.
func (*Output) Descriptor() ([]byte, []int) {
	return file_gc_ssl_gc_api_proto_rawDescGZIP(), []int{0}
}

func (x *Output) GetMatchState() *State {
	if x != nil {
		return x.MatchState
	}
	return nil
}

func (x *Output) GetGcState() *GcState {
	if x != nil {
		return x.GcState
	}
	return nil
}

func (x *Output) GetProtocol() *Protocol {
	if x != nil {
		return x.Protocol
	}
	return nil
}

func (x *Output) GetConfig() *Config {
	if x != nil {
		return x.Config
	}
	return nil
}

// The game protocol
type Protocol struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Is this a delta only?
	// Entries that were already sent are not sent again, because the protocol is immutable anyway.
	// But if the game is reset, the whole protocol must be replaced. That's what this flag is for.
	Delta *bool `protobuf:"varint,1,opt,name=delta" json:"delta,omitempty"`
	// The (delta) list of entries
	Entry []*ProtocolEntry `protobuf:"bytes,2,rep,name=entry" json:"entry,omitempty"`
}

func (x *Protocol) Reset() {
	*x = Protocol{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gc_ssl_gc_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Protocol) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Protocol) ProtoMessage() {}

func (x *Protocol) ProtoReflect() protoreflect.Message {
	mi := &file_gc_ssl_gc_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Protocol.ProtoReflect.Descriptor instead.
func (*Protocol) Descriptor() ([]byte, []int) {
	return file_gc_ssl_gc_api_proto_rawDescGZIP(), []int{1}
}

func (x *Protocol) GetDelta() bool {
	if x != nil && x.Delta != nil {
		return *x.Delta
	}
	return false
}

func (x *Protocol) GetEntry() []*ProtocolEntry {
	if x != nil {
		return x.Entry
	}
	return nil
}

// A protocol entry of a change
type ProtocolEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Id of the entry
	Id *int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// The change that was made
	Change *Change `protobuf:"bytes,2,opt,name=change" json:"change,omitempty"`
	// The match time elapsed when this change was made
	MatchTimeElapsed *durationpb.Duration `protobuf:"bytes,3,opt,name=match_time_elapsed,json=matchTimeElapsed" json:"match_time_elapsed,omitempty"`
	// The stage time elapsed when this change was made
	StageTimeElapsed *durationpb.Duration `protobuf:"bytes,4,opt,name=stage_time_elapsed,json=stageTimeElapsed" json:"stage_time_elapsed,omitempty"`
}

func (x *ProtocolEntry) Reset() {
	*x = ProtocolEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gc_ssl_gc_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtocolEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtocolEntry) ProtoMessage() {}

func (x *ProtocolEntry) ProtoReflect() protoreflect.Message {
	mi := &file_gc_ssl_gc_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtocolEntry.ProtoReflect.Descriptor instead.
func (*ProtocolEntry) Descriptor() ([]byte, []int) {
	return file_gc_ssl_gc_api_proto_rawDescGZIP(), []int{2}
}

func (x *ProtocolEntry) GetId() int32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *ProtocolEntry) GetChange() *Change {
	if x != nil {
		return x.Change
	}
	return nil
}

func (x *ProtocolEntry) GetMatchTimeElapsed() *durationpb.Duration {
	if x != nil {
		return x.MatchTimeElapsed
	}
	return nil
}

func (x *ProtocolEntry) GetStageTimeElapsed() *durationpb.Duration {
	if x != nil {
		return x.StageTimeElapsed
	}
	return nil
}

// Message format that can be send from the client to the GC
type Input struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A change to be enqueued into the GC engine
	Change *Change `protobuf:"bytes,1,opt,name=change" json:"change,omitempty"`
	// Reset the match
	ResetMatch *bool `protobuf:"varint,2,opt,name=reset_match,json=resetMatch" json:"reset_match,omitempty"`
	// An updated config delta
	ConfigDelta *Config `protobuf:"bytes,3,opt,name=config_delta,json=configDelta" json:"config_delta,omitempty"`
	// Continue with action
	ContinueAction *ContinueAction `protobuf:"bytes,4,opt,name=continue_action,json=continueAction" json:"continue_action,omitempty"`
}

func (x *Input) Reset() {
	*x = Input{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gc_ssl_gc_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Input) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Input) ProtoMessage() {}

func (x *Input) ProtoReflect() protoreflect.Message {
	mi := &file_gc_ssl_gc_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Input.ProtoReflect.Descriptor instead.
func (*Input) Descriptor() ([]byte, []int) {
	return file_gc_ssl_gc_api_proto_rawDescGZIP(), []int{3}
}

func (x *Input) GetChange() *Change {
	if x != nil {
		return x.Change
	}
	return nil
}

func (x *Input) GetResetMatch() bool {
	if x != nil && x.ResetMatch != nil {
		return *x.ResetMatch
	}
	return false
}

func (x *Input) GetConfigDelta() *Config {
	if x != nil {
		return x.ConfigDelta
	}
	return nil
}

func (x *Input) GetContinueAction() *ContinueAction {
	if x != nil {
		return x.ContinueAction
	}
	return nil
}

var File_gc_ssl_gc_api_proto protoreflect.FileDescriptor

var file_gc_ssl_gc_api_proto_rawDesc = []byte{
	0x0a, 0x13, 0x67, 0x63, 0x2f, 0x73, 0x73, 0x6c, 0x5f, 0x67, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x67, 0x63, 0x2f, 0x73, 0x73, 0x6c, 0x5f, 0x67, 0x63,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x67, 0x63,
	0x2f, 0x73, 0x73, 0x6c, 0x5f, 0x67, 0x63, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x67, 0x63, 0x2f, 0x73, 0x73, 0x6c, 0x5f, 0x67, 0x63, 0x5f,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x63,
	0x2f, 0x73, 0x73, 0x6c, 0x5f, 0x67, 0x63, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x01, 0x0a, 0x06,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x27, 0x0a, 0x0b, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x0a, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x23, 0x0a, 0x08, 0x67, 0x63, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x08, 0x2e, 0x47, 0x63, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x07, 0x67, 0x63, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x25, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x1f, 0x0a, 0x06, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x46, 0x0a, 0x08,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x6c, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x12, 0x24,
	0x0a, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x65,
	0x6e, 0x74, 0x72, 0x79, 0x22, 0xd2, 0x01, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x06, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52,
	0x06, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x47, 0x0a, 0x12, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x65, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x45, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x64,
	0x12, 0x47, 0x0a, 0x12, 0x73, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x65,
	0x6c, 0x61, 0x70, 0x73, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x73, 0x74, 0x61, 0x67, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x45, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x64, 0x22, 0xaf, 0x01, 0x0a, 0x05, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x12, 0x1f, 0x0a, 0x06, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x06, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x65, 0x74,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x2a, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f,
	0x64, 0x65, 0x6c, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x44, 0x65, 0x6c, 0x74,
	0x61, 0x12, 0x38, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x65, 0x5f, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x43, 0x6f, 0x6e,
	0x74, 0x69, 0x6e, 0x75, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x63, 0x6f, 0x6e,
	0x74, 0x69, 0x6e, 0x75, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x25, 0x5a, 0x23, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x69, 0x55, 0x2d, 0x53, 0x65,
	0x65, 0x47, 0x6f, 0x61, 0x6c, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x67, 0x6f, 0x2f,
	0x67, 0x63,
}

var (
	file_gc_ssl_gc_api_proto_rawDescOnce sync.Once
	file_gc_ssl_gc_api_proto_rawDescData = file_gc_ssl_gc_api_proto_rawDesc
)

func file_gc_ssl_gc_api_proto_rawDescGZIP() []byte {
	file_gc_ssl_gc_api_proto_rawDescOnce.Do(func() {
		file_gc_ssl_gc_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_gc_ssl_gc_api_proto_rawDescData)
	})
	return file_gc_ssl_gc_api_proto_rawDescData
}

var file_gc_ssl_gc_api_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_gc_ssl_gc_api_proto_goTypes = []any{
	(*Output)(nil),              // 0: Output
	(*Protocol)(nil),            // 1: Protocol
	(*ProtocolEntry)(nil),       // 2: ProtocolEntry
	(*Input)(nil),               // 3: Input
	(*State)(nil),               // 4: State
	(*GcState)(nil),             // 5: GcState
	(*Config)(nil),              // 6: Config
	(*Change)(nil),              // 7: Change
	(*durationpb.Duration)(nil), // 8: google.protobuf.Duration
	(*ContinueAction)(nil),      // 9: ContinueAction
}
var file_gc_ssl_gc_api_proto_depIdxs = []int32{
	4,  // 0: Output.match_state:type_name -> State
	5,  // 1: Output.gc_state:type_name -> GcState
	1,  // 2: Output.protocol:type_name -> Protocol
	6,  // 3: Output.config:type_name -> Config
	2,  // 4: Protocol.entry:type_name -> ProtocolEntry
	7,  // 5: ProtocolEntry.change:type_name -> Change
	8,  // 6: ProtocolEntry.match_time_elapsed:type_name -> google.protobuf.Duration
	8,  // 7: ProtocolEntry.stage_time_elapsed:type_name -> google.protobuf.Duration
	7,  // 8: Input.change:type_name -> Change
	6,  // 9: Input.config_delta:type_name -> Config
	9,  // 10: Input.continue_action:type_name -> ContinueAction
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_gc_ssl_gc_api_proto_init() }
func file_gc_ssl_gc_api_proto_init() {
	if File_gc_ssl_gc_api_proto != nil {
		return
	}
	file_gc_ssl_gc_state_proto_init()
	file_gc_ssl_gc_change_proto_init()
	file_gc_ssl_gc_engine_proto_init()
	file_gc_ssl_gc_engine_config_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_gc_ssl_gc_api_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Output); i {
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
		file_gc_ssl_gc_api_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Protocol); i {
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
		file_gc_ssl_gc_api_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ProtocolEntry); i {
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
		file_gc_ssl_gc_api_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Input); i {
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
			RawDescriptor: file_gc_ssl_gc_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gc_ssl_gc_api_proto_goTypes,
		DependencyIndexes: file_gc_ssl_gc_api_proto_depIdxs,
		MessageInfos:      file_gc_ssl_gc_api_proto_msgTypes,
	}.Build()
	File_gc_ssl_gc_api_proto = out.File
	file_gc_ssl_gc_api_proto_rawDesc = nil
	file_gc_ssl_gc_api_proto_goTypes = nil
	file_gc_ssl_gc_api_proto_depIdxs = nil
}
