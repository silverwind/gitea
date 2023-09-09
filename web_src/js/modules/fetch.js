// based on https://github.com/sindresorhus/p-retry
const networkErrorMsgs = new Set([
  'Failed to fetch', // Chrome
  'NetworkError when attempting to fetch resource.', // Firefox
  'The Internet connection appears to be offline.', // Safari
]);

export function isNetworkError(errorMessage) {
  return networkErrorMsgs.has(errorMessage);
}
