package msgresolution

import (
	"reflect"
	"testing"

	"github.com/carousell/aggproto/pkg/dsl"
	"github.com/carousell/aggproto/pkg/registry"
)

func Test_msgResolver_Resolve(t *testing.T) {
	type fields struct {
		r registry.Registry
	}
	type args struct {
		fds []dsl.FieldDescriptor
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   AdaptorContext
	}{
		{
			name: "mergeableStaticPrimitives",
			args: args{
				fds: dsl.GetFieldDescriptors(
					"a.b.c.field_1=42",
					"a.b.d.field_1=3.14",
					"a.b.c.field_2=false",
				),
			},
			want: &adaptorContext{
				adaptorUnits: []adaptorUnit{
					&nestedAdaptorUnit{
						name: "a",
						nestedUnit: []adaptorUnit{
							&nestedAdaptorUnit{
								name: "b",
								nestedUnit: []adaptorUnit{
									&nestedAdaptorUnit{
										name: "c",
										nestedUnit: []adaptorUnit{
											&staticPrimitiveAdaptorUnit{
												fieldName:   "field_1",
												primitiveFD: &dsl.IntValueFieldDescriptor{Value: 42},
											},
											&staticPrimitiveAdaptorUnit{
												fieldName:   "field_2",
												primitiveFD: &dsl.BoolValueFieldDescriptor{Value: false},
											},
										},
									},
									&nestedAdaptorUnit{
										name: "d",
										nestedUnit: []adaptorUnit{
											&staticPrimitiveAdaptorUnit{
												fieldName:   "field_1",
												primitiveFD: &dsl.FloatValueFieldDescriptor{Value: 3.14},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &msgResolver{
				r: tt.fields.r,
			}
			if got := m.Resolve(tt.args.fds); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Resolve() = %v, want %v", got, tt.want)
			}
		})
	}
}
