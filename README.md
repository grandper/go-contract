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

# Reference
- [Microsoft's code contract library](https://learn.microsoft.com/en-us/dotnet/framework/debug-trace-profile/code-contracts)
- Reusable Code Contract Helpers described in *Patterns, Principles, and Practices of Domain-Driven Design*, Millett, Tune, 2015 p.343
- [Terraform's validation](https://developer.hashicorp.com/terraform/language/expressions/custom-conditions)
