<script lang="ts">
	import { StravaService } from '@features/integrations/strava/strava.service';

	let modal: any;

	let loading = false;

	let loadTrainingData = () => {
		loading = true;
		modal.showModal();
		StravaService.loadTrainingData().then(() => {
			loading = false;
			modal.close();
		});
	};
</script>

<div class="card w-96 bg-base-100 shadow-xl">
	<div class="card-body">
		<h3 class="card-title">Strava</h3>
		<p class="my-4">
			This application is powered by <a class="link" href="https://strava.com">Strava</a>
		</p>
		<button class="btn btn-primary" on:click={loadTrainingData}>Von Strava laden</button>
	</div>
</div>

<!-- Open the modal using ID.showModal() method -->
<dialog bind:this={modal} id="import-training-modal" class="modal">
	<div class="modal-box">
		{#if loading}
			<div class="loading loading-lg" />
		{:else}
			<div class="modal-header">
				<h3>Importing training data from Strava</h3>
			</div>
			<div class="modal-body">
				<p>Importing training data from Strava can take a while. Please be patient.</p>
			</div>
		{/if}
	</div>
</dialog>
