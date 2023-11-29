import { c as create_ssr_component, d as add_attribute, v as validate_component, f as each, e as escape } from "../../../../chunks/ssr.js";
import "js-cookie";
const Strava_import_trainings = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let modal;
  return `<div class="card w-96 bg-base-100 shadow-xl"><div class="card-body"><h3 class="card-title" data-svelte-h="svelte-3owe1v">Strava</h3> <p class="my-4" data-svelte-h="svelte-191aaq0">This application is powered by <a class="link" href="https://strava.com">Strava</a></p> <button class="btn btn-primary" data-svelte-h="svelte-ehflca">Von Strava laden</button></div></div>  <dialog id="import-training-modal" class="modal"${add_attribute("this", modal, 0)}><div class="modal-box">${`<div class="modal-header" data-svelte-h="svelte-1mmesro"><h3>Importing training data from Strava</h3></div> <div class="modal-body" data-svelte-h="svelte-pmy5m0"><p>Importing training data from Strava can take a while. Please be patient.</p></div>`}</div></dialog>`;
});
const No_entries = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  return `<p data-svelte-h="svelte-pybrwt">Seems to be pretty empty bro. Either you lazy af or didn&#39;t import. You can do so now with one of
	your <a class="link" href="/settings">integrations</a></p> ${validate_component(Strava_import_trainings, "StravaImportTrainings").$$render($$result, {}, {}, {})}`;
});
const Page = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let { data } = $$props;
  if ($$props.data === void 0 && $$bindings.data && data !== void 0)
    $$bindings.data(data);
  return `<h1 class="text-3xl mb-6" data-svelte-h="svelte-o9wa5x">Diary</h1> <div class="grid grid-cols-1 gap-6 md:grid-cols-2">${each(data.trainings, (activity) => {
    return `<div class="bg-white rounded-lg shadow-md p-6"${add_attribute("key", activity.trainingId, 0)}><h2 class="text-xl font-semibold mb-2">${escape(activity.Name)}</h2> <p class="text-gray-600 mb-2">Distance: ${escape(activity.Distance)} meters</p> <p class="text-gray-600 mb-2">Moving Time: ${escape(activity.MovingTime)} seconds</p> <p class="text-gray-600 mb-2">Elevation Gain: ${escape(activity.TotalElevationGain)} meters</p>  </div>`;
  })}</div> ${validate_component(No_entries, "NoEntries").$$render($$result, {}, {}, {})}`;
});
export {
  Page as default
};
