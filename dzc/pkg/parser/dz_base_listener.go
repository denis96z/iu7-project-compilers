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

// EnterDecls is called when production decls is entered.
func (s *BaseDZListener) EnterDecls(ctx *DeclsContext) {}

// ExitDecls is called when production decls is exited.
func (s *BaseDZListener) ExitDecls(ctx *DeclsContext) {}

// EnterDecl is called when production decl is entered.
func (s *BaseDZListener) EnterDecl(ctx *DeclContext) {}

// ExitDecl is called when production decl is exited.
func (s *BaseDZListener) ExitDecl(ctx *DeclContext) {}

// EnterSubdecl is called when production subdecl is entered.
func (s *BaseDZListener) EnterSubdecl(ctx *SubdeclContext) {}

// ExitSubdecl is called when production subdecl is exited.
func (s *BaseDZListener) ExitSubdecl(ctx *SubdeclContext) {}

// EnterProcdecl is called when production procdecl is entered.
func (s *BaseDZListener) EnterProcdecl(ctx *ProcdeclContext) {}

// ExitProcdecl is called when production procdecl is exited.
func (s *BaseDZListener) ExitProcdecl(ctx *ProcdeclContext) {}

// EnterFuncdecl is called when production funcdecl is entered.
func (s *BaseDZListener) EnterFuncdecl(ctx *FuncdeclContext) {}

// ExitFuncdecl is called when production funcdecl is exited.
func (s *BaseDZListener) ExitFuncdecl(ctx *FuncdeclContext) {}

// EnterProcheader is called when production procheader is entered.
func (s *BaseDZListener) EnterProcheader(ctx *ProcheaderContext) {}

// ExitProcheader is called when production procheader is exited.
func (s *BaseDZListener) ExitProcheader(ctx *ProcheaderContext) {}

// EnterFuncheader is called when production funcheader is entered.
func (s *BaseDZListener) EnterFuncheader(ctx *FuncheaderContext) {}

// ExitFuncheader is called when production funcheader is exited.
func (s *BaseDZListener) ExitFuncheader(ctx *FuncheaderContext) {}

// EnterArgs is called when production args is entered.
func (s *BaseDZListener) EnterArgs(ctx *ArgsContext) {}

// ExitArgs is called when production args is exited.
func (s *BaseDZListener) ExitArgs(ctx *ArgsContext) {}

// EnterArgdecls is called when production argdecls is entered.
func (s *BaseDZListener) EnterArgdecls(ctx *ArgdeclsContext) {}

// ExitArgdecls is called when production argdecls is exited.
func (s *BaseDZListener) ExitArgdecls(ctx *ArgdeclsContext) {}

// EnterArgdecl is called when production argdecl is entered.
func (s *BaseDZListener) EnterArgdecl(ctx *ArgdeclContext) {}

// ExitArgdecl is called when production argdecl is exited.
func (s *BaseDZListener) ExitArgdecl(ctx *ArgdeclContext) {}

// EnterFuncret is called when production funcret is entered.
func (s *BaseDZListener) EnterFuncret(ctx *FuncretContext) {}

// ExitFuncret is called when production funcret is exited.
func (s *BaseDZListener) ExitFuncret(ctx *FuncretContext) {}

// EnterComplexdecl is called when production complexdecl is entered.
func (s *BaseDZListener) EnterComplexdecl(ctx *ComplexdeclContext) {}

// ExitComplexdecl is called when production complexdecl is exited.
func (s *BaseDZListener) ExitComplexdecl(ctx *ComplexdeclContext) {}

// EnterEnumdecl is called when production enumdecl is entered.
func (s *BaseDZListener) EnterEnumdecl(ctx *EnumdeclContext) {}

// ExitEnumdecl is called when production enumdecl is exited.
func (s *BaseDZListener) ExitEnumdecl(ctx *EnumdeclContext) {}

// EnterEnumoptions is called when production enumoptions is entered.
func (s *BaseDZListener) EnterEnumoptions(ctx *EnumoptionsContext) {}

// ExitEnumoptions is called when production enumoptions is exited.
func (s *BaseDZListener) ExitEnumoptions(ctx *EnumoptionsContext) {}

// EnterEnumoption is called when production enumoption is entered.
func (s *BaseDZListener) EnterEnumoption(ctx *EnumoptionContext) {}

// ExitEnumoption is called when production enumoption is exited.
func (s *BaseDZListener) ExitEnumoption(ctx *EnumoptionContext) {}

// EnterStructdecl is called when production structdecl is entered.
func (s *BaseDZListener) EnterStructdecl(ctx *StructdeclContext) {}

// ExitStructdecl is called when production structdecl is exited.
func (s *BaseDZListener) ExitStructdecl(ctx *StructdeclContext) {}

