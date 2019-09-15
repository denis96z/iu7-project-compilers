// Code generated from /home/denis/Documents/iu7-project-compilers/stc/SmallTalk.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // SmallTalk

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSmallTalkListener is a complete listener for a parse tree produced by SmallTalkParser.
type BaseSmallTalkListener struct{}

var _ SmallTalkListener = &BaseSmallTalkListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSmallTalkListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSmallTalkListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSmallTalkListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSmallTalkListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterScript is called when production script is entered.
func (s *BaseSmallTalkListener) EnterScript(ctx *ScriptContext) {}

// ExitScript is called when production script is exited.
func (s *BaseSmallTalkListener) ExitScript(ctx *ScriptContext) {}

// EnterClassDef is called when production classDef is entered.
func (s *BaseSmallTalkListener) EnterClassDef(ctx *ClassDefContext) {}

// ExitClassDef is called when production classDef is exited.
func (s *BaseSmallTalkListener) ExitClassDef(ctx *ClassDefContext) {}

// EnterMain is called when production main is entered.
func (s *BaseSmallTalkListener) EnterMain(ctx *MainContext) {}

// ExitMain is called when production main is exited.
func (s *BaseSmallTalkListener) ExitMain(ctx *MainContext) {}

// EnterInstanceVars is called when production instanceVars is entered.
func (s *BaseSmallTalkListener) EnterInstanceVars(ctx *InstanceVarsContext) {}

// ExitInstanceVars is called when production instanceVars is exited.
func (s *BaseSmallTalkListener) ExitInstanceVars(ctx *InstanceVarsContext) {}

// EnterInstanceVar is called when production instanceVar is entered.
func (s *BaseSmallTalkListener) EnterInstanceVar(ctx *InstanceVarContext) {}

// ExitInstanceVar is called when production instanceVar is exited.
func (s *BaseSmallTalkListener) ExitInstanceVar(ctx *InstanceVarContext) {}

// EnterMethods is called when production methods is entered.
func (s *BaseSmallTalkListener) EnterMethods(ctx *MethodsContext) {}

// ExitMethods is called when production methods is exited.
func (s *BaseSmallTalkListener) ExitMethods(ctx *MethodsContext) {}

// EnterMethod is called when production method is entered.
func (s *BaseSmallTalkListener) EnterMethod(ctx *MethodContext) {}

// ExitMethod is called when production method is exited.
func (s *BaseSmallTalkListener) ExitMethod(ctx *MethodContext) {}

// EnterNamedMethod is called when production namedMethod is entered.
func (s *BaseSmallTalkListener) EnterNamedMethod(ctx *NamedMethodContext) {}

// ExitNamedMethod is called when production namedMethod is exited.
func (s *BaseSmallTalkListener) ExitNamedMethod(ctx *NamedMethodContext) {}

// EnterKeywordMethod is called when production keywordMethod is entered.
func (s *BaseSmallTalkListener) EnterKeywordMethod(ctx *KeywordMethodContext) {}

// ExitKeywordMethod is called when production keywordMethod is exited.
func (s *BaseSmallTalkListener) ExitKeywordMethod(ctx *KeywordMethodContext) {}

// EnterLocalVars is called when production localVars is entered.
func (s *BaseSmallTalkListener) EnterLocalVars(ctx *LocalVarsContext) {}

// ExitLocalVars is called when production localVars is exited.
func (s *BaseSmallTalkListener) ExitLocalVars(ctx *LocalVarsContext) {}

// EnterLocalVar is called when production localVar is entered.
func (s *BaseSmallTalkListener) EnterLocalVar(ctx *LocalVarContext) {}

// ExitLocalVar is called when production localVar is exited.
func (s *BaseSmallTalkListener) ExitLocalVar(ctx *LocalVarContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseSmallTalkListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseSmallTalkListener) ExitBlock(ctx *BlockContext) {}

// EnterBlockArgs is called when production blockArgs is entered.
func (s *BaseSmallTalkListener) EnterBlockArgs(ctx *BlockArgsContext) {}

// ExitBlockArgs is called when production blockArgs is exited.
func (s *BaseSmallTalkListener) ExitBlockArgs(ctx *BlockArgsContext) {}

// EnterBlockArg is called when production blockArg is entered.
func (s *BaseSmallTalkListener) EnterBlockArg(ctx *BlockArgContext) {}

// ExitBlockArg is called when production blockArg is exited.
func (s *BaseSmallTalkListener) ExitBlockArg(ctx *BlockArgContext) {}

// EnterBody is called when production body is entered.
func (s *BaseSmallTalkListener) EnterBody(ctx *BodyContext) {}

// ExitBody is called when production body is exited.
func (s *BaseSmallTalkListener) ExitBody(ctx *BodyContext) {}

// EnterStatements is called when production statements is entered.
func (s *BaseSmallTalkListener) EnterStatements(ctx *StatementsContext) {}

