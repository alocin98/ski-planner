import admin, { type ServiceAccount } from 'firebase-admin';

import serviceAccount from '@root/.google/sunloungery-prod-firebase-adminsdk-1k6f9-afbfbc5319.json';

export let app: admin.app.App;
export function initServerFirebase() {
	if (!app) {
		app = admin.initializeApp({
			credential: admin.credential.cert(serviceAccount as ServiceAccount)
		});
	}
}
