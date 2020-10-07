Elsa.tests({
  "test fs - remove": function () {
    Elsa.writeFile(
      "testing/fs/to_remove.txt",
      `
    I am written by Elsa.writeFile
    ...and then removed by Elsa.writeFile
    `
    );

    Elsa.remove("testing/fs/to_remove.txt");
  },
});
