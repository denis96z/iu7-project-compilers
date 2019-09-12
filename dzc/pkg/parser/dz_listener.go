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

	// EnterDecls is called when entering the decls production.
	EnterDecls(c *DeclsContext)

	// EnterDecl is called when entering the decl production.
	EnterDecl(c *DeclContext)

	// EnterSubdecl is called when entering the subdecl production.
	EnterSubdecl(c *SubdeclContext)

	// EnterProcdecl is called when entering the procdecl production.
	EnterProcdecl(c *ProcdeclContext)

	// EnterFuncdecl is called when entering the funcdecl production.
	EnterFuncdecl(c *FuncdeclContext)

	// EnterProcheader is called when entering the procheader production.
	EnterProcheader(c *ProcheaderContext)

	// EnterFuncheader is called when entering the funcheader production.
	EnterFuncheader(c *FuncheaderContext)

	// EnterArgs is called when entering the args production.
	EnterArgs(c *ArgsContext)

	// EnterArgdecls is called when entering the argdecls production.
	EnterArgdecls(c *ArgdeclsContext)

	// EnterArgdecl is called when entering the argdecl production.
	EnterArgdecl(c *ArgdeclContext)

	// EnterFuncret is called when entering the funcret production.
	EnterFuncret(c *FuncretContext)

	// EnterComplexdecl is called when entering the complexdecl production.
	EnterComplexdecl(c *ComplexdeclContext)

	// EnterEnumdecl is called when entering the enumdecl production.
	EnterEnumdecl(c *EnumdeclContext)

	// EnterEnumoptions is called when entering the enumoptions production.
	EnterEnumoptions(c *EnumoptionsContext)

	// EnterEnumoption is called when entering the enumoption production.
	EnterEnumoption(c *EnumoptionContext)

	// EnterStructdecl is called when entering the structdecl production.
	EnterStructdecl(c *StructdeclContext)

	// EnterStructattrs is called when entering the structattrs production.
	EnterStructattrs(c *StructattrsContext)

	// EnterStructattr is called when entering the structattr production.
	EnterStructattr(c *StructattrContext)

	// EnterConstdecl is called when entering the constdecl production.
	EnterConstdecl(c *ConstdeclContext)

	// EnterConstexpr is called when entering the constexpr production.
	EnterConstexpr(c *ConstexprContext)

	// EnterIntexpr is called when entering the intexpr production.
	EnterIntexpr(c *IntexprContext)

	// EnterBoolexpr is called when entering the boolexpr production.
	EnterBoolexpr(c *BoolexprContext)

	// EnterConstval is called when entering the constval production.
	EnterConstval(c *ConstvalContext)

	// EnterTypedecl is called when entering the typedecl production.
	EnterTypedecl(c *TypedeclContext)

	// EnterTypespec is called when entering the typespec production.
	EnterTypespec(c *TypespecContext)

	// EnterSimpletypespec is called when entering the simpletypespec production.
	EnterSimpletypespec(c *SimpletypespecContext)

	// EnterBasictypespec is called when entering the basictypespec production.
	EnterBasictypespec(c *BasictypespecContext)

	// EnterNamedtypespec is called when entering the namedtypespec production.
	EnterNamedtypespec(c *NamedtypespecContext)

	// EnterReftypespec is called when entering the reftypespec production.
	EnterReftypespec(c *ReftypespecContext)

	// EnterArraytypespec is called when entering the arraytypespec production.
	EnterArraytypespec(c *ArraytypespecContext)

	// EnterSizespec is called when entering the sizespec production.
	EnterSizespec(c *SizespecContext)

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

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterProcCall is called when entering the procCall production.
	EnterProcCall(c *ProcCallContext)

	// EnterProcParam is called when entering the procParam production.
	EnterProcParam(c *ProcParamContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

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

	// ExitDecls is called when exiting the decls production.
	ExitDecls(c *DeclsContext)

	// ExitDecl is called when exiting the decl production.
	ExitDecl(c *DeclContext)

	// ExitSubdecl is called when exiting the subdecl production.
	ExitSubdecl(c *SubdeclContext)

	// ExitProcdecl is called when exiting the procdecl production.
	ExitProcdecl(c *ProcdeclContext)

	// ExitFuncdecl is called when exiting the funcdecl production.
	ExitFuncdecl(c *FuncdeclContext)

	// ExitProcheader is called when exiting the procheader production.
	ExitProcheader(c *ProcheaderContext)

	// ExitFuncheader is called when exiting the funcheader production.
	ExitFuncheader(c *FuncheaderContext)

	// ExitArgs is called when exiting the args production.
	ExitArgs(c *ArgsContext)

	// ExitArgdecls is called when exiting the argdecls production.
	ExitArgdecls(c *ArgdeclsContext)

	// ExitArgdecl is called when exiting the argdecl production.
	ExitArgdecl(c *ArgdeclContext)

	// ExitFuncret is called when exiting the funcret production.
	ExitFuncret(c *FuncretContext)

	// ExitComplexdecl is called when exiting the complexdecl production.
	ExitComplexdecl(c *ComplexdeclContext)

	// ExitEnumdecl is called when exiting the enumdecl production.
	ExitEnumdecl(c *EnumdeclContext)

	// ExitEnumoptions is called when exiting the enumoptions production.
	ExitEnumoptions(c *EnumoptionsContext)

	// ExitEnumoption is called when exiting the enumoption production.
	ExitEnumoption(c *EnumoptionContext)

	// ExitStructdecl is called when exiting the structdecl production.
	ExitStructdecl(c *StructdeclContext)

	// ExitStructattrs is called when exiting the structattrs production.
	ExitStructattrs(c *StructattrsContext)

	// ExitStructattr is called when exiting the structattr production.
	ExitStructattr(c *StructattrContext)

	// ExitConstdecl is called when exiting the constdecl production.
	ExitConstdecl(c *ConstdeclContext)

	// ExitConstexpr is called when exiting the constexpr production.
	ExitConstexpr(c *ConstexprContext)

	// ExitIntexpr is called when exiting the intexpr production.
	ExitIntexpr(c *IntexprContext)

	// ExitBoolexpr is called when exiting the boolexpr production.
	ExitBoolexpr(c *BoolexprContext)

	// ExitConstval is called when exiting the constval production.
	ExitConstval(c *ConstvalContext)

	// ExitTypedecl is called when exiting the typedecl production.
	ExitTypedecl(c *TypedeclContext)

	// ExitTypespec is called when exiting the typespec production.
	ExitTypespec(c *TypespecContext)

	// ExitSimpletypespec is called when exiting the simpletypespec production.
	ExitSimpletypespec(c *SimpletypespecContext)

	// ExitBasictypespec is called when exiting the basictypespec production.
	ExitBasictypespec(c *BasictypespecContext)

	// ExitNamedtypespec is called when exiting the namedtypespec production.
	ExitNamedtypespec(c *NamedtypespecContext)

	// ExitReftypespec is called when exiting the reftypespec production.
	ExitReftypespec(c *ReftypespecContext)

	// ExitArraytypespec is called when exiting the arraytypespec production.
	ExitArraytypespec(c *ArraytypespecContext)

	// ExitSizespec is called when exiting the sizespec production.
	ExitSizespec(c *SizespecContext)

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

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitProcCall is called when exiting the procCall production.
	ExitProcCall(c *ProcCallContext)

	// ExitProcParam is called when exiting the procParam production.
	ExitProcParam(c *ProcParamContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitFuncParam is called when exiting the funcParam production.
	ExitFuncParam(c *FuncParamContext)

	// ExitUnaryOperator is called when exiting the unaryOperator production.
	ExitUnaryOperator(c *UnaryOperatorContext)

	// ExitBinaryOperator is called when exiting the binaryOperator production.
	ExitBinaryOperator(c *BinaryOperatorContext)

	// ExitReturnStatement is called when exiting the returnStatement production.
	ExitReturnStatement(c *ReturnStatementContext)
}
