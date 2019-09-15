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

// EnterInstanceMethods is called when production instanceMethods is entered.
func (s *BaseSmallTalkListener) EnterInstanceMethods(ctx *InstanceMethodsContext) {}

// ExitInstanceMethods is called when production instanceMethods is exited.
func (s *BaseSmallTalkListener) ExitInstanceMethods(ctx *InstanceMethodsContext) {}

// EnterMethod is called when production method is entered.
func (s *BaseSmallTalkListener) EnterMethod(ctx *MethodContext) {}

// ExitMethod is called when production method is exited.
func (s *BaseSmallTalkListener) ExitMethod(ctx *MethodContext) {}

// EnterUnaryMethod is called when production unaryMethod is entered.
func (s *BaseSmallTalkListener) EnterUnaryMethod(ctx *UnaryMethodContext) {}

// ExitUnaryMethod is called when production unaryMethod is exited.
func (s *BaseSmallTalkListener) ExitUnaryMethod(ctx *UnaryMethodContext) {}

// EnterBinaryMethod is called when production binaryMethod is entered.
func (s *BaseSmallTalkListener) EnterBinaryMethod(ctx *BinaryMethodContext) {}

// ExitBinaryMethod is called when production binaryMethod is exited.
func (s *BaseSmallTalkListener) ExitBinaryMethod(ctx *BinaryMethodContext) {}

// EnterKeywordMethod is called when production keywordMethod is entered.
func (s *BaseSmallTalkListener) EnterKeywordMethod(ctx *KeywordMethodContext) {}

// ExitKeywordMethod is called when production keywordMethod is exited.
func (s *BaseSmallTalkListener) ExitKeywordMethod(ctx *KeywordMethodContext) {}

// EnterMethodArgs is called when production methodArgs is entered.
func (s *BaseSmallTalkListener) EnterMethodArgs(ctx *MethodArgsContext) {}

// ExitMethodArgs is called when production methodArgs is exited.
func (s *BaseSmallTalkListener) ExitMethodArgs(ctx *MethodArgsContext) {}

// EnterMethodArg is called when production methodArg is entered.
func (s *BaseSmallTalkListener) EnterMethodArg(ctx *MethodArgContext) {}

// ExitMethodArg is called when production methodArg is exited.
func (s *BaseSmallTalkListener) ExitMethodArg(ctx *MethodArgContext) {}

// EnterLocalVars is called when production localVars is entered.
func (s *BaseSmallTalkListener) EnterLocalVars(ctx *LocalVarsContext) {}

// ExitLocalVars is called when production localVars is exited.
func (s *BaseSmallTalkListener) ExitLocalVars(ctx *LocalVarsContext) {}

// EnterLocalVar is called when production localVar is entered.
func (s *BaseSmallTalkListener) EnterLocalVar(ctx *LocalVarContext) {}

// ExitLocalVar is called when production localVar is exited.
func (s *BaseSmallTalkListener) ExitLocalVar(ctx *LocalVarContext) {}

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

// EnterMessageStatement is called when production messageStatement is entered.
func (s *BaseSmallTalkListener) EnterMessageStatement(ctx *MessageStatementContext) {}

// ExitMessageStatement is called when production messageStatement is exited.
func (s *BaseSmallTalkListener) ExitMessageStatement(ctx *MessageStatementContext) {}

// EnterMessageExpression is called when production messageExpression is entered.
func (s *BaseSmallTalkListener) EnterMessageExpression(ctx *MessageExpressionContext) {}

// ExitMessageExpression is called when production messageExpression is exited.
func (s *BaseSmallTalkListener) ExitMessageExpression(ctx *MessageExpressionContext) {}

// EnterMessage is called when production message is entered.
func (s *BaseSmallTalkListener) EnterMessage(ctx *MessageContext) {}

// ExitMessage is called when production message is exited.
func (s *BaseSmallTalkListener) ExitMessage(ctx *MessageContext) {}

// EnterMessageRange is called when production messageRange is entered.
func (s *BaseSmallTalkListener) EnterMessageRange(ctx *MessageRangeContext) {}

// ExitMessageRange is called when production messageRange is exited.
func (s *BaseSmallTalkListener) ExitMessageRange(ctx *MessageRangeContext) {}

