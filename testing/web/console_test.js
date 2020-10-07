Elsa.tests({
  "test console.log - string": function () {
    console.log("This is a string");
  },
  "test console.log - number": function () {
    console.log(12345);
  },
  "test console.log (number + string)": function () {
    console.log(1 + " is a string!");
  },
  "test console.log (nested object)": function () {
    console.log({
      super: {
        nested: {
          object: {
            it: {
              is: {
                indeed: true,
              },
            },
          },
        },
      },
    });
  },
  "test console.log - array with nested types": function () {
    console.log([
      1,
      2,
      3,
      "woop, a string",
      {
        oops: {
          coords: [
            {
              x: 1,
              y: 2,
            },
            {
              x: 2,
              y: 1,
            },
          ],
        },
      },
      ["umm", "nested", ["arrays"]],
    ]);
  },
  "test console.log - functions": function () {
    console.log(console.log);
  },
});
