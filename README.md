# big math

***This package provides mathematical functions that are not provided by the math/big package.***

> ### How is computed the _logarithm_ of a number ?

It's not that difficult since we have the `Log10` function provided by the math package. So we just have to convert the big number into a `float64` (type input required by `Log10`), and we're good.<br>
But if an _overflow_ happens, it's bad!<br>
So we check if the number is _too big_ to be converted into a `float64`.
If so, we know that:

![equation log(A*B) = log(A) + log(B)](https://bit.ly/3J0stAR)

So we can do:

![equation log(x) = log(sqrt(x)) + log(sqrt(x))](https://bit.ly/3PP5Oda)

<sup>(let x be a big number)</sup><br>
By computing the square root of _x_ (with the [`Sqrt` function provided by the math/big package](https://pkg.go.dev/math/big)), that value will be smaller (![equation sqrt(x) < x](https://bit.ly/3Ot8TOI)) and may not cause an _overflow_.<br>
If it doesn't _overflow_, by _adding_ the _logarithm_ of this value with itself (![equation log(sqrt(x)) + log(sqrt(x))](https://bit.ly/3J60r75)), we can compute the logarithm of the big number.
Now if ![equation sqrt(x)](https://bit.ly/3cAJ7uT) still _overflows_, we just have to compute ![equation sqrt(sqrt(x))](https://bit.ly/3aYI737). In other terms we get the _fourth root_. So to get ![equation log(x)](https://bit.ly/3v8oYmq), we have to multiply ![equation sqrt(sqrt(x)) * 4](https://bit.ly/3v95opS).
And if ![equation sqrt(sqrt(x))](https://bit.ly/3aYI737) is still too big, we continue the same process by computing its square root, and multiply it by _8_.
So the general formula is:

![equation log(2^n√x) \* 2^n = log(x)](https://bit.ly/3BbvL2h)

<sup>(n ∈ N being the number of times we've computed the square root of the previous square root until we found a decent value that doesn't _overflow_. Meaning when ![equation 2^n√x](https://bit.ly/3BjcJaa) <= `max number before overflow`)</sup>
