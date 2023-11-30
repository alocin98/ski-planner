<script lang="ts">
	import NoEntries from '@features/training-diary/components/no-entries.svelte';
	import TrainingsTable from '@features/training-diary/components/trainings-table.svelte';

	export let data;
</script>

<h1 class="text-3xl mb-6">Diary</h1>
{#if !data.trainings || data.trainings.length === 0}
	<NoEntries />
{:else}
	<TrainingsTable trainings={data.trainings} />
{/if}
