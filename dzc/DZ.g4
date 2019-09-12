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

SHL: '<<';
SHR: '>>';
AND: '&';
OR:  '|';
XOR: '^';

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

REF: '@';

INT_CONST:   [-]?[0-9]+;
TRUE:        'true';
FALSE:       'false';

CONST:      [A-Z]([_]?[A-Z0-9])*;
TYPE:       [a-z]([_]?[a-z0-9])*[_][t];
IDENTIFIER: [a-z]([_]?[a-z0-9])*;

WHITESPACE: [ \r\n\t]+ -> skip;

start : pkg decls EOF;

pkg: KW_PKG name=IDENTIFIER SEMICOLON;

decls : decl*;
decl  : constdecl | typedecl | complexdecl | subdecl;

subdecl : procdecl | funcdecl;

procdecl : KW_PROC procheader block;
funcdecl : KW_FUNC funcheader block;

procheader : id=IDENTIFIER args;
funcheader : id=IDENTIFIER args funcret;

args     : LEFT_PRT argdecls RIGHT_PRT;
argdecls : (argdecl (COMMA argdecl)*)?;
argdecl  : id=IDENTIFIER COLON t=simpletypespec;

funcret : COLON t=simpletypespec;

complexdecl : enumdecl | structdecl;

enumdecl    : KW_ENUM id=TYPE LEFT_BRC enumoptions RIGHT_BRC;
enumoptions : enumoption+;
enumoption  : id=CONST COMMA;

structdecl  : KW_STRUCT id=TYPE LEFT_BRC structattrs RIGHT_BRC;
structattrs : structattr*;
structattr  : id=IDENTIFIER COLON t=typespec SEMICOLON;

constdecl : KW_CONST id=CONST COLON t=basictypespec ASGN v=constexpr SEMICOLON;
constexpr : intexpr | boolexpr | constval;
intexpr   : v=INT_CONST;
boolexpr  : v=(TRUE | FALSE);
constval  : id=CONST;

typedecl : KW_TYPE id=TYPE ASGN t=typespec SEMICOLON;

typespec       : simpletypespec | reftypespec | arraytypespec;
simpletypespec : basictypespec | namedtypespec;
basictypespec  : id=(
                     I8_T   | U8_T    |
                     I16_T  | U16_T   |
                     I32_T  | U32_T   |
                     I64_T  | U64_T   |
                     CHAR_T | BOOL_T  |
                     SIZE_T | SSIZE_T
                 );
namedtypespec  : id=TYPE;
reftypespec    : REF id=simpletypespec;
arraytypespec  : LEFT_BRK id=simpletypespec COLON size=sizespec RIGHT_BRK;
sizespec       : INT_CONST | name=CONST;

block
    : statement*;

statement
    : condition
    | loop
    | breakStatement
    | continueStatement
    | procCall
    | declaration SEMICOLON
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

declaration
    : KW_LET name=IDENTIFIER ASGN value=expression;

procCall
    : procName=IDENTIFIER LEFT_PRT (procParam (COMMA procParam)*)? RIGHT_PRT SEMICOLON;

procParam
    : value = expression;

expression
    : varName=IDENTIFIER
    | intValue=INT_CONST
    | boolValue=(TRUE | FALSE)
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
    ;

returnStatement
    : KW_RETURN
    | KW_RETURN value=expression SEMICOLON;
