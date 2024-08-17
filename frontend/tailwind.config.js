/** @type {import('tailwindcss').Config} */
const textWidth = '60ch';

export default {
  content: ['./index.html', './src/**/*.{js,ts,jsx,tsx}'],
  theme: {
    extend: {
      transitionProperty: {
        rotate: 'rotate',
      },
      width: {
        text: textWidth,
      },
      maxWidth: {
        text: `min(100%, ${textWidth})`,
      },
    },
  },
  plugins: [],
};
