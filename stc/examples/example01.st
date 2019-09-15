class Class1 [
    | a b c |
]

class Class2 [
    | a b c |
]

"
Multi
    line
    comment
"

class Class3 : Class2 []

class Class2 [
    method1 [ :arg1 :arg2 |
        | x y | x := y := 3. [ :a | | i j | r := $A. #s1.].
    ]

    kw_method1: [ :arg1 :arg2 |
        | x y | x + (y * 1) - (1 / 7) % x + y.
    ]

    method2 [
        2 doSth: 1 2 5; thenDoSthElse: #a; pass.
    ]
]

| a b c | a + b + c.
