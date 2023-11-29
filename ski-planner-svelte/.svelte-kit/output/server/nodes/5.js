

export const index = 5;
let component_cache;
export const component = async () => component_cache ??= (await import('../entries/pages/(public)/about/_page.svelte.js')).default;
export const imports = ["_app/immutable/nodes/5.6b0585e1.js","_app/immutable/chunks/scheduler.e108d1fd.js","_app/immutable/chunks/index.2b52f78d.js"];
export const stylesheets = [];
export const fonts = [];
