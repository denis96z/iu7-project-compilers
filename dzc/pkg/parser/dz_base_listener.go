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

// EnterProcargs is called when production procargs is entered.
func (s *BaseDZListener) EnterProcargs(ctx *ProcargsContext) {}

// ExitProcargs is called when production procargs is exited.
func (s *BaseDZListener) ExitProcargs(ctx *ProcargsContext) {}

// EnterFuncargs is called when production funcargs is entered.
func (s *BaseDZListener) EnterFuncargs(ctx *FuncargsContext) {}

// ExitFuncargs is called when production funcargs is exited.
func (s *BaseDZListener) ExitFuncargs(ctx *FuncargsContext) {}

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

// EnterUniondecl is called when production uniondecl is entered.
func (s *BaseDZListener) EnterUniondecl(ctx *UniondeclContext) {}

// ExitUniondecl is called when production uniondecl is exited.
func (s *BaseDZListener) ExitUniondecl(ctx *UniondeclContext) {}

// EnterUnionattrs is called when production unionattrs is entered.
func (s *BaseDZListener) EnterUnionattrs(ctx *UnionattrsContext) {}

// ExitUnionattrs is called when production unionattrs is exited.
func (s *BaseDZListener) ExitUnionattrs(ctx *UnionattrsContext) {}

// EnterUnionattr is called when production unionattr is entered.
func (s *BaseDZListener) EnterUnionattr(ctx *UnionattrContext) {}

// ExitUnionattr is called when production unionattr is exited.
func (s *BaseDZListener) ExitUnionattr(ctx *UnionattrContext) {}

// EnterConstdecl is called when production constdecl is entered.
func (s *BaseDZListener) EnterConstdecl(ctx *ConstdeclContext) {}

// ExitConstdecl is called when production constdecl is exited.
func (s *BaseDZListener) ExitConstdecl(ctx *ConstdeclContext) {}

// EnterConstasgn is called when production constasgn is entered.
func (s *BaseDZListener) EnterConstasgn(ctx *ConstasgnContext) {}

// ExitConstasgn is called when production constasgn is exited.
func (s *BaseDZListener) ExitConstasgn(ctx *ConstasgnContext) {}

// EnterCasgn is called when production casgn is entered.
func (s *BaseDZListener) EnterCasgn(ctx *CasgnContext) {}

// ExitCasgn is called when production casgn is exited.
func (s *BaseDZListener) ExitCasgn(ctx *CasgnContext) {}

// EnterIntasgn is called when production intasgn is entered.
func (s *BaseDZListener) EnterIntasgn(ctx *IntasgnContext) {}

// ExitIntasgn is called when production intasgn is exited.
func (s *BaseDZListener) ExitIntasgn(ctx *IntasgnContext) {}

// EnterFloatasgn is called when production floatasgn is entered.
func (s *BaseDZListener) EnterFloatasgn(ctx *FloatasgnContext) {}

// ExitFloatasgn is called when production floatasgn is exited.
func (s *BaseDZListener) ExitFloatasgn(ctx *FloatasgnContext) {}

// EnterBoolconst is called when production boolconst is entered.
func (s *BaseDZListener) EnterBoolconst(ctx *BoolconstContext) {}

// ExitBoolconst is called when production boolconst is exited.
func (s *BaseDZListener) ExitBoolconst(ctx *BoolconstContext) {}

// EnterTypedecl is called when production typedecl is entered.
func (s *BaseDZListener) EnterTypedecl(ctx *TypedeclContext) {}

// ExitTypedecl is called when production typedecl is exited.
func (s *BaseDZListener) ExitTypedecl(ctx *TypedeclContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseDZListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseDZListener) ExitBlock(ctx *BlockContext) {}

// EnterTypespec is called when production typespec is entered.
func (s *BaseDZListener) EnterTypespec(ctx *TypespecContext) {}

// ExitTypespec is called when production typespec is exited.
func (s *BaseDZListener) ExitTypespec(ctx *TypespecContext) {}

// EnterBasictypespec is called when production basictypespec is entered.
func (s *BaseDZListener) EnterBasictypespec(ctx *BasictypespecContext) {}

