<script lang="ts">
	import { Formaters } from '../utils/formaters';
	export let trainings: any[] = [];

	const { distance, time, stringIsoDate } = Formaters;
</script>

<table class="table">
	<!-- head -->
	<thead>
		<tr>
			<th>Name</th>
			<th>Date</th>
			<th>Distance</th>
			<th>Time</th>
			<th>sport</th>
		</tr>
	</thead>
	<tbody>
		{#each trainings as training}
			<tr>
				<th>{training.Name}</th>
				<td>{stringIsoDate(training.StartDate)}</td>
				<td>{distance(training.Distance)}</td>
				<td>{time(training.ElapsedTime)}</td>
				<td>{training.SportType}</td>
			</tr>
		{/each}
	</tbody>
</table>