// EnterStructattrs is called when production structattrs is entered.
func (s *BaseDZListener) EnterStructattrs(ctx *StructattrsContext) {}

// ExitStructattrs is called when production structattrs is exited.
func (s *BaseDZListener) ExitStructattrs(ctx *StructattrsContext) {}

// EnterStructattr is called when production structattr is entered.
func (s *BaseDZListener) EnterStructattr(ctx *StructattrContext) {}

// ExitStructattr is called when production structattr is exited.
func (s *BaseDZListener) ExitStructattr(ctx *StructattrContext) {}

// EnterConstdecl is called when production constdecl is entered.
func (s *BaseDZListener) EnterConstdecl(ctx *ConstdeclContext) {}

// ExitConstdecl is called when production constdecl is exited.
func (s *BaseDZListener) ExitConstdecl(ctx *ConstdeclContext) {}

// EnterConstexpr is called when production constexpr is entered.
func (s *BaseDZListener) EnterConstexpr(ctx *ConstexprContext) {}

// ExitConstexpr is called when production constexpr is exited.
func (s *BaseDZListener) ExitConstexpr(ctx *ConstexprContext) {}

// EnterIntexpr is called when production intexpr is entered.
func (s *BaseDZListener) EnterIntexpr(ctx *IntexprContext) {}

// ExitIntexpr is called when production intexpr is exited.
func (s *BaseDZListener) ExitIntexpr(ctx *IntexprContext) {}

// EnterBoolexpr is called when production boolexpr is entered.
func (s *BaseDZListener) EnterBoolexpr(ctx *BoolexprContext) {}

// ExitBoolexpr is called when production boolexpr is exited.
func (s *BaseDZListener) ExitBoolexpr(ctx *BoolexprContext) {}

// EnterConstval is called when production constval is entered.
func (s *BaseDZListener) EnterConstval(ctx *ConstvalContext) {}

// ExitConstval is called when production constval is exited.
func (s *BaseDZListener) ExitConstval(ctx *ConstvalContext) {}

// EnterTypedecl is called when production typedecl is entered.
func (s *BaseDZListener) EnterTypedecl(ctx *TypedeclContext) {}

// ExitTypedecl is called when production typedecl is exited.
func (s *BaseDZListener) ExitTypedecl(ctx *TypedeclContext) {}

// EnterTypespec is called when production typespec is entered.
func (s *BaseDZListener) EnterTypespec(ctx *TypespecContext) {}

// ExitTypespec is called when production typespec is exited.
func (s *BaseDZListener) ExitTypespec(ctx *TypespecContext) {}

// EnterSimpletypespec is called when production simpletypespec is entered.
func (s *BaseDZListener) EnterSimpletypespec(ctx *SimpletypespecContext) {}

// ExitSimpletypespec is called when production simpletypespec is exited.
func (s *BaseDZListener) ExitSimpletypespec(ctx *SimpletypespecContext) {}

// EnterBasictypespec is called when production basictypespec is entered.
func (s *BaseDZListener) EnterBasictypespec(ctx *BasictypespecContext) {}

// ExitBasictypespec is called when production basictypespec is exited.
func (s *BaseDZListener) ExitBasictypespec(ctx *BasictypespecContext) {}

// EnterNamedtypespec is called when production namedtypespec is entered.
func (s *BaseDZListener) EnterNamedtypespec(ctx *NamedtypespecContext) {}

// ExitNamedtypespec is called when production namedtypespec is exited.
func (s *BaseDZListener) ExitNamedtypespec(ctx *NamedtypespecContext) {}

// EnterReftypespec is called when production reftypespec is entered.
func (s *BaseDZListener) EnterReftypespec(ctx *ReftypespecContext) {}

// ExitReftypespec is called when production reftypespec is exited.
func (s *BaseDZListener) ExitReftypespec(ctx *ReftypespecContext) {}

// EnterArraytypespec is called when production arraytypespec is entered.
func (s *BaseDZListener) EnterArraytypespec(ctx *ArraytypespecContext) {}

// ExitArraytypespec is called when production arraytypespec is exited.
func (s *BaseDZListener) ExitArraytypespec(ctx *ArraytypespecContext) {}

// EnterSizespec is called when production sizespec is entered.
func (s *BaseDZListener) EnterSizespec(ctx *SizespecContext) {}

// ExitSizespec is called when production sizespec is exited.
func (s *BaseDZListener) ExitSizespec(ctx *SizespecContext) {}

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

// EnterDeclaration is called when production declaration is entered.
func (s *BaseDZListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseDZListener) ExitDeclaration(ctx *DeclarationContext) {}

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