// ExitBasictypespec is called when production basictypespec is exited.
func (s *BaseDZListener) ExitBasictypespec(ctx *BasictypespecContext) {}

// EnterSimpletypespec is called when production simpletypespec is entered.
func (s *BaseDZListener) EnterSimpletypespec(ctx *SimpletypespecContext) {}

// ExitSimpletypespec is called when production simpletypespec is exited.
func (s *BaseDZListener) ExitSimpletypespec(ctx *SimpletypespecContext) {}

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

// EnterStatements is called when production statements is entered.
func (s *BaseDZListener) EnterStatements(ctx *StatementsContext) {}

// ExitStatements is called when production statements is exited.
func (s *BaseDZListener) ExitStatements(ctx *StatementsContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseDZListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseDZListener) ExitStatement(ctx *StatementContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseDZListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseDZListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterAsgnop is called when production asgnop is entered.
func (s *BaseDZListener) EnterAsgnop(ctx *AsgnopContext) {}

// ExitAsgnop is called when production asgnop is exited.
func (s *BaseDZListener) ExitAsgnop(ctx *AsgnopContext) {}

// EnterCondition is called when production condition is entered.
func (s *BaseDZListener) EnterCondition(ctx *ConditionContext) {}

// ExitCondition is called when production condition is exited.
func (s *BaseDZListener) ExitCondition(ctx *ConditionContext) {}

// EnterIfblock is called when production ifblock is entered.
func (s *BaseDZListener) EnterIfblock(ctx *IfblockContext) {}

// ExitIfblock is called when production ifblock is exited.
func (s *BaseDZListener) ExitIfblock(ctx *IfblockContext) {}

// EnterElseblocks is called when production elseblocks is entered.
func (s *BaseDZListener) EnterElseblocks(ctx *ElseblocksContext) {}

// ExitElseblocks is called when production elseblocks is exited.
func (s *BaseDZListener) ExitElseblocks(ctx *ElseblocksContext) {}

// EnterElifblock is called when production elifblock is entered.
func (s *BaseDZListener) EnterElifblock(ctx *ElifblockContext) {}

// ExitElifblock is called when production elifblock is exited.
func (s *BaseDZListener) ExitElifblock(ctx *ElifblockContext) {}

// EnterElseblock is called when production elseblock is entered.
func (s *BaseDZListener) EnterElseblock(ctx *ElseblockContext) {}

// ExitElseblock is called when production elseblock is exited.
func (s *BaseDZListener) ExitElseblock(ctx *ElseblockContext) {}

// EnterLoop is called when production loop is entered.
func (s *BaseDZListener) EnterLoop(ctx *LoopContext) {}

// ExitLoop is called when production loop is exited.
func (s *BaseDZListener) ExitLoop(ctx *LoopContext) {}

// EnterTrueloop is called when production trueloop is entered.
func (s *BaseDZListener) EnterTrueloop(ctx *TrueloopContext) {}

// ExitTrueloop is called when production trueloop is exited.
func (s *BaseDZListener) ExitTrueloop(ctx *TrueloopContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDZListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDZListener) ExitExpression(ctx *ExpressionContext) {}

// EnterRetstatement is called when production retstatement is entered.
func (s *BaseDZListener) EnterRetstatement(ctx *RetstatementContext) {}

// ExitRetstatement is called when production retstatement is exited.
func (s *BaseDZListener) ExitRetstatement(ctx *RetstatementContext) {}

// EnterProcretstatement is called when production procretstatement is entered.
func (s *BaseDZListener) EnterProcretstatement(ctx *ProcretstatementContext) {}

// ExitProcretstatement is called when production procretstatement is exited.
func (s *BaseDZListener) ExitProcretstatement(ctx *ProcretstatementContext) {}

// EnterFuncretstatement is called when production funcretstatement is entered.
func (s *BaseDZListener) EnterFuncretstatement(ctx *FuncretstatementContext) {}

// ExitFuncretstatement is called when production funcretstatement is exited.
func (s *BaseDZListener) ExitFuncretstatement(ctx *FuncretstatementContext) {}
