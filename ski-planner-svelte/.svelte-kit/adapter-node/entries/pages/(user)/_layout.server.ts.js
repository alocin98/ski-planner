import { signInWithEmailAndPassword, createUserWithEmailAndPassword } from "firebase/auth";
import Cookies from "js-cookie";
import "firebase/app";
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
export {
  load
};
