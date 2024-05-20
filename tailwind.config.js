/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["./**/*.html", "./**/*.templ", "./**/*.go",],
    darkMode: 'class',
    theme: {
        extend: {
            extend: {
                fontFamily: {
                    'exo': ['Exo 2', 'sans-serif'],
                },
            },
            colors: {
                primary: {
                    '50': '#edfcfe',
                    '100': '#d2f6fb',
                    '200': '#aaebf7',
                    '300': '#70daf0',
                    '400': '#2ec1e2',
                    '500': '#14b3da',
                    '600': '#1283a8',
                    '700': '#166a88',
                    '800': '#1b576f',
                    '900': '#1b495e',
                    '950': '#0c2e40',
                },
            }
        },
    },
    plugins: [],
}