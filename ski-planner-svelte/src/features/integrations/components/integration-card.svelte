<script lang="ts">
	export let title: string;
</script>

<div class="card w-96 bg-base-100 shadow-xl">
	<div class="card-body">
		<h3 class="card-title">{title}</h3>
		<slot />
	</div>
</div>
