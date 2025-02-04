# Go Programming Practical Tips

English &nbsp; | &nbsp; [Chinese](./README.zh-CN.md)

## Tips

### 1. Measure the execution time of a function in just one line of code

### 2. Multistage defer

### 3. Pre-allocate slices for performance

### 4. Parse an Array into a Slice

### 5. Method Chaining

### 6. Underscore Import

### 7. Wrapping Errors

### 8. Compile-Time Interface Verification

### 9. Avoid Naked Parameters

### 10. Umeric Separators

### 11. Avoid using math/rand, use crypto/rand for keys instead

### 12. Error messages should not be capitalized or end with punctuation

### 13. When to use Dot Import and Blank Import

### 14. Don't Return -1 or nil to Indicate Error

### 15. Understanding "Return fast, return early" to avoid nested code

### 16. Define interfaces in the consumer package, not the producer

### 17. Avoid named results unless necessary for documentation

### 18. Pass values, not pointers

### 19. Prefer using a pointer receiver when defining methods

### 20. Prefer using a pointer receiver when defining methods

### 21. Simplify function signatures with structs or variadic options

### 22. Skip the 'Get' prefix for getters

### 23. Avoid repetition in naming

### 24. Prefer 'chan struct{}' over 'chan bool' for signaling between goroutines

### 25. Explicitly ignore values with blank identifier (_) instead of silently ignoring them

### 26. Filter without any allocation

### 27. Converting multiple if-else statements into switch cases

### 28. Avoid context.Background(), make your goroutines promisable

### 29. Keep contexts going with context.WithoutCancel()

### 30. Loop labels for cleaner breaks and continues

### 31. Scheduling functions after context cancellation with context.AfterFunc

### 32. Just Don’t Panic

### 33. Lead with context, end with options, and always close with an error

### 34. Prefer strconv over fmt for converting to/from string

### 35. Naming Unexported Global Variables with an Underscore (_) Prefix

### 36. Using Unexported Empty Struct as Context Key

### 37. Make your errors clear with fmt.Errorf, don't just leave them bare

### 38. Avoid defer in loops, or your memory might blow up

### 39. Handle errors while using defer to prevent silent failures

### 40. Sort your fields in your struct from largest to smallest

### 41. Single Touch Error Handling, Less Noise

### 42. Gracefully Shut Down Your Application

### 43. Intentionally Stop with Must Functions

### 44. Always Manage Your Goroutine Lifetime

### 45. Avoid using break in switch cases, except when paired with labels

### 46. Table-driven tests, subtests, and parallel tests

### 47. Table-driven tests, subtests, and parallel tests

### 48. Avoid Global Variables, Especially Mutable Ones

### 49. Give the Caller the Right to Make Decisions

### 50. Make Structs Non-comparable
