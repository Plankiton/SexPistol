package Sex

import (
    "testing"
    "fmt"
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

func CopyTestMapStructFunc (v interface{}, v2 interface{}, isMatching bool) (*okErr, error) {
    value, ok := v.([]interface{})
    if !ok {
        return nil, fmt.Errorf("value need to be a list")
    }
    vmap, ok := value[0].(map[string]interface{})
    if !ok {
        return nil, fmt.Errorf("value need to be a map")
    }
    vstr, ok := value[1].(okErr)
    if !ok {
        return nil, fmt.Errorf("value need to be a struct")
    }
    match, ok := v2.(okErr)
    if !ok {
        return nil, fmt.Errorf("value need to be a struct")
    }

    if err := Copy(vmap, &vstr); err != nil {
        Err(err)
        return nil, err
    }

    match_string := "is not"
    if !isMatching {
        match_string = "is"
    }

    if (vstr == match) != isMatching {
        err := fmt.Errorf(`Copy(%v, %v): %v %s matching with "%v"`, vmap, vstr, vstr, match_string, match)
        Err(vstr == match, isMatching, err)
        return nil, err
    }

    return &vstr, nil
}

func TestCopyMapStruct(t *testing.T) {
    tests := []testCase {
        {
            Value: []interface{}{
                map[string]interface{}{
                    "ok": true,
                    "err": false,
                    "Oth": "thing",
                },
                okErr {},
            },
            MatchValue: okErr {
                Ok: true,
                Err: false,
                Oth: "thing",
            },
            Match: true,
        },
        {
            Value: []interface{}{
                map[string]interface{}{
                    "ok": true,
                    "err": false,
                    "Oth": "thing",
                },
                okErr {},
            },
            MatchValue: okErr {},
            Match: false,
        },
        {
            Value: []interface{}{
                map[string]interface{}{
                    "ok": true,
                    "err": false,
                    "Oth": "thing",
                },
                okErr {},
            },
            MatchValue: okErr {
                Ok: true,
            },
            Match: false,
        },
        {
            Value: []interface{}{
                map[string]interface{}{
                    "ok": true,
                    "err": false,
                    "Oth": "thing",
                },
                okErr {},
            },
            MatchValue: okErr {
                Ok: true,
                Err: false,
            },
            Match: false,
        },
        {
            Value: []interface{}{
                map[string]interface{}{
                    "ok": true,
                    "err": false,
                },
                okErr {},
            },
            MatchValue: okErr {
                Ok: true,
                Err: false,
            },
            Match: true,
        },
    }

    for _, test := range tests {
        v, err := CopyTestMapStructFunc(test.Value, test.MatchValue, test.Match.(bool))
        if err != nil {
            t.Error(err)
        }

        if v != nil {
            test.Res = *v
        }

        if err := doTestCase(test); err != nil {
            t.Error(err)
        }
    }
}
