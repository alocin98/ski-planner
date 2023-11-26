import { UserService } from '@features/auth/service/user-service';

export async function load({ cookies, fetch }) {
	const language = cookies.get('LANG') ?? 'de';
	const user = await UserService.getUser(fetch)

	return { language, user };
}
