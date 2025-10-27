package parser

import (
	"strings"
	"testing"

	"github.com/weaweawe01/ParserOgnl/lexer"
)

// TestMaxParseIterations 测试最大解析迭代次数保护机制
func TestMaxParseIterations(t *testing.T) {
	t.Run("Normal expression within limit", func(t *testing.T) {
		// 正常的链式表达式，不会超过限制
		input := "a.b.c.d.e.f.g.h.i.j"
		l := lexer.NewLexer(input)
		p := New(l)

		expr, err := p.ParseTopLevelExpression()
		if err != nil {
			t.Errorf("Expected successful parsing, got error: %v", err)
		}
		if expr == nil {
			t.Error("Expected non-nil expression")
		}
		if p.iterationCount >= MaxParseIterations {
			t.Errorf("Iteration count %d should be much less than limit %d", p.iterationCount, MaxParseIterations)
		}
		t.Logf("Normal expression used %d iterations", p.iterationCount)
	})

	t.Run("Very long chain expression", func(t *testing.T) {
		// 生成一个非常长的链式表达式（但仍在合理范围内）
		var parts []string
		for i := 0; i < 1000; i++ {
			parts = append(parts, "prop")
		}
		input := "obj." + strings.Join(parts, ".")

		l := lexer.NewLexer(input)
		p := New(l)

		expr, err := p.ParseTopLevelExpression()
		if err != nil {
			t.Errorf("Expected successful parsing, got error: %v", err)
		}
		if expr == nil {
			t.Error("Expected non-nil expression")
		}
		if p.iterationCount >= MaxParseIterations {
			t.Errorf("Iteration count %d exceeded limit %d", p.iterationCount, MaxParseIterations)
		}
		t.Logf("Long chain (1000 properties) used %d iterations", p.iterationCount)
	})

	t.Run("Complex nested expression", func(t *testing.T) {
		// 复杂的嵌套表达式
		var builder strings.Builder
		builder.WriteString("obj")
		for i := 0; i < 100; i++ {
			builder.WriteString(".method()")
			builder.WriteString(".field")
			builder.WriteString("[0]")
		}
		input := builder.String()

		l := lexer.NewLexer(input)
		p := New(l)

		expr, err := p.ParseTopLevelExpression()
		if err != nil {
			t.Errorf("Expected successful parsing, got error: %v", err)
		}
		if expr == nil {
			t.Error("Expected non-nil expression")
		}
		t.Logf("Complex nested expression (100 levels) used %d iterations", p.iterationCount)
	})

	t.Run("Extremely long chain near limit", func(t *testing.T) {
		// 生成一个接近但不超过限制的超长链
		// 每个属性访问大约需要 2-3 次迭代
		// 所以大约 5000-6000 个属性应该接近但不超过 20000 次迭代
		var parts []string
		for i := 0; i < 5000; i++ {
			parts = append(parts, "p")
		}
		input := "x." + strings.Join(parts, ".")

		l := lexer.NewLexer(input)
		p := New(l)

		expr, err := p.ParseTopLevelExpression()
		if err != nil {
			t.Errorf("Expected successful parsing, got error: %v", err)
		}
		if expr == nil {
			t.Error("Expected non-nil expression")
		}
		t.Logf("Extremely long chain (5000 properties) used %d iterations (limit: %d)",
			p.iterationCount, MaxParseIterations)

		// 验证接近但未超过限制
		if p.iterationCount >= MaxParseIterations {
			t.Errorf("Iteration count %d exceeded limit %d", p.iterationCount, MaxParseIterations)
		}
		if p.iterationCount < MaxParseIterations/2 {
			t.Logf("Warning: Iteration count %d is far from limit, test may not be challenging enough",
				p.iterationCount)
		}
	})

	t.Run("Multiple projection operations", func(t *testing.T) {
		// 多个投影操作
		var builder strings.Builder
		builder.WriteString("list")
		for i := 0; i < 50; i++ {
			builder.WriteString(".{#this.field}")
		}
		input := builder.String()

		l := lexer.NewLexer(input)
		p := New(l)

		expr, err := p.ParseTopLevelExpression()
		if err != nil {
			t.Errorf("Expected successful parsing, got error: %v", err)
		}
		if expr == nil {
			t.Error("Expected non-nil expression")
		}
		t.Logf("Multiple projections (50 times) used %d iterations", p.iterationCount)
	})
}

