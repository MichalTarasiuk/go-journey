# Go Journey

## Terminology

- **Package**: A package in Go is a collection of source files located in the same directory, and they are compiled together as a unit.

- **Module**: A module in Go refers to a collection of related Go packages that are released together. Modules help manage dependencies and versioning of packages within a project.

## Go Naming Conventions

In Go, there are some common naming conventions that are widely followed to write clean and idiomatic code. These conventions help make your code more readable and maintainable. Here are some of the key naming conventions in Go:

### Package Names

- Package names should be short and concise, in lowercase letters. For multi-word package names, use lowercase letters with no underscores or hyphens (e.g., "math," "encoding/json," "net/http").

### Variable Names

- Variable names should be short and meaningful, using mixedCase (camelCase) for readability. For example, `userName`, `httpRequest`, `totalAmount`, etc.
- Avoid using single-letter variable names unless they have very short scopes (e.g., `i`, `j`, `x`, `y`, etc.).

### Function Names

- Function names should be concise and describe the action they perform. Use mixedCase (camelCase) for function names. For example, `getUserInfo()`, `calculateTotal()`, `openFile()`, etc.

### Constant Names

- Constants should be in uppercase letters with words separated by underscores. For example, `MAX_VALUE`, `PI`, `API_KEY`, etc.

### Type Names

- Type names should be short and meaningful, using mixedCase (camelCase). For example, `Person`, `httpRequest`, `TimeSpan`, etc.

### Exported Identifiers

- If an identifier (variable, function, type, or constant) starts with an uppercase letter, it is considered exported and can be accessed from other packages. If it starts with a lowercase letter, it is unexported (private) and can only be accessed within the same package.

### Acronyms

- Acronyms should be treated as single words. For example, use `httpRequest`, not `HTTPRequest`.

### Abbreviations

- Avoid using abbreviations unless they are widely accepted or part of your project's naming convention. Clarity and readability should be a priority.

### Receiver Names

- When defining methods on types, the receiver variable's name should be a short, meaningful abbreviation of the type name itself. Common abbreviations include `r`, `rec`, or the first letter of the type name itself.

Remember that consistency is key when following naming conventions. It's a good practice to adhere to these conventions in your Go code to make it more understandable to other developers and to ensure consistency across your project.
