function load({ cookies }) {
  const language = cookies.get("LANG") ?? "de";
  return { language };
}

var _layout_server_ts = /*#__PURE__*/Object.freeze({
  __proto__: null,
  load: load
});

const index = 2;
let component_cache;
const component = async () => component_cache ??= (await import('./_layout.svelte-2dd03c6b.js')).default;
const server_id = "src/routes/(public)/+layout.server.ts";
const imports = ["_app/immutable/nodes/2.5505d2be.js","_app/immutable/chunks/scheduler.e108d1fd.js","_app/immutable/chunks/index.2b52f78d.js","_app/immutable/chunks/app.7d0d17cc.js"];
const stylesheets = ["_app/immutable/assets/app.c5e7b66d.css"];
const fonts = [];

export { component, fonts, imports, index, _layout_server_ts as server, server_id, stylesheets };
//# sourceMappingURL=2-9d3a6ab3.js.map
