package msgresolution

import (
	"testing"

	"github.com/carousell/aggproto/pkg/dsl"
	"github.com/carousell/aggproto/pkg/registry"
	"github.com/carousell/aggproto/pkg/registry/parser"
	"github.com/go-test/deep"
)

var (
	testBasicNestedEntityBMockMessage = &parser.MessageContainer{
		PackageName: "pkg_a",
		MessageName: "entity_b",
		MessageFields: []registry.Field{
			&parser.MessageField{
				FieldType: registry.FieldTypeString,
				FieldName: "field_1",
			},
			&parser.MessageField{
				FieldType: registry.FieldTypeBool,
				FieldName: "field_2",
			},
		},
	}
	testBasicNestedMockMessage = &parser.MessageContainer{
		PackageName: "pkg_a",
		MessageName: "entity_a",
		MessageFields: []registry.Field{
			&parser.MessageField{
				FieldName:        "entity_b_field_1",
				FieldType:        registry.FieldTypeMessage,
				FieldMessageType: testBasicNestedEntityBMockMessage,
			},
		},
	}
	testComposedNestedWithPrimitiveMock = &parser.MessageContainer{
		MessageName: "entity_c",
		PackageName: "pkg_a",
		MessageFields: []registry.Field{
			&parser.MessageField{
				FieldType: registry.FieldTypeString,
				FieldName: "field_1",
			},
		},
	}
)

func stitchMessage(mc *parser.MessageContainer) {
	for _, f := range mc.MessageFields {
		f := f.(*parser.MessageField)
		f.FieldContext = mc
	}
	for _, d := range mc.MessageDefinitions {
		d := d.(*parser.MessageContainer)
		d.MessageParent = mc
		stitchMessage(d)
	}
}

func init() {
	deep.CompareUnexportedFields = true
	deep.NilMapsAreEmpty = true
	deep.NilSlicesAreEmpty = true
	stitchMessage(testBasicNestedEntityBMockMessage)
	stitchMessage(testBasicNestedMockMessage)
}

func Test_msgResolver_Resolve(t *testing.T) {
	api := dsl.GetApiDescriptor("test", "test", 1)
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
				apiDescriptor: api,
				adaptorUnits: []adaptorUnit{
					&nestedAdaptorUnit{
						fieldName: "a",
						nestedUnit: []adaptorUnit{
							&nestedAdaptorUnit{
								fieldName: "b",
								nestedUnit: []adaptorUnit{
									&nestedAdaptorUnit{
										fieldName: "c",
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
										fieldName: "d",
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
		{
			name: "basicNested",
			args: args{fds: dsl.GetFieldDescriptors(
				"pkg_a.entity_a.entity_b_field_1.field_1",
				"pkg_a.entity_a.entity_b_field_1.field_2",
			)},
			fields: fields{r: registry.Mock().OnListMessageMatchPrefix("pkg_a.entity_a", []registry.Message{
				testBasicNestedMockMessage,
			})},
			want: &adaptorContext{
				apiDescriptor: api,
				adaptorUnits: []adaptorUnit{
					&nestedAdaptorUnit{
						fieldName: "pkg_a",
						nestedUnit: []adaptorUnit{
							&nestedAdaptorUnit{
								fieldName: "entity_a",
								nestedUnit: []adaptorUnit{
									&nestedAdaptorUnit{
										fieldName: "entity_b_field_1",
										nestedUnit: []adaptorUnit{
											&messageFieldAdaptorUnit{
												fieldName:  "field_1",
												underlying: testBasicNestedEntityBMockMessage.MessageFields[0],
												fieldMessageDependencies: []fieldMessageDependency{
													{"entity_a", testBasicNestedMockMessage},
													{"entity_b_field_1", testBasicNestedEntityBMockMessage},
												},
											},
											&messageFieldAdaptorUnit{
												fieldName:  "field_2",
												underlying: testBasicNestedEntityBMockMessage.MessageFields[1],
												fieldMessageDependencies: []fieldMessageDependency{
													{"entity_a", testBasicNestedMockMessage},
													{"entity_b_field_1", testBasicNestedEntityBMockMessage},
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
		},
		{
			name: "composedNestedWithPrimitives",
			args: args{fds: dsl.GetFieldDescriptors("pkg_a.entity_a.entity_b_field_1.field_1",
				"pkg_a.entity_a.entity_b_field_1.new_field_1=pkg_a.entity_c.field_1",
				"pkg_a.entity_a.entity_b_field_1.new_field_2=42",
			)},
			fields: fields{r: registry.Mock().
				OnListMessageMatchPrefix("pkg_a.entity_a", []registry.Message{
					testBasicNestedMockMessage,
				}).
				OnListMessageMatchPrefix("pkg_a.entity_c", []registry.Message{
					testComposedNestedWithPrimitiveMock,
				}),
			},
			want: &adaptorContext{
				apiDescriptor: api,
				adaptorUnits: []adaptorUnit{
					&nestedAdaptorUnit{
						fieldName: "pkg_a",
						nestedUnit: []adaptorUnit{
							&nestedAdaptorUnit{
								fieldName: "entity_a",
								nestedUnit: []adaptorUnit{
									&nestedAdaptorUnit{
										fieldName: "entity_b_field_1",
										nestedUnit: []adaptorUnit{
											&messageFieldAdaptorUnit{
												fieldName:  "field_1",
												underlying: testBasicNestedEntityBMockMessage.MessageFields[0],
												fieldMessageDependencies: []fieldMessageDependency{
													{"entity_a", testBasicNestedMockMessage},
													{"entity_b_field_1", testBasicNestedEntityBMockMessage},
												},
											},
											&messageFieldAdaptorUnit{
												fieldName:  "new_field_1",
												underlying: testComposedNestedWithPrimitiveMock.MessageFields[0],
												fieldMessageDependencies: []fieldMessageDependency{
													{"entity_c", testComposedNestedWithPrimitiveMock},
												},
											},
											&staticPrimitiveAdaptorUnit{
												fieldName:   "new_field_2",
												primitiveFD: &dsl.IntValueFieldDescriptor{Value: 42},
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
			got := m.Resolve(api, tt.args.fds)
			if diff := deep.Equal(got, tt.want); diff != nil {
				t.Error(diff)
			}
		})
	}
}
