// eslint-disable-next-line no-undef
module.exports = {
  extends: '../../.eslintrc.js',
  ignorePatterns: ['src/contract-artifacts.ts'],
  overrides: [
    {
      files: ['**/*.ts'],
      rules: {
        'no-constant-condition': ['error', { checkLoops: false }],
      },
    },
  ],
}
