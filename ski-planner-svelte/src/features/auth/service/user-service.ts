import { createUserWithEmailAndPassword, signInWithEmailAndPassword } from 'firebase/auth';
import Cookies from 'js-cookie';
import { auth } from '@providers/firebase/client';
import { LocalStorageService } from '@utils/constants/local-storage';
import type { User } from '@types';

async function loginWithEmailAndPassword(email: string, password: string) {
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
    return fetch('/api/login', {
        credentials: 'include',
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
    Cookies.remove('AccessToken');
    LocalStorageService.clear();
    return auth.signOut();
}


function generateUserAndSendPasswordEmail(hotelUuid: string, email: string) {
    return fetch('/api/hotel/' + hotelUuid + '/users', {
        method: 'POST',
        body: JSON.stringify({ email })
    });
}

function getUser(fetch?: typeof window.fetch) {
    if (!fetch) {
        fetch = window.fetch;
    }
    return fetch('/api/user', {
        credentials: 'include'
    }).then(async (response) => {
        if (response.status === 401) {
            return null;
        }

        return response.json();
    }).catch(console.error)
}

function finishTutorial() {
    return fetch('/api/user/finish-tutorial');
}

export const UserService = {
    loginWithEmailAndPassword,
    newUserWithEmailAndPassword,
    getUser,
    logout,
    generateUserAndSendPasswordEmail,
    finishTutorial
};
