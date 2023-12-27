export default {
  trailingComma: "none",
  plugins: ["prettier-plugin-svelte"],
  overrides: [
    {
      files: "*.svelte",
      options: {
        parser: "svelte"
      }
    }
  ]
};
