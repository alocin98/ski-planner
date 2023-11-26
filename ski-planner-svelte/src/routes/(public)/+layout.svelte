<script>
	import SpaceLayout from '@components/SpaceLayout.svelte';
	import '../../app.postcss';
	import Header from '@components/Header.svelte';
</script>

<Header />
<SpaceLayout>
	<slot />
</SpaceLayout>
