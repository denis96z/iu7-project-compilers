// Code generated from /home/denis/Documents/iu7-project-compilers/dzc/DZ.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // DZ

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDZListener is a complete listener for a parse tree produced by DZParser.
type BaseDZListener struct{}

var _ DZListener = &BaseDZListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDZListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDZListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDZListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDZListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseDZListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseDZListener) ExitStart(ctx *StartContext) {}

// EnterPkg is called when production pkg is entered.
func (s *BaseDZListener) EnterPkg(ctx *PkgContext) {}

// ExitPkg is called when production pkg is exited.
func (s *BaseDZListener) ExitPkg(ctx *PkgContext) {}

// EnterDecl is called when production decl is entered.
func (s *BaseDZListener) EnterDecl(ctx *DeclContext) {}

// ExitDecl is called when production decl is exited.
func (s *BaseDZListener) ExitDecl(ctx *DeclContext) {}

// EnterConstDecl is called when production constDecl is entered.
func (s *BaseDZListener) EnterConstDecl(ctx *ConstDeclContext) {}

// ExitConstDecl is called when production constDecl is exited.
func (s *BaseDZListener) ExitConstDecl(ctx *ConstDeclContext) {}

// EnterTypeDecl is called when production typeDecl is entered.
func (s *BaseDZListener) EnterTypeDecl(ctx *TypeDeclContext) {}

// ExitTypeDecl is called when production typeDecl is exited.
func (s *BaseDZListener) ExitTypeDecl(ctx *TypeDeclContext) {}

// EnterTypeSpec is called when production typeSpec is entered.
func (s *BaseDZListener) EnterTypeSpec(ctx *TypeSpecContext) {}

// ExitTypeSpec is called when production typeSpec is exited.
func (s *BaseDZListener) ExitTypeSpec(ctx *TypeSpecContext) {}

// EnterSimpleTypeSpec is called when production simpleTypeSpec is entered.
func (s *BaseDZListener) EnterSimpleTypeSpec(ctx *SimpleTypeSpecContext) {}

// ExitSimpleTypeSpec is called when production simpleTypeSpec is exited.
func (s *BaseDZListener) ExitSimpleTypeSpec(ctx *SimpleTypeSpecContext) {}

// EnterBasicTypeSpec is called when production basicTypeSpec is entered.
func (s *BaseDZListener) EnterBasicTypeSpec(ctx *BasicTypeSpecContext) {}

// ExitBasicTypeSpec is called when production basicTypeSpec is exited.
func (s *BaseDZListener) ExitBasicTypeSpec(ctx *BasicTypeSpecContext) {}

// EnterNamedTypeSpec is called when production namedTypeSpec is entered.
func (s *BaseDZListener) EnterNamedTypeSpec(ctx *NamedTypeSpecContext) {}

// ExitNamedTypeSpec is called when production namedTypeSpec is exited.
func (s *BaseDZListener) ExitNamedTypeSpec(ctx *NamedTypeSpecContext) {}

// EnterRefTypeSpec is called when production refTypeSpec is entered.
func (s *BaseDZListener) EnterRefTypeSpec(ctx *RefTypeSpecContext) {}

// ExitRefTypeSpec is called when production refTypeSpec is exited.
func (s *BaseDZListener) ExitRefTypeSpec(ctx *RefTypeSpecContext) {}

// EnterArrayTypeSpec is called when production arrayTypeSpec is entered.
func (s *BaseDZListener) EnterArrayTypeSpec(ctx *ArrayTypeSpecContext) {}

// ExitArrayTypeSpec is called when production arrayTypeSpec is exited.
func (s *BaseDZListener) ExitArrayTypeSpec(ctx *ArrayTypeSpecContext) {}

// EnterSliceTypeSpec is called when production sliceTypeSpec is entered.
func (s *BaseDZListener) EnterSliceTypeSpec(ctx *SliceTypeSpecContext) {}

// ExitSliceTypeSpec is called when production sliceTypeSpec is exited.
func (s *BaseDZListener) ExitSliceTypeSpec(ctx *SliceTypeSpecContext) {}

