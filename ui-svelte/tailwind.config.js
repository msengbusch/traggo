/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./src/index.html",
    "./src/**/*.{svelte,js,ts}",
  ],
  plugins: [
      require('daisyui')
  ],
}

