<script lang="ts">
	export const openModal = () => {
		modal.showModal();
	};

	let modal: HTMLDialogElement;
	export let open = false;

	$: if (modal && open) {
		modal.showModal();
	}
</script>

<!-- Open the modal using ID.showModal() method -->
<dialog bind:this={modal} class="modal" {...$$restProps}>
	<div class="modal-box max-w-4xl">
		<slot />
	</div>
</dialog>
