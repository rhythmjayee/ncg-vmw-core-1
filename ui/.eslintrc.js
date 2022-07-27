module.exports = {
  root: true,
  env: {
    node: true,
  },
  extends: [
    'plugin:vue/essential',
    '@vue/airbnb',
  ],
  rules: {
    'no-console': process.env.NODE_ENV === 'production' ? 'error' : 'off',
    'no-debugger': process.env.NODE_ENV === 'production' ? 'error' : 'off',
    'vue/multi-word-component-names': 'off',
    'vue/no-reserved-component-names': 'off',
    'no-param-reassign': 0,
    'no-shadow': 0,
    'object-shorthand': 0,
    'max-len':'off',
  },
  parserOptions: {
    parser: '@babel/eslint-parser',
    ecmaVersion: 8,
    requireConfigFile: false,

  },
};
