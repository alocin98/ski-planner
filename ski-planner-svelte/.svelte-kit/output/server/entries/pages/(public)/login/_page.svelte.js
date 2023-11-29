import { c as create_ssr_component, e as escape, d as add_attribute, v as validate_component } from "../../../../chunks/ssr.js";
import { t, g as goto } from "../../../../chunks/index.js";
import "firebase/auth";
import "js-cookie";
import "firebase/app";
const Login_form = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let { afterLogin } = $$props;
  let { registerUrl = "/register" } = $$props;
  const data = { email: "", password: "" };
  if ($$props.afterLogin === void 0 && $$bindings.afterLogin && afterLogin !== void 0)
    $$bindings.afterLogin(afterLogin);
  if ($$props.registerUrl === void 0 && $$bindings.registerUrl && registerUrl !== void 0)
    $$bindings.registerUrl(registerUrl);
  return `<div class="card w-96 bg-base-100 shadow-xl"><form class="card-body"><div class="form-control"><label for="email-input" class="label"><span class="label-text">${escape(t("login.email"))}</span></label> <input id="email-input" type="email" autocomplete="email" placeholder="email" class="input input-bordered"${add_attribute("value", data.email, 0)}></div> <div class="form-control"><label for="password-input" class="label"><span class="label-text">${escape(t("login.password"))}</span></label> <input id="password-input" type="password" placeholder="password" class="input input-bordered"${add_attribute("value", data.password, 0)}></div> ${``} <div class="form-control mt-2 gap-2"><button class="btn btn-primary">${`${escape(t("login.loginButton"))}`}</button> <p>${escape(t("login.noAccountYet"))}</p> <a class="link"${add_attribute("href", registerUrl, 0)}>${escape(t("register"))}</a></div></form></div>`;
});
const Page = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  const afterLogin = () => {
    goto("/home");
  };
  return `<div class="grid lg:gap-8 xl:gap-0 lg:py-16 lg:grid-cols-12"><div class="mx-auto place-self-center lg:col-span-7">${validate_component(Login_form, "LoginForm").$$render($$result, { afterLogin }, {}, {})}</div> <div class="hidden lg:mt-0 lg:col-span-5 lg:flex" data-svelte-h="svelte-115o7sw"><img src="images/hero.jpeg" alt="mockup"></div></div>`;
});
export {
  Page as default
};
