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

func (r *registry) searchMessage(prefix string) []*models.Message{
	messages := r.reg.ListMessages(registry2.LMOWithPrefixMatch(prefix))
	var ret []*models.Message
	for _, msg := range messages {
		ret = append(ret, adaptMessage(msg))
	}
	fmt.Printf("len %d\n", len(messages))
	return ret
}

func adaptMessage(msg registry2.Message) *models.Message {
	message := &models.Message{
		PackageName: msg.Package(),
		Name:        msg.Name(),
	}
	for _, f := range msg.Fields() {
		switch f.Type() {
		case registry2.FieldTypeMessage:
			message.Fields = append(message.Fields, &models.MessageField{
				Name:     f.Name(),
				Type:     f.Type().GoTypeString(),
				Message:  adaptMessage(f.Message()),
				Selected: false,
			})
		default:
			message.Fields = append(message.Fields, &models.MessageField{
				Name:     f.Name(),
				Type:     f.Type().GoTypeString(),
				Message:  nil,
				Selected: false,
			})
		}
	}
	return message
}
