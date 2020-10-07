/**
 * Copyright (c) elsa land 2020.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 *
 */

const _red = "\u001b[31m";
const _black = "\u001b[30m";
const _green = "\u001b[32m";
const _yellow = "\u001b[33m";
const _blue = "\u001b[34m";
const _magenta = "\u001b[35m";
const _cyan = "\u001b[36m";
const _white = "\u001b[37m";
const _reset = "\u001b[0m";

// * 16bits colors
const _brightBlack = "\u001b[30;1m";
const _brightRed = "\u001b[31;1m";
const _brightGreen = "\u001b[32;1m";
const _brightYellow = "\u001b[33;1m";
const _brightBlue = "\u001b[34;1m";
const _brightMagenta = "\u001b[35;1m";
const _brightCyan = "\u001b[36;1m";
const _brightWhite = "\u001b[37;1m";

// * backgroud color
const _backgroundBlack = "\u001b[40m";
const _backgroundRed = "\u001b[41m";
const _backgroundGreen = "\u001b[42m";
const _backgroundYellow = "\u001b[43m";
const _backgroundBlue = "\u001b[44m";
const _backgroundMagenta = "\u001b[45m";
const _backgroundCyan = "\u001b[46m";
const _backgroundWhite = "\u001b[47m";

// * backgroud Bright color
const _backgroundBrightBlack = "\u001b[40;1m";
const _backgroundBrightRed = "\u001b[41;1m";
const _backgroundBrightGreen = "\u001b[42;1m";
const _backgroundBrightYellow = "\u001b[43;1m";
const _backgroundBrightBlue = "\u001b[44;1m";
const _backgroundBrightMagenta = "\u001b[45;1m";
const _backgroundBrightCyan = "\u001b[46;1m";
const _backgroundBrightWhite = "\u001b[47;1m";

// * decorations
const _bold = "\u001b[1m";
const _underline = "\u001b[4m";
const _reversed = "\u001b[7m";

/**
 * 8bits red color
 * @param text
 */
export function red(text: string) {
  return `${_red}${text}${_reset}`;
}

/**
 * 8bits black color
 * @param text
 */
export function black(text: string) {
  return `${_black}${text}${_reset}`;
}

/**
 * 8bits green color
 * @param text
 */
export function green(text: string) {
  return `${_green}${text}${_reset}`;
}

/**
 * 8bits yellow color
 * @param text
 */
export function yellow(text: string) {
  return `${_yellow}${text}${_reset}`;
}

/**
 * 8bits blue color
 * @param text
 */
export function blue(text: string) {
  return `${_blue}${text}${_reset}`;
}

/**
 * 8bits magenta color
 * @param text
 */
export function magenta(text: string) {
  return `${_magenta}${text}${_reset}`;
}

/**
 * 8bits cyan color
 * @param text
 */
export function cyan(text: string) {
  return `${_cyan}${text}${_reset}`;
}

/**
 * 8bits white color
 * @param text
 */
export function white(text: string) {
  return `${_white}${text}${_reset}`;
}

/**
 * brightBlack color
 * @param text
 */
export function brightBlack(text: string) {
  return `${_brightBlack}${text}${_reset}`;
}

/**
 * brightRed color
 * @param text
 */
export function brightRed(text: string) {
  return `${_brightRed}${text}${_reset}`;
}

/**
 * brightGreen color
 * @param text
 */
export function brightGreen(text: string) {
  return `${_brightGreen}${text}${_reset}`;
}

/**
 * brightYellow color
 * @param text
 */
export function brightYellow(text: string) {
  return `${_brightYellow}${text}${_reset}`;
}

/**
 * brightBlue color
 * @param text
 */
export function brightBlue(text: string) {
  return `${_brightBlue}${text}${_reset}`;
}

/**
 * brightMagenta color
 * @param text
 */
export function brightMagenta(text: string) {
  return `${_brightMagenta}${text}${_reset}`;
}

