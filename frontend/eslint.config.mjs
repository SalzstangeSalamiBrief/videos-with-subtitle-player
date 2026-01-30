//  @ts-check
import js from '@eslint/js';
import { tanstackConfig } from '@tanstack/eslint-config';
import eslintConfigPrettier from 'eslint-config-prettier/flat';
import globals from 'globals';
import tseslint from 'typescript-eslint';

export default [
  eslintConfigPrettier,
  ...tanstackConfig,
  { ignores: ['dist', 'api.ts', 'src/client/generated'] },
  {
    files: ['**/*.{ts,tsx,js}'],
    languageOptions: {
      ecmaVersion: 2020,
      globals: globals.browser,
    },
    rules: {
      '@typescript-eslint/no-explicit-any': 'error',
      'no-console': ['error', { allow: ['warn', 'error'] }],
      '@typescript-eslint/array-type': [
        'error',
        {
          default: 'array',
        },
      ],
      '@typescript-eslint/prefer-for-of': 'off',
      'import-x/order': 'off',
      'import/order': 'off',
    },
  },
  js.configs.recommended,
  ...tseslint.configs.recommended,
];
