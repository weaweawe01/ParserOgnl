import ognl.Ognl;
import ognl.OgnlException;
import ognl.SimpleNode;

public class ognlParser {
    public static void main(String[] args) throws OgnlException {
        // 待解析的 OGNL 表达式
        String expression = "(true ? 'tabHeader' : '') + (false ? 'tabHeader' : '')";

        // 解析表达式生成 AST 根节点（SimpleNode 是 OGNL AST 节点的基类）
        SimpleNode rootNode = (SimpleNode) Ognl.parseExpression(expression);
        System.out.println("rootNode:"+rootNode);
        // 打印 AST 结构（递归遍历节点）
        System.out.println("OGNL 表达式: " + expression);
        System.out.println("AST 结构:");
        printNode(rootNode, 0);
    }

    // 递归打印节点信息（缩进表示层级）
    private static void printNode(SimpleNode node, int depth) {
        // 缩进
        StringBuilder indent = new StringBuilder();
        for (int i = 0; i < depth; i++) {
            indent.append("  ");
        }

        // 打印节点类型和关键信息
        String nodeInfo = String.format("%s%s  %s 表达式片段: %s",
                indent,
                depth == 0 ? "" : "",
                node.getClass().getSimpleName(),
                node.toString() // 节点对应的表达式片段
        );
        System.out.println(nodeInfo);

        // 递归打印子节点
        for (int i = 0; i < node.jjtGetNumChildren(); i++) {
            printNode((SimpleNode) node.jjtGetChild(i), depth + 1);
        }
    }
}