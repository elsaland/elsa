Elsa.tests = function (tests) {
  let failures = 0;
  for (let testName in tests) {
    let testAction = tests[testName];
    try {
      testAction();
      console.log(`TEST ${testName} ... OK`);
    } catch (e) {
      failures++;
      console.error(`TEST ${testName} ... FAILED ${e}`);
      console.error(e.stack);
    }
  }
};
