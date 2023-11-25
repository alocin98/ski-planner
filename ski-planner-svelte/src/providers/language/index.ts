import en from './en.json';
import de from './de.json';

export const locale = 'de';

export let currentLanguageFile = en;

/**
 *
 * @param {string} lang
 */
export function changeLanguageFile(lang) {
	switch (lang) {
		case 'en':
			currentLanguageFile = en;
			break;
		case 'de':
			currentLanguageFile = de;
			break;
		default:
			currentLanguageFile = en;
	}
}

export function t(key: string) {
	const keys = key.split('.'); // Split the key by '.' to handle nested keys
	let value = currentLanguageFile;

	for (const k of keys) {
		if (value && value.hasOwnProperty(k)) {
			value = value[k];
		} else {
			return undefined; // Key not found, return undefined or handle as needed
		}
	}

	return value;
}
