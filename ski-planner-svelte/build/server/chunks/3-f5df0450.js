import { signInWithEmailAndPassword, createUserWithEmailAndPassword } from 'firebase/auth';
import Cookies from 'js-cookie';
import 'firebase/app';

let auth;
async function loginWithEmailAndPassword(email, password) {
  return signInWithEmailAndPassword(auth, email, password).then(async (user) => {
    const dbUser = {
      email: user.user.email ?? "NO-EMAIL"
    };
    const token = await user.user.getIdToken();
    Cookies.set("AccessToken", token);
    return backendLogin(dbUser).then(() => {
      return user;
    });
  });
}
async function backendLogin(user) {
  return fetch("/api/login", {
    credentials: "include",
    method: "POST",
    body: JSON.stringify(user)
  });
}
function newUserWithEmailAndPassword(email, password) {
  return createUserWithEmailAndPassword(auth, email, password).then(async (user) => {
    const dbUser = {
      email: user.user.email ?? "NO-EMAIL"
    };
    return fetch("/api/user/login", {
      method: "POST",
      body: JSON.stringify(dbUser)
    }).then(() => user);
  });
}
function logout() {
  Cookies.remove("AccessToken");
  return auth.signOut();
}
function generateUserAndSendPasswordEmail(hotelUuid, email) {
  return fetch("/api/hotel/" + hotelUuid + "/users", {
    method: "POST",
    body: JSON.stringify({ email })
  });
}
function getUser(fetch2) {
  if (!fetch2) {
    fetch2 = window.fetch;
  }
  return fetch2("/api/user", {
    credentials: "include"
  }).then(async (response) => {
    if (response.status === 401) {
      return null;
    }
    return response.json();
  }).catch(console.error);
}
const UserService = {
  loginWithEmailAndPassword,
  newUserWithEmailAndPassword,
  getUser,
  logout,
  generateUserAndSendPasswordEmail
};
async function load({ cookies, fetch: fetch2 }) {
  const language = cookies.get("LANG") ?? "de";
  const user = await UserService.getUser(fetch2);
  return { language, user };
}

var _layout_server_ts = /*#__PURE__*/Object.freeze({
  __proto__: null,
  load: load
});

const index = 3;
let component_cache;
const component = async () => component_cache ??= (await import('./_layout.svelte-37156280.js')).default;
const server_id = "src/routes/(user)/+layout.server.ts";
const imports = ["_app/immutable/nodes/3.b506d2b5.js","_app/immutable/chunks/scheduler.e108d1fd.js","_app/immutable/chunks/index.2b52f78d.js","_app/immutable/chunks/app.7d0d17cc.js","_app/immutable/chunks/navigation.e0f8f435.js","_app/immutable/chunks/singletons.2f93ef9a.js","_app/immutable/chunks/js.cookie.edb2da2a.js"];
const stylesheets = ["_app/immutable/assets/app.c5e7b66d.css"];
const fonts = [];

export { component, fonts, imports, index, _layout_server_ts as server, server_id, stylesheets };
//# sourceMappingURL=3-f5df0450.js.map
