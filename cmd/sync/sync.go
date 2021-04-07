package sync

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/carousell/aggproto/pkg/dsl"
	"github.com/carousell/aggproto/pkg/generator"
	"github.com/carousell/aggproto/pkg/registry/parser"
)

func createFileWithParentDirs(filepath string) (*os.File, error) {
	splits := strings.Split(filepath, "/")
	if len(splits) > 1 {
		parentDir := strings.Join(splits[:len(splits)-1], "/")
		err := os.MkdirAll(parentDir, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}
	return os.Create(filepath)

}
func APIs(registryDir string, apiSpecDir string, protoOutDir string, goOutDir string) error {
	reg := parser.Load(registryDir)
	gen := generator.Generator(reg)
	return filepath.WalkDir(apiSpecDir, func(path string, d fs.DirEntry, err error) error {
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
		outFiles := gen.PlanAndGenerate(context)
		for fileName, fileContent := range outFiles {
			var filePath string
			if strings.HasSuffix(fileName, ".proto") {
				filePath = fmt.Sprintf("%s/%s", protoOutDir, fileName)
			} else if strings.HasSuffix(fileName, ".go") {
				filePath = fmt.Sprintf("%s/%s", goOutDir, fileName)
			} else {
				return errors.New("unknown file format " + fileName)
			}
			outFile, err := createFileWithParentDirs(filePath)
			if err != nil {
				return err
			}
			_, err = outFile.WriteString(fileContent)
			if err != nil {
				return err
			}
			err = outFile.Close()
			if err != nil {
				return err
			}

		}
		return nil
	})

}
