// TODO(@qu4k): remove when tla is supported
(async () => {
  try {
    const res = await fetch("Hello World")
    console.log(res)
  } catch (err) {
    console.error(err.message)
  }
})();

throw "ERROR";
