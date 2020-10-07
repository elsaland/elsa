import { eq } from "./utils.ts";

Elsa.tests({
  "test mode == `test`": function () {
    eq(Elsa.mode, "test");
  },
});
