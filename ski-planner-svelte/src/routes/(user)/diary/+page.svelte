<script lang="ts">
	import NoEntries from '@features/training-diary/components/no-entries.svelte';
	import { TrainingService } from '@features/training-diary/service/training-service';

	export let data;
</script>

<h1 class="text-3xl mb-6">Diary</h1>
<div class="grid grid-cols-1 gap-6 md:grid-cols-2">
	{#each data.trainings as activity}
		<div class="bg-white rounded-lg shadow-md p-6" key={activity.trainingId}>
			<h2 class="text-xl font-semibold mb-2">{activity.Name}</h2>
			<p class="text-gray-600 mb-2">Distance: {activity.Distance} meters</p>
			<p class="text-gray-600 mb-2">Moving Time: {activity.MovingTime} seconds</p>
			<p class="text-gray-600 mb-2">Elevation Gain: {activity.TotalElevationGain} meters</p>
			<!-- Add other key details here -->
		</div>
	{/each}
</div>
<NoEntries />
