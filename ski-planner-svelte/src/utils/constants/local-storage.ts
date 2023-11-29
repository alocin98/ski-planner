export enum LocalStorageKey {
    welcomeStep = 'welcomeStep',
}

export let LocalStorageService = {
    set: (key: LocalStorageKey, value: string) => {
        localStorage.setItem(key, value);
    },
    get: (key: LocalStorageKey) => {
        return localStorage.getItem(key);
    },
    clear: () => {
        localStorage.clear();
    }
}
