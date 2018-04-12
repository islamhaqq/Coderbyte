const assert = require('assert');

const FirstReverse = require('./');

assert.equal(FirstReverse('javascript'), 'tpircsavaj');
assert.equal(FirstReverse('nodejs'), 'sjedon');
assert.equal(FirstReverse('es6'), '6se');
