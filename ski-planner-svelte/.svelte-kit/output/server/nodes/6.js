

export const index = 6;
let component_cache;
export const component = async () => component_cache ??= (await import('../entries/pages/(public)/contact/_page.svelte.js')).default;
export const imports = ["_app/immutable/nodes/6.7050e596.js","_app/immutable/chunks/scheduler.e108d1fd.js","_app/immutable/chunks/index.2b52f78d.js"];
export const stylesheets = [];
export const fonts = [];
