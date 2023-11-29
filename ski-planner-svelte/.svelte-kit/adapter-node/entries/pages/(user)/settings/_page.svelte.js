import { c as create_ssr_component, e as escape, v as validate_component } from "../../../../chunks/ssr.js";
import "js-cookie";
const Integration_card = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let { title } = $$props;
  if ($$props.title === void 0 && $$bindings.title && title !== void 0)
    $$bindings.title(title);
  return `<div class="card w-96 bg-base-100 shadow-xl"><div class="card-body"><h3 class="card-title">${escape(title)}</h3> ${slots.default ? slots.default({}) : ``}</div></div>`;
});
const Strava_integration = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let { user } = $$props;
  let isStravaConnected = user.stravaConnected;
  if ($$props.user === void 0 && $$bindings.user && user !== void 0)
    $$bindings.user(user);
  return `${validate_component(Integration_card, "IntegrationCard").$$render($$result, { title: "strava" }, {}, {
    default: () => {
      return `<p class="my-4" data-svelte-h="svelte-1996tdk">This application is powered by <a class="link" href="https://strava.com">Strava</a></p> ${isStravaConnected ? `<p data-svelte-h="svelte-1hyhngw">You are already connected.</p> <button class="btn btn-secondary" data-svelte-h="svelte-yb1imh">Click to disconnect</button>` : `<button data-svelte-h="svelte-8h23sc"><img src="branding/strava/btn-connect-with-strava.png"></button>`}`;
    }
  })}`;
});
const Page = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let { data } = $$props;
  if ($$props.data === void 0 && $$bindings.data && data !== void 0)
    $$bindings.data(data);
  return `<h1 class="text-2xl mb-6" data-svelte-h="svelte-5vn5pa">Settings</h1> <h2 class="text-xl mb-6" data-svelte-h="svelte-1fk1g4k">Partner Integrations</h2> ${validate_component(Strava_integration, "StravaIntegration").$$render($$result, { user: data.user }, {}, {})}`;
});
export {
  Page as default
};
