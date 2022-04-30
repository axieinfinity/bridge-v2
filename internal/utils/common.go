package utils

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
	"reflect"
)

type IUtils interface {
	Invoke(any interface{}, name string, args ...interface{}) (reflect.Value, error)
	LoadAbi(path string) (*abi.ABI, error)
	GetArguments(a abi.ABI, name string, data []byte, isInput bool) (abi.Arguments, error)
	UnpackToInterface(a abi.ABI, name string, data []byte, isInput bool, v interface{}) error
	Title(text string) string
}

type Utils struct{}

func (u *Utils) Invoke(any interface{}, name string, args ...interface{}) (reflect.Value, error) {
	method := reflect.ValueOf(any).MethodByName(name)
	methodType := method.Type()
	numIn := methodType.NumIn()
	if numIn > len(args) {
		return reflect.ValueOf(nil), fmt.Errorf("method %s must have minimum %d params. Have %d", name, numIn, len(args))
	}
	if numIn != len(args) && !methodType.IsVariadic() {
		return reflect.ValueOf(nil), fmt.Errorf("method %s must have %d params. Have %d", name, numIn, len(args))
	}
	in := make([]reflect.Value, len(args))
	for i := 0; i < len(args); i++ {
		var inType reflect.Type
		if methodType.IsVariadic() && i >= numIn-1 {
			inType = methodType.In(numIn - 1).Elem()
		} else {
			inType = methodType.In(i)
		}
		argValue := reflect.ValueOf(args[i])
		if !argValue.IsValid() {
			return reflect.ValueOf(nil), fmt.Errorf("Method %s. Param[%d] must be %s. Have %s", name, i, inType, argValue.String())
		}
		argType := argValue.Type()
		if argType.ConvertibleTo(inType) {
			in[i] = argValue.Convert(inType)
		} else {
			return reflect.ValueOf(nil), fmt.Errorf("Method %s. Param[%d] must be %s. Have %s", name, i, inType, argType)
		}
	}
	return method.Call(in)[0], nil
}

func (u *Utils) LoadAbi(path string) (*abi.ABI, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	a, err := abi.JSON(f)
	if err != nil {
		return nil, err
	}
	return &a, nil
}

func (u *Utils) GetArguments(a abi.ABI, name string, data []byte, isInput bool) (abi.Arguments, error) {
	// since there can't be naming collisions with contracts and events,
	// we need to decide whether we're calling a method or an event
	var args abi.Arguments
	if method, ok := a.Methods[name]; ok {
		if len(data)%32 != 0 {
			return nil, fmt.Errorf("abi: improperly formatted output: %s - Bytes: [%+v]", string(data), data)
		}
		if isInput {
			args = method.Inputs
		} else {
			args = method.Outputs
		}
	}
	if event, ok := a.Events[name]; ok {
		args = event.Inputs
	}
	if args == nil {
		return nil, errors.New("abi: could not locate named method or event")
	}
	return args, nil
}

func (u *Utils) UnpackToInterface(a abi.ABI, name string, data []byte, isInput bool, v interface{}) error {
	args, err := u.GetArguments(a, name, data, isInput)
	if err != nil {
		return err
	}
	unpacked, err := args.Unpack(data)
	if err != nil {
		return err
	}
	return args.Copy(v, unpacked)
}

func (u *Utils) Title(text string) string {
	c := cases.Title(language.English)
	return c.String(text)
}
