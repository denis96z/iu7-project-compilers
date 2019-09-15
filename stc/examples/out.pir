.namespace ['B']
.sub 'get_string' :method :vtable
    .return('B')
.end
.sub 'doSth' :method
    .local pmc x
    .lex 'x', x
    .local pmc y
    .lex 'y', y
    $P1 = box 1
    y = $P1
    x = $P1
    say x
    .return(y)
.end
.sub 'doSth2' :method
    $P3 = box 1
    $P4 = box 0b10
    $P5 = $P3 + $P4
    $P6 = box 0xFF
    $P7 = $P5 * $P6
    $P8 = self.'doSth'()
    $P9 = box 'MySQL'
    $P10 = $P8.'doSthBin'($P9)
    .return($P10)
.end
.sub 'doSthBin' :method
    .param pmc x
    .lex 'x', x
    say x
    $P12 = box 0o75
    x = $P12
    $P13 = box 100
    $P14 = self.'swap:'($P13,x)
    .return($P12)
.end
.sub 'doSthBlock' :method
    .local pmc x
    .lex 'x', x
.end
.sub 'block_15' :outer('doSthBlock')
    .param pmc y
    .lex 'y', y
    .local pmc x
    find_lex x, 'x'
    $P16 = x + y
    .return($P16)
.end
.sub 'swap:' :method
    .param pmc a
    .lex 'a', a
    .param pmc b
    .lex 'b', b
    .local pmc tmp
    .lex 'tmp', tmp
    tmp = a
    a = b
    b = tmp
    .return(tmp)
.end
.namespace ['A']
.sub 'get_string' :method :vtable
    .return('A')
.end
.sub 'doSthKW:' :method
    .param pmc a
    .lex 'a', a
    .param pmc b
    .lex 'b', b
    $P18 = box 'str1'
    .return($P18)
.end
.namespace []
.sub 'main' :main
    $P0 = newclass 'B'
    $P17 = subclass 'B','A'
    .local pmc a
    .lex 'a', a
    .local pmc b
    .lex 'b', b
    .local pmc c
    .lex 'c', c
    $P19 = new 'B'
    a = $P19
    $P20 = a.'doSth'()
    say a
.end
.sub 'block_21' :outer('main')
    .param pmc x
    .lex 'x', x
    .local pmc y
    .lex 'y', y
    .local pmc a
    find_lex a, 'a'
    .local pmc b
    find_lex b, 'b'
    .local pmc c
    find_lex c, 'c'
    $P22 = x * y
    .return($P22)
.end
