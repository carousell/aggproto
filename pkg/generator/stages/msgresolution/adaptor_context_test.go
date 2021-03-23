package msgresolution

import (
	"testing"

	"github.com/carousell/aggproto/pkg/dsl"
	"github.com/carousell/aggproto/pkg/generator/printer"
)

var (
	emptyTestExpected = `message EmptyTestResponse{
}
`
	nestedPrimitiveExpected = `message NestedPrimitiveTestResponse{
	message DefGen{
		string field1 = 1;
	}
	DefGen def = 1;
}
`
)

func Test_adaptorContext_PrintProto(t *testing.T) {
	type fields struct {
		apiDescriptor dsl.ApiDescriptor
		adaptorUnits  []adaptorUnit
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "empty units",
			fields: fields{
				apiDescriptor: dsl.GetApiDescriptor("adaptorTest", "emptyTest", 1),
				adaptorUnits:  nil,
			},
			want: emptyTestExpected,
		},
		{
			name: "Nested primitive units",
			fields: fields{
				apiDescriptor: dsl.GetApiDescriptor("adaptorTest", "nestedPrimitiveTest", 1),
				adaptorUnits: []adaptorUnit{
					&nestedAdaptorUnit{
						fieldName: "def",
						nestedUnit: []adaptorUnit{
							&staticPrimitiveAdaptorUnit{
								fieldName:   "field1",
								primitiveFD: &dsl.StringValueFieldDescriptor{Value: "hello"},
							},
						},
					},
				},
			},
			want: nestedPrimitiveExpected,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := printer.New()
			a := &adaptorContext{
				apiDescriptor: tt.fields.apiDescriptor,
				adaptorUnits:  tt.fields.adaptorUnits,
			}
			a.PrintProto(p)
			got := p.String()
			if got != tt.want {
				t.Errorf("PrintProto got=%s want %s", got, tt.want)
			}
		})
	}
}

var (
	emptyTestCodeExpected = `package adaptorTest_v1
func AdaptEmptyTestResponse()*EmptyTestResponse{
	resp := &EmptyTestResponse{}
	return resp
}
`
	nestedPrimitiveCodeExpected = `package adaptorTest_v1
func AdaptNestedPrimitiveTestResponse()*NestedPrimitiveTestResponse{
	resp := &NestedPrimitiveTestResponse{}
	resp.def = &DefGen{}
	resp.def.field1 = "hello"
	return resp
}
`
)
func Test_adaptorContext_PrintCode(t *testing.T) {
	type fields struct {
		apiDescriptor dsl.ApiDescriptor
		adaptorUnits  []adaptorUnit
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "empty units",
			fields: fields{
				apiDescriptor: dsl.GetApiDescriptor("adaptorTest", "emptyTest", 1),
				adaptorUnits:  nil,
			},
			want: emptyTestCodeExpected,
		},
		{
			name: "Nested primitive units",
			fields: fields{
				apiDescriptor: dsl.GetApiDescriptor("adaptorTest", "nestedPrimitiveTest", 1),
				adaptorUnits: []adaptorUnit{
					&nestedAdaptorUnit{
						fieldName: "def",
						nestedUnit: []adaptorUnit{
							&staticPrimitiveAdaptorUnit{
								fieldName:   "field1",
								primitiveFD: &dsl.StringValueFieldDescriptor{Value: "hello"},
							},
						},
					},
				},
			},
			want: nestedPrimitiveCodeExpected,
		},
	}
	for _, tt := range tests {
		p := printer.New()
		t.Run(tt.name, func(t *testing.T) {
			a := &adaptorContext{
				apiDescriptor: tt.fields.apiDescriptor,
				adaptorUnits:  tt.fields.adaptorUnits,
			}
			a.PrintCode(p)
			got := p.String()
			if got != tt.want {
				t.Errorf("PrintCode() got %s want %s", got, tt.want)
			}
		})
	}
}
