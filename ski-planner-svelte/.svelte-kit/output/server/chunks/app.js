import { c as create_ssr_component } from "./ssr.js";
const SpaceLayout = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  return `<section class="bg-white dark:bg-gray-900 body-height overflow-y-scroll"><div class="max-w-screen-xl px-4 py-8 mx-auto">${slots.default ? slots.default({}) : ``}</div></section>`;
});
const app = "";
export {
  SpaceLayout as S
};
