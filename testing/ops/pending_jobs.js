// TODO(@qu4k): remove when tla is supported
(async () => {
  const promise = new Promise((resolve) => {
    resolve(anotherHello());
  });
  console.log(await await promise);
})();

async function anotherHello() {
  const promise = new Promise((resolve) => {
    resolve("Hello World");
  });
  return promise;
}
