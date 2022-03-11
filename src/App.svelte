<script>
	export let name;
</script>

<svelte:options tag="my-app" />

<main class="text-center">
	<h1 class="text-orange-600">Hello {name}!</h1>
	<p>Visit the <a href="https://svelte.dev/tutorial" class="text-blue-600 underline">Svelte tutorial</a> to learn how to build Svelte apps.</p>
</main>

<style >
	@tailwind base;
	@tailwind components;
	@tailwind utilities;
</style>

