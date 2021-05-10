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
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/carousell/aggproto/pkg/registry"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

type persistentRegistry interface {
	registry.Registry
	addMessages(msgs ...*MessageContainer)
	addOperations(ops ...*UnaryOperationContainer)
	addFileDescriptor(fd *descriptorpb.FileDescriptorProto)
	addEnums(enums ...*EnumContainer)
	save() error
}

type store struct {
	dirName         string
	fileDescriptors map[string]*descriptorpb.FileDescriptorProto
}

func (s *store) save() error {
	err := os.MkdirAll(s.dirName, os.ModePerm)
	if err != nil {
		return err
	}
	for _, v := range s.fileDescriptors {
		splits := strings.Split(v.GetName(), "/")
		protoName := splits[len(splits)-1]
		filePath := fmt.Sprintf("%s/%s_%s.fdp", s.dirName, v.GetPackage(), protoName)
		file, err := os.Create(filePath)
		if err != nil {
			return err
		}
		err = func() error {
			defer file.Close()
			fdpEncoded, err := proto.Marshal(v)
			if err != nil {
				return err
			}
			_, err = file.Write(fdpEncoded)
			if err != nil {
				return err
			}
			return nil
		}()
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *store) add(fd *descriptorpb.FileDescriptorProto) {
	protoName := fmt.Sprintf("%s_%s", fd.GetPackage(), fd.GetName())
	if _, ok := s.fileDescriptors[protoName]; !ok {
		s.fileDescriptors[protoName] = fd
	} else {
		log.Fatalf("FD already exists %s", protoName)
	}
}

type registryImpl struct {
	s     *store
	msgs  map[string]*MessageContainer
	ops   map[string]*UnaryOperationContainer
	enums map[string]*EnumContainer
}

func (r *registryImpl) save() error {
	return r.s.save()
}

func (r *registryImpl) ListMessages(opts ...registry.ListMessageOption) []registry.Message {
	options := registry.ListMessageOptions{}
	for _, opt := range opts {
		options = opt(options)
	}
	if options.ExactFullName != nil {
		if *options.ExactFullName == "" {
			return flattenAll(r.msgs)
		}
		return fetchMsgBySubDefinitionTraversal(r.msgs, *options.ExactFullName)
	}
	if options.PrefixMatch != nil {
		if *options.PrefixMatch == "" {
			return flattenAll(r.msgs)
		}
		return fetchShallowByParentName(r.msgs, *options.PrefixMatch)

	}
	return flattenAll(r.msgs)
}

func (r *registryImpl) ListEnums(opts ...registry.ListEnumOption) []registry.Enum {
	options := registry.ListEnumOptions{}

	for _, opt := range opts {
		options = opt(options)
	}

	if options.ExactFullName != nil {
		if *options.ExactFullName == "" {
			return flattenAllEnums(r.enums)
		}

		enum, ok := r.enums[*options.ExactFullName]
		if ok {
			return []registry.Enum{enum}
		}
		return fetchEnumBySubDefinitionTraversal(r.msgs, *options.ExactFullName)
	}
	return nil

}

func (r *registryImpl) ListOperations(opts ...registry.ListServiceOption) []registry.UnaryOperation {
	options := registry.ListServiceOptions{}
	for _, opt := range opts {
		options = opt(options)
	}
	if options.MatchReturn != nil {
		return fetchOpsByReturnMatch(r.ops, options.MatchReturn)
	}
	if len(options.OpNameIn) > 0 {
		return fetchOpsFilterByName(r.ops, options.OpNameIn)
	}
	return flattenAllOps(r.ops)
}

func flattenAllOps(ops map[string]*UnaryOperationContainer) []registry.UnaryOperation {
	var ret []registry.UnaryOperation
	for _, op := range ops {
		ret = append(ret, op)
	}
	return ret
}

func fetchOpsFilterByName(ops map[string]*UnaryOperationContainer, in []string) []registry.UnaryOperation {
	var ret []registry.UnaryOperation
	for _, opName := range in {
		if op, ok := ops[opName]; ok {
			ret = append(ret, op)
		}
	}
	return ret
}

func fetchOpsByReturnMatch(ops map[string]*UnaryOperationContainer, matchReturn registry.Message) []registry.UnaryOperation {
	var ret []registry.UnaryOperation
	for _, op := range ops {
		if op.Output() == matchReturn {
			ret = append(ret, op)
		}
	}
	return ret
}

func fetchShallowByParentName(msgs map[string]*MessageContainer, msgName string) []registry.Message {
	splits := strings.Split(msgName, ".")
	if len(splits) == 1 {
		return fetchMsgByPackageName(msgs, splits[0])
	}
	exactName := fmt.Sprintf("%s.%s", splits[0], splits[1])
	msg, ok := msgs[exactName]
	if !ok {
		return nil
	}
	return []registry.Message{msg}
}

func flattenAll(msgs map[string]*MessageContainer) []registry.Message {
	var all []registry.Message
	for _, msg := range msgs {
		all = append(all, msg)
	}
	return all
}

func flattenAllEnums(enums map[string]*EnumContainer) []registry.Enum {
	var all []registry.Enum
	for _, enum := range enums {
		all = append(all, enum)
	}
	return all
}

func fetchMsgBySubDefinitionTraversal(msgs map[string]*MessageContainer, msgName string) []registry.Message {
	splits := strings.Split(msgName, ".")
	if len(splits) == 1 {
		return fetchMsgByPackageName(msgs, splits[0])
	}
	exactName := fmt.Sprintf("%s.%s", splits[0], splits[1])
	msg, ok := msgs[exactName]
	if !ok {
		return nil
	}
	for _, defName := range splits[2:] { // is this a assumption that the nested message will be only upto 3 levels ? @vv
		var matchedDef *MessageContainer
		for _, subDef := range msg.MessageDefinitions {
			if subDef.Name() == defName {
				matchedDef = subDef
				break
			}
		}
		if matchedDef == nil {
			return nil
		}
		msg = matchedDef
	}
	return []registry.Message{msg}
}

func fetchEnumBySubDefinitionTraversal(msgs map[string]*MessageContainer, enumName string) []registry.Enum {
	splits := strings.Split(enumName, ".")

	enumDefName := splits[len(splits)-1]
	exactName := fmt.Sprintf("%s.%s", splits[0], splits[1])
	msg, ok := msgs[exactName]
	if !ok {
		return nil
	}

	// Add support for more nesting in enums
	for _, enumDef := range msg.EnumDefinitions {
		if enumDef.Name() == enumDefName {
			return []registry.Enum{enumDef}
		}
	}

	return nil
}

func fetchMsgByPackageName(msgs map[string]*MessageContainer, packageName string) []registry.Message {
	var ret []registry.Message
	for name, msg := range msgs {
		if strings.HasPrefix(name, packageName) {
			ret = append(ret, msg)
		}
	}
	return ret
}

func (r *registryImpl) addMessages(msgs ...*MessageContainer) {
	for _, msg := range msgs {
		if _, ok := r.msgs[msg.FullName()]; !ok {
			r.msgs[msg.FullName()] = msg
		} else {
			// TODO check if same message then ignore.
			log.Fatalf("message name already registered %s", msg.FullName())
		}
	}
}

func (r *registryImpl) addEnums(enums ...*EnumContainer) {
	for _, enum := range enums {
		if _, ok := r.enums[enum.FullName()]; !ok {
			r.enums[enum.FullName()] = enum
		} else {
			// TODO check if same message then ignore.
			log.Fatalf("message name already registered %s", enum.Name())
		}
	}

}

func (r *registryImpl) addOperations(ops ...*UnaryOperationContainer) {
	for _, op := range ops {
		if _, ok := r.ops[op.FullName()]; !ok {
			r.ops[op.FullName()] = op
		} else {
			log.Fatalf("service operation already registered %s", op.FullName())
		}
	}
}

func (r *registryImpl) addFileDescriptor(fd *descriptorpb.FileDescriptorProto) {
	pc := &protoContainer{fd, r}
	r.addMessages(pc.messages()...)
	r.addEnums(pc.enums()...)
	pc.populateMessageFields()
	r.addOperations(pc.operations()...)
	r.s.add(fd)
}

func Load(dirName string) persistentRegistry {
	r := &registryImpl{
		s: &store{
			dirName:         dirName,
			fileDescriptors: map[string]*descriptorpb.FileDescriptorProto{},
		},
		msgs:  map[string]*MessageContainer{},
		ops:   map[string]*UnaryOperationContainer{},
		enums: map[string]*EnumContainer{},
	}
	err := filepath.WalkDir(dirName, func(path string, d fs.DirEntry, err error) error {
		if d == nil || d.IsDir() {
			return nil
		}
		if err != nil {
			return err
		}
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		fdpEncoded, err := ioutil.ReadAll(file)
		if err != nil {
			return err
		}
		fdp := &descriptorpb.FileDescriptorProto{}
		err = proto.Unmarshal(fdpEncoded, fdp)
		if err != nil {
			return err
		}
		r.addFileDescriptor(fdp)
		return nil
	})
	if err != nil {
		panic(err)
	}
	return r
}
