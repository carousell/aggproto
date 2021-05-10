//
//  This source file is part of the carousell/aggproto open source project
//
//  Copyright © 2021 Carousell and the project authors
//  Licensed under Apache License v2.0
//
//  See https://github.com/carousell/aggproto/blob/master/LICENSE for license information
//  See https://github.com/carousell/aggproto/graphs/contributors for the list of project authors
//
package parser

import (
	"fmt"

	"github.com/carousell/aggproto/pkg/registry"
	"google.golang.org/protobuf/types/descriptorpb"
)

type MessageContainer struct {
	MessageName        string
	PackageName        string
	MessageParent      *MessageContainer
	MessageDefinitions []*MessageContainer
	EnumDefinitions    []*EnumContainer
	MessageFields      []*MessageField
}

func (m *MessageContainer) FullName() string {
	return fmt.Sprintf("%s.%s", m.PackageName, m.MessageName)
}

func (m *MessageContainer) Parent() registry.Message {
	if m.MessageParent != nil {
		return m.MessageParent
	}
	return nil
}

func (m *MessageContainer) Fields() []registry.Field {
	ret := make([]registry.Field, len(m.MessageFields))
	for i, f := range m.MessageFields {
		ret[i] = f
	}
	return ret
}

func (m *MessageContainer) Definitions() []registry.Message {
	ret := make([]registry.Message, len(m.MessageDefinitions))
	for i, m := range m.MessageDefinitions {
		ret[i] = m
	}
	return ret
}

func (m *MessageContainer) Package() string {
	return m.PackageName
}

func (m *MessageContainer) Name() string {
	return m.MessageName
}

func parseMessage(r registry.Registry, packageName string, msgType *descriptorpb.DescriptorProto) *MessageContainer {
	if msgType == nil {
		return nil
	}
	msg := &MessageContainer{
		PackageName: packageName,
		MessageName: msgType.GetName(),
	}
	var definitions []*MessageContainer
	for _, dp := range msgType.NestedType {
		subMsg := parseMessage(r, packageName, dp)
		subMsg.MessageParent = msg
		if subMsg != nil {
			definitions = append(definitions, subMsg)
		}
	}

	var enumDefinition []*EnumContainer
	for _, en := range msgType.EnumType {
		subEnum := parseEnum(r, packageName, en)
		subEnum.EnumParent = msg
		if subEnum != nil {
			// Should we do a add enum here to registry so that we do not need to traverse all the messages
			// when finding the enum type ?? @vinay
			enumDefinition = append(enumDefinition, subEnum)
		}
	}

	msg.EnumDefinitions = enumDefinition
	msg.MessageDefinitions = definitions
	return msg
}

func populateMessageField(r registry.Registry, packageName string, msgType *descriptorpb.DescriptorProto) {
	msgName := fmt.Sprintf("%s.%s", packageName, msgType.GetName())
	msgs := r.ListMessages(registry.LMOWithFullName(msgName))
	if len(msgs) != 1 {
		panic(fmt.Sprintf("message not found %s", msgName))
	}
	msg := msgs[0].(*MessageContainer)
	var fields []*MessageField
	for _, ft := range msgType.Field {
		fields = append(fields, parseField(r, msg, ft))
	}
	msg.MessageFields = fields

	for _, dp := range msgType.NestedType {
		populateMessageField(r, msgName, dp)
	}
}
