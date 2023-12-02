/** @type {import('tailwindcss').Config} */
export default {
  content: [
      './src/**/*.{html,css,svelte}',
      './index.html',
  ],
  theme: {
    extend: {},
  },
  plugins: [
      require("daisyui")
  ],
}

