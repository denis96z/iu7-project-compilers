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

	// EnterStructdecl is called when entering the structdecl production.
	EnterStructdecl(c *StructdeclContext)

	// EnterStructattrs is called when entering the structattrs production.
	EnterStructattrs(c *StructattrsContext)

	// EnterStructattr is called when entering the structattr production.
	EnterStructattr(c *StructattrContext)

	// EnterEnumdecl is called when entering the enumdecl production.
	EnterEnumdecl(c *EnumdeclContext)

	// EnterEnumoptions is called when entering the enumoptions production.
	EnterEnumoptions(c *EnumoptionsContext)

	// EnterEnumoption is called when entering the enumoption production.
	EnterEnumoption(c *EnumoptionContext)

	// EnterUniondecl is called when entering the uniondecl production.
	EnterUniondecl(c *UniondeclContext)

	// EnterUnionattrs is called when entering the unionattrs production.
	EnterUnionattrs(c *UnionattrsContext)

	// EnterUnionattr is called when entering the unionattr production.
	EnterUnionattr(c *UnionattrContext)

	// EnterConstdecl is called when entering the constdecl production.
	EnterConstdecl(c *ConstdeclContext)

	// EnterConstasgn is called when entering the constasgn production.
	EnterConstasgn(c *ConstasgnContext)

	// EnterCasgn is called when entering the casgn production.
	EnterCasgn(c *CasgnContext)

	// EnterIntasgn is called when entering the intasgn production.
	EnterIntasgn(c *IntasgnContext)

	// EnterFloatasgn is called when entering the floatasgn production.
	EnterFloatasgn(c *FloatasgnContext)

	// EnterBoolconst is called when entering the boolconst production.
	EnterBoolconst(c *BoolconstContext)

	// EnterTypedecl is called when entering the typedecl production.
	EnterTypedecl(c *TypedeclContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterTypespec is called when entering the typespec production.
	EnterTypespec(c *TypespecContext)

	// EnterBasictypespec is called when entering the basictypespec production.
	EnterBasictypespec(c *BasictypespecContext)

	// EnterSimpletypespec is called when entering the simpletypespec production.
	EnterSimpletypespec(c *SimpletypespecContext)

	// EnterReftypespec is called when entering the reftypespec production.
	EnterReftypespec(c *ReftypespecContext)

	// EnterArraytypespec is called when entering the arraytypespec production.
	EnterArraytypespec(c *ArraytypespecContext)

	// EnterSizespec is called when entering the sizespec production.
	EnterSizespec(c *SizespecContext)

	// EnterStatements is called when entering the statements production.
	EnterStatements(c *StatementsContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterAsgnop is called when entering the asgnop production.
	EnterAsgnop(c *AsgnopContext)

	// EnterCondition is called when entering the condition production.
	EnterCondition(c *ConditionContext)

	// EnterIfblock is called when entering the ifblock production.
	EnterIfblock(c *IfblockContext)

	// EnterElseblocks is called when entering the elseblocks production.
	EnterElseblocks(c *ElseblocksContext)

	// EnterElifblock is called when entering the elifblock production.
	EnterElifblock(c *ElifblockContext)

	// EnterElseblock is called when entering the elseblock production.
	EnterElseblock(c *ElseblockContext)

	// EnterLoop is called when entering the loop production.
	EnterLoop(c *LoopContext)

	// EnterTrueloop is called when entering the trueloop production.
	EnterTrueloop(c *TrueloopContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterRetstatement is called when entering the retstatement production.
	EnterRetstatement(c *RetstatementContext)

	// EnterProcretstatement is called when entering the procretstatement production.
	EnterProcretstatement(c *ProcretstatementContext)

	// EnterFuncretstatement is called when entering the funcretstatement production.
	EnterFuncretstatement(c *FuncretstatementContext)

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

	// ExitStructdecl is called when exiting the structdecl production.
	ExitStructdecl(c *StructdeclContext)

	// ExitStructattrs is called when exiting the structattrs production.
	ExitStructattrs(c *StructattrsContext)

	// ExitStructattr is called when exiting the structattr production.
	ExitStructattr(c *StructattrContext)

	// ExitEnumdecl is called when exiting the enumdecl production.
	ExitEnumdecl(c *EnumdeclContext)

	// ExitEnumoptions is called when exiting the enumoptions production.
	ExitEnumoptions(c *EnumoptionsContext)

	// ExitEnumoption is called when exiting the enumoption production.
	ExitEnumoption(c *EnumoptionContext)

	// ExitUniondecl is called when exiting the uniondecl production.
	ExitUniondecl(c *UniondeclContext)

	// ExitUnionattrs is called when exiting the unionattrs production.
	ExitUnionattrs(c *UnionattrsContext)

	// ExitUnionattr is called when exiting the unionattr production.
	ExitUnionattr(c *UnionattrContext)

	// ExitConstdecl is called when exiting the constdecl production.
	ExitConstdecl(c *ConstdeclContext)

	// ExitConstasgn is called when exiting the constasgn production.
	ExitConstasgn(c *ConstasgnContext)

	// ExitCasgn is called when exiting the casgn production.
	ExitCasgn(c *CasgnContext)

	// ExitIntasgn is called when exiting the intasgn production.
	ExitIntasgn(c *IntasgnContext)

	// ExitFloatasgn is called when exiting the floatasgn production.
	ExitFloatasgn(c *FloatasgnContext)

	// ExitBoolconst is called when exiting the boolconst production.
	ExitBoolconst(c *BoolconstContext)

	// ExitTypedecl is called when exiting the typedecl production.
	ExitTypedecl(c *TypedeclContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitTypespec is called when exiting the typespec production.
	ExitTypespec(c *TypespecContext)

	// ExitBasictypespec is called when exiting the basictypespec production.
	ExitBasictypespec(c *BasictypespecContext)

	// ExitSimpletypespec is called when exiting the simpletypespec production.
	ExitSimpletypespec(c *SimpletypespecContext)

	// ExitReftypespec is called when exiting the reftypespec production.
	ExitReftypespec(c *ReftypespecContext)

	// ExitArraytypespec is called when exiting the arraytypespec production.
	ExitArraytypespec(c *ArraytypespecContext)

	// ExitSizespec is called when exiting the sizespec production.
	ExitSizespec(c *SizespecContext)

	// ExitStatements is called when exiting the statements production.
	ExitStatements(c *StatementsContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitAsgnop is called when exiting the asgnop production.
	ExitAsgnop(c *AsgnopContext)

	// ExitCondition is called when exiting the condition production.
	ExitCondition(c *ConditionContext)

	// ExitIfblock is called when exiting the ifblock production.
	ExitIfblock(c *IfblockContext)

	// ExitElseblocks is called when exiting the elseblocks production.
	ExitElseblocks(c *ElseblocksContext)

	// ExitElifblock is called when exiting the elifblock production.
	ExitElifblock(c *ElifblockContext)

	// ExitElseblock is called when exiting the elseblock production.
	ExitElseblock(c *ElseblockContext)

	// ExitLoop is called when exiting the loop production.
	ExitLoop(c *LoopContext)

	// ExitTrueloop is called when exiting the trueloop production.
	ExitTrueloop(c *TrueloopContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitRetstatement is called when exiting the retstatement production.
	ExitRetstatement(c *RetstatementContext)

	// ExitProcretstatement is called when exiting the procretstatement production.
	ExitProcretstatement(c *ProcretstatementContext)

	// ExitFuncretstatement is called when exiting the funcretstatement production.
	ExitFuncretstatement(c *FuncretstatementContext)
}
