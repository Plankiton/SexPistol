package Sex

import (
	"fmt"
	"reflect"
	"testing"
)

type testCase struct {
    Value interface{}
    MatchValue interface{}
    Res    interface{}

    Match  interface{}
}

func doTestCase(t testCase) error {
    isMatching, ok := t.Match.(bool)
    if !ok {
        isMatching = true
    }

    if (t.Res == t.MatchValue) != isMatching {
        return fmt.Errorf(`Result "%v" does not match with "%v"`, t.Res, t.MatchValue)
    }

    return nil
}

func TestFixPath(t *testing.T) {
    tests := []testCase {
        // Happy cases
        {
            Value: "///",
            MatchValue: "/",
        },
        {
            Value: "/joao//",
            MatchValue: "/joao",
        },
        {
            Value: "//joao",
            MatchValue: "/joao",
        },
        {
            Value: "",
            MatchValue: "/",
        },

        // Sad cases
        {
            Value: "",
            MatchValue: "",
            Match: false,
        },
        {
            Value: "//joao",
            MatchValue: "/joao/",
            Match: false,
        },
        {
            Value: "/joao///",
            MatchValue: "/joao/",
            Match: false,
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
    Ok  bool        `gorm:"ok"`
    Err bool        `gorm:"err"`
    Oth interface{}
}

func CopyTestFunc (v interface{}, r interface{}, isMatching bool) (interface{}, error) {
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

    v2_v := reflect.ValueOf(v2)
    r_v := reflect.ValueOf(r)
    if (reflect.DeepEqual(v2_v, r_v)) != isMatching {
        err := fmt.Errorf(`Copy(%v, %v): %v %s matching with "%v"`, v1, v2, v2, match_string, r)
        Err(reflect.DeepEqual(v1, v2), isMatching, err)
        return nil, err
    }

    return v2, nil
}

func TestCopy(t *testing.T) {
    tests := []testCase {
        {
            Value: []interface{}{
                Dict {
                    "ok": true,
                    "err": false,
                    "Oth": "thing",
                },
                &okErr {},
            },
            MatchValue: &okErr {
                Ok: true,
                Err: false,
                Oth: "thing",
            },
            Match: true,
        },
        {
            Value: []interface{}{
                Dict {
                    "ok": true,
                    "err": false,
                    "Oth": "thing",
                },
                &okErr {},
            },
            MatchValue: okErr {},
            Match: false,
        },
        {
            Value: []interface{}{
                Dict {
                    "ok": true,
                    "err": false,
                    "Oth": "thing",
                },
                &okErr {},
            },
            MatchValue: okErr {
                Ok: true,
            },
            Match: false,
        },
        {
            Value: []interface{}{
                Dict {
                    "ok": true,
                    "err": false,
                    "Oth": "thing",
                },
                &okErr {},
            },
            MatchValue: okErr {
                Ok: true,
                Err: false,
            },
            Match: false,
        },
        {
            Value: []interface{}{
                Dict {
                    "ok": true,
                    "err": false,
                },
                &okErr {},
            },
            MatchValue: okErr {
                Ok: true,
                Err: false,
            },
            Match: true,
        },
    }

    for _, test := range tests {
        v, err := CopyTestFunc(test.Value, test.MatchValue, test.Match.(bool))
        if err != nil {
            t.Error(err)
        }

        if v != nil {
            test.Res = v
        }

        if err := doTestCase(test); err != nil {
            t.Error(err)
        }
    }
}
