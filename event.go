package data

import(
  "github.com/ugorji/go-msgpack"
  "encoding/json"
  "errors"
)

type Event struct {
  Payload interface{}
}

// Map type asserts to `map`
func (j *Event) Map() (map[string]interface{}, error) {
  var copy map[string]interface{}
  
  switch (j.Payload).(type) {
    default:
        return nil, errors.New("type assertion to map[string]interface{} failed")
    case map[string]interface{}:
        copy, _ = (j.Payload).(map[string]interface{})
    case map[interface{}]interface{}:
        copy = make(map[string]interface{})
        attrs, _ := (j.Payload).(map[interface{}]interface{})
        for i, k := range(attrs) {
          st, _ := i.(string)
          copy[st] = k
        }
  }
  
  return copy, nil

}

// Get returns a pointer to a new `Json` object 
// for `key` in its `map` representation
// 
// useful for chaining operations (to traverse a nested JSON):
//    js.Get("top_level").Get("dict").Get("value").Int()
func (j *Event) Get(key string) *Event {
  m, err := j.Map()
  if err == nil {
    if val, ok := m[key]; ok {
      return &Event{val}
    }
  }
  return &Event{nil}
}

// Array type asserts to an `array`
func (j *Event) Array() ([]interface{}, error) {
  if a, ok := (j.Payload).([]interface{}); ok {
    return a, nil
  }
  return nil, errors.New("type assertion to []interface{} failed")
}

// Bool type asserts to `bool`
func (j *Event) Bool() (bool, error) {
  if s, ok := (j.Payload).(bool); ok {
    return s, nil
  }
  return false, errors.New("type assertion to bool failed")
}

// String type asserts to `string`
func (j *Event) String() (string, error) {
  if s, ok := (j.Payload).(string); ok {
    return s, nil
  }
  return "", errors.New("type assertion to string failed")
}

// Float64 type asserts to `float64`
func (j *Event) Float64() (float64, error) {
  if i, ok := (j.Payload).(float64); ok {
    return i, nil
  }
  return -1, errors.New("type assertion to float64 failed")
}

// Int type asserts to `float64` then converts to `int`
func (j *Event) Int() (int, error) {
  
  if f, ok := (j.Payload).(int); ok {
    return int(f), nil
  }
  if f, ok := (j.Payload).(int8); ok {
    return int(f), nil
  }
  if f, ok := (j.Payload).(float64); ok {
    return int(f), nil
  }
  
  return -1, errors.New("type assertion to float64 failed")
}

// Int type asserts to `float64` then converts to `int64`
func (j *Event) Int64() (int64, error) {
  if f, ok := (j.Payload).(float64); ok {
    return int64(f), nil
  }
  
  return -1, errors.New("type assertion to float64 failed")
}

// Bytes type asserts to `[]byte`
func (j *Event) Bytes() ([]byte, error) {
  if s, ok := (j.Payload).(string); ok {
    return []byte(s), nil
  }
  return nil, errors.New("type assertion to []byte failed")
}




type EventsChannel chan *Event

func Decode(payload []byte) (*Event, error) {
  event := &Event{}
  err := msgpack.Unmarshal(payload, &event.Payload, nil)
  return event, err
}

func Encode(event *Event) (data []byte, err error) {
  data, err = msgpack.Marshal(&event.Payload)
  return
}

func DecodeJSON(payload []byte) (*Event, error) {
  event := &Event{}
  err := json.Unmarshal(payload, &event.Payload)
  return event, err
}

func EncodeJSON(event *Event) (data []byte, err error) {
  data, err = json.Marshal(&event.Payload)
  return
}