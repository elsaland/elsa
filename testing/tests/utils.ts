// Utility functions for testing the test runner ;)
// (also tests imports...for some reason it didn't feel good without them)

export function add(a: number, b: number) {
  return a + b;
}

export function eq(a: number | string | boolean, b: number | string | boolean) {
  if (!deepEqual(a, b)) {
    throw new Error(`Assertion failed: ${a} !== ${b}\n`);
  }
}

export function deepEqual(a: any, b: any) {
  if (Object.is(a, b)) {
    // items are identical
    return true;
  } else if (
    typeof a === "object" &&
    a !== null &&
    typeof b === "object" &&
    b !== null
  ) {
    // items are objects - do a deep property value compare
    // join keys from both objects together in one array
    let keys = Object.keys(a).concat(Object.keys(b));
    // filter out duplicate keys
    keys = keys.filter(function (value, index, self) {
      return self.indexOf(value) === index;
    });
    for (const p of keys) {
      if (typeof a[p] === "object" && typeof b[p] === "object") {
        if (!deepEqual(a[p], b[p])) {
          return false;
        }
      } else if (a[p] !== b[p]) {
        return false;
      }
    }
    return true;
  } else {
    return false;
  }
}
