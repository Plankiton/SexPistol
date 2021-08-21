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

func fixPathTests(t *testing.T) {
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
        },
        {
            Value: "//joao",
            MatchValue: "/joao/",
        },
        {
            Value: "/joao///",
            MatchValue: "/joao/",
        },
    }

    for _, test := range tests {
        test.Res = fixPath(test.Value.(string))
        if err := doTestCase(test); err != nil {
            t.Error(err)
        }
    }
}

