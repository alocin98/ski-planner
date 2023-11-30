<script>
	import { goto } from '$app/navigation';
	import LoginForm from '@features/auth/components/login-form.svelte';

	const afterLogin = () => {
		goto('/diary');
	};
</script>

<div class="grid lg:gap-8 xl:gap-0 lg:py-16 lg:grid-cols-12">
	<div class="mx-auto w-full place-self-center lg:col-span-7">
		<LoginForm {afterLogin} />
	</div>
	<div class="hidden lg:mt-0 lg:col-span-5 lg:flex">
		<img src="images/hero.jpeg" alt="mockup" />
	</div>
</div>
