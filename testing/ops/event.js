var ee = new EventEmitter();
ee.defineEvents(["01", "02", "03"]);

ee.addListener("01", (x) => {
  console.log(x);
});

ee.emitEvent("01", ["some emitted data"]);
