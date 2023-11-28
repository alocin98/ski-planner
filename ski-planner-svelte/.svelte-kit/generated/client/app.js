import * as client_hooks from '../../../src/hooks.client.ts';

export { matchers } from './matchers.js';

export const nodes = [
	() => import('./nodes/0'),
	() => import('./nodes/1'),
	() => import('./nodes/2'),
	() => import('./nodes/3'),
	() => import('./nodes/4'),
	() => import('./nodes/5'),
	() => import('./nodes/6'),
	() => import('./nodes/7'),
	() => import('./nodes/8'),
	() => import('./nodes/9'),
	() => import('./nodes/10'),
	() => import('./nodes/11'),
	() => import('./nodes/12'),
	() => import('./nodes/13')
];

export const server_loads = [2,3];

export const dictionary = {
		"/(public)": [4,[2]],
		"/(public)/about": [5,[2]],
		"/(public)/contact": [6,[2]],
		"/(user)/coolfeatures": [9,[3]],
		"/(user)/diary": [10,[3]],
		"/(user)/home": [11,[3]],
		"/(public)/login": [7,[2]],
		"/(public)/register": [8,[2]],
		"/(user)/settings": [12,[3]],
		"/(user)/stats": [13,[3]]
	};

export const hooks = {
	handleError: client_hooks.handleError || (({ error }) => { console.error(error) }),
};

export { default as root } from '../root.svelte';