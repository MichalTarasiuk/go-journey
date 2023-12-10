# Go Journey

## Terminology

- **Package**: A package in Go is a collection of source files located in the same directory, and they are compiled together as a unit.

- **Module**: A module in Go refers to a collection of related Go packages that are released together. Modules help manage dependencies and versioning of packages within a project.

## Commands

- **go verseion**: This command is used to print the Go version installed on your system. Running go version in the terminal will display the installed Go version.
- **go build**: The command is used to compile Go programs. It compiles the Go source code files in the current directory and produces an executable file. For example, running go build in a directory with Go source code will generate an executable with the same name as the directory.
- **go install**: This command is used to install the Go binary (executable) to the $GOPATH/bin directory. If the command is run in a directory containing Go source files, it compiles and installs the binary. If run without arguments, it installs the Go binaries for the packages specified.
- **go run**: The command is used to compile and execute Go programs directly from the source code.
- **go mod init**: The command is used to initialize a new Go module in the current directory. It creates a go.mod file, which serves as the module definition. The go.mod file contains information about the module, its dependencies, and other metadata.
- **go mod tidy**: The command is used to remove any dependencies from your go.mod file that are no longer required by your code. It ensures that your go.mod file reflects the actual dependencies used in your project.
- **go work init**: The command is used to initialize a new Go workspace in the current directory. It creates a go.work file. With multi-module workspaces, you can tell the Go command that you’re writing code in multiple modules at the same time and easily build and run code in those modules.
- **go work use**: The command which adds a module to the workspace.
- **go get**: The command is used to download and install packages and dependencies from the internet. It adds the specified packages to your go.mod file and downloads them.

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
