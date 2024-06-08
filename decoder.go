package ethidxer

import (
	"fmt"
	"math/big"
	"reflect"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// UnmarshalLog 从事件日志中解析数据到结构体中的字段
func unmarshalLogTopic[T any](log *types.Log, event *T) error {
	eventType := reflect.TypeOf(event)
	eventValue := reflect.ValueOf(event)

	if eventType.Kind() != reflect.Ptr || eventType.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("event parameter must be a pointer to a struct")
	}

	fieldCount := eventType.Elem().NumField()
	topicCount := len(log.Topics) - 1
	if fieldCount > topicCount {
		fieldCount = topicCount
	}
	for i := 0; i < fieldCount; i++ {
		fieldValue := eventValue.Elem().Field(i)
		if !fieldValue.CanSet() {
			continue
		}

		fieldType := eventType.Elem().Field(i).Type
		if err := parseField(log.Topics, i+1, fieldType, fieldValue); err != nil {
			return err
		}
	}
	return nil
}

func parseField(topics []common.Hash, start int, fieldType reflect.Type, fieldValue reflect.Value) error {
	switch fieldType.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		value := new(big.Int).SetBytes(topics[start].Bytes())
		fieldValue.SetUint(value.Uint64())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		value := new(big.Int).SetBytes(topics[start].Bytes())
		if value.Bit(255) == 1 {
			value.Sub(value, new(big.Int).Lsh(big.NewInt(1), 256))
		}
		fieldValue.SetInt(value.Int64())
	case reflect.Bool:
		fieldValue.SetBool(topics[start].Big().Uint64() == 1)
	case reflect.Slice:
		elemType := fieldType.Elem()
		if elemType.Kind() == reflect.Uint8 {
			bytes := topics[start].Bytes()
			fieldValue.SetBytes(bytes[:])
		} else {
			return fmt.Errorf("slice only support []byte")
		}
	case reflect.Array:
		elemType := fieldType.Elem()
		dataLength := fieldType.Len()
		if elemType.Kind() == reflect.Uint8 {
			bytes := topics[start].Bytes()
			if len(bytes) < dataLength {
				return fmt.Errorf("not enough bytes to fill the array")
			}
			for i := 0; i < dataLength; i++ {
				fieldValue.Index(i).Set(reflect.ValueOf(bytes[i]))
			}
		} else {
			return fmt.Errorf("only fixed-size byte arrays are supported")
		}
	case reflect.Ptr:
		if fieldValue.IsNil() {
			// Initialize the pointer if it's nil
			fieldValue.Set(reflect.New(fieldType.Elem()))
		}
		fieldValue = fieldValue.Elem()
		fieldType = fieldValue.Type()
		return parseField(topics, start, fieldType, fieldValue)
	case reflect.Struct:
		if fieldType == reflect.TypeOf(big.Int{}) {
			bigIntValue := new(big.Int).SetBytes(topics[start].Bytes())
			fieldValue.Set(reflect.ValueOf(*bigIntValue))
		} else {
			return fmt.Errorf("struct only support big.Int")
		}
	default:
		return fmt.Errorf("unsupported field type: %s", fieldType)
	}
	return nil
}
