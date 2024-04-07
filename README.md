# expr

Package expr provides a type-safe way to build expressions.

Example:

```go
expr.Build[bool](func (user *User) expr.Bool {
  return expr.EqualRV(&user.Email, "user@example.com")
})
```

Note how pointer is used to reference a field, instead of name of it, which is not type-safe. And thanks to that,
the types of the field and value are guaranteed to match, by the use of generics.

Expressions can be evaluated and reflected. Evaluation can be used, for example, by a test database to filter rows.
Reflection can be used, for example, by ORM frameworks to generate SQL queries.

Check [documentation](https://pkg.go.dev/github.com/primego/expr#Expr) for details.
