import * as server from '../entries/pages/(user)/_layout.server.ts.js';

export const index = 3;
let component_cache;
export const component = async () => component_cache ??= (await import('../entries/pages/(user)/_layout.svelte.js')).default;
export { server };
export const server_id = "src/routes/(user)/+layout.server.ts";
export const imports = ["_app/immutable/nodes/3.b506d2b5.js","_app/immutable/chunks/scheduler.e108d1fd.js","_app/immutable/chunks/index.2b52f78d.js","_app/immutable/chunks/app.7d0d17cc.js","_app/immutable/chunks/navigation.e0f8f435.js","_app/immutable/chunks/singletons.2f93ef9a.js","_app/immutable/chunks/js.cookie.edb2da2a.js"];
export const stylesheets = ["_app/immutable/assets/app.c5e7b66d.css"];
export const fonts = [];
