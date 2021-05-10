//
//  This source file is part of the carousell/aggproto open source project
//
//  Copyright © 2021 Carousell and the project authors
//  Licensed under Apache License v2.0
//
//  See https://github.com/carousell/aggproto/blob/master/LICENSE for license information
//  See https://github.com/carousell/aggproto/graphs/contributors for the list of project authors
//
package dsl

import (
	"reflect"
	"testing"
)

func Test_parseOutputFields(t *testing.T) {
	type args struct {
		fields []string
	}
	tests := []struct {
		name string
		args args
		want []FieldDescriptor
	}{
		{
			name: "Test Int",
			args: args{
				fields: []string{
					"pkg_a.a.f1=42",
				},
			},
			want: []FieldDescriptor{
				&parseFd{
					output: &NamespacedMessageFieldDescriptor{NamespacedField: "pkg_a.a.f1"},
					input:  &IntValueFieldDescriptor{Value: 42},
				},
			},
		},
		{
			name: "Test Float",
			args: args{
				fields: []string{
					"pkg_a.a.f2=3.14",
				},
			},
			want: []FieldDescriptor{
				&parseFd{
					output: &NamespacedMessageFieldDescriptor{NamespacedField: "pkg_a.a.f2"},
					input:  &FloatValueFieldDescriptor{Value: 3.14},
				},
			},
		},
		{
			name: "Test String",
			args: args{
				fields: []string{
					"pkg_a.a.f3=\"hello world\"",
				},
			},
			want: []FieldDescriptor{
				&parseFd{
					output: &NamespacedMessageFieldDescriptor{NamespacedField: "pkg_a.a.f3"},
					input:  &StringValueFieldDescriptor{Value: "hello world"},
				},
			},
		},
		{
			name: "Test Namespaced",
			args: args{
				fields: []string{
					"pkg_a.a.f4=pkg_a.entity1.field1",
					"pkg_a.a.f5",
				},
			},
			want: []FieldDescriptor{
				&parseFd{
					output: &NamespacedMessageFieldDescriptor{NamespacedField: "pkg_a.a.f4"},
					input:  &NamespacedMessageFieldDescriptor{NamespacedField: "pkg_a.entity1.field1"},
				},
				&parseFd{
					output: &NamespacedMessageFieldDescriptor{NamespacedField: "pkg_a.a.f5"},
					input:  &NamespacedMessageFieldDescriptor{NamespacedField: "pkg_a.a.f5"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseOutputFields(tt.args.fields...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseOutputFields() = %v, want %v", got, tt.want)
			}
		})
	}
}