// EnterUnaryMessage is called when production unaryMessage is entered.
func (s *BaseSmallTalkListener) EnterUnaryMessage(ctx *UnaryMessageContext) {}

// ExitUnaryMessage is called when production unaryMessage is exited.
func (s *BaseSmallTalkListener) ExitUnaryMessage(ctx *UnaryMessageContext) {}

// EnterBinaryMessage is called when production binaryMessage is entered.
func (s *BaseSmallTalkListener) EnterBinaryMessage(ctx *BinaryMessageContext) {}

// ExitBinaryMessage is called when production binaryMessage is exited.
func (s *BaseSmallTalkListener) ExitBinaryMessage(ctx *BinaryMessageContext) {}

// EnterKeywordMessage is called when production keywordMessage is entered.
func (s *BaseSmallTalkListener) EnterKeywordMessage(ctx *KeywordMessageContext) {}

// ExitKeywordMessage is called when production keywordMessage is exited.
func (s *BaseSmallTalkListener) ExitKeywordMessage(ctx *KeywordMessageContext) {}

// EnterKeywordMessageArgs is called when production keywordMessageArgs is entered.
func (s *BaseSmallTalkListener) EnterKeywordMessageArgs(ctx *KeywordMessageArgsContext) {}

// ExitKeywordMessageArgs is called when production keywordMessageArgs is exited.
func (s *BaseSmallTalkListener) ExitKeywordMessageArgs(ctx *KeywordMessageArgsContext) {}

// EnterMessageArg is called when production messageArg is entered.
func (s *BaseSmallTalkListener) EnterMessageArg(ctx *MessageArgContext) {}

// ExitMessageArg is called when production messageArg is exited.
func (s *BaseSmallTalkListener) ExitMessageArg(ctx *MessageArgContext) {}

// EnterPrtStatement is called when production prtStatement is entered.
func (s *BaseSmallTalkListener) EnterPrtStatement(ctx *PrtStatementContext) {}

// ExitPrtStatement is called when production prtStatement is exited.
func (s *BaseSmallTalkListener) ExitPrtStatement(ctx *PrtStatementContext) {}

// EnterValue is called when production value is entered.
func (s *BaseSmallTalkListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseSmallTalkListener) ExitValue(ctx *ValueContext) {}

// EnterValueVar is called when production valueVar is entered.
func (s *BaseSmallTalkListener) EnterValueVar(ctx *ValueVarContext) {}

// ExitValueVar is called when production valueVar is exited.
func (s *BaseSmallTalkListener) ExitValueVar(ctx *ValueVarContext) {}

// EnterValueSelf is called when production valueSelf is entered.
func (s *BaseSmallTalkListener) EnterValueSelf(ctx *ValueSelfContext) {}

// ExitValueSelf is called when production valueSelf is exited.
func (s *BaseSmallTalkListener) ExitValueSelf(ctx *ValueSelfContext) {}

// EnterValueSuper is called when production valueSuper is entered.
func (s *BaseSmallTalkListener) EnterValueSuper(ctx *ValueSuperContext) {}

// ExitValueSuper is called when production valueSuper is exited.
func (s *BaseSmallTalkListener) ExitValueSuper(ctx *ValueSuperContext) {}

// EnterValueBlock is called when production valueBlock is entered.
func (s *BaseSmallTalkListener) EnterValueBlock(ctx *ValueBlockContext) {}

// ExitValueBlock is called when production valueBlock is exited.
func (s *BaseSmallTalkListener) ExitValueBlock(ctx *ValueBlockContext) {}

// EnterBlockArgs is called when production blockArgs is entered.
func (s *BaseSmallTalkListener) EnterBlockArgs(ctx *BlockArgsContext) {}

// ExitBlockArgs is called when production blockArgs is exited.
func (s *BaseSmallTalkListener) ExitBlockArgs(ctx *BlockArgsContext) {}

// EnterBlockArg is called when production blockArg is entered.
func (s *BaseSmallTalkListener) EnterBlockArg(ctx *BlockArgContext) {}

// ExitBlockArg is called when production blockArg is exited.
func (s *BaseSmallTalkListener) ExitBlockArg(ctx *BlockArgContext) {}

