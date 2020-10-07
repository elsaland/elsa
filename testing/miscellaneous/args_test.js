Elsa.tests({
  "test Elsa.args": function () {
    if (!Elsa.args || typeof Elsa.args == "undefined") {
      throw new Error("Elsa.args is undefined");
    }
  },
});
