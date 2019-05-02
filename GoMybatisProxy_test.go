package easyrpc_discovery

import (
	"fmt"
	"reflect"
	"testing"
)

type TestMapper struct {
	SelectByIds func(id string) (string, error) `mapperParams:"ids"`
}

func TestUseMapperValue(t *testing.T) {
	var test = TestMapper{}
	ProxyValue(reflect.ValueOf(&test), func(funcField reflect.StructField, field reflect.Value) func(arg ProxyArg) []reflect.Value {
		return func(arg ProxyArg) []reflect.Value {
			if len(arg.Args) <= 0 {
				t.Fatal("AopProxy() args len = 0")
			}
			if len(arg.TagArgs) <= 0 {
				t.Fatal("AopProxy() tagArgs len = 0")
			}
			var e error
			var returns = make([]reflect.Value, 0)
			returns = append(returns, reflect.ValueOf("yes return string="+arg.Args[0].Interface().(string)))
			returns = append(returns, reflect.Zero(reflect.TypeOf(&e).Elem()))
			return returns
		}
	})

	var result, e = test.SelectByIds("1234")
	fmt.Println(result, e)
	if e != nil {
		t.Fatal(e)
	}
}

func TestUseMapper(t *testing.T) {
	var test = TestMapper{}
	Proxy(&test, func(funcField reflect.StructField, field reflect.Value) func(arg ProxyArg) []reflect.Value {
		return func(arg ProxyArg) []reflect.Value {
			if len(arg.Args) <= 0 {
				t.Fatal("AopProxy() args len = 0")
			}
			if len(arg.TagArgs) <= 0 {
				t.Fatal("AopProxy() tagArgs len = 0")
			}
			var e error
			var returns = make([]reflect.Value, 0)
			returns = append(returns, reflect.ValueOf("yes return string="+arg.Args[0].Interface().(string)))
			returns = append(returns, reflect.Zero(reflect.TypeOf(&e).Elem()))
			return returns
		}
	})

	var result, e = test.SelectByIds("1234")
	fmt.Println(result, e)
	if e != nil {
		t.Fatal(e)
	}
}
