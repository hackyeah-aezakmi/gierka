const { createGlobPatternsForDependencies } = require('@nx/angular/tailwind');
const { join } = require('path');

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    join(__dirname, 'src/**/!(*.stories|*.spec).{ts,html}'),
    join(__dirname, 'state/**/!(*.stories|*.spec).{ts,html}'),
    join(__dirname, 'landing/**/!(*.stories|*.spec).{ts,html}'),
    join(__dirname, 'join/**/!(*.stories|*.spec).{ts,html}'),
    join(__dirname, 'error/**/!(*.stories|*.spec).{ts,html}'),
    join(__dirname, 'characters/**/!(*.stories|*.spec).{ts,html}'),
    ...createGlobPatternsForDependencies(__dirname),
  ],
  theme: {
    extend: {},
  },
  plugins: [],
};
