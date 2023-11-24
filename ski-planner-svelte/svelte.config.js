import { vitePreprocess } from '@sveltejs/kit/vite';
import adapter from '@sveltejs/adapter-node';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	kit: {
		// adapter-auto only supports some environments, see https://kit.svelte.dev/docs/adapter-auto for a list.
		// If your environment is not supported or you settled on a specific environment, switch out the adapter.
		// See https://kit.svelte.dev/docs/adapters for more information about adapters.
		adapter: adapter(),
		alias: {
			'@assets/*': './src/assets/*',
			'@features/*': './src/features/*',
			'@packages/*': './src/packages/*',
			'@providers/*': './src/providers/*',
			'@components/*': './src/lib/components/*',
			'@utils/*': './src/utils/*',
			'@types': './src/types',
			'@root/*': './*'
		}
	},

	preprocess: [vitePreprocess({})]
};

export default config;