// TestIterationCounterReset 测试迭代计数器在每次解析时重置
func TestIterationCounterReset(t *testing.T) {
	input := "a.b.c.d.e"

	// 第一次解析
	l1 := lexer.NewLexer(input)
	p1 := New(l1)
	_, err1 := p1.ParseTopLevelExpression()
	if err1 != nil {
		t.Errorf("First parse failed: %v", err1)
	}
	count1 := p1.iterationCount
	t.Logf("First parse used %d iterations", count1)

	// 第二次解析（新的 Parser 实例）
	l2 := lexer.NewLexer(input)
	p2 := New(l2)
	_, err2 := p2.ParseTopLevelExpression()
	if err2 != nil {
		t.Errorf("Second parse failed: %v", err2)
	}
	count2 := p2.iterationCount
	t.Logf("Second parse used %d iterations", count2)

	// 两次解析应该使用相同数量的迭代
	if count1 != count2 {
		t.Errorf("Iteration counts differ: first=%d, second=%d", count1, count2)
	}
}

// BenchmarkParseChainLength 基准测试：不同长度链式表达式的性能
func BenchmarkParseChainLength(b *testing.B) {
	lengths := []int{10, 50, 100, 500, 1000}

	for _, length := range lengths {
		// 构造链式表达式
		var parts []string
		for i := 0; i < length; i++ {
			parts = append(parts, "prop")
		}
		input := "obj." + strings.Join(parts, ".")

		b.Run(string(rune(length))+" properties", func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				l := lexer.NewLexer(input)
				p := New(l)
				_, err := p.ParseTopLevelExpression()
				if err != nil {
					b.Fatalf("Parse error: %v", err)
				}
			}
		})
	}
}

// TestIterationLimitEffectiveness 测试迭代限制的有效性
func TestIterationLimitEffectiveness(t *testing.T) {
	t.Run("Iteration counter increases during parsing", func(t *testing.T) {
		input := "a.b.c.d.e.f.g.h.i.j.k.l.m.n.o.p"
		l := lexer.NewLexer(input)
		p := New(l)

		// 解析前计数器应该为 0
		if p.iterationCount != 0 {
			t.Errorf("Initial iteration count should be 0, got %d", p.iterationCount)
		}

		expr, err := p.ParseTopLevelExpression()
		if err != nil {
			t.Errorf("Parse error: %v", err)
		}
		if expr == nil {
			t.Error("Expected non-nil expression")
		}

		// 解析后计数器应该增加
		if p.iterationCount == 0 {
			t.Error("Iteration count should be > 0 after parsing")
		}
		t.Logf("Parse completed with %d iterations", p.iterationCount)
	})

	t.Run("Check iteration limit in navigation loop", func(t *testing.T) {
		// 创建一个会进入 navigationLoop 的表达式
		input := "obj.a.b.c.d.e.f.g.h.i.j"
		l := lexer.NewLexer(input)
		p := New(l)

		expr, err := p.ParseTopLevelExpression()
		if err != nil {
			t.Errorf("Parse error: %v", err)
		}
		if expr == nil {
			t.Error("Expected non-nil expression")
		}

		// 验证迭代计数
		if p.iterationCount == 0 {
			t.Error("Expected positive iteration count")
		}
		if p.iterationCount >= MaxParseIterations {
			t.Errorf("Iteration count %d exceeded limit %d", p.iterationCount, MaxParseIterations)
		}

		t.Logf("Navigation loop test used %d iterations (limit: %d)",
			p.iterationCount, MaxParseIterations)
	})
}
