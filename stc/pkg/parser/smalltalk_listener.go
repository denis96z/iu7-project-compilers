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

	// EnterInstanceVars is called when entering the instanceVars production.
	EnterInstanceVars(c *InstanceVarsContext)

	// EnterInstanceVar is called when entering the instanceVar production.
	EnterInstanceVar(c *InstanceVarContext)

	// EnterMethods is called when entering the methods production.
	EnterMethods(c *MethodsContext)

	// EnterMethod is called when entering the method production.
	EnterMethod(c *MethodContext)

	// EnterNamedMethod is called when entering the namedMethod production.
	EnterNamedMethod(c *NamedMethodContext)

	// EnterKeywordMethod is called when entering the keywordMethod production.
	EnterKeywordMethod(c *KeywordMethodContext)

	// EnterLocalVars is called when entering the localVars production.
	EnterLocalVars(c *LocalVarsContext)

	// EnterLocalVar is called when entering the localVar production.
	EnterLocalVar(c *LocalVarContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterBlockArgs is called when entering the blockArgs production.
	EnterBlockArgs(c *BlockArgsContext)

	// EnterBlockArg is called when entering the blockArg production.
	EnterBlockArg(c *BlockArgContext)

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

	// EnterMessageStatements is called when entering the messageStatements production.
	EnterMessageStatements(c *MessageStatementsContext)

	// EnterMessageStatement is called when entering the messageStatement production.
	EnterMessageStatement(c *MessageStatementContext)

	// EnterMessageExpression is called when entering the messageExpression production.
	EnterMessageExpression(c *MessageExpressionContext)

	// EnterMethodExpression is called when entering the methodExpression production.
	EnterMethodExpression(c *MethodExpressionContext)

	// EnterMethodSend is called when entering the methodSend production.
	EnterMethodSend(c *MethodSendContext)

	// EnterMethodSendRange is called when entering the methodSendRange production.
	EnterMethodSendRange(c *MethodSendRangeContext)

	// EnterMethodSendItem is called when entering the methodSendItem production.
	EnterMethodSendItem(c *MethodSendItemContext)

	// EnterBinaryExpression is called when entering the binaryExpression production.
	EnterBinaryExpression(c *BinaryExpressionContext)

	// EnterBinExprTailItem is called when entering the binExprTailItem production.
	EnterBinExprTailItem(c *BinExprTailItemContext)

	// EnterUnaryExpression is called when entering the unaryExpression production.
	EnterUnaryExpression(c *UnaryExpressionContext)

	// EnterPrimaryLiteral is called when entering the primaryLiteral production.
	EnterPrimaryLiteral(c *PrimaryLiteralContext)

	// EnterPrimaryIdentifier is called when entering the primaryIdentifier production.
	EnterPrimaryIdentifier(c *PrimaryIdentifierContext)

	// EnterPrimaryBlock is called when entering the primaryBlock production.
	EnterPrimaryBlock(c *PrimaryBlockContext)

	// EnterPrimaryExpression is called when entering the primaryExpression production.
	EnterPrimaryExpression(c *PrimaryExpressionContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterLiteralNumber is called when entering the literalNumber production.
	EnterLiteralNumber(c *LiteralNumberContext)

	// EnterLiteralChar is called when entering the literalChar production.
	EnterLiteralChar(c *LiteralCharContext)

	// EnterLiteralSymbol is called when entering the literalSymbol production.
	EnterLiteralSymbol(c *LiteralSymbolContext)

	// EnterSymbol is called when entering the symbol production.
	EnterSymbol(c *SymbolContext)

	// EnterLiteralArray is called when entering the literalArray production.
	EnterLiteralArray(c *LiteralArrayContext)

	// EnterLiteralArrayItem is called when entering the literalArrayItem production.
	EnterLiteralArrayItem(c *LiteralArrayItemContext)

	// EnterLiteralString is called when entering the literalString production.
	EnterLiteralString(c *LiteralStringContext)

	// EnterLiteralNil is called when entering the literalNil production.
	EnterLiteralNil(c *LiteralNilContext)

	// EnterLiteralSelf is called when entering the literalSelf production.
	EnterLiteralSelf(c *LiteralSelfContext)

	// EnterLiteralBool is called when entering the literalBool production.
	EnterLiteralBool(c *LiteralBoolContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterBinaryOperator is called when entering the binaryOperator production.
	EnterBinaryOperator(c *BinaryOperatorContext)

	// EnterOpchar is called when entering the opchar production.
	EnterOpchar(c *OpcharContext)

	// ExitScript is called when exiting the script production.
	ExitScript(c *ScriptContext)

	// ExitClassDef is called when exiting the classDef production.
	ExitClassDef(c *ClassDefContext)

	// ExitMain is called when exiting the main production.
	ExitMain(c *MainContext)

	// ExitInstanceVars is called when exiting the instanceVars production.
	ExitInstanceVars(c *InstanceVarsContext)

	// ExitInstanceVar is called when exiting the instanceVar production.
	ExitInstanceVar(c *InstanceVarContext)

	// ExitMethods is called when exiting the methods production.
	ExitMethods(c *MethodsContext)

	// ExitMethod is called when exiting the method production.
	ExitMethod(c *MethodContext)

	// ExitNamedMethod is called when exiting the namedMethod production.
	ExitNamedMethod(c *NamedMethodContext)

	// ExitKeywordMethod is called when exiting the keywordMethod production.
	ExitKeywordMethod(c *KeywordMethodContext)

	// ExitLocalVars is called when exiting the localVars production.
	ExitLocalVars(c *LocalVarsContext)

	// ExitLocalVar is called when exiting the localVar production.
	ExitLocalVar(c *LocalVarContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitBlockArgs is called when exiting the blockArgs production.
	ExitBlockArgs(c *BlockArgsContext)

	// ExitBlockArg is called when exiting the blockArg production.
	ExitBlockArg(c *BlockArgContext)

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

	// ExitMessageStatements is called when exiting the messageStatements production.
	ExitMessageStatements(c *MessageStatementsContext)

	// ExitMessageStatement is called when exiting the messageStatement production.
	ExitMessageStatement(c *MessageStatementContext)

	// ExitMessageExpression is called when exiting the messageExpression production.
	ExitMessageExpression(c *MessageExpressionContext)

	// ExitMethodExpression is called when exiting the methodExpression production.
	ExitMethodExpression(c *MethodExpressionContext)

	// ExitMethodSend is called when exiting the methodSend production.
	ExitMethodSend(c *MethodSendContext)

	// ExitMethodSendRange is called when exiting the methodSendRange production.
	ExitMethodSendRange(c *MethodSendRangeContext)

	// ExitMethodSendItem is called when exiting the methodSendItem production.
	ExitMethodSendItem(c *MethodSendItemContext)

	// ExitBinaryExpression is called when exiting the binaryExpression production.
	ExitBinaryExpression(c *BinaryExpressionContext)

	// ExitBinExprTailItem is called when exiting the binExprTailItem production.
	ExitBinExprTailItem(c *BinExprTailItemContext)

	// ExitUnaryExpression is called when exiting the unaryExpression production.
	ExitUnaryExpression(c *UnaryExpressionContext)

	// ExitPrimaryLiteral is called when exiting the primaryLiteral production.
	ExitPrimaryLiteral(c *PrimaryLiteralContext)

	// ExitPrimaryIdentifier is called when exiting the primaryIdentifier production.
	ExitPrimaryIdentifier(c *PrimaryIdentifierContext)

	// ExitPrimaryBlock is called when exiting the primaryBlock production.
	ExitPrimaryBlock(c *PrimaryBlockContext)

	// ExitPrimaryExpression is called when exiting the primaryExpression production.
	ExitPrimaryExpression(c *PrimaryExpressionContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitLiteralNumber is called when exiting the literalNumber production.
	ExitLiteralNumber(c *LiteralNumberContext)

	// ExitLiteralChar is called when exiting the literalChar production.
	ExitLiteralChar(c *LiteralCharContext)

	// ExitLiteralSymbol is called when exiting the literalSymbol production.
	ExitLiteralSymbol(c *LiteralSymbolContext)

	// ExitSymbol is called when exiting the symbol production.
	ExitSymbol(c *SymbolContext)

	// ExitLiteralArray is called when exiting the literalArray production.
	ExitLiteralArray(c *LiteralArrayContext)

	// ExitLiteralArrayItem is called when exiting the literalArrayItem production.
	ExitLiteralArrayItem(c *LiteralArrayItemContext)

	// ExitLiteralString is called when exiting the literalString production.
	ExitLiteralString(c *LiteralStringContext)

	// ExitLiteralNil is called when exiting the literalNil production.
	ExitLiteralNil(c *LiteralNilContext)

	// ExitLiteralSelf is called when exiting the literalSelf production.
	ExitLiteralSelf(c *LiteralSelfContext)

	// ExitLiteralBool is called when exiting the literalBool production.
	ExitLiteralBool(c *LiteralBoolContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitBinaryOperator is called when exiting the binaryOperator production.
	ExitBinaryOperator(c *BinaryOperatorContext)

	// ExitOpchar is called when exiting the opchar production.
	ExitOpchar(c *OpcharContext)
}
