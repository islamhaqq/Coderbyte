const assert = require('assert');

const firstReverse = require('./');

assert.equal(firstReverse('javascript'), 'tpircsavaj');
assert.equal(firstReverse('nodejs'), 'sjedon');
assert.equal(firstReverse('es6'), '6se');
