package msgresolution

import (
	"testing"

	"github.com/carousell/aggproto/pkg/dsl"
	"github.com/carousell/aggproto/pkg/generator/printer"
	"github.com/kylelemons/godebug/diff"
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
	namespacedMessageExpected = `message NamespacedMessageResponse{
	message EntityAGen{
		message EntityBField1Gen{
			string field_1 = 1;
		}
		EntityBField1Gen entity_b_field_1 = 1;
	}
	EntityAGen entity_a = 1;
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
		{
			name: "namespaced message",
			fields: fields{
				apiDescriptor: dsl.GetApiDescriptor("adaptorTest", "namespacedMessage", 1),
				adaptorUnits: []adaptorUnit{
					&nestedAdaptorUnit{
						fieldName: "entity_a",
						nestedUnit: []adaptorUnit{
							&nestedAdaptorUnit{
								fieldName: "entity_b_field_1",
								nestedUnit: []adaptorUnit{
									&messageFieldAdaptorUnit{
										fieldName:  "field_1",
										underlying: testBasicNestedEntityBMockMessage.MessageFields[0],
									},
								},
							},
						},
					},
				},
			},
			want: namespacedMessageExpected,
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
func AdaptEmptyTestResponse() *EmptyTestResponse{
	resp := &EmptyTestResponse{}
	return resp
}
`
	nestedPrimitiveCodeExpected = `package adaptorTest_v1
func AdaptNestedPrimitiveTestResponse() *NestedPrimitiveTestResponse{
	resp := &NestedPrimitiveTestResponse{}
	resp.Def = &DefGen{}
	resp.Def.Field1 = "hello"
	return resp
}
`
	namespacedMessageCodeExpected = `package adaptorTest_v1
func AdaptNamespacedMessageResponse(entityA *pkg_a.EntityA) *NamespacedMessageResponse{
	entityBField1 := entityA.entityBField1
	resp := &NamespacedMessageResponse{}
	resp.EntityA = &EntityAGen{}
	resp.EntityA.EntityBField1 = &EntityBField1Gen{}
	resp.EntityA.EntityBField1.Field1 = entityBField1.Field1
	return resp
}
`
	namespacedComposedHybridExpected = `package adaptorTest_v1
func AdaptNamespacedComposedHybridResponse(entityA *pkg_a.EntityA, entityC *pkg_a.EntityC) *NamespacedComposedHybridResponse{
	entityBField1 := entityA.entityBField1
	resp := &NamespacedComposedHybridResponse{}
	resp.PkgA = &PkgAGen{}
	resp.PkgA.EntityA = &EntityAGen{}
	resp.PkgA.EntityA.EntityBField1 = &EntityBField1Gen{}
	resp.PkgA.EntityA.EntityBField1.Field1 = entityBField1.Field1
	resp.PkgA.EntityA.EntityBField1.NewField1 = entityC.NewField1
	resp.PkgA.EntityA.EntityBField1.NewField2 = 42
	return resp
}
`
)

// todo add imports to print code
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
		{
			name: "namespaced message",
			fields: fields{
				apiDescriptor: dsl.GetApiDescriptor("adaptorTest", "namespacedMessage", 1),
				adaptorUnits: []adaptorUnit{
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
											{
												"entity_a", testBasicNestedMockMessage,
											},
											{
												"entity_b_field_1", testBasicNestedEntityBMockMessage,
											},
										},
									},
								},
							},
						},
					},
				},
			},
			want: namespacedMessageCodeExpected,
		},
		{
			name: "namespaced composed hybrid",
			fields: fields{
				apiDescriptor: dsl.GetApiDescriptor("adaptorTest", "namespacedComposedHybrid", 1),
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
			want: namespacedComposedHybridExpected,
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
				t.Errorf("PrintCode() got %s want %s diff %s", got, tt.want, diff.Diff(got, tt.want))
			}
		})
	}
}
