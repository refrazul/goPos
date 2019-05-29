package interpreter

import "testing"


func TestSuma(t *testing.T) {
	cadena := []byte{49, 57, 78}
	expected := "31394E"

	res := Interpret(cadena, 0)

	if expected != string(res) {
        t.Errorf("got '%s' expected '%s'", expected, string(res))
    }
}