/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
        "./cmd/web/**/*.html", "./cmd/web/**/*.templ",
    ],
    theme: {
        extend: {
            colors: {
                'grayBG': 'rgb(110,107,122)',
                'highlight': 'rgb(70,67,82)',
                'highlightHover': 'rgb(98,92,122)',
            }
        },
    },
    plugins: [],
}

