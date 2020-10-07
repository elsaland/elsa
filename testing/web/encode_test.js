const encoder = new TextEncoder();
const decoder = new TextDecoder();
import { eq } from "../tests/utils.ts";

const text = "Elsa";

Elsa.tests({
  "encode & decode": function () {
    const encoded = encoder.encode(text);
    const decoded = decoder.decode(encoded);
    eq(text, decoded);
  },
  "atob & btoa": function () {
    eq(btoa("Hello, world!"), "SGVsbG8sIHdvcmxkIQ==");
    eq(atob("SGVsbG8sIHdvcmxkIQ=="), "Hello, world!");
  },
});
