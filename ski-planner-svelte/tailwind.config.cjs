/** @type {import('tailwindcss').Config}*/
const config = {
    content: [
        './src/**/*.{html,js,svelte,ts}',
        './node_modules/flowbite-svelte/**/*.{html,js,svelte,ts}'
    ],
    plugins: [require('flowbite/plugin'), require('daisyui')],
    darkMode: 'class',
    theme: {
        extend: {
            gridTemplateColumns: {
                applayout: '15rem auto'
            },
            gridTemplateRows: {
                applayout: '2.5rem auto'
            }
        }
    },
    daisyui: {
        themes: [
            {
                skiyeti: {
                    primary: '#0C284A', // oxford blue
                    'primary-content': '#ffffff',
                    secondary: '#9BCBC4', // Tiffany Blue
                    accent: '#37cdbe',
                    neutral: '#566C7D',
                    'base-100': '#ffffff'
                }
            }
        ]
    }
};

module.exports = config;
