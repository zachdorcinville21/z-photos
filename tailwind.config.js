/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./templates/**/*.{html,js}"],
  theme: {
    extend: {
      colors: {
        primary: "#6a5acd",
        blue: '#5072A7'
      },
      gridTemplateColumns: {
        'auto-fit-xl': 'repeat(auto-fit, 600px)',
        'auto-fit-lg': 'repeat(auto-fit, 500px)',
        'auto-fit-md': 'repeat(auto-fit, 400px),'
      }
    },
    fontFamily: {
      titillium: ["Titillium Web"],
    },
  },
  plugins: [],
}

