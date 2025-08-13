/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./views/**/*.html",
    "./views/partials/**/*.html"
  ],
  theme: {
    extend: {
      colors: {
        w: '#fefeff', //-- white
        f: '#f0f0f4', //-- off-white
        d: '#1a1b23', //-- black
        t: '#49433d', //-- text
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