// EnterValueConst is called when production valueConst is entered.
func (s *BaseSmallTalkListener) EnterValueConst(ctx *ValueConstContext) {}

// ExitValueConst is called when production valueConst is exited.
func (s *BaseSmallTalkListener) ExitValueConst(ctx *ValueConstContext) {}

// EnterValueConstNum is called when production valueConstNum is entered.
func (s *BaseSmallTalkListener) EnterValueConstNum(ctx *ValueConstNumContext) {}

// ExitValueConstNum is called when production valueConstNum is exited.
func (s *BaseSmallTalkListener) ExitValueConstNum(ctx *ValueConstNumContext) {}

// EnterValueConstBinInt is called when production valueConstBinInt is entered.
func (s *BaseSmallTalkListener) EnterValueConstBinInt(ctx *ValueConstBinIntContext) {}

// ExitValueConstBinInt is called when production valueConstBinInt is exited.
func (s *BaseSmallTalkListener) ExitValueConstBinInt(ctx *ValueConstBinIntContext) {}

// EnterValueConstOctInt is called when production valueConstOctInt is entered.
func (s *BaseSmallTalkListener) EnterValueConstOctInt(ctx *ValueConstOctIntContext) {}

// ExitValueConstOctInt is called when production valueConstOctInt is exited.
func (s *BaseSmallTalkListener) ExitValueConstOctInt(ctx *ValueConstOctIntContext) {}

// EnterValueConstDecInt is called when production valueConstDecInt is entered.
func (s *BaseSmallTalkListener) EnterValueConstDecInt(ctx *ValueConstDecIntContext) {}

// ExitValueConstDecInt is called when production valueConstDecInt is exited.
func (s *BaseSmallTalkListener) ExitValueConstDecInt(ctx *ValueConstDecIntContext) {}

// EnterValueConstHexInt is called when production valueConstHexInt is entered.
func (s *BaseSmallTalkListener) EnterValueConstHexInt(ctx *ValueConstHexIntContext) {}

// ExitValueConstHexInt is called when production valueConstHexInt is exited.
func (s *BaseSmallTalkListener) ExitValueConstHexInt(ctx *ValueConstHexIntContext) {}

// EnterValueConstFloat is called when production valueConstFloat is entered.
func (s *BaseSmallTalkListener) EnterValueConstFloat(ctx *ValueConstFloatContext) {}

// ExitValueConstFloat is called when production valueConstFloat is exited.
func (s *BaseSmallTalkListener) ExitValueConstFloat(ctx *ValueConstFloatContext) {}

// EnterValueConstChar is called when production valueConstChar is entered.
func (s *BaseSmallTalkListener) EnterValueConstChar(ctx *ValueConstCharContext) {}

// ExitValueConstChar is called when production valueConstChar is exited.
func (s *BaseSmallTalkListener) ExitValueConstChar(ctx *ValueConstCharContext) {}

// EnterValueConstString is called when production valueConstString is entered.
func (s *BaseSmallTalkListener) EnterValueConstString(ctx *ValueConstStringContext) {}

// ExitValueConstString is called when production valueConstString is exited.
func (s *BaseSmallTalkListener) ExitValueConstString(ctx *ValueConstStringContext) {}

// EnterValueConstBool is called when production valueConstBool is entered.
func (s *BaseSmallTalkListener) EnterValueConstBool(ctx *ValueConstBoolContext) {}

// ExitValueConstBool is called when production valueConstBool is exited.
func (s *BaseSmallTalkListener) ExitValueConstBool(ctx *ValueConstBoolContext) {}

// EnterValueNil is called when production valueNil is entered.
func (s *BaseSmallTalkListener) EnterValueNil(ctx *ValueNilContext) {}

// ExitValueNil is called when production valueNil is exited.
func (s *BaseSmallTalkListener) ExitValueNil(ctx *ValueNilContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseSmallTalkListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseSmallTalkListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterOperator is called when production operator is entered.
func (s *BaseSmallTalkListener) EnterOperator(ctx *OperatorContext) {}

// ExitOperator is called when production operator is exited.
func (s *BaseSmallTalkListener) ExitOperator(ctx *OperatorContext) {}
