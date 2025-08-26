/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./views/**/*.html",
    "./views/partials/**/*.html"
  ],
  theme: {
    extend: {
      colors: {
        w: '#FEFEFF', //-- white
        f: '#EDEDF0', //-- off-white
        d: '#1A1B23', //-- black
        t: '#49433D', //-- text
        b: '#387CFF', //-- blue
      },
      padding: {
        x: '5vw',
        y: '3.5rem',
      },
    },
  },
  plugins: [],
}
