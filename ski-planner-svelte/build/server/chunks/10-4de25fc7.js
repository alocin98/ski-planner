function getTrainings(fetch) {
  if (!fetch) {
    fetch = window.fetch;
  }
  return fetch("/api/trainings").then((res) => res.json());
}
const TrainingService = {
  getTrainings
};
async function load({ fetch }) {
  const trainings = await TrainingService.getTrainings(fetch);
  return { trainings };
}

var _page_server_ts = /*#__PURE__*/Object.freeze({
  __proto__: null,
  load: load
});

const index = 10;
let component_cache;
const component = async () => component_cache ??= (await import('./_page.svelte-889f27c7.js')).default;
const server_id = "src/routes/(user)/diary/+page.server.ts";
const imports = ["_app/immutable/nodes/10.a3ba431a.js","_app/immutable/chunks/scheduler.e108d1fd.js","_app/immutable/chunks/index.2b52f78d.js","_app/immutable/chunks/strava.service.cffbef42.js","_app/immutable/chunks/navigation.e0f8f435.js","_app/immutable/chunks/singletons.2f93ef9a.js","_app/immutable/chunks/js.cookie.edb2da2a.js"];
const stylesheets = [];
const fonts = [];

export { component, fonts, imports, index, _page_server_ts as server, server_id, stylesheets };
//# sourceMappingURL=10-4de25fc7.js.map
