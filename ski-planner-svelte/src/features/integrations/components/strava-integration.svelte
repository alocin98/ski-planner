<script>
	import { StravaService } from '../strava/strava.service';
	import IntegrationCard from './integration-card.svelte';

	export let user;

	let onClickAuthorizeStrava = () => {
		const redirectUrl = window.location.href;
		StravaService.authorize(redirectUrl);
	};

	let isStravaConnected = user.stravaConnected;
</script>

<IntegrationCard title="strava">
	<p class="my-4">
		This application is powered by <a class="link" href="https://strava.com">Strava</a>
	</p>

	{#if isStravaConnected}
		<p>You are already connected.</p>
		<button class="btn btn-secondary">Click to disconnect</button>
	{:else}
		<button on:click={onClickAuthorizeStrava}
			><img src="branding/strava/btn-connect-with-strava.png" /></button
		>
	{/if}
</IntegrationCard>
