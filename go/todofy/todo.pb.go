// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.12
// source: proto/todofy/todo.proto

package todofy

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

type TodoApp int32

const (
	TodoApp_TODO_APP_UNSPECIFIED TodoApp = 0
	TodoApp_TODO_APP_DIDA365     TodoApp = 1
	TodoApp_TODO_APP_TICKTICK    TodoApp = 2
	TodoApp_TODO_APP_TODOIST     TodoApp = 3
)

// Enum value maps for TodoApp.
var (
	TodoApp_name = map[int32]string{
		0: "TODO_APP_UNSPECIFIED",
		1: "TODO_APP_DIDA365",
		2: "TODO_APP_TICKTICK",
		3: "TODO_APP_TODOIST",
	}
	TodoApp_value = map[string]int32{
		"TODO_APP_UNSPECIFIED": 0,
		"TODO_APP_DIDA365":     1,
		"TODO_APP_TICKTICK":    2,
		"TODO_APP_TODOIST":     3,
	}
)

func (x TodoApp) Enum() *TodoApp {
	p := new(TodoApp)
	*p = x
	return p
}

func (x TodoApp) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TodoApp) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_todofy_todo_proto_enumTypes[0].Descriptor()
}

func (TodoApp) Type() protoreflect.EnumType {
	return &file_proto_todofy_todo_proto_enumTypes[0]
}

func (x TodoApp) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TodoApp.Descriptor instead.
func (TodoApp) EnumDescriptor() ([]byte, []int) {
	return file_proto_todofy_todo_proto_rawDescGZIP(), []int{0}
}

type PopullateTodoMethod int32

const (
	PopullateTodoMethod_POPULLATE_TODO_METHOD_UNSPECIFIED PopullateTodoMethod = 0
	PopullateTodoMethod_POPULLATE_TODO_METHOD_MAILJET     PopullateTodoMethod = 1
	PopullateTodoMethod_POPULLATE_TODO_METHOD_API         PopullateTodoMethod = 2
)

// Enum value maps for PopullateTodoMethod.
var (
	PopullateTodoMethod_name = map[int32]string{
		0: "POPULLATE_TODO_METHOD_UNSPECIFIED",
		1: "POPULLATE_TODO_METHOD_MAILJET",
		2: "POPULLATE_TODO_METHOD_API",
	}
	PopullateTodoMethod_value = map[string]int32{
		"POPULLATE_TODO_METHOD_UNSPECIFIED": 0,
		"POPULLATE_TODO_METHOD_MAILJET":     1,
		"POPULLATE_TODO_METHOD_API":         2,
	}
)

func (x PopullateTodoMethod) Enum() *PopullateTodoMethod {
	p := new(PopullateTodoMethod)
	*p = x
	return p
}

func (x PopullateTodoMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PopullateTodoMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_todofy_todo_proto_enumTypes[1].Descriptor()
}

func (PopullateTodoMethod) Type() protoreflect.EnumType {
	return &file_proto_todofy_todo_proto_enumTypes[1]
}

func (x PopullateTodoMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PopullateTodoMethod.Descriptor instead.
func (PopullateTodoMethod) EnumDescriptor() ([]byte, []int) {
	return file_proto_todofy_todo_proto_rawDescGZIP(), []int{1}
}

type TodoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of todo app to use.
	App TodoApp `protobuf:"varint,1,opt,name=app,proto3,enum=todofy.TodoApp" json:"app,omitempty"`
	// The method to populate the todo.
	Method PopullateTodoMethod `protobuf:"varint,2,opt,name=method,proto3,enum=todofy.PopullateTodoMethod" json:"method,omitempty"`
	// The subject of the todo.
	Subject string `protobuf:"bytes,3,opt,name=subject,proto3" json:"subject,omitempty"`
	// The body of the todo.
	Body string `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	// The tags of the todo.
	Tags []string `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
	// The category of the todo.
	Category string `protobuf:"bytes,6,opt,name=category,proto3" json:"category,omitempty"`
	// The todo task is collected from.
	From string `protobuf:"bytes,7,opt,name=from,proto3" json:"from,omitempty"`
	// Optional. Override email address this todo is sent to.
	To string `protobuf:"bytes,8,opt,name=to,proto3" json:"to,omitempty"`
	// Optional. Override the email address name this todo is sent to.
	ToName string `protobuf:"bytes,9,opt,name=to_name,json=toName,proto3" json:"to_name,omitempty"`
}

func (x *TodoRequest) Reset() {
	*x = TodoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todofy_todo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoRequest) ProtoMessage() {}

func (x *TodoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todofy_todo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoRequest.ProtoReflect.Descriptor instead.
func (*TodoRequest) Descriptor() ([]byte, []int) {
	return file_proto_todofy_todo_proto_rawDescGZIP(), []int{0}
}

func (x *TodoRequest) GetApp() TodoApp {
	if x != nil {
		return x.App
	}
	return TodoApp_TODO_APP_UNSPECIFIED
}

func (x *TodoRequest) GetMethod() PopullateTodoMethod {
	if x != nil {
		return x.Method
	}
	return PopullateTodoMethod_POPULLATE_TODO_METHOD_UNSPECIFIED
}

func (x *TodoRequest) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *TodoRequest) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *TodoRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *TodoRequest) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *TodoRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *TodoRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *TodoRequest) GetToName() string {
	if x != nil {
		return x.ToName
	}
	return ""
}

type TodoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id of the todo.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The message returned by the populating service.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *TodoResponse) Reset() {
	*x = TodoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todofy_todo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoResponse) ProtoMessage() {}

func (x *TodoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todofy_todo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoResponse.ProtoReflect.Descriptor instead.
func (*TodoResponse) Descriptor() ([]byte, []int) {
	return file_proto_todofy_todo_proto_rawDescGZIP(), []int{1}
}

func (x *TodoResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TodoResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_todofy_todo_proto protoreflect.FileDescriptor

var file_proto_todofy_todo_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x66, 0x79, 0x2f, 0x74,
	0x6f, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x74, 0x6f, 0x64, 0x6f, 0x66,
	0x79, 0x22, 0x80, 0x02, 0x0a, 0x0b, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x21, 0x0a, 0x03, 0x61, 0x70, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f,
	0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x66, 0x79, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x41, 0x70, 0x70, 0x52,
	0x03, 0x61, 0x70, 0x70, 0x12, 0x33, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x66, 0x79, 0x2e, 0x50, 0x6f,
	0x70, 0x75, 0x6c, 0x6c, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74,
	0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x74,
	0x6f, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0x38, 0x0a, 0x0c, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x66,
	0x0a, 0x07, 0x54, 0x6f, 0x64, 0x6f, 0x41, 0x70, 0x70, 0x12, 0x18, 0x0a, 0x14, 0x54, 0x4f, 0x44,
	0x4f, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x4f, 0x44, 0x4f, 0x5f, 0x41, 0x50, 0x50, 0x5f,
	0x44, 0x49, 0x44, 0x41, 0x33, 0x36, 0x35, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x54, 0x4f, 0x44,
	0x4f, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x54, 0x49, 0x43, 0x4b, 0x54, 0x49, 0x43, 0x4b, 0x10, 0x02,
	0x12, 0x14, 0x0a, 0x10, 0x54, 0x4f, 0x44, 0x4f, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x54, 0x4f, 0x44,
	0x4f, 0x49, 0x53, 0x54, 0x10, 0x03, 0x2a, 0x7e, 0x0a, 0x13, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x6c,
	0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x25, 0x0a,
	0x21, 0x50, 0x4f, 0x50, 0x55, 0x4c, 0x4c, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x4f, 0x44, 0x4f, 0x5f,
	0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x21, 0x0a, 0x1d, 0x50, 0x4f, 0x50, 0x55, 0x4c, 0x4c, 0x41, 0x54,
	0x45, 0x5f, 0x54, 0x4f, 0x44, 0x4f, 0x5f, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x4d, 0x41,
	0x49, 0x4c, 0x4a, 0x45, 0x54, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x50, 0x4f, 0x50, 0x55, 0x4c,
	0x4c, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x4f, 0x44, 0x4f, 0x5f, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44,
	0x5f, 0x41, 0x50, 0x49, 0x10, 0x02, 0x32, 0x48, 0x0a, 0x0b, 0x54, 0x6f, 0x64, 0x6f, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x0c, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x74,
	0x65, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x13, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x66, 0x79, 0x2e, 0x54,
	0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x74, 0x6f, 0x64,
	0x6f, 0x66, 0x79, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a,
	0x69, 0x79, 0x69, 0x78, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x6f, 0x2f,
	0x74, 0x6f, 0x64, 0x6f, 0x66, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_todofy_todo_proto_rawDescOnce sync.Once
	file_proto_todofy_todo_proto_rawDescData = file_proto_todofy_todo_proto_rawDesc
)

func file_proto_todofy_todo_proto_rawDescGZIP() []byte {
	file_proto_todofy_todo_proto_rawDescOnce.Do(func() {
		file_proto_todofy_todo_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_todofy_todo_proto_rawDescData)
	})
	return file_proto_todofy_todo_proto_rawDescData
}

var file_proto_todofy_todo_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_todofy_todo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_todofy_todo_proto_goTypes = []interface{}{
	(TodoApp)(0),             // 0: todofy.TodoApp
	(PopullateTodoMethod)(0), // 1: todofy.PopullateTodoMethod
	(*TodoRequest)(nil),      // 2: todofy.TodoRequest
	(*TodoResponse)(nil),     // 3: todofy.TodoResponse
}
var file_proto_todofy_todo_proto_depIdxs = []int32{
	0, // 0: todofy.TodoRequest.app:type_name -> todofy.TodoApp
	1, // 1: todofy.TodoRequest.method:type_name -> todofy.PopullateTodoMethod
	2, // 2: todofy.TodoService.PopulateTodo:input_type -> todofy.TodoRequest
	3, // 3: todofy.TodoService.PopulateTodo:output_type -> todofy.TodoResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_todofy_todo_proto_init() }
func file_proto_todofy_todo_proto_init() {
	if File_proto_todofy_todo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_todofy_todo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoRequest); i {
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
		file_proto_todofy_todo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoResponse); i {
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
			RawDescriptor: file_proto_todofy_todo_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_todofy_todo_proto_goTypes,
		DependencyIndexes: file_proto_todofy_todo_proto_depIdxs,
		EnumInfos:         file_proto_todofy_todo_proto_enumTypes,
		MessageInfos:      file_proto_todofy_todo_proto_msgTypes,
	}.Build()
	File_proto_todofy_todo_proto = out.File
	file_proto_todofy_todo_proto_rawDesc = nil
	file_proto_todofy_todo_proto_goTypes = nil
	file_proto_todofy_todo_proto_depIdxs = nil
}
