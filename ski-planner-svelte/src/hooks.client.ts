//const analytics = getAnalytics(app);

import type { HandleClientError } from '@sveltejs/kit';
import { auth, initFirebase } from '@providers/firebase/client';
import type { User } from 'firebase/auth';
import Cookies from 'js-cookie';

initFirebase();

auth.onIdTokenChanged(async (user: User | null) => {
	if (!user) return;
	const token = await user?.getIdTokenResult();
	Cookies.set('AccessToken', token.token, {
		expires: new Date(token.expirationTime),
		secure: true
	});
});

export const handleError: HandleClientError = async ({ error, event }) => {
	const errorId = crypto.randomUUID();
	console.log(error, event, errorId);
};
