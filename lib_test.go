package makeweb

import (
	"reflect"
	"testing"
)

func TestJoinmaps(t *testing.T) {
	// TODO: more tests
	map1 := map[string]interface{}{"1": 1, "test": "a"}
	map2 := map[string]interface{}{"test": "b"}
	result, err := joinmaps(map1, map2)
	if err != nil {
		t.Fail()
	}
	if !reflect.DeepEqual(result, map[string]interface{}{"1": 1, "test": "b"}) {
		t.Fail()
	}

}

func TestJoinmapsEmpty(t *testing.T) {
	map1 := map[string]interface{}{}
	map2 := map[string]interface{}{}
	result, err := joinmaps(map1, map2)
	if err != nil {
		t.Fail()
	}
	if !reflect.DeepEqual(result, map[string]interface{}{}) {
		t.Fail()
	}
}
func TestJoinmapsOneEmpty(t *testing.T) {
	map1 := map[string]interface{}{}
	map2 := map[string]interface{}{"test": "b"}
	result, err := joinmaps(map1, map2)
	if err != nil {
		t.Fail()
	}
	if !reflect.DeepEqual(result, map[string]interface{}{"test": "b"}) {
		t.Fail()
	}
}

func TestJoinmapsRecursive(t *testing.T) {
	map1 := map[string]interface{}{"test1": "a"}
	map3 := map[string]interface{}{"test2": 3}
	map2 := map[string]interface{}{"test": map3}
	result, err := joinmaps(map1, map2)
	if err != nil {
		t.Fail()
	}
	if !reflect.DeepEqual(result, map[string]interface{}{"test1": "a", "test": map[string]interface{}{"test2": 3}}) {
		t.Fail()
	}
}

func TestSplitVarsContent(t *testing.T) {
	vars, content := splitVarsContent("a\n---\nb")
	if vars != "a" {
		t.Fail()
	}
	if content != "b" {
		t.Fail()
	}
}
func TestSplitVarsContent2(t *testing.T) {
	vars, content := splitVarsContent("a--b\n---\n")
	if vars != "a--b" {
		t.Fail()
	}
	if content != "" {
		t.Fail()
	}
}
func TestSplitVarsContentNoSeparator(t *testing.T) {
	vars, content := splitVarsContent("a--b--")
	if vars != "{}" {
		t.Fail()
	}
	if content != "a--b--" {
		t.Fail()
	}
}
func TestSplitVarsContentTwoSeparators(t *testing.T) {
	vars, content := splitVarsContent("a\n---\nb\n---\nc")
	if vars != "a" {
		t.Fail()
	}
	if content != "b\n---\nc" {
		t.Fail()
	}
}
func TestRecursiveLs(t *testing.T) {
	result, err := recursiveLs("test/recursivels")
	if err != nil {
		t.Fail()
	}
	if !reflect.DeepEqual(result, []string{"test/recursivels/1", "test/recursivels/t/2"}) {
		t.Fail()
	}
}
func TestRecursiveLsNonexistentDir(t *testing.T) {
	result, err := recursiveLs("ablabfdljv,abvhjb")
	if err == nil {
		t.Fail()
	}
	if result != nil {
		t.Fail()
	}
}
func TestExists(t *testing.T) {
	result, err := exists("test/file")
	if result != true {
		t.Fail()
	}
	if err != nil {
		t.Fail()
	}
}
func TestExistsDir(t *testing.T) {
	result, err := exists("test/folder")
	if result != true {
		t.Fail()
	}
	if err != nil {
		t.Fail()
	}
}
func TestExistsNonexistent(t *testing.T) {
	result, err := exists("nflandfblabvlajldbl")
	if result == true {
		t.Fail()
	}
	if err != nil {
		t.Fail()
	}
}
