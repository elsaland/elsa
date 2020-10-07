Elsa.tests({
  "test Elsa.mode": function () {
    if (!Elsa.mode || typeof Elsa.mode == "undefined") {
      throw new Error("Elsa.mode is undefined");
    }
  },
});