// EnterEnumDecl is called when production enumDecl is entered.
func (s *BaseDZListener) EnterEnumDecl(ctx *EnumDeclContext) {}

// ExitEnumDecl is called when production enumDecl is exited.
func (s *BaseDZListener) ExitEnumDecl(ctx *EnumDeclContext) {}

// EnterEnumOption is called when production enumOption is entered.
func (s *BaseDZListener) EnterEnumOption(ctx *EnumOptionContext) {}

// ExitEnumOption is called when production enumOption is exited.
func (s *BaseDZListener) ExitEnumOption(ctx *EnumOptionContext) {}

// EnterStructDecl is called when production structDecl is entered.
func (s *BaseDZListener) EnterStructDecl(ctx *StructDeclContext) {}

// ExitStructDecl is called when production structDecl is exited.
func (s *BaseDZListener) ExitStructDecl(ctx *StructDeclContext) {}

// EnterStructAttr is called when production structAttr is entered.
func (s *BaseDZListener) EnterStructAttr(ctx *StructAttrContext) {}

// ExitStructAttr is called when production structAttr is exited.
func (s *BaseDZListener) ExitStructAttr(ctx *StructAttrContext) {}

// EnterProcDecl is called when production procDecl is entered.
func (s *BaseDZListener) EnterProcDecl(ctx *ProcDeclContext) {}

// ExitProcDecl is called when production procDecl is exited.
func (s *BaseDZListener) ExitProcDecl(ctx *ProcDeclContext) {}

// EnterProcArg is called when production procArg is entered.
func (s *BaseDZListener) EnterProcArg(ctx *ProcArgContext) {}

// ExitProcArg is called when production procArg is exited.
func (s *BaseDZListener) ExitProcArg(ctx *ProcArgContext) {}

// EnterFuncDecl is called when production funcDecl is entered.
func (s *BaseDZListener) EnterFuncDecl(ctx *FuncDeclContext) {}

// ExitFuncDecl is called when production funcDecl is exited.
func (s *BaseDZListener) ExitFuncDecl(ctx *FuncDeclContext) {}

// EnterFuncArg is called when production funcArg is entered.
func (s *BaseDZListener) EnterFuncArg(ctx *FuncArgContext) {}

// ExitFuncArg is called when production funcArg is exited.
func (s *BaseDZListener) ExitFuncArg(ctx *FuncArgContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseDZListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseDZListener) ExitBlock(ctx *BlockContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseDZListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseDZListener) ExitStatement(ctx *StatementContext) {}

// EnterCondition is called when production condition is entered.
func (s *BaseDZListener) EnterCondition(ctx *ConditionContext) {}

// ExitCondition is called when production condition is exited.
func (s *BaseDZListener) ExitCondition(ctx *ConditionContext) {}

// EnterIfConditionBranch is called when production ifConditionBranch is entered.
func (s *BaseDZListener) EnterIfConditionBranch(ctx *IfConditionBranchContext) {}

// ExitIfConditionBranch is called when production ifConditionBranch is exited.
func (s *BaseDZListener) ExitIfConditionBranch(ctx *IfConditionBranchContext) {}

// EnterElifConditionBranch is called when production elifConditionBranch is entered.
func (s *BaseDZListener) EnterElifConditionBranch(ctx *ElifConditionBranchContext) {}

// ExitElifConditionBranch is called when production elifConditionBranch is exited.
func (s *BaseDZListener) ExitElifConditionBranch(ctx *ElifConditionBranchContext) {}

// EnterElseConditionBranch is called when production elseConditionBranch is entered.
func (s *BaseDZListener) EnterElseConditionBranch(ctx *ElseConditionBranchContext) {}

// ExitElseConditionBranch is called when production elseConditionBranch is exited.
func (s *BaseDZListener) ExitElseConditionBranch(ctx *ElseConditionBranchContext) {}

// EnterLoop is called when production loop is entered.
func (s *BaseDZListener) EnterLoop(ctx *LoopContext) {}

