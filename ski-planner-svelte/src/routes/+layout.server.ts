import type { Cookies } from '@sveltejs/kit';

export function load({ cookies }: { cookies: Cookies }) {
	const language = cookies.get('LANG') ?? 'de';

	return { language };
}
