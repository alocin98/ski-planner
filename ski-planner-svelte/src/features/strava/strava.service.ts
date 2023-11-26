import { goto } from "$app/navigation";
import Cookies from "js-cookie";

const authorize = (redirect: string) => {
    Cookies.set('strava-authorization-redirect-url', redirect, {"expiresIn": "1000"});
    goto('/api/strava/authorize');

}

export const StravaService = {
    authorize
}