# big math

**_This package provides mathematical functions that are not provided by the math/big package._**

> ### How is computed the _logarithm_ of a number ?

It's not that difficult since we have the [`Log10` function provided by the math package](https://pkg.go.dev/math#Log10). So we just have to convert the big number into a `float64` (type input required by [`Log10`](https://pkg.go.dev/math#Log10)), and we're good.

But if an _overflow_ happens, it's bad!

So we check if the number is _too big_ to be converted into a `float64`.
If so, we know that:

![equation log(A*B) = log(A) + log(B)](./equations/equation-1.jpg)

So we can do:

![equation log(x) = log(sqrt(x)) + log(sqrt(x))](./equations/equation-2.jpg)

<sup>(let x be a big number)</sup><br>
By computing the square root of _x_ (with the [`Sqrt` function provided by the math/big package](https://pkg.go.dev/math/big)), that value will be smaller (![equation sqrt(x) < x](./equations/equation-3.jpg)) and may not cause an _overflow_.<br>
If it doesn't _overflow_, by _adding_ the _logarithm_ of this value with itself (![equation log(sqrt(x)) + log(sqrt(x))](./equations/equation-4.jpg)), we can compute the logarithm of the big number.<br>
Now if ![equation sqrt(x)](./equations/equation-5.jpg) still _overflows_, we just have to compute ![equation sqrt(sqrt(x))](./equations/equation-6.jpg). In other terms we get the _fourth root_. So to get ![equation log(x)](./equations/equation-7.jpg), we have to multiply ![equation sqrt(sqrt(x)) * 4](./equations/equation-8.jpg).<br>
And if ![equation sqrt(sqrt(x))](./equations/equation-9.jpg) is still too big, we continue the same process by computing its square root, and multiply it by _8_.
So the general formula is:

![equation log((2^n)√x) * 2^n = log(x)](./equations/equation-10.jpg)

<sup>(n ∈ N being the number of times we've computed the square root of the previous square root until we found a decent value that doesn't _overflow_. Meaning when ![equation 2^n√x](./equations/equation-11.jpg) <= `max number before overflow`)</sup>
