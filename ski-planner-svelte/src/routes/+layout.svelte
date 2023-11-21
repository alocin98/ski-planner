<script>
	import { NavBrand, NavHamburger, NavLi, NavUl, Navbar } from 'flowbite-svelte';
	import '../app.postcss';
</script>

<Navbar>
	<NavBrand href="/">
		<img src="/branding/logo/logo.jpeg" class="mr-3 h-6 sm:h-9" alt="Flowbite Logo" />
		<span class="self-center whitespace-nowrap text-xl font-semibold dark:text-white">Skiyeti</span>
	</NavBrand>
	<NavHamburger />
	<NavUl>
		<NavLi href="/">Home</NavLi>
		<NavLi href="/about">About</NavLi>
		<NavLi href="/contact">Contact</NavLi>
	</NavUl>
</Navbar>
<section class="bg-white dark:bg-gray-900">
	<div class="max-w-screen-xl px-4 py-8 mx-auto">
		<slot />
	</div>
</section>
