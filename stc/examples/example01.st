class B [
  doSth [
    | x y | x := y := 1. Console write x. y.
  ]

  doSth2 [
    1 + 2r10 * 16rFF. self doSth; doSthBin 'MySQL'.
  ]

  doSthBin [ :x |
    Console write x. x := 8r75. self swap: x 100.
  ]

  doSthBlock [
    | x | [ :y | x + y. ].
  ]

  swap: [ :a :b |
    | tmp | tmp := a. a := b. b := tmp.
  ]
]

class A : B [
  doSthKW: [ :a :b |
    'str1'.
  ]
]


| a b c |
a := B new. a doSth. [:x || y | x * y.].
Console write a.
