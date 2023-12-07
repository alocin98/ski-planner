<script lang="ts">
	import { Formaters } from '../utils/formaters';
	export let trainings: any[] = [];

	$: sorted = trainings.sort((a, b) => {
		return new Date(b.StartDate).getMilliseconds() - new Date(a.StartDate).getMilliseconds();
	});

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
		{#each sorted as training}
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
