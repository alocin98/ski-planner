<script>
	import { StravaService } from '../strava/strava.service';
	import IntegrationCard from './integration-card.svelte';

	export let user;

	let onClickAuthorizeStrava = () => {
		const redirectUrl = window.location.href;

		console.log(redirectUrl);
		//StravaService.authorize(redirectUrl);
	};

	let isStravaConnected = user.stravaInfo.stravaConnected;
</script>

<IntegrationCard title="strava">
	{#if isStravaConnected}
		<p>You are already connected.</p>
		<button class="btn btn-secondary">Click to disconnect</button>
	{:else}
		<button on:click={onClickAuthorizeStrava}
			><img src="branding/strava/btn-connect-with-strava.png" /></button
		>
	{/if}
</IntegrationCard>
