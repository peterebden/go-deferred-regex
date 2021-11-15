// Package deferredregex offers a wrapper around the standard regexp.Regexp type which
// initialises itself on first use. This helps avoid spending excessive time compiling
// regexes at init time (which is unconditional and must complete before main is reached).
package deferredregex

import (
	"io"
	"regexp"
	"sync"
)

// A DeferredRegex is like a normal regexp but defers its initialisation until first use.
//
// Note that it uses MustCompile internally in order to mimic the regexp interface,
// so you only want to use it for static regexes that you know are valid (which is
// typically the only use case you would want this for anyway).
//
// It is safe for concurrent use, except for configuration methods such as Longest.
// It should not be copied.
type DeferredRegex struct {
	Re   string
	once sync.Once
	re   *regexp.Regexp
}

func (dr *DeferredRegex) init() {
	dr.once.Do(func() {
		dr.re = regexp.MustCompile(dr.Re)
	})
}

// UnmarshalText implements the encoding.TextUnmarshaler interface
// Note that it still defers the parse at this point.
func (dr *DeferredRegex) UnmarshalText(text []byte) error {
	dr.Re = string(text)
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface
func (dr *DeferredRegex) MarshalText() ([]byte, error) {
	return []byte(dr.Re), nil
}

// UnmarshalFlag implements the Unmarshaler interface from go-flags
// Note that it still defers the parse at this point.
func (dr *DeferredRegex) UnmarshalFlag(in string) error {
	dr.Re = in
	return nil
}

// MarshalFlag implements the Marshaler interface from go-flags
func (dr *DeferredRegex) MarshalFlag() (string, error) {
	return dr.Re, nil
}

func (dr *DeferredRegex) Expand(dst []byte, template []byte, src []byte, match []int) []byte {
	dr.init()
	return dr.re.Expand(dst, template, src, match)
}

func (dr *DeferredRegex) ExpandString(dst []byte, template string, src string, match []int) []byte {
	dr.init()
	return dr.re.ExpandString(dst, template, src, match)
}

func (dr *DeferredRegex) Find(b []byte) []byte {
	dr.init()
	return dr.re.Find(b)
}

func (dr *DeferredRegex) FindAll(b []byte, n int) [][]byte {
	dr.init()
	return dr.re.FindAll(b, n)
}

func (dr *DeferredRegex) FindAllIndex(b []byte, n int) [][]int {
	dr.init()
	return dr.re.FindAllIndex(b, n)
}

func (dr *DeferredRegex) FindAllString(s string, n int) []string {
	dr.init()
	return dr.re.FindAllString(s, n)
}

func (dr *DeferredRegex) FindAllStringIndex(s string, n int) [][]int {
	dr.init()
	return dr.re.FindAllStringIndex(s, n)
}

func (dr *DeferredRegex) FindAllStringSubmatch(s string, n int) [][]string {
	dr.init()
	return dr.re.FindAllStringSubmatch(s, n)
}

func (dr *DeferredRegex) FindAllStringSubmatchIndex(s string, n int) [][]int {
	dr.init()
	return dr.re.FindAllStringSubmatchIndex(s, n)
}

func (dr *DeferredRegex) FindAllSubmatch(b []byte, n int) [][][]byte {
	dr.init()
	return dr.re.FindAllSubmatch(b, n)
}

func (dr *DeferredRegex) FindAllSubmatchIndex(b []byte, n int) [][]int {
	dr.init()
	return dr.re.FindAllSubmatchIndex(b, n)
}

func (dr *DeferredRegex) FindIndex(b []byte) (loc []int) {
	dr.init()
	return dr.re.FindIndex(b)
}

func (dr *DeferredRegex) FindReaderIndex(r io.RuneReader) (loc []int) {
	dr.init()
	return dr.re.FindReaderIndex(r)
}

func (dr *DeferredRegex) FindReaderSubmatchIndex(r io.RuneReader) []int {
	dr.init()
	return dr.re.FindReaderSubmatchIndex(r)
}

func (dr *DeferredRegex) FindString(s string) string {
	dr.init()
	return dr.re.FindString(s)
}

func (dr *DeferredRegex) FindStringIndex(s string) (loc []int) {
	dr.init()
	return dr.re.FindStringIndex(s)
}

func (dr *DeferredRegex) FindStringSubmatch(s string) []string {
	dr.init()
	return dr.re.FindStringSubmatch(s)
}

func (dr *DeferredRegex) FindStringSubmatchIndex(s string) []int {
	dr.init()
	return dr.re.FindStringSubmatchIndex(s)
}

func (dr *DeferredRegex) FindSubmatch(b []byte) [][]byte {
	dr.init()
	return dr.re.FindSubmatch(b)
}

func (dr *DeferredRegex) FindSubmatchIndex(b []byte) []int {
	dr.init()
	return dr.re.FindSubmatchIndex(b)
}

func (dr *DeferredRegex) LiteralPrefix() (prefix string, complete bool) {
	dr.init()
	return dr.re.LiteralPrefix()
}

func (dr *DeferredRegex) Longest() {
	dr.init()
	dr.re.Longest()
}

func (dr *DeferredRegex) Match(b []byte) bool {
	dr.init()
	return dr.re.Match(b)
}

func (dr *DeferredRegex) MatchReader(r io.RuneReader) bool {
	dr.init()
	return dr.re.MatchReader(r)
}

func (dr *DeferredRegex) MatchString(s string) bool {
	dr.init()
	return dr.re.MatchString(s)
}

func (dr *DeferredRegex) NumSubexp() int {
	dr.init()
	return dr.re.NumSubexp()
}

func (dr *DeferredRegex) ReplaceAll(src, repl []byte) []byte {
	dr.init()
	return dr.re.ReplaceAll(src, repl)
}

func (dr *DeferredRegex) ReplaceAllFunc(src []byte, repl func([]byte) []byte) []byte {
	dr.init()
	return dr.re.ReplaceAllFunc(src, repl)
}

func (dr *DeferredRegex) ReplaceAllLiteral(src, repl []byte) []byte {
	dr.init()
	return dr.re.ReplaceAllLiteral(src, repl)
}

func (dr *DeferredRegex) ReplaceAllLiteralString(src, repl string) string {
	dr.init()
	return dr.re.ReplaceAllLiteralString(src, repl)
}

func (dr *DeferredRegex) ReplaceAllString(src, repl string) string {
	dr.init()
	return dr.re.ReplaceAllString(src, repl)
}

func (dr *DeferredRegex) ReplaceAllStringFunc(src string, repl func(string) string) string {
	dr.init()
	return dr.re.ReplaceAllStringFunc(src, repl)
}

func (dr *DeferredRegex) Split(s string, n int) []string {
	dr.init()
	return dr.re.Split(s, n)
}

func (dr *DeferredRegex) String() string {
	dr.init()
	return dr.re.String()
}

func (dr *DeferredRegex) SubexpIndex(name string) int {
	dr.init()
	return dr.re.SubexpIndex(name)
}

func (dr *DeferredRegex) SubexpNames() []string {
	dr.init()
	return dr.re.SubexpNames()
}
