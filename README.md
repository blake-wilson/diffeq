# diffeq
Diffeq is a REST server for solving differential equations.  A request for a solution is sent as JSON and should have
the following format:


```
{
    'timestep': NUMBER, // the timestep to use for the solver
    'final_time': NUMBER, // the last time value to evaluate
    'method': STRING // the method to use ('Euler' or 'Taylor')
    'expression': STRING // the derivative of the solution as a string expression
}
```

#### Format of an expression

Expressions can use the operators ```'+', '-', '*', '/', and '^'```.  They are entered as they would intuitively
be described.  For exmample, ```3*x^2``` is a valid expression (whose solution as a differential
equation x' = 3*x^2 is x^3).
