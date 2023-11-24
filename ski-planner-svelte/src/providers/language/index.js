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
/**
 *
 * @param {string} key
 * @returns string
 */
export function t(key) {
	return currentLanguageFile[key];
}
