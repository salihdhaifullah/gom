package main 

import (
	"fmt"
	"strings"
)

func ifElse[T any](t bool, x T, y T) T {
	if t {
		return x
	}
	return y
}

func fmtAround(around, content string) string {
	b := strings.Builder{}
	runes := []rune(content)

	start := 0
	for ; start < len(runes) && (runes[start] == ' ' || runes[start] == '\n'); start++ {
		b.WriteRune(runes[start])
	}

	if start < len(runes) {
		b.WriteString(around)
	}

	end := len(runes) - 1
	for ; end >= 0 && (runes[end] == ' ' || runes[end] == '\n'); end-- {
	}

	if start <= end {
		b.WriteString(string(runes[start : end+1]))
		b.WriteString(around)
	}

	for i := end + 1; i < len(runes); i++ {
		b.WriteRune(runes[i])
	}

	return b.String()
}

var (
	punctuationMap = map[rune]bool{
		'!':  true,
		'"':  true,
		'#':  true,
		'$':  true,
		'%':  true,
		'&':  true,
		'\'': true,
		'(':  true,
		')':  true,
		'*':  true,
		'+':  true,
		',':  true,
		'-':  true,
		'.':  true,
		'/':  true,
		':':  true,
		';':  true,
		'<':  true,
		'=':  true,
		'>':  true,
		'?':  true,
		'@':  true,
		'[':  true,
		'\\': true,
		']':  true,
		'^':  true,
		'_':  true,
		'`':  true,
		'{':  true,
		'|':  true,
		'}':  true,
		'~':  true,
	}
)

var HR string = "\n\n---\n\n"
var L string = "\n"

func Escape(input string) string {
	escaped := strings.Builder{}

	for _, ch := range input {
		if punctuationMap[ch] {
			escaped.WriteRune('\\')
		}
		escaped.WriteRune(ch)
	}

	return escaped.String()
}

func Doc(nodes ...string) string {
	return strings.Join(nodes, "")
}

func H1(nodes ...string) string {
	return fmt.Sprintf("# %s\n", strings.Join(nodes, ""))
}

func H2(nodes ...string) string {
	return fmt.Sprintf("## %s\n", strings.Join(nodes, ""))
}

func H3(nodes ...string) string {
	return fmt.Sprintf("### %s\n", strings.Join(nodes, ""))
}

func H4(nodes ...string) string {
	return fmt.Sprintf("#### %s\n", strings.Join(nodes, ""))
}

func H5(nodes ...string) string {
	return fmt.Sprintf("##### %s\n", strings.Join(nodes, ""))
}

func H6(nodes ...string) string {
	return fmt.Sprintf("###### %s\n", strings.Join(nodes, ""))
}

func Italic(nodes ...string) string {
    return fmtAround("*", strings.Join(nodes, ""))
}

func Bold(nodes ...string) string {
    return fmtAround("**", strings.Join(nodes, ""))
}

func Quote(nodes ...string) string {
	return fmt.Sprintf("> %s\n", strings.Join(nodes, ""))
}

func UL(nodes ...string) string {
	list := make([]string, len(nodes))
	for i := 0; i < len(nodes); i++ {
		list[i] = fmt.Sprintf("- %s", nodes[i])
	}

	return strings.Join(list, "\n")
}

func OL(nodes ...string) string {
	list := make([]string, len(nodes))
	for i := 0; i < len(nodes); i++ {
		list[i] = fmt.Sprintf("%d. %s", i+1, nodes[i])
	}

	return strings.Join(list, "\n")
}

func Task(done bool, nodes ...string) string {
	return fmt.Sprintf("- [%s] %s\n", ifElse(done, "x", " "), strings.Join(nodes, ""))
}

func Strikethrough(nodes ...string) string {
    return fmtAround("~~", strings.Join(nodes, ""))
}

func Code(text string) string {
    return fmtAround("`", text)
}

func Link(link, text string) string {
	return fmt.Sprintf("[%s](%s)", text, link)
}

func Img(link, text string) string {
	return fmt.Sprintf("![%s](%s)\n", text, link)
}

func CodeBlock(lang, text string) string {
	return fmt.Sprintf("\n```%s\n%s\n```\n", lang, text)
}

func If(val bool, nodes ...string) string {
	return ifElse(val, strings.Join(nodes, ""), "")
}

func IfElse(val bool, a, b string) string {
	return ifElse(val, a, b)
}

func For[T any](list []T, fn func(arg T) string) string {
	nodes := make([]string, len(list))
	for i := 0; i < len(list); i++ {
		nodes[i] = fn(list[i])
	}

	return strings.Join(nodes, "")
}
