<script>
	import { goto } from '$app/navigation';
	import RegisterForm from '@features/auth/components/register-form.svelte';

	const afterRegister = () => {
		goto('/diary');
	};
</script>

<div class="grid lg:gap-8 xl:gap-0 lg:py-16 lg:grid-cols-12">
	<div class="mx-auto place-self-center lg:col-span-7">
		<RegisterForm {afterRegister} />
	</div>
	<div class="hidden lg:mt-0 lg:col-span-5 lg:flex">
		<img src="images/hero.jpeg" alt="mockup" />
	</div>
</div>
