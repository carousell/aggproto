package parser

import (
	_ "embed"
	"reflect"
	"testing"

	registry2 "github.com/carousell/aggproto/pkg/registry"
	"github.com/carousell/aggproto/test"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

var bSubMessage = &MessageContainer{
	PackageName: "pkgb.BMessage",
	MessageName: "BSubMessage",
	MessageFields: []registry2.Field{
		&MessageField{
			FieldName: "field_one",
			FieldType: registry2.FieldTypeString,
		},
	},
}
var bMessage = &MessageContainer{
	PackageName: "pkgb",
	MessageName: "BMessage",
	MessageDefinitions: []registry2.Message{
		bSubMessage,
	},
	MessageFields: []registry2.Field{
		&MessageField{
			FieldName: "field_one",
			FieldType: registry2.FieldTypeString,
		},
		&MessageField{
			FieldName:        "field_three",
			FieldType:        registry2.FieldTypeMessage,
			FieldMessageType: bSubMessage,
		},
		&MessageField{
			FieldName:        "field_seven",
			FieldType:        registry2.FieldTypeMessage,
			RepeatedField:    true,
			FieldMessageType: bSubMessage,
		},
	},
}

var bReq = &MessageContainer{
	MessageName: "BReq",
	PackageName: "pkgb",
	MessageFields: []registry2.Field{
		&MessageField{
			FieldName: "name",
			FieldType: registry2.FieldTypeString,
		},
	},
}

func stitchMessage(mc *MessageContainer) {
	for _, f := range mc.MessageFields {
		f := f.(*MessageField)
		f.FieldContext = mc
	}
	for _, d := range mc.MessageDefinitions {
		d := d.(*MessageContainer)
		d.MessageParent = mc
		stitchMessage(d)
	}
}

func init() {
	stitchMessage(bMessage)
}

func TestParser(t *testing.T) {
	reg := registry2.Mock()
	p := New(reg)

	var testCases = []struct {
		name         string
		descriptors  [][]byte
		wantPanic    bool
		wantMessages []registry2.Message
	}{
		{
			name:        "baseTest",
			descriptors: [][]byte{test.PkgaAProto},
			wantPanic:   true, // TODO unhandled types
		},
		{
			name:        "naiveTest",
			descriptors: [][]byte{test.PkgbBProto},
			wantPanic:   false,
			wantMessages: []registry2.Message{
				bMessage,
				bReq,
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			for _, fdpEncoded := range tt.descriptors {
				defer func() {
					err := recover()
					if err != nil && tt.wantPanic {
						t.Log("recovered", err)
					} else if err != nil && !tt.wantPanic {
						t.Error(err)
						t.Fail()
					} else if err == nil && tt.wantPanic {
						t.Error("wanted panic did not get panic")
						t.Fail()
					}
				}()
				fdp := &descriptorpb.FileDescriptorProto{}
				err := proto.Unmarshal(fdpEncoded, fdp)
				if err != nil {
					t.Error(err)
					t.Fail()
				}
				err = p.AddProtoContainer(fdp)
				if err != nil {
					t.Error(err)
					t.Fail()
				}
			}
			for i, msg := range reg.ListMessages() {
				wantMsg := tt.wantMessages[i]
				if !reflect.DeepEqual(wantMsg, msg) {
					t.Errorf("want %v got %v", wantMsg, msg)
					t.Fail()
				}
			}
		})

	}

}
