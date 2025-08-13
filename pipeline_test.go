package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestPythonHelloWorld(t *testing.T) {
	pyCode := []byte("print('Hello, World!')\n")
	pyFile := "hello.py"
	if err := os.WriteFile(pyFile, pyCode, 0644); err != nil {
		t.Fatalf("Failed to write Python file: %v", err)
	}
	defer os.Remove(pyFile)

	cmd := exec.Command("python3", pyFile)
	if err := cmd.Run(); err != nil {
		t.Fatalf("Python Hello World failed: %v", err)
	}
}

func TestZigHelloWorld(t *testing.T) {
	zigCode := []byte(`const std = @import("std");
pub fn main() void {
    std.debug.print("Hello, World!\n", .{});
}`)
	zigFile := "hello.zig"
	if err := os.WriteFile(zigFile, zigCode, 0644); err != nil {
		t.Fatalf("Failed to write Zig file: %v", err)
	}
	defer os.Remove(zigFile)

	cmd := exec.Command("zig", "run", zigFile)
	if err := cmd.Run(); err != nil {
		t.Fatalf("Zig Hello World failed: %v", err)
	}
}
