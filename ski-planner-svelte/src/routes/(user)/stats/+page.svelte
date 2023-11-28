<script>
	import { StravaService } from '@features/integrations/strava/strava.service';

	let loadTrainingData = () => {
		StravaService.loadTrainingData().then(console.log);
	};
</script>

<h1>Planner</h1>
<button on:click={loadTrainingData}>Load training data</button>
