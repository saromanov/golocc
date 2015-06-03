package main

import (
	"testing"
)

func TestCLOC(t *testing.T) {

	parser := Parser{}
	result := parser.ParseDir("./fixture")

	if result.CLOC != 8 {
		t.Error("expected 8 comment lines")
	}
}

func TestStructCount(t *testing.T) {

	parser := Parser{}
	result := parser.ParseDir("./fixture")

	if result.Struct != 2 {
		t.Error("expected 2 structs")
	}
}

func TestInterfaceCount(t *testing.T) {

	parser := Parser{}
	result := parser.ParseDir("./fixture")

	if result.Interface != 1 {
		t.Error("expected 1 interface")
	}
}

func TestMethodCount(t *testing.T) {

	parser := Parser{}
	result := parser.ParseDir("./fixture")

	if result.Method != 2 {
		t.Error("expected 2 methods")
	}

	if result.ExportedMethod != 1 {
		t.Error("expected 1 exported method")
	}
}

func TestFunctionCount(t *testing.T) {

	parser := Parser{}
	result := parser.ParseDir("./fixture")

	if result.Function != 2 {
		t.Error("expected 2 function")
	}

	if result.ExportedFunction != 1 {
		t.Error("expected 1 exported function")
	}
}

func TestImportCount(t *testing.T) {

	parser := Parser{}
	result := parser.ParseDir("./fixture")

	if result.Import != 2 {
		t.Error("expected 2 imports got", result.Import)
	}
}

