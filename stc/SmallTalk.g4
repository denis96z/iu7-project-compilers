grammar SmallTalk;

NIL
    : 'nil'
    ;
TRUE
    : 'true'
    ;
FALSE
    : 'false'
    ;

CLASS
    : 'class'
    ;
SELF
    : 'self'
    ;
SUPER
    : 'super'
    ;

COMMENT
	: '"' ('""' | ~'"')* '"' -> skip
	;

WHITESPACE
    : [ \r\n\t]+ -> skip
    ;

DOT
    : '.'
    ;
PIPE
    : '|'
    ;
CARROT
    : '^'
    ;
COLON
    : ':'
    ;
SEMICOLON
    : ';'
    ;
ASSIGNMENT
    : ':='
    ;
LEFT_BRK
    : '['
    ;
RIGHT_BRK
    : ']'
    ;
LEFT_PRT
    : '('
    ;
RIGHT_PRT
    : ')'
    ;
LEFT_BRC
    : '{'
    ;
RIGHT_BRC
    : '}'
    ;
LEFT_GLM
    : '<'
    ;
RIGHT_GLM
    : '>'
    ;

PRIMITIVE
    : 'primitive:'
    ;

IDENTIFIER
    : [a-zA-Z][_a-zA-Z0-9]*
    ;

KEYWORD
    : IDENTIFIER COLON
    ;

HEX
    : '16r';

INT_NUMBER
    : [-]?[0-9]+
    ;
HEX_INT_NUMBER
    : HEX [-]?[0-9A-F]+
    ;
FLOAT_NUMBER
    : [-]?[0-9]+ '.' [0-9]+
    ;

SYMBOL_PREFIX
    : '#'
    ;

CHAR
    : '$' ~('@'|'\n'|'\t'|' ')
    ;
STRING
	: '\'' ('\'\'' | ~'\'')* '\''
	;

script
    : classDef* main EOF
    ;

classDef
	: CLASS name=IDENTIFIER (COLON sup=IDENTIFIER)? LEFT_BRK instanceVars? methods? RIGHT_BRK
	;

main
    : body;

instanceVars
    : PIPE instanceVar+ PIPE
    ;

instanceVar
    : name=IDENTIFIER
    ;

methods
    : method+
    ;

method
	: namedMethod
	| keywordMethod
	;

namedMethod
    : name=IDENTIFIER block
    ;

keywordMethod
    : name=KEYWORD block
    ;

localVars
	:	PIPE localVar+ PIPE
	;

localVar
    : name=IDENTIFIER
    ;

block
	: LEFT_BRK blockArgs? body RIGHT_BRK
	;

blockArgs
    : blockArg+ PIPE
    ;

blockArg
    : COLON name=IDENTIFIER
    ;

body
    : localVars? statements?
    ;

statements
    : statement (DOT statement)* DOT?
    ;

statement
    : assignmentStatement
	| messageStatements
	;

assignmentStatement
    : assignmentItem assignmentItem* messageExpression
    ;

assignmentItem
    : name=IDENTIFIER ASSIGNMENT
    ;

messageStatements
    : messageStatement
    ;

messageStatement
    : messageExpression
    ;

messageExpression
    : methodExpression
    ;

methodExpression
	: methodSend
	;

methodSend
    : recv=binaryExpression methodSendRange?
    ;

methodSendRange
    : methodSendItem (SEMICOLON? methodSendItem)*
    ;

methodSendItem
    : name=(IDENTIFIER|KEYWORD) (arg=binaryExpression)*
    ;

binaryExpression
	: unaryExpression binExprTailItem+
	;

binExprTailItem
    : op=binaryOperator unaryExpression
    ;

unaryExpression
	: primaryLiteral
	| primaryIdentifier
	| primaryBlock
	| primaryExpression
	;

primaryLiteral
    : literal
    ;

primaryIdentifier
    : identifier
    ;

primaryBlock
    : block
    ;

primaryExpression
    : LEFT_PRT statement RIGHT_PRT
    ;

literal
	: literalNumber
	| literalChar
	| literalString
	| literalSymbol
	| literalArray
	| literalNil
	| literalSelf
	| literalBool
	;

literalNumber
    : value=(INT_NUMBER | HEX_INT_NUMBER | FLOAT_NUMBER)
    ;

literalChar
    : value=CHAR
    ;

literalSymbol
    : symbol
    ;

symbol
    : SYMBOL_PREFIX name=IDENTIFIER
    ;

literalArray
    : SYMBOL_PREFIX LEFT_PRT literalArrayItem+ RIGHT_PRT
    ;

literalArrayItem
    : value=literal
    ;

literalString
    : value=STRING
    ;

literalNil
    : NIL
    ;

literalSelf
    : SELF
    ;

literalBool
    : value=(TRUE | FALSE)
    ;

identifier
	: name=IDENTIFIER
	;

binaryOperator
    : opchar
    | opchar opchar
    ;

opchar
	: '+' | '-' | '*' | '/' | '%'
	;
