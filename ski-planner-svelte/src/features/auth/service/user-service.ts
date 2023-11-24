import { createUserWithEmailAndPassword, signInWithEmailAndPassword } from 'firebase/auth';
import { auth } from '../../../providers/firebase/client';
import type { User } from '@types';
import Cookies from 'js-cookie';

async function loginWithEmailndPassword(email: string, password: string) {
	return signInWithEmailAndPassword(auth, email, password).then(async (user) => {
		const dbUser: Partial<User> = {
			email: user.user.email ?? 'NO-EMAIL',
		};
		const token = await user.user.getIdToken();
		Cookies.set('AccessToken', token);
        return backendLogin(dbUser as User).then(() => {
			return user;
		});
	});
}

async function backendLogin(user: User) {
    return fetch('/api/user/current', {
        method: 'POST',
        body: JSON.stringify(user)
    });
}

function newUserWithEmailAndPassword(email: string, password: string) {
	return createUserWithEmailAndPassword(auth, email, password).then(async (user) => {
		const dbUser: Partial<User> = {
			email: user.user.email ?? 'NO-EMAIL'
		};
		return fetch('/api/user/login', {
			method: 'POST',
			body: JSON.stringify(dbUser)
		}).then(() => user);
	});
}

function logout() {
	return auth.signOut();
}


function generateUserAndSendPasswordEmail(hotelUuid: string, email: string) {
	return fetch('/api/hotel/' + hotelUuid + '/users', {
		method: 'POST',
		body: JSON.stringify({ email})
	});
}

export const UserService = {
	loginWithEmailndPassword,
	newUserWithEmailAndPassword,
	logout,
	generateUserAndSendPasswordEmail
};
