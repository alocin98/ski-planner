

export const index = 1;
let component_cache;
export const component = async () => component_cache ??= (await import('../entries/fallbacks/error.svelte.js')).default;
export const imports = ["_app/immutable/nodes/1.5cc2ac80.js","_app/immutable/chunks/scheduler.e108d1fd.js","_app/immutable/chunks/index.2b52f78d.js","_app/immutable/chunks/singletons.2f93ef9a.js"];
export const stylesheets = [];
export const fonts = [];
