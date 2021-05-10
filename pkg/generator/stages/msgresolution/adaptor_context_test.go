//
//  This source file is part of the carousell/aggproto open source project
//
//  Copyright © 2021 Carousell and the project authors
//  Licensed under Apache License v2.0
//
//  See https://github.com/carousell/aggproto/blob/master/LICENSE for license information
//  See https://github.com/carousell/aggproto/graphs/contributors for the list of project authors
//
package msgresolution

import (
	"testing"

	"github.com/carousell/aggproto/pkg/dsl"
	"github.com/carousell/aggproto/pkg/generator/printer"
	"github.com/go-test/deep"
)

var (
	emptyTestExpected = `message EmptyTestResponse {
}
`
	nestedPrimitiveExpected = `message NestedPrimitiveTestResponse {
	message DefGen {
		string field1 = 1;
	}
	DefGen def = 1;
}
`
	namespacedMessageExpected = `message NamespacedMessageResponse {
	message EntityAGen {
		message EntityBField1Gen {
			string field_1 = 1;
		}
		EntityBField1Gen entity_b_field_1 = 1;
	}
	EntityAGen entity_a = 1;
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
		want   map[string]string
	}{
		{
			name: "empty units",
			fields: fields{
				apiDescriptor: dsl.GetApiDescriptor("adaptorTest", "emptyTest", 1),
				adaptorUnits:  nil,
			},
			want: map[string]string{
				"adaptorTest_v1/emptyTest.proto": emptyTestExpected,
			},
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
			want: map[string]string{
				"adaptorTest_v1/nestedPrimitiveTest.proto": nestedPrimitiveExpected,
			},
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
			want: map[string]string{
				"adaptorTest_v1/namespacedMessage.proto": namespacedMessageExpected,
			},
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
			got := p.Out()

			for k, v := range tt.want {
				if _, ok := got[k]; !ok {
					t.Errorf("Did not find key need %s", k)
					t.Fail()
				}
				if diff := deep.Equal(got[k], v); diff != nil {
					t.Errorf("PrintCode() got %v want %v diff %s", got, tt.want, diff)
				}

			}
		})
	}
}

var (
	emptyTestCodeExpected = `package adaptorTest_v1
func adaptEmptyTestResponse() *EmptyTestResponse{
	resp := &EmptyTestResponse{}
	return resp
}
`
	nestedPrimitiveCodeExpected = `package adaptorTest_v1
func adaptNestedPrimitiveTestResponse() *NestedPrimitiveTestResponse{
	resp := &NestedPrimitiveTestResponse{}
	resp.Def = &NestedPrimitiveTestResponse_DefGen{}
	resp.Def.Field1 = "hello"
	return resp
}
`
	namespacedMessageCodeExpected = `package adaptorTest_v1

import (
	"test/pkg_a"
)

func adaptNamespacedMessageResponse(entityA *pkg_a.EntityA) *NamespacedMessageResponse{
	entityBField1 := entityA.entityBField1
	resp := &NamespacedMessageResponse{}
	resp.EntityA = &NamespacedMessageResponse_EntityAGen{}
	resp.EntityA.EntityBField1 = &NamespacedMessageResponse_EntityAGen_EntityBField1Gen{}
	resp.EntityA.EntityBField1.Field1 = entityBField1.Field1
	return resp
}
`
	namespacedComposedHybridExpected = `package adaptorTest_v1
func adaptNamespacedComposedHybridResponse(entityA *pkg_a.EntityA, entityC *pkg_a.EntityC) *NamespacedComposedHybridResponse{
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
		want   map[string]string
	}{
		{
			name: "empty units",
			fields: fields{
				apiDescriptor: dsl.GetApiDescriptor("adaptorTest", "emptyTest", 1),
				adaptorUnits:  nil,
			},
			want: map[string]string{
				"adaptorTest_v1/emptyTest_adaptor.go": emptyTestCodeExpected,
			},
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
			want: map[string]string{
				"adaptorTest_v1/nestedPrimitiveTest_adaptor.go": nestedPrimitiveCodeExpected,
			},
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
												"entity_a", testBasicNestedMockMessage, false,
											},
											{
												"entity_b_field_1", testBasicNestedEntityBMockMessage, false,
											},
										},
									},
								},
							},
						},
					},
				},
			},
			want: map[string]string{
				"adaptorTest_v1/namespacedMessage_adaptor.go": namespacedMessageCodeExpected,
			},
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
													{"entity_a", testBasicNestedMockMessage, false},
													{"entity_b_field_1", testBasicNestedEntityBMockMessage, false},
												},
											},
											&messageFieldAdaptorUnit{
												fieldName:  "new_field_1",
												underlying: testComposedNestedWithPrimitiveMock.MessageFields[0],
												fieldMessageDependencies: []fieldMessageDependency{
													{"entity_c", testComposedNestedWithPrimitiveMock, false},
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
			want: map[string]string{
				"adaptorTest_v1/namespacedComposedHybrid_adaptor.go": namespacedComposedHybridExpected,
			},
		},
	}
	for _, tt := range tests {
		factory := printer.New()
		t.Run(tt.name, func(t *testing.T) {
			a := &adaptorContext{
				apiDescriptor: tt.fields.apiDescriptor,
				adaptorUnits:  tt.fields.adaptorUnits,
				meta:          dsl.Meta{GoPackage: "test"},
			}
			a.PrintCode(factory)

			got := factory.Out()
			for filename, body := range tt.want {
				if gotBody, ok := got[filename]; ok {
					if diff := deep.Equal(body, gotBody); diff != nil {
						t.Errorf(" got %s want %s diff %v", gotBody, body, diff)
					}
				}
			}
		})
	}
}
