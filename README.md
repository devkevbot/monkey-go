# The Monkey Programming Language

Monkey is an interpreted language written in Go based on [this](https://interpreterbook.com/) book by Thorsten Ball.

## Example program using the Monkey REPL

```
go run main.go
>> let name = "hello";
>> puts(name + " world");
hello world
```

## Features

### Data Types

| Data Type | Example |
| -------- | ------------------- |
| string | `"foo"`
| integer | `5` |
| boolean | `true`
| null | `null`
| array | `["foo", 5, null]`
| hash map | `{"name": "Alice", true: 5}`
| function | `fn foo() { return "bar"; };`


### Operators

| Operator | Description   | Supported Data Type |
| -------- | ------------- | ------------------- |
| `=`      | assignment    | All                 |
| `==`     | equals        | integer, boolean    |
| `!`      | not           | boolean             |
| `!=`     | not equals    | integer, boolean    |
| `<`      | less than     | integer             |
| `>`      | greater than  | integer             |
| `/`      | division      | integer             |
| `*`      | product       | integer             |
| `-`      | subtraction   | integer             |
| `+`      | addition      | integer             |
| `+`      | concatenation | string              |
| `[]`     | indexing      | arrays, hash maps   |

### Keywords

| Keyword | Description   | Usage |
| -------- | ------------- | ------------------- |
| `let` | binds a value to an identifier | `let hello = "world";` |
| `return` | returns from the current scope with a value | `return 5;` |
| `fn` | declare a function | `fn(a, b) { return a + b; };` |
| `if` | the main branch of a conditional statement | `if (len("foo") > 0) {...};` |
| `else` | the alternate branch of a conditional statement | `if (...) {...} else {...}` |
| `true` | the boolean value true | `true == true` |
| `false` | the boolean value false | `true == false` |

### Builtins

| Builtin | Description   | Usage |
| -------- | ------------- | ------------------- |
| `len(...)` | Returns the length of the type | `len("hello"); // 5` |
| `first(...)` | Return the first element of an array |  `first([5, 3]); // 5` |
| `last(...)` | Return the last element of an array |  `last([5, 3]); // 3` |
| `push(...)` | Return a **new** array with the value appended |  `push([5, 3], 4); // [5, 3, 4]` |
| `rest(...)` | Return all but the first element of an array |  `rest([5, 3]); // 3` |
| `puts(...)` | Prints the value to `stdout` | `puts("foo"); // foo `



## Structure

Monkey is interpreted using Go in the following sequence:

1. An input string of potential source code is read and tokenized using a lexer
2. The output of the lexer is given to a parser which creates an abstract syntax tree (AST) using Pratt parsing
methods
3. The parsed AST is recursively walked using a top-down approach and each node in the AST is evaluated as appropriate
4. The parsed expressions are mapped into the constructs
recognized by the Monkey programming language
5. An output is produced