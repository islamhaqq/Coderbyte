# Calculate the factorial of a number.
def Factorial(n):
    # We know !1 is 1.
    if n == 1:
        return 1
    # Because !5 is 5 * !4, we can say !n is n * !(n-1).
    return n * Factorial(n-1)
