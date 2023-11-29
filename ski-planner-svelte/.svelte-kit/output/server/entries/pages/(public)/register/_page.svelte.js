import { c as create_ssr_component, e as escape, d as add_attribute, v as validate_component } from "../../../../chunks/ssr.js";
import { t, g as goto } from "../../../../chunks/index.js";
import "firebase/auth";
import "js-cookie";
import "firebase/app";
const Register_form = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let { afterRegister } = $$props;
  let { loginUrl = "/login" } = $$props;
  const data = { email: "", password: "" };
  if ($$props.afterRegister === void 0 && $$bindings.afterRegister && afterRegister !== void 0)
    $$bindings.afterRegister(afterRegister);
  if ($$props.loginUrl === void 0 && $$bindings.loginUrl && loginUrl !== void 0)
    $$bindings.loginUrl(loginUrl);
  return `<div class="card w-96 bg-base-100 shadow-xl"><form class="card-body"><div class="form-control"><label for="email-input" class="label"><span class="label-text">${escape(t("login.email"))}</span></label> <input id="email-input" type="email" autocomplete="email" placeholder="email" class="input input-bordered"${add_attribute("value", data.email, 0)}></div> <div class="form-control"><label for="password-input" class="label"><span class="label-text">${escape(t("login.password"))}</span></label> <input id="password-input" type="password" placeholder="password" class="input input-bordered"${add_attribute("value", data.password, 0)}></div> ${``} <div class="form-control mt-2 gap-2"><button class="btn btn-secondary">${`${escape(t("register"))} `}</button> <p>${escape(t("login.alreadyHaveAnAccount"))}</p> <a${add_attribute("href", loginUrl, 0)} class="link">${escape(t("login.loginLink"))}</a></div></form></div>`;
});
const Page = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  const afterRegister = () => {
    goto("/home");
  };
  return `<div class="grid lg:gap-8 xl:gap-0 lg:py-16 lg:grid-cols-12"><div class="mx-auto place-self-center lg:col-span-7">${validate_component(Register_form, "RegisterForm").$$render($$result, { afterRegister }, {}, {})}</div> <div class="hidden lg:mt-0 lg:col-span-5 lg:flex" data-svelte-h="svelte-115o7sw"><img src="images/hero.jpeg" alt="mockup"></div></div>`;
});
export {
  Page as default
};
