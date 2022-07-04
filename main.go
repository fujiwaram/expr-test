package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/antonmedv/expr"
)

type Env map[string]interface{}

func (Env) Age(in, out string) string {
	return strings.Replace(in, "world", "user", 1) + out
}

// フロントから渡ってくる式（保存の時にだけ使用する）
var rawExp = "Age({プロフィール>生年月日}) + {加算日数}"

var reservedValue = regexp.MustCompile(`\$p[0-9]+`)
var convertValue = regexp.MustCompile(`\{.[^ }]+\}`)

// 変換後の式
var convertedExp = "Age($p1) + $p2"

// 式の値とIDのマッピング
var expMap = map[string]string{
	"$p1": "4cab30b6-61f6-4068-83f0-3a95cad3d29e",
	"$p2": "4cab30b6-61f6-4068-83f0-3a95cad3d29f",
}

func main() {
	rawExp := "$p1 + $p5"
	if x := reservedValue.FindString(rawExp); x != "" {
		fmt.Println(x)
	}
	rawExp2 := "p1 + p5"
	if x := reservedValue.FindString(rawExp2); x != "" {
		fmt.Println(x)
	}

	// env := Env{
	// 	"$p1":                                   "",
	// 	"a4cab30b6_61f6_4068_83f0_3a95cad3d29f": "",
	// }
	env2 := Env{
		"$p1":                                   "hello world 2",
		"a4cab30b6_61f6_4068_83f0_3a95cad3d29f": "!",
	}
	env3 := Env{
		"$p1":                                   "hell",
		"a4cab30b6_61f6_4068_83f0_3a95cad3d29f": "?",
	}

	exp := `Age($p1,a4cab30b6_61f6_4068_83f0_3a95cad3d29f)+a4cab30b6_61f6_4068_83f0_3a95cad3d29f`
	newExp := exp
	// prefix := "aaa"
	// idx := 1
	// newExp := regexpSlugV4.ReplaceAllStringFunc(exp, func(inStr string) string {
	// 	outStr := prefix + strconv.Itoa(idx)
	// 	idx++
	// 	return outStr
	// })

	// program, err := expr.Compile(newExp, expr.Env(env))
	program, err := expr.Compile(newExp) // 引数指定なくても動く
	if err != nil {
		fmt.Println(err)
		return
	}
	output, err := expr.Run(program, env2)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", output)

	output, err = expr.Run(program, env3)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", output)
}
