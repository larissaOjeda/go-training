# Chapter 5 - Tests and Benchmarks

## Testing

### Test file location

By convention the test files are collocated with the actual code files.
We have the following convention for test file naming, given that the name of the file which contains the code to be tested is `repository.go`:

#### Unit tests

Unit tests reside in the `repository_test.go` file

#### Integration tests

Integration tests reside in the `integration_test.go` file. If there is need for more integration files per package we should prefix the file `repository_integration_test.go`.

There are 2 option in order to mark the integration tests:

##### Using a build tags

Using a build tag `// +build integration` at the beginning of a file. This way when you call `go test ./...` only the unit tests are build and run. If you want to run the integration tests also you have to include the build tag `go test ./... -tags=integration`.

- **[PRO]** Support for more test types like: end to end tests (e2e), contract testing etc
- **[PRO]** The CLI is cleaner since the default `go test ./...` will run only unit tests and you have to opt in for other test types
- **[CONS]** Since they are build tags the code is only build if you ask with the tags

##### Using `testing.Short()` to mark unit tests

Add to every test the `testing.Short()` at the beginning and use the `go test -short ./...` in order to run the unit tests.

- **[PRO]** No need to add a tag to include in the build process
- **[CONS]** Unit tests have to be tested using the `short` flag `go test -short ./...`
- **[CONS]** No support for other test types, there are only unit tests and the rest.

### CLI

In order to run tests we can use the following commands:

```bash
go test ./...
```

which will test all packages recursively from where the command runs.

In order to run the integration tests we just append the tag instruction to the above command:

```bash
go test ./... -tags=integration
```

Some helpful flags are the following:

- `-race`, which enables the race detector
- `-cover`, which reports the code coverage of each package
- `-timeout`, which sets the timeout e.g. `-timeout=60s`

Other flags can be found by running `go help test`.

### Stubs, mocks, spies

