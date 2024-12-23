# Shell Project

## Installation

*a*

## About the Project

This project implements a POSIX-compliant interactive and minimalist shell, designed to interpret shell commands, execute external programs, and provide support for various builtin commands. It also implements advanced features such as input/output redirection and support for single and double quote handling.

The goal is to create a lightweight and flexible experience while maintaining compatibility with [POSIX standards](https://www.gnu.org/software/bash/manual/html_node/Bash-Builtins.html), making the shell ideal for learning and experimentation.

## Builtin Commands

Below is the list of builtin commands implemented in the shell or planned for future implementation:

| Command                 | Description                                                                                | Usage Example                      | Status |
| ----------------------- | ------------------------------------------------------------------------------------------ | ---------------------------------- | ------ |
| `exit`                  | Terminates the shell execution with the specified status code.                             | `exit 0`                           | ✅      |
| `echo`                  | Prints the provided arguments to the terminal.                                             | `echo hello world`                 | ✅      |
| `type`                  | Displays how the shell interprets a command.                                               | `type echo`                        | ✅      |
| `pwd`                   | Prints the current working directory.                                                      | `pwd`                              | ✅      |
| `cd`                    | Changes the current working directory.                                                     | `cd /usr/local/bin`                | ✅      |
| Output Redirection (1>) | Redirects the output of a command to a file.                                               | `echo "Hello World" > file.txt`    | ✅      |
| Error Redirection (2>)  | Redirects error messages to a file.                                                        | `ls nonexistent 2> error.txt`      | ✅      |
| Append Output (1>>)     | Appends the output of a command to the end of a file.                                      | `echo "More text" >> file.txt`     | ✅      |
| Append Error (2>>)      | Appends error messages to the end of a file.                                               | `ls nonexistent 2>> error.txt`     | ✅      |
| `:`                     | A null command that does nothing but expand arguments and perform redirections.            | `: [arguments]`                    | 🛠️      |
| `.`                     | Reads and executes commands from a specified file in the current shell environment.        | `. filename [arguments]`           | 🛠️      |
| `eval`                  | Concatenates arguments into a single command, then executes it.                            | `eval [arguments]`                 | ❌      |
| `exec`                  | Replaces the shell with the specified command without creating a new process.              | `exec [-cl] [-a name] [command]`   | ❌      |

> [!NOTE]  
> Note: There are other commands considered builtins in Bash, as listed in the [Bash Builtins Manual](https://www.gnu.org/software/bash/manual/html_node/Bash-Builtins.html). However, since this is primarily a learning project, these are unlikely to be implemented.

## Advanced Features

- **Support for single and double quotes:** Preserves literal values or interprets special characters, respectively.
- **External program execution:** Locates and executes external programs using the PATH.

#### Examples:

```bash
$ echo "Text with 'single quotes' and \"double quotes\""
Text with 'single quotes' and "double quotes"
```



## Codecrafters Challenge

This shell project is part of the [Codecrafters Shell Course](https://app.codecrafters.io/courses/shell/). The course provides a guided journey to build your own shell from scratch, offering an in-depth understanding of shell functionalities and POSIX compliance.

If you're interested in taking the course, you can use my referral link to join: [https://app.codecrafters.io/r/witty-leopard-861910](https://app.codecrafters.io/r/witty-leopard-861910).

