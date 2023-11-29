function client_method(key) {
  {
    if (key === "before_navigate" || key === "after_navigate" || key === "on_navigate") {
      return () => {
      };
    } else {
      const name_lookup = {
        disable_scroll_handling: "disableScrollHandling",
        preload_data: "preloadData",
        preload_code: "preloadCode",
        invalidate_all: "invalidateAll"
      };
      return () => {
        throw new Error(`Cannot call ${name_lookup[key] ?? key}(...) on the server`);
      };
    }
  }
}
const goto = /* @__PURE__ */ client_method("goto");
const register = "register";
const firebaseErrors = {
  "auth/invalid-email": "Email konnte nicht erkannt werden"
};
const date = "Datum";
const price = "Preis";
const taste = "Geschmack";
const comment = "Kommentar";
const city = "Stadt";
const title = "Sun Loungery";
const subTitle = "Stop fighting for a spot by the pool. Just book it.";
const name = "name";
const optional = "optional";
const uniqueName = "unique name";
const streetAddress = "street address";
const postalCode = "postal code";
const country = "country";
const poolMap = "pool map";
const login = {
  loginButton: "login",
  password: "password",
  email: "e-mail",
  noAccountYet: "Don't have an account yet? better get one!",
  alreadyHaveAnAccount: "Already have an account?",
  loginLink: "login"
};
const en = {
  register,
  firebaseErrors,
  date,
  price,
  taste,
  comment,
  city,
  title,
  subTitle,
  name,
  optional,
  uniqueName,
  streetAddress,
  postalCode,
  country,
  poolMap,
  login
};
let currentLanguageFile = en;
function t(key) {
  const keys = key.split(".");
  let value = currentLanguageFile;
  for (const k of keys) {
    if (value && value.hasOwnProperty(k)) {
      value = value[k];
    } else {
      return void 0;
    }
  }
  return value;
}
export {
  goto as g,
  t
};
