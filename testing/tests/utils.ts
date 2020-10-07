// Utility functions for testing the test runner ;)
// (also tests imports...for some reason it didn't feel good without them)

export function add(a: number, b: number) {
  return a + b;
}

export function eq(a: number | string | boolean, b: number | string | boolean) {
  if (a !== b) {
    throw new Error(`Assertion failed: ${a} !== ${b}`);
  }
}
