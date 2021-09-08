package Sex

import (
	"fmt"
	"reflect"
	"testing"
)

type testCase struct {
	Value      interface{}
	MatchValue interface{}
	Res        interface{}

	Match interface{}
}

func doTestCase(t testCase) error {
	isMatching, ok := t.Match.(bool)
	if !ok {
		isMatching = true
	}

	if reflect.DeepEqual(t.Res, t.MatchValue) != isMatching {
		return fmt.Errorf(`Result "%v" does not match with "%v"`, t.Res, t.MatchValue)
	}

	return nil
}

func TestFixPath(t *testing.T) {
	tests := []testCase{
		// Happy cases
		{
			Value:      "///",
			MatchValue: "/",
		},
		{
			Value:      "/joao//",
			MatchValue: "/joao",
		},
		{
			Value:      "//joao",
			MatchValue: "/joao",
		},
		{
			Value:      "",
			MatchValue: "/",
		},

		// Sad cases
		{
			Value:      "",
			MatchValue: "",
			Match:      false,
		},
		{
			Value:      "//joao",
			MatchValue: "/joao/",
			Match:      false,
		},
		{
			Value:      "/joao///",
			MatchValue: "/joao/",
			Match:      false,
		},
	}

	for _, test := range tests {
		test.Res = fixPath(test.Value.(string))
		if err := doTestCase(test); err != nil {
			t.Error(err)
		}
	}
}

type okErr struct {
	Ok  bool `json:"ok"`
	Err bool `json:"err"`
	Oth interface{}
}

func CopyTestFunc(v interface{}, r interface{}, isMatching bool) (interface{}, error) {
	value, ok := v.([]interface{})
	if !ok {
		return nil, fmt.Errorf("Copy(..., ...): v need to be a []interface{}")
	}

	v1 := value[0]
	v2 := value[1]

	if err := Copy(v1, v2); err != nil {
		Err(err)
		return nil, err
	}

	match_string := "is not"
	if !isMatching {
		match_string = "is"
	}

	v2_v := reflect.ValueOf(v2).Interface()
	r_v := reflect.ValueOf(r).Interface()

	if reflect.TypeOf(v2).Kind() == reflect.Ptr {
		v2_v = reflect.ValueOf(v2).Elem().Interface()
	}
	if reflect.TypeOf(r).Kind() == reflect.Ptr {
		r_v = reflect.ValueOf(r).Elem().Interface()
	}

	if reflect.DeepEqual(v2_v, r_v) != isMatching {
		err := fmt.Errorf(`Copy(%v, %v): %v %s matching with "%v"`, v1, v2, v2_v, match_string, r_v)
		Err(err)
		return nil, err
	}

	return v2, nil
}

func TestCopy(t *testing.T) {
	tests := []testCase{
		{
			Value: []interface{}{
				Dict{
					"ok":  true,
					"err": false,
					"Oth": "thing",
				},
				new(okErr),
			},
			MatchValue: &okErr{
				Ok:  true,
				Err: false,
				Oth: "thing",
			},
			Match: true,
		},
		{
			Value: []interface{}{
				Dict{
					"ok":  true,
					"err": false,
					"Oth": "thing",
				},
				new(Dict),
			},
			MatchValue: &Dict{
				"ok":  true,
				"err": false,
				"Oth": "thing",
			},
			Match: true,
		},
		{
			Value: []interface{}{
				okErr{
					Ok:  true,
					Err: false,
					Oth: "thing",
				},
				new(Dict),
			},
			MatchValue: &Dict{
				"ok":  true,
				"err": false,
				"Oth": "thing",
			},
			Match: true,
		},
		{
			Value: []interface{}{
				okErr{
					Ok:  true,
					Err: false,
					Oth: "thing",
				},
				new(Dict),
			},
			MatchValue: &Dict{
				"Ok":  true,
				"Err": false,
				"Oth": "thing",
			},
			Match: false,
		},
		{
			Value: []interface{}{
				okErr{
					Ok:  true,
					Err: false,
					Oth: "thing",
				},
				new(okErr),
			},
			MatchValue: &okErr{
				Ok:  true,
				Err: false,
				Oth: "thing",
			},
			Match: true,
		},
	}

	for i, test := range tests {
		v, err := CopyTestFunc(test.Value, test.MatchValue, test.Match.(bool))
		if err != nil {
			t.Errorf("Error trying %dº test case: %v", i+1, err)
			continue
		}

		test.Res = v
		if err := doTestCase(test); err != nil {
			t.Error(err)
		}
	}
}

