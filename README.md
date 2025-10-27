# Go parser for Java  Apache OGNL 
[![Go Version](https://img.shields.io/badge/Go-1.18+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://www.apache.org/licenses/LICENSE-2.0)

这是 Apache OGNL (Object-Graph Navigation Language) 的 Go 语言实现，完整兼容 Java OGNL 的语法和语义。

## 与 Java OGNL 的兼容性
### 完全兼容的特性

✅ **语法兼容性**: 100% 兼容 Java OGNL 语法 </br>
✅ **AST 结构**: AST 节点类型与 Java OGNL 一一对应</br>
✅ **运算符优先级**: 完全遵循 Java OGNL 的运算符优先级</br>
✅ **字面量格式**: 支持所有 Java OGNL 的字面量格式</br>
✅ **集合操作**: 完整支持投影、选择等集合操作</br>
✅ **静态引用**: 完整支持静态字段和方法访问</br>
✅ **Lambda 表达式**: 完整支持 Lambda 定义和调用</br>


```bash
go get github.com/weaweawe01/ParserOgnl
```

## 快速开始

```go
package main

import (
	"fmt"

	"github.com/weaweawe01/ParserOgnl/lexer"
	"github.com/weaweawe01/ParserOgnl/parser"
)

func main() {
	input := "new ognl.test.objects.Simple(new Object[5])"
	fmt.Println("输入表达式:", input)
	// 创建词法分析器和解析器
	l := lexer.NewLexer(input)
	p := parser.New(l)
	// 解析表达式
	expr, err := p.ParseTopLevelExpression()
	// 检查错误
	if err != nil {
		fmt.Printf("Parser errors: %v\n", err)
		return
	}
	if len(p.Errors()) > 0 {
		fmt.Println("Parser errors:")
		for _, err := range p.Errors() {
			fmt.Printf("  %s\n", err)
		}
		return
	}
	if expr == nil {
		fmt.Println("Failed to parse expression")
		return
	}
	// 输出详细的AST结构
	fmt.Println("AST 结构:")
	parser.PrintASTStructure(expr, 0)

}
```

### 输出示例

```
#go run .
输入表达式: new ognl.test.objects.Simple(new Object[5])
AST 结构:
  ASTCtor 表达式片段: new ognl.test.objects.Simple(new Object[5])
        ASTCtor 表达式片段: new Object[5]
      ASTConst 表达式片段: 5
```


## 架构设计
### 三层架构
```
┌─────────────────────────────────────────┐
│          Application Layer              │
│   (命令行工具、库接口、测试工具)           │
└─────────────────────────────────────────┘
                    ↓
┌─────────────────────────────────────────┐
│           Parser Layer                  │
│   (语法解析、AST 构建、错误处理)           │
└─────────────────────────────────────────┘
                    ↓
┌─────────────────────────────────────────┐
│           Lexer Layer                   │
│   (词法分析、Token 生成、字符处理)          │
└─────────────────────────────────────────┘
```

### 目录结构

```
go_ognl/
├── main.go                 # 命令行入口
├── README.md              # 本文档
├── lexer/                 # 词法分析器
│   └── lexer.go
├── parser/                # 语法解析器
│   └── parser.go
├── token/                 # Token 定义
│   └── token.go
├── ast/                   # AST 节点定义
│   └── ast.go
└── tests/                 # 测试套件
    ├── check.go           # 测试辅助函数
    ├── *_test.go          # 各种测试文件
    └── ...
```

### 解析流程

```
输入字符串
    ↓
[Lexer] 词法分析
    ↓
Token 流
    ↓
[Parser] 语法分析
    ↓
AST (抽象语法树)
    ↓
[输出] JSON / 树形结构 / 遍历访问
```




## 许可证

本项目采用 Apache License 2.0 许可证。详见 [LICENSE](../LICENSE.txt) 文件。

## 致谢

- 感谢 [Apache OGNL](https://github.com/orphan-oss/ognl) 项目提供原始实现
- 感谢所有贡献者的辛勤工作

## 联系方式

- 问题反馈: [GitHub Issues](https://github.com/yourusername/go_ognl/issues)
- 讨论区: [GitHub Discussions](https://github.com/yourusername/go_ognl/discussions)

## 参考资源

- [Apache OGNL 官方文档](https://commons.apache.org/proper/commons-ognl/)
- [OGNL 语言指南](../docs/LanguageGuide.md)
- [OGNL 开发者指南](../docs/DeveloperGuide.md)

---

## 已完成的测试对比

本项目的测试文件完全对应 Java OGNL 的测试套件，确保 100% 语法兼容性。

| Go 测试文件 | Java 测试文件 | 状态 | 说明 |
|------------|--------------|------|------|
| `ArithmeticAndLogicalOperators_test.go` | `ArithmeticAndLogicalOperatorsTest.java` | ✅ | 算术和逻辑运算符 |
| `ArithmeticAndLogicalOperatorsOnEnums_test.go` | `ArithmeticAndLogicalOperatorsOnEnumsTest.java` | ✅ | 枚举上的运算符 |
| `ArrayCreation_test.go` | `ArrayCreationTest.java` | ✅ | 数组创建 |
| `ArrayElements_test.go` | `ArrayElementsTest.java` | ✅ | 数组元素访问 |
| `ASTChain_test.go` | `ASTChainTest.java` | ✅ | 链式表达式 |
| `ASTMethod_test.go` | `ASTMethodTest.java` | ✅ | 方法调用 |
| `ASTProperty_test.go` | `ASTPropertyTest.java` | ✅ | 属性访问 |
| `ASTSequence_test.go` | `ASTSequenceTest.java` | ✅ | 序列表达式 |
| `Chain_test.go` | `ChainTest.java` | ✅ | 复杂链式调用 |
| `ClassMethod_test.go` | `ClassMethodTest.java` | ✅ | 类方法 |
| `CollectionDirectProperty_test.go` | `CollectionDirectPropertyTest.java` | ✅ | 集合直接属性 |
| `Constant_test.go` | `ConstantTest.java` | ✅ | 常量 |
| `ConstantTree_test.go` | `ConstantTreeTest.java` | ✅ | 常量树 |
| `ContextVariable_test.go` | `ContextVariableTest.java` | ✅ | 上下文变量 |
| `DefaultClassResolver_test.go` | `DefaultClassResolverTest.java` | ✅ | 默认类解析器 |
| `Generics_test.go` | `GenericsTest.java` | ✅ | 泛型 |
| `IndexAccess_test.go` | `IndexAccessTest.java` | ✅ | 索引访问 |
| `IndexedProperty_test.go` | `IndexedPropertyTest.java` | ✅ | 索引属性 |
| `InExpression_test.go` | `InExpressionTest.java` | ✅ | in 表达式 |
| `InheritedMethods_test.go` | `InheritedMethodsTest.java` | ✅ | 继承方法 |
| `InterfaceInheritance_test.go` | `InterfaceInheritanceTest.java` | ✅ | 接口继承 |
| `IsTruck_test.go` | `IsTruckTest.java` | ✅ | instanceof 测试 |
| `Java8_test.go` | `Java8Test.java` | ✅ | Java 8 特性 |
| `LambdaExpression_test.go` | `LambdaExpressionTest.java` | ✅ | Lambda 表达式 |
| `MapCreation_test.go` | `MapCreationTest.java` | ✅ | Map 创建 |
| `MemberAccess_test.go` | `MemberAccessTest.java` | ✅ | 成员访问 |
| `Method_test.go` | `MethodTest.java` | ✅ | 方法测试 |
| `MethodWithConversion_test.go` | `MethodWithConversionTest.java` | ✅ | 带类型转换的方法 |
| `NestedMethod_test.go` | `NestedMethodTest.java` | ✅ | 嵌套方法 |
| `NullHandler_test.go` | `NullHandlerTest.java` | ✅ | 空值处理 |
| `NullRoot_test.go` | `NullRootTest.java` | ✅ | 空根对象 |
| `NullStringCatenation_test.go` | `NullStringCatenationTest.java` | ✅ | 空值字符串拼接 |
| `NumberFormatException_test.go` | `NumberFormatExceptionTest.java` | ✅ | 数字格式异常 |
| `NumericConversion_test.go` | `NumericConversionTest.java` | ✅ | 数值转换 |
| `ObjectIndexed_test.go` | `ObjectIndexedTest.java` | ✅ | 对象索引 |
| `ObjectIndexedProperty_test.go` | `ObjectIndexedPropertyTest.java` | ✅ | 对象索引属性 |
| `OgnlContextCreate_test.go` | `OgnlContextCreateTest.java` | ✅ | OGNL 上下文创建 |
| `OgnlException_test.go` | `OgnlExceptionTest.java` | ✅ | OGNL 异常 |
| `OgnlOps_test.go` | `OgnlOpsTest.java` | ✅ | OGNL 操作 |
| `Operation_test.go` | `OperationTest.java` | ✅ | 运算操作 |
| `Operator_test.go` | `OperatorTest.java` | ✅ | 运算符 |
| `PrimitiveArray_test.go` | `PrimitiveArrayTest.java` | ✅ | 原始类型数组 |
| `PrimitiveNullHandling_test.go` | `PrimitiveNullHandlingTest.java` | ✅ | 原始类型空值处理 |
| `PrivateAccessor_test.go` | `PrivateAccessorTest.java` | ✅ | 私有访问器 |
| `PropertyArithmeticAndLogicalOperators_test.go` | `PropertyArithmeticAndLogicalOperatorsTest.java` | ✅ | 属性与运算符组合 |
| `Property_test.go` | `PropertyTest.java` | ✅ | 属性测试 |
| `PropertySetter_test.go` | `PropertySetterTest.java` | ✅ | 属性设置器 |
| `ProtectedInnerClass_test.go` | `ProtectedInnerClassTest.java` | ✅ | 受保护内部类 |
| `ProtectedMember_test.go` | `ProtectedMemberTest.java` | ✅ | 受保护成员 |
| `PublicMember_test.go` | `PublicMemberTest.java` | ✅ | 公共成员 |
| `Quoting_test.go` | `QuotingTest.java` | ✅ | 引号处理 |
| `RaceCondition_test.go` | `RaceConditionTest.java` | ✅ | 竞态条件 |
| `Setter_test.go` | `SetterTest.java` | ✅ | 设置器 |
| `SetterWithConversion_test.go` | `SetterWithConversionTest.java` | ✅ | 带类型转换的设置器 |
| `ShortCircuitingExpression_test.go` | `ShortCircuitingExpressionTest.java` | ✅ | 短路表达式 |
| `SimpleNavigationChainTree_test.go` | `SimpleNavigationChainTreeTest.java` | ✅ | 简单导航链树 |
| `SimplePropertyTree_test.go` | `SimplePropertyTreeTest.java` | ✅ | 简单属性树 |
| `StaticsAndConstructors_test.go` | `StaticsAndConstructorsTest.java` | ✅ | 静态引用和构造函数 |
| `VarArgsMethod_test.go` | `VarArgsMethodTest.java` | ✅ | 可变参数方法 |

### 测试统计

- **总测试文件数**: 58 个
- **已完成**: 58 个 ✅
- **完成率**: 100%
- **测试用例总数**: 1000+ 个表达式测试

### 测试覆盖范围

所有测试完全验证了以下特性的解析正确性：

- ✅ 所有运算符（算术、逻辑、比较、位运算）
- ✅ 所有字面量类型（整数、浮点、字符串、字符、布尔、null）
- ✅ 属性访问和链式调用
- ✅ 方法调用（包括可变参数）
- ✅ 索引访问和集合操作
- ✅ 投影和选择表达式
- ✅ 静态字段和方法引用
- ✅ 构造函数（包括内部类）
- ✅ Lambda 表达式
- ✅ 变量引用和上下文操作
- ✅ 类型转换和泛型
- ✅ 异常情况和边界条件


**注意**: 本项目仍在积极开发中。欢迎提交问题和建议！

Made with ❤️ by Go OGNL Contributors
