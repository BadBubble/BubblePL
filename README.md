# The Bubble Programming Language

## Overview
```
let five = 5;
let ten = 10;
let add = fn(x, y) {
	x + y;
};
let result = add(five, ten);
```

### Tutorial

### Variable
* int
```
let num = 5;
```
* string
```
let s = "hello,world";
```
* boolean
```
let f = false;
let t = true;
```
* array
```
let array = [5, -1, "haha", false];
let first = array[0];
```

* hash
```
let h = {1: "hi", "hello": "world", false: true};
let second = h["hello"];
```
### Functions
```
let foo = fn(x) {x * x};
let bar = fn(x, y, f) {f(x + y)};
let n = bar(1, 2, foo);
```
### Return
```
let foo = fn(x) {return x * x;};
```
### Condition
```
let x = 1;
if (x == 1) { let x = 2;} else {let x = 3;};
```
### Built-in Functions
* the length of string
```
let l = len("123");
```
* the length of array
```
let l = len([1, 2, 3]);
```

* the first object of array
```
let f = first([1, 2, 3]);
```

* the last object of array
```
let l = last([1, 2, 3]);
```
* delete the first object of array
```
let l = rest([1, 2, 3]);
```
* push an object to the tail of array
```
let l = [1, 2, 3];
push(l , 4);
```
* pop the last object of array
```
let l = [1, 2, 3];
pop(l);
```


* output
```
print("123");
```

## Features & TODOs

* [ ] bigint
* [ ] utf-8
* [ ] for
* [ ] for range
