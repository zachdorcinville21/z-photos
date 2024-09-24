/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./templates/**/*.{html,js}"],
  theme: {
    extend: {
      colors: {
        primary: "#6a5acd",
      },
    },
    fontFamily: {
      titillium: ["Titillium Web"],
    },
  },
  plugins: [],
}

