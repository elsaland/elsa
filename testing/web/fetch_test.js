Elsa.tests({
  "test fetch API": async function () {
    try {
      fetch("https://google.com");
    } catch (e) {
      throw new Error(e);
    }
  },
});