// ExitStatements is called when production statements is exited.
func (s *BaseSmallTalkListener) ExitStatements(ctx *StatementsContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseSmallTalkListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseSmallTalkListener) ExitStatement(ctx *StatementContext) {}

// EnterAssignmentStatement is called when production assignmentStatement is entered.
func (s *BaseSmallTalkListener) EnterAssignmentStatement(ctx *AssignmentStatementContext) {}

// ExitAssignmentStatement is called when production assignmentStatement is exited.
func (s *BaseSmallTalkListener) ExitAssignmentStatement(ctx *AssignmentStatementContext) {}

// EnterAssignmentItem is called when production assignmentItem is entered.
func (s *BaseSmallTalkListener) EnterAssignmentItem(ctx *AssignmentItemContext) {}

// ExitAssignmentItem is called when production assignmentItem is exited.
func (s *BaseSmallTalkListener) ExitAssignmentItem(ctx *AssignmentItemContext) {}

// EnterMessageStatements is called when production messageStatements is entered.
func (s *BaseSmallTalkListener) EnterMessageStatements(ctx *MessageStatementsContext) {}

// ExitMessageStatements is called when production messageStatements is exited.
func (s *BaseSmallTalkListener) ExitMessageStatements(ctx *MessageStatementsContext) {}

// EnterMessageStatement is called when production messageStatement is entered.
func (s *BaseSmallTalkListener) EnterMessageStatement(ctx *MessageStatementContext) {}

// ExitMessageStatement is called when production messageStatement is exited.
func (s *BaseSmallTalkListener) ExitMessageStatement(ctx *MessageStatementContext) {}

// EnterMessageExpression is called when production messageExpression is entered.
func (s *BaseSmallTalkListener) EnterMessageExpression(ctx *MessageExpressionContext) {}

// ExitMessageExpression is called when production messageExpression is exited.
func (s *BaseSmallTalkListener) ExitMessageExpression(ctx *MessageExpressionContext) {}

// EnterMethodExpression is called when production methodExpression is entered.
func (s *BaseSmallTalkListener) EnterMethodExpression(ctx *MethodExpressionContext) {}

// ExitMethodExpression is called when production methodExpression is exited.
func (s *BaseSmallTalkListener) ExitMethodExpression(ctx *MethodExpressionContext) {}

// EnterMethodSend is called when production methodSend is entered.
func (s *BaseSmallTalkListener) EnterMethodSend(ctx *MethodSendContext) {}

// ExitMethodSend is called when production methodSend is exited.
func (s *BaseSmallTalkListener) ExitMethodSend(ctx *MethodSendContext) {}

// EnterMethodSendRange is called when production methodSendRange is entered.
func (s *BaseSmallTalkListener) EnterMethodSendRange(ctx *MethodSendRangeContext) {}

// ExitMethodSendRange is called when production methodSendRange is exited.
func (s *BaseSmallTalkListener) ExitMethodSendRange(ctx *MethodSendRangeContext) {}

// EnterMethodSendItem is called when production methodSendItem is entered.
func (s *BaseSmallTalkListener) EnterMethodSendItem(ctx *MethodSendItemContext) {}

// ExitMethodSendItem is called when production methodSendItem is exited.
func (s *BaseSmallTalkListener) ExitMethodSendItem(ctx *MethodSendItemContext) {}

// EnterBinaryExpression is called when production binaryExpression is entered.
func (s *BaseSmallTalkListener) EnterBinaryExpression(ctx *BinaryExpressionContext) {}

// ExitBinaryExpression is called when production binaryExpression is exited.
func (s *BaseSmallTalkListener) ExitBinaryExpression(ctx *BinaryExpressionContext) {}

// EnterBinExprTailItem is called when production binExprTailItem is entered.
func (s *BaseSmallTalkListener) EnterBinExprTailItem(ctx *BinExprTailItemContext) {}

// ExitBinExprTailItem is called when production binExprTailItem is exited.
func (s *BaseSmallTalkListener) ExitBinExprTailItem(ctx *BinExprTailItemContext) {}

// EnterUnaryExpression is called when production unaryExpression is entered.
func (s *BaseSmallTalkListener) EnterUnaryExpression(ctx *UnaryExpressionContext) {}

// ExitUnaryExpression is called when production unaryExpression is exited.
func (s *BaseSmallTalkListener) ExitUnaryExpression(ctx *UnaryExpressionContext) {}

// EnterPrimaryLiteral is called when production primaryLiteral is entered.
func (s *BaseSmallTalkListener) EnterPrimaryLiteral(ctx *PrimaryLiteralContext) {}

// ExitPrimaryLiteral is called when production primaryLiteral is exited.
func (s *BaseSmallTalkListener) ExitPrimaryLiteral(ctx *PrimaryLiteralContext) {}

// EnterPrimaryIdentifier is called when production primaryIdentifier is entered.
func (s *BaseSmallTalkListener) EnterPrimaryIdentifier(ctx *PrimaryIdentifierContext) {}

