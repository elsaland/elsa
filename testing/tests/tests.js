function add(a, b) {
  return a + b;
}

function eq(a, b) {
  if (a !== b) {
    throw new Error(`Assertion failed: ${a} !== ${b}`);
  }
}

Elsa.tests({
  "adds numbers": function () {
    eq(6, add(2, 4));
    eq(6.6, add(2.6, 4));
  },

  "subtracts numbers": function () {
    eq(-2, add(2, -4));
  },

  "add fail": function () {
    eq(2, add(2, 1));
  },
});
