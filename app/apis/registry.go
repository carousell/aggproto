//
//  This source file is part of the carousell/aggproto open source project
//
//  Copyright © 2021 Carousell and the project authors
//  Licensed under Apache License v2.0
//
//  See https://github.com/carousell/aggproto/blob/master/LICENSE for license information
//  See https://github.com/carousell/aggproto/graphs/contributors for the list of project authors
//
package apis

import (
	"fmt"

	"github.com/carousell/aggproto/app/models"
	registry2 "github.com/carousell/aggproto/pkg/registry"
	"github.com/carousell/aggproto/pkg/registry/parser"
)

type registry struct {
	registryDir string
	reg         registry2.Registry
}

func (r *registry) init() {
	r.reg = parser.Load(r.registryDir)
}

func (r *registry) searchMessage(prefix string) []*models.Message {
	messages := r.reg.ListMessages(registry2.LMOWithPrefixMatch(prefix))
	var ret []*models.Message
	for _, msg := range messages {
		ret = append(ret, adaptMessage(msg, msg.Package()))
	}
	return ret
}

func adaptMessage(msg registry2.Message, packageName string) *models.Message {
	message := &models.Message{
		PackageName: packageName,
		Name:        msg.Name(),
	}
	for _, f := range msg.Fields() {
		switch f.Type() {
		case registry2.FieldTypeMessage:
			message.Fields = append(message.Fields, &models.MessageField{
				Name:     f.Name(),
				Type:     f.Type().GoTypeString(),
				Message:  adaptMessage(f.Message(), fmt.Sprintf("%s.%s", packageName, msg.Name())),
				Selected: false,
				Repeated: f.Repeated(),
			})
		default:
			message.Fields = append(message.Fields, &models.MessageField{
				Name:     f.Name(),
				Type:     f.Type().GoTypeString(),
				Message:  nil,
				Selected: false,
				Repeated: f.Repeated(),
			})
		}
	}
	return message
}
