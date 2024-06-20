/** @type {import('tailwindcss').Config} */
const plugin = require("tailwindcss/plugin");
module.exports = {
  content: ["./Templates//*.html"], // This is where your HTML templates / JSX files are located
  plugins: [require("@tailwindcss/forms")],
  theme: {
    extend: {
      textShadow: {
        sm: "0 1px 2px var(--tw-shadow-color)",
        DEFAULT: "0 2px 4px var(--tw-shadow-color)",
        lg: "0 8px 16px var(--tw-shadow-color)",
      },
      fontFamily: {
        sans: ["Iosevka Aile Iaso", "sans-serif"],
        mono: ["Iosevka Curly Iaso", "monospace"],
        serif: ["Iosevka Etoile Iaso", "serif"],
      },
      boxShadow: {
        col: "0px 35px 60px 0px rgba(47, 72, 88 ,0.4)",
        Rcol: "20px 20px 20px 0px rgba(47, 72, 88 ,0.4)",
      },
      backgroundImage: {
        "man-ship": "url('/Static/Images/a_space_ship_man_rocky.jpg')",
        "city-map": "url('/Static/Images/a_map_of_a_city.jpg')",
      },
      keyframes: {
        size: {
          "30%, 70%": {
            transform: "scale(1.3)",
            background: "red",
          },
          "75%": {
            transform: "scale(1),",
            background: "#374151",
          },
        },
        wiggle: {
          "0%, 100%": { transform: "rotate(-3deg)" },
          "50%": { transform: "rotate(3deg)" },
        },
      },
      animation: {
        wiggle_anim: "spin 3s linear infinite",
      },
    },
  },
  plugins: [
    plugin(function ({ matchUtilities, theme }) {
      matchUtilities(
        {
          "text-shadow": (value) => ({
            textShadow: value,
          }),
        },
        { values: theme("textShadow") },
      );
    }),
  ],
};
