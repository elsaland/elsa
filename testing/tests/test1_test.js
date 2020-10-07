import { add, eq } from "./utils.ts";

Elsa.tests({
  "adds numbers #2": function () {
    eq(6, add(2, 4));
    eq(6.6, add(2.6, 4));
  },

  "subtracts numbers #2": function () {
    eq(-2, add(2, -4));
  },
});
