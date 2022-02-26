package main

import "testing"

func Test_trimDuplicateSpaces(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"empty string", "", ""},
		{"single character", " ", " "},
		{"beginning", "     word", " word"},
		{"end", "word       ", "word "},
		{"both ends", "    word    ", " word "},
		{"single space", " word ", " word "},
		{"tab", "\tword", " word"},
		{"newline", "\nword", " word"},
		{"vertical tab", "\vword", " word"},
		{"form field", "\fword", " word"},
		{"carriage return", "\rword", " word"},
		{"combination", "\r\f\t\nword", " word"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := trimDuplicateSpaces(tt.in)
			if got != tt.want {
				t.Errorf("\nGot: \t%#v;\nWant: \t%#v", got, tt.want)
			}
		})
	}
}
