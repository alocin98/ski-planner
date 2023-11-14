

export const index = 2;
let component_cache;
export const component = async () => component_cache ??= (await import('../entries/pages/_page.svelte.js')).default;
export const imports = ["_app/immutable/nodes/2.3b30f4b3.js","_app/immutable/chunks/scheduler.e7d17915.js","_app/immutable/chunks/index.0b315888.js","_app/immutable/chunks/ToolbarButton.b4dd5b23.js"];
export const stylesheets = [];
export const fonts = [];
