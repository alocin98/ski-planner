function getTrainings(fetch) {
  if (!fetch) {
    fetch = window.fetch;
  }
  return fetch("/api/trainings").then((res) => res.json());
}
const TrainingService = {
  getTrainings
};
async function load({ fetch }) {
  const trainings = await TrainingService.getTrainings(fetch);
  return { trainings };
}
export {
  load
};
