/** @type {import('tailwindcss').Config} */
export default {
    content: ['./templates/**/*.html'],
    theme: {
        fontFamily: {
            sans: ['"Nunito Sans"', 'ui-sans-serif', 'sans-serif'],
            serif: ['"Red Rose"', 'ui-serif', 'serif'],
            mono: ['"Fragment Mono"', 'ui-monospace', 'monospace'],
        },
        extend: {
            colors: {
                signal: {
                    yellow: '#e6b019',
                    orange: '#bc602d',
                    red:    '#8f1e24',
                    blue:   '#134a85',
                    green:  '#417e57',
                    white:  '#ebecea',
                    black:  '#2f3133',
                },
            },
        },
    },
    plugins: [],
}
