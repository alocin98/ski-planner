<script lang="ts">
	import NoEntries from '@features/training-diary/components/no-entries.svelte';
	import { TrainingService } from '@features/training-diary/service/training-service';

	export let data;
</script>

<h1 class="text-3xl mb-6">Diary</h1>
{JSON.stringify(data.trainings)}
<NoEntries />
