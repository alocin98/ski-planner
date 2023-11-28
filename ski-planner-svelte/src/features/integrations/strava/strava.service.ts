import { goto } from "$app/navigation";
import Cookies from "js-cookie";

const authorize = (redirect: string) => {
    Cookies.set('strava-authorization-redirect-url', redirect, {"expiresIn": "1000"});
    goto('/api/strava/authorize');

}

const loadTrainingData = () => {
    return fetch('/api/strava/load-training-data')
        .then(response => response.json())
        .then(data => {
            return data;
        })
}

export const StravaService = {
    authorize,
    loadTrainingData
}