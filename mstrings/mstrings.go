package mstrings

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
)

func Compress(input string) string {
	g := groupBy(input)
	m := lo.Map(g, func(item string, index int) string {
		return compressOne(item)
	})
	return concat(m)
}

func groupBy(input string) []string {
	s := strings.Split(input, "")
	g := lo.GroupBy(s, func(i string) string {
		return i
	})
	result := lo.MapToSlice(g, func(k string, v []string) string {
		return strings.Join(v, "")
	})
	return result
}

func concat(input []string) string {
	return strings.Join(input, "")
}

func compressOne(input string) string {
	return fmt.Sprintf("%v%v", string(input[0]), len(input))
}
