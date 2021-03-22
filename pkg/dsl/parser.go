package dsl

import (
	"strconv"
	"strings"
)

type parse_fd struct {
	input  InputFieldDescriptor
	output OutputFieldDescriptor
}

func (p *parse_fd) Input() InputFieldDescriptor {
	return p.input
}

func (p *parse_fd) Output() OutputFieldDescriptor {
	return p.output
}

func parseOutputFields(fields ...string) []FieldDescriptor {
	var fds []FieldDescriptor
	for _, f := range fields {
		var inputString, outputString string
		if strings.Contains(f, "=") {
			splits := strings.Split(f, "=")
			if len(splits) > 2 {
				panic("unhandled output field spec with more than one =")
			}
			outputString = splits[0]
			inputString = splits[1]
		} else {
			outputString = f
			inputString = f
		}
		ofd := &NamespacedMessageFieldDescriptor{outputString}
		if b, er := strconv.ParseBool(inputString); er == nil {
			fds = append(fds, &parse_fd{output: ofd, input: &BoolValueFieldDescriptor{b}})
			continue
		}
		if i, er := strconv.ParseInt(inputString, 10, 64); er == nil {
			fds = append(fds, &parse_fd{output: ofd, input: &IntValueFieldDescriptor{i}})
			continue
		}
		if f, er := strconv.ParseFloat(inputString, 10); er == nil {
			fds = append(fds, &parse_fd{output: ofd, input: &FloatValueFieldDescriptor{f}})
			continue
		}
		if strings.HasPrefix(inputString, "\"") && strings.HasSuffix(inputString, "\"") {
			fds = append(fds, &parse_fd{output: ofd, input: &StringValueFieldDescriptor{Value: strings.Trim(inputString, "\"")}})
			continue
		}
		fds = append(fds, &parse_fd{output: ofd, input: &NamespacedMessageFieldDescriptor{NamespacedField: inputString}})
	}
	return fds
}
