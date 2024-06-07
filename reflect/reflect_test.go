package reflect

import (
	"reflect"
	"testing"
)

type PersonTest struct {
	Name    string
	Profile ProfileTest
}

type ProfileTest struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Tom"},
			[]string{"Tom"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				Town string
			}{"Tom", "Oslo"},
			[]string{"Tom", "Oslo"},
		},
		{
			"struct with int field",
			struct {
				Name string
				Sum  int
			}{"Timmy", 42},
			[]string{"Timmy"},
		},
		{
			"nested fields",
			PersonTest{
				"Timmy",
				ProfileTest{31, "Oslo"},
			},
			[]string{"Timmy", "Oslo"},
		},
		{
			"introduce pointer",
			&PersonTest{
				"Timmy",
				ProfileTest{31, "Oslo"},
			},
			[]string{"Timmy", "Oslo"},
		},
		{
			"give me a slice",
			[]ProfileTest{
				{25, "Oslo"},
				{30, "Trondheim"},
			},
			[]string{"Oslo", "Trondheim"},
		},
		{
			"give me array",
			[2]ProfileTest{
				{25, "Oslo"},
				{30, "Trondheim"},
			},
			[]string{"Oslo", "Trondheim"},
		},
		{
			"give me map",
			map[string]string{
				"Place": "Oslo",
				"Town":  "Trondheim",
			},
			[]string{"Oslo", "Trondheim"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			Walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("want %v, got %v", test.ExpectedCalls, got)
			}
		})
	}
}
