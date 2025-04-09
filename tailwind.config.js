/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
        "./cmd/web/**/*.html", "./cmd/web/**/*.templ",
    ],
    theme: {
        extend: {
            colors: {
                'bgColor': 'rgb(68,104,142)',
                'highlight': 'rgb(46,71,96)',
                'highlightHover': 'rgb(104,141,182)',
            }
        },
    },
    plugins: [],
}

