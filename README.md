tlog Package provides a simple logging framework with support for colorful output, log file wrappers, and customizable loggers.
 
 This package includes a TLogger type, which provides methods for logging messages with various levels of severity, such as Println() and Panicln().
 Each TLogger can be customized with options such as whether to print to standard output, whether to use a log file wrapper, and the number of stack frames to skip when determining the caller's location.

 In addition, this package provides a Trace() function, which can be used to obtain the filename, line number, and function name of the caller at a specified number of stack frames away. This can be useful for adding more information to log messages.

 To use this package, create a new TLogger with CreateLogger() function and use the logging methods to output messages. You can also customize the TLogger options to fit your use case. By default, all TLoggers print to standard output and use log file wrappers.

 Example usage:

```go
     package main
     import (
         "github.com/tommynanny/tlog"
     )

     func main() {
        tlog.Println("Hello World!")
     }
```