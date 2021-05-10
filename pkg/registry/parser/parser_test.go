package parser

import (
	_ "embed"
	"os"
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
	MessageFields: []*MessageField{
		{
			FieldName: "field_one",
			FieldType: registry2.FieldTypeString,
		},
	},
}
var bMessage = &MessageContainer{
	PackageName: "pkgb",
	MessageName: "BMessage",
	MessageDefinitions: []*MessageContainer{
		bSubMessage,
	},
	MessageFields: []*MessageField{
		{
			FieldName: "field_one",
			FieldType: registry2.FieldTypeString,
		},
		{
			FieldName:        "field_three",
			FieldType:        registry2.FieldTypeMessage,
			FieldMessageType: bSubMessage,
		},
		{
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
	MessageFields: []*MessageField{
		{
			FieldName: "name",
			FieldType: registry2.FieldTypeString,
		},
	},
}

func stitchMessage(mc *MessageContainer) {
	for _, f := range mc.MessageFields {
		f.FieldContext = mc
	}
	for _, d := range mc.MessageDefinitions {
		d.MessageParent = mc
		stitchMessage(d)
	}
}

func init() {
	stitchMessage(bMessage)
	stitchMessage(bReq)
}

func TestParser(t *testing.T) {

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
			testDir := os.TempDir()
			reg := Load(testDir)
			p := &parser{r: reg}
			for _, fdpEncoded := range tt.descriptors {
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
			registeredMsgs := map[string]registry2.Message{}
			for _, msg := range reg.ListMessages() {
				registeredMsgs[msg.FullName()] = msg
			}
			for _, msg := range tt.wantMessages {
				gotMsg, ok := registeredMsgs[msg.FullName()]
				if !ok {
					t.Errorf("want %v got %v", msg, nil)
					t.Fail()
				}
				if !reflect.DeepEqual(msg, gotMsg) {
					t.Errorf("want %v got %v", msg, gotMsg)
					t.Fail()
				}
			}
		})

	}

}
