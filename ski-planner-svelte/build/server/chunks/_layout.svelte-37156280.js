import { c as create_ssr_component, v as validate_component } from './ssr-b1ec9c41.js';
import { S as SpaceLayout } from './app-53c25568.js';
import 'js-cookie';

const LoggedInHeader = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  return `<div class="navbar bg-base-100 h-10"><div class="navbar-start" data-svelte-h="svelte-uex9qc"><div class="dropdown">  <label for="my-drawer" tabindex="0" class="btn btn-ghost md:hidden"><svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h8m-8 6h16"></path></svg></label></div></div> <div class="navbar-end"><button data-svelte-h="svelte-hd9vso">Logout</button></div></div>`;
});
const Footer = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  return `<footer class="footer footer-center p-2 bg-base-300 text-base-content h-12" data-svelte-h="svelte-1lrtarc"><aside><p>Copyright © 2023 - skiyeti.io</p></aside></footer>`;
});
const Layout = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let { data } = $$props;
  if ($$props.data === void 0 && $$bindings.data && data !== void 0)
    $$bindings.data(data);
  return `<div class=""><div class="drawer lg:drawer-open"><input id="my-drawer" type="checkbox" class="drawer-toggle"> <div class="drawer-content lg:pl-60">${validate_component(LoggedInHeader, "LoggedInHeader").$$render($$result, {}, {}, {})} ${validate_component(SpaceLayout, "SpaceLayout").$$render($$result, {}, {}, {
    default: () => {
      return `${slots.default ? slots.default({}) : ``}`;
    }
  })} ${validate_component(Footer, "Footer").$$render($$result, {}, {}, {})}</div> <div class="drawer-side text-base-100" data-svelte-h="svelte-rlpg5h"><div class="p-4 min-h-full bg-primary fixed"><label for="my-drawer col-start-1" aria-label="close sidebar" class="drawer-overlay"></label> <div class="menu w-52"><label for="my-drawer" tabindex="0" class="btn btn-ghost lg:hidden"><svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h8m-8 6h16"></path></svg></label> <p class="text-xl pb-8">⛷️ Skiyeti</p> <ul> <li><a href="/diary">Diary</a></li> <li><a href="/stats">Stats</a></li> <li><a href="/coolfeatures">Some other really cool features</a></li> <li><a href="/settings">Settings</a></li></ul></div> <img class="absolute bottom-1 left-0 w-24" src="/branding/strava/powered-by-strava.png" alt="powered by strava"></div></div></div></div>`;
});

export { Layout as default };
//# sourceMappingURL=_layout.svelte-37156280.js.map
