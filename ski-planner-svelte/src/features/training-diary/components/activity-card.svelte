<script lang="ts">
	import { Formaters } from '../utils/formaters';
	export let activity;

	const { distance, time, elevation } = Formaters;
</script>

<div class="grid grid-cols-1 gap-6 md:grid-cols-2">
	<div class="bg-white rounded-lg shadow-md p-6" key={activity.trainingId}>
		<h2 class="text-xl font-semibold mb-2">{activity.Name}</h2>
		<p class="text-gray-600 mb-2">Distance: {distance(activity.Distance)}</p>
		<p class="text-gray-600 mb-2">Moving Time: {time(activity.MovingTime)}</p>
		<p class="text-gray-600 mb-2">
			Elevation Gain: {elevation(activity.TotalElevationGain)} meters
		</p>
		<!-- Add other key details here -->
	</div>
</div>
