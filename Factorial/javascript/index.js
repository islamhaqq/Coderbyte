// Calculates the factorial of a number recursively using the: !n = n * !(n-1) algorithm.
const factorial = n => n > 1 ? n * factorial(n-1) : 1;

module.exports = factorial;
