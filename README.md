# big math

**_This package provides mathematical functions that are not provided by the math/big package._**

## How is computed the _logarithm_ of a number ?

It's not that difficult since we have the [`Log10` function provided by the math package](https://pkg.go.dev/math#Log10). So we just have to convert the big number into a `float64` (type input required by [`Log10`](https://pkg.go.dev/math#Log10)), and we're good.

But if an _overflow_ happens, it's bad!

So we check if the number is _too big_ to be converted into a `float64`.
If so, we know that:

$log\big(a \times b\big) = log(a) + log(b)$

So we can do:

$log(x) = log(\sqrt{x}) + log(\sqrt{x})$&emsp;<sub>(let x be a big number)</sub>

By computing the square root of _x_ (with the [`Sqrt` function provided by the math/big package](https://pkg.go.dev/math/big)), that value will be smaller ($\sqrt{x} < x$) and _may_ not cause an _overflow_.<br>
If it doesn't _overflow_, by _adding_ the _logarithm_ of this value with itself $(log(\sqrt{x}) + log(\sqrt{x}))$, we can compute the _logarithm_ of the big number.<br>
Now if $\sqrt{x}$ still _overflows_, we just have to compute $\sqrt{\sqrt{x}}$. In other terms we get the _fourth root_. So to get $log(x)$, we have to multiply $\sqrt{\sqrt{x}} \times 4$.<br>
And if $\sqrt{\sqrt{x}}$ is still too big, we continue the same process by computing its square root, and multiply it by $8$.
So the general formula is:

$$log\big(\sqrt[2^{n}]{x}\big) \times 2^{n} = log(x)$$

<sup>(n ∈ N being the number of times we've computed the square root of the previous square root until we found a decent value that doesn't _overflow_. Meaning when $\sqrt[2^{n}]{x} \leq$ `max number before overflow`)</sup>

## Limitations of `Log10` (`IntLog10`, `FloatLog10`, `RatLog10`)

The largest number a `float64` (output type of `Log10`) can handle before _overflowing_ is $2^{1024} - 1$ (see [math.MaxFloat64](https://pkg.go.dev/math#MaxFloat64)).<br>
And we know that $log\big(10^{p}\big) = p$.<br>
With that, we can conclude that the **biggest** number we can compute its _logarithm_ is $10^{(2^{1024} - 1)}$.

Also, the variable that can _overflows_ before reaching $10^{(2^{1024} - 1)}$ is $2^{n}$ (which is of type `int64` and therefore can handle up to $2^{63} - 1$).<br>
We can know the **largest** number that can be computed before causing an _overflow_ to this value ($2^{n}$) by resolving this inequation:<br>
$\sqrt[2^{63} - 1]{x} \leq 2^{63} - 1$ for `IntLog10` (max value of `int64`: $2^{63} - 1$)<br>
$\sqrt[2^{63} - 1]{x} \leq 2^{1024} - 1$ for `FloatLog10` and `RatLog10` (max value of `float64`: $2^{1024} - 1$)

So for `IntLog10`: $x \leq 10^{\frac{log\big(2^{63} - 1\big)}{(\frac{1}{2^{63} - 1})}}$ or $x \leq \sqrt[\frac{1}{2^{63} - 1}]{2^{63} - 1}$

$x \leq 10^{10^{20.24284004863006}}$

And for `FloatLog10`/`RatLog10`: $x \leq \sqrt[\frac{1}{2^{63} - 1}]{2^{1024} - 1}$

$x \leq 10^{10^{21.45379945581629}}$
