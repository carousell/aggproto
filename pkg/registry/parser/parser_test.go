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

var bSubMessage = &messageContainer{
	packageName: "pkgb.BMessage",
	name:        "BSubMessage",
	fields: []registry2.Field{
		&messageField{
			name:      "field_one",
			fieldType: registry2.FieldTypeString,
		},
	},
}
var bMessage = &messageContainer{
	packageName: "pkgb",
	name:        "BMessage",
	definitions: []registry2.Message{
		bSubMessage,
	},
	fields: []registry2.Field{
		&messageField{
			name:      "field_one",
			fieldType: registry2.FieldTypeString,
		},
		&messageField{
			name:      "field_three",
			fieldType: registry2.FieldTypeMessage,
			message:   bSubMessage,
		},
		&messageField{
			name:      "field_seven",
			fieldType: registry2.FieldTypeMessage,
			repeated:  true,
			message:   bSubMessage,
		},
	},
}

var bReq = &messageContainer{
	name:        "BReq",
	packageName: "pkgb",
	fields: []registry2.Field{
		&messageField{
			name:      "name",
			fieldType: registry2.FieldTypeString,
		},
	},
}

func stitchMessage(mc *messageContainer) {
	for _, f := range mc.fields {
		f := f.(*messageField)
		f.context = mc
	}
	for _, d := range mc.definitions {
		d := d.(*messageContainer)
		d.parent = mc
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
