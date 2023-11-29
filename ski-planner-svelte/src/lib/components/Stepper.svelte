<script lang="ts">
	export let currentStep = 1;
	export let steps: { title: string }[] = [];

	export const nextStep = () => {
		currentStep++;
	};

	let stepList: HTMLElement | null;

	$: if (steps && stepList && currentStep) {
		const steps = stepList.querySelectorAll('li');
		steps.forEach((step, index) => {
			if (index < currentStep) {
				step.classList.add('step-primary');
			} else {
				step.classList.remove('step-primary');
			}
		});
	}
</script>

<slot />
<ul bind:this={stepList} class="steps w-full">
	{#each steps as step}
		<li class="step">{step.title}</li>
	{/each}
</ul>
