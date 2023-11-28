import { TrainingService } from '@features/training-diary/service/training-service';

export async function load({ fetch }) {
    const trainings = await TrainingService.getTrainings(fetch);

    return { trainings };
}
