// Import the functions you need from the SDKs you need

import { initializeApp, type FirebaseApp } from 'firebase/app';
import { getAuth, browserLocalPersistence, type Auth, signInAnonymously } from 'firebase/auth';
import {
	PUBLIC_FIREBASE_API_KEY,
	PUBLIC_FIREBASE_APPID,
	PUBLIC_FIREBASE_AUTHDOMAIN,
	PUBLIC_FIREBASE_MEASUREMENTID,
	PUBLIC_FIREBASE_PROJECTID,
	PUBLIC_FIREBASE_STORAGEBUCKET,
	PUBLIC_FIREBASE_MESSAGINGSENDERID
} from '$env/static/public';
import Cookies from 'js-cookie';
// TODO: Add SDKs for Firebase products that you want to use
// https://firebase.google.com/docs/web/setup#available-libraries
// Your web app's Firebase configuration
// For Firebase JS SDK v7.20.0 and later, measurementId is optional
const firebaseConfig = {
	apiKey: PUBLIC_FIREBASE_API_KEY,
	authDomain: PUBLIC_FIREBASE_AUTHDOMAIN,
	projectId: PUBLIC_FIREBASE_PROJECTID,
	storageBucket: PUBLIC_FIREBASE_STORAGEBUCKET,
	messagingSenderId: PUBLIC_FIREBASE_MESSAGINGSENDERID,
	appId: PUBLIC_FIREBASE_APPID,
	measurementId: PUBLIC_FIREBASE_MEASUREMENTID
};

export let app: FirebaseApp;
export let auth: Auth;

export async function initFirebase() {
	app = initializeApp(firebaseConfig);
	auth = getAuth(app);
	await auth.setPersistence(browserLocalPersistence);
	if (auth.currentUser) {
		const token = await auth.currentUser.getIdTokenResult();
		Cookies.set('AccessToken', token.token, {
			expires: new Date(token.expirationTime),
			secure: true
		});
		return;
	}
	await signInAnonymously(auth);
}
// Initialize
