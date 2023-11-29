<script lang="ts">
	import Modal from '@components/Modal.svelte';
	import Stepper from '@components/Stepper.svelte';

	let currentStep = 1;
</script>

<Modal open={true}>
	<div class="flex flex-col items-center">
		<h1 class="text-3xl mb-6">Welcome to Skiyeti!</h1>
		<button class="btn btn-primary" on:click={() => currentStep++}>Next</button>
	</div>
	<Stepper {currentStep} steps={[{ title: 'Step 1' }, { title: 'Step 2' }, { title: 'Step 3' }]} />
</Modal>
