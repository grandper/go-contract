# go-contract

**go-contract** is a small library to help you write code where all the expectations on inputs and outputs, and invariants are clearly stated.

You will write your contract by specifying
- Preconditions
- Postconditions
- Invariants

## Preconditions, Postconditions, and Invariants
Preconditions, postconditions and invariants are different ways to express requirements that must be met at some point in your code.

- A **precondition** is a requirement that must be met at the beginning of a scope.

- A **postcondition** is arequirement that must be met when exiting the scope.

- An **invariant** is a condition or a relation that is always true.
