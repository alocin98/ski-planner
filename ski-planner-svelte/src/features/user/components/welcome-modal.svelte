<script lang="ts">
	import Modal from '@components/Modal.svelte';
	import Stepper from '@components/Stepper.svelte';
	import Integrations from '@features/integrations/components/integrations.svelte';
	import { StravaService } from '@features/integrations/strava/strava.service';
	import { onMount } from 'svelte';
	import { LocalStorageKey } from '@utils/constants/local-storage';
	import { goto, invalidate, invalidateAll } from '$app/navigation';
	import { UserService } from '@features/auth/service/user-service';

	let currentStep = 1;

	let nextStep = () => {
		currentStep++;
		localStorage.setItem(LocalStorageKey.welcomeStep, currentStep.toString());
	};

	onMount(() => {
		const step = localStorage.getItem(LocalStorageKey.welcomeStep);
		if (step) {
			currentStep = parseInt(step);
		}
	});
	export let user;

	let isConnected = false;

	$: if (isConnected && currentStep == 2) {
		nextStep();
	}

	$: if (currentStep == 3) {
		StravaService.loadTrainingData().then(() => {
			nextStep();
		});
	}

	const finish = () => {
		UserService.finishTutorial().then(() => {
			localStorage.removeItem(LocalStorageKey.welcomeStep);
			invalidateAll();
			goto('/diary');
		});
	};

	const steps = [
		{ title: 'Welcome' },
		{ title: 'Choose provider' },
		{ title: 'Import data' },
		{ title: 'Ski on' }
	];
</script>

<Modal open={true}>
	<div class="flex flex-col items-start">
		{#if currentStep == 1}
			<div class="m-4 mb-8 w-full flex-col flex items-center">
				<h2 class="text-3xl mb-6">Welcome to Skiyeti!</h2>
				<p class="text-gray-600 mb-2">We are excited to have you here. Let's get started by</p>
				<button class="btn btn-primary" on:click={nextStep}>connecting to a provider</button>
			</div>
		{:else if currentStep == 2}
			<div class="m-4 w-full">
				<h2 class="text-3xl mb-6">Choose provider</h2>
				<p>
					You can choose one of the following providers to fetch your trainings. When connected, we
					will <b>fetch your trainings</b>, as long as you are active on the platform. Furthermore,
					we read some of your <b>profile information</b> to fill your accont. The data stays yours
					and won't be exposed. Note that you can<b> only connect to one</b> at the same time.
				</p>
				<div class="w-full flex justify-center">
					<Integrations {user} bind:isConnected />
				</div>
			</div>
		{:else if currentStep == 3}
			<div class="w-full flex flex-col items-center">
				<h1 class="text-3xl mb-6">Loading data</h1>
				<p>
					We are now fetching some of your last trainings to get some data. This might take a while,
					please be patient.
				</p>
				<span class="loading loading-ring loading-lg" />
			</div>
		{:else if currentStep == 4}
			<div class="flex flex-col items-center gap-3 mb-3">
				<h1 class="text-3xl mb-6">Ski on</h1>
				<p>
					You are all set up. You can now start tracking your trainings and get some insights about
					your performance.
				</p>
				<i class="fa-regular fa-circle-check text-4xl" />
				<form method="dialog">
					<buton class="btn btn-primary" on:click={finish}>Let's go</buton>
				</form>
			</div>
		{/if}
	</div>
	<Stepper {currentStep} {steps} />
</Modal>
