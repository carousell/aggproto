package apis

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/carousell/aggproto/app/models"
	"github.com/carousell/aggproto/pkg/dsl"
)

type apiSpecs struct {
	specDir string
}

func (as *apiSpecs) list() ([]models.SpecInfo, error) {
	var desc []models.SpecInfo
	err := filepath.WalkDir(as.specDir, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		if err != nil {
			return err
		}
		context, er := dsl.ParseApiSpec(path)
		if er != nil {
			return er
		}
		splits := strings.Split(path, "/")
		desc = append(desc, models.SpecInfo{
			APIDescriptor: models.APIDescriptor{
				Name:    context.Api.Name(),
				Group:   context.Api.Group(),
				Version: context.Api.Version(),
			},
			Filename: splits[len(splits)-1],
		})
		return nil
	})
	if err != nil {
		return nil, err
	}
	return desc, nil
}

func (as *apiSpecs) load(file string) (*models.APIContext, error) {
	context, err := dsl.ParseApiSpec(fmt.Sprintf("%s/%s", as.specDir, file))
	if err != nil {
		return nil, err
	}
	apiContext := &models.APIContext{
		API: models.APIDescriptor{
			Name:    context.Api.Name(),
			Group:   context.Api.Group(),
			Version: context.Api.Version(),
		},
		Meta: struct {
			GoPackage string
		}{
			GoPackage: context.Meta.GoPackage,
		},
		Operations: struct {
			AllowedOperations []string
		}{
			context.Operation.AllowedOperations,
		},
	}
	for _, fd := range context.Output {
		opCtx := models.APIContextOutput{}
		if nmfd, ok := fd.Output().(*dsl.NamespacedMessageFieldDescriptor); ok {
			opCtx.Output.Value = nmfd.NamespacedField
		} else {
			log.Println("unknown output type")
			continue
		}
		switch tfd := fd.Input().(type) {
		case *dsl.StringValueFieldDescriptor:
			opCtx.Input.Type = models.PrimitiveStringType
			opCtx.Input.Value = tfd.Value
		case *dsl.IntValueFieldDescriptor:
			opCtx.Input.Type = models.PrimitiveStringType
			opCtx.Input.Value = strconv.FormatInt(tfd.Value, 10)
		case *dsl.BoolValueFieldDescriptor:
			opCtx.Input.Type = models.PrimitiveStringType
			opCtx.Input.Value = strconv.FormatBool(tfd.Value)
		case *dsl.NamespacedMessageFieldDescriptor:
			opCtx.Input.Type = models.MessageType
			opCtx.Input.Value = tfd.NamespacedField
		default:
			log.Println("unknown input type")
			continue
		}
		apiContext.Outputs = append(apiContext.Outputs, opCtx)
	}
	return apiContext, nil
}
