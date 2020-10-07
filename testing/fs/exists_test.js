Elsa.tests({
  "test fs - exists (true)": function () {
    if (!Elsa.exists("testing/fs/sample.txt"))
      throw new Error("Elsa.exists failed");
  },
  "test fs - exists (false)": function () {
    if (Elsa.exists("testing/idonotexist.txt"))
      throw new Error("Elsa.exists failed");
  },
});
