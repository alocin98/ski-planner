/** @type {import('tailwindcss').Config}*/
const config = {
	content: [
		'./src/**/*.{html,js,svelte,ts}',
		'./node_modules/flowbite-svelte/**/*.{html,js,svelte,ts}'
	],
	plugins: [require('flowbite/plugin'), require('daisyui')],
	darkMode: 'class',
	daisyui: {
		themes: [
			{
				skiyeti: {
					primary: '#6ED3F2',
					secondary: '#f6d860',
					accent: '#37cdbe',
					neutral: '#3d4451',
					'base-100': '#ffffff'
				}
			}
		]
	}
};

module.exports = config;
