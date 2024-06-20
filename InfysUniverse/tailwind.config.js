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
      boxShadow: {
        col: "0px 35px 60px 0px rgba(47, 72, 88 ,0.4)",
        Rcol: "20px 20px 20px 0px rgba(47, 72, 88 ,0.4)",
      },
      backgroundImage: {
        "man-ship": "url('/Static/Images/a_space_ship_man_rocky.jpg')",
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
  plugins: [],
};
