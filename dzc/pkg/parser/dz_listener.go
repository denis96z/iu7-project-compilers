// Code generated from /home/denis/Documents/iu7-project-compilers/dzc/DZ.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // DZ

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DZListener is a complete listener for a parse tree produced by DZParser.
type DZListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterPkg is called when entering the pkg production.
	EnterPkg(c *PkgContext)

	// EnterDecl is called when entering the decl production.
	EnterDecl(c *DeclContext)

	// EnterConstDecl is called when entering the constDecl production.
	EnterConstDecl(c *ConstDeclContext)

	// EnterTypeDecl is called when entering the typeDecl production.
	EnterTypeDecl(c *TypeDeclContext)

	// EnterTypeSpec is called when entering the typeSpec production.
	EnterTypeSpec(c *TypeSpecContext)

	// EnterSimpleTypeSpec is called when entering the simpleTypeSpec production.
	EnterSimpleTypeSpec(c *SimpleTypeSpecContext)

	// EnterBasicTypeSpec is called when entering the basicTypeSpec production.
	EnterBasicTypeSpec(c *BasicTypeSpecContext)

	// EnterNamedTypeSpec is called when entering the namedTypeSpec production.
	EnterNamedTypeSpec(c *NamedTypeSpecContext)

	// EnterRefTypeSpec is called when entering the refTypeSpec production.
	EnterRefTypeSpec(c *RefTypeSpecContext)

	// EnterArrayTypeSpec is called when entering the arrayTypeSpec production.
	EnterArrayTypeSpec(c *ArrayTypeSpecContext)

	// EnterSliceTypeSpec is called when entering the sliceTypeSpec production.
	EnterSliceTypeSpec(c *SliceTypeSpecContext)

	// EnterEnumDecl is called when entering the enumDecl production.
	EnterEnumDecl(c *EnumDeclContext)

	// EnterEnumOption is called when entering the enumOption production.
	EnterEnumOption(c *EnumOptionContext)

	// EnterStructDecl is called when entering the structDecl production.
	EnterStructDecl(c *StructDeclContext)

	// EnterStructAttr is called when entering the structAttr production.
	EnterStructAttr(c *StructAttrContext)

	// EnterProcDecl is called when entering the procDecl production.
	EnterProcDecl(c *ProcDeclContext)

	// EnterProcArg is called when entering the procArg production.
	EnterProcArg(c *ProcArgContext)

	// EnterFuncDecl is called when entering the funcDecl production.
	EnterFuncDecl(c *FuncDeclContext)

	// EnterFuncArg is called when entering the funcArg production.
	EnterFuncArg(c *FuncArgContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterCondition is called when entering the condition production.
	EnterCondition(c *ConditionContext)

	// EnterIfConditionBranch is called when entering the ifConditionBranch production.
	EnterIfConditionBranch(c *IfConditionBranchContext)

	// EnterElifConditionBranch is called when entering the elifConditionBranch production.
	EnterElifConditionBranch(c *ElifConditionBranchContext)

	// EnterElseConditionBranch is called when entering the elseConditionBranch production.
	EnterElseConditionBranch(c *ElseConditionBranchContext)

	// EnterLoop is called when entering the loop production.
	EnterLoop(c *LoopContext)

	// EnterTrueLoop is called when entering the trueLoop production.
	EnterTrueLoop(c *TrueLoopContext)

	// EnterWhileLoop is called when entering the whileLoop production.
	EnterWhileLoop(c *WhileLoopContext)

	// EnterBreakStatement is called when entering the breakStatement production.
	EnterBreakStatement(c *BreakStatementContext)

	// EnterContinueStatement is called when entering the continueStatement production.
	EnterContinueStatement(c *ContinueStatementContext)

	// EnterVarDecl is called when entering the varDecl production.
	EnterVarDecl(c *VarDeclContext)

	// EnterProcCall is called when entering the procCall production.
	EnterProcCall(c *ProcCallContext)

	// EnterProcParam is called when entering the procParam production.
	EnterProcParam(c *ProcParamContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterConstValue is called when entering the constValue production.
	EnterConstValue(c *ConstValueContext)

	// EnterFuncParam is called when entering the funcParam production.
	EnterFuncParam(c *FuncParamContext)

	// EnterUnaryOperator is called when entering the unaryOperator production.
	EnterUnaryOperator(c *UnaryOperatorContext)

	// EnterBinaryOperator is called when entering the binaryOperator production.
	EnterBinaryOperator(c *BinaryOperatorContext)

	// EnterReturnStatement is called when entering the returnStatement production.
	EnterReturnStatement(c *ReturnStatementContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitPkg is called when exiting the pkg production.
	ExitPkg(c *PkgContext)

	// ExitDecl is called when exiting the decl production.
	ExitDecl(c *DeclContext)

	// ExitConstDecl is called when exiting the constDecl production.
	ExitConstDecl(c *ConstDeclContext)

	// ExitTypeDecl is called when exiting the typeDecl production.
	ExitTypeDecl(c *TypeDeclContext)

	// ExitTypeSpec is called when exiting the typeSpec production.
	ExitTypeSpec(c *TypeSpecContext)

	// ExitSimpleTypeSpec is called when exiting the simpleTypeSpec production.
	ExitSimpleTypeSpec(c *SimpleTypeSpecContext)

	// ExitBasicTypeSpec is called when exiting the basicTypeSpec production.
	ExitBasicTypeSpec(c *BasicTypeSpecContext)

	// ExitNamedTypeSpec is called when exiting the namedTypeSpec production.
	ExitNamedTypeSpec(c *NamedTypeSpecContext)

	// ExitRefTypeSpec is called when exiting the refTypeSpec production.
	ExitRefTypeSpec(c *RefTypeSpecContext)

	// ExitArrayTypeSpec is called when exiting the arrayTypeSpec production.
	ExitArrayTypeSpec(c *ArrayTypeSpecContext)

	// ExitSliceTypeSpec is called when exiting the sliceTypeSpec production.
	ExitSliceTypeSpec(c *SliceTypeSpecContext)

	// ExitEnumDecl is called when exiting the enumDecl production.
	ExitEnumDecl(c *EnumDeclContext)

	// ExitEnumOption is called when exiting the enumOption production.
	ExitEnumOption(c *EnumOptionContext)

	// ExitStructDecl is called when exiting the structDecl production.
	ExitStructDecl(c *StructDeclContext)

	// ExitStructAttr is called when exiting the structAttr production.
	ExitStructAttr(c *StructAttrContext)

	// ExitProcDecl is called when exiting the procDecl production.
	ExitProcDecl(c *ProcDeclContext)

	// ExitProcArg is called when exiting the procArg production.
	ExitProcArg(c *ProcArgContext)

	// ExitFuncDecl is called when exiting the funcDecl production.
	ExitFuncDecl(c *FuncDeclContext)

	// ExitFuncArg is called when exiting the funcArg production.
	ExitFuncArg(c *FuncArgContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitCondition is called when exiting the condition production.
	ExitCondition(c *ConditionContext)

	// ExitIfConditionBranch is called when exiting the ifConditionBranch production.
	ExitIfConditionBranch(c *IfConditionBranchContext)

	// ExitElifConditionBranch is called when exiting the elifConditionBranch production.
	ExitElifConditionBranch(c *ElifConditionBranchContext)

	// ExitElseConditionBranch is called when exiting the elseConditionBranch production.
	ExitElseConditionBranch(c *ElseConditionBranchContext)

	// ExitLoop is called when exiting the loop production.
	ExitLoop(c *LoopContext)

	// ExitTrueLoop is called when exiting the trueLoop production.
	ExitTrueLoop(c *TrueLoopContext)

	// ExitWhileLoop is called when exiting the whileLoop production.
	ExitWhileLoop(c *WhileLoopContext)

	// ExitBreakStatement is called when exiting the breakStatement production.
	ExitBreakStatement(c *BreakStatementContext)

	// ExitContinueStatement is called when exiting the continueStatement production.
	ExitContinueStatement(c *ContinueStatementContext)

	// ExitVarDecl is called when exiting the varDecl production.
	ExitVarDecl(c *VarDeclContext)

	// ExitProcCall is called when exiting the procCall production.
	ExitProcCall(c *ProcCallContext)

	// ExitProcParam is called when exiting the procParam production.
	ExitProcParam(c *ProcParamContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitConstValue is called when exiting the constValue production.
	ExitConstValue(c *ConstValueContext)

	// ExitFuncParam is called when exiting the funcParam production.
	ExitFuncParam(c *FuncParamContext)

	// ExitUnaryOperator is called when exiting the unaryOperator production.
	ExitUnaryOperator(c *UnaryOperatorContext)

	// ExitBinaryOperator is called when exiting the binaryOperator production.
	ExitBinaryOperator(c *BinaryOperatorContext)

	// ExitReturnStatement is called when exiting the returnStatement production.
	ExitReturnStatement(c *ReturnStatementContext)
}
