const assert = require('assert');

const factorial = require('./');

assert.equal(factorial(5), 120);
assert.equal(factorial(1), 1);
assert.equal(factorial(15), 1307674368000)
