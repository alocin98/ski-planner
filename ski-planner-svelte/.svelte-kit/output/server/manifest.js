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
		client: {"start":"_app/immutable/entry/start.592d1404.js","app":"_app/immutable/entry/app.c5f313b5.js","imports":["_app/immutable/entry/start.592d1404.js","_app/immutable/chunks/scheduler.e7d17915.js","_app/immutable/chunks/singletons.e7e69fd3.js","_app/immutable/chunks/index.e21d44e1.js","_app/immutable/entry/app.c5f313b5.js","_app/immutable/chunks/scheduler.e7d17915.js","_app/immutable/chunks/index.0b315888.js"],"stylesheets":[],"fonts":[]},
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
