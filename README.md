# go-contract

**go-contract** is a small library to help you write code where all the expectations on inputs and outputs, and invariants are clearly stated.

You will write your contract by specifying
- Preconditions
- Postconditions
- Invariants

## Preconditions, Postconditions, and Invariants
Preconditions, postconditions and invariants are different ways to express requirements that must be met at some point in your code.

- Use **precondition** for assumptions. An assumption is a requirement that must be true at the beginning of a scope. For example, that a user must be logged in.

- Use **postcondition** for guarantees. A guarantee is characteristic or behavior of an object that must be true when exiting the scope. The rest of the code should be able to rely on it.

- Use **invariant** to make sure a condition or a relation remains true.

## How To

### Conditions
To write your expectations on the code, you need to either use *expressions* or *conditions".

There exists some functions in the library that let you directly write an *expression* that evaluate into a `bool`. For example if you want to check that a variable `x`of type `int` is positive, you will simply write `x > 0`.

Another possibility is to use a *condition*. The type `Condition` is simply defined as a function that accept a variable of a given type and that returns an error if the condition is not fulfilled.
```go
type Condition[T any] func(T) error
```

Although it might be a little bit more work than just writing an *expression*, a *condition* can help you encapsulate complex conditions into a more readable format. For example, you could create a condition `isValidEmail` that you can apply on a string containing an email. Using the libary will see in the code:
```go
err := contract.RequiresThat(email, isValidEmail).ErrorOnFailure()
```
*Conditions* can also be combined:
```go
err := contract.RequiresThat(product, isWellDesigned.And(isNotTooExpensive)).ErrorOnFailure()
```

### Imposing a Precondition
If you want to check requirement use the `Requires` function.
```go
contract.Requires(x != nil)
```
Then use special functions to define the behavior when the precondition fails:
```go
// Return an error.
err := contract.Requires(x != nil).ErrorOnFailure()

// Panic
contract.Requires(x != nil).PanicOnFailure()

// Log
contract.Requires(x != nil).LogOnFailure()
```

### Imposing a Postcondition
To define a precondition use the `Ensures` function:
```go
contract.Ensures(a.Field > 0)
```
Then use one of the function of the condition to define the behavior when the postcondition fails:
```go
// Return an error.
err := contract.Ensures(x != nil).ErrorOnFailure()

// Panic
contract.Ensures(x != nil).PanicOnFailure()

// Log
contract.Ensures(x != nil).LogOnFailure()
```
### Imposing an Invariant
To define a check on an invariant call `defer` on the `invariant` function. Then use one of the function of the invariant to define the behavior when the invariant fails:
```go
// Panic
defer contract.Invariant(number, isPositive).PanicOnFailure()

// Log
defer contract.Invariant(number, isPositive).LogOnFailure()
```
The library offers a specal call for checking for immutability. If you want to ensure that an object remains immutable use the `Immutable` function
```go
defer Immutable(&a).AssertVerified(t)
```
#### Note
Although you technically can check for immutability directly in your code, it's better to ensure the immutability of your code within your tests.

### Chaining
You might be wondering what is the benefit of this over just wrapping the email validation using a simple function. When you use function you get to actually chain all the requirement on the different variable and avoid a long list of `if` statements:
```go
// This is what a typical verification code looks like.
if !check1(variable) {
  return fmt.Errorf("check 1 failed")
}
if !check2(variable) {
  return fmt.Errorf("check 2 failed")
}
if !check3(variable) {
  return fmt.Errorf("check 2 failed")
}

// This is what you can write instead
if err := contract.RequiresThat(variable1,   
  hasProperty1.
  And(hasProperty2).
  And(hasProperty2)).
  ErrorOnFailure(); err != nil {
  return err
}

## Testing
Of course you may want to use preconditions, postconditions, and invariants in your tests. The library includes method to directly write expectation and asserting within your code.

```go
// Preconditions
contract.Requires(x != nil).AssertFulfilled(t)
contract.Requires(x != nil).AssertUnfulfilled(t)
contract.RequiresThat(number, isPositive).AssertFulfilled(t)
contract.RequiresThat(number, isPositive).AssertUnfulfilled(t)

// Postconditions
contract.Ensures(x != nil).AssertFulfilled(t)
contract.Ensures(x != nil).AssertUnfulfilled(t)
contract.EnsuresThat(number, isPositive).AssertFulfilled(t)
contract.EnsuresThat(number, isPositive).AssertUnfulfilled(t)

// Invariants
defer contract.Invariant(number, isPositive).AssertVerified(t)
defer contract.Invariant(number, isPositive).AssertViolated(t)
defer Immutable(&a).AssertVerified(t)
defer Immutable(&a).AssertViolated(t)
```

# References
- [Microsoft's code contract library](https://learn.microsoft.com/en-us/dotnet/framework/debug-trace-profile/code-contracts)
- Reusable Code Contract Helpers described in *Patterns, Principles, and Practices of Domain-Driven Design*, Millett, Tune, 2015 p.343
- [Terraform's validation](https://developer.hashicorp.com/terraform/language/expressions/custom-conditions)
