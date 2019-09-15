// Code generated from /home/denis/Documents/iu7-project-compilers/stc/SmallTalk.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // SmallTalk

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SmallTalkListener is a complete listener for a parse tree produced by SmallTalkParser.
type SmallTalkListener interface {
	antlr.ParseTreeListener

	// EnterScript is called when entering the script production.
	EnterScript(c *ScriptContext)

	// EnterClassDef is called when entering the classDef production.
	EnterClassDef(c *ClassDefContext)

	// EnterMain is called when entering the main production.
	EnterMain(c *MainContext)

	// EnterInstanceMethods is called when entering the instanceMethods production.
	EnterInstanceMethods(c *InstanceMethodsContext)

	// EnterMethod is called when entering the method production.
	EnterMethod(c *MethodContext)

	// EnterUnaryMethod is called when entering the unaryMethod production.
	EnterUnaryMethod(c *UnaryMethodContext)

	// EnterBinaryMethod is called when entering the binaryMethod production.
	EnterBinaryMethod(c *BinaryMethodContext)

	// EnterKeywordMethod is called when entering the keywordMethod production.
	EnterKeywordMethod(c *KeywordMethodContext)

	// EnterMethodArgs is called when entering the methodArgs production.
	EnterMethodArgs(c *MethodArgsContext)

	// EnterMethodArg is called when entering the methodArg production.
	EnterMethodArg(c *MethodArgContext)

	// EnterLocalVars is called when entering the localVars production.
	EnterLocalVars(c *LocalVarsContext)

	// EnterLocalVar is called when entering the localVar production.
	EnterLocalVar(c *LocalVarContext)

	// EnterBody is called when entering the body production.
	EnterBody(c *BodyContext)

	// EnterStatements is called when entering the statements production.
	EnterStatements(c *StatementsContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterAssignmentStatement is called when entering the assignmentStatement production.
	EnterAssignmentStatement(c *AssignmentStatementContext)

	// EnterAssignmentItem is called when entering the assignmentItem production.
	EnterAssignmentItem(c *AssignmentItemContext)

	// EnterMessageStatement is called when entering the messageStatement production.
	EnterMessageStatement(c *MessageStatementContext)

	// EnterMessageExpression is called when entering the messageExpression production.
	EnterMessageExpression(c *MessageExpressionContext)

	// EnterMessage is called when entering the message production.
	EnterMessage(c *MessageContext)

	// EnterMessageRange is called when entering the messageRange production.
	EnterMessageRange(c *MessageRangeContext)

	// EnterUnaryMessage is called when entering the unaryMessage production.
	EnterUnaryMessage(c *UnaryMessageContext)

	// EnterBinaryMessage is called when entering the binaryMessage production.
	EnterBinaryMessage(c *BinaryMessageContext)

	// EnterKeywordMessage is called when entering the keywordMessage production.
	EnterKeywordMessage(c *KeywordMessageContext)

	// EnterKeywordMessageArgs is called when entering the keywordMessageArgs production.
	EnterKeywordMessageArgs(c *KeywordMessageArgsContext)

	// EnterMessageArg is called when entering the messageArg production.
	EnterMessageArg(c *MessageArgContext)

	// EnterPrtStatement is called when entering the prtStatement production.
	EnterPrtStatement(c *PrtStatementContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterValueVar is called when entering the valueVar production.
	EnterValueVar(c *ValueVarContext)

	// EnterValueSelf is called when entering the valueSelf production.
	EnterValueSelf(c *ValueSelfContext)

	// EnterValueSuper is called when entering the valueSuper production.
	EnterValueSuper(c *ValueSuperContext)

	// EnterValueBlock is called when entering the valueBlock production.
	EnterValueBlock(c *ValueBlockContext)

	// EnterBlockArgs is called when entering the blockArgs production.
	EnterBlockArgs(c *BlockArgsContext)

	// EnterBlockArg is called when entering the blockArg production.
	EnterBlockArg(c *BlockArgContext)

	// EnterValueConst is called when entering the valueConst production.
	EnterValueConst(c *ValueConstContext)

	// EnterValueConstNum is called when entering the valueConstNum production.
	EnterValueConstNum(c *ValueConstNumContext)

	// EnterValueConstBinInt is called when entering the valueConstBinInt production.
	EnterValueConstBinInt(c *ValueConstBinIntContext)

	// EnterValueConstOctInt is called when entering the valueConstOctInt production.
	EnterValueConstOctInt(c *ValueConstOctIntContext)

	// EnterValueConstDecInt is called when entering the valueConstDecInt production.
	EnterValueConstDecInt(c *ValueConstDecIntContext)

	// EnterValueConstHexInt is called when entering the valueConstHexInt production.
	EnterValueConstHexInt(c *ValueConstHexIntContext)

	// EnterValueConstFloat is called when entering the valueConstFloat production.
	EnterValueConstFloat(c *ValueConstFloatContext)

	// EnterValueConstChar is called when entering the valueConstChar production.
	EnterValueConstChar(c *ValueConstCharContext)

	// EnterValueConstString is called when entering the valueConstString production.
	EnterValueConstString(c *ValueConstStringContext)

	// EnterValueConstBool is called when entering the valueConstBool production.
	EnterValueConstBool(c *ValueConstBoolContext)

	// EnterValueNil is called when entering the valueNil production.
	EnterValueNil(c *ValueNilContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterOperator is called when entering the operator production.
	EnterOperator(c *OperatorContext)

	// ExitScript is called when exiting the script production.
	ExitScript(c *ScriptContext)

	// ExitClassDef is called when exiting the classDef production.
	ExitClassDef(c *ClassDefContext)

	// ExitMain is called when exiting the main production.
	ExitMain(c *MainContext)

	// ExitInstanceMethods is called when exiting the instanceMethods production.
	ExitInstanceMethods(c *InstanceMethodsContext)

	// ExitMethod is called when exiting the method production.
	ExitMethod(c *MethodContext)

	// ExitUnaryMethod is called when exiting the unaryMethod production.
	ExitUnaryMethod(c *UnaryMethodContext)

	// ExitBinaryMethod is called when exiting the binaryMethod production.
	ExitBinaryMethod(c *BinaryMethodContext)

	// ExitKeywordMethod is called when exiting the keywordMethod production.
	ExitKeywordMethod(c *KeywordMethodContext)

	// ExitMethodArgs is called when exiting the methodArgs production.
	ExitMethodArgs(c *MethodArgsContext)

	// ExitMethodArg is called when exiting the methodArg production.
	ExitMethodArg(c *MethodArgContext)

	// ExitLocalVars is called when exiting the localVars production.
	ExitLocalVars(c *LocalVarsContext)

	// ExitLocalVar is called when exiting the localVar production.
	ExitLocalVar(c *LocalVarContext)

	// ExitBody is called when exiting the body production.
	ExitBody(c *BodyContext)

	// ExitStatements is called when exiting the statements production.
	ExitStatements(c *StatementsContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitAssignmentStatement is called when exiting the assignmentStatement production.
	ExitAssignmentStatement(c *AssignmentStatementContext)

	// ExitAssignmentItem is called when exiting the assignmentItem production.
	ExitAssignmentItem(c *AssignmentItemContext)

	// ExitMessageStatement is called when exiting the messageStatement production.
	ExitMessageStatement(c *MessageStatementContext)

	// ExitMessageExpression is called when exiting the messageExpression production.
	ExitMessageExpression(c *MessageExpressionContext)

	// ExitMessage is called when exiting the message production.
	ExitMessage(c *MessageContext)

	// ExitMessageRange is called when exiting the messageRange production.
	ExitMessageRange(c *MessageRangeContext)

	// ExitUnaryMessage is called when exiting the unaryMessage production.
	ExitUnaryMessage(c *UnaryMessageContext)

	// ExitBinaryMessage is called when exiting the binaryMessage production.
	ExitBinaryMessage(c *BinaryMessageContext)

	// ExitKeywordMessage is called when exiting the keywordMessage production.
	ExitKeywordMessage(c *KeywordMessageContext)

	// ExitKeywordMessageArgs is called when exiting the keywordMessageArgs production.
	ExitKeywordMessageArgs(c *KeywordMessageArgsContext)

	// ExitMessageArg is called when exiting the messageArg production.
	ExitMessageArg(c *MessageArgContext)

	// ExitPrtStatement is called when exiting the prtStatement production.
	ExitPrtStatement(c *PrtStatementContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitValueVar is called when exiting the valueVar production.
	ExitValueVar(c *ValueVarContext)

	// ExitValueSelf is called when exiting the valueSelf production.
	ExitValueSelf(c *ValueSelfContext)

	// ExitValueSuper is called when exiting the valueSuper production.
	ExitValueSuper(c *ValueSuperContext)

	// ExitValueBlock is called when exiting the valueBlock production.
	ExitValueBlock(c *ValueBlockContext)

	// ExitBlockArgs is called when exiting the blockArgs production.
	ExitBlockArgs(c *BlockArgsContext)

	// ExitBlockArg is called when exiting the blockArg production.
	ExitBlockArg(c *BlockArgContext)

	// ExitValueConst is called when exiting the valueConst production.
	ExitValueConst(c *ValueConstContext)

	// ExitValueConstNum is called when exiting the valueConstNum production.
	ExitValueConstNum(c *ValueConstNumContext)

	// ExitValueConstBinInt is called when exiting the valueConstBinInt production.
	ExitValueConstBinInt(c *ValueConstBinIntContext)

	// ExitValueConstOctInt is called when exiting the valueConstOctInt production.
	ExitValueConstOctInt(c *ValueConstOctIntContext)

	// ExitValueConstDecInt is called when exiting the valueConstDecInt production.
	ExitValueConstDecInt(c *ValueConstDecIntContext)

	// ExitValueConstHexInt is called when exiting the valueConstHexInt production.
	ExitValueConstHexInt(c *ValueConstHexIntContext)

	// ExitValueConstFloat is called when exiting the valueConstFloat production.
	ExitValueConstFloat(c *ValueConstFloatContext)

	// ExitValueConstChar is called when exiting the valueConstChar production.
	ExitValueConstChar(c *ValueConstCharContext)

	// ExitValueConstString is called when exiting the valueConstString production.
	ExitValueConstString(c *ValueConstStringContext)

	// ExitValueConstBool is called when exiting the valueConstBool production.
	ExitValueConstBool(c *ValueConstBoolContext)

	// ExitValueNil is called when exiting the valueNil production.
	ExitValueNil(c *ValueNilContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitOperator is called when exiting the operator production.
	ExitOperator(c *OperatorContext)
}
