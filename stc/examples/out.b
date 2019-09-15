defcls Class1, Object
defcat Class1, a
defcmt Class1, doSth
defcma Class1, doSth, :x
defcma Class1, doSth, :y
defcmv Class1, doSth, a
addcmc Class1, doSth, pop, :x
addcmc Class1, doSth, pop, :y
defint %r1, 15
defflt %r2, 16
addcmc Class1, doSth, add, %r4, %r1, %r3