// ExitPrimaryIdentifier is called when production primaryIdentifier is exited.
func (s *BaseSmallTalkListener) ExitPrimaryIdentifier(ctx *PrimaryIdentifierContext) {}

// EnterPrimaryBlock is called when production primaryBlock is entered.
func (s *BaseSmallTalkListener) EnterPrimaryBlock(ctx *PrimaryBlockContext) {}

// ExitPrimaryBlock is called when production primaryBlock is exited.
func (s *BaseSmallTalkListener) ExitPrimaryBlock(ctx *PrimaryBlockContext) {}

// EnterPrimaryExpression is called when production primaryExpression is entered.
func (s *BaseSmallTalkListener) EnterPrimaryExpression(ctx *PrimaryExpressionContext) {}

// ExitPrimaryExpression is called when production primaryExpression is exited.
func (s *BaseSmallTalkListener) ExitPrimaryExpression(ctx *PrimaryExpressionContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseSmallTalkListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseSmallTalkListener) ExitLiteral(ctx *LiteralContext) {}

// EnterLiteralNumber is called when production literalNumber is entered.
func (s *BaseSmallTalkListener) EnterLiteralNumber(ctx *LiteralNumberContext) {}

// ExitLiteralNumber is called when production literalNumber is exited.
func (s *BaseSmallTalkListener) ExitLiteralNumber(ctx *LiteralNumberContext) {}

// EnterLiteralChar is called when production literalChar is entered.
func (s *BaseSmallTalkListener) EnterLiteralChar(ctx *LiteralCharContext) {}

// ExitLiteralChar is called when production literalChar is exited.
func (s *BaseSmallTalkListener) ExitLiteralChar(ctx *LiteralCharContext) {}

// EnterLiteralSymbol is called when production literalSymbol is entered.
func (s *BaseSmallTalkListener) EnterLiteralSymbol(ctx *LiteralSymbolContext) {}

// ExitLiteralSymbol is called when production literalSymbol is exited.
func (s *BaseSmallTalkListener) ExitLiteralSymbol(ctx *LiteralSymbolContext) {}

// EnterSymbol is called when production symbol is entered.
func (s *BaseSmallTalkListener) EnterSymbol(ctx *SymbolContext) {}

// ExitSymbol is called when production symbol is exited.
func (s *BaseSmallTalkListener) ExitSymbol(ctx *SymbolContext) {}

// EnterLiteralArray is called when production literalArray is entered.
func (s *BaseSmallTalkListener) EnterLiteralArray(ctx *LiteralArrayContext) {}

// ExitLiteralArray is called when production literalArray is exited.
func (s *BaseSmallTalkListener) ExitLiteralArray(ctx *LiteralArrayContext) {}

// EnterLiteralArrayItem is called when production literalArrayItem is entered.
func (s *BaseSmallTalkListener) EnterLiteralArrayItem(ctx *LiteralArrayItemContext) {}

// ExitLiteralArrayItem is called when production literalArrayItem is exited.
func (s *BaseSmallTalkListener) ExitLiteralArrayItem(ctx *LiteralArrayItemContext) {}

// EnterLiteralString is called when production literalString is entered.
func (s *BaseSmallTalkListener) EnterLiteralString(ctx *LiteralStringContext) {}

// ExitLiteralString is called when production literalString is exited.
func (s *BaseSmallTalkListener) ExitLiteralString(ctx *LiteralStringContext) {}

// EnterLiteralNil is called when production literalNil is entered.
func (s *BaseSmallTalkListener) EnterLiteralNil(ctx *LiteralNilContext) {}

// ExitLiteralNil is called when production literalNil is exited.
func (s *BaseSmallTalkListener) ExitLiteralNil(ctx *LiteralNilContext) {}

// EnterLiteralSelf is called when production literalSelf is entered.
func (s *BaseSmallTalkListener) EnterLiteralSelf(ctx *LiteralSelfContext) {}

// ExitLiteralSelf is called when production literalSelf is exited.
func (s *BaseSmallTalkListener) ExitLiteralSelf(ctx *LiteralSelfContext) {}

// EnterLiteralBool is called when production literalBool is entered.
func (s *BaseSmallTalkListener) EnterLiteralBool(ctx *LiteralBoolContext) {}

// ExitLiteralBool is called when production literalBool is exited.
func (s *BaseSmallTalkListener) ExitLiteralBool(ctx *LiteralBoolContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseSmallTalkListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseSmallTalkListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterBinaryOperator is called when production binaryOperator is entered.
func (s *BaseSmallTalkListener) EnterBinaryOperator(ctx *BinaryOperatorContext) {}

// ExitBinaryOperator is called when production binaryOperator is exited.
func (s *BaseSmallTalkListener) ExitBinaryOperator(ctx *BinaryOperatorContext) {}

// EnterOpchar is called when production opchar is entered.
func (s *BaseSmallTalkListener) EnterOpchar(ctx *OpcharContext) {}

// ExitOpchar is called when production opchar is exited.
func (s *BaseSmallTalkListener) ExitOpchar(ctx *OpcharContext) {}
