import prettierBase from "@videos-with-subtitle-player/prettier-base";

/**
 * @type {import("prettier").Config}
 */
const config = {
  ...prettierBase,
  plugins: ["prettier-plugin-tailwindcss"],
};

export default config;
