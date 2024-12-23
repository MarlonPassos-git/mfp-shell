package main

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	tests := []struct {
		input        string
		expectedCmd  string
		expectedArgs []string
	}{
		// {"cmd arg1 arg2", "cmd", []string{"arg1", "arg2"}},
		// {"echo world\\ \\ \\ \\ \\ \\ script", "echo", []string{"world      script"}},
		// {"echo world\nexample", "echo", []string{"worldnexample"}},
		{"echo echo\\world", "echo", []string{"testnhello"}},
		// {"cmd 'arg with spaces'", "cmd", []string{"arg with spaces"}},
		// {"cmd \"arg with spaces\"", "cmd", []string{"arg with spaces"}},
		// {"cmd 'arg1' 'arg2'", "cmd", []string{"arg1", "arg2"}},
		// {"cmd \"arg1\" \"arg2\"", "cmd", []string{"arg1", "arg2"}},
	}

	for _, test := range tests {
		cmd, args := ParseInput(test.input)
		if cmd != test.expectedCmd || !reflect.DeepEqual(args, test.expectedArgs) {
			t.Errorf("ParseInput(%q) = %q, %q; want %q, %q", test.input, cmd, args, test.expectedCmd, test.expectedArgs)
		}
	}
}
