/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./Templates//*.html"], // This is where your HTML templates / JSX files are located
  plugins: [require("@tailwindcss/forms")],
  theme: {
    extend: {
      fontFamily: {
        sans: ["Iosevka Aile Iaso", "sans-serif"],
        mono: ["Iosevka Curly Iaso", "monospace"],
        serif: ["Iosevka Etoile Iaso", "serif"],
      },
      backgroundImage: {
        "man-ship": "url('/Static/Images/a_space_ship_man_rocky.jpg')",
      },
    },
  },
  plugins: [],
};