// ExitLoop is called when production loop is exited.
func (s *BaseDZListener) ExitLoop(ctx *LoopContext) {}

// EnterTrueLoop is called when production trueLoop is entered.
func (s *BaseDZListener) EnterTrueLoop(ctx *TrueLoopContext) {}

// ExitTrueLoop is called when production trueLoop is exited.
func (s *BaseDZListener) ExitTrueLoop(ctx *TrueLoopContext) {}

// EnterWhileLoop is called when production whileLoop is entered.
func (s *BaseDZListener) EnterWhileLoop(ctx *WhileLoopContext) {}

// ExitWhileLoop is called when production whileLoop is exited.
func (s *BaseDZListener) ExitWhileLoop(ctx *WhileLoopContext) {}

// EnterBreakStatement is called when production breakStatement is entered.
func (s *BaseDZListener) EnterBreakStatement(ctx *BreakStatementContext) {}

// ExitBreakStatement is called when production breakStatement is exited.
func (s *BaseDZListener) ExitBreakStatement(ctx *BreakStatementContext) {}

// EnterContinueStatement is called when production continueStatement is entered.
func (s *BaseDZListener) EnterContinueStatement(ctx *ContinueStatementContext) {}

// ExitContinueStatement is called when production continueStatement is exited.
func (s *BaseDZListener) ExitContinueStatement(ctx *ContinueStatementContext) {}

// EnterVarDecl is called when production varDecl is entered.
func (s *BaseDZListener) EnterVarDecl(ctx *VarDeclContext) {}

// ExitVarDecl is called when production varDecl is exited.
func (s *BaseDZListener) ExitVarDecl(ctx *VarDeclContext) {}

// EnterProcCall is called when production procCall is entered.
func (s *BaseDZListener) EnterProcCall(ctx *ProcCallContext) {}

// ExitProcCall is called when production procCall is exited.
func (s *BaseDZListener) ExitProcCall(ctx *ProcCallContext) {}

// EnterProcParam is called when production procParam is entered.
func (s *BaseDZListener) EnterProcParam(ctx *ProcParamContext) {}

// ExitProcParam is called when production procParam is exited.
func (s *BaseDZListener) ExitProcParam(ctx *ProcParamContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDZListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDZListener) ExitExpression(ctx *ExpressionContext) {}

// EnterConstValue is called when production constValue is entered.
func (s *BaseDZListener) EnterConstValue(ctx *ConstValueContext) {}

// ExitConstValue is called when production constValue is exited.
func (s *BaseDZListener) ExitConstValue(ctx *ConstValueContext) {}

// EnterFuncCall is called when production funcCall is entered.
func (s *BaseDZListener) EnterFuncCall(ctx *FuncCallContext) {}

// ExitFuncCall is called when production funcCall is exited.
func (s *BaseDZListener) ExitFuncCall(ctx *FuncCallContext) {}

// EnterFuncParam is called when production funcParam is entered.
func (s *BaseDZListener) EnterFuncParam(ctx *FuncParamContext) {}

// ExitFuncParam is called when production funcParam is exited.
func (s *BaseDZListener) ExitFuncParam(ctx *FuncParamContext) {}

// EnterUnaryOperator is called when production unaryOperator is entered.
func (s *BaseDZListener) EnterUnaryOperator(ctx *UnaryOperatorContext) {}

// ExitUnaryOperator is called when production unaryOperator is exited.
func (s *BaseDZListener) ExitUnaryOperator(ctx *UnaryOperatorContext) {}

// EnterBinaryOperator is called when production binaryOperator is entered.
func (s *BaseDZListener) EnterBinaryOperator(ctx *BinaryOperatorContext) {}

// ExitBinaryOperator is called when production binaryOperator is exited.
func (s *BaseDZListener) ExitBinaryOperator(ctx *BinaryOperatorContext) {}

// EnterReturnStatement is called when production returnStatement is entered.
func (s *BaseDZListener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production returnStatement is exited.
func (s *BaseDZListener) ExitReturnStatement(ctx *ReturnStatementContext) {}
