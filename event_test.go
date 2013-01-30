package data

import(
  "testing"
  "github.com/bmizerany/assert"
  // "log"
)

func TestEncode(t *testing.T) {
  payload := make(map[string]interface{})
  payload["type"] = "order"
  payload["blah"] = 1
  
  nested := make(map[string]interface{})
  nested["account"] = "foo"
  nested["nested_int"] = 10
  nested["nested_float"] = 2.4
  
  payload["nested"] = nested
  
  event := &Event{
    payload: payload,
  }
  
  // #Get()
  title, _ := event.Get("type").String()
  assert.Equal(t, "order", title)
  
  blah1, _ := event.Get("blah").Int()
  
  assert.Equal(t, 1, blah1)
  
  // Encode to msgpack
  raw, _ := Encode(event)
  
  // Decode into new event
  event2, _ := Decode(raw)
  title2, _ := event2.Get("type").String()
  
  assert.Equal(t, title, title2)
  
  // nested
  account, _  := event2.Get("nested").Get("account").String()
  nint, _     := event2.Get("nested").Get("nested_int").Int()
  nfloat, _   := event2.Get("nested").Get("nested_float").Float64()
  
  assert.Equal(t, "foo", account)
  assert.Equal(t, 10, nint)
  assert.Equal(t, 2.4, nfloat)
}

