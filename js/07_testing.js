Elsa.tests = function (tests) {
  // Run tests only when mode is `test`
  // Will not run when running a script
  if (Elsa.mode !== "test") return;

  // Keep track for failures
  let failures = 0;
  // Loop through tests and run the function passed.
  // It basically `try...catch`'s the run and determines its failure.
  // No additional helper utilities are provided.
  for (let testName in tests) {
    let testAction = tests[testName];
    try {
      testAction();
      console.log(`TEST ${testName} ... OK`);
    } catch (e) {
      // Increments faliure, prints stact trace
      failures++;
      console.error(`TEST ${testName} ... FAILED ${e}`);
      console.error(e.stack);
    }
  }
};
