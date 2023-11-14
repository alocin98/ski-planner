/** @type {import('tailwindcss').Config}*/
const config = {
	content: ['./src/**/*.{html,js,svelte,ts}', './node_modules/flowbite-svelte/**/*.{html,js,svelte,ts}'],
	plugins: [require('flowbite/plugin')],
	darkMode: 'class',

	theme: {
		extend: {
			colors: {
				// flowbite-svelte
				primary: {
					50: '#6ED3F2',     // Light Icy Blue
					100: '#6EC9F2',    // Slightly lighter
					200: '#6EBBF2',    // Lighter Icy Blue
					300: '#6EAAF2',    // A bit more saturated
					400: '#6E94F2',    // Even more saturated
					500: '#6E76F2',    // Primary Icy Blue
					600: '#5E5CCB',    // Evergreen Green for deeper shades
					700: '#3F4073',    // Darker tone for contrast
					800: '#3B385A',    // Dark Frosty Gray
					900: '#292741'     // Deepest Frosty Gray
				  }
				  
			  }
		}
	},

	plugins: []
};

module.exports = config;
