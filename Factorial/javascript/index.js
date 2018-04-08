function factorial(n) {
  // We know !1 = 1.
  if (n === 1) return 1;
  // Because !5 = 5 * !4, we can say !n = n * !(n-1).
  return n * factorial(n-1)
}

module.exports = factorial;
