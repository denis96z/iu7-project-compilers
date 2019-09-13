grammar DZ;

KW_PKG:      'pkg';
KW_TYPE:     'type';
KW_FOR:      'for';
KW_WHILE:    'while';
KW_LOOP:     'loop';
KW_BREAK:    'break';
KW_CONTINUE: 'continue';
KW_IF:       'if';
KW_ELIF:     'elif';
KW_ELSE:     'else';
KW_PROC:     'proc';
KW_FUNC:     'func';
KW_CONST:    'const';
KW_LET:      'let';
KW_STRUCT:   'struct';
KW_ENUM:     'enum';
KW_RETURN:   'return';
KW_TRUE:     'true';
KW_FALSE:    'false';

I8_T:    'i8_t';
U8_T:    'u8_t';
I16_T:   'i16_t';
U16_T:   'u16_t';
I32_T:   'i32_t';
U32_T:   'u32_t';
I64_T:   'i64_t';
U64_T:   'u64_t';
CHAR_T:  'char_t';
BOOL_T:  'bool_t';
SIZE_T:  'size_t';
SSIZE_T: 'ssize_t';

LEFT_PRT:  '(';
RIGHT_PRT: ')';
LEFT_BRC:  '{';
RIGHT_BRC: '}';
LEFT_BRK:  '[';
RIGHT_BRK: ']';

COLON:     ':';
COMMA:     ',';
SEMICOLON: ';';

ADD: '+';
SUB: '-';
MUL: '*';
DIV: '/';
MOD: '%';

NOT: '~';
REF: '@';

SHL: '<<';
SHR: '>>';
AND: '&';
OR:  '|';
XOR: '^';

EQL:      '==';
NOT_EQL:  '!=';
GRT:      '>';
GRT_EQL:  '>=';
LESS:     '<';
LESS_EQL: '<=';

ASGN:     '=';
ADD_ASGN: '+=';
SUB_ASGN: '-=';
MUL_ASGN: '*=';
DIV_ASGN: '/=';
MOD_ASGN: '%=';

SHL_ASGN: '<<=';
SHR_ASGN: '>>=';
AND_ASGN: '&=';
OR_ASGN:  '|=';
XOR_ASGN: '^=';

CONST:      [A-Z]([_]?[A-Z0-9])*;
INT_VALUE:  [-]?[0-9]+;
TYPE:       [a-z]([_]?[a-z0-9])*[_][t];
IDENTIFIER: [a-z]([_]?[a-z0-9])*;

WHITESPACE: [ \r\n\t]+ -> skip;

start
    : pkg decl* EOF
    ;

pkg
    : KW_PKG name=IDENTIFIER SEMICOLON
    ;

decl
    : constDecl
    | typeDecl
    | enumDecl
    | structDecl
    | procDecl
    | funcDecl
    ;


constDecl
    : KW_CONST name=CONST COLON tName=basicTypeSpec ASGN value=(INT_VALUE | KW_TRUE | KW_FALSE | CONST) SEMICOLON
    ;

typeDecl
    : KW_TYPE name=TYPE ASGN tName=typeSpec SEMICOLON
    ;

typeSpec
    : simpleTypeSpec
    | refTypeSpec
    | arrayTypeSpec
    | sliceTypeSpec
    ;

simpleTypeSpec
    : basicTypeSpec
    | namedTypeSpec
    ;

basicTypeSpec
    : I8_T | U8_T
    | I16_T | U16_T
    | I32_T | U32_T
    | I64_T | U64_T
    | SIZE_T | SSIZE_T
    | CHAR_T | BOOL_T
    ;

namedTypeSpec
    : name=TYPE
    ;

refTypeSpec
    : REF tName=simpleTypeSpec
    ;

arrayTypeSpec
    : LEFT_BRK tName=simpleTypeSpec COLON size=(INT_VALUE | CONST) RIGHT_BRK
    ;

sliceTypeSpec
    : LEFT_BRK tName=simpleTypeSpec RIGHT_BRK
    ;

enumDecl
    : KW_ENUM name=TYPE LEFT_BRC enumOption (COMMA enumOption)* RIGHT_BRC
    ;

enumOption
    : name=CONST
    ;

structDecl
    : KW_STRUCT name=TYPE LEFT_BRC structAttr* RIGHT_BRC
    ;

structAttr
    : name=IDENTIFIER COLON tName=typeSpec SEMICOLON
    ;

procDecl
    : KW_PROC name=IDENTIFIER LEFT_PRT (procArg (COMMA procArg)*)? RIGHT_PRT body=block
    ;

procArg
    : name=IDENTIFIER COLON tName=typeSpec
    ;

funcDecl
    : KW_FUNC name=IDENTIFIER LEFT_PRT (funcArg (COMMA funcArg)*)? COLON tName=typeSpec RIGHT_PRT body=block
    ;

funcArg
    : name=IDENTIFIER COLON tName=typeSpec
    ;

block
    : statement*;

statement
    : condition
    | loop
    | breakStatement
    | continueStatement
    | procCall
    | varDecl SEMICOLON
    | expression SEMICOLON
    | returnStatement
    ;

condition
    : ifConditionBranch elifConditionBranch* elseConditionBranch?;

ifConditionBranch
    : KW_IF cond=expression LEFT_BRC body=block RIGHT_BRC;

elifConditionBranch
    : KW_ELIF cond=expression LEFT_BRC body=block RIGHT_BRC;

elseConditionBranch
    : KW_ELSE LEFT_BRC body=block RIGHT_BRC;

loop
    : trueLoop | whileLoop;

trueLoop
    : KW_LOOP LEFT_BRC body=block RIGHT_BRC;

whileLoop
    : KW_WHILE cond=expression LEFT_BRC body=block RIGHT_BRC;

breakStatement
    : KW_BREAK
    ;

continueStatement
    : KW_CONTINUE
    ;

varDecl
    : KW_LET name=IDENTIFIER ASGN value=expression;

procCall
    : procName=IDENTIFIER LEFT_PRT (procParam (COMMA procParam)*)? RIGHT_PRT SEMICOLON;

procParam
    : value = expression;

expression
    : varName=IDENTIFIER
    | intValue=INT_VALUE
    | boolValue=(KW_TRUE | KW_FALSE)
    | constName=CONST
    | LEFT_BRK brkExpr=expression RIGHT_BRK
    | LEFT_PRT prtExpr=expression RIGHT_PRT
    | funcName=IDENTIFIER LEFT_PRT (funcParam (COMMA funcParam)*)? RIGHT_PRT
    | lBinExpr=expression binOp=binaryOperator rBinExpr=expression
    | unOp=unaryOperator unExpr=expression
    ;

funcParam
    : value=expression;

unaryOperator
    : SUB
    | NOT
    ;

binaryOperator
    : ASGN
    | ADD | SUB | MUL | DIV | MOD
    | AND | OR  | XOR | SHL | SHR
    | ADD_ASGN  | SUB_ASGN  | MUL_ASGN | DIV_ASGN | MOD_ASGN
    | AND_ASGN  | OR_ASGN   | XOR_ASGN | SHL_ASGN | SHR_ASGN
    | EQL | GRT | LESS | GRT_EQL | LESS_EQL
    ;

returnStatement
    : KW_RETURN
    | KW_RETURN value=expression SEMICOLON;