/**
 * brightCyan color
 * @param text
 */
export function brightCyan(text: string) {
  return `${_brightCyan}${text}${_reset}`;
}

/**
 * brightWhite color
 * @param text
 */
export function brightWhite(text: string) {
  return `${_brightWhite}${text}${_reset}`;
}

/**
 * backgroundBlack color
 * @param text
 */
export function backgroundBlack(text: string) {
  return `${_backgroundBlack}${text}${_reset}`;
}

/**
 * backgroundRed color
 * @param text
 */
export function backgroundRed(text: string) {
  return `${_backgroundRed}${text}${_reset}`;
}

/**
 * backgroundGreen color
 * @param text
 */
export function backgroundGreen(text: string) {
  return `${_backgroundGreen}${text}${_reset}`;
}

/**
 * backgroundYellow color
 * @param text
 */
export function backgroundYellow(text: string) {
  return `${_backgroundYellow}${text}${_reset}`;
}

/**
 * backgroundBlue color
 * @param text
 */
export function backgroundBlue(text: string) {
  return `${_backgroundBlue}${text}${_reset}`;
}

/**
 * backgroundMagenta color
 * @param text
 */
export function backgroundMagenta(text: string) {
  return `${_backgroundMagenta}${text}${_reset}`;
}

/**
 * backgroundCyan color
 * @param text
 */
export function backgroundCyan(text: string) {
  return `${_backgroundCyan}${text}${_reset}`;
}

/**
 * backgroundWhite color
 * @param text
 */
export function backgroundWhite(text: string) {
  return `${_backgroundWhite}${text}${_reset}`;
}

/**
 * backgroundBrightBlack color
 * @param text
 */
export function backgroundBrightBlack(text: string) {
  return `${_backgroundBrightBlack}${text}${_reset}`;
}

/**
 * backgroundBrightRed color
 * @param text
 */
export function backgroundBrightRed(text: string) {
  return `${_backgroundBrightRed}${text}${_reset}`;
}

/**
 * backgroundBrightGreen color
 * @param text
 */
export function backgroundBrightGreen(text: string) {
  return `${_backgroundBrightGreen}${text}${_reset}`;
}

/**
 * backgroundBrightYellow color
 * @param text
 */
export function backgroundBrightYellow(text: string) {
  return `${_backgroundBrightYellow}${text}${_reset}`;
}

/**
 * backgroundBrightBlue color
 * @param text
 */
export function backgroundBrightBlue(text: string) {
  return `${_backgroundBrightBlue}${text}${_reset}`;
}

/**
 * backgroundBrightMagenta color
 * @param text
 */
export function backgroundBrightMagenta(text: string) {
  return `${_backgroundBrightMagenta}${text}${_reset}`;
}

/**
 * backgroundBrightCyan color
 * @param text
 */
export function backgroundBrightCyan(text: string) {
  return `${_backgroundBrightCyan}${text}${_reset}`;
}

/**
 * backgroundBrightWhite color
 * @param text
 */
export function backgroundBrightWhite(text: string) {
  return `${_backgroundBrightWhite}${text}${_reset}`;
}

/**
 * bold text
 * @param text
 */
export function bold(text: string) {
  return `${_bold}${text}${_reset}`;
}

/**
 * underline text
 * @param text
 */
export function underline(text: string) {
  return `${_underline}${text}${_reset}`;
}

/**
 * reverse color text
 * @param text
 */
export function reversed(text: string) {
  return `${_reversed}${text}${_reset}`;
}

/**
 * ANSI 256 color text
 * @param text
 */
export function color256(text: string, code: number) {
  if (code < 0 || code > 255) {
    console.log(
      red(
        bold(
          "only numbers between 0 and 255 can be passed in the color256 functio"
        )
      )
    );
    throw new Error();
  }

  return `\u001b[38;5;${code}m${text}${_reset}`;
}
