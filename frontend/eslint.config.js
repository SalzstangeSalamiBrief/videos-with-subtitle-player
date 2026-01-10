//  @ts-check
import js from '@eslint/js';
import { tanstackConfig } from '@tanstack/eslint-config';
import eslintConfigPrettier from 'eslint-config-prettier/flat';
import globals from 'globals';
import tseslint from 'typescript-eslint';

export default [
  eslintConfigPrettier,
  ...tanstackConfig,
  { ignores: ['dist', 'api.ts'] },
  {
    files: ['**/*.{ts,tsx,js}'],
    languageOptions: {
      ecmaVersion: 2020,
      globals: globals.browser,
    },
    rules: {
      '@typescript-eslint/no-explicit-any': 'error',
      'no-console': ['error', { allow: ['warn', 'error'] }],
      '@typescript-eslint/array-type': 'generic',
    },
  },
  js.configs.recommended,
  ...tseslint.configs.recommended,
];