func MergeTestFunc(v interface{}, r interface{}, isMatching bool) (interface{}, error) {
	value, ok := v.([]interface{})
	if !ok {
		return nil, fmt.Errorf("Merge(..., ...): v need to be a []interface{}")
	}

	v1 := value[0]
	v2 := value[1]

	merged, err := Merge(v1, v2)

	if err != nil {
		Err(err)
		return nil, err
	}

	match_string := "is not"
	if !isMatching {
		match_string = "is"
	}

	m_v := reflect.ValueOf(merged).Interface()
	r_v := reflect.ValueOf(r).Interface()

	if reflect.TypeOf(merged).Kind() == reflect.Ptr {
		m_v = reflect.ValueOf(m_v).Elem().Interface()
	}
	if reflect.TypeOf(r).Kind() == reflect.Ptr {
		r_v = reflect.ValueOf(r).Elem().Interface()
	}

	if reflect.DeepEqual(m_v, r_v) != isMatching {
		err := fmt.Errorf(`Merge(%v, %v): %v %s matching with "%v"`, v1, v2, m_v, match_string, r_v)
		Err(err)
		return nil, err
	}

	return v2, nil
}

func TestMerge(t *testing.T) {
	tests := []testCase{
		{
			Value: []interface{}{
				Dict{
					"ok":  true,
					"err": false,
					"Oth": "thing",
				},
				new(okErr),
			},
			MatchValue: &okErr{
				Ok:  true,
				Err: false,
				Oth: "thing",
			},
			Match: true,
		},
		{
			Value: []interface{}{
				Dict{
					"ok":  true,
					"err": false,
					"Oth": "thing",
				},
				new(Dict),
			},
			MatchValue: &Dict{
				"ok":  true,
				"err": false,
				"Oth": "thing",
			},
			Match: true,
		},
		{
			Value: []interface{}{
				okErr{
					Ok:  true,
					Err: false,
					Oth: "thing",
				},
				new(Dict),
			},
			MatchValue: &Dict{
				"ok":  true,
				"err": false,
				"Oth": "thing",
			},
			Match: true,
		},
		{
			Value: []interface{}{
				okErr{
					Ok:  true,
					Err: false,
					Oth: "thing",
				},
				new(okErr),
			},
			MatchValue: &okErr{
				Ok:  true,
				Err: false,
				Oth: "thing",
			},
			Match: true,
		},
	}

	for i, test := range tests {
		v, err := MergeTestFunc(test.Value, test.MatchValue, test.Match.(bool))
		if err != nil {
			t.Errorf("Error trying %dº test case: %v", i+1, err)
			continue
		}

		test.Res = v
		if err := doTestCase(test); err != nil {
			t.Error(err)
		}
	}
}

func TestFromJSON(t *testing.T) {
	tests := map[string][3]interface{}{
		`{"ok":true, "err": true, "Oth": "Maria"}`: [3]interface{}{
			new(okErr),
			&okErr{
				Ok:  true,
				Err: true,
				Oth: "Maria",
			},
			true,
		},
	}

	for in, out := range tests {
		if err := FromJSON([]byte(in), out[0]); err != nil {
			t.Error(err)
		}

		o := reflect.ValueOf(out[0]).Interface()
		r := reflect.ValueOf(out[1]).Interface()

		if reflect.TypeOf(out[0]).Kind() == reflect.Ptr {
			o = reflect.ValueOf(out[0]).Elem().Interface()
		}
		if reflect.TypeOf(out[1]).Kind() == reflect.Ptr {
			r = reflect.ValueOf(out[1]).Elem().Interface()
		}

		if ok := reflect.DeepEqual(o, r); ok != out[2] {
			t.Errorf("FromJSON(%s, %v): %v == %v -> %v",
				in, out[0], o, r, ok)
		}
	}
}
