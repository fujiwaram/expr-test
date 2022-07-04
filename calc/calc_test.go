package calc_test

import (
	"reflect"
	"testing"

	"github.com/fujiwaram/expr-test/calc"
)

func TestCalc(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		ev         calc.Env
		want       interface{}
		wantErr    bool
	}{
		{
			expression: "$1 + $2",
			ev:         calc.Env{"$1": 1, "$2": 2},
			want:       3,
		},
		{
			expression: `map(filter($1, {#["$4"] >= 16}), {#["$2"]})`,
			ev: calc.Env{
				"$1": []map[string]interface{}{
					{"$2": "1", "$3": "foo", "$4": 16}, // 16以上
					{"$2": "2", "$3": "bar", "$4": 15},
					{"$2": "3", "$3": "baz", "$4": 18}, // 16以上
					{"$2": "4", "$3": "biz", "$4": 5},
					{"$2": "5", "$3": "buz", "$4": 51}, // 16以上
				},
			},
			want: []interface{}{"1", "3", "5"},
		},
		{
			expression: `map(filter($1, {#["$4"] < 16}), {#["$2"]})`,
			ev: calc.Env{
				"$1": []map[string]interface{}{
					{"$2": "1", "$3": "foo", "$4": 16},
					{"$2": "2", "$3": "bar", "$4": 15}, // 16未満
					{"$2": "3", "$3": "baz", "$4": 18},
					{"$2": "4", "$3": "biz", "$4": 5}, // 16未満
					{"$2": "5", "$3": "buz", "$4": 51},
				},
			},
			want: []interface{}{"2", "4"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calc.Calc(tt.expression, tt.ev)
			if (err != nil) != tt.wantErr {
				t.Errorf("Calc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Calc() = %v, want %v", got, tt.want)
			}
		})
	}
}
