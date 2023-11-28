<script>
	import StravaIntegration from '@features/integrations/components/strava-integration.svelte';

	export let data;
</script>

<h1 class="text-2xl mb-6">Settings</h1>
<h2 class="text-xl mb-6">Partner Integrations</h2>

<StravaIntegration user={data.user} />
