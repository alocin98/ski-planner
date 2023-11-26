<script lang="ts">
	import { StravaService } from '@features/strava/strava.service.js';

	export let data;

	let onClickAuthorizeStrava = () => {
		const redirectUrl = window.location.href;
		StravaService.authorize(redirectUrl);
	};
</script>

<h1>Strava</h1>

{#if data.user.stravaConnected}
	<p>Connected to strava</p>
{:else}
	<p>Not connected to strava</p>
	<button class="btn" on:click={onClickAuthorizeStrava}>authorize</button>
{/if}
<p>{JSON.stringify(data.user)}</p>
