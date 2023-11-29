

export const index = 9;
let component_cache;
export const component = async () => component_cache ??= (await import('../entries/pages/(user)/coolfeatures/_page.svelte.js')).default;
export const imports = ["_app/immutable/nodes/9.3e8ac265.js","_app/immutable/chunks/scheduler.e108d1fd.js","_app/immutable/chunks/index.2b52f78d.js"];
export const stylesheets = [];
export const fonts = [];
