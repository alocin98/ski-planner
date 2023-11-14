

export const index = 1;
let component_cache;
export const component = async () => component_cache ??= (await import('../entries/fallbacks/error.svelte.js')).default;
export const imports = ["_app/immutable/nodes/1.099b25a7.js","_app/immutable/chunks/scheduler.e7d17915.js","_app/immutable/chunks/index.0b315888.js","_app/immutable/chunks/singletons.e7e69fd3.js","_app/immutable/chunks/index.e21d44e1.js"];
export const stylesheets = [];
export const fonts = [];
