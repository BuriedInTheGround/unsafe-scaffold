/** @type {import('tailwindcss').Config} */
export default {
    content: ['./templates/*.html'],
    theme: {
        fontFamily: {
            sans: [
                '"Work Sans"',
                'ui-sans-serif',
                'system-ui',
                '-apple-system',
                'BlinkMacSystemFont',
                '"Segoe UI"',
                'Roboto',
                '"Helvetica Neue"',
                'Arial',
                '"Noto Sans"',
                'sans-serif',
                '"Apple Color Emoji"',
                '"Segoe UI Emoji"',
                '"Segoe UI Symbol"',
                '"Noto Color Emoji"',
            ],
            serif: [
                '"Red Rose"',
                'ui-serif',
                'Georgia',
                'Cambria',
                '"Times New Roman"',
                'Times',
                'serif',
            ],
            mono: [
                '"Fragment Mono"',
                'ui-monospace',
                'SFMono-Regular',
                'Menlo',
                'Monaco',
                'Consolas',
                '"Liberation Mono"',
                '"Courier New"',
                'monospace',
            ],
        },
        extend: {
            colors: {
                luminous: {
                    yellow: '#ffff00',
                    orange: '#ff4612',
                    red:    '#ee1729',
                    green:  '#20a339',
                },
                pastel: {
                    yellow: '#d9a156',
                    orange: '#e17c30',
                    blue:   {
                        13:      '#8eb8de',
                        DEFAULT: '#6c8daa', // K=~33
                    },
                    green:  '#b8cfad',
                },
                signal: {
                    yellow: '#e6b019',
                    orange: '#bc602d',
                    red:    '#8f1e24',
                    blue: {
                        DEFAULT: '#134a85', // K=48
                        49:      '#124982',
                        59:      '#0f3b69',
                    },
                    green: {
                        36:      '#55a371',
                        DEFAULT: '#417e57', // K=~51
                    },
                    white: '#ebecea',
                    black: '#2f3133',
                },
            },
            fontSize: {
                '2xs': ['0.625rem', '0.75rem'],
            },
        },
    },
    plugins: [],
}