An excellent article on the differences is [Mocks Aren't Stubs](https://martinfowler.com/articles/mocksArentStubs.html) by Martin Fowler.

If we follow the small interface approach we can implement stubs by hand and don't need to use a package for that. Avoid adding dependencies.

If there is really need for mocks, there are some packages out there tha can help. Check out the below mentioned useful packages.

### Tests and sub-test

Any function that follows the below convention is a test. Keep in mind that in `Xxx` the first letter is capital.

```go
func TestXxx(*testing.T)
```

Assume we have the following code:

```go
func division(a, b float64) (float64, error) {
    if b == 0.0 {
        return 0.0, errors.New("division by zero")
    }
    return a / b, nil
}
```

A simple test with standard assertion:

```go
func TestDivision1_StdAssertion(t *testing.T) {
    res, err := division(3.0, 1.0)
    if err != nil {
        t.Errorf("division() returned an error %v where none was expected", err)
    }
    if res != 3.0 {
        t.Errorf("division() = %v, want 3.0", res)
    }
}
```

I would urge you to use a testing package ([testify](https://github.com/stretchr/testify)) which makes the above code smaller and easier to read.

```go
func TestDivision1(t *testing.T) {
    res, err := division(3.0, 1.0)
    assert.NoError(t, err)
    assert.Equal(t, 3.0, res)
}
```

A test with sub-tests:

```go
func TestDivision2(t *testing.T) {
    t.Run("success", func(t *testing.T) {
        res, err := division(3.0, 1.0)
        assert.NoError(t, err)
        assert.Equal(t, 3.0, res)
    })
    t.Run("failure", func(t *testing.T) {
        res, err := division(3.0, 0.0)
        assert.EqualError(t, err, "division by zero")
        assert.Equal(t, 0.0, res)
    })
}
```

A table driven test:

```go
func TestDivision3(t *testing.T) {
    type args struct {
        a float64
        b float64
    }
    tests := map[string]struct {
        args        args
        want        float64
        expectedErr string
    }{
        "success": {args: args{a: 3.0, b: 1.0}, want: 3.0},
        "failure": {args: args{a: 3.0, b: 0.0}, expectedErr: "division by zero"},
    }
    for name, tt := range tests {
        t.Run(name, func(t *testing.T) {
            got, err := division(tt.args.a, tt.args.b)
            if tt.expectedErr != "" {
                assert.EqualError(t, err, tt.expectedErr)
                assert.Equal(t, 0.0, got)
            } else {
                assert.NoError(t, err)
                assert.Equal(t, tt.want, got)
            }
        })
    }
}
```

The args struct can be omitted for simple cases but provides a good separation of what is an input and what an output.

In order to cover code paths there might be the need to check for specific errors in your tests. The code can handle this with the following 2 options:

- Check the error string `err.Error()` for an expected error. Checking the error message might lead to brittle tests especially if you are not in control of the error message
- Create custom error types and check for the specific error type

Hint: there is a nice tool, part of the Visual Studio Code extension and in available also Goland, that writes the whole skeleton test which can be invoked with `Ctrl+Shift+P` and select "Go: Generate Unit Test For Function".

**Prefer to write table-driven test since having one test for each test case will be hard to maintain.**

### Sample data

If for some reasons you need a fixture with data e.g. json you should create a folder inside the package folder named `testdata` and put the fixture in it. Then you can use the following code in your test to load that fixture:

```go
data, err := ioutil.ReadFile("testdata/abc.json")
```

### Useful packages

In order to make our life easier we can use some packages that help us to avoid writing repetitive code and make our tests more readable.

#### [testify](https://github.com/stretchr/testify)

So instead of writing:

```go
if age != 18 {
    t.Errorf("Age = %d; want 18", age)
}
```

we can write:

```go
assert.Equalf(t, 18, age, "Age = %d; want 18", age)
```

we can also use require which calls explicitly `t.FailNow()` on failure.

Testify contains also a mock package which can be used as a mocking framework.

## Code coverage and Visualization

If you are using VS Code or Goland you can actually see visually the code coverage and identify untested code paths.

For VS Code:

- Go to the top of the [test file](src/example_test.go)
- Click on the `run package tests` link over the package instruction
- Head to the code file and see the 100% coverage

## Benchmarks

Any function that follows the pattern:

```go
func BenchmarkXxx(*testing.B)
```

is considered a benchmark. There is no convention for the location of the benchmarks.
Usually they are inside the unit test files.

An example of a benchmark is the following:

```go
var res float64
var err error

func BenchmarkDivision(b *testing.B) {

    // Any initialization code comes here
    var res1 float64
    var err1 error

    b.ReportAllocs()
    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        res1, err1 = division(3.0, 1.0)
    }
    res = res1
    err = err1
}
```

Some ground rules here:

- In order to avoid compiler optimizations(inlining) the results of the division function are assigned to package level variable
- The `b.ReportAllocs()` enables allocation reporting which should be used by default
- The `b.ResetTimer()` can be used in order to start the timer after the reset. This way we can initialize any data we want without compromising the benchmark results
- The value of `b.N` will increase each time until the benchmark runner is satisfied with the stability of the benchmark.
- Sub and table driven benchmark are supported in the same way as the tests

In order to run the benchmark execute the following command:

```bash
go test -bench=Benchmark_division
```

The output of a run is the following:

```bash
goos: linux
goarch: amd64
pkg: examples
Benchmark_division-16           2000000000               0.74 ns/op            0 B/op          0 allocs/op
PASS
ok      examples        1.565s
```

where:

- `Benchmark_division-16` is the number of the benchmark along with a number that indicates how many cores the benchmark has used, in this case 16 cores
- 2000000000 is the number `b.N`

## Exercises

Create a new project and implement a method that calculates the [haversine distance](https://en.wikipedia.org/wiki/Haversine_formula) between two points Lat, Lng.
A point consists of a Longitude and Latitude.

### Write a table-driven test

Test the Distance of the following city pairs:

Athens - Amsterdam
Amsterdam - Berlin
Berlin - Athens

Where (Lat/Lng):

- Athens: 37.983972, 23.727806
- Amsterdam: 52.366667, 4.9
- Berlin: 52.516667, 13.388889

### Write a Benchmark

We should benchmark the performance of this function with Athens and Amsterdam as input.

[-> Next&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;: **Chapter 6**](../chapter6/README.md)  
[<- Previous&nbsp;: **Chapter 4**](../chapter4/README.md)
