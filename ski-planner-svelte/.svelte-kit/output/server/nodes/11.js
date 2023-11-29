

export const index = 11;
let component_cache;
export const component = async () => component_cache ??= (await import('../entries/pages/(user)/home/_page.svelte.js')).default;
export const imports = ["_app/immutable/nodes/11.3aa413eb.js","_app/immutable/chunks/scheduler.e108d1fd.js","_app/immutable/chunks/index.2b52f78d.js"];
export const stylesheets = [];
export const fonts = [];
