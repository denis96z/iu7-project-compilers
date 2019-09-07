grammar DZ;

KW_PKG:    'pkg';
KW_USE:    'use';
KW_TYPE:   'type';
KW_FOR:    'for';
KW_WHILE:  'while';
KW_LOOP:   'loop';
KW_IF:     'if';
KW_ELIF:   'elif';
KW_ELSE:   'else';
KW_PROC:   'proc';
KW_FUNC:   'func';
KW_CONST:  'const';
KW_LET:    'let';
KW_MUT:    'mut';
KW_STRUCT: 'struct';
KW_ENUM:   'enum';
KW_UNION:  'union';
KW_RETURN: 'return';

I8_T:  'i8_t';
U8_T:  'u8_t';
I16_T: 'i16_t';
U16_T: 'u16_t';
I32_T: 'i32_t';
U32_T: 'u32_t';
I64_T: 'i64_t';
U64_T: 'u64_t';

CHAR_T: 'char_t';
BOOL_T: 'bool_t';

SIZE_T:  'size_t';
SSIZE_T: 'ssize_t';

LEFT_PRT:  '(';
RIGHT_PRT: ')';

LEFT_BRC:  '{';
RIGHT_BRC: '}';

LEFT_BRK:  '[';
RIGHT_BRK: ']';

DCOLON: '::';

COLON:     ':';
COMMA:     ',';
SEMICOLON: ';';

ADD: '+';
SUB: '-';
MUL: '*';
DIV: '/';
MOD: '%';
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

CONST:      [A-Z]([_]?[A-Z0-9])*;
TYPE:       [a-z]([_]?[a-z0-9])*[_][t];
IDENTIFIER: [a-z]([_]?[a-z0-9])*;

INT_CONST:   [-]?[0-9]+;
FLOAT_CONST: [-]?[0-9]+([.][0-9]+)?;
TRUE:        'true';
FALSE:       'false';

WHITESPACE: [ \r\n\t]+ -> skip;

start : pkg decls EOF;

pkg: KW_PKG name=IDENTIFIER SEMICOLON;

decls : decl*;
decl  : subdecl | complexdecl | typedecl | constdecl;

subdecl : procdecl | funcdecl;

procdecl : KW_PROC procheader block;
funcdecl : KW_FUNC funcheader block;

procheader : id=IDENTIFIER args;
funcheader : id=IDENTIFIER args funcret;

args     : LEFT_PRT argdecls RIGHT_PRT;
argdecls : (argdecl (COMMA argdecl)*)?;
argdecl  : id=IDENTIFIER COLON t=typespec;

funcret : COLON t=typespec;

complexdecl : structdecl | enumdecl | uniondecl;

structdecl  : KW_STRUCT id=TYPE LEFT_BRC structattrs RIGHT_BRC;
structattrs : structattr*;
structattr  : id=IDENTIFIER COLON t=typespec;

enumdecl    : KW_ENUM id=TYPE LEFT_BRC enumoptions RIGHT_BRC;
enumoptions : enumoption*;
enumoption  : id=CONST;

uniondecl : KW_UNION id=TYPE LEFT_BRC unionattrs RIGHT_BRC;
unionattrs : unionattr*;
unionattr  : id=IDENTIFIER COLON t=typespec;

constdecl : KW_CONST id=CONST COLON basictypespec ASGN constasgn SEMICOLON;
constasgn : casgn | intasgn | floatasgn | boolconst;
casgn     : id=CONST;
intasgn   : v=INT_CONST;
floatasgn : v=FLOAT_CONST;
boolconst : v=(TRUE | FALSE);

typedecl : KW_TYPE id=TYPE ASGN typespec SEMICOLON;

block : LEFT_BRC statements RIGHT_BRC;

typespec       : basictypespec | simpletypespec | reftypespec | arraytypespec;
basictypespec  : id=(
                     I8_T   | U8_T    |
                     I16_T  | U16_T   |
                     I32_T  | U32_T   |
                     I64_T  | U64_T   |
                     CHAR_T | BOOL_T  |
                     SIZE_T | SSIZE_T
                 );
simpletypespec : id=TYPE;
reftypespec    : REF id=TYPE;
arraytypespec  : LEFT_BRK id=TYPE COLON size=sizespec RIGHT_BRK;
sizespec       : INT_CONST | name=CONST;

statements : statement*;
statement  : assignment | condition | loop | retstatement SEMICOLON;

assignment : id=IDENTIFIER asgnop expression SEMICOLON;
asgnop     : id=(ASGN | ADD_ASGN | SUB_ASGN | MUL_ASGN | DIV_ASGN);

condition : ifblock elseblocks;

ifblock    : KW_IF expression block;
elseblocks : elifblock* elseblock?;
elifblock  : KW_ELIF expression block;
elseblock  : KW_ELSE block;

loop : trueloop;
trueloop : KW_LOOP block;

expression :;

retstatement
    : procretstatement
    | funcretstatement
    ;
procretstatement : KW_RETURN;
funcretstatement : KW_RETURN v=expression;
