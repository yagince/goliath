package test_util

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"testing"
)

type Matcher interface {
	IsSatisfied() bool
	Expected() string
	Actual() string
	ErrorMessage() string
}

func printErrors(matcher Matcher) {
	_, filename, line_number, _ := runtime.Caller(2)
	buffer, _ := ioutil.ReadFile(filename)
	fmt.Println(green(filename))
	line := strings.Split(string(buffer), "\n")[line_number-1]
	fmt.Println(grey(fmt.Sprintf("%d: %s", line_number, line)))
	fmt.Println(red(matcher.ErrorMessage()))
	fmt.Println(grey(fmt.Sprintf("expected %s", matcher.Expected())))
	fmt.Println(grey(fmt.Sprintf("got      %s", matcher.Actual())))
}

func Assert(t *testing.T, matcher Matcher) {
	if !matcher.IsSatisfied() {
		printErrors(matcher)
		t.FailNow()
	}
}

func Verify(t *testing.T, matcher Matcher) {
	if !matcher.IsSatisfied() {
		printErrors(matcher)
		t.Fail()
	}
}

type Equal struct {
	ExpectedValue interface{}
	ActualValue   interface{}
}

func (matcher Equal) IsSatisfied() bool {
	return reflect.DeepEqual(matcher.ExpectedValue, matcher.ActualValue)
}
func (matcher Equal) ErrorMessage() string {
	return fmt.Sprintf("`%v` is not `%v`", matcher.ActualValue, matcher.ExpectedValue)
}
func (matcher Equal) Expected() string {
	return fmt.Sprintf("%v", matcher.ExpectedValue)
}
func (matcher Equal) Actual() string {
	return fmt.Sprintf("%v", matcher.ActualValue)
}

type IsTrue struct {
	ActualValue bool
}

func (matcher IsTrue) IsSatisfied() bool {
	return matcher.ActualValue == true
}
func (matcher IsTrue) ErrorMessage() string {
	return fmt.Sprintf("expected true but got false")
}
func (matcher IsTrue) Expected() string {
	return strconv.FormatBool(true)
}
func (matcher IsTrue) Actual() string {
	return strconv.FormatBool(matcher.ActualValue)
}

type IsFalse struct {
	ActualValue bool
}

func (matcher IsFalse) IsSatisfied() bool {
	return matcher.ActualValue == false
}
func (matcher IsFalse) ErrorMessage() string {
	return fmt.Sprintf("expected false but got true")
}
func (matcher IsFalse) Expected() string {
	return strconv.FormatBool(false)
}
func (matcher IsFalse) Actual() string {
	return strconv.FormatBool(matcher.ActualValue)
}

// Add red terminal ANSI color
func red(str string) string {
	return "\033[31m\033[1m" + str + "\033[0m"
}

// Add green terminal ANSI color
func green(str string) string {
	return "\033[32m\033[1m" + str + "\033[0m"
}

// Add grey terminal ANSI color
func grey(str string) string {
	return "\x1B[90m" + str + "\033[0m"
}
