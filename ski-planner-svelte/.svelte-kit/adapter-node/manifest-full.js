export const manifest = (() => {
function __memo(fn) {
	let value;
	return () => value ??= (value = fn());
}

return {
	appDir: "_app",
	appPath: "_app",
	assets: new Set(["branding/logo/favicon.ico","branding/logo/logo.jpeg","branding/strava/btn-connect-with-strava.png","branding/strava/powered-by-strava.png","images/hero.jpeg","powered by Strava/.DS_Store","powered by Strava/pwrdBy_strava_gray/api_logo_pwrdBy_strava_horiz_gray.eps","powered by Strava/pwrdBy_strava_gray/api_logo_pwrdBy_strava_horiz_gray.png","powered by Strava/pwrdBy_strava_gray/api_logo_pwrdBy_strava_horiz_gray.svg","powered by Strava/pwrdBy_strava_gray/api_logo_pwrdBy_strava_stack_gray.eps","powered by Strava/pwrdBy_strava_gray/api_logo_pwrdBy_strava_stack_gray.png","powered by Strava/pwrdBy_strava_gray/api_logo_pwrdBy_strava_stack_gray.svg","powered by Strava/pwrdBy_strava_light/api_logo_pwrdBy_strava_horiz_light.eps","powered by Strava/pwrdBy_strava_light/api_logo_pwrdBy_strava_horiz_light.png","powered by Strava/pwrdBy_strava_light/api_logo_pwrdBy_strava_horiz_light.svg","powered by Strava/pwrdBy_strava_light/api_logo_pwrdBy_strava_stack_light.eps","powered by Strava/pwrdBy_strava_light/api_logo_pwrdBy_strava_stack_light.png","powered by Strava/pwrdBy_strava_light/api_logo_pwrdBy_strava_stack_light.svg","powered by Strava/pwrdBy_strava_white/api_logo_pwrdBy_strava_horiz_white.eps","powered by Strava/pwrdBy_strava_white/api_logo_pwrdBy_strava_horiz_white.svg","powered by Strava/pwrdBy_strava_white/api_logo_pwrdBy_strava_stack_white.eps","powered by Strava/pwrdBy_strava_white/api_logo_pwrdBy_strava_stack_white.png","powered by Strava/pwrdBy_strava_white/api_logo_pwrdBy_strava_stack_white.svg"]),
	mimeTypes: {".jpeg":"image/jpeg",".png":"image/png",".eps":"application/postscript",".svg":"image/svg+xml"},
	_: {
		client: {"start":"_app/immutable/entry/start.1678d003.js","app":"_app/immutable/entry/app.bdc7e561.js","imports":["_app/immutable/entry/start.1678d003.js","_app/immutable/chunks/scheduler.e108d1fd.js","_app/immutable/chunks/singletons.2f93ef9a.js","_app/immutable/entry/app.bdc7e561.js","_app/immutable/chunks/client.5020f725.js","_app/immutable/chunks/js.cookie.edb2da2a.js","_app/immutable/chunks/scheduler.e108d1fd.js","_app/immutable/chunks/index.2b52f78d.js"],"stylesheets":[],"fonts":[]},
		nodes: [
			__memo(() => import('./nodes/0.js')),
			__memo(() => import('./nodes/1.js')),
			__memo(() => import('./nodes/2.js')),
			__memo(() => import('./nodes/3.js')),
			__memo(() => import('./nodes/4.js')),
			__memo(() => import('./nodes/5.js')),
			__memo(() => import('./nodes/6.js')),
			__memo(() => import('./nodes/7.js')),
			__memo(() => import('./nodes/8.js')),
			__memo(() => import('./nodes/9.js')),
			__memo(() => import('./nodes/10.js')),
			__memo(() => import('./nodes/11.js')),
			__memo(() => import('./nodes/12.js')),
			__memo(() => import('./nodes/13.js'))
		],
		routes: [
			{
				id: "/(public)",
				pattern: /^\/?$/,
				params: [],
				page: { layouts: [0,2,], errors: [1,,], leaf: 4 },
				endpoint: null
			},
			{
				id: "/(public)/about",
				pattern: /^\/about\/?$/,
				params: [],
				page: { layouts: [0,2,], errors: [1,,], leaf: 5 },
				endpoint: null
			},
			{
				id: "/(public)/contact",
				pattern: /^\/contact\/?$/,
				params: [],
				page: { layouts: [0,2,], errors: [1,,], leaf: 6 },
				endpoint: null
			},
			{
				id: "/(user)/coolfeatures",
				pattern: /^\/coolfeatures\/?$/,
				params: [],
				page: { layouts: [0,3,], errors: [1,,], leaf: 9 },
				endpoint: null
			},
			{
				id: "/(user)/diary",
				pattern: /^\/diary\/?$/,
				params: [],
				page: { layouts: [0,3,], errors: [1,,], leaf: 10 },
				endpoint: null
			},
			{
				id: "/(user)/home",
				pattern: /^\/home\/?$/,
				params: [],
				page: { layouts: [0,3,], errors: [1,,], leaf: 11 },
				endpoint: null
			},
			{
				id: "/(public)/login",
				pattern: /^\/login\/?$/,
				params: [],
				page: { layouts: [0,2,], errors: [1,,], leaf: 7 },
				endpoint: null
			},
			{
				id: "/(public)/register",
				pattern: /^\/register\/?$/,
				params: [],
				page: { layouts: [0,2,], errors: [1,,], leaf: 8 },
				endpoint: null
			},
			{
				id: "/(user)/settings",
				pattern: /^\/settings\/?$/,
				params: [],
				page: { layouts: [0,3,], errors: [1,,], leaf: 12 },
				endpoint: null
			},
			{
				id: "/(user)/stats",
				pattern: /^\/stats\/?$/,
				params: [],
				page: { layouts: [0,3,], errors: [1,,], leaf: 13 },
				endpoint: null
			}
		],
		matchers: async () => {
			
			return {  };
		}
	}
}
})();
