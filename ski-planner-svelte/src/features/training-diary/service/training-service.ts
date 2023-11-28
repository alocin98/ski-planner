function getTrainings(fetch?: typeof window.fetch) {
    if (!fetch) {
        fetch = window.fetch
    }
    return fetch('/api/trainings').then((res) => res.json())
}

export const TrainingService = {
    getTrainings
}
