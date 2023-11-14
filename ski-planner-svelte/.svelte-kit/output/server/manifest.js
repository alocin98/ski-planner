export const manifest = (() => {
function __memo(fn) {
	let value;
	return () => value ??= (value = fn());
}

return {
	appDir: "_app",
	appPath: "_app",
	assets: new Set(["favicon.png"]),
	mimeTypes: {".png":"image/png"},
	_: {
		client: {"start":"_app/immutable/entry/start.d4124f1e.js","app":"_app/immutable/entry/app.fc18729f.js","imports":["_app/immutable/entry/start.d4124f1e.js","_app/immutable/chunks/scheduler.e7d17915.js","_app/immutable/chunks/singletons.d9368c0f.js","_app/immutable/chunks/index.e21d44e1.js","_app/immutable/entry/app.fc18729f.js","_app/immutable/chunks/scheduler.e7d17915.js","_app/immutable/chunks/index.0b315888.js"],"stylesheets":[],"fonts":[]},
		nodes: [
			__memo(() => import('./nodes/0.js')),
			__memo(() => import('./nodes/1.js')),
			__memo(() => import('./nodes/2.js'))
		],
		routes: [
			{
				id: "/",
				pattern: /^\/$/,
				params: [],
				page: { layouts: [0,], errors: [1,], leaf: 2 },
				endpoint: null
			}
		],
		matchers: async () => {
			
			return {  };
		}
	}
}
})();
