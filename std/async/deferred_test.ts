import { deferred } from "./deferred.ts";

Elsa.tests({
  "[async] deferred: resolve": async function () {
    const d = deferred<string>();
    d.resolve("❄");
    if ((await d) !== "❄") throw new Error("Assertion failed");
  },
  "[async] deferred: reject": async function (): Promise<void> {
    const d = deferred<number>();
    d.reject(new Error("An elsa error ❄"));
    d.then(() => {
      throw new Error("should fail");
    });
  },
});
