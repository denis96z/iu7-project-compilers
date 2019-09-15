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

IDENTIFIER
    : [a-zA-Z][_a-zA-Z0-9]*
    ;
KEYWORD
    : IDENTIFIER COLON
    ;

BIN
    : '2r'
    ;
OCT
    : '8r'
    ;
HEX
    : '16r';

BIN_INT_NUMBER
    : BIN [-]?[0-1]+
    ;
OCT_INT_NUMBER
    : OCT [-]?[0-7]+
    ;
DEC_INT_NUMBER
    : [-]?[0-9]+
    ;
HEX_INT_NUMBER
    : HEX [-]?[0-9A-F]+
    ;
FLOAT_NUMBER
    : [-]?[0-9]+ DOT [0-9]+
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
	: CLASS name=IDENTIFIER (COLON sup=IDENTIFIER)? LEFT_BRK instanceMethods? RIGHT_BRK
	;

main
    : body;

instanceMethods
    : method+
    ;

method
	: unaryMethod
	| binaryMethod
	| keywordMethod
	;

unaryMethod
    : name=IDENTIFIER LEFT_BRK body RIGHT_BRK
    ;

binaryMethod
    : name=IDENTIFIER LEFT_BRK methodArg PIPE body RIGHT_BRK
    ;

keywordMethod
    : name=KEYWORD LEFT_BRK methodArgs body RIGHT_BRK
    ;

methodArgs
    : methodArg+ PIPE
    ;

methodArg
    : COLON name=IDENTIFIER
    ;

localVars
	: PIPE localVar+ PIPE
	;

localVar
    : name=IDENTIFIER
    ;

body
    : localVars? statements?
    ;

statements
    : statement (DOT statement)* DOT
    ;

statement
    : messageStatement
    | assignmentStatement
	;

assignmentStatement
    : assignmentItem assignmentItem* messageExpression
    ;

assignmentItem
    : name=IDENTIFIER ASSIGNMENT
    ;

messageStatement
    : messageExpression
    ;

messageExpression
    : messageExpression messageRange
    | prtStatement
    | value
    ;

message
    : keywordMessage
    | binaryMessage
    | unaryMessage
    ;

messageRange
    : message (SEMICOLON message)*
    ;

unaryMessage
    : methodName=IDENTIFIER
    ;

binaryMessage
    : (idName=identifier | opName=operator) messageArg
    ;

keywordMessage
    : methodName=KEYWORD keywordMessageArgs
    ;

keywordMessageArgs
    : messageArg+
    ;

messageArg
    : value | prtStatement
    ;

prtStatement
    : LEFT_PRT statement RIGHT_PRT
    ;

value
    : valueVar
    | valueSelf
    | valueSuper
    | valueBlock
    | valueConst
    ;

valueVar
    : IDENTIFIER
    ;

valueSelf
    : SELF
    ;

valueSuper
    : SUPER
    ;

valueBlock
	: LEFT_BRK blockArgs? body RIGHT_BRK
	;

blockArgs
    : blockArg+ PIPE
    ;

blockArg
    : COLON name=IDENTIFIER
    ;

valueConst
    : valueConstNum
    | valueConstChar
    | valueConstString
    | valueConstBool
    | valueNil
    ;

valueConstNum
    : valueConstBinInt
    | valueConstOctInt
    | valueConstDecInt
    | valueConstHexInt
    | valueConstFloat
    ;

valueConstBinInt
    : BIN_INT_NUMBER
    ;
valueConstOctInt
    : OCT_INT_NUMBER
    ;
valueConstDecInt
    : DEC_INT_NUMBER
    ;
valueConstHexInt
    : HEX_INT_NUMBER
    ;
valueConstFloat
    : FLOAT_NUMBER
    ;

valueConstChar
    : CHAR
    ;

valueConstString
    : STRING
    ;

valueConstBool
    : TRUE | FALSE
    ;

valueNil
    : NIL
    ;

identifier
    : IDENTIFIER
    ;

operator
    : '+' | '-' | '*' | '/'
    ;
