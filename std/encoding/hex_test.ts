// Ported from Go
// https://github.com/golang/go/blob/go1.12.5/src/encoding/hex/hex.go
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Copyright 2018-2020 the Deno authors. All rights reserved. MIT license.
import { eq } from "../../testing/tests/utils.ts";

import {
  decode,
  decodedLen,
  decodeString,
  encode,
  encodedLen,
  encodeToString,
  errInvalidByte,
  errLength,
} from "./hex.ts";

function toByte(s: string): number {
  return new TextEncoder().encode(s)[0];
}

const testCases = [
  // encoded(hex) / decoded(Uint8Array)
  ["", []],
  ["0001020304050607", [0, 1, 2, 3, 4, 5, 6, 7]],
  ["08090a0b0c0d0e0f", [8, 9, 10, 11, 12, 13, 14, 15]],
  ["f0f1f2f3f4f5f6f7", [0xf0, 0xf1, 0xf2, 0xf3, 0xf4, 0xf5, 0xf6, 0xf7]],
  ["f8f9fafbfcfdfeff", [0xf8, 0xf9, 0xfa, 0xfb, 0xfc, 0xfd, 0xfe, 0xff]],
  ["67", Array.from(new TextEncoder().encode("g"))],
  ["e3a1", [0xe3, 0xa1]],
];

const errCases = [
  // encoded(hex) / error
  ["0", errLength()],
  ["zd4aa", errInvalidByte(toByte("z"))],
  ["d4aaz", errInvalidByte(toByte("z"))],
  ["30313", errLength()],
  ["0g", errInvalidByte(new TextEncoder().encode("g")[0])],
  ["00gg", errInvalidByte(new TextEncoder().encode("g")[0])],
  ["0\x01", errInvalidByte(new TextEncoder().encode("\x01")[0])],
  ["ffeed", errLength()],
];

Elsa.test({
  name: "[encoding.hex] encodedLen",
  fn(): void {
    eq(encodedLen(0), 0);
    eq(encodedLen(1), 2);
    eq(encodedLen(2), 4);
    eq(encodedLen(3), 6);
    eq(encodedLen(4), 8);
  },
});

Elsa.test({
  name: "[encoding.hex] encode",
  fn(): void {
    {
      const srcStr = "abc";
      const src = new TextEncoder().encode(srcStr);
      const dest = encode(src);
      eq(src, new Uint8Array([97, 98, 99]));
      eq(dest.length, 6);
    }

    for (const [enc, dec] of testCases) {
      const src = new Uint8Array(dec as number[]);
      const dest = encode(src);
      eq(dest.length, src.length * 2);
      eq(new TextDecoder().decode(dest), enc);
    }
  },
});

Elsa.test({
  name: "[encoding.hex] encodeToString",
  fn(): void {
    for (const [enc, dec] of testCases) {
      eq(encodeToString(new Uint8Array(dec as number[])), enc);
    }
  },
});

Elsa.test({
  name: "[encoding.hex] decodedLen",
  fn(): void {
    eq(decodedLen(0), 0);
    eq(decodedLen(2), 1);
    eq(decodedLen(4), 2);
    eq(decodedLen(6), 3);
    eq(decodedLen(8), 4);
  },
});

Elsa.test({
  name: "[encoding.hex] decode",
  fn(): void {
    // Case for decoding uppercase hex characters, since
    // Encode always uses lowercase.
    const extraTestcase = [
      ["F8F9FAFBFCFDFEFF", [0xf8, 0xf9, 0xfa, 0xfb, 0xfc, 0xfd, 0xfe, 0xff]],
    ];

    const cases = testCases.concat(extraTestcase);

    for (const [enc, dec] of cases) {
      const src = new TextEncoder().encode(enc as string);
      const dest = decode(src);
      eq(Array.from(dest), Array.from(dec as number[]));
    }
  },
});

Elsa.test({
  name: "[encoding.hex] decodeString",
  fn(): void {
    for (const [enc, dec] of testCases) {
      const dst = decodeString(enc as string);

      eq(dec, Array.from(dst));
    }
  },
});

Elsa.test({
  name: "[encoding.hex] decode error",
  fn(): void {
    for (const [input, expectedErr] of errCases) {
      assertThrows(
        () => decode(new TextEncoder().encode(input as string)),
        Error,
        (expectedErr as Error).message
      );
    }
  },
});

Elsa.test({
  name: "[encoding.hex] decodeString error",
  fn(): void {
    for (const [input, expectedErr] of errCases) {
      assertThrows(
        (): void => {
          decodeString(input as string);
        },
        Error,
        (expectedErr as Error).message
      );
    }
  },
});
