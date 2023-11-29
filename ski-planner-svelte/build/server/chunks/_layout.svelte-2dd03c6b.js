import { c as create_ssr_component, v as validate_component } from './ssr-b1ec9c41.js';
import { S as SpaceLayout } from './app-53c25568.js';

const Header = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  return `<div class="navbar bg-base-100" data-svelte-h="svelte-1yqvytu"><div class="navbar-start"><div class="dropdown">  <label role="navigation" tabindex="0" class="btn btn-ghost md:hidden"><svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h8m-8 6h16"></path></svg></label> <ul class="menu menu-sm dropdown-content mt-3 z-[1] p-2 shadow bg-base-100 rounded-box w-52"><li><a href="about">about</a></li> <li><a href="contact">contact</a></li></ul></div> <div class="flex-1 flex flex-row"><img src="/branding/logo/logo.jpeg" class="mr-3 h-6 sm:h-9" alt="Skiyeti logo"> <a href="/" class="self-center whitespace-nowrap text-xl font-semibold dark:text-white">Skiyeti</a></div></div> <div class="navbar-end"><ul class="menu menu-horizontal px-1 hidden md:flex text-lg"><li><a href="about">about</a></li> <li><a href="contact">contact</a></li></ul> <a href="/login">Login</a></div></div>`;
});
const Layout = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  return `${validate_component(Header, "Header").$$render($$result, {}, {}, {})} ${validate_component(SpaceLayout, "SpaceLayout").$$render($$result, {}, {}, {
    default: () => {
      return `${slots.default ? slots.default({}) : ``}`;
    }
  })}`;
});

export { Layout as default };
//# sourceMappingURL=_layout.svelte-2dd03c6b.js.map
