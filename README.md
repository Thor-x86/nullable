# Nullable SQL Data Types for Golang (GO)

[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://github.com/Thor-x86/nullable/blob/master/LICENSE)
[![Open Issues](https://img.shields.io/github/issues-raw/Thor-x86/nullable)](https://github.com/Thor-x86/nullable/issues)
[![Open Pull Request](https://img.shields.io/github/issues-pr-raw/Thor-x86/nullable)](https://github.com/Thor-x86/nullable/pulls)
[![Unit Test Result](https://img.shields.io/travis/Thor-x86/nullable)](https://travis-ci.org/Thor-x86/nullable)
[![Unit Test Coverage](https://img.shields.io/codecov/c/github/Thor-x86/nullable?token=41e528f35dbe410eb642eeb4c4d3a9e1)](https://codecov.io/github/Thor-x86/nullable)

Helps you convert every SQL [nullable](https://www.w3schools.com/sql/sql_null_values.asp) data types into Golang's supported types. So you don't have to create your own [scanner](https://www.geeksforgeeks.org/fmt-scan-function-in-golang-with-examples/) and [valuer](https://documentation.help/Golang/database_sql_drive.htm#Valuer) only for.. let's say... `BIGINT UNSIGNED NULL`

**Nullable** handles all SQL<->Go data type conversion burden just for you :)

## Features
- 100% [GORM](https://gorm.io/) support
- Can be marshalled into JSON
- Can be unmarshal from JSON
- Convenient Set/Get operation
- Support MySQL, MariaDB, SQLite, and PostgreSQL
- Zero configuration, just use it as normal data type.
- Heavily tested! So you don't have to worry of many bugs :D

## Supported Data Types
- bool
- byte
- string
- time.Time (capable of handling `DATETIME`, `TIME`, `DATE`, and `TIMESTAMP`)
- []byte
- float32
- float64
- int
- int8
- int16
- int32
- int64
- uint
- uint8
- uint16
- uint32
- uint64

**WARNING:** PostgreSQL [won't support any form of unsigned integers](https://www.postgresql.org/message-id/CAEcSYX+Arn7y4FeYPp6ZgbiiiMfZYmsn9aUyotZB-MA1n5hTOw@mail.gmail.com). However, you still able to use any uint variants with a drawback: **they will be stored in form of raw binary instead of normal integer**. Thus, PostgreSQL-side uint comparation is impossible, you have to compare them in your Go application. MySQL, MariaDB, and SQLite won't affected by this issue, don't worry.

# How to Use?

Very easy! first of all, let's install like normal Go packages

```bash
go get github.com/Thor-x86/nullable
```

## Create a new variable

Remember, all you need to have is a basic variable and a nullable variable created with `nullable.New...(&yourBasicVar)`. Example:

```go
import (
    "fmt"
    "gorm.io/gorm"
    "github.com/Thor-x86/nullable"
)

func main() {
    // Create new
    myBasicString := "Hello blяoW!"
    myNullableString := nullable.NewString(&myBasicString)

    // Create new but already nil
    myAlreadyNullString := nullable.NewString(nil)

    // Get and print to command console
    fmt.Println(myNullableString.Get()) // Output: Hello blяoW!
    fmt.Println(myAlreadyNullString.Get()) // Output: nil
}
```

## Change existing variable

You'll use `.Set(&anotherBasicVar)` to change existing variable. Example:

```go
import (
    "fmt"
    "gorm.io/gorm"
    "github.com/Thor-x86/nullable"
)

func main() {
    // Create new
    myBasicString := "Hello blяoW!"
    myNullableString := nullable.NewString(&myBasicString)
    fmt.Println(myNullableString.Get()) // Output: Hello blяoW!

    // Change existing variable
    anotherString := "Hello World!"
    myNullableString.Set(&anotherString)
    fmt.Println(myNullableString.Get()) // Output: Hello World!

    // Change with nil
    myNullableString.Set(nil)
    fmt.Println(myNullableString.Get()) // Output: nil
}
```

Also another example for uint64:

```go
import (
    "fmt"
    "gorm.io/gorm"
    "github.com/Thor-x86/nullable"
)

func main() {
    // Create new
    var theNumber uint64 = 70
    nullableNumber := nullable.NewUint64(&theNumber)
    fmt.Println(nullableNumber.Get()) // Output: 70

    // Change to another number
    var anotherNumber uint64 = 3306
    nullableNumber.Set(&anotherNumber)
    fmt.Println(nullableNumber.Get()) // Output: 3306

    // Change to nil
    nullableNumber.Set(nil)
    fmt.Println(nullableNumber.Get()) // Output: nil
}
```

## Change without creating a new basic variable

If you thinking it's not really convenient to create a basic variable first, then you can use `.Scan(...)` instead of `.Set(&yourBasicVar)`. Example:

```go
import (
    "fmt"
    "gorm.io/gorm"
    "github.com/Thor-x86/nullable"
)

func main() {
    // Create new
    nullableString := nullable.NewString(nil)

    // Directly write string
    nullableString.Scan("Hello blяoW!")
    fmt.Println(nullableString.Get()) // Output: Hello blяoW!

    // Change existing string
    nullableString.Scan("Hello World!")
    fmt.Println(nullableString.Get()) // Output: Hello World!

    // Scan also works with nil
    nullableString.Scan(nil)
    fmt.Println(nullableString.Get()) // Output: nil
}
```

**WARNING:** Mostly `.Scan(...)` won't cause compile-time error when you did something wrong, please be careful.

# For Contributors

Feel free to clone, fork, pull request, and open a new issue on this repository. However, you must test your work before asking for pull request. Here's how to execute the test:

1. For Windows users, [install WSL](https://docs.microsoft.com/en-us/windows/wsl/install-win10) first. MacOS and Linux users can skip to the next step.
2. Open your bash terminal. Then [install Docker](https://docs.docker.com/get-docker/) and [Docker Compose](https://docs.docker.com/compose/install/).
3. Make sure your bash terminal is currently pointing **nullable** project directory.
4. Run this to start the databases: ```docker-compose up```
5. Wait for a minute, just to make sure all databases are ready.
6. Run this to start the testing process: ```./test_all.sh```
7. If everything OK, you're good to pull request. Otherwise, check what's the root of problem.

# Frequently Asked Question (FAQ)

Q: If everything set with pointer, will it causes memory leak or something?
A: No, it won't store the pointer but the real value itself. Much like `sql.NullString` but every nil checking already done at **nullable** package.

Q: Why Scan is unsafe?
A: The `Scan` method actually meant to be implemented for Golang's SQL package to read query response from SQL server. However, you're allowed to use that as you like within risk of runtime error if you did something wrong.

Q: Will you add Microsoft SQL (MsSQL) Server support?
A: MsSQL throws me a lot of error while developing this project. I don't have much that time to tame MsSQL. Anyway, feel free to pull request. I already added the test configuration at `test_all.sh`, `main_test.go`, and `.github/workflows/tests.yml`. Just uncomment them and run the test as usual.

Q: How about key-value server support like Redis, Cassandra, and Memcached?
A: First of all, another ORM need to be implemented like [redis-orm](https://github.com/ezbuy/redis-orm). The challenge is we need to make sure it won't conflict with GORM. Pull request always welcome :)

Q: Is this support NoSQL like MongoDB?
A: If it JSON-marshallable, then the answer is yes. **Nullable** can be unmarshalled from JSON.

Q: I found security issue, can I open issue for this?
A: Please don't. You risking another users' security. Email me instead at athaariqa@gmail.com

Q: This FAQ section didn't answer my problem.
A: Feel free to [find issue](https://github.com/Thor-x86/nullable/issues) regarding the problem. If you found nothing, you are allowed to open a new issue.
