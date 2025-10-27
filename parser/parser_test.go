package parser

import (
	"testing"

	"github.com/weaweawe01/ParserOgnl/ast"
	"github.com/weaweawe01/ParserOgnl/lexer"
)

func TestParseSimpleExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"42", "42"},
		{"true", "true"},
		{"false", "false"},
		{"null", "null"},
		{"\"hello\"", "\"hello\""},
		{"'c'", "'c'"},
		{"name", "name"},
		{"#this", "#this"},
		{"#root", "#root"},
		{"#variable", "#variable"},
	}

	for _, tt := range tests {
		l := lexer.NewLexer(tt.input)
		p := New(l)
		expr, err := p.ParseTopLevelExpression()
		if err != nil {
			t.Fatalf("ParseTopLevelExpression() returned error for input: %s, error: %v", tt.input, err)
		}

		if expr == nil {
			t.Fatalf("ParseTopLevelExpression() returned nil for input: %s", tt.input)
		}

		if len(p.Errors()) != 0 {
			t.Fatalf("parser has %d errors for input '%s': %v",
				len(p.Errors()), tt.input, p.Errors())
		}

		if expr.String() != tt.expected {
			t.Errorf("expr.String() wrong. expected=%q, got=%q",
				tt.expected, expr.String())
		}
	}
}

func TestParseBinaryExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"1 + 2", "(1 + 2)"},
		{"1 - 2", "(1 - 2)"},
		{"1 * 2", "(1 * 2)"},
		{"1 / 2", "(1 / 2)"},
		{"1 % 2", "(1 % 2)"},
		{"1 == 2", "(1 == 2)"},
		{"1 != 2", "(1 != 2)"},
		{"1 < 2", "(1 < 2)"},
		{"1 > 2", "(1 > 2)"},
		{"1 <= 2", "(1 <= 2)"},
		{"1 >= 2", "(1 >= 2)"},
		{"1 && 2", "(1 && 2)"},
		{"1 || 2", "(1 || 2)"},
		{"1 & 2", "(1 & 2)"},
		{"1 | 2", "(1 | 2)"},
		{"1 ^ 2", "(1 ^ 2)"},
		{"1 << 2", "(1 << 2)"},
		{"1 >> 2", "(1 >> 2)"},
		{"1 >>> 2", "(1 >>> 2)"},
	}

	for _, tt := range tests {
		l := lexer.NewLexer(tt.input)
		p := New(l)
		expr, err := p.ParseTopLevelExpression()
		if err != nil {
			t.Fatalf("ParseTopLevelExpression() returned error for input: %s, error: %v", tt.input, err)
		}

		if expr == nil {
			t.Fatalf("ParseTopLevelExpression() returned nil for input: %s", tt.input)
		}

		if len(p.Errors()) != 0 {
			t.Fatalf("parser has %d errors for input '%s': %v",
				len(p.Errors()), tt.input, p.Errors())
		}

		if expr.String() != tt.expected {
			t.Errorf("expr.String() wrong. expected=%q, got=%q",
				tt.expected, expr.String())
		}
	}
}

func TestParseUnaryExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"-42", "(-42)"},
		{"+42", "(+42)"},
		{"!true", "(!true)"},
		{"~1", "(~1)"},
		{"!!true", "(!(true))"},
		{"-(-42)", "(-((-42)))"},
	}

	for _, tt := range tests {
		l := lexer.NewLexer(tt.input)
		p := New(l)
		expr, err := p.ParseTopLevelExpression()
		if err != nil {
			t.Fatalf("ParseTopLevelExpression() returned error for input: %s, error: %v", tt.input, err)
		}
		if expr == nil {
			t.Fatalf("ParseTopLevelExpression() returned nil for input: %s", tt.input)
		}

		if len(p.Errors()) != 0 {
			t.Fatalf("parser has %d errors for input '%s': %v",
				len(p.Errors()), tt.input, p.Errors())
		}

		if expr.String() != tt.expected {
			t.Errorf("expr.String() wrong. expected=%q, got=%q",
				tt.expected, expr.String())
		}
	}
}

func TestParseAssignmentExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"a = 1", "(a = 1)"},
		{"a = b = 2", "(a = (b = 2))"},
		{"x = y + z", "(x = (y + z))"},
	}

	for _, tt := range tests {
		l := lexer.NewLexer(tt.input)
		p := New(l)
		expr, err := p.ParseTopLevelExpression()
		if err != nil {
			t.Fatalf("ParseTopLevelExpression() returned error for input: %s, error: %v", tt.input, err)
		}
		if expr == nil {
			t.Fatalf("ParseTopLevelExpression() returned nil for input: %s", tt.input)
		}

		if len(p.Errors()) != 0 {
			t.Fatalf("parser has %d errors for input '%s': %v",
				len(p.Errors()), tt.input, p.Errors())
		}

		if expr.String() != tt.expected {
			t.Errorf("expr.String() wrong. expected=%q, got=%q",
				tt.expected, expr.String())
		}
	}
}

func TestParseConditionalExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"a ? b : c", "(a ? b : c)"},
		{"x > 0 ? x : -x", "((x > 0) ? x : (-x))"},
		{"true ? 1 : false ? 2 : 3", "(true ? 1 : (false ? 2 : 3))"},
	}

	for _, tt := range tests {
		l := lexer.NewLexer(tt.input)
		p := New(l)
		expr, err := p.ParseTopLevelExpression()
		if err != nil {
			t.Fatalf("ParseTopLevelExpression() returned error for input: %s, error: %v", tt.input, err)
		}
		if expr == nil {
			t.Fatalf("ParseTopLevelExpression() returned nil for input: %s", tt.input)
		}

		if len(p.Errors()) != 0 {
			t.Fatalf("parser has %d errors for input '%s': %v",
				len(p.Errors()), tt.input, p.Errors())
		}

		if expr.String() != tt.expected {
			t.Errorf("expr.String() wrong. expected=%q, got=%q",
				tt.expected, expr.String())
		}
	}
}

func TestParseSequenceExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"a, b", "a, b"},
		{"1, 2, 3", "1, 2, 3"},
		{"a = 1, b = 2", "(a = 1), (b = 2)"},
	}

	for _, tt := range tests {
		l := lexer.NewLexer(tt.input)
		p := New(l)
		expr, err := p.ParseTopLevelExpression()
		if err != nil {
			t.Fatalf("ParseTopLevelExpression() returned error for input: %s, error: %v", tt.input, err)
		}
		if expr == nil {
			t.Fatalf("ParseTopLevelExpression() returned nil for input: %s", tt.input)
		}

		if len(p.Errors()) != 0 {
			t.Fatalf("parser has %d errors for input '%s': %v",
				len(p.Errors()), tt.input, p.Errors())
		}

		if expr.String() != tt.expected {
			t.Errorf("expr.String() wrong. expected=%q, got=%q",
				tt.expected, expr.String())
		}
	}
}

func TestParseChainExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"a.b", "a.b"},
		{"a.b.c", "a.b.c"},
		{"obj.method()", "obj.method()"},
		{"obj.method(1, 2)", "obj.method(1, 2)"},
	}

	for _, tt := range tests {
		l := lexer.NewLexer(tt.input)
		p := New(l)
		expr, err := p.ParseTopLevelExpression()
		if err != nil {
			t.Fatalf("ParseTopLevelExpression() returned error for input: %s, error: %v", tt.input, err)
		}
		if expr == nil {
			t.Fatalf("ParseTopLevelExpression() returned nil for input: %s", tt.input)
		}

		if len(p.Errors()) != 0 {
			t.Fatalf("parser has %d errors for input '%s': %v",
				len(p.Errors()), tt.input, p.Errors())
		}

		if expr.String() != tt.expected {
			t.Errorf("expr.String() wrong. expected=%q, got=%q",
				tt.expected, expr.String())
		}
	}
}

func TestParseInstanceofExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"obj instanceof String", "(obj instanceof String)"},
		{"x instanceof java.util.List", "(x instanceof java.util.List)"},
	}

	for _, tt := range tests {
		l := lexer.NewLexer(tt.input)
		p := New(l)
		expr, err := p.ParseTopLevelExpression()
		if err != nil {
			t.Fatalf("ParseTopLevelExpression() returned error for input: %s, error: %v", tt.input, err)
		}
		if expr == nil {
			t.Fatalf("ParseTopLevelExpression() returned nil for input: %s", tt.input)
		}

		if len(p.Errors()) != 0 {
			t.Fatalf("parser has %d errors for input '%s': %v",
				len(p.Errors()), tt.input, p.Errors())
		}

		if expr.String() != tt.expected {
			t.Errorf("expr.String() wrong. expected=%q, got=%q",
				tt.expected, expr.String())
		}
	}
}

func TestParseConstructorExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"new String()", "new String()"},
		{"new String(\"hello\")", "new String(\"hello\")"},
		{"new int[]{1, 2, 3}", "new int[]{ 1, 2, 3 }"},
		{"new java.util.ArrayList(10)", "new java.util.ArrayList(10)"},
	}

	for _, tt := range tests {
		l := lexer.NewLexer(tt.input)
		p := New(l)
		expr, err := p.ParseTopLevelExpression()
		if err != nil {
			t.Fatalf("ParseTopLevelExpression() returned error for input: %s, error: %v", tt.input, err)
		}
		if expr == nil {
			t.Fatalf("ParseTopLevelExpression() returned nil for input: %s", tt.input)
		}

		if len(p.Errors()) != 0 {
			t.Fatalf("parser has %d errors for input '%s': %v",
				len(p.Errors()), tt.input, p.Errors())
		}

		if expr.String() != tt.expected {
			t.Errorf("expr.String() wrong. expected=%q, got=%q",
				tt.expected, expr.String())
		}
	}
}

func TestParseStaticExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"@Math@PI", "@java.lang.Math@PI"},
		{"@Math@max(1, 2)", "@java.lang.Math@max(1, 2)"},
		{"@java.lang.System@out", "@java.lang.System@out"},
	}

	for _, tt := range tests {
		l := lexer.NewLexer(tt.input)
		p := New(l)
		expr, err := p.ParseTopLevelExpression()
		if err != nil {
			t.Fatalf("ParseTopLevelExpression() returned error for input: %s, error: %v", tt.input, err)
		}
		if expr == nil {
			t.Fatalf("ParseTopLevelExpression() returned nil for input: %s", tt.input)
		}

		if len(p.Errors()) != 0 {
			t.Fatalf("parser has %d errors for input '%s': %v",
				len(p.Errors()), tt.input, p.Errors())
		}

		if expr.String() != tt.expected {
			t.Errorf("expr.String() wrong. expected=%q, got=%q",
				tt.expected, expr.String())
		}
	}
}

func TestParseArrayExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"{}", "{ }"},
		{"{1}", "{ 1 }"},
		{"{1, 2, 3}", "{ 1, 2, 3 }"},
		{"{\"a\", \"b\"}", "{ \"a\", \"b\" }"},
	}

	for _, tt := range tests {
		l := lexer.NewLexer(tt.input)
		p := New(l)
		expr, err := p.ParseTopLevelExpression()
		if err != nil {
			t.Fatalf("ParseTopLevelExpression() returned error for input: %s, error: %v", tt.input, err)
		}
		if expr == nil {
			t.Fatalf("ParseTopLevelExpression() returned nil for input: %s", tt.input)
		}

		if len(p.Errors()) != 0 {
			t.Fatalf("parser has %d errors for input '%s': %v",
				len(p.Errors()), tt.input, p.Errors())
		}

		if expr.String() != tt.expected {
			t.Errorf("expr.String() wrong. expected=%q, got=%q",
				tt.expected, expr.String())
		}
	}
}

func TestParseMapExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"#{\"key\": \"value\"}", "#{ \"key\" : \"value\" }"},
		{"#{\"a\": 1, \"b\": 2}", "#{ \"a\" : 1, \"b\" : 2 }"},
		{"#{key: value}", "#{ key : value }"},
	}

	for _, tt := range tests {
		l := lexer.NewLexer(tt.input)
		p := New(l)
		expr, err := p.ParseTopLevelExpression()
		if err != nil {
			t.Fatalf("ParseTopLevelExpression() returned error for input: %s, error: %v", tt.input, err)
		}
		if expr == nil {
			t.Fatalf("ParseTopLevelExpression() returned nil for input: %s", tt.input)
		}

		if len(p.Errors()) != 0 {
			t.Fatalf("parser has %d errors for input '%s': %v",
				len(p.Errors()), tt.input, p.Errors())
		}

		if expr.String() != tt.expected {
			t.Errorf("expr.String() wrong. expected=%q, got=%q",
				tt.expected, expr.String())
		}
	}
}

func TestParseComplexExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"1 + 2 * 3", "(1 + (2 * 3))"},
		{"(1 + 2) * 3", "((1 + 2) * 3)"},
		{"a.b.c.method(x, y)", "a.b.c.method(x, y)"},
		{"list.{? #this > 5}", "list.{? (#this > 5)}"},
		{"array.{#this * 2}", "array.{(#this * 2)}"},
		{"obj instanceof String ? obj.length() : 0", "((obj instanceof String) ? obj.length() : 0)"},
		{"x = y > 0 ? y : -y", "(x = ((y > 0) ? y : (-y)))"},
	}

	for _, tt := range tests {
		l := lexer.NewLexer(tt.input)
		p := New(l)
		expr, err := p.ParseTopLevelExpression()
		if err != nil {
			t.Fatalf("ParseTopLevelExpression() returned error for input: %s, error: %v", tt.input, err)
		}
		if expr == nil {
			t.Fatalf("ParseTopLevelExpression() returned nil for input: %s", tt.input)
		}

		if len(p.Errors()) != 0 {
			t.Fatalf("parser has %d errors for input '%s': %v",
				len(p.Errors()), tt.input, p.Errors())
		}

		if expr.String() != tt.expected {
			t.Errorf("expr.String() wrong. expected=%q, got=%q",
				tt.expected, expr.String())
		}
	}
}

// Benchmark tests
func BenchmarkParseSimpleExpression(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := lexer.NewLexer("a + b * c")
		p := New(l)
		p.ParseTopLevelExpression()
	}
}

func BenchmarkParseComplexExpression(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := lexer.NewLexer("obj.method(x, y).property.{? #this > 0}")
		p := New(l)
		p.ParseTopLevelExpression()
	}
}

// Helper function to check AST node types
func testLiteralExpression(t *testing.T, exp ast.Expression, expected interface{}) bool {
	switch v := expected.(type) {
	case int:
		return testIntegerLiteral(t, exp, int64(v))
	case int64:
		return testIntegerLiteral(t, exp, v)
	case float64:
		return testFloatLiteral(t, exp, v)
	case string:
		return testStringLiteral(t, exp, v)
	case bool:
		return testBooleanLiteral(t, exp, v)
	}
	t.Errorf("type of exp not handled. got=%T", exp)
	return false
}

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	literal, ok := il.(*ast.Literal)
	if !ok {
		t.Errorf("il not *ast.Literal. got=%T", il)
		return false
	}

	intVal, ok := literal.Value.(int64)
	if !ok {
		t.Errorf("literal.Value not int64. got=%T", literal.Value)
		return false
	}

	if intVal != value {
		t.Errorf("literal.Value not %d. got=%d", value, intVal)
		return false
	}

	return true
}

func testFloatLiteral(t *testing.T, fl ast.Expression, value float64) bool {
	literal, ok := fl.(*ast.Literal)
	if !ok {
		t.Errorf("fl not *ast.Literal. got=%T", fl)
		return false
	}

	floatVal, ok := literal.Value.(float64)
	if !ok {
		t.Errorf("literal.Value not float64. got=%T", literal.Value)
		return false
	}

	if floatVal != value {
		t.Errorf("literal.Value not %f. got=%f", value, floatVal)
		return false
	}

	return true
}

func testStringLiteral(t *testing.T, sl ast.Expression, value string) bool {
	literal, ok := sl.(*ast.Literal)
	if !ok {
		t.Errorf("sl not *ast.Literal. got=%T", sl)
		return false
	}

	strVal, ok := literal.Value.(string)
	if !ok {
		t.Errorf("literal.Value not string. got=%T", literal.Value)
		return false
	}

	if strVal != value {
		t.Errorf("literal.Value not %q. got=%q", value, strVal)
		return false
	}

	return true
}

func testBooleanLiteral(t *testing.T, bl ast.Expression, value bool) bool {
	literal, ok := bl.(*ast.Literal)
	if !ok {
		t.Errorf("bl not *ast.Literal. got=%T", bl)
		return false
	}

	boolVal, ok := literal.Value.(bool)
	if !ok {
		t.Errorf("literal.Value not bool. got=%T", literal.Value)
		return false
	}

	if boolVal != value {
		t.Errorf("literal.Value not %t. got=%t", value, boolVal)
		return false
	}

	return true
}

func testIdentifier(t *testing.T, exp ast.Expression, value string) bool {
	ident, ok := exp.(*ast.Identifier)
	if !ok {
		t.Errorf("exp not *ast.Identifier. got=%T", exp)
		return false
	}

	if ident.Value != value {
		t.Errorf("ident.Value not %s. got=%s", value, ident.Value)
		return false
	}

	return true
}
