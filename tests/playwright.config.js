const { defineConfig } = require('@playwright/test');

module.exports = defineConfig({
  testDir: '.',
  timeout: 30000,
  use: {
    browserName: 'firefox',
    headless: true,
    baseURL: 'http://127.0.0.1:8080',
  },
});
