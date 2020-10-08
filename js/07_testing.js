Elsa.tests = function (tests) {
  // Run tests only when mode is `test`
  // Will not run when running a script
  if (Elsa.mode !== "test") return;

  // Keep track for failures
  let failures = 0;
  // Loop through tests and run the function passed.
  // It basically `try...catch`'s the run and determines its failure.
  // No additional helper utilities are provided.
  let passed = 0;
  const startTestTime = Date.now();

  const red = "\u001b[31m";
  const gree = "\u001b[32m";
  const reset = "\u001b[0m";
  const gray = "\u001b[38;5;8m";
  const greenBG = "\u001b[42m";
  const redBG = "\u001b[41m";

  for (let testName in tests) {
    let testAction = tests[testName];
    const timeBeforeStart = Date.now();
    try {
      testAction();

      console.log(
        `TEST ${testName} ... ${greenBG} OK ${reset} ${gray}(${
          Date.now() - timeBeforeStart
        }ms)${reset}\n`
      );

      passed++;
    } catch (e) {
      failures++;
      console.error(
        `TEST ${testName} ... ${redBG} FAILED ${reset} ${gray}(${
          Date.now() - timeBeforeStart
        }ms)${reset} \n ${e}`
      );
      console.error(e.stack);
    }
  }
  const endTestTime = Date.now() - startTestTime;

  console.log(
    `TEST results: ${passed} passed, ${failures} failed, ${
      passed + failures
    } total, ${gray}(${endTestTime}ms)${reset}`
  );
};
