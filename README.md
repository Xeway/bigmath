# big math

**_This package provides mathematical functions that are not provided by the math/big package._**

## How is computed the _logarithm_ of a number ?

It's not that difficult since we have the [`Log10` function provided by the math package](https://pkg.go.dev/math#Log10). So we just have to convert the big number into a `float64` (type input required by [`Log10`](https://pkg.go.dev/math#Log10)), and we're good.

But if an _overflow_ happens, it's bad!

So we check if the number is _too big_ to be converted into a `float64`.
If so, we know that:

![log(A*B) = log(A) + log(B)](./images/image-1.jpg)

So we can do:

![log(x) = log(sqrt(x)) + log(sqrt(x))](./images/image-2.jpg)

<sup>(let x be a big number)</sup><br>
By computing the square root of _x_ (with the [`Sqrt` function provided by the math/big package](https://pkg.go.dev/math/big)), that value will be smaller (![sqrt(x) < x](./images/image-3.jpg)) and may not cause an _overflow_.<br>
If it doesn't _overflow_, by _adding_ the _logarithm_ of this value with itself (![log(sqrt(x)) + log(sqrt(x))](./images/image-4.jpg)), we can compute the _logarithm_ of the big number.<br>
Now if ![sqrt(x)](./images/image-5.jpg) still _overflows_, we just have to compute ![sqrt(sqrt(x))](./images/image-6.jpg). In other terms we get the _fourth root_. So to get ![log(x)](./images/image-7.jpg), we have to multiply ![sqrt(sqrt(x)) * 4](./images/image-8.jpg).<br>
And if ![sqrt(sqrt(x))](./images/image-9.jpg) is still too big, we continue the same process by computing its square root, and multiply it by _8_.
So the general formula is:

![log((2^n)√x) * 2^n = log(x)](./images/image-10.jpg)

<sup>(n ∈ N being the number of times we've computed the square root of the previous square root until we found a decent value that doesn't _overflow_. Meaning when ![2^n√x](./images/image-11.jpg) <= `max number before overflow`)</sup>

## Limitations of `Log10` (`IntLog10`, `FloatLog10`, `RatLog10`)

The largest number a `float64` (output type of `Log10`) can handle before _overflowing_ is ![2^1024 - 1](./images/image-12.jpg) (see [math.MaxFloat64](https://pkg.go.dev/math#MaxFloat64)).<br>
And we know that ![log(10^p) = p](./images/image-13.jpg).<br>
With that, we can conclude that the **biggest** number we can compute its _logarithm_ is ![10^(2^1024 - 1)](./images/image-14.jpg).

Also, the variable that can _overflows_ before reaching ![10^(2^1024 - 1)](./images/image-15.jpg) is ![2^n](./images/image-16.jpg) (which is of type `int64` and therefore can handle up to ![2^63 - 1](./images/image-17.jpg)).<br>
We can know the **largest** number that can be computed before causing an _overflow_ to this value (![2^n](./images/image-18.jpg)) by resolving this inequation:<br>
![(2^63 -1)sqrt(x) <= 2^63 - 1](./images/image-19.jpg) for `IntLog10` (max value of `int64`: ![2^63 - 1](./images/image-20.jpg))<br>
![(2^63 -1)sqrt(x) <= 2^1024 - 1](./images/image-21.jpg) for `FloatLog10` and `RatLog10` (max value of `float64`: ![2^1024 - 1](./images/image-22.jpg))

So for `IntLog10`: ![x <= 10^(log(2^63 - 1)/(1/(2^63 - 1)))](./images/image-23.jpg) or ![x <= (1/(2^63 - 1))sqrt(2^63 - 1)](./images/image-24.jpg)

![x <= 10^(10^20.242840)](./images/image-25.jpg)

And for `FloatLog10`/`RatLog10`: ![x <= (1/(2^63 - 1))sqrt(2^1024 - 1)](./images/image-26.jpg)

![x <= 10^(10^21.453799)](./images/image-27.jpg)
