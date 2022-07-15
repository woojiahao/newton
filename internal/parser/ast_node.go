package parser

type ASTNodeType int

const (
	Undefined ASTNodeType = iota
	OperatorPlus
	OperatorMinus
	OperatorMul
	OperatorDiv
	UnaryMinus
	NumberValue
)

type ASTNode struct {
	ASTNodeType ASTNodeType
	Value       Value
	Left        *ASTNode
	Right       *ASTNode
}

func emptyASTNode() *ASTNode {
	return &ASTNode{Undefined, 0.0, nil, nil}
}

func createASTNode(astNodeType ASTNodeType, left, right *ASTNode) *ASTNode {
	return &ASTNode{astNodeType, 0.0, left, right}
}

func createUnaryASTNode(left *ASTNode) *ASTNode {
	return &ASTNode{UnaryMinus, 0.0, left, nil}
}

func createNumberASTNode(value Value) *ASTNode {
	return &ASTNode{NumberValue, value, nil, nil}
}
